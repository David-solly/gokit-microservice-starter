// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro_service.proto

package v1

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

// Custom business logic response type here
// will be passed via Service Response through grpc
// response path
// Edit fields to hearts content
type BusinessResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusinessResponse) Reset()         { *m = BusinessResponse{} }
func (m *BusinessResponse) String() string { return proto.CompactTextString(m) }
func (*BusinessResponse) ProtoMessage()    {}
func (*BusinessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ca88c0b3d1b4a2, []int{0}
}

func (m *BusinessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessResponse.Unmarshal(m, b)
}
func (m *BusinessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessResponse.Marshal(b, m, deterministic)
}
func (m *BusinessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessResponse.Merge(m, src)
}
func (m *BusinessResponse) XXX_Size() int {
	return xxx_messageInfo_BusinessResponse.Size(m)
}
func (m *BusinessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessResponse proto.InternalMessageInfo

func (m *BusinessResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Request sent to the grpc server from any axternal client
// Edit fields only to avoid integration failure
// Edit fields to hearts content
type ServiceRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceRequest) Reset()         { *m = ServiceRequest{} }
func (m *ServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceRequest) ProtoMessage()    {}
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ca88c0b3d1b4a2, []int{1}
}

func (m *ServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRequest.Unmarshal(m, b)
}
func (m *ServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRequest.Marshal(b, m, deterministic)
}
func (m *ServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRequest.Merge(m, src)
}
func (m *ServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceRequest.Size(m)
}
func (m *ServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRequest proto.InternalMessageInfo

func (m *ServiceRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

// Returned by the server-
// Template object. WARNING - changing this will require
// changing the corresponding endpoint and transport files.
// Its Recommended to leave this as is unless you are confident.
type ServiceResponse struct {
	Info                 *BusinessResponse `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Error                *ServiceError     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceResponse) Reset()         { *m = ServiceResponse{} }
func (m *ServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceResponse) ProtoMessage()    {}
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ca88c0b3d1b4a2, []int{2}
}

func (m *ServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceResponse.Unmarshal(m, b)
}
func (m *ServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceResponse.Marshal(b, m, deterministic)
}
func (m *ServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceResponse.Merge(m, src)
}
func (m *ServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceResponse.Size(m)
}
func (m *ServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceResponse proto.InternalMessageInfo

func (m *ServiceResponse) GetInfo() *BusinessResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ServiceResponse) GetError() *ServiceError {
	if m != nil {
		return m.Error
	}
	return nil
}

// Template object. WARNING - Leave alone if possible unless you are
// prepared to edit multiple files to regain cohesion.
type ServiceError struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceError) Reset()         { *m = ServiceError{} }
func (m *ServiceError) String() string { return proto.CompactTextString(m) }
func (*ServiceError) ProtoMessage()    {}
func (*ServiceError) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ca88c0b3d1b4a2, []int{3}
}

func (m *ServiceError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceError.Unmarshal(m, b)
}
func (m *ServiceError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceError.Marshal(b, m, deterministic)
}
func (m *ServiceError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceError.Merge(m, src)
}
func (m *ServiceError) XXX_Size() int {
	return xxx_messageInfo_ServiceError.Size(m)
}
func (m *ServiceError) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceError.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceError proto.InternalMessageInfo

func (m *ServiceError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ServiceError) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*BusinessResponse)(nil), "v1.BusinessResponse")
	proto.RegisterType((*ServiceRequest)(nil), "v1.ServiceRequest")
	proto.RegisterType((*ServiceResponse)(nil), "v1.ServiceResponse")
	proto.RegisterType((*ServiceError)(nil), "v1.ServiceError")
}

func init() {
	proto.RegisterFile("micro_service.proto", fileDescriptor_07ca88c0b3d1b4a2)
}

var fileDescriptor_07ca88c0b3d1b4a2 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0x86, 0x40,
	0x10, 0xc6, 0x79, 0x5f, 0x5e, 0x8b, 0x46, 0x29, 0x19, 0x3d, 0x48, 0xa7, 0xd8, 0x83, 0x48, 0x07,
	0x41, 0x83, 0xe8, 0x5c, 0xd4, 0xad, 0xcb, 0xf6, 0x01, 0xc2, 0xd6, 0x09, 0x16, 0xcc, 0xb5, 0x5d,
	0xf5, 0xf3, 0xc7, 0xfe, 0x31, 0xac, 0xdb, 0xcc, 0xb3, 0xbf, 0xe7, 0xd9, 0x99, 0x81, 0xec, 0x4b,
	0x0a, 0xad, 0xde, 0x0d, 0xe9, 0x55, 0x0a, 0xaa, 0x27, 0xad, 0x66, 0x85, 0xc7, 0xb5, 0x61, 0x25,
	0xa4, 0x8f, 0x8b, 0x91, 0x23, 0x19, 0xc3, 0xc9, 0x4c, 0x6a, 0x34, 0x84, 0x08, 0xa7, 0xbe, 0x9b,
	0xbb, 0xe2, 0x70, 0x73, 0xa8, 0x2e, 0xb8, 0xab, 0xd9, 0x2d, 0x5c, 0xbe, 0x79, 0x33, 0xa7, 0xef,
	0x85, 0xcc, 0x8c, 0x05, 0x9c, 0x6b, 0x5f, 0x06, 0x70, 0x6b, 0x99, 0x80, 0xab, 0x5f, 0x36, 0x44,
	0x56, 0x70, 0x92, 0xe3, 0xa7, 0x72, 0x64, 0xdc, 0xe6, 0xf5, 0xda, 0xd4, 0xff, 0xbf, 0xe5, 0x8e,
	0xc0, 0x12, 0x22, 0xd2, 0x5a, 0xe9, 0xe2, 0xe8, 0xd0, 0xd4, 0xa2, 0x21, 0xed, 0xd9, 0xea, 0xdc,
	0x3f, 0xb3, 0x07, 0x48, 0xf6, 0x32, 0xe6, 0x9b, 0xcf, 0x0f, 0xe3, 0x1b, 0xbb, 0x8a, 0x50, 0x3d,
	0xb9, 0xb0, 0x88, 0xbb, 0xba, 0x7d, 0x81, 0xe4, 0xd5, 0x5e, 0x23, 0xd8, 0xf1, 0x1e, 0xe2, 0xa7,
	0x6e, 0x18, 0xb6, 0x16, 0x77, 0x3f, 0x86, 0x5d, 0xaf, 0xb3, 0x3f, 0x9a, 0x9f, 0xf7, 0xe3, 0xcc,
	0x5d, 0xf1, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x07, 0xd0, 0xf9, 0x82, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MicroServiceClient is the client API for MicroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MicroServiceClient interface {
	CallService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type microServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroServiceClient(cc grpc.ClientConnInterface) MicroServiceClient {
	return &microServiceClient{cc}
}

func (c *microServiceClient) CallService(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/v1.MicroService/CallService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroServiceServer is the server API for MicroService service.
type MicroServiceServer interface {
	CallService(context.Context, *ServiceRequest) (*ServiceResponse, error)
}

// UnimplementedMicroServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMicroServiceServer struct {
}

func (*UnimplementedMicroServiceServer) CallService(ctx context.Context, req *ServiceRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallService not implemented")
}

func RegisterMicroServiceServer(s *grpc.Server, srv MicroServiceServer) {
	s.RegisterService(&_MicroService_serviceDesc, srv)
}

func _MicroService_CallService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).CallService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MicroService/CallService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).CallService(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MicroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MicroService",
	HandlerType: (*MicroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallService",
			Handler:    _MicroService_CallService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro_service.proto",
}