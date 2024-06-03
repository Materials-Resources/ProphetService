package ports

import "github.com/materials-resources/s-prophet/internal/billing/core/domain"

type ApiPort interface {
	GetInvoicesByOrderID(orderID string) ([]*domain.Invoice, error)
}
