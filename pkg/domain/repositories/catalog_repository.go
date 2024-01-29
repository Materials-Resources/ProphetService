package repositories

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
)

type CatalogRepository interface {
	CreateProduct()
	FindProductByID(id string) (*entities.ValidatedProduct, error)
	ReadProductByGroup(id string) ([]*entities.ValidatedProduct, error)
	UpdateProduct()
	DeleteProduct()

	CreateGroup()
	ListGroup() ([]*entities.ValidatedProductGroup, error)
	FindGroupByID(id string) (*entities.ValidatedProductGroup, error)
	UpdateGroup()
	DeleteGroup()
}
