// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recharge_protocol.proto

package recharge_protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	InviteCode           string   `protobuf:"bytes,3,opt,name=invite_code,json=inviteCode,proto3" json:"invite_code,omitempty"`
	Passwd               string   `protobuf:"bytes,4,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recharge_protocol_3c8ce8dc04521942, []int{0}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetInviteCode() string {
	if m != nil {
		return m.InviteCode
	}
	return ""
}

func (m *RegisterRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type RegisterReply struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_recharge_protocol_3c8ce8dc04521942, []int{1}
}
func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (dst *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(dst, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *RegisterReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "recharge_protocol.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "recharge_protocol.RegisterReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/recharge_protocol.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recharge_protocol.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recharge_protocol.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recharge_protocol.proto",
}

func init() {
	proto.RegisterFile("recharge_protocol.proto", fileDescriptor_recharge_protocol_3c8ce8dc04521942)
}

var fileDescriptor_recharge_protocol_3c8ce8dc04521942 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x69, 0x5a, 0x4b, 0xfb, 0xc4, 0xaa, 0xa1, 0xe8, 0x50, 0x04, 0x4b, 0x56, 0xe2, 0xa2,
	0x03, 0xba, 0x73, 0x27, 0xde, 0x60, 0xc0, 0x75, 0x89, 0x93, 0x67, 0x1a, 0x9a, 0xe6, 0xc5, 0x24,
	0x56, 0xba, 0xf5, 0x0a, 0x1e, 0xcd, 0x2b, 0x78, 0x10, 0x99, 0xcc, 0x88, 0x60, 0xc1, 0xdd, 0xfb,
	0xbf, 0xff, 0x25, 0x1f, 0x09, 0x9c, 0x07, 0xac, 0x57, 0x32, 0x68, 0x5c, 0xfa, 0x40, 0x89, 0x6a,
	0xb2, 0x8b, 0x3c, 0xf0, 0xd3, 0xbd, 0x62, 0x76, 0xa1, 0x89, 0xb4, 0xc5, 0x52, 0x7a, 0x53, 0x4a,
	0xe7, 0x28, 0xc9, 0x64, 0xc8, 0xc5, 0xf6, 0x80, 0xd8, 0xc2, 0x71, 0x85, 0xda, 0xc4, 0x84, 0xa1,
	0xc2, 0x97, 0x57, 0x8c, 0x89, 0x4f, 0xe1, 0xc0, 0xaf, 0xc8, 0x61, 0xd1, 0x9b, 0xf7, 0xae, 0xc6,
	0x55, 0x1b, 0x1a, 0x8a, 0x1b, 0x69, 0x6c, 0xc1, 0x5a, 0x9a, 0x03, 0xbf, 0x84, 0x43, 0xe3, 0xb6,
	0x26, 0xe1, 0xb2, 0x26, 0x85, 0x45, 0x3f, 0x77, 0xd0, 0xa2, 0x07, 0x52, 0xc8, 0xcf, 0x60, 0xe8,
	0x65, 0x8c, 0x6f, 0xaa, 0x18, 0xe4, 0xae, 0x4b, 0xe2, 0x1e, 0x8e, 0x7e, 0xbd, 0xde, 0xee, 0xf8,
	0x04, 0x18, 0xad, 0xb3, 0x72, 0x54, 0x31, 0x5a, 0x37, 0xd9, 0xa8, 0x4e, 0xc6, 0x8c, 0xe2, 0x27,
	0xd0, 0xdf, 0x44, 0xdd, 0x19, 0x9a, 0xf1, 0xc6, 0xc1, 0xe0, 0x31, 0x62, 0xe0, 0xcf, 0x30, 0xfa,
	0xb9, 0x8a, 0x8b, 0xc5, 0xfe, 0xcf, 0xfc, 0x79, 0xdf, 0x6c, 0xfe, 0xef, 0x8e, 0xb7, 0x3b, 0x31,
	0x7d, 0xff, 0xfc, 0xfa, 0x60, 0x13, 0x31, 0x2e, 0x43, 0xc7, 0xef, 0x7a, 0xd7, 0x4f, 0xc3, 0xbc,
	0x7d, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x66, 0x22, 0x7c, 0xfe, 0x7d, 0x01, 0x00, 0x00,
}