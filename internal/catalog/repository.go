package catalog

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

// Repository is the interface that wraps the basic CRUD operations for the catalog
type Repository interface {
	CreateProduct()
	// SelectProduct returns a product by id
	SelectProduct(ctx context.Context, id int32) (*domain.Product, error)

	SelectProductPrice(ctx context.Context, uid []int32) ([]*domain.ProductPrice, error)

	// AddProductSupplier adds a product supplier to an existing product in the database
	AddProductSupplier(ctx context.Context, supplier *domain.ValidatedProductSupplier) error
	// SelectProductSupplier selects a product supplier by product and supplier id
	SelectProductSupplier(ctx context.Context, productId, supplierId string) (*domain.ProductSupplier, error)
	// UpdateProductSupplier updates a product supplier by product and supplier id
	UpdateProductSupplier(ctx context.Context, productId, supplierId string, p *domain.ProductSupplierPatch) error
	// SetPrimaryProductSupplier sets a product supplier as the primary supplier for a product
	SetPrimaryProductSupplier(ctx context.Context, productId, supplierId string) error
}
