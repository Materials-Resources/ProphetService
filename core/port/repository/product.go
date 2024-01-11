package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context)
	//ReadProduct Returns product details given a identifier
	ReadProduct(ctx context.Context, id string) (*domain.Product, error)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	List(ctx context.Context, filter domain.ProductFilter) ([]*domain.Product, error)

	//ReadProductBySupplier Returns product serial number details given a supplier and its part number
	ReadProductBySupplier(ctx context.Context, supplierId string, supplierPartNo string) (*string, error)

	CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error
	ReadProductGroup(ctx context.Context, id string) (*domain.ProductGroup, error)
}
