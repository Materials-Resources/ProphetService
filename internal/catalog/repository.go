package catalog

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

// Repository is the interface that wraps the basic CRUD operations for the catalog
type Repository interface {
	CreateProduct()
	// ListProducts by given filter
	ListProducts(filter *domain.ProductFilter) ([]*domain.Product, error)

	FilterProductByGroup(filter *domain.ProductFilter) ([]*domain.Product, error)
	// SelectProduct returns a product by id
	SelectProduct(id string) (*domain.Product, error)
	ListProduct(ctx context.Context, cursor int32, count int) (res []*domain.Product, nextCursor int, err error)
	ReadProductByGroup(id string) ([]*domain.Product, error)
	// UpdateProduct updates a product by id
	UpdateProduct(ctx context.Context, p *domain.ValidatedProduct)
	// DeleteProduct deletes a product by id
	DeleteProduct(ctx context.Context, id string) error

	// AddProductSupplier adds a product supplier to an existing product in the database
	AddProductSupplier(ctx context.Context, supplier *domain.ValidatedProductSupplier) error
	// SelectProductSupplier selects a product supplier by product and supplier id
	SelectProductSupplier(ctx context.Context, productId, supplierId string) (*domain.ProductSupplier, error)
	// UpdateProductSupplier updates a product supplier by product and supplier id
	UpdateProductSupplier(ctx context.Context, productId, supplierId string, p *domain.ProductSupplierPatch) error
	// SetPrimaryProductSupplier sets a product supplier as the primary supplier for a product
	SetPrimaryProductSupplier(ctx context.Context, productId, supplierId string) error

	CreateGroup(ctx context.Context, group *domain.ValidatedProductGroup) error
	ListGroup() ([]*domain.ProductGroup, error)
	FindGroupByID(ctx context.Context, id string) (*domain.ProductGroup, error)
	UpdateGroup(ctx context.Context, group *domain.ProductGroup) error
	DeleteGroup()
}
