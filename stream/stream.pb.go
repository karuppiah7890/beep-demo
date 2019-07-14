// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package stream

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

type PositionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionRequest) Reset()         { *m = PositionRequest{} }
func (m *PositionRequest) String() string { return proto.CompactTextString(m) }
func (*PositionRequest) ProtoMessage()    {}
func (*PositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *PositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionRequest.Unmarshal(m, b)
}
func (m *PositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionRequest.Marshal(b, m, deterministic)
}
func (m *PositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionRequest.Merge(m, src)
}
func (m *PositionRequest) XXX_Size() int {
	return xxx_messageInfo_PositionRequest.Size(m)
}
func (m *PositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PositionRequest proto.InternalMessageInfo

type PositionResponse struct {
	Position             int64    `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionResponse) Reset()         { *m = PositionResponse{} }
func (m *PositionResponse) String() string { return proto.CompactTextString(m) }
func (*PositionResponse) ProtoMessage()    {}
func (*PositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *PositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionResponse.Unmarshal(m, b)
}
func (m *PositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionResponse.Marshal(b, m, deterministic)
}
func (m *PositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionResponse.Merge(m, src)
}
func (m *PositionResponse) XXX_Size() int {
	return xxx_messageInfo_PositionResponse.Size(m)
}
func (m *PositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PositionResponse proto.InternalMessageInfo

func (m *PositionResponse) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func init() {
	proto.RegisterType((*PositionRequest)(nil), "stream.PositionRequest")
	proto.RegisterType((*PositionResponse)(nil), "stream.PositionResponse")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x04, 0xb9, 0xf8,
	0x03, 0xf2, 0x8b, 0x33, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94,
	0xf4, 0xb8, 0x04, 0x10, 0x42, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x05,
	0x50, 0x31, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x38, 0xdf, 0xc8, 0x93, 0x8b, 0x2d, 0x18,
	0x6c, 0x98, 0x90, 0x3d, 0x17, 0x07, 0x4c, 0xa7, 0x90, 0xb8, 0x1e, 0xd4, 0x3e, 0x34, 0xe3, 0xa5,
	0x24, 0x30, 0x25, 0x20, 0x96, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x67, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xe4, 0x80, 0x7c, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error) {
	out := new(PositionResponse)
	err := c.cc.Invoke(ctx, "/stream.Stream/Position", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Position(context.Context, *PositionRequest) (*PositionResponse, error)
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Position(ctx context.Context, req *PositionRequest) (*PositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.Stream/Position",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).Position(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Position",
			Handler:    _Stream_Position_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream.proto",
}