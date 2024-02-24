package catalog

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type Repository interface {
	CreateProduct()
	FindProductByID(id string) (*domain.ValidatedProduct, error)
	ListProduct(ctx context.Context, cursor int32, count int) (res []*domain.Product, nextCursor int, err error)
	ReadProductByGroup(id string) ([]*domain.Product, error)
	UpdateProduct()
	DeleteProduct(ctx context.Context, id string) error

	GetProductSupplier(ctx context.Context, productId, supplierId string) (*domain.ProductSupplier, error)
	UpdateProductSupplier(ctx context.Context, p *domain.ValidatedProductSupplier) error
	SetPrimaryProductSupplier(ctx context.Context, productId, supplierId string) error

	CreateGroup(ctx context.Context, group *domain.ValidatedProductGroup) error
	ListGroup() ([]*domain.ValidatedProductGroup, error)
	FindGroupByID(ctx context.Context, id string) (*domain.ValidatedProductGroup, error)
	UpdateGroup(ctx context.Context, group *domain.ProductGroup) error
	DeleteGroup()
}
