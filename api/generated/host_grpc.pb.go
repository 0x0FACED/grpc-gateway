// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.3
// source: api/proto/host.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HostService_SetHostname_FullMethodName     = "/api.HostService/SetHostname"
	HostService_GetDNSServers_FullMethodName   = "/api.HostService/GetDNSServers"
	HostService_AddDNSServer_FullMethodName    = "/api.HostService/AddDNSServer"
	HostService_RemoveDNSServer_FullMethodName = "/api.HostService/RemoveDNSServer"
)

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostServiceClient interface {
	SetHostname(ctx context.Context, in *SetHostnameRequest, opts ...grpc.CallOption) (*SetHostnameResponse, error)
	GetDNSServers(ctx context.Context, in *GetDNSServersRequest, opts ...grpc.CallOption) (*GetDNSServersResponse, error)
	AddDNSServer(ctx context.Context, in *AddDNSServerRequest, opts ...grpc.CallOption) (*AddDNSServerResponse, error)
	RemoveDNSServer(ctx context.Context, in *RemoveDNSServerRequest, opts ...grpc.CallOption) (*RemoveDNSServerResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) SetHostname(ctx context.Context, in *SetHostnameRequest, opts ...grpc.CallOption) (*SetHostnameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetHostnameResponse)
	err := c.cc.Invoke(ctx, HostService_SetHostname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) GetDNSServers(ctx context.Context, in *GetDNSServersRequest, opts ...grpc.CallOption) (*GetDNSServersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDNSServersResponse)
	err := c.cc.Invoke(ctx, HostService_GetDNSServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) AddDNSServer(ctx context.Context, in *AddDNSServerRequest, opts ...grpc.CallOption) (*AddDNSServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDNSServerResponse)
	err := c.cc.Invoke(ctx, HostService_AddDNSServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostServiceClient) RemoveDNSServer(ctx context.Context, in *RemoveDNSServerRequest, opts ...grpc.CallOption) (*RemoveDNSServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveDNSServerResponse)
	err := c.cc.Invoke(ctx, HostService_RemoveDNSServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
// All implementations must embed UnimplementedHostServiceServer
// for forward compatibility
type HostServiceServer interface {
	SetHostname(context.Context, *SetHostnameRequest) (*SetHostnameResponse, error)
	GetDNSServers(context.Context, *GetDNSServersRequest) (*GetDNSServersResponse, error)
	AddDNSServer(context.Context, *AddDNSServerRequest) (*AddDNSServerResponse, error)
	RemoveDNSServer(context.Context, *RemoveDNSServerRequest) (*RemoveDNSServerResponse, error)
	mustEmbedUnimplementedHostServiceServer()
}

// UnimplementedHostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (UnimplementedHostServiceServer) SetHostname(context.Context, *SetHostnameRequest) (*SetHostnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostname not implemented")
}
func (UnimplementedHostServiceServer) GetDNSServers(context.Context, *GetDNSServersRequest) (*GetDNSServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDNSServers not implemented")
}
func (UnimplementedHostServiceServer) AddDNSServer(context.Context, *AddDNSServerRequest) (*AddDNSServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDNSServer not implemented")
}
func (UnimplementedHostServiceServer) RemoveDNSServer(context.Context, *RemoveDNSServerRequest) (*RemoveDNSServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDNSServer not implemented")
}
func (UnimplementedHostServiceServer) mustEmbedUnimplementedHostServiceServer() {}

// UnsafeHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServiceServer will
// result in compilation errors.
type UnsafeHostServiceServer interface {
	mustEmbedUnimplementedHostServiceServer()
}

func RegisterHostServiceServer(s grpc.ServiceRegistrar, srv HostServiceServer) {
	s.RegisterService(&HostService_ServiceDesc, srv)
}

func _HostService_SetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).SetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostService_SetHostname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).SetHostname(ctx, req.(*SetHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_GetDNSServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDNSServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).GetDNSServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostService_GetDNSServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).GetDNSServers(ctx, req.(*GetDNSServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_AddDNSServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDNSServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).AddDNSServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostService_AddDNSServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).AddDNSServer(ctx, req.(*AddDNSServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostService_RemoveDNSServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDNSServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).RemoveDNSServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostService_RemoveDNSServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).RemoveDNSServer(ctx, req.(*RemoveDNSServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostService_ServiceDesc is the grpc.ServiceDesc for HostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHostname",
			Handler:    _HostService_SetHostname_Handler,
		},
		{
			MethodName: "GetDNSServers",
			Handler:    _HostService_GetDNSServers_Handler,
		},
		{
			MethodName: "AddDNSServer",
			Handler:    _HostService_AddDNSServer_Handler,
		},
		{
			MethodName: "RemoveDNSServer",
			Handler:    _HostService_RemoveDNSServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/host.proto",
}
