// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FoobarClient is the client API for Foobar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoobarClient interface {
	AFoo(ctx context.Context, in *AFooRequest, opts ...grpc.CallOption) (*AFooResponse, error)
	Baz(ctx context.Context, in *BazRequest, opts ...grpc.CallOption) (*BazResponse, error)
	Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
	Empty(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	WellKnown(ctx context.Context, in *WellKnownRequest, opts ...grpc.CallOption) (*WellKnownResponse, error)
}

type foobarClient struct {
	cc grpc.ClientConnInterface
}

func NewFoobarClient(cc grpc.ClientConnInterface) FoobarClient {
	return &foobarClient{cc}
}

func (c *foobarClient) AFoo(ctx context.Context, in *AFooRequest, opts ...grpc.CallOption) (*AFooResponse, error) {
	out := new(AFooResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/AFoo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarClient) Baz(ctx context.Context, in *BazRequest, opts ...grpc.CallOption) (*BazResponse, error) {
	out := new(BazResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/Baz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarClient) Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/Bar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarClient) Empty(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/Empty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foobarClient) WellKnown(ctx context.Context, in *WellKnownRequest, opts ...grpc.CallOption) (*WellKnownResponse, error) {
	out := new(WellKnownResponse)
	err := c.cc.Invoke(ctx, "/wombat.v1.Foobar/WellKnown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoobarServer is the server API for Foobar service.
// All implementations must embed UnimplementedFoobarServer
// for forward compatibility
type FoobarServer interface {
	AFoo(context.Context, *AFooRequest) (*AFooResponse, error)
	Baz(context.Context, *BazRequest) (*BazResponse, error)
	Bar(context.Context, *BarRequest) (*BarResponse, error)
	Foo(context.Context, *FooRequest) (*FooResponse, error)
	Empty(context.Context, *EmptyRequest) (*EmptyResponse, error)
	WellKnown(context.Context, *WellKnownRequest) (*WellKnownResponse, error)
	mustEmbedUnimplementedFoobarServer()
}

// UnimplementedFoobarServer must be embedded to have forward compatible implementations.
type UnimplementedFoobarServer struct {
}

func (UnimplementedFoobarServer) AFoo(context.Context, *AFooRequest) (*AFooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AFoo not implemented")
}
func (UnimplementedFoobarServer) Baz(context.Context, *BazRequest) (*BazResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Baz not implemented")
}
func (UnimplementedFoobarServer) Bar(context.Context, *BarRequest) (*BarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bar not implemented")
}
func (UnimplementedFoobarServer) Foo(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedFoobarServer) Empty(context.Context, *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Empty not implemented")
}
func (UnimplementedFoobarServer) WellKnown(context.Context, *WellKnownRequest) (*WellKnownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WellKnown not implemented")
}
func (UnimplementedFoobarServer) mustEmbedUnimplementedFoobarServer() {}

// UnsafeFoobarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoobarServer will
// result in compilation errors.
type UnsafeFoobarServer interface {
	mustEmbedUnimplementedFoobarServer()
}

func RegisterFoobarServer(s grpc.ServiceRegistrar, srv FoobarServer) {
	s.RegisterService(&_Foobar_serviceDesc, srv)
}

func _Foobar_AFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AFooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).AFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/AFoo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).AFoo(ctx, req.(*AFooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foobar_Baz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BazRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).Baz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/Baz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).Baz(ctx, req.(*BazRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foobar_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).Bar(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foobar_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foobar_Empty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).Empty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/Empty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).Empty(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foobar_WellKnown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WellKnownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoobarServer).WellKnown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wombat.v1.Foobar/WellKnown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoobarServer).WellKnown(ctx, req.(*WellKnownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foobar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wombat.v1.Foobar",
	HandlerType: (*FoobarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AFoo",
			Handler:    _Foobar_AFoo_Handler,
		},
		{
			MethodName: "Baz",
			Handler:    _Foobar_Baz_Handler,
		},
		{
			MethodName: "Bar",
			Handler:    _Foobar_Bar_Handler,
		},
		{
			MethodName: "Foo",
			Handler:    _Foobar_Foo_Handler,
		},
		{
			MethodName: "Empty",
			Handler:    _Foobar_Empty_Handler,
		},
		{
			MethodName: "WellKnown",
			Handler:    _Foobar_WellKnown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foobar.proto",
}
