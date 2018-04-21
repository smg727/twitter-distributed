// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protodef.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	protodef.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	Credentials
	RegisterReply
	LoginReply
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Credentials struct {
	Uname string `protobuf:"bytes,1,opt,name=uname" json:"uname,omitempty"`
	Pwd   string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Credentials) GetUname() string {
	if m != nil {
		return m.Uname
	}
	return ""
}

func (m *Credentials) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type RegisterReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RegisterReply) Reset()                    { *m = RegisterReply{} }
func (m *RegisterReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()               {}
func (*RegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LoginReply struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*Credentials)(nil), "helloworld.Credentials")
	proto.RegisterType((*RegisterReply)(nil), "helloworld.RegisterReply")
	proto.RegisterType((*LoginReply)(nil), "helloworld.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Register(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
	Register(context.Context, *Credentials) (*RegisterReply, error)
	Login(context.Context, *Credentials) (*LoginReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Register(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Greeter_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Greeter_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protodef.proto",
}

func init() { proto.RegisterFile("protodef.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xab, 0x6d, 0xe3, 0x68, 0x55, 0x06, 0xa9, 0xb1, 0x5e, 0x64, 0x11, 0xd1, 0x4b, 0x10,
	0xc5, 0x8b, 0x07, 0xd1, 0xf6, 0xa0, 0x07, 0x0f, 0x25, 0x1e, 0x3c, 0xaf, 0x66, 0x8c, 0x81, 0x4d,
	0x76, 0xdd, 0xdd, 0x50, 0xf3, 0x1f, 0xfc, 0xd1, 0x92, 0x6d, 0x43, 0xb7, 0xa0, 0x08, 0xde, 0xde,
	0xdb, 0x7d, 0xf3, 0xf1, 0x1e, 0x03, 0xdb, 0x4a, 0x4b, 0x2b, 0x13, 0x7a, 0x8b, 0x1c, 0x40, 0x78,
	0x27, 0x21, 0xe4, 0x4c, 0x6a, 0x91, 0x30, 0x06, 0x5b, 0x0f, 0x35, 0x8b, 0xe9, 0xa3, 0x24, 0x63,
	0x11, 0x61, 0xbd, 0xe0, 0x39, 0x85, 0xed, 0xa3, 0xf6, 0xe9, 0x46, 0xec, 0x30, 0x3b, 0x01, 0x58,
	0x68, 0x94, 0xa8, 0x30, 0x84, 0x7e, 0x4e, 0xc6, 0xf0, 0xb4, 0x11, 0x35, 0x94, 0x5d, 0xc1, 0xe6,
	0x44, 0x53, 0x42, 0x85, 0xcd, 0xb8, 0x30, 0xb8, 0x07, 0xdd, 0xd2, 0xeb, 0x35, 0x27, 0xb8, 0x0b,
	0x6b, 0x6a, 0x96, 0x84, 0x1d, 0xf7, 0x56, 0x43, 0x76, 0x06, 0x83, 0x98, 0xd2, 0xcc, 0x58, 0xd2,
	0x7f, 0x4d, 0x38, 0x06, 0x78, 0x94, 0x69, 0x56, 0xcc, 0x75, 0x43, 0xe8, 0x19, 0xcb, 0x6d, 0x69,
	0x9c, 0x2c, 0x88, 0x17, 0xec, 0xe2, 0xab, 0x03, 0xfd, 0x7b, 0x4d, 0x64, 0x49, 0xe3, 0x0d, 0x04,
	0x4f, 0xbc, 0x72, 0xeb, 0x63, 0x18, 0x2d, 0x8d, 0x47, 0xbe, 0xeb, 0xd1, 0xf0, 0x87, 0x1f, 0x25,
	0x2a, 0xd6, 0xc2, 0x09, 0x0c, 0x9a, 0xfa, 0xbb, 0x94, 0x67, 0xc5, 0xbf, 0x9a, 0xdc, 0x42, 0xd0,
	0x38, 0xc4, 0x7d, 0x5f, 0xe5, 0xc5, 0x35, 0x3a, 0xf0, 0x3f, 0x56, 0x02, 0x61, 0x2d, 0xbc, 0x86,
	0xae, 0x33, 0xfe, 0x7b, 0xf9, 0xca, 0xf4, 0x65, 0x48, 0xac, 0x35, 0x3e, 0x87, 0xc3, 0x4c, 0x46,
	0xa9, 0x56, 0xaf, 0x11, 0x7d, 0xf2, 0x5c, 0x09, 0x32, 0x9e, 0x76, 0xbc, 0xe3, 0x56, 0x7d, 0xae,
	0xf1, 0xb4, 0x3e, 0x8f, 0x69, 0xfb, 0xa5, 0xe7, 0xee, 0xe4, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xb9, 0x3c, 0xf1, 0x39, 0x02, 0x00, 0x00,
}