// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

type MessageReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReq) Reset()         { *m = MessageReq{} }
func (m *MessageReq) String() string { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()    {}
func (*MessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReq.Unmarshal(m, b)
}
func (m *MessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReq.Marshal(b, m, deterministic)
}
func (m *MessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReq.Merge(m, src)
}
func (m *MessageReq) XXX_Size() int {
	return xxx_messageInfo_MessageReq.Size(m)
}
func (m *MessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReq proto.InternalMessageInfo

func (m *MessageReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MessageResp struct {
	Greet                string   `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResp) Reset()         { *m = MessageResp{} }
func (m *MessageResp) String() string { return proto.CompactTextString(m) }
func (*MessageResp) ProtoMessage()    {}
func (*MessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *MessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResp.Unmarshal(m, b)
}
func (m *MessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResp.Marshal(b, m, deterministic)
}
func (m *MessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResp.Merge(m, src)
}
func (m *MessageResp) XXX_Size() int {
	return xxx_messageInfo_MessageResp.Size(m)
}
func (m *MessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResp proto.InternalMessageInfo

func (m *MessageResp) GetGreet() string {
	if m != nil {
		return m.Greet
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageReq)(nil), "message.MessageReq")
	proto.RegisterType((*MessageResp)(nil), "message.MessageResp")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x14, 0xb8,
	0xb8, 0x7c, 0x21, 0xcc, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x99, 0x8b, 0x1b, 0xae, 0xa2, 0xb8, 0x40,
	0x48, 0x84, 0x8b, 0x35, 0xbd, 0x28, 0x35, 0xb5, 0x04, 0xaa, 0x06, 0xc2, 0x31, 0x72, 0x85, 0x1b,
	0x13, 0x5c, 0x56, 0x24, 0x64, 0xce, 0xc5, 0x95, 0x9e, 0x5a, 0x02, 0x15, 0x10, 0x12, 0xd6, 0x83,
	0xd9, 0x8d, 0xb0, 0x49, 0x4a, 0x04, 0x53, 0xb0, 0xb8, 0x20, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x41, 0xf9, 0xd3, 0xae, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageSvrClient is the client API for MessageSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageSvrClient interface {
	GetMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageResp, error)
}

type messageSvrClient struct {
	cc *grpc.ClientConn
}

func NewMessageSvrClient(cc *grpc.ClientConn) MessageSvrClient {
	return &messageSvrClient{cc}
}

func (c *messageSvrClient) GetMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageResp, error) {
	out := new(MessageResp)
	err := c.cc.Invoke(ctx, "/message.MessageSvr/getMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageSvrServer is the server API for MessageSvr service.
type MessageSvrServer interface {
	GetMessage(context.Context, *MessageReq) (*MessageResp, error)
}

// UnimplementedMessageSvrServer can be embedded to have forward compatible implementations.
type UnimplementedMessageSvrServer struct {
}

func (*UnimplementedMessageSvrServer) GetMessage(ctx context.Context, req *MessageReq) (*MessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func RegisterMessageSvrServer(s *grpc.Server, srv MessageSvrServer) {
	s.RegisterService(&_MessageSvr_serviceDesc, srv)
}

func _MessageSvr_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvrServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageSvr/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvrServer).GetMessage(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageSvr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageSvr",
	HandlerType: (*MessageSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMessage",
			Handler:    _MessageSvr_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
