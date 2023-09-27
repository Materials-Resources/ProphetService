package repository

import (
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
)

type shippingRepository struct {
	db repository.Database
}

func (s shippingRepository) GetPickTicket(id string) (*domain.PickTicket, error) {
	var oePickTicket model.OePickTicket
	err := s.db.GetDB().Model(&model.OePickTicket{}).Preload("OeHdr").Preload("OeHdr.Contact").First(
		&oePickTicket,
		id,
	).Error
	if err != nil {
		return nil, err
	}

	pickTicket := oePickTicketToDomain(&oePickTicket)
	return &pickTicket, nil

}

func NewShippingRepository(db repository.Database) repository.ShippingRepository {
	return &shippingRepository{db: db}
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
