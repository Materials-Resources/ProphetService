package repository

import (
	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
)

type OrderRepository interface {
	Create() error                         //Creates a new order
	Read(id string) (*domain.Order, error) //Reads an existing order
	Update() error                         //Updates an existing order
	Delete() error                         //Deletes an existing order
	AddItem() error                        //Adds an item to an existing order
}
