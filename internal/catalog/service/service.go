package service

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type CatalogService interface {
	// GetProduct returns a Product by the given UID
	GetProduct(ctx context.Context, uid int32) (domain.Product, error)
	// GetProductBySupplierPartNumber returns a Product by the given part number and supplier
	GetProductBySupplierPartNumber(ctx context.Context, partNumber string, supplier float64) (domain.Product, error)
	// ListProduct returns a list of Product
	ListProduct(ctx context.Context, filter domain.Filter) ([]*domain.Product, error)
	// UpdateProduct updates a Product for specified locations
	UpdateProduct(ctx context.Context, product *domain.Product, locations []float64) error
	// DeleteProduct deletes a Product
	DeleteProduct(ctx context.Context, uid int32) error

	UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error

	// ListProductGroup returns a list of ProductGroup
	ListProductGroup(ctx context.Context) ([]*domain.ProductGroup, error)
	// GetProductGroup returns ProductGroup by the given SN
	GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, []*domain.Product, error)
	// UpdateProductGroup updates a ProductGroup
	UpdateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
	// CreateProductGroup creates a new product group
	CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
}

type Producer interface {
	PublishUpdateProduct(ctx context.Context, product *domain.Product) error
}
