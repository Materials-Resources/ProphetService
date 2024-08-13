package repository

import "github.com/uptrace/bun"

type Repository struct {
	Order    *OrderRepository
	Shipment *ShipmentRepository
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{Order: NewOrderRepository(db), Shipment: NewShipmentRepository(db)}
}
