// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: analytics/analytics.proto

package pb

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
	Analytics_GetCharts_FullMethodName            = "/analytics.Analytics/GetCharts"
	Analytics_GetDescriptionStream_FullMethodName = "/analytics.Analytics/GetDescriptionStream"
)

// AnalyticsClient is the client API for Analytics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsClient interface {
	GetCharts(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Report, error)
	GetDescriptionStream(ctx context.Context, in *Params, opts ...grpc.CallOption) (Analytics_GetDescriptionStreamClient, error)
}

type analyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsClient(cc grpc.ClientConnInterface) AnalyticsClient {
	return &analyticsClient{cc}
}

func (c *analyticsClient) GetCharts(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, Analytics_GetCharts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsClient) GetDescriptionStream(ctx context.Context, in *Params, opts ...grpc.CallOption) (Analytics_GetDescriptionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Analytics_ServiceDesc.Streams[0], Analytics_GetDescriptionStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &analyticsGetDescriptionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Analytics_GetDescriptionStreamClient interface {
	Recv() (*Report, error)
	grpc.ClientStream
}

type analyticsGetDescriptionStreamClient struct {
	grpc.ClientStream
}

func (x *analyticsGetDescriptionStreamClient) Recv() (*Report, error) {
	m := new(Report)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticsServer is the server API for Analytics service.
// All implementations must embed UnimplementedAnalyticsServer
// for forward compatibility
type AnalyticsServer interface {
	GetCharts(context.Context, *Params) (*Report, error)
	GetDescriptionStream(*Params, Analytics_GetDescriptionStreamServer) error
	mustEmbedUnimplementedAnalyticsServer()
}

// UnimplementedAnalyticsServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServer struct {
}

func (UnimplementedAnalyticsServer) GetCharts(context.Context, *Params) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharts not implemented")
}
func (UnimplementedAnalyticsServer) GetDescriptionStream(*Params, Analytics_GetDescriptionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDescriptionStream not implemented")
}
func (UnimplementedAnalyticsServer) mustEmbedUnimplementedAnalyticsServer() {}

// UnsafeAnalyticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServer will
// result in compilation errors.
type UnsafeAnalyticsServer interface {
	mustEmbedUnimplementedAnalyticsServer()
}

func RegisterAnalyticsServer(s grpc.ServiceRegistrar, srv AnalyticsServer) {
	s.RegisterService(&Analytics_ServiceDesc, srv)
}

func _Analytics_GetCharts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).GetCharts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analytics_GetCharts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).GetCharts(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analytics_GetDescriptionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Params)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnalyticsServer).GetDescriptionStream(m, &analyticsGetDescriptionStreamServer{stream})
}

type Analytics_GetDescriptionStreamServer interface {
	Send(*Report) error
	grpc.ServerStream
}

type analyticsGetDescriptionStreamServer struct {
	grpc.ServerStream
}

func (x *analyticsGetDescriptionStreamServer) Send(m *Report) error {
	return x.ServerStream.SendMsg(m)
}

// Analytics_ServiceDesc is the grpc.ServiceDesc for Analytics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analytics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.Analytics",
	HandlerType: (*AnalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharts",
			Handler:    _Analytics_GetCharts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDescriptionStream",
			Handler:       _Analytics_GetDescriptionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "analytics/analytics.proto",
}