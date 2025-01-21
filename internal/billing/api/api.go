package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/service"
	billingv1 "github.com/materials-resources/s-prophet/internal/grpc/billing"
	"github.com/materials-resources/s-prophet/internal/grpc/billing/billingconnect"
)

func NewBillingServiceHandler(svc service.Billing) BillingServiceHandler {
	return BillingServiceHandler{service: svc}
}

type BillingServiceHandler struct {
	service service.Billing
}

func (h BillingServiceHandler) GetInvoicesByOrder(ctx context.Context, req *connect.Request[billingv1.GetInvoicesByOrderRequest]) (*connect.Response[billingv1.GetInvoicesByOrderResponse], error) {
	res, err := h.service.GetInvoicesByOrder(ctx, req.Msg)

	return connect.NewResponse(res), err
}

var _ billingconnect.BillingServiceHandler = (*BillingServiceHandler)(nil)
