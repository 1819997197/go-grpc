// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: hello_http.proto

package hello

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "go-grpc/ch06/proto/google/api"
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

// HelloRequest 请求结构
type HelloHTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloHTTPRequest) Reset() {
	*x = HelloHTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloHTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloHTTPRequest) ProtoMessage() {}

func (x *HelloHTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloHTTPRequest.ProtoReflect.Descriptor instead.
func (*HelloHTTPRequest) Descriptor() ([]byte, []int) {
	return file_hello_http_proto_rawDescGZIP(), []int{0}
}

func (x *HelloHTTPRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// HelloResponse 响应结构
type HelloHTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloHTTPResponse) Reset() {
	*x = HelloHTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloHTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloHTTPResponse) ProtoMessage() {}

func (x *HelloHTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloHTTPResponse.ProtoReflect.Descriptor instead.
func (*HelloHTTPResponse) Descriptor() ([]byte, []int) {
	return file_hello_http_proto_rawDescGZIP(), []int{1}
}

func (x *HelloHTTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hello_http_proto protoreflect.FileDescriptor

var file_hello_http_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2d, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x64,
	0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x12, 0x57, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x22, 0x0d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_http_proto_rawDescOnce sync.Once
	file_hello_http_proto_rawDescData = file_hello_http_proto_rawDesc
)

func file_hello_http_proto_rawDescGZIP() []byte {
	file_hello_http_proto_rawDescOnce.Do(func() {
		file_hello_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_http_proto_rawDescData)
	})
	return file_hello_http_proto_rawDescData
}

var file_hello_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hello_http_proto_goTypes = []interface{}{
	(*HelloHTTPRequest)(nil),  // 0: hello.HelloHTTPRequest
	(*HelloHTTPResponse)(nil), // 1: hello.HelloHTTPResponse
}
var file_hello_http_proto_depIdxs = []int32{
	0, // 0: hello.HelloHTTP.SayHello:input_type -> hello.HelloHTTPRequest
	1, // 1: hello.HelloHTTP.SayHello:output_type -> hello.HelloHTTPResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_http_proto_init() }
func file_hello_http_proto_init() {
	if File_hello_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloHTTPRequest); i {
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
		file_hello_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloHTTPResponse); i {
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
			RawDescriptor: file_hello_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_http_proto_goTypes,
		DependencyIndexes: file_hello_http_proto_depIdxs,
		MessageInfos:      file_hello_http_proto_msgTypes,
	}.Build()
	File_hello_http_proto = out.File
	file_hello_http_proto_rawDesc = nil
	file_hello_http_proto_goTypes = nil
	file_hello_http_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloHTTPClient is the client API for HelloHTTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloHTTPClient interface {
	// 定义SayHello方法
	SayHello(ctx context.Context, in *HelloHTTPRequest, opts ...grpc.CallOption) (*HelloHTTPResponse, error)
}

type helloHTTPClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloHTTPClient(cc grpc.ClientConnInterface) HelloHTTPClient {
	return &helloHTTPClient{cc}
}

func (c *helloHTTPClient) SayHello(ctx context.Context, in *HelloHTTPRequest, opts ...grpc.CallOption) (*HelloHTTPResponse, error) {
	out := new(HelloHTTPResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloHTTP/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloHTTPServer is the server API for HelloHTTP service.
type HelloHTTPServer interface {
	// 定义SayHello方法
	SayHello(context.Context, *HelloHTTPRequest) (*HelloHTTPResponse, error)
}

// UnimplementedHelloHTTPServer can be embedded to have forward compatible implementations.
type UnimplementedHelloHTTPServer struct {
}

func (*UnimplementedHelloHTTPServer) SayHello(context.Context, *HelloHTTPRequest) (*HelloHTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloHTTPServer(s *grpc.Server, srv HelloHTTPServer) {
	s.RegisterService(&_HelloHTTP_serviceDesc, srv)
}

func _HelloHTTP_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloHTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloHTTPServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloHTTP/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloHTTPServer).SayHello(ctx, req.(*HelloHTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloHTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloHTTP",
	HandlerType: (*HelloHTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloHTTP_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_http.proto",
}
