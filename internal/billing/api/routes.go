package api

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/service"
	billingpb "github.com/materials-resources/s-prophet/proto/billing/v1"
)

var _ billingpb.BillingServiceServer = Billing{}

func NewBillingRoutes(service *service.Billing) Billing {
	return Billing{service: service}
}

type Billing struct {
	service *service.Billing
}

func (r Billing) GetInvoicesByOrder(ctx context.Context, request *billingpb.GetInvoicesByOrderRequest) (*billingpb.GetInvoicesByOrderResponse, error) {
	invoices, err := r.service.GetInvoicesByOrder(ctx, request.OrderId)

	if err != nil {
		return &billingpb.GetInvoicesByOrderResponse{}, err
	}

	res := &billingpb.GetInvoicesByOrderResponse{
		Invoices: make([]*billingpb.SimplifiedInvoice, len(invoices)),
	}

	for i, invoice := range invoices {
		res.Invoices[i] = &billingpb.SimplifiedInvoice{
			Id:         invoice.Id,
			OrderId:    invoice.OrderId,
			CreatedAt:  invoice.CreatedAt.Format(""),
			Total:      invoice.Total,
			AmountPaid: invoice.AmountPaid,
		}
	}

	//TODO implement me
	panic("implement me")
}
