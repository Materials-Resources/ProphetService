package prophet_19_1_3668

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
	"github.com/uptrace/bun"
)

func NewBunOrderRepository(db *bun.DB) repositories.OrderRepository {
	return &BunOrderRepository{db: db}
}

type BunOrderRepository struct {
	db *bun.DB
}

func (b BunOrderRepository) FindOrderById(id string) (*entities.ValidatedOrder, error) {
	ctx := context.Background()
	dbOrder := new(model.OeHdr)

	err := b.db.NewSelect().Model(dbOrder).Relation("OeLineItems").Relation("Contact").Relation("OeLineItems.InvMast").Where("oe_hdr.order_no = ?",
		id,
	).Scan(ctx)
	fmt.Println(dbOrder.Contact)
	if err != nil {
		return nil, errors.New("could not find requested order")
	}
	return FromDBOrder(dbOrder)
}

func (b BunOrderRepository) FindPickTicketByID(id string) (*entities.ValidatedPickTicket, error) {
	ctx := context.Background()

	dbID, err := strconv.ParseFloat(id, 64)
	if err != nil {
		return nil, errors.New("invalid id provided")
	}

	dbPickTicket := new(model.OePickTicket)

	err = b.db.NewSelect().Model(dbPickTicket).Relation("OeHdr").Where("oe_pick_ticket.pick_ticket_no = ?", dbID).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not find requested pick ticket")
	}
	return FromDBPickTicket(dbPickTicket)
}
