package order

import (
	entities "github.com/materials-resources/s_prophet/internal/order/domain"
)

type Repository interface {
	CreateOrder(order entities.ValidatedOrder) error
	AddItemToOrder(orderId string, items []entities.ValidatedOrderItem) error
	FindOrderById(id string) (*entities.ValidatedOrder, error)

	FindPickTicketByID(id string) (*entities.ValidatedPickTicket, error)
}
