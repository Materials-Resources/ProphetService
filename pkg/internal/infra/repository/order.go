package repository

import (
	"context"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

type orderRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (o orderRepository) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) Read(ctx context.Context, id string) (*domain.Order, error) {
	oeHdr := new(model.OeHdr)

	err := o.bun.NewSelect().Model(oeHdr).Relation("OeLineItems").Relation("OeLineItems.InvMast").Where(
		"order_no = ?",
		id,
	).Scan(ctx)
	if err != nil {
		return nil, gormToRepositoryError(err)
	}

	order := oeHdrToDomain(oeHdr)

	return &order, nil
}

func (o orderRepository) Update(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) AddItem(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db repository.Database) repository.OrderRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	return &orderRepository{
		db:  db,
		bun: bundb,
	}
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
