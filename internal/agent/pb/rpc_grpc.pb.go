// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: rpc.proto

package pb

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
	AgentService_ListContainers_FullMethodName         = "/protobuf.AgentService/ListContainers"
	AgentService_FindContainer_FullMethodName          = "/protobuf.AgentService/FindContainer"
	AgentService_StreamLogs_FullMethodName             = "/protobuf.AgentService/StreamLogs"
	AgentService_LogsBetweenDates_FullMethodName       = "/protobuf.AgentService/LogsBetweenDates"
	AgentService_StreamRawBytes_FullMethodName         = "/protobuf.AgentService/StreamRawBytes"
	AgentService_StreamEvents_FullMethodName           = "/protobuf.AgentService/StreamEvents"
	AgentService_StreamStats_FullMethodName            = "/protobuf.AgentService/StreamStats"
	AgentService_StreamContainerStarted_FullMethodName = "/protobuf.AgentService/StreamContainerStarted"
	AgentService_HostInfo_FullMethodName               = "/protobuf.AgentService/HostInfo"
	AgentService_ContainerAction_FullMethodName        = "/protobuf.AgentService/ContainerAction"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	ListContainers(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error)
	FindContainer(ctx context.Context, in *FindContainerRequest, opts ...grpc.CallOption) (*FindContainerResponse, error)
	StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (AgentService_StreamLogsClient, error)
	LogsBetweenDates(ctx context.Context, in *LogsBetweenDatesRequest, opts ...grpc.CallOption) (AgentService_LogsBetweenDatesClient, error)
	StreamRawBytes(ctx context.Context, in *StreamRawBytesRequest, opts ...grpc.CallOption) (AgentService_StreamRawBytesClient, error)
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (AgentService_StreamEventsClient, error)
	StreamStats(ctx context.Context, in *StreamStatsRequest, opts ...grpc.CallOption) (AgentService_StreamStatsClient, error)
	StreamContainerStarted(ctx context.Context, in *StreamContainerStartedRequest, opts ...grpc.CallOption) (AgentService_StreamContainerStartedClient, error)
	HostInfo(ctx context.Context, in *HostInfoRequest, opts ...grpc.CallOption) (*HostInfoResponse, error)
	ContainerAction(ctx context.Context, in *ContainerActionRequest, opts ...grpc.CallOption) (*ContainerActionResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) ListContainers(ctx context.Context, in *ListContainersRequest, opts ...grpc.CallOption) (*ListContainersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListContainersResponse)
	err := c.cc.Invoke(ctx, AgentService_ListContainers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) FindContainer(ctx context.Context, in *FindContainerRequest, opts ...grpc.CallOption) (*FindContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindContainerResponse)
	err := c.cc.Invoke(ctx, AgentService_FindContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (AgentService_StreamLogsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], AgentService_StreamLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceStreamLogsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_StreamLogsClient interface {
	Recv() (*StreamLogsResponse, error)
	grpc.ClientStream
}

type agentServiceStreamLogsClient struct {
	grpc.ClientStream
}

func (x *agentServiceStreamLogsClient) Recv() (*StreamLogsResponse, error) {
	m := new(StreamLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) LogsBetweenDates(ctx context.Context, in *LogsBetweenDatesRequest, opts ...grpc.CallOption) (AgentService_LogsBetweenDatesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[1], AgentService_LogsBetweenDates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceLogsBetweenDatesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_LogsBetweenDatesClient interface {
	Recv() (*StreamLogsResponse, error)
	grpc.ClientStream
}

type agentServiceLogsBetweenDatesClient struct {
	grpc.ClientStream
}

func (x *agentServiceLogsBetweenDatesClient) Recv() (*StreamLogsResponse, error) {
	m := new(StreamLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) StreamRawBytes(ctx context.Context, in *StreamRawBytesRequest, opts ...grpc.CallOption) (AgentService_StreamRawBytesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[2], AgentService_StreamRawBytes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceStreamRawBytesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_StreamRawBytesClient interface {
	Recv() (*StreamRawBytesResponse, error)
	grpc.ClientStream
}

type agentServiceStreamRawBytesClient struct {
	grpc.ClientStream
}

func (x *agentServiceStreamRawBytesClient) Recv() (*StreamRawBytesResponse, error) {
	m := new(StreamRawBytesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (AgentService_StreamEventsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[3], AgentService_StreamEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceStreamEventsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_StreamEventsClient interface {
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type agentServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *agentServiceStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) StreamStats(ctx context.Context, in *StreamStatsRequest, opts ...grpc.CallOption) (AgentService_StreamStatsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[4], AgentService_StreamStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceStreamStatsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_StreamStatsClient interface {
	Recv() (*StreamStatsResponse, error)
	grpc.ClientStream
}

type agentServiceStreamStatsClient struct {
	grpc.ClientStream
}

func (x *agentServiceStreamStatsClient) Recv() (*StreamStatsResponse, error) {
	m := new(StreamStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) StreamContainerStarted(ctx context.Context, in *StreamContainerStartedRequest, opts ...grpc.CallOption) (AgentService_StreamContainerStartedClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[5], AgentService_StreamContainerStarted_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceStreamContainerStartedClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_StreamContainerStartedClient interface {
	Recv() (*StreamContainerStartedResponse, error)
	grpc.ClientStream
}

type agentServiceStreamContainerStartedClient struct {
	grpc.ClientStream
}

func (x *agentServiceStreamContainerStartedClient) Recv() (*StreamContainerStartedResponse, error) {
	m := new(StreamContainerStartedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) HostInfo(ctx context.Context, in *HostInfoRequest, opts ...grpc.CallOption) (*HostInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HostInfoResponse)
	err := c.cc.Invoke(ctx, AgentService_HostInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ContainerAction(ctx context.Context, in *ContainerActionRequest, opts ...grpc.CallOption) (*ContainerActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContainerActionResponse)
	err := c.cc.Invoke(ctx, AgentService_ContainerAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	ListContainers(context.Context, *ListContainersRequest) (*ListContainersResponse, error)
	FindContainer(context.Context, *FindContainerRequest) (*FindContainerResponse, error)
	StreamLogs(*StreamLogsRequest, AgentService_StreamLogsServer) error
	LogsBetweenDates(*LogsBetweenDatesRequest, AgentService_LogsBetweenDatesServer) error
	StreamRawBytes(*StreamRawBytesRequest, AgentService_StreamRawBytesServer) error
	StreamEvents(*StreamEventsRequest, AgentService_StreamEventsServer) error
	StreamStats(*StreamStatsRequest, AgentService_StreamStatsServer) error
	StreamContainerStarted(*StreamContainerStartedRequest, AgentService_StreamContainerStartedServer) error
	HostInfo(context.Context, *HostInfoRequest) (*HostInfoResponse, error)
	ContainerAction(context.Context, *ContainerActionRequest) (*ContainerActionResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) ListContainers(context.Context, *ListContainersRequest) (*ListContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContainers not implemented")
}
func (UnimplementedAgentServiceServer) FindContainer(context.Context, *FindContainerRequest) (*FindContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContainer not implemented")
}
func (UnimplementedAgentServiceServer) StreamLogs(*StreamLogsRequest, AgentService_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedAgentServiceServer) LogsBetweenDates(*LogsBetweenDatesRequest, AgentService_LogsBetweenDatesServer) error {
	return status.Errorf(codes.Unimplemented, "method LogsBetweenDates not implemented")
}
func (UnimplementedAgentServiceServer) StreamRawBytes(*StreamRawBytesRequest, AgentService_StreamRawBytesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRawBytes not implemented")
}
func (UnimplementedAgentServiceServer) StreamEvents(*StreamEventsRequest, AgentService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedAgentServiceServer) StreamStats(*StreamStatsRequest, AgentService_StreamStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStats not implemented")
}
func (UnimplementedAgentServiceServer) StreamContainerStarted(*StreamContainerStartedRequest, AgentService_StreamContainerStartedServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamContainerStarted not implemented")
}
func (UnimplementedAgentServiceServer) HostInfo(context.Context, *HostInfoRequest) (*HostInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostInfo not implemented")
}
func (UnimplementedAgentServiceServer) ContainerAction(context.Context, *ContainerActionRequest) (*ContainerActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerAction not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_ListContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ListContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListContainers(ctx, req.(*ListContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_FindContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).FindContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_FindContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).FindContainer(ctx, req.(*FindContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamLogs(m, &agentServiceStreamLogsServer{ServerStream: stream})
}

type AgentService_StreamLogsServer interface {
	Send(*StreamLogsResponse) error
	grpc.ServerStream
}

type agentServiceStreamLogsServer struct {
	grpc.ServerStream
}

func (x *agentServiceStreamLogsServer) Send(m *StreamLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_LogsBetweenDates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsBetweenDatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).LogsBetweenDates(m, &agentServiceLogsBetweenDatesServer{ServerStream: stream})
}

type AgentService_LogsBetweenDatesServer interface {
	Send(*StreamLogsResponse) error
	grpc.ServerStream
}

type agentServiceLogsBetweenDatesServer struct {
	grpc.ServerStream
}

func (x *agentServiceLogsBetweenDatesServer) Send(m *StreamLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_StreamRawBytes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRawBytesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamRawBytes(m, &agentServiceStreamRawBytesServer{ServerStream: stream})
}

type AgentService_StreamRawBytesServer interface {
	Send(*StreamRawBytesResponse) error
	grpc.ServerStream
}

type agentServiceStreamRawBytesServer struct {
	grpc.ServerStream
}

func (x *agentServiceStreamRawBytesServer) Send(m *StreamRawBytesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamEvents(m, &agentServiceStreamEventsServer{ServerStream: stream})
}

type AgentService_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	grpc.ServerStream
}

type agentServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *agentServiceStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_StreamStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamStats(m, &agentServiceStreamStatsServer{ServerStream: stream})
}

type AgentService_StreamStatsServer interface {
	Send(*StreamStatsResponse) error
	grpc.ServerStream
}

type agentServiceStreamStatsServer struct {
	grpc.ServerStream
}

func (x *agentServiceStreamStatsServer) Send(m *StreamStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_StreamContainerStarted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamContainerStartedRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).StreamContainerStarted(m, &agentServiceStreamContainerStartedServer{ServerStream: stream})
}

type AgentService_StreamContainerStartedServer interface {
	Send(*StreamContainerStartedResponse) error
	grpc.ServerStream
}

type agentServiceStreamContainerStartedServer struct {
	grpc.ServerStream
}

func (x *agentServiceStreamContainerStartedServer) Send(m *StreamContainerStartedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_HostInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).HostInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_HostInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).HostInfo(ctx, req.(*HostInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ContainerAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ContainerAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ContainerAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ContainerAction(ctx, req.(*ContainerActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContainers",
			Handler:    _AgentService_ListContainers_Handler,
		},
		{
			MethodName: "FindContainer",
			Handler:    _AgentService_FindContainer_Handler,
		},
		{
			MethodName: "HostInfo",
			Handler:    _AgentService_HostInfo_Handler,
		},
		{
			MethodName: "ContainerAction",
			Handler:    _AgentService_ContainerAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLogs",
			Handler:       _AgentService_StreamLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogsBetweenDates",
			Handler:       _AgentService_LogsBetweenDates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamRawBytes",
			Handler:       _AgentService_StreamRawBytes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEvents",
			Handler:       _AgentService_StreamEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamStats",
			Handler:       _AgentService_StreamStats_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamContainerStarted",
			Handler:       _AgentService_StreamContainerStarted_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
