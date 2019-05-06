// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package models

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserProto struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	HashPassword         string   `protobuf:"bytes,5,opt,name=hashPassword,proto3" json:"hashPassword,omitempty"`
	Score                uint64   `protobuf:"varint,6,opt,name=score,proto3" json:"score"`
	Picture              string   `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
	Games                uint64   `protobuf:"varint,8,opt,name=games,proto3" json:"games"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProto) Reset()         { *m = UserProto{} }
func (m *UserProto) String() string { return proto.CompactTextString(m) }
func (*UserProto) ProtoMessage()    {}
func (*UserProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *UserProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProto.Unmarshal(m, b)
}
func (m *UserProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProto.Marshal(b, m, deterministic)
}
func (m *UserProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProto.Merge(m, src)
}
func (m *UserProto) XXX_Size() int {
	return xxx_messageInfo_UserProto.Size(m)
}
func (m *UserProto) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProto.DiscardUnknown(m)
}

var xxx_messageInfo_UserProto proto.InternalMessageInfo

func (m *UserProto) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UserProto) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UserProto) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProto) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserProto) GetHashPassword() string {
	if m != nil {
		return m.HashPassword
	}
	return ""
}

func (m *UserProto) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *UserProto) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *UserProto) GetGames() uint64 {
	if m != nil {
		return m.Games
	}
	return 0
}

type JWT struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWT) Reset()         { *m = JWT{} }
func (m *JWT) String() string { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()    {}
func (*JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *JWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWT.Unmarshal(m, b)
}
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWT.Marshal(b, m, deterministic)
}
func (m *JWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWT.Merge(m, src)
}
func (m *JWT) XXX_Size() int {
	return xxx_messageInfo_JWT.Size(m)
}
func (m *JWT) XXX_DiscardUnknown() {
	xxx_messageInfo_JWT.DiscardUnknown(m)
}

var xxx_messageInfo_JWT proto.InternalMessageInfo

func (m *JWT) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserProto)(nil), "models.UserProto")
	proto.RegisterType((*JWT)(nil), "models.JWT")
	proto.RegisterType((*Nothing)(nil), "models.Nothing")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0x32, 0x31,
	0x14, 0xc5, 0x1d, 0xff, 0x8d, 0x73, 0x15, 0x3f, 0xbe, 0xd0, 0x45, 0xb0, 0x1b, 0xc9, 0x4a, 0x68,
	0x2b, 0xc5, 0x3e, 0x41, 0x51, 0x28, 0x95, 0x52, 0x64, 0xb0, 0xb8, 0x4e, 0xf5, 0x92, 0x04, 0x67,
	0x26, 0x92, 0x44, 0xfa, 0x94, 0x7d, 0xa4, 0x42, 0x49, 0xd2, 0x91, 0x16, 0x29, 0x74, 0xf9, 0x3b,
	0xe7, 0xdc, 0x5c, 0xee, 0x21, 0x30, 0x2c, 0xd1, 0x5a, 0x2e, 0xd0, 0x4e, 0x0f, 0x46, 0x3b, 0x4d,
	0xba, 0xa5, 0xde, 0x61, 0x61, 0xd9, 0x7b, 0x02, 0xd9, 0x8b, 0x45, 0xb3, 0x0a, 0xea, 0x10, 0x9a,
	0x8f, 0x0b, 0x9a, 0x8c, 0x93, 0x49, 0x3b, 0x6f, 0xaa, 0x05, 0xb9, 0x80, 0x4e, 0xa1, 0x85, 0xaa,
	0x68, 0x73, 0x9c, 0x4c, 0xb2, 0x3c, 0x82, 0x57, 0xb1, 0xe4, 0xaa, 0xa0, 0xad, 0xa8, 0x06, 0x20,
	0x23, 0xe8, 0x1d, 0xb8, 0xb5, 0x6f, 0xda, 0xec, 0x68, 0x3b, 0x18, 0x27, 0x26, 0x0c, 0x06, 0x92,
	0x5b, 0xb9, 0xaa, 0xfd, 0x4e, 0xf0, 0x7f, 0x68, 0xfe, 0x55, 0xbb, 0xd5, 0x06, 0x69, 0x37, 0xac,
	0x8f, 0x40, 0x28, 0xa4, 0x07, 0xb5, 0x75, 0x47, 0x83, 0x34, 0x0d, 0x43, 0x35, 0xfa, 0xbc, 0xe0,
	0x25, 0x5a, 0xda, 0x8b, 0xf9, 0x00, 0xec, 0x12, 0x5a, 0xcb, 0xcd, 0xda, 0x9b, 0x4e, 0xef, 0xb1,
	0x0a, 0xb7, 0x64, 0x79, 0x04, 0x96, 0x41, 0xfa, 0xac, 0x9d, 0x54, 0x95, 0x98, 0x7d, 0x24, 0xd0,
	0xbf, 0x3f, 0x3a, 0x39, 0x97, 0xb8, 0xdd, 0xa3, 0x21, 0x37, 0x90, 0x3d, 0xf9, 0xe3, 0x7c, 0x17,
	0xe4, 0xff, 0x34, 0xb6, 0x33, 0x3d, 0x35, 0x33, 0xea, 0xd7, 0xd2, 0x72, 0xb3, 0x66, 0x0d, 0x72,
	0x0b, 0x83, 0x1c, 0x85, 0xb2, 0x0e, 0xcd, 0x1f, 0x27, 0xae, 0x20, 0x7d, 0x40, 0x17, 0xc2, 0xdf,
	0x9d, 0xd1, 0xf9, 0x24, 0x6b, 0x90, 0x19, 0xc0, 0x5c, 0xf2, 0x4a, 0xe0, 0x6f, 0x8f, 0xff, 0xab,
	0xa5, 0xaf, 0x7b, 0x58, 0x83, 0x5c, 0x03, 0x2c, 0xb0, 0x40, 0x87, 0xe7, 0x3b, 0xce, 0xd3, 0xaf,
	0xdd, 0xf0, 0x0d, 0xee, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xee, 0x7e, 0x99, 0x18, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	LoginUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*JWT, error)
	RegisterUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*JWT, error)
	//GetUser() also checks JWT. If token if invalid, nil will be returned.
	GetUser(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*UserProto, error)
	ChangeUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*Nothing, error)
	DeleteUser(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*Nothing, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) LoginUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*JWT, error) {
	out := new(JWT)
	err := c.cc.Invoke(ctx, "/models.AuthChecker/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) RegisterUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*JWT, error) {
	out := new(JWT)
	err := c.cc.Invoke(ctx, "/models.AuthChecker/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetUser(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*UserProto, error) {
	out := new(UserProto)
	err := c.cc.Invoke(ctx, "/models.AuthChecker/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) ChangeUser(ctx context.Context, in *UserProto, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/models.AuthChecker/ChangeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) DeleteUser(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/models.AuthChecker/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	LoginUser(context.Context, *UserProto) (*JWT, error)
	RegisterUser(context.Context, *UserProto) (*JWT, error)
	//GetUser() also checks JWT. If token if invalid, nil will be returned.
	GetUser(context.Context, *JWT) (*UserProto, error)
	ChangeUser(context.Context, *UserProto) (*Nothing, error)
	DeleteUser(context.Context, *JWT) (*Nothing, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.AuthChecker/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).LoginUser(ctx, req.(*UserProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.AuthChecker/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).RegisterUser(ctx, req.(*UserProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.AuthChecker/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetUser(ctx, req.(*JWT))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.AuthChecker/ChangeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).ChangeUser(ctx, req.(*UserProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.AuthChecker/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).DeleteUser(ctx, req.(*JWT))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _AuthChecker_LoginUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _AuthChecker_RegisterUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthChecker_GetUser_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _AuthChecker_ChangeUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthChecker_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}