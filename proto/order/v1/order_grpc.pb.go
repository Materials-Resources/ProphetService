// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: order/v1/order.proto

package order_v1alpha0

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
	OrderService_ListOrdersByCustomerBranch_FullMethodName = "/order.v1.OrderService/ListOrdersByCustomerBranch"
	OrderService_ListOrdersByCustomer_FullMethodName       = "/order.v1.OrderService/ListOrdersByCustomer"
	OrderService_ListOrdersByTaker_FullMethodName          = "/order.v1.OrderService/ListOrdersByTaker"
	OrderService_GetOrder_FullMethodName                   = "/order.v1.OrderService/GetOrder"
	OrderService_CreateOrder_FullMethodName                = "/order.v1.OrderService/CreateOrder"
	OrderService_CreateQuote_FullMethodName                = "/order.v1.OrderService/CreateQuote"
	OrderService_GetPickTicketById_FullMethodName          = "/order.v1.OrderService/GetPickTicketById"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	ListOrdersByCustomerBranch(ctx context.Context, in *ListOrdersByCustomerBranchRequest, opts ...grpc.CallOption) (*ListOrdersByCustomerBranchResponse, error)
	ListOrdersByCustomer(ctx context.Context, in *ListOrdersByCustomerRequest, opts ...grpc.CallOption) (*ListOrdersByCustomerResponse, error)
	ListOrdersByTaker(ctx context.Context, in *ListOrdersByTakerRequest, opts ...grpc.CallOption) (*ListOrdersByTakerResponse, error)
	// GetOrder returns the order details for a given order id
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*CreateQuoteResponse, error)
	GetPickTicketById(ctx context.Context, in *GetPickTicketByIdRequest, opts ...grpc.CallOption) (*GetPickTicketByIdResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) ListOrdersByCustomerBranch(ctx context.Context, in *ListOrdersByCustomerBranchRequest, opts ...grpc.CallOption) (*ListOrdersByCustomerBranchResponse, error) {
	out := new(ListOrdersByCustomerBranchResponse)
	err := c.cc.Invoke(ctx, OrderService_ListOrdersByCustomerBranch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrdersByCustomer(ctx context.Context, in *ListOrdersByCustomerRequest, opts ...grpc.CallOption) (*ListOrdersByCustomerResponse, error) {
	out := new(ListOrdersByCustomerResponse)
	err := c.cc.Invoke(ctx, OrderService_ListOrdersByCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrdersByTaker(ctx context.Context, in *ListOrdersByTakerRequest, opts ...grpc.CallOption) (*ListOrdersByTakerResponse, error) {
	out := new(ListOrdersByTakerResponse)
	err := c.cc.Invoke(ctx, OrderService_ListOrdersByTaker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*CreateQuoteResponse, error) {
	out := new(CreateQuoteResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetPickTicketById(ctx context.Context, in *GetPickTicketByIdRequest, opts ...grpc.CallOption) (*GetPickTicketByIdResponse, error) {
	out := new(GetPickTicketByIdResponse)
	err := c.cc.Invoke(ctx, OrderService_GetPickTicketById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	ListOrdersByCustomerBranch(context.Context, *ListOrdersByCustomerBranchRequest) (*ListOrdersByCustomerBranchResponse, error)
	ListOrdersByCustomer(context.Context, *ListOrdersByCustomerRequest) (*ListOrdersByCustomerResponse, error)
	ListOrdersByTaker(context.Context, *ListOrdersByTakerRequest) (*ListOrdersByTakerResponse, error)
	// GetOrder returns the order details for a given order id
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateQuote(context.Context, *CreateQuoteRequest) (*CreateQuoteResponse, error)
	GetPickTicketById(context.Context, *GetPickTicketByIdRequest) (*GetPickTicketByIdResponse, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) ListOrdersByCustomerBranch(context.Context, *ListOrdersByCustomerBranchRequest) (*ListOrdersByCustomerBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrdersByCustomerBranch not implemented")
}
func (UnimplementedOrderServiceServer) ListOrdersByCustomer(context.Context, *ListOrdersByCustomerRequest) (*ListOrdersByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrdersByCustomer not implemented")
}
func (UnimplementedOrderServiceServer) ListOrdersByTaker(context.Context, *ListOrdersByTakerRequest) (*ListOrdersByTakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrdersByTaker not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateQuote(context.Context, *CreateQuoteRequest) (*CreateQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuote not implemented")
}
func (UnimplementedOrderServiceServer) GetPickTicketById(context.Context, *GetPickTicketByIdRequest) (*GetPickTicketByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPickTicketById not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_ListOrdersByCustomerBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersByCustomerBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrdersByCustomerBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrdersByCustomerBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrdersByCustomerBranch(ctx, req.(*ListOrdersByCustomerBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrdersByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrdersByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrdersByCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrdersByCustomer(ctx, req.(*ListOrdersByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrdersByTaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersByTakerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrdersByTaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrdersByTaker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrdersByTaker(ctx, req.(*ListOrdersByTakerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateQuote(ctx, req.(*CreateQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetPickTicketById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPickTicketByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetPickTicketById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetPickTicketById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetPickTicketById(ctx, req.(*GetPickTicketByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.v1.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrdersByCustomerBranch",
			Handler:    _OrderService_ListOrdersByCustomerBranch_Handler,
		},
		{
			MethodName: "ListOrdersByCustomer",
			Handler:    _OrderService_ListOrdersByCustomer_Handler,
		},
		{
			MethodName: "ListOrdersByTaker",
			Handler:    _OrderService_ListOrdersByTaker_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "CreateQuote",
			Handler:    _OrderService_CreateQuote_Handler,
		},
		{
			MethodName: "GetPickTicketById",
			Handler:    _OrderService_GetPickTicketById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/v1/order.proto",
}
