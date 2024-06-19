package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
)

type CatalogService interface {
	CustomerGetProduct(ctx context.Context, uid string) (*domain.Product, error)

	ClerkListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error)
	ClerkGetProduct(ctx context.Context, uid string) (*domain.Product, error)
	ClerkUpdateProduct(ctx context.Context, product *domain.Product) error
	ClerkDeleteProduct(ctx context.Context, uid string, deleteMode domain.DeleteMode) error

	ClerkListGroups(ctx context.Context) ([]*domain.ProductGroup, error)
	ClerkCreateGroup(ctx context.Context, productGroup *domain.ProductGroup) error
	ClerkReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error)
	ClerkUpdateGroup(ctx context.Context, productGroup *domain.ProductGroup) error

	// GetProductBySupplierPartNumber returns a Product by the given part number and supplier
	GetProductBySupplierPartNumber(ctx context.Context, partNumber string, supplier float64) (domain.Product, error)
	// DeleteProduct deletes a Product
	DeleteProduct(ctx context.Context, uid string) error

	UpdateGroup(ctx context.Context, productGroup *domain.ProductGroup) error

	UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error

	// SetPrimaryProductSupplier sets the primary product supplier
	SetPrimaryProductSupplier(ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error
}

type Producer interface {
	PublishUpdateProduct(ctx context.Context, product *domain.Product) error
}
