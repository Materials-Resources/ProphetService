package repositories

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
)

type OrderRepository interface {
	FindOrderById(id string) (*entities.ValidatedOrder, error)

	FindPickTicketByID(id string) (*entities.ValidatedPickTicket, error)
}
