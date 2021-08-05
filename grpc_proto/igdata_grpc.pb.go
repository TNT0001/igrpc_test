// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: igdata.proto

package igrpcproto

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

// IgdataClient is the client API for Igdata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IgdataClient interface {
	// Ping Service
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// ESQuery Service
	EsQuery(ctx context.Context, in *ESQuery, opts ...grpc.CallOption) (*ESResponse, error)
	// SQLQuery Service
	SQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SQLQueryResponse, error)
	// Stream ESQuery service
	EsQueryStream(ctx context.Context, opts ...grpc.CallOption) (Igdata_EsQueryStreamClient, error)
	// Stream SQLQuery service
	SQLQueryStream(ctx context.Context, opts ...grpc.CallOption) (Igdata_SQLQueryStreamClient, error)
}

type igdataClient struct {
	cc grpc.ClientConnInterface
}

func NewIgdataClient(cc grpc.ClientConnInterface) IgdataClient {
	return &igdataClient{cc}
}

func (c *igdataClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/igrpcproto.igdata/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *igdataClient) EsQuery(ctx context.Context, in *ESQuery, opts ...grpc.CallOption) (*ESResponse, error) {
	out := new(ESResponse)
	err := c.cc.Invoke(ctx, "/igrpcproto.igdata/EsQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *igdataClient) SQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SQLQueryResponse, error) {
	out := new(SQLQueryResponse)
	err := c.cc.Invoke(ctx, "/igrpcproto.igdata/SQLQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *igdataClient) EsQueryStream(ctx context.Context, opts ...grpc.CallOption) (Igdata_EsQueryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Igdata_ServiceDesc.Streams[0], "/igrpcproto.igdata/EsQueryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &igdataEsQueryStreamClient{stream}
	return x, nil
}

type Igdata_EsQueryStreamClient interface {
	Send(*ESQuery) error
	Recv() (*ESResponse, error)
	grpc.ClientStream
}

type igdataEsQueryStreamClient struct {
	grpc.ClientStream
}

func (x *igdataEsQueryStreamClient) Send(m *ESQuery) error {
	return x.ClientStream.SendMsg(m)
}

func (x *igdataEsQueryStreamClient) Recv() (*ESResponse, error) {
	m := new(ESResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *igdataClient) SQLQueryStream(ctx context.Context, opts ...grpc.CallOption) (Igdata_SQLQueryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Igdata_ServiceDesc.Streams[1], "/igrpcproto.igdata/SQLQueryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &igdataSQLQueryStreamClient{stream}
	return x, nil
}

type Igdata_SQLQueryStreamClient interface {
	Send(*SQLQueryRequest) error
	Recv() (*SQLQueryResponse, error)
	grpc.ClientStream
}

type igdataSQLQueryStreamClient struct {
	grpc.ClientStream
}

func (x *igdataSQLQueryStreamClient) Send(m *SQLQueryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *igdataSQLQueryStreamClient) Recv() (*SQLQueryResponse, error) {
	m := new(SQLQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IgdataServer is the server API for Igdata service.
// All implementations must embed UnimplementedIgdataServer
// for forward compatibility
type IgdataServer interface {
	// Ping Service
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// ESQuery Service
	EsQuery(context.Context, *ESQuery) (*ESResponse, error)
	// SQLQuery Service
	SQLQuery(context.Context, *SQLQueryRequest) (*SQLQueryResponse, error)
	// Stream ESQuery service
	EsQueryStream(Igdata_EsQueryStreamServer) error
	// Stream SQLQuery service
	SQLQueryStream(Igdata_SQLQueryStreamServer) error
	mustEmbedUnimplementedIgdataServer()
}

// UnimplementedIgdataServer must be embedded to have forward compatible implementations.
type UnimplementedIgdataServer struct {
}

func (UnimplementedIgdataServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedIgdataServer) EsQuery(context.Context, *ESQuery) (*ESResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsQuery not implemented")
}
func (UnimplementedIgdataServer) SQLQuery(context.Context, *SQLQueryRequest) (*SQLQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLQuery not implemented")
}
func (UnimplementedIgdataServer) EsQueryStream(Igdata_EsQueryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EsQueryStream not implemented")
}
func (UnimplementedIgdataServer) SQLQueryStream(Igdata_SQLQueryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SQLQueryStream not implemented")
}
func (UnimplementedIgdataServer) mustEmbedUnimplementedIgdataServer() {}

// UnsafeIgdataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IgdataServer will
// result in compilation errors.
type UnsafeIgdataServer interface {
	mustEmbedUnimplementedIgdataServer()
}

func RegisterIgdataServer(s grpc.ServiceRegistrar, srv IgdataServer) {
	s.RegisterService(&Igdata_ServiceDesc, srv)
}

func _Igdata_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgdataServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpcproto.igdata/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgdataServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Igdata_EsQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ESQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgdataServer).EsQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpcproto.igdata/EsQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgdataServer).EsQuery(ctx, req.(*ESQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Igdata_SQLQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgdataServer).SQLQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpcproto.igdata/SQLQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgdataServer).SQLQuery(ctx, req.(*SQLQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Igdata_EsQueryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IgdataServer).EsQueryStream(&igdataEsQueryStreamServer{stream})
}

type Igdata_EsQueryStreamServer interface {
	Send(*ESResponse) error
	Recv() (*ESQuery, error)
	grpc.ServerStream
}

type igdataEsQueryStreamServer struct {
	grpc.ServerStream
}

func (x *igdataEsQueryStreamServer) Send(m *ESResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *igdataEsQueryStreamServer) Recv() (*ESQuery, error) {
	m := new(ESQuery)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Igdata_SQLQueryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IgdataServer).SQLQueryStream(&igdataSQLQueryStreamServer{stream})
}

type Igdata_SQLQueryStreamServer interface {
	Send(*SQLQueryResponse) error
	Recv() (*SQLQueryRequest, error)
	grpc.ServerStream
}

type igdataSQLQueryStreamServer struct {
	grpc.ServerStream
}

func (x *igdataSQLQueryStreamServer) Send(m *SQLQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *igdataSQLQueryStreamServer) Recv() (*SQLQueryRequest, error) {
	m := new(SQLQueryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Igdata_ServiceDesc is the grpc.ServiceDesc for Igdata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Igdata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "igrpcproto.igdata",
	HandlerType: (*IgdataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Igdata_Ping_Handler,
		},
		{
			MethodName: "EsQuery",
			Handler:    _Igdata_EsQuery_Handler,
		},
		{
			MethodName: "SQLQuery",
			Handler:    _Igdata_SQLQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EsQueryStream",
			Handler:       _Igdata_EsQueryStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SQLQueryStream",
			Handler:       _Igdata_SQLQueryStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "igdata.proto",
}