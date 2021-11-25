package grpc_backend

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kei6u/dogfood/pkg/ddconfig"
	dogfoodpb "github.com/kei6u/dogfood/proto/v1/dogfood"
	healthcheckpb "github.com/kei6u/dogfood/proto/v1/healthcheck"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	grpc_dd "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	http_dd "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var _ dogfoodpb.DogFoodServiceServer = (*Server)(nil)
var _ healthcheckpb.HealthCheckServiceServer = (*Server)(nil)

type Server struct {
	// gRPCAddr is the address where gRPC Server listens to.
	gRPCAddr string
	// gRPCGWAddr is the address where gRPC-Gateway Server listens to.
	gRPCGWAddr string
	logger     *zap.Logger
	db         *sql.DB
}

func NewServer(gRPCAddr, gRPCGWAddr string, logger *zap.Logger, db *sql.DB) *Server {
	if !strings.HasPrefix(gRPCAddr, ":") {
		gRPCAddr = fmt.Sprintf(":%s", gRPCAddr)
	}
	if !strings.HasPrefix(gRPCGWAddr, ":") {
		gRPCGWAddr = fmt.Sprintf(":%s", gRPCGWAddr)
	}
	return &Server{
		gRPCAddr:   gRPCAddr,
		gRPCGWAddr: gRPCGWAddr,
		logger:     logger,
		db:         db,
	}
}

func (s *Server) Start(ctx context.Context) {
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := prometheus.NewRegistry()
		if err := r.Register(grpcMetrics); err != nil {
			s.logger.Warn("failed to register gRPC metrics to prometheus", zap.Error(err))
			return
		}
		if err := r.Register(dogfoodGramGuage); err != nil {
			s.logger.Warn("failed to register dogfood gram guage to prometheus", zap.Error(err))
			return
		}
		if err := r.Register(dogfoodNameCount); err != nil {
			s.logger.Warn("failed to register dogfood name count to prometheus", zap.Error(err))
			return
		}
		if err := (&http.Server{
			Handler: promhttp.HandlerFor(r, promhttp.HandlerOpts{}),
			Addr:    ":9092",
		}).ListenAndServe(); err != nil {
			s.logger.Warn("prometheus Server fails to start", zap.Error(err))
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", s.gRPCAddr)
		if err != nil {
			s.logger.Warn("faild to listen to gRPC Address", zap.Error(err), zap.String("address", s.gRPCAddr))
			return
		}
		defer lis.Close()

		ignoreMethods := make([]string, len(healthcheckpb.HealthCheckService_ServiceDesc.Methods))
		for i, m := range healthcheckpb.HealthCheckService_ServiceDesc.Methods {
			ignoreMethods[i] = fmt.Sprintf(
				"/%s/%s",
				healthcheckpb.HealthCheckService_ServiceDesc.ServiceName,
				m.MethodName,
			)
		}

		grpcsvc := grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_dd.UnaryServerInterceptor(
					grpc_dd.WithAnalytics(true),
					grpc_dd.WithIgnoredMethods(ignoreMethods...),
				),
				metricsUnaryServerInterceptor(),
				grpcMetrics.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(
					s.logger,
					grpc_zap.WithDecider(func(fullMethodName string, _ error) bool {
						return !strings.Contains(fullMethodName, "healthcheck")
					}),
					grpc_zap.WithMessageProducer(func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
						// inject trace ID into logs
						if dds, ok := tracer.SpanFromContext(ctx); ok {
							grpc_zap.AddFields(
								ctx,
								zap.Uint64("dd.trace_id", dds.Context().TraceID()),
								zap.Uint64("dd.span_id", dds.Context().SpanID()),
							)
						}
						grpc_zap.DefaultMessageProducer(ctx, msg, level, code, err, duration)
					}),
				),
			),
		)
		dogfoodpb.RegisterDogFoodServiceServer(grpcsvc, s)
		healthcheckpb.RegisterHealthCheckServiceServer(grpcsvc, s)
		grpcMetrics.InitializeMetrics(grpcsvc)

		s.logger.Info("gRPC Server starts", zap.String("address", s.gRPCAddr))
		if err := grpcsvc.Serve(lis); err != nil {
			s.logger.Warn("gRPC Server fails to start", zap.Error(err), zap.String("address", s.gRPCAddr))
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := grpc.DialContext(
			ctx,
			s.gRPCAddr,
			grpc.WithBlock(),
			grpc.WithInsecure(),
			grpc.WithDisableHealthCheck(),
		)
		if err != nil {
			s.logger.Warn("failed to dial gRPC Server", zap.Error(err))
			return
		}
		defer conn.Close()

		gwmux := runtime.NewServeMux(
			runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
				return metadata.New(map[string]string{
					tracer.DefaultTraceIDHeader:  r.Header.Get(tracer.DefaultTraceIDHeader),
					tracer.DefaultParentIDHeader: r.Header.Get(tracer.DefaultParentIDHeader),
					tracer.DefaultPriorityHeader: r.Header.Get(tracer.DefaultPriorityHeader),
				})
			}),
		)
		if err := dogfoodpb.RegisterDogFoodServiceHandler(ctx, gwmux, conn); err != nil {
			s.logger.Warn("failed to regiser handler", zap.Error(err))
			return
		}
		if err := healthcheckpb.RegisterHealthCheckServiceHandler(ctx, gwmux, conn); err != nil {
			s.logger.Warn("failed to regiser handler", zap.Error(err))
			return
		}
		s.logger.Info("gRPC-Gateway Server starts", zap.String("address", s.gRPCGWAddr))
		if err := (&http.Server{
			Addr: s.gRPCGWAddr,
			Handler: http_dd.WrapHandler(
				gwmux,
				ddconfig.GetService(ddconfig.WithServiceSuffix("grpcgateway")),
				"",
				http_dd.WithAnalytics(true),
				http_dd.WithIgnoreRequest(func(r *http.Request) bool {
					return strings.Contains(strings.ToLower(r.RequestURI), "healthcheck")
				}),
			),
		}).ListenAndServe(); err != nil {
			s.logger.Warn("gRPC-Gateway fails to start", zap.Error(err))
			return
		}
	}()

	wg.Wait()
}
