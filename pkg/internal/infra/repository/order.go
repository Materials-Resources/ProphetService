package repository

import (
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
)

type orderRepository struct {
	db repository.Database
}

func (o orderRepository) Create() error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) Read(id string) (*domain.Order, error) {
	var oeHdr model.OeHdr

	err := o.db.GetDB().Model(&model.OeHdr{}).Preload("OeLineItems").Preload("OeLineItems.InvMast").First(
		&oeHdr,
		id,
	).Error
	if err != nil {
		return nil, gormToRepositoryError(err)
	}

	order := oeHdrToDomain(&oeHdr)

	return &order, nil
}

func (o orderRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) AddItem() error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db repository.Database) repository.OrderRepository {
	return &orderRepository{db: db}
}

func oeHdrToDomain(oeHdr *model.OeHdr) domain.Order {
	order := domain.Order{
		ID: oeHdr.OrderNo,
		ShippingAddress: domain.OrderShippingAddress{
			Name:                 oeHdr.Ship2Name.String,
			LineOne:              oeHdr.Ship2Add1.String,
			LineTwo:              oeHdr.Ship2Add2.String,
			City:                 oeHdr.Ship2City.String,
			State:                oeHdr.Ship2State.String,
			DeliveryInstructions: oeHdr.DeliveryInstructions.String,
		},
	}

	for _, oeLineItem := range oeHdr.OeLineItems {
		order.OrderItems = append(
			order.OrderItems,
			domain.OrderItem{
				ID:            strconv.Itoa(int(oeLineItem.InvMastUid)),
				SN:            oeLineItem.InvMast.ItemId,
				Name:          oeLineItem.InvMast.ItemDesc,
				CostPerUnit:   int64(oeLineItem.UnitPrice.Float64 * 100),
				UnitPurchased: int64(oeLineItem.QtyOrdered.Float64),
				UnitLabel:     oeLineItem.UnitOfMeasure.String,
			},
		)
	}

	return order
}
