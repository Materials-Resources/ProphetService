package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
)

type OrderRepository interface {
	Create(ctx context.Context) error                           //Creates a new order
	Read(ctx context.Context, id string) (*domain.Order, error) //Reads an existing order
	Update(ctx context.Context) error                           //Updates an existing order
	Delete(ctx context.Context) error                           //Deletes an existing order
	AddItem(ctx context.Context) error                          //Adds an item to an existing order
}
