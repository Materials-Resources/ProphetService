package repositories

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order entities.ValidatedOrder) error
	AddItemToOrder(orderId string, items []entities.ValidatedOrderItem) error
	FindOrderById(id string) (*entities.ValidatedOrder, error)

	FindPickTicketByID(id string) (*entities.ValidatedPickTicket, error)
}
