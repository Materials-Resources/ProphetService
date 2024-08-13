package repository

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/uptrace/bun"
	"strconv"
)

type OrderRepository struct {
	db bun.IDB
}

func NewOrderRepository(db bun.IDB) *OrderRepository {
	return &OrderRepository{db: db}
}

type ListOrdersParams struct {
	CustomerId       string
	AdminId          string
	OnlyActiveOrders bool
	Page             int
	PageSize         int
}

func (r *OrderRepository) ListOrders(ctx context.Context, params ListOrdersParams) ([]*domain.Order, error) {
	var dbOrders []*oeHdr

	query := r.db.NewSelect().
		Model(&dbOrders).Relation("Customer").Relation("Contact")
	query = query.Limit(100).OrderExpr("oe_hdr.order_no")
	if params.OnlyActiveOrders {
		query = query.Where("oe_hdr.delete_flag = 'N'")
	}

	err := query.Scan(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}

	orders := make([]*domain.Order, len(dbOrders))
	for i, dbOrder := range dbOrders {
		orders[i] = mapDbOrderToDomainWithoutItems(dbOrder)
	}

	return orders, nil
}

func mapDbOrderToDomainWithoutItems(dbOrder *oeHdr) *domain.Order {

	order := &domain.Order{
		Id: dbOrder.OrderNo,
		Customer: domain.Customer{
			Id:   strconv.FormatFloat(dbOrder.Customer.CustomerId, 'f', 0, 64),
			Name: dbOrder.Customer.CustomerName,
		},
		PurchaseOrder:        dbOrder.PoNo,
		Taker:                dbOrder.Taker,
		DateCreated:          dbOrder.DateCreated,
		DeliveryInstructions: dbOrder.DeliveryInstructions,
	}

	if dbOrder.Contact != nil {
		order.Contact = domain.Contact{
			Id:        dbOrder.Contact.Id,
			FirstName: dbOrder.Contact.FirstName,
			LastName:  dbOrder.Contact.LastName,
		}
	}

	return order

}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	var dbOrder oeHdr

	err := r.db.NewSelect().Model(&dbOrder).Relation("OeLines", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Relation("InvMast")
	}).Relation("Customer").Relation("Contact").Where("order_no = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return mapDbOrderToDomainOrder(&dbOrder), nil
}

func mapDbOrderToDomainOrder(dbOrder *oeHdr) *domain.Order {
	order := &domain.Order{
		Id:                   dbOrder.OrderNo,
		PurchaseOrder:        dbOrder.PoNo,
		DeliveryInstructions: dbOrder.DeliveryInstructions,
		DateCreated:          dbOrder.DateCreated,
		DateRequested:        dbOrder.RequestedDate,
		Taker:                dbOrder.Taker,
		Customer: domain.Customer{
			Id:   strconv.FormatFloat(dbOrder.CustomerId, 'f', 0, 64),
			Name: dbOrder.Customer.CustomerName,
		},
	}

	if dbOrder.Contact != nil {
		order.Contact = domain.Contact{
			Id:        dbOrder.Contact.Id,
			FirstName: dbOrder.Contact.FirstName,
			LastName:  dbOrder.Contact.LastName,
			Title:     dbOrder.Contact.Title,
		}
	}

	if dbOrder.OeLines != nil {
		for _, line := range dbOrder.OeLines {
			order.Items = append(order.Items, mapDbLineToDomainOrderItem(line))
		}
	}

	return order
}

func mapDbLineToDomainOrderItem(dbLine *oeLine) *domain.OrderItem {
	item := &domain.OrderItem{
		Id:                fmt.Sprintf("%0.f", dbLine.LineNo),
		ProductUid:        strconv.Itoa(int(dbLine.InvMastUid)),
		ProductSn:         dbLine.InvMast.ItemId,
		ProductName:       dbLine.InvMast.ItemDesc,
		CustomerProductSn: dbLine.CustomerPartNumber,
		OrderQuantity:     *dbLine.QtyOrdered,
		OrderQuantityUnit: *dbLine.UnitOfMeasure,
		UnitPrice:         *dbLine.UnitPrice,
		TotalPrice:        *dbLine.ExtendedPrice,
		ShippedQuantity:   *dbLine.QtyInvoiced,
	}
	return item
}
