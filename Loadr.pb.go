// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Loadr.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ControlRequest struct {
	Spec                 *TestSpecification `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ControlRequest) Reset()         { *m = ControlRequest{} }
func (m *ControlRequest) String() string { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()    {}
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{0}
}

func (m *ControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlRequest.Unmarshal(m, b)
}
func (m *ControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlRequest.Marshal(b, m, deterministic)
}
func (m *ControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRequest.Merge(m, src)
}
func (m *ControlRequest) XXX_Size() int {
	return xxx_messageInfo_ControlRequest.Size(m)
}
func (m *ControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRequest proto.InternalMessageInfo

func (m *ControlRequest) GetSpec() *TestSpecification {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ControlResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlResponse) Reset()         { *m = ControlResponse{} }
func (m *ControlResponse) String() string { return proto.CompactTextString(m) }
func (*ControlResponse) ProtoMessage()    {}
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{1}
}

func (m *ControlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlResponse.Unmarshal(m, b)
}
func (m *ControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlResponse.Marshal(b, m, deterministic)
}
func (m *ControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlResponse.Merge(m, src)
}
func (m *ControlResponse) XXX_Size() int {
	return xxx_messageInfo_ControlResponse.Size(m)
}
func (m *ControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlResponse proto.InternalMessageInfo

type TestSpecification struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestSpecification) Reset()         { *m = TestSpecification{} }
func (m *TestSpecification) String() string { return proto.CompactTextString(m) }
func (*TestSpecification) ProtoMessage()    {}
func (*TestSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{2}
}

func (m *TestSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestSpecification.Unmarshal(m, b)
}
func (m *TestSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestSpecification.Marshal(b, m, deterministic)
}
func (m *TestSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestSpecification.Merge(m, src)
}
func (m *TestSpecification) XXX_Size() int {
	return xxx_messageInfo_TestSpecification.Size(m)
}
func (m *TestSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_TestSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_TestSpecification proto.InternalMessageInfo

type SendTestResultRequest struct {
	Result               *TestResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendTestResultRequest) Reset()         { *m = SendTestResultRequest{} }
func (m *SendTestResultRequest) String() string { return proto.CompactTextString(m) }
func (*SendTestResultRequest) ProtoMessage()    {}
func (*SendTestResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{3}
}

func (m *SendTestResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTestResultRequest.Unmarshal(m, b)
}
func (m *SendTestResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTestResultRequest.Marshal(b, m, deterministic)
}
func (m *SendTestResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTestResultRequest.Merge(m, src)
}
func (m *SendTestResultRequest) XXX_Size() int {
	return xxx_messageInfo_SendTestResultRequest.Size(m)
}
func (m *SendTestResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTestResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendTestResultRequest proto.InternalMessageInfo

func (m *SendTestResultRequest) GetResult() *TestResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type SendTestResultResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendTestResultResponse) Reset()         { *m = SendTestResultResponse{} }
func (m *SendTestResultResponse) String() string { return proto.CompactTextString(m) }
func (*SendTestResultResponse) ProtoMessage()    {}
func (*SendTestResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{4}
}

func (m *SendTestResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTestResultResponse.Unmarshal(m, b)
}
func (m *SendTestResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTestResultResponse.Marshal(b, m, deterministic)
}
func (m *SendTestResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTestResultResponse.Merge(m, src)
}
func (m *SendTestResultResponse) XXX_Size() int {
	return xxx_messageInfo_SendTestResultResponse.Size(m)
}
func (m *SendTestResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTestResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendTestResultResponse proto.InternalMessageInfo

type TestResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_121cb0d4cb144d55, []int{5}
}

func (m *TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResult.Unmarshal(m, b)
}
func (m *TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResult.Marshal(b, m, deterministic)
}
func (m *TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResult.Merge(m, src)
}
func (m *TestResult) XXX_Size() int {
	return xxx_messageInfo_TestResult.Size(m)
}
func (m *TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ControlRequest)(nil), "main.ControlRequest")
	proto.RegisterType((*ControlResponse)(nil), "main.ControlResponse")
	proto.RegisterType((*TestSpecification)(nil), "main.TestSpecification")
	proto.RegisterType((*SendTestResultRequest)(nil), "main.SendTestResultRequest")
	proto.RegisterType((*SendTestResultResponse)(nil), "main.SendTestResultResponse")
	proto.RegisterType((*TestResult)(nil), "main.TestResult")
}

func init() { proto.RegisterFile("Loadr.proto", fileDescriptor_121cb0d4cb144d55) }

var fileDescriptor_121cb0d4cb144d55 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x09, 0x2c, 0x2b, 0xcc, 0x2e, 0xab, 0x1b, 0x5d, 0x2d, 0xd5, 0x83, 0xe4, 0x54, 0x10,
	0x8a, 0xd4, 0x9b, 0xe0, 0x41, 0x7a, 0xd4, 0x53, 0xeb, 0x0b, 0xd4, 0x74, 0x84, 0x40, 0xcd, 0xc4,
	0x24, 0xf5, 0x55, 0x7c, 0x5d, 0x69, 0x1a, 0x2d, 0xd6, 0x5e, 0xff, 0xff, 0xcb, 0x7c, 0x33, 0x81,
	0xcd, 0x33, 0x35, 0xad, 0xcd, 0x8d, 0x25, 0x4f, 0x7c, 0xf5, 0xde, 0x28, 0x2d, 0x1e, 0x60, 0x57,
	0x92, 0xf6, 0x96, 0xba, 0x0a, 0x3f, 0x7a, 0x74, 0x9e, 0xdf, 0xc0, 0xca, 0x19, 0x94, 0x09, 0xbb,
	0x66, 0xd9, 0xa6, 0xb8, 0xc8, 0x07, 0x2c, 0x7f, 0x41, 0xe7, 0x6b, 0x83, 0x52, 0xbd, 0x29, 0xd9,
	0x78, 0x45, 0xba, 0x0a, 0x90, 0xd8, 0xc3, 0xf1, 0xef, 0x73, 0x67, 0x48, 0x3b, 0x14, 0xa7, 0xb0,
	0xff, 0x47, 0x8b, 0x47, 0x38, 0xd4, 0xa8, 0xdb, 0xa1, 0xa8, 0xd0, 0xf5, 0x9d, 0xff, 0xb1, 0x65,
	0xb0, 0xb6, 0x21, 0x88, 0xbe, 0x93, 0xc9, 0x17, 0xc1, 0xd8, 0x8b, 0x04, 0xce, 0xe7, 0x23, 0xa2,
	0x71, 0x0b, 0x30, 0xa5, 0xc5, 0x17, 0x83, 0x6d, 0xb8, 0xb3, 0x46, 0xfb, 0xa9, 0x24, 0xf2, 0x7b,
	0x38, 0x2a, 0x49, 0x6b, 0x94, 0x9e, 0x1f, 0xc6, 0xe9, 0xb3, 0x95, 0xd3, 0xb3, 0x59, 0x1c, 0x56,
	0xcb, 0xd8, 0x2d, 0xe3, 0x4f, 0xb0, 0xfb, 0x2b, 0xe5, 0x97, 0x23, 0xbb, 0x78, 0x4d, 0x7a, 0xb5,
	0x5c, 0x8e, 0x9a, 0xd7, 0x75, 0xf8, 0xf8, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x0b,
	0x1d, 0x69, 0x87, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadrServiceClient is the client API for LoadrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadrServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (LoadrService_ConnectClient, error)
	SendTestResult(ctx context.Context, in *SendTestResultRequest, opts ...grpc.CallOption) (*SendTestResultResponse, error)
}

type loadrServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadrServiceClient(cc *grpc.ClientConn) LoadrServiceClient {
	return &loadrServiceClient{cc}
}

func (c *loadrServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (LoadrService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadrService_serviceDesc.Streams[0], "/main.LoadrService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadrServiceConnectClient{stream}
	return x, nil
}

type LoadrService_ConnectClient interface {
	Send(*ControlResponse) error
	Recv() (*ControlRequest, error)
	grpc.ClientStream
}

type loadrServiceConnectClient struct {
	grpc.ClientStream
}

func (x *loadrServiceConnectClient) Send(m *ControlResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadrServiceConnectClient) Recv() (*ControlRequest, error) {
	m := new(ControlRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loadrServiceClient) SendTestResult(ctx context.Context, in *SendTestResultRequest, opts ...grpc.CallOption) (*SendTestResultResponse, error) {
	out := new(SendTestResultResponse)
	err := c.cc.Invoke(ctx, "/main.LoadrService/SendTestResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadrServiceServer is the server API for LoadrService service.
type LoadrServiceServer interface {
	Connect(LoadrService_ConnectServer) error
	SendTestResult(context.Context, *SendTestResultRequest) (*SendTestResultResponse, error)
}

// UnimplementedLoadrServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadrServiceServer struct {
}

func (*UnimplementedLoadrServiceServer) Connect(srv LoadrService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedLoadrServiceServer) SendTestResult(ctx context.Context, req *SendTestResultRequest) (*SendTestResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestResult not implemented")
}

func RegisterLoadrServiceServer(s *grpc.Server, srv LoadrServiceServer) {
	s.RegisterService(&_LoadrService_serviceDesc, srv)
}

func _LoadrService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadrServiceServer).Connect(&loadrServiceConnectServer{stream})
}

type LoadrService_ConnectServer interface {
	Send(*ControlRequest) error
	Recv() (*ControlResponse, error)
	grpc.ServerStream
}

type loadrServiceConnectServer struct {
	grpc.ServerStream
}

func (x *loadrServiceConnectServer) Send(m *ControlRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadrServiceConnectServer) Recv() (*ControlResponse, error) {
	m := new(ControlResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LoadrService_SendTestResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTestResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadrServiceServer).SendTestResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.LoadrService/SendTestResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadrServiceServer).SendTestResult(ctx, req.(*SendTestResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.LoadrService",
	HandlerType: (*LoadrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTestResult",
			Handler:    _LoadrService_SendTestResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _LoadrService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Loadr.proto",
}
