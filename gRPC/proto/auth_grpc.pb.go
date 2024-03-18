// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: gRPC/proto/auth.proto

package auth

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateAuth(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenRes, error)
	VerifyAuth(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateAuth(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenRes, error) {
	out := new(CreateTokenRes)
	err := c.cc.Invoke(ctx, "/AuthService/CreateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyAuth(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error) {
	out := new(VerifyTokenRes)
	err := c.cc.Invoke(ctx, "/AuthService/VerifyAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	CreateAuth(context.Context, *CreateTokenReq) (*CreateTokenRes, error)
	VerifyAuth(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) CreateAuth(context.Context, *CreateTokenReq) (*CreateTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuth not implemented")
}
func (UnimplementedAuthServiceServer) VerifyAuth(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAuth not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/CreateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateAuth(ctx, req.(*CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/VerifyAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyAuth(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuth",
			Handler:    _AuthService_CreateAuth_Handler,
		},
		{
			MethodName: "VerifyAuth",
			Handler:    _AuthService_VerifyAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto/auth.proto",
}
