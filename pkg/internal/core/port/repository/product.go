package repository

import (
	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
)

type ProductRepository interface {
	Create()
	Read(id string) (*domain.Product, error) // Returns product details given a identifier
	Update()
	Delete()
}
