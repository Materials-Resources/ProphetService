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
	// GetBasicProductDetails returns a list of Product with basic details
	GetBasicProductDetails(ctx context.Context, uids []int32) ([]domain.Product, error)

	UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error

	// SetPrimaryProductSupplier sets the primary product supplier
	SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error

	// RequestDeleteProduct adds a request to deletes a Product
	RequestDeleteProduct(ctx context.Context, uid int32) error

	// ListProductGroup returns a list of ProductGroup
	ListProductGroup(ctx context.Context) ([]*domain.ProductGroup, error)
	// GetProductGroup returns ProductGroup by the given SN
	GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, []int32, error)
	// UpdateProductGroup updates a ProductGroup
	UpdateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
	// CreateProductGroup creates a new product group
	CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
}

type Producer interface {
	PublishUpdateProduct(ctx context.Context, product *domain.Product) error
}
