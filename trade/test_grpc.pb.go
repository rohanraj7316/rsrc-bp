// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: test.proto

package trade

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

// TradingMSClient is the client API for TradingMS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingMSClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type tradingMSClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingMSClient(cc grpc.ClientConnInterface) TradingMSClient {
	return &tradingMSClient{cc}
}

func (c *tradingMSClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/TradingMS/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingMSServer is the server API for TradingMS service.
// All implementations must embed UnimplementedTradingMSServer
// for forward compatibility
type TradingMSServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedTradingMSServer()
}

// UnimplementedTradingMSServer must be embedded to have forward compatible implementations.
type UnimplementedTradingMSServer struct {
}

func (UnimplementedTradingMSServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedTradingMSServer) mustEmbedUnimplementedTradingMSServer() {}

// UnsafeTradingMSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradingMSServer will
// result in compilation errors.
type UnsafeTradingMSServer interface {
	mustEmbedUnimplementedTradingMSServer()
}

func RegisterTradingMSServer(s grpc.ServiceRegistrar, srv TradingMSServer) {
	s.RegisterService(&TradingMS_ServiceDesc, srv)
}

func _TradingMS_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingMSServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TradingMS/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingMSServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradingMS_ServiceDesc is the grpc.ServiceDesc for TradingMS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradingMS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TradingMS",
	HandlerType: (*TradingMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _TradingMS_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}