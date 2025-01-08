package repository

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
	"strconv"
	"time"
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

type CreateQuoteItemParams struct {
	ProductId string
	Quantity  int
}
type CreateQuoteParams struct {
	PurchaseOrder        string
	CustomerId           string
	ContactId            string
	RequestDate          time.Time
	DeliveryInstructions string
	Items                []CreateQuoteItemParams
}

func (r *OrderRepository) CreateQuote(ctx context.Context, quote *CreateQuoteParams) error {
	customerId, err := strconv.ParseFloat(quote.CustomerId, 64)

	// TODO validate contact belongs to customer
	customerRec, err := r.getCustomer(ctx, customerId)
	if err != nil {
		return err
	}

	addressRec, err := r.getAddress(ctx, customerId)
	if err != nil {
		return err
	}
	contactRec, err := r.getContact(ctx, quote.ContactId)
	if err != nil {
		return err
	}
	oeHdrUid, err := r.getNextCounter(ctx, "oe_hdr")
	if err != nil {
		return fmt.Errorf("failed to get oe_hdr_uid: %w", err)
	}
	orderNo, err := r.getNextCounter(ctx, "WO")
	if err != nil {
		return fmt.Errorf("failed to get order_no: %w", err)
	}

	// Create new oeHdr record with eSTORE taker

	carrierId, err := strconv.ParseFloat(*addressRec.CarrierId, 64)
	oeHdrNew := prophet.OeHdr{
		OeHdrUid:             oeHdrUid,
		OrderNo:              strconv.Itoa(int(orderNo)),
		CustomerId:           customerId,
		ContactId:            &quote.ContactId,
		ProjectedOrder:       ptrTo("Y"),
		PoNo:                 &quote.PurchaseOrder,
		DeliveryInstructions: &quote.DeliveryInstructions,
		PickTicketType:       customerRec.PickTicketType,
		Terms:                &customerRec.TermsId,
		AddressId:            &addressRec.Id,
		Ship2Name:            &addressRec.Name,
		Ship2Add1:            addressRec.MailAddress1,
		Ship2Add2:            addressRec.MailAddress2,
		Ship2Add3:            addressRec.MailAddress3,
		Ship2City:            addressRec.MailCity,
		Ship2State:           addressRec.MailState,
		Ship2Country:         addressRec.MailCountry,
		ShipToPhone:          contactRec.DirectPhone,
		CarrierId:            &carrierId,
	}

	if err := r.createOeHdr(ctx, &oeHdrNew); err != nil {
		return err
	}

	// create new oeHdrSalesRep record

	if err := r.createOeHdrSalesRep(ctx, oeHdrNew.OrderNo, *customerRec.SalesrepId); err != nil {
		return err
	}

	// create new QuoteHdr record

	quoteHdrUid, err := r.getNextCounter(ctx, "quote_hdr")
	if err != nil {
		return fmt.Errorf("failed to get quote_hdr_uid: %w", err)
	}

	quoteHdrNew := prophet.QuoteHdr{
		LastMaintainedBy: "admin",
		QuoteHdrUid:      quoteHdrUid,
		OeHdrUid:         oeHdrUid,
	}

	_, err = r.db.NewInsert().Model(&quoteHdrNew).Exec(ctx)

	// create new oeLine record for each item

	return nil
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

func (r *OrderRepository) getNextCounter(ctx context.Context, counterName string) (int32, error) {
	var uid int32
	query := fmt.Sprintf(`DECLARE @uid int
		EXEC @uid = p21_get_counter '%s', 1
		SELECT @uid`, counterName)
	err := r.db.QueryRowContext(ctx, query).Scan(&uid)
	return uid, err
}

func (r *OrderRepository) getCustomer(ctx context.Context, customerId float64) (*customer, error) {
	customerRec := customer{}
	err := r.db.NewSelect().Model(&customerRec).Where("customer_id = ? AND company_id = ?", customerId, "MRS").Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &customerRec, nil
}

func (r *OrderRepository) getAddress(ctx context.Context, customerId float64) (*address, error) {
	addressRec := address{}
	err := r.db.NewSelect().Model(&addressRec).Where("id = ?", customerId).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &addressRec, nil
}

func (r *OrderRepository) getContact(ctx context.Context, id string) (*contacts, error) {
	contactRec := contacts{}
	err := r.db.NewSelect().Model(&contactRec).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &contactRec, nil
}

func (r *OrderRepository) createOeHdr(ctx context.Context, oeHdr *prophet.OeHdr) error {
	_, err := r.db.NewInsert().Model(oeHdr).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create oeHdr: %w", err)
	}
	return nil
}
func (r *OrderRepository) createOeHdrSalesRep(ctx context.Context, orderNo string, salesRepId string) error {
	oeHdrSalesRep := prophet.OeHdrSalesrep{
		OrderNumber:     orderNo,
		SalesrepId:      salesRepId,
		CommissionSplit: 100,
		PrimarySalesrep: ptrTo("Y"),
	}

	_, err := r.db.NewInsert().Model(&oeHdrSalesRep).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create oeHdrSalesRep: %w", err)
	}
	return nil
}

func ptrTo[T any](v T) *T {
	return &v
}
