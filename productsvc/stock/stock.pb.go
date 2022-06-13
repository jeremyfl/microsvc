// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: grpc/stock.proto

package stock

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body int32 `protobuf:"varint,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_grpc_stock_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() int32 {
	if x != nil {
		return x.Body
	}
	return 0
}

var File_grpc_stock_proto protoreflect.FileDescriptor

var file_grpc_stock_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x7a, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_stock_proto_rawDescOnce sync.Once
	file_grpc_stock_proto_rawDescData = file_grpc_stock_proto_rawDesc
)

func file_grpc_stock_proto_rawDescGZIP() []byte {
	file_grpc_stock_proto_rawDescOnce.Do(func() {
		file_grpc_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_stock_proto_rawDescData)
	})
	return file_grpc_stock_proto_rawDescData
}

var file_grpc_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_stock_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: stock.Message
}
var file_grpc_stock_proto_depIdxs = []int32{
	0, // 0: stock.StockService.FindStockByProduct:input_type -> stock.Message
	0, // 1: stock.StockService.FindEmptyStock:input_type -> stock.Message
	0, // 2: stock.StockService.FindStockByProduct:output_type -> stock.Message
	0, // 3: stock.StockService.FindEmptyStock:output_type -> stock.Message
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_stock_proto_init() }
func file_grpc_stock_proto_init() {
	if File_grpc_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_grpc_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_stock_proto_goTypes,
		DependencyIndexes: file_grpc_stock_proto_depIdxs,
		MessageInfos:      file_grpc_stock_proto_msgTypes,
	}.Build()
	File_grpc_stock_proto = out.File
	file_grpc_stock_proto_rawDesc = nil
	file_grpc_stock_proto_goTypes = nil
	file_grpc_stock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockServiceClient interface {
	FindStockByProduct(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	FindEmptyStock(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) FindStockByProduct(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/stock.StockService/FindStockByProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) FindEmptyStock(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/stock.StockService/FindEmptyStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
type StockServiceServer interface {
	FindStockByProduct(context.Context, *Message) (*Message, error)
	FindEmptyStock(context.Context, *Message) (*Message, error)
}

// UnimplementedStockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (*UnimplementedStockServiceServer) FindStockByProduct(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStockByProduct not implemented")
}
func (*UnimplementedStockServiceServer) FindEmptyStock(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEmptyStock not implemented")
}

func RegisterStockServiceServer(s *grpc.Server, srv StockServiceServer) {
	s.RegisterService(&_StockService_serviceDesc, srv)
}

func _StockService_FindStockByProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).FindStockByProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.StockService/FindStockByProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).FindStockByProduct(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_FindEmptyStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).FindEmptyStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.StockService/FindEmptyStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).FindEmptyStock(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stock.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindStockByProduct",
			Handler:    _StockService_FindStockByProduct_Handler,
		},
		{
			MethodName: "FindEmptyStock",
			Handler:    _StockService_FindEmptyStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/stock.proto",
}