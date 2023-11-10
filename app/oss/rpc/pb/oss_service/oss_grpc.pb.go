// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: oss.proto

package oss_service

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

const (
	Oss_CreatUpToken_FullMethodName     = "/service.Oss/CreatUpToken"
	Oss_GetEndpoint_FullMethodName      = "/service.Oss/GetEndpoint"
	Oss_GetFileAccessUrl_FullMethodName = "/service.Oss/GetFileAccessUrl"
)

// OssClient is the client API for Oss service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OssClient interface {
	CreatUpToken(ctx context.Context, in *CreateUpTokenRequest, opts ...grpc.CallOption) (*CreateUpTokenResponse, error)
	GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointResponse, error)
	GetFileAccessUrl(ctx context.Context, in *GetFileAccessUrlRequest, opts ...grpc.CallOption) (*GetFileAccessUrlResponse, error)
}

type ossClient struct {
	cc grpc.ClientConnInterface
}

func NewOssClient(cc grpc.ClientConnInterface) OssClient {
	return &ossClient{cc}
}

func (c *ossClient) CreatUpToken(ctx context.Context, in *CreateUpTokenRequest, opts ...grpc.CallOption) (*CreateUpTokenResponse, error) {
	out := new(CreateUpTokenResponse)
	err := c.cc.Invoke(ctx, Oss_CreatUpToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossClient) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointResponse, error) {
	out := new(GetEndpointResponse)
	err := c.cc.Invoke(ctx, Oss_GetEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ossClient) GetFileAccessUrl(ctx context.Context, in *GetFileAccessUrlRequest, opts ...grpc.CallOption) (*GetFileAccessUrlResponse, error) {
	out := new(GetFileAccessUrlResponse)
	err := c.cc.Invoke(ctx, Oss_GetFileAccessUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OssServer is the server API for Oss service.
// All implementations must embed UnimplementedOssServer
// for forward compatibility
type OssServer interface {
	CreatUpToken(context.Context, *CreateUpTokenRequest) (*CreateUpTokenResponse, error)
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointResponse, error)
	GetFileAccessUrl(context.Context, *GetFileAccessUrlRequest) (*GetFileAccessUrlResponse, error)
	mustEmbedUnimplementedOssServer()
}

// UnimplementedOssServer must be embedded to have forward compatible implementations.
type UnimplementedOssServer struct {
}

func (UnimplementedOssServer) CreatUpToken(context.Context, *CreateUpTokenRequest) (*CreateUpTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatUpToken not implemented")
}
func (UnimplementedOssServer) GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoint not implemented")
}
func (UnimplementedOssServer) GetFileAccessUrl(context.Context, *GetFileAccessUrlRequest) (*GetFileAccessUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileAccessUrl not implemented")
}
func (UnimplementedOssServer) mustEmbedUnimplementedOssServer() {}

// UnsafeOssServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OssServer will
// result in compilation errors.
type UnsafeOssServer interface {
	mustEmbedUnimplementedOssServer()
}

func RegisterOssServer(s grpc.ServiceRegistrar, srv OssServer) {
	s.RegisterService(&Oss_ServiceDesc, srv)
}

func _Oss_CreatUpToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUpTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServer).CreatUpToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oss_CreatUpToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServer).CreatUpToken(ctx, req.(*CreateUpTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oss_GetEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServer).GetEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oss_GetEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServer).GetEndpoint(ctx, req.(*GetEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oss_GetFileAccessUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileAccessUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OssServer).GetFileAccessUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oss_GetFileAccessUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OssServer).GetFileAccessUrl(ctx, req.(*GetFileAccessUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oss_ServiceDesc is the grpc.ServiceDesc for Oss service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oss_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Oss",
	HandlerType: (*OssServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatUpToken",
			Handler:    _Oss_CreatUpToken_Handler,
		},
		{
			MethodName: "GetEndpoint",
			Handler:    _Oss_GetEndpoint_Handler,
		},
		{
			MethodName: "GetFileAccessUrl",
			Handler:    _Oss_GetFileAccessUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oss.proto",
}