package repository

import "github.com/uptrace/bun"

type OrderRepository struct {
	db bun.IDB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}
