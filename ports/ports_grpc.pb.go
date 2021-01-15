// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ports

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	// Sends a greeting
	AddPort(ctx context.Context, in *AddPortRequest, opts ...grpc.CallOption) (*AddPortReply, error)
	// Sends another greeting
	GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*GetPortReply, error)
}

type portDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainServiceClient(cc grpc.ClientConnInterface) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) AddPort(ctx context.Context, in *AddPortRequest, opts ...grpc.CallOption) (*AddPortReply, error) {
	out := new(AddPortReply)
	err := c.cc.Invoke(ctx, "/PortDomainService/AddPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*GetPortReply, error) {
	out := new(GetPortReply)
	err := c.cc.Invoke(ctx, "/PortDomainService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
// All implementations must embed UnimplementedPortDomainServiceServer
// for forward compatibility
type PortDomainServiceServer interface {
	// Sends a greeting
	AddPort(context.Context, *AddPortRequest) (*AddPortReply, error)
	// Sends another greeting
	GetPort(context.Context, *PortRequest) (*GetPortReply, error)
	mustEmbedUnimplementedPortDomainServiceServer()
}

// UnimplementedPortDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServiceServer struct {
}

func (UnimplementedPortDomainServiceServer) AddPort(context.Context, *AddPortRequest) (*AddPortReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPort not implemented")
}
func (UnimplementedPortDomainServiceServer) GetPort(context.Context, *PortRequest) (*GetPortReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortDomainServiceServer) mustEmbedUnimplementedPortDomainServiceServer() {}

// UnsafePortDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServiceServer will
// result in compilation errors.
type UnsafePortDomainServiceServer interface {
	mustEmbedUnimplementedPortDomainServiceServer()
}

func RegisterPortDomainServiceServer(s grpc.ServiceRegistrar, srv PortDomainServiceServer) {
	s.RegisterService(&_PortDomainService_serviceDesc, srv)
}

func _PortDomainService_AddPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).AddPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomainService/AddPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).AddPort(ctx, req.(*AddPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortDomainService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).GetPort(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortDomainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPort",
			Handler:    _PortDomainService_AddPort_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _PortDomainService_GetPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports.proto",
}