package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

type shippingRepository struct {
	db  repository.Database
	bun *bun.DB
}

func (r shippingRepository) GetPickTicket(ctx context.Context, id string) (*domain.PickTicket, error) {
	oePickTicket := new(model.OePickTicket)
	pickTicketId, err := strconv.ParseFloat(
		id,
		32,
	)
	if err != nil {
		return nil, err
	}
	err = r.bun.NewSelect().Model(oePickTicket).Relation("OeHdr").Relation("OeHdr.Contact").Where(
		"pick_ticket_no = ?",
		pickTicketId,
	).Scan(ctx)
	fmt.Println(oePickTicket.OeHdr)
	if err != nil {
		return nil, err
	}

	pickTicket := oePickTicketToDomain(oePickTicket)
	return &pickTicket, nil

}

func NewShippingRepository(db repository.Database) repository.ShippingRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	return &shippingRepository{
		db:  db,
		bun: bundb,
	}
}

func oePickTicketToDomain(oePickTicket *model.OePickTicket) domain.PickTicket {
	pickTicket := domain.PickTicket{
		ID: strconv.FormatFloat(
			oePickTicket.PickTicketNo,
			'f',
			-1,
			64,
		),
		ShippingAddress: domain.PickTicketShippingAddress{
			Name:                 oePickTicket.OeHdr.Ship2Name.String,
			LineOne:              oePickTicket.OeHdr.Ship2Add1.String,
			LineTwo:              oePickTicket.OeHdr.Ship2Add2.String,
			City:                 oePickTicket.OeHdr.Ship2City.String,
			State:                oePickTicket.OeHdr.Ship2State.String,
			DeliveryInstructions: oePickTicket.Instructions.String,
		},
		OrderId:            oePickTicket.OrderNo,
		OrderPurchaseOrder: oePickTicket.OeHdr.PoNo.String,
		OrderContactName: fmt.Sprintf(
			"%s %s",
			oePickTicket.OeHdr.Contact.FirstName,
			oePickTicket.OeHdr.Contact.LastName,
		),
	}

	return pickTicket
}
