package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/s-prophet/internal/customer/service"
	customerv1 "github.com/materials-resources/s-prophet/internal/grpc/customer"
	"github.com/materials-resources/s-prophet/internal/grpc/customer/customerconnect"
)

func NewCustomerServiceHandler(svc service.CustomerService) *CustomerServiceHandler {
	return &CustomerServiceHandler{
		service: svc,
	}
}

type CustomerServiceHandler struct {
	service service.CustomerService
}

func (c CustomerServiceHandler) GetRecentPurchasesByBranch(ctx context.Context, req *connect.Request[customerv1.GetRecentPurchasesByBranchRequest]) (*connect.Response[customerv1.GetRecentPurchasesByBranchResponse], error) {
	res, err := c.service.GetRecentPurchasesByBranch(ctx, req.Msg)
	return connect.NewResponse(res), err
}

func (c CustomerServiceHandler) GetBranch(ctx context.Context, req *connect.Request[customerv1.GetBranchRequest]) (
	*connect.Response[customerv1.GetBranchResponse], error) {
	res, err := c.service.GetBranch(ctx, req.Msg)
	return connect.NewResponse(res), err
}

func (c CustomerServiceHandler) GetBranchesForContact(ctx context.Context, req *connect.Request[customerv1.GetBranchesForContactRequest]) (*connect.Response[customerv1.GetBranchesForContactResponse], error) {
	res, err := c.service.GetBranchesForContact(ctx, req.Msg)
	return connect.NewResponse(res), err
}

var _ customerconnect.CustomerServiceHandler = (*CustomerServiceHandler)(nil)
