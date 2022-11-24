// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: grpc/auction.proto

package grpc

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

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	Bid(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*Response, error)
	Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Amount, error)
	Crashed(ctx context.Context, in *ServerId, opts ...grpc.CallOption) (*Void, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) Bid(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.auction/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Amount, error) {
	out := new(Amount)
	err := c.cc.Invoke(ctx, "/grpc.auction/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Crashed(ctx context.Context, in *ServerId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.auction/crashed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility
type AuctionServer interface {
	Bid(context.Context, *Amount) (*Response, error)
	Result(context.Context, *Void) (*Amount, error)
	Crashed(context.Context, *ServerId) (*Void, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (UnimplementedAuctionServer) Bid(context.Context, *Amount) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServer) Result(context.Context, *Void) (*Amount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServer) Crashed(context.Context, *ServerId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crashed not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Amount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.auction/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Bid(ctx, req.(*Amount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.auction/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Result(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Crashed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Crashed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.auction/crashed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Crashed(ctx, req.(*ServerId))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bid",
			Handler:    _Auction_Bid_Handler,
		},
		{
			MethodName: "result",
			Handler:    _Auction_Result_Handler,
		},
		{
			MethodName: "crashed",
			Handler:    _Auction_Crashed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/auction.proto",
}

// ServerNodeClient is the client API for ServerNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerNodeClient interface {
	AddOperationToExecutionOrder(ctx context.Context, in *OperationId, opts ...grpc.CallOption) (*Void, error)
}

type serverNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewServerNodeClient(cc grpc.ClientConnInterface) ServerNodeClient {
	return &serverNodeClient{cc}
}

func (c *serverNodeClient) AddOperationToExecutionOrder(ctx context.Context, in *OperationId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.serverNode/addOperationToExecutionOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerNodeServer is the server API for ServerNode service.
// All implementations must embed UnimplementedServerNodeServer
// for forward compatibility
type ServerNodeServer interface {
	AddOperationToExecutionOrder(context.Context, *OperationId) (*Void, error)
	mustEmbedUnimplementedServerNodeServer()
}

// UnimplementedServerNodeServer must be embedded to have forward compatible implementations.
type UnimplementedServerNodeServer struct {
}

func (UnimplementedServerNodeServer) AddOperationToExecutionOrder(context.Context, *OperationId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOperationToExecutionOrder not implemented")
}
func (UnimplementedServerNodeServer) mustEmbedUnimplementedServerNodeServer() {}

// UnsafeServerNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerNodeServer will
// result in compilation errors.
type UnsafeServerNodeServer interface {
	mustEmbedUnimplementedServerNodeServer()
}

func RegisterServerNodeServer(s grpc.ServiceRegistrar, srv ServerNodeServer) {
	s.RegisterService(&ServerNode_ServiceDesc, srv)
}

func _ServerNode_AddOperationToExecutionOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).AddOperationToExecutionOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.serverNode/addOperationToExecutionOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).AddOperationToExecutionOrder(ctx, req.(*OperationId))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerNode_ServiceDesc is the grpc.ServiceDesc for ServerNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.serverNode",
	HandlerType: (*ServerNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addOperationToExecutionOrder",
			Handler:    _ServerNode_AddOperationToExecutionOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/auction.proto",
}
