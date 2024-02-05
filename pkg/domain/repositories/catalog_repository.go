package repositories

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
)

type CatalogRepository interface {
	CreateProduct()
	FindProductByID(id string) (*entities.ValidatedProduct, error)
	ReadProductByGroup(id string) ([]*entities.ValidatedProduct, error)
	UpdateProduct()
	DeleteProduct(ctx context.Context, id string) error

	CreateGroup()
	ListGroup() ([]*entities.ValidatedProductGroup, error)
	FindGroupByID(ctx context.Context, id string) (*entities.ValidatedProductGroup, error)
	UpdateGroup()
	DeleteGroup()
}
