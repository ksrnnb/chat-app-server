// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chatpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatAppServiceClient is the client API for ChatAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatAppServiceClient interface {
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (ChatAppService_GetMessagesClient, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	// ユーザーのログイン処理
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// 動作確認用
	HelloMessage(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type chatAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatAppServiceClient(cc grpc.ClientConnInterface) ChatAppServiceClient {
	return &chatAppServiceClient{cc}
}

func (c *chatAppServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (ChatAppService_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatAppService_ServiceDesc.Streams[0], "/chatpb.ChatAppService/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatAppServiceGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatAppService_GetMessagesClient interface {
	Recv() (*GetMessagesResponse, error)
	grpc.ClientStream
}

type chatAppServiceGetMessagesClient struct {
	grpc.ClientStream
}

func (x *chatAppServiceGetMessagesClient) Recv() (*GetMessagesResponse, error) {
	m := new(GetMessagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatAppServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/chatpb.ChatAppService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAppServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/chatpb.ChatAppService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAppServiceClient) HelloMessage(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/chatpb.ChatAppService/HelloMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatAppServiceServer is the server API for ChatAppService service.
// All implementations must embed UnimplementedChatAppServiceServer
// for forward compatibility
type ChatAppServiceServer interface {
	GetMessages(*GetMessagesRequest, ChatAppService_GetMessagesServer) error
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	// ユーザーのログイン処理
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// 動作確認用
	HelloMessage(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedChatAppServiceServer()
}

// UnimplementedChatAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatAppServiceServer struct {
}

func (UnimplementedChatAppServiceServer) GetMessages(*GetMessagesRequest, ChatAppService_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatAppServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedChatAppServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatAppServiceServer) HelloMessage(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloMessage not implemented")
}
func (UnimplementedChatAppServiceServer) mustEmbedUnimplementedChatAppServiceServer() {}

// UnsafeChatAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatAppServiceServer will
// result in compilation errors.
type UnsafeChatAppServiceServer interface {
	mustEmbedUnimplementedChatAppServiceServer()
}

func RegisterChatAppServiceServer(s grpc.ServiceRegistrar, srv ChatAppServiceServer) {
	s.RegisterService(&ChatAppService_ServiceDesc, srv)
}

func _ChatAppService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatAppServiceServer).GetMessages(m, &chatAppServiceGetMessagesServer{stream})
}

type ChatAppService_GetMessagesServer interface {
	Send(*GetMessagesResponse) error
	grpc.ServerStream
}

type chatAppServiceGetMessagesServer struct {
	grpc.ServerStream
}

func (x *chatAppServiceGetMessagesServer) Send(m *GetMessagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatAppService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAppServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.ChatAppService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAppServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAppService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAppServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.ChatAppService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAppServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAppService_HelloMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAppServiceServer).HelloMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.ChatAppService/HelloMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAppServiceServer).HelloMessage(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatAppService_ServiceDesc is the grpc.ServiceDesc for ChatAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatpb.ChatAppService",
	HandlerType: (*ChatAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _ChatAppService_CreateMessage_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ChatAppService_Login_Handler,
		},
		{
			MethodName: "HelloMessage",
			Handler:    _ChatAppService_HelloMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessages",
			Handler:       _ChatAppService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat-app.proto",
}
