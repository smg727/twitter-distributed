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
	UserExistsReply
	UserExistsRequest
	AddTweetRequest
	AddTweetReply
	Tweet
	OwnTweetsReply
	OwnTweetsRequest
	DeleteReply
	User
	UsersToFollowRequest
	UsersToFollowResponse
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

type UserExistsReply struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *UserExistsReply) Reset()                    { *m = UserExistsReply{} }
func (m *UserExistsReply) String() string            { return proto.CompactTextString(m) }
func (*UserExistsReply) ProtoMessage()               {}
func (*UserExistsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserExistsReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type UserExistsRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UserExistsRequest) Reset()                    { *m = UserExistsRequest{} }
func (m *UserExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*UserExistsRequest) ProtoMessage()               {}
func (*UserExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserExistsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AddTweetRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	TweetText string `protobuf:"bytes,2,opt,name=tweet_text,json=tweetText" json:"tweet_text,omitempty"`
}

func (m *AddTweetRequest) Reset()                    { *m = AddTweetRequest{} }
func (m *AddTweetRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTweetRequest) ProtoMessage()               {}
func (*AddTweetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddTweetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddTweetRequest) GetTweetText() string {
	if m != nil {
		return m.TweetText
	}
	return ""
}

type AddTweetReply struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *AddTweetReply) Reset()                    { *m = AddTweetReply{} }
func (m *AddTweetReply) String() string            { return proto.CompactTextString(m) }
func (*AddTweetReply) ProtoMessage()               {}
func (*AddTweetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddTweetReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type Tweet struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Tweet) Reset()                    { *m = Tweet{} }
func (m *Tweet) String() string            { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()               {}
func (*Tweet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Tweet) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type OwnTweetsReply struct {
	TweetList []*Tweet `protobuf:"bytes,1,rep,name=tweetList" json:"tweetList,omitempty"`
}

func (m *OwnTweetsReply) Reset()                    { *m = OwnTweetsReply{} }
func (m *OwnTweetsReply) String() string            { return proto.CompactTextString(m) }
func (*OwnTweetsReply) ProtoMessage()               {}
func (*OwnTweetsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *OwnTweetsReply) GetTweetList() []*Tweet {
	if m != nil {
		return m.TweetList
	}
	return nil
}

type OwnTweetsRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *OwnTweetsRequest) Reset()                    { *m = OwnTweetsRequest{} }
func (m *OwnTweetsRequest) String() string            { return proto.CompactTextString(m) }
func (*OwnTweetsRequest) ProtoMessage()               {}
func (*OwnTweetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *OwnTweetsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DeleteReply struct {
	DeleteStatus bool `protobuf:"varint,1,opt,name=deleteStatus" json:"deleteStatus,omitempty"`
}

func (m *DeleteReply) Reset()                    { *m = DeleteReply{} }
func (m *DeleteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()               {}
func (*DeleteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteReply) GetDeleteStatus() bool {
	if m != nil {
		return m.DeleteStatus
	}
	return false
}

type User struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UsersToFollowRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *UsersToFollowRequest) Reset()                    { *m = UsersToFollowRequest{} }
func (m *UsersToFollowRequest) String() string            { return proto.CompactTextString(m) }
func (*UsersToFollowRequest) ProtoMessage()               {}
func (*UsersToFollowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UsersToFollowRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UsersToFollowResponse struct {
	UsersToFollowList []*User `protobuf:"bytes,1,rep,name=usersToFollowList" json:"usersToFollowList,omitempty"`
}

func (m *UsersToFollowResponse) Reset()                    { *m = UsersToFollowResponse{} }
func (m *UsersToFollowResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersToFollowResponse) ProtoMessage()               {}
func (*UsersToFollowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UsersToFollowResponse) GetUsersToFollowList() []*User {
	if m != nil {
		return m.UsersToFollowList
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*Credentials)(nil), "helloworld.Credentials")
	proto.RegisterType((*RegisterReply)(nil), "helloworld.RegisterReply")
	proto.RegisterType((*LoginReply)(nil), "helloworld.LoginReply")
	proto.RegisterType((*UserExistsReply)(nil), "helloworld.UserExistsReply")
	proto.RegisterType((*UserExistsRequest)(nil), "helloworld.UserExistsRequest")
	proto.RegisterType((*AddTweetRequest)(nil), "helloworld.AddTweetRequest")
	proto.RegisterType((*AddTweetReply)(nil), "helloworld.AddTweetReply")
	proto.RegisterType((*Tweet)(nil), "helloworld.Tweet")
	proto.RegisterType((*OwnTweetsReply)(nil), "helloworld.OwnTweetsReply")
	proto.RegisterType((*OwnTweetsRequest)(nil), "helloworld.OwnTweetsRequest")
	proto.RegisterType((*DeleteReply)(nil), "helloworld.DeleteReply")
	proto.RegisterType((*User)(nil), "helloworld.User")
	proto.RegisterType((*UsersToFollowRequest)(nil), "helloworld.UsersToFollowRequest")
	proto.RegisterType((*UsersToFollowResponse)(nil), "helloworld.UsersToFollowResponse")
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
	UserExists(ctx context.Context, in *UserExistsRequest, opts ...grpc.CallOption) (*UserExistsReply, error)
	AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*AddTweetReply, error)
	OwnTweets(ctx context.Context, in *OwnTweetsRequest, opts ...grpc.CallOption) (*OwnTweetsReply, error)
	DeleteUser(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*DeleteReply, error)
	UsersToFollow(ctx context.Context, in *UsersToFollowRequest, opts ...grpc.CallOption) (*UsersToFollowResponse, error)
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

func (c *greeterClient) UserExists(ctx context.Context, in *UserExistsRequest, opts ...grpc.CallOption) (*UserExistsReply, error) {
	out := new(UserExistsReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/UserExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*AddTweetReply, error) {
	out := new(AddTweetReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/AddTweet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) OwnTweets(ctx context.Context, in *OwnTweetsRequest, opts ...grpc.CallOption) (*OwnTweetsReply, error) {
	out := new(OwnTweetsReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/OwnTweets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeleteUser(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UsersToFollow(ctx context.Context, in *UsersToFollowRequest, opts ...grpc.CallOption) (*UsersToFollowResponse, error) {
	out := new(UsersToFollowResponse)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/UsersToFollow", in, out, c.cc, opts...)
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
	UserExists(context.Context, *UserExistsRequest) (*UserExistsReply, error)
	AddTweet(context.Context, *AddTweetRequest) (*AddTweetReply, error)
	OwnTweets(context.Context, *OwnTweetsRequest) (*OwnTweetsReply, error)
	DeleteUser(context.Context, *Credentials) (*DeleteReply, error)
	UsersToFollow(context.Context, *UsersToFollowRequest) (*UsersToFollowResponse, error)
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

func _Greeter_UserExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UserExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/UserExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UserExists(ctx, req.(*UserExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_AddTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AddTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/AddTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AddTweet(ctx, req.(*AddTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_OwnTweets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnTweetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).OwnTweets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/OwnTweets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).OwnTweets(ctx, req.(*OwnTweetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeleteUser(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UsersToFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersToFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UsersToFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/UsersToFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UsersToFollow(ctx, req.(*UsersToFollowRequest))
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
		{
			MethodName: "UserExists",
			Handler:    _Greeter_UserExists_Handler,
		},
		{
			MethodName: "AddTweet",
			Handler:    _Greeter_AddTweet_Handler,
		},
		{
			MethodName: "OwnTweets",
			Handler:    _Greeter_OwnTweets_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Greeter_DeleteUser_Handler,
		},
		{
			MethodName: "UsersToFollow",
			Handler:    _Greeter_UsersToFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protodef.proto",
}

func init() { proto.RegisterFile("protodef.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6f, 0xd3, 0x4c,
	0x10, 0x75, 0xbe, 0x36, 0x6d, 0xf2, 0xd2, 0x34, 0xc9, 0xaa, 0x5f, 0x09, 0x0e, 0x95, 0xca, 0x0a,
	0x41, 0x7b, 0x71, 0x21, 0x88, 0x0b, 0x87, 0x8a, 0xb4, 0x85, 0x22, 0x14, 0x89, 0x2a, 0x0d, 0xf4,
	0x88, 0x0c, 0x1e, 0x82, 0x25, 0xc7, 0x36, 0xde, 0x8d, 0x92, 0x1c, 0xf9, 0xcf, 0xd1, 0xae, 0xe3,
	0x66, 0x6d, 0x48, 0x5a, 0x71, 0x9b, 0xd9, 0x79, 0xf3, 0x63, 0x67, 0xdf, 0x5b, 0xec, 0xc6, 0x49,
	0x24, 0x23, 0x8f, 0xbe, 0x3b, 0xda, 0x60, 0xf8, 0x41, 0x41, 0x10, 0x4d, 0xa3, 0x24, 0xf0, 0x38,
	0xc7, 0xce, 0x7b, 0xe5, 0x0d, 0xe8, 0xe7, 0x84, 0x84, 0x64, 0x0c, 0x9b, 0xa1, 0x3b, 0xa6, 0x76,
	0xe9, 0xb0, 0x74, 0x54, 0x1d, 0x68, 0x9b, 0x3f, 0x05, 0x16, 0x98, 0x38, 0x98, 0xb3, 0x36, 0xb6,
	0xc7, 0x24, 0x84, 0x3b, 0xca, 0x40, 0x99, 0xcb, 0x5f, 0xa1, 0x76, 0x9e, 0x90, 0x47, 0xa1, 0xf4,
	0xdd, 0x40, 0xb0, 0x3d, 0x94, 0x27, 0x46, 0xad, 0xd4, 0x61, 0x4d, 0x6c, 0xc4, 0x53, 0xaf, 0xfd,
	0x9f, 0x3e, 0x53, 0x26, 0x3f, 0x46, 0x7d, 0x40, 0x23, 0x5f, 0x48, 0x4a, 0xee, 0xea, 0xf0, 0x04,
	0xe8, 0x47, 0x23, 0x3f, 0x4c, 0x71, 0xfb, 0xd8, 0x12, 0xd2, 0x95, 0x13, 0xa1, 0x61, 0x95, 0xc1,
	0xc2, 0xe3, 0xc7, 0x68, 0x7c, 0x12, 0x94, 0xbc, 0x9d, 0xf9, 0x42, 0x8a, 0xf5, 0xd0, 0x13, 0xb4,
	0x4c, 0x68, 0xba, 0x03, 0x1b, 0x95, 0x89, 0xa0, 0xc4, 0x98, 0xfd, 0xd6, 0xe7, 0x7d, 0x34, 0x7a,
	0x9e, 0x37, 0x9c, 0x12, 0xc9, 0x7b, 0xc0, 0xd9, 0x01, 0x20, 0x15, 0xf6, 0x8b, 0xa4, 0x99, 0x5c,
	0x5c, 0xba, 0xaa, 0x4f, 0x86, 0x34, 0x93, 0xfc, 0x19, 0xea, 0xcb, 0x6a, 0xeb, 0xe6, 0xec, 0xa0,
	0xac, 0x51, 0xea, 0x7d, 0x74, 0xa9, 0xc5, 0xfb, 0x28, 0x9b, 0xf7, 0xb0, 0xfb, 0x71, 0x1a, 0xea,
	0xf8, 0xe2, 0xba, 0x27, 0x48, 0x9b, 0xf4, 0x7d, 0xa1, 0xa0, 0x1b, 0x47, 0xb5, 0x6e, 0xcb, 0x59,
	0xbe, 0xba, 0x93, 0x76, 0x5c, 0x62, 0xb8, 0x83, 0xa6, 0x51, 0xe2, 0xee, 0x35, 0xbc, 0x40, 0xed,
	0x82, 0x02, 0x92, 0x94, 0xf6, 0xe3, 0xd8, 0xf1, 0xb4, 0x7b, 0x6d, 0x0e, 0x9f, 0x3b, 0xe3, 0x1c,
	0x9b, 0x6a, 0xd5, 0x6b, 0xcb, 0x76, 0xb1, 0xa7, 0x30, 0x62, 0x18, 0xbd, 0x8b, 0xd4, 0xb0, 0xf7,
	0x19, 0xe5, 0x06, 0xff, 0x17, 0x72, 0x44, 0x1c, 0x85, 0x82, 0xd8, 0x29, 0x5a, 0x13, 0x33, 0x60,
	0x2c, 0xa3, 0x69, 0x2e, 0x43, 0x65, 0x0f, 0xfe, 0x84, 0x76, 0x7f, 0x95, 0xb1, 0x7d, 0x99, 0x10,
	0x49, 0x4a, 0xd8, 0x29, 0x2a, 0xd7, 0xee, 0x5c, 0xab, 0x80, 0xb5, 0xcd, 0x64, 0x53, 0x3c, 0xf6,
	0xfe, 0x5f, 0x22, 0x71, 0x30, 0xe7, 0x16, 0x3b, 0x47, 0x3d, 0xcb, 0xef, 0x8d, 0x5c, 0x3f, 0xfc,
	0xa7, 0x22, 0x6f, 0x50, 0xc9, 0x84, 0xc2, 0x1e, 0x98, 0x28, 0x43, 0x75, 0xf6, 0x43, 0x33, 0x90,
	0xd3, 0x15, 0xb7, 0xd8, 0x6b, 0x94, 0xb5, 0x7e, 0x56, 0xa7, 0xe7, 0xba, 0x2f, 0xb5, 0xc6, 0x2d,
	0xf6, 0x01, 0x58, 0x4a, 0x85, 0x1d, 0x14, 0x37, 0x98, 0x93, 0x90, 0xdd, 0x59, 0x15, 0x4e, 0x6b,
	0x5d, 0xa0, 0x92, 0xf1, 0x9e, 0xe5, 0xa0, 0x05, 0x6d, 0xe5, 0x6f, 0x93, 0x93, 0x0a, 0xb7, 0xd8,
	0x25, 0xaa, 0xb7, 0xa4, 0x65, 0x8f, 0x4c, 0x64, 0x91, 0xcb, 0xb6, 0xbd, 0x22, 0x9a, 0x2d, 0x16,
	0x29, 0x9b, 0x35, 0x41, 0x57, 0xee, 0x26, 0x17, 0x30, 0xe8, 0xcf, 0x2d, 0xf6, 0x19, 0xf5, 0x1c,
	0x09, 0xd9, 0x61, 0x71, 0x01, 0x45, 0x4e, 0xdb, 0x8f, 0xd7, 0x20, 0x52, 0x06, 0x73, 0xeb, 0xec,
	0x39, 0x3a, 0x7e, 0xe4, 0x8c, 0x92, 0xf8, 0x9b, 0x43, 0x33, 0x77, 0x1c, 0x07, 0x24, 0x8c, 0xb4,
	0xb3, 0x86, 0xe6, 0xc7, 0x8d, 0xb2, 0xaf, 0xd4, 0xd7, 0x7e, 0x55, 0xfa, 0xba, 0xa5, 0xff, 0xf8,
	0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xe1, 0x80, 0x60, 0xf5, 0x05, 0x00, 0x00,
}
