package grpc

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/ports"
	billingV1 "github.com/materials-resources/s-prophet/proto/billing/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillingService struct {
	api ports.ApiPort
	billingV1.UnsafeBillingServiceServer
}

func (s *BillingService) GetInvoicesByOrder(ctx context.Context, request *billingV1.GetInvoicesByOrderRequest) (*billingV1.GetInvoicesByOrderResponse, error) {
	invoices, err := s.api.GetInvoicesByOrderID(request.OrderId)
	if err != nil {
		return &billingV1.GetInvoicesByOrderResponse{}, status.Error(codes.Internal, "a server error occurred")

	}

	billingInvoices := make([]*billingV1.SimplifiedInvoice, len(invoices))
	for i, invoice := range invoices {
		billingInvoices[i] = &billingV1.SimplifiedInvoice{
			Id:         invoice.Id,
			OrderId:    invoice.OrderId,
			Total:      invoice.Total,
			AmountPaid: invoice.AmountPaid,
			CreatedAt:  invoice.CreatedAt.Format("2006-01-02 15:04:05"),
		}

	}
	return &billingV1.GetInvoicesByOrderResponse{Invoices: billingInvoices}, nil
}

func (s *BillingService) registerService(server *grpc.Server) {
	billingV1.RegisterBillingServiceServer(server, s)
}
