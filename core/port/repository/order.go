package repository

import (
	"context"

	orderpPb "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
)

type OrderRepository interface {
	Create(ctx context.Context) error                                        //Creates a new order
	Read(ctx context.Context, id string) (*orderpPb.GetOrderResponse, error) //Reads an existing order
	Update(ctx context.Context) error                                        //Updates an existing order
	Delete(ctx context.Context) error                                        //Deletes an existing order
	AddItem(ctx context.Context) error                                       //Adds an item to an existing order
}
