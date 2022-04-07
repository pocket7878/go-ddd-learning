// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: go_ddd_learning.proto

package go_ddd_learning

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoDDDLearningServiceClient is the client API for GoDDDLearningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoDDDLearningServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	PostponeTask(ctx context.Context, in *PostponeTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type goDDDLearningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoDDDLearningServiceClient(cc grpc.ClientConnInterface) GoDDDLearningServiceClient {
	return &goDDDLearningServiceClient{cc}
}

func (c *goDDDLearningServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/go_ddd_learning.GoDDDLearningService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goDDDLearningServiceClient) DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/go_ddd_learning.GoDDDLearningService/DeactivateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goDDDLearningServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/go_ddd_learning.GoDDDLearningService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goDDDLearningServiceClient) PostponeTask(ctx context.Context, in *PostponeTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/go_ddd_learning.GoDDDLearningService/PostponeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoDDDLearningServiceServer is the server API for GoDDDLearningService service.
// All implementations must embed UnimplementedGoDDDLearningServiceServer
// for forward compatibility
type GoDDDLearningServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	DeactivateUser(context.Context, *DeactivateUserRequest) (*UserResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error)
	PostponeTask(context.Context, *PostponeTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGoDDDLearningServiceServer()
}

// UnimplementedGoDDDLearningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoDDDLearningServiceServer struct {
}

func (UnimplementedGoDDDLearningServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGoDDDLearningServiceServer) DeactivateUser(context.Context, *DeactivateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateUser not implemented")
}
func (UnimplementedGoDDDLearningServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedGoDDDLearningServiceServer) PostponeTask(context.Context, *PostponeTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostponeTask not implemented")
}
func (UnimplementedGoDDDLearningServiceServer) mustEmbedUnimplementedGoDDDLearningServiceServer() {}

// UnsafeGoDDDLearningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoDDDLearningServiceServer will
// result in compilation errors.
type UnsafeGoDDDLearningServiceServer interface {
	mustEmbedUnimplementedGoDDDLearningServiceServer()
}

func RegisterGoDDDLearningServiceServer(s grpc.ServiceRegistrar, srv GoDDDLearningServiceServer) {
	s.RegisterService(&GoDDDLearningService_ServiceDesc, srv)
}

func _GoDDDLearningService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoDDDLearningServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_ddd_learning.GoDDDLearningService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoDDDLearningServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoDDDLearningService_DeactivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoDDDLearningServiceServer).DeactivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_ddd_learning.GoDDDLearningService/DeactivateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoDDDLearningServiceServer).DeactivateUser(ctx, req.(*DeactivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoDDDLearningService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoDDDLearningServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_ddd_learning.GoDDDLearningService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoDDDLearningServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoDDDLearningService_PostponeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostponeTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoDDDLearningServiceServer).PostponeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_ddd_learning.GoDDDLearningService/PostponeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoDDDLearningServiceServer).PostponeTask(ctx, req.(*PostponeTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoDDDLearningService_ServiceDesc is the grpc.ServiceDesc for GoDDDLearningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoDDDLearningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_ddd_learning.GoDDDLearningService",
	HandlerType: (*GoDDDLearningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _GoDDDLearningService_CreateUser_Handler,
		},
		{
			MethodName: "DeactivateUser",
			Handler:    _GoDDDLearningService_DeactivateUser_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _GoDDDLearningService_CreateTask_Handler,
		},
		{
			MethodName: "PostponeTask",
			Handler:    _GoDDDLearningService_PostponeTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_ddd_learning.proto",
}