package service

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/billing/repository"
	billingv1 "github.com/materials-resources/s-prophet/internal/grpc/billing"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrNotAuthorized    = errors.New("not authorized")
	ErrResourceNotFound = errors.New("resource not found")
)

type BillingService interface {
	GetInvoicesByOrder(ctx context.Context, req *billingv1.GetInvoicesByOrderRequest) (*billingv1.
		GetInvoicesByOrderResponse, error)
}

func NewBillingService(a *app.App) Billing {
	return Billing{repository: repository.NewRepository(a.GetDB())}
}

type Billing struct {
	repository *repository.Repository
}

func (s Billing) GetInvoicesByOrder(ctx context.Context, req *billingv1.GetInvoicesByOrderRequest) (*billingv1.
	GetInvoicesByOrderResponse, error) {
	invoices, err := s.repository.Invoice.GetByOrder(ctx, req.OrderId)
	if err != nil {
		return &billingv1.GetInvoicesByOrderResponse{}, err
	}

	response := &billingv1.GetInvoicesByOrderResponse{}
	for _, invoice := range invoices {

		response.Invoices = append(response.Invoices, &billingv1.InvoiceSummary{
			Id:           invoice.Id,
			OrderId:      invoice.OrderId,
			AmountPaid:   invoice.AmountPaid,
			Total:        invoice.Total,
			DateInvoiced: timestamppb.New(invoice.DateInvoiced),
		})

	}

	return response, err
}

var _ BillingService = (*Billing)(nil)
