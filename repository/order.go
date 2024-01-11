package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/core/port/repository"
	"github.com/materials-resources/s_prophet/model"
	orderpPb "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
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

func (o orderRepository) Read(ctx context.Context, id string) (*orderpPb.GetOrderResponse, error) {
	oeHdr := new(model.OeHdr)

	err := o.bun.NewSelect().Model(oeHdr).Relation("OeLineItems").Relation("OeLineItems.InvMast").Where(
		"order_no = ?",
		id,
	).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return toPbGetOrderResponse(oeHdr), nil
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
	//bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return &orderRepository{
		db:  db,
		bun: bundb,
	}
}

func toPbGetOrderResponse(mod *model.OeHdr) (pb *orderpPb.GetOrderResponse) {
	fmt.Println(mod.QuoteType)
	pb = &orderpPb.GetOrderResponse{
		Id:            mod.OrderNo,
		PurchaseOrder: mod.PoNo.String,
		Status:        mod.ValidationStatus.String,
	}

	for _, item := range mod.OeLineItems {
		pb.OrderItems = append(
			pb.OrderItems,
			&orderpPb.OrderItem{
				Id:            strconv.Itoa(int(item.InvMastUid)),
				Sn:            item.InvMast.ItemId,
				Name:          item.InvMast.ItemDesc,
				UnitPurchased: item.QtyOrdered.Float64,
				UnitLabel:     item.UnitOfMeasure.String,
				CostPerUnit:   item.UnitPrice.Float64,
				TotalPrice:    item.ExtendedPrice.Float64,
			},
		)
	}
	return
}
