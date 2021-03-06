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
	return fileDescriptor_recharge_protocol_54bf3c7750172340, []int{0}
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
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_recharge_protocol_54bf3c7750172340, []int{1}
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

func (m *RegisterReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RegisterReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterReply) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RechargeRequest struct {
	FGoldinFlowId        string   `protobuf:"bytes,1,opt,name=f_goldin_flow_id,json=fGoldinFlowId,proto3" json:"f_goldin_flow_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount               string   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeRequest) Reset()         { *m = RechargeRequest{} }
func (m *RechargeRequest) String() string { return proto.CompactTextString(m) }
func (*RechargeRequest) ProtoMessage()    {}
func (*RechargeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recharge_protocol_54bf3c7750172340, []int{2}
}
func (m *RechargeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeRequest.Unmarshal(m, b)
}
func (m *RechargeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeRequest.Marshal(b, m, deterministic)
}
func (dst *RechargeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeRequest.Merge(dst, src)
}
func (m *RechargeRequest) XXX_Size() int {
	return xxx_messageInfo_RechargeRequest.Size(m)
}
func (m *RechargeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeRequest proto.InternalMessageInfo

func (m *RechargeRequest) GetFGoldinFlowId() string {
	if m != nil {
		return m.FGoldinFlowId
	}
	return ""
}

func (m *RechargeRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RechargeRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type RechargeReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeReply) Reset()         { *m = RechargeReply{} }
func (m *RechargeReply) String() string { return proto.CompactTextString(m) }
func (*RechargeReply) ProtoMessage()    {}
func (*RechargeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_recharge_protocol_54bf3c7750172340, []int{3}
}
func (m *RechargeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeReply.Unmarshal(m, b)
}
func (m *RechargeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeReply.Marshal(b, m, deterministic)
}
func (dst *RechargeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeReply.Merge(dst, src)
}
func (m *RechargeReply) XXX_Size() int {
	return xxx_messageInfo_RechargeReply.Size(m)
}
func (m *RechargeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeReply.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeReply proto.InternalMessageInfo

func (m *RechargeReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RechargeReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "recharge_protocol.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "recharge_protocol.RegisterReply")
	proto.RegisterType((*RechargeRequest)(nil), "recharge_protocol.RechargeRequest")
	proto.RegisterType((*RechargeReply)(nil), "recharge_protocol.RechargeReply")
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

// RechargeClient is the client API for Recharge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RechargeClient interface {
	Recharge(ctx context.Context, in *RechargeRequest, opts ...grpc.CallOption) (*RechargeReply, error)
}

type rechargeClient struct {
	cc *grpc.ClientConn
}

func NewRechargeClient(cc *grpc.ClientConn) RechargeClient {
	return &rechargeClient{cc}
}

func (c *rechargeClient) Recharge(ctx context.Context, in *RechargeRequest, opts ...grpc.CallOption) (*RechargeReply, error) {
	out := new(RechargeReply)
	err := c.cc.Invoke(ctx, "/recharge_protocol.Recharge/Recharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RechargeServer is the server API for Recharge service.
type RechargeServer interface {
	Recharge(context.Context, *RechargeRequest) (*RechargeReply, error)
}

func RegisterRechargeServer(s *grpc.Server, srv RechargeServer) {
	s.RegisterService(&_Recharge_serviceDesc, srv)
}

func _Recharge_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RechargeServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recharge_protocol.Recharge/Recharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RechargeServer).Recharge(ctx, req.(*RechargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recharge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recharge_protocol.Recharge",
	HandlerType: (*RechargeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recharge",
			Handler:    _Recharge_Recharge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recharge_protocol.proto",
}

func init() {
	proto.RegisterFile("recharge_protocol.proto", fileDescriptor_recharge_protocol_54bf3c7750172340)
}

var fileDescriptor_recharge_protocol_54bf3c7750172340 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x99, 0xfe, 0xfb, 0xbe, 0x46, 0xaa, 0x35, 0x14, 0x3b, 0x14, 0xc1, 0x32, 0x1b, 0xa5,
	0x8b, 0x16, 0xea, 0xca, 0x6e, 0x05, 0xa5, 0xdb, 0x01, 0xd7, 0x43, 0xec, 0xdc, 0x99, 0x06, 0xd3,
	0xdc, 0x31, 0xc9, 0xb4, 0x74, 0xeb, 0x2b, 0xf8, 0x68, 0xbe, 0x82, 0x0f, 0x22, 0x99, 0xcc, 0x60,
	0x6b, 0x55, 0x70, 0x77, 0xcf, 0xb9, 0xf7, 0xe6, 0x97, 0x93, 0x90, 0xbe, 0x82, 0xc5, 0x92, 0xa9,
	0x14, 0xa2, 0x4c, 0xa1, 0xc1, 0x05, 0x8a, 0x71, 0x51, 0xd0, 0xd3, 0x83, 0xc6, 0xe0, 0x3c, 0x45,
	0x4c, 0x05, 0x4c, 0x58, 0xc6, 0x27, 0x4c, 0x4a, 0x34, 0xcc, 0x70, 0x94, 0xda, 0x2d, 0x04, 0x6b,
	0x72, 0x12, 0x42, 0xca, 0xb5, 0x01, 0x15, 0xc2, 0x73, 0x0e, 0xda, 0xd0, 0x1e, 0x69, 0x66, 0x4b,
	0x94, 0xe0, 0x7b, 0x43, 0xef, 0xaa, 0x1d, 0x3a, 0x61, 0x5d, 0x58, 0x31, 0x2e, 0xfc, 0x9a, 0x73,
	0x0b, 0x41, 0x2f, 0xc8, 0x11, 0x97, 0x6b, 0x6e, 0x20, 0x5a, 0x60, 0x0c, 0x7e, 0xbd, 0xe8, 0x11,
	0x67, 0xdd, 0x62, 0x0c, 0xf4, 0x8c, 0xb4, 0x32, 0xa6, 0xf5, 0x26, 0xf6, 0x1b, 0x45, 0xaf, 0x54,
	0x41, 0x48, 0x3a, 0x9f, 0xdc, 0x4c, 0x6c, 0xed, 0xa0, 0x36, 0xcc, 0xe4, 0xba, 0xc0, 0x36, 0xc3,
	0x52, 0xd1, 0x2e, 0xa9, 0xaf, 0x74, 0x5a, 0x52, 0x6d, 0x49, 0xfb, 0xe4, 0x5f, 0xae, 0x41, 0x45,
	0x3c, 0x2e, 0x79, 0x2d, 0x2b, 0xe7, 0x71, 0xf0, 0x64, 0xb3, 0xb8, 0xf8, 0x55, 0x96, 0x4b, 0xd2,
	0x4d, 0xa2, 0x14, 0x45, 0xcc, 0x65, 0x94, 0x08, 0xdc, 0xd8, 0x25, 0x17, 0xab, 0x93, 0xdc, 0x17,
	0xf6, 0x9d, 0xc0, 0xcd, 0x3c, 0xde, 0x3d, 0xb4, 0xb6, 0x7b, 0xa8, 0xbd, 0x17, 0x5b, 0x61, 0x2e,
	0x4d, 0x05, 0x73, 0x2a, 0xb8, 0xb1, 0x01, 0x2a, 0xd8, 0x9f, 0x02, 0x4c, 0x25, 0x69, 0x3c, 0x68,
	0x50, 0x34, 0x21, 0xff, 0xab, 0x37, 0xa0, 0xc1, 0xf8, 0xf0, 0x4b, 0xbf, 0x7c, 0xcc, 0x60, 0xf8,
	0xeb, 0x4c, 0x26, 0xb6, 0x41, 0xef, 0xe5, 0xed, 0xfd, 0xb5, 0x76, 0x1c, 0xb4, 0x27, 0xaa, 0xf4,
	0x67, 0xde, 0x68, 0xaa, 0x2c, 0xc7, 0x2d, 0x3a, 0x66, 0x59, 0x7f, 0xcf, 0xdc, 0x7b, 0xc0, 0x1f,
	0x98, 0x3b, 0xb9, 0xf7, 0x98, 0xce, 0x9f, 0x79, 0xa3, 0xc7, 0x56, 0x31, 0x7d, 0xfd, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x6a, 0x6b, 0xdb, 0xaa, 0x02, 0x00, 0x00,
}
