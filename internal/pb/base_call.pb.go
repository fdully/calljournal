// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: base_call.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BaseCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Clid string               `protobuf:"bytes,4,opt,name=clid,proto3" json:"clid,omitempty"`
	Clna string               `protobuf:"bytes,7,opt,name=clna,proto3" json:"clna,omitempty"`
	Dest string               `protobuf:"bytes,10,opt,name=dest,proto3" json:"dest,omitempty"`
	Dirc string               `protobuf:"bytes,13,opt,name=dirc,proto3" json:"dirc,omitempty"`
	Stti *timestamp.Timestamp `protobuf:"bytes,16,opt,name=stti,proto3" json:"stti,omitempty"`
	Durs uint32               `protobuf:"varint,19,opt,name=durs,proto3" json:"durs,omitempty"`
	Bils uint32               `protobuf:"varint,22,opt,name=bils,proto3" json:"bils,omitempty"`
	Recd bool                 `protobuf:"varint,25,opt,name=recd,proto3" json:"recd,omitempty"`
	Recs uint32               `protobuf:"varint,28,opt,name=recs,proto3" json:"recs,omitempty"`
	Recl string               `protobuf:"bytes,31,opt,name=recl,proto3" json:"recl,omitempty"`
	Rtag string               `protobuf:"bytes,34,opt,name=rtag,proto3" json:"rtag,omitempty"`
	Epos int64                `protobuf:"varint,37,opt,name=epos,proto3" json:"epos,omitempty"`
	Epoa int64                `protobuf:"varint,40,opt,name=epoa,proto3" json:"epoa,omitempty"`
	Epoe int64                `protobuf:"varint,43,opt,name=epoe,proto3" json:"epoe,omitempty"`
	Wbye string               `protobuf:"bytes,46,opt,name=wbye,proto3" json:"wbye,omitempty"`
	Hang string               `protobuf:"bytes,49,opt,name=hang,proto3" json:"hang,omitempty"`
	Code string               `protobuf:"bytes,52,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BaseCall) Reset() {
	*x = BaseCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCall) ProtoMessage() {}

func (x *BaseCall) ProtoReflect() protoreflect.Message {
	mi := &file_base_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCall.ProtoReflect.Descriptor instead.
func (*BaseCall) Descriptor() ([]byte, []int) {
	return file_base_call_proto_rawDescGZIP(), []int{0}
}

func (x *BaseCall) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BaseCall) GetClid() string {
	if x != nil {
		return x.Clid
	}
	return ""
}

func (x *BaseCall) GetClna() string {
	if x != nil {
		return x.Clna
	}
	return ""
}

func (x *BaseCall) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *BaseCall) GetDirc() string {
	if x != nil {
		return x.Dirc
	}
	return ""
}

func (x *BaseCall) GetStti() *timestamp.Timestamp {
	if x != nil {
		return x.Stti
	}
	return nil
}

func (x *BaseCall) GetDurs() uint32 {
	if x != nil {
		return x.Durs
	}
	return 0
}

func (x *BaseCall) GetBils() uint32 {
	if x != nil {
		return x.Bils
	}
	return 0
}

func (x *BaseCall) GetRecd() bool {
	if x != nil {
		return x.Recd
	}
	return false
}

func (x *BaseCall) GetRecs() uint32 {
	if x != nil {
		return x.Recs
	}
	return 0
}

func (x *BaseCall) GetRecl() string {
	if x != nil {
		return x.Recl
	}
	return ""
}

func (x *BaseCall) GetRtag() string {
	if x != nil {
		return x.Rtag
	}
	return ""
}

func (x *BaseCall) GetEpos() int64 {
	if x != nil {
		return x.Epos
	}
	return 0
}

func (x *BaseCall) GetEpoa() int64 {
	if x != nil {
		return x.Epoa
	}
	return 0
}

func (x *BaseCall) GetEpoe() int64 {
	if x != nil {
		return x.Epoe
	}
	return 0
}

func (x *BaseCall) GetWbye() string {
	if x != nil {
		return x.Wbye
	}
	return ""
}

func (x *BaseCall) GetHang() string {
	if x != nil {
		return x.Hang
	}
	return ""
}

func (x *BaseCall) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SaveBaseCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseCall *BaseCall `protobuf:"bytes,1,opt,name=base_call,json=baseCall,proto3" json:"base_call,omitempty"`
}

func (x *SaveBaseCallRequest) Reset() {
	*x = SaveBaseCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBaseCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBaseCallRequest) ProtoMessage() {}

func (x *SaveBaseCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBaseCallRequest.ProtoReflect.Descriptor instead.
func (*SaveBaseCallRequest) Descriptor() ([]byte, []int) {
	return file_base_call_proto_rawDescGZIP(), []int{1}
}

func (x *SaveBaseCallRequest) GetBaseCall() *BaseCall {
	if x != nil {
		return x.BaseCall
	}
	return nil
}

var File_base_call_proto protoreflect.FileDescriptor

var file_base_call_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8e, 0x03, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6c, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x6e, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6c, 0x6e, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x69, 0x72, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x72, 0x63,
	0x12, 0x2e, 0x0a, 0x04, 0x73, 0x74, 0x74, 0x69, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x73, 0x74, 0x74, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x75, 0x72, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x64, 0x75, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x73, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x62, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x63, 0x64,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x63, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x63, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x65, 0x63, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x63, 0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x65, 0x63, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x74, 0x61, 0x67, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x70, 0x6f, 0x73,
	0x18, 0x25, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x70, 0x6f, 0x61, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x65, 0x70, 0x6f, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x70, 0x6f, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x65, 0x70, 0x6f, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x62, 0x79, 0x65, 0x18, 0x2e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x77, 0x62, 0x79, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x6e, 0x67,
	0x18, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x3d, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x32,
	0x4f, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x73, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_call_proto_rawDescOnce sync.Once
	file_base_call_proto_rawDescData = file_base_call_proto_rawDesc
)

func file_base_call_proto_rawDescGZIP() []byte {
	file_base_call_proto_rawDescOnce.Do(func() {
		file_base_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_call_proto_rawDescData)
	})
	return file_base_call_proto_rawDescData
}

var file_base_call_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_call_proto_goTypes = []interface{}{
	(*BaseCall)(nil),            // 0: BaseCall
	(*SaveBaseCallRequest)(nil), // 1: SaveBaseCallRequest
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_base_call_proto_depIdxs = []int32{
	2, // 0: BaseCall.stti:type_name -> google.protobuf.Timestamp
	0, // 1: SaveBaseCallRequest.base_call:type_name -> BaseCall
	1, // 2: BaseCallService.SaveBaseCall:input_type -> SaveBaseCallRequest
	3, // 3: BaseCallService.SaveBaseCall:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_base_call_proto_init() }
func file_base_call_proto_init() {
	if File_base_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCall); i {
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
		file_base_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBaseCallRequest); i {
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
			RawDescriptor: file_base_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_base_call_proto_goTypes,
		DependencyIndexes: file_base_call_proto_depIdxs,
		MessageInfos:      file_base_call_proto_msgTypes,
	}.Build()
	File_base_call_proto = out.File
	file_base_call_proto_rawDesc = nil
	file_base_call_proto_goTypes = nil
	file_base_call_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BaseCallServiceClient is the client API for BaseCallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaseCallServiceClient interface {
	SaveBaseCall(ctx context.Context, in *SaveBaseCallRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type baseCallServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseCallServiceClient(cc grpc.ClientConnInterface) BaseCallServiceClient {
	return &baseCallServiceClient{cc}
}

func (c *baseCallServiceClient) SaveBaseCall(ctx context.Context, in *SaveBaseCallRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/BaseCallService/SaveBaseCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseCallServiceServer is the server API for BaseCallService service.
type BaseCallServiceServer interface {
	SaveBaseCall(context.Context, *SaveBaseCallRequest) (*empty.Empty, error)
}

// UnimplementedBaseCallServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBaseCallServiceServer struct {
}

func (*UnimplementedBaseCallServiceServer) SaveBaseCall(context.Context, *SaveBaseCallRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBaseCall not implemented")
}

func RegisterBaseCallServiceServer(s *grpc.Server, srv BaseCallServiceServer) {
	s.RegisterService(&_BaseCallService_serviceDesc, srv)
}

func _BaseCallService_SaveBaseCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBaseCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseCallServiceServer).SaveBaseCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BaseCallService/SaveBaseCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseCallServiceServer).SaveBaseCall(ctx, req.(*SaveBaseCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaseCallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BaseCallService",
	HandlerType: (*BaseCallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveBaseCall",
			Handler:    _BaseCallService_SaveBaseCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base_call.proto",
}