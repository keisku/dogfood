// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/v1/healthcheck/healthcheck.proto

package healthcheckpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LivenessProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessProbeRequest) Reset() {
	*x = LivenessProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessProbeRequest) ProtoMessage() {}

func (x *LivenessProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessProbeRequest.ProtoReflect.Descriptor instead.
func (*LivenessProbeRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{0}
}

type LivenessProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessProbeResponse) Reset() {
	*x = LivenessProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessProbeResponse) ProtoMessage() {}

func (x *LivenessProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessProbeResponse.ProtoReflect.Descriptor instead.
func (*LivenessProbeResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{1}
}

type ReadinessProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadinessProbeRequest) Reset() {
	*x = ReadinessProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessProbeRequest) ProtoMessage() {}

func (x *ReadinessProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessProbeRequest.ProtoReflect.Descriptor instead.
func (*ReadinessProbeRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{2}
}

type ReadinessProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadinessProbeResponse) Reset() {
	*x = ReadinessProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadinessProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessProbeResponse) ProtoMessage() {}

func (x *ReadinessProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessProbeResponse.ProtoReflect.Descriptor instead.
func (*ReadinessProbeResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{3}
}

type StartupProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartupProbeRequest) Reset() {
	*x = StartupProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartupProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartupProbeRequest) ProtoMessage() {}

func (x *StartupProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartupProbeRequest.ProtoReflect.Descriptor instead.
func (*StartupProbeRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{4}
}

type StartupProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartupProbeResponse) Reset() {
	*x = StartupProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartupProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartupProbeResponse) ProtoMessage() {}

func (x *StartupProbeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_healthcheck_healthcheck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartupProbeResponse.ProtoReflect.Descriptor instead.
func (*StartupProbeResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP(), []int{5}
}

var File_proto_v1_healthcheck_healthcheck_proto protoreflect.FileDescriptor

var file_proto_v1_healthcheck_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x03, 0x0a, 0x12,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x12, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x8b, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12,
	0x27, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x25, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x42, 0x18, 0x5a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_v1_healthcheck_healthcheck_proto_rawDescOnce sync.Once
	file_proto_v1_healthcheck_healthcheck_proto_rawDescData = file_proto_v1_healthcheck_healthcheck_proto_rawDesc
)

func file_proto_v1_healthcheck_healthcheck_proto_rawDescGZIP() []byte {
	file_proto_v1_healthcheck_healthcheck_proto_rawDescOnce.Do(func() {
		file_proto_v1_healthcheck_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_healthcheck_healthcheck_proto_rawDescData)
	})
	return file_proto_v1_healthcheck_healthcheck_proto_rawDescData
}

var file_proto_v1_healthcheck_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_v1_healthcheck_healthcheck_proto_goTypes = []interface{}{
	(*LivenessProbeRequest)(nil),   // 0: healthcheckpb.v1.LivenessProbeRequest
	(*LivenessProbeResponse)(nil),  // 1: healthcheckpb.v1.LivenessProbeResponse
	(*ReadinessProbeRequest)(nil),  // 2: healthcheckpb.v1.ReadinessProbeRequest
	(*ReadinessProbeResponse)(nil), // 3: healthcheckpb.v1.ReadinessProbeResponse
	(*StartupProbeRequest)(nil),    // 4: healthcheckpb.v1.StartupProbeRequest
	(*StartupProbeResponse)(nil),   // 5: healthcheckpb.v1.StartupProbeResponse
}
var file_proto_v1_healthcheck_healthcheck_proto_depIdxs = []int32{
	0, // 0: healthcheckpb.v1.HealthCheckService.LivenessProbe:input_type -> healthcheckpb.v1.LivenessProbeRequest
	2, // 1: healthcheckpb.v1.HealthCheckService.ReadinessProbe:input_type -> healthcheckpb.v1.ReadinessProbeRequest
	4, // 2: healthcheckpb.v1.HealthCheckService.StartupProbe:input_type -> healthcheckpb.v1.StartupProbeRequest
	1, // 3: healthcheckpb.v1.HealthCheckService.LivenessProbe:output_type -> healthcheckpb.v1.LivenessProbeResponse
	3, // 4: healthcheckpb.v1.HealthCheckService.ReadinessProbe:output_type -> healthcheckpb.v1.ReadinessProbeResponse
	5, // 5: healthcheckpb.v1.HealthCheckService.StartupProbe:output_type -> healthcheckpb.v1.StartupProbeResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_healthcheck_healthcheck_proto_init() }
func file_proto_v1_healthcheck_healthcheck_proto_init() {
	if File_proto_v1_healthcheck_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LivenessProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LivenessProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadinessProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadinessProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartupProbeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1_healthcheck_healthcheck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartupProbeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_healthcheck_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_healthcheck_healthcheck_proto_goTypes,
		DependencyIndexes: file_proto_v1_healthcheck_healthcheck_proto_depIdxs,
		MessageInfos:      file_proto_v1_healthcheck_healthcheck_proto_msgTypes,
	}.Build()
	File_proto_v1_healthcheck_healthcheck_proto = out.File
	file_proto_v1_healthcheck_healthcheck_proto_rawDesc = nil
	file_proto_v1_healthcheck_healthcheck_proto_goTypes = nil
	file_proto_v1_healthcheck_healthcheck_proto_depIdxs = nil
}