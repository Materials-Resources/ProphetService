package catalog

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type Repository interface {
	CreateProduct()
	FindProductByID(id string) (*domain.ValidatedProduct, error)
	ReadProductByGroup(id string) ([]*domain.ValidatedProduct, error)
	UpdateProduct()
	DeleteProduct(ctx context.Context, id string) error

	CreateGroup()
	ListGroup() ([]*domain.ValidatedProductGroup, error)
	FindGroupByID(ctx context.Context, id string) (*domain.ValidatedProductGroup, error)
	UpdateGroup()
	DeleteGroup()
}
