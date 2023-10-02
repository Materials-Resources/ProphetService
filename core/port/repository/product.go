package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
)

type ProductRepository interface {
	Create(ctx context.Context)
	Read(ctx context.Context, id string) (*domain.Product, error) // Returns product details given a identifier
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
