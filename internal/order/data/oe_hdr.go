package data

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/uptrace/bun"
	"strconv"
	"time"
)

type oeHdr struct {
	data.OeHdr `bun:",extend"`
	Customer   *customer `bun:"rel:has-one,join:customer_id=customer_id,join:company_id=company_id"`
	Contact    *contacts `bun:"rel:has-one,join:contact_id=id"`
	Lines      []*oeLine `bun:"rel:has-many,join:order_no=order_no"`
}

type OeHdrModel struct {
	bun bun.IDB
}

func NewOeHdrModel(bun bun.IDB) OeHdrModel {
	return OeHdrModel{bun: bun}
}

var _ bun.BeforeAppendModelHook = (*oeHdr)(nil)

func (m *oeHdr) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

func (m *oeHdr) SelectDefaultColumns() []string {
	return []string{
		"order_no",
		"address_id",
		"customer_id",
		"contact_id",
		"ship2_name",
		"ship2_add1",
		"ship2_add2",
		"ship2_city",
		"ship2_state",
		"ship2_zip",
		"po_no",
		"delivery_instructions",
		"taker",
		"order_date",
	}
}

func oeHdrToOrder(rec *oeHdr) *domain.Order {
	order := &domain.Order{
		Id:                        rec.OrderNo,
		CustomerBranchId:          fmt.Sprintf("%.0f", rec.AddressId.Float64),
		CustomerId:                fmt.Sprintf("%.0f", rec.CustomerId),
		ContactId:                 rec.ContactId.String,
		PurchaseOrder:             rec.PoNo.String,
		DeliveryInstructions:      rec.DeliveryInstructions.String,
		Taker:                     rec.Taker.String,
		OrderDate:                 rec.OrderDate.Time,
		ShippingAddressId:         fmt.Sprintf("%.0f", rec.AddressId.Float64),
		ShippingAddressName:       rec.Ship2Name.String,
		ShippingAddressLineOne:    rec.Ship2Add1.String,
		ShippingAddressLineTwo:    rec.Ship2Add2.String,
		ShippingAddressCity:       rec.Ship2City.String,
		ShippingAddressState:      rec.Ship2State.String,
		ShippingAddressPostalCode: rec.Ship2Zip.String,
		ShippingAddressCountry:    rec.Ship2Country.String,
	}

	if rec.Lines != nil {
		order.Items = make([]*domain.OrderItem, len(rec.Lines))
		for i, line := range rec.Lines {
			order.Items[i] = &domain.OrderItem{
				Id:                  fmt.Sprintf("%0.f", line.LineNo),
				ProductUid:          strconv.Itoa(int(line.InvMastUid)),
				ProductSn:           line.InvMast.ItemId,
				ProductName:         line.InvMast.ItemDesc,
				CustomerProductSn:   line.CustomerPartNumber,
				OrderQuantity:       line.QtyOrdered.Float64,
				OrderQuantityUnit:   line.UnitOfMeasure.String,
				PriceUnit:           "",
				Price:               line.UnitPrice.Float64,
				TotalPrice:          line.ExtendedPrice.Float64,
				ShippedQuantity:     line.QtyInvoiced.Float64,
				BackOrderedQuantity: line.QtyOrdered.Float64 - line.QtyInvoiced.Float64,
			}
		}

	}

	if rec.Customer != nil {
		order.CustomerId = strconv.Itoa(int(rec.Customer.CustomerId))
		order.CustomerName = rec.Customer.CustomerName.String

	}

	if rec.Contact != nil {
		order.ContactName = fmt.Sprintf("%s %s", rec.Contact.FirstName, rec.Contact.LastName)
	}
	return order
}

func (m *OeHdrModel) Get(ctx context.Context, id string) (*domain.Order, error) {
	var oeHdrRec oeHdr

	err := m.bun.NewSelect().Model(&oeHdrRec).Relation("Lines", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Relation("InvMast")
	}).Relation("Customer").Relation("Contact").Where("order_no = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return oeHdrToOrder(&oeHdrRec), nil

}

type OeHdrListOptions struct {
	CustomerId       *[]string
	CustomerBranchId *[]string
}

func (m *OeHdrModel) List(ctx context.Context, opt *OeHdrListOptions) ([]*domain.Order, error) {
	var oeHdrRecs []*oeHdr

	q := m.bun.NewSelect().Model(&oeHdrRecs)
	if opt.CustomerId != nil {
		customerIds, err := convertSliceStringToFloat64(*opt.CustomerId)
		if err != nil {
			// TODO return invalid argument error
			return nil, err
		}
		q.Where("customer_id IN (?)", bun.In(customerIds))
	}
	if opt.CustomerBranchId != nil {
		customerBranchIds, err := convertSliceStringToFloat64(*opt.CustomerBranchId)
		if err != nil {
			// TODO return invalid argument error
			return nil, err
		}
		q.Where("address_id IN (?)", bun.In(customerBranchIds))

	}

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]*domain.Order, len(oeHdrRecs))
	for i, rec := range oeHdrRecs {
		orders[i] = oeHdrToOrder(rec)
	}

	return orders, nil
}

func convertSliceStringToFloat64(slice []string) ([]float64, error) {
	var floatSlice []float64
	for _, str := range slice {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floatSlice = append(floatSlice, val)
	}
	return floatSlice, nil
}
