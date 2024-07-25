// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: withdraw.proto

package withdraw

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

// WithdrawClient is the client API for Withdraw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WithdrawClient interface {
	FindAddressByCoinId(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*AddressSimpleList, error)
	SendCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*NoRes, error)
	WithdrawCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*NoRes, error)
	WithdrawRecord(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*RecordList, error)
}

type withdrawClient struct {
	cc grpc.ClientConnInterface
}

func NewWithdrawClient(cc grpc.ClientConnInterface) WithdrawClient {
	return &withdrawClient{cc}
}

func (c *withdrawClient) FindAddressByCoinId(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*AddressSimpleList, error) {
	out := new(AddressSimpleList)
	err := c.cc.Invoke(ctx, "/withdraw.Withdraw/findAddressByCoinId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawClient) SendCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*NoRes, error) {
	out := new(NoRes)
	err := c.cc.Invoke(ctx, "/withdraw.Withdraw/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawClient) WithdrawCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*NoRes, error) {
	out := new(NoRes)
	err := c.cc.Invoke(ctx, "/withdraw.Withdraw/WithdrawCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawClient) WithdrawRecord(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*RecordList, error) {
	out := new(RecordList)
	err := c.cc.Invoke(ctx, "/withdraw.Withdraw/WithdrawRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawServer is the server API for Withdraw service.
// All implementations must embed UnimplementedWithdrawServer
// for forward compatibility
type WithdrawServer interface {
	FindAddressByCoinId(context.Context, *WithdrawReq) (*AddressSimpleList, error)
	SendCode(context.Context, *WithdrawReq) (*NoRes, error)
	WithdrawCode(context.Context, *WithdrawReq) (*NoRes, error)
	WithdrawRecord(context.Context, *WithdrawReq) (*RecordList, error)
	mustEmbedUnimplementedWithdrawServer()
}

// UnimplementedWithdrawServer must be embedded to have forward compatible implementations.
type UnimplementedWithdrawServer struct {
}

func (UnimplementedWithdrawServer) FindAddressByCoinId(context.Context, *WithdrawReq) (*AddressSimpleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAddressByCoinId not implemented")
}
func (UnimplementedWithdrawServer) SendCode(context.Context, *WithdrawReq) (*NoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedWithdrawServer) WithdrawCode(context.Context, *WithdrawReq) (*NoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawCode not implemented")
}
func (UnimplementedWithdrawServer) WithdrawRecord(context.Context, *WithdrawReq) (*RecordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawRecord not implemented")
}
func (UnimplementedWithdrawServer) mustEmbedUnimplementedWithdrawServer() {}

// UnsafeWithdrawServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WithdrawServer will
// result in compilation errors.
type UnsafeWithdrawServer interface {
	mustEmbedUnimplementedWithdrawServer()
}

func RegisterWithdrawServer(s grpc.ServiceRegistrar, srv WithdrawServer) {
	s.RegisterService(&Withdraw_ServiceDesc, srv)
}

func _Withdraw_FindAddressByCoinId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServer).FindAddressByCoinId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/withdraw.Withdraw/findAddressByCoinId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServer).FindAddressByCoinId(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Withdraw_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/withdraw.Withdraw/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServer).SendCode(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Withdraw_WithdrawCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServer).WithdrawCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/withdraw.Withdraw/WithdrawCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServer).WithdrawCode(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Withdraw_WithdrawRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServer).WithdrawRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/withdraw.Withdraw/WithdrawRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServer).WithdrawRecord(ctx, req.(*WithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Withdraw_ServiceDesc is the grpc.ServiceDesc for Withdraw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Withdraw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "withdraw.Withdraw",
	HandlerType: (*WithdrawServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAddressByCoinId",
			Handler:    _Withdraw_FindAddressByCoinId_Handler,
		},
		{
			MethodName: "SendCode",
			Handler:    _Withdraw_SendCode_Handler,
		},
		{
			MethodName: "WithdrawCode",
			Handler:    _Withdraw_WithdrawCode_Handler,
		},
		{
			MethodName: "WithdrawRecord",
			Handler:    _Withdraw_WithdrawRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "withdraw.proto",
}
