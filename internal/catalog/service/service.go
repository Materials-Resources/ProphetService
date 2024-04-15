package service

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type CatalogService interface {
	// GetProduct returns a Product by the given UID
	GetProduct(ctx context.Context, uid int32) (domain.Product, error)

	// ListProductGroup returns a list of ProductGroup
	ListProductGroup(ctx context.Context) ([]*domain.ProductGroup, error)
	// GetProductGroup returns ProductGroup by the given SN
	GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, error)
	// UpdateProductGroup updates a ProductGroup
	UpdateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
	// CreateProductGroup creates a new product group
	CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
}
