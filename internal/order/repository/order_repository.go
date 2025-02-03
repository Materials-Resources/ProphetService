package repository

import (
	"context"
	"database/sql"
	"errors"
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
	BranchId         string
	CustomerId       string
	Page             int
	PageSize         int
	AdminId          string
	OnlyActiveOrders bool
}

type CreateQuoteItemParams struct {
	ProductId string
	Quantity  float64
}
type CreateQuoteParams struct {
	Notes       string
	ContactId   string
	BranchId    string
	RequestDate time.Time
	Items       []CreateQuoteItemParams
}

type ListQuotesParams struct {
	BranchId         string
	CustomerId       string
	Page             int
	PageSize         int
	AdminId          string
	OnlyActiveOrders bool
}

func (r *OrderRepository) ListQuotes(ctx context.Context, params ListQuotesParams) ([]*domain.QuoteSummary,
	int32, error) {
	var oeHdrRecs []*OeHdr
	branchId, err := strconv.ParseFloat(params.BranchId, 64)
	if err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.PageSize
	err = r.db.NewSelect().Model(&oeHdrRecs).Relation("QuoteHdr").Where(
		"address_id = ? and projected_order = ? and delete_flag = ? and rma_flag <> ?", branchId, "Y", "N", "Y").Limit(params.PageSize).Offset(offset).Order("date_created DESC").Scan(ctx)

	if err != nil {
		return nil, 0, err
	}
	var quotes []*domain.QuoteSummary

	for _, oeHdrRec := range oeHdrRecs {
		quotes = append(quotes, &domain.QuoteSummary{
			Id:        oeHdrRec.OrderNo,
			ContactId: getOptionalValue(oeHdrRec.ContactId, ""),
			BranchId: convertOptionalValue(oeHdrRec.AddressId, func(v float64) string {
				return strconv.FormatFloat(v, 'f', 0, 64)
			}, ""),
			PurchaseOrder: getOptionalValue(oeHdrRec.PoNo, ""),
			DateCreated:   oeHdrRec.DateCreated,
			DateRequested: offSetToUtc(oeHdrRec.RequestedDate),
			Status:        determineQuoteStatus(oeHdrRec),
		})
	}

	// Fetch the total count of records disregarding pagination
	totalRecordCount, err := r.db.NewSelect().
		Model(&oeHdrRecs).
		Where("address_id = ? and projected_order = ? and delete_flag = ? and rma_flag <> ?", branchId, "Y", "N", "Y").
		Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return quotes, int32(totalRecordCount), err
}

func (r *OrderRepository) GetQuote(ctx context.Context, id string) (*domain.Quote, error) {
	var oeHdrRec OeHdr
	err := r.db.NewSelect().Model(&oeHdrRec).Relation("QuoteHdr").Relation("OeLines", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Relation("InvMast")
	}).Where(
		"order_no = ?",
		id).Where("projected_order = ? and delete_flag = ? and rma_flag <> ?", "Y", "N", "Y").Scan(ctx)
	if err != nil {
		return nil, err
	}

	quote := domain.Quote{
		Id: oeHdrRec.OrderNo,
		Branch: &domain.Branch{
			Id: convertOptionalValue(oeHdrRec.AddressId, func(v float64) string {
				return strconv.FormatFloat(v, 'f', 0, 64)
			}, ""),
			Name: getOptionalValue(oeHdrRec.Ship2Name, ""),
		},
		Customer: domain.Customer{},
		Contact: domain.Contact{
			Id: getOptionalValue(oeHdrRec.ContactId, ""),
		},
		PurchaseOrder: getOptionalValue(oeHdrRec.PoNo, ""),
		DateCreated:   oeHdrRec.DateCreated,
		DateRequested: offSetToUtc(oeHdrRec.RequestedDate),
		DateExpires:   offSetToUtc(oeHdrRec.QuoteHdr.ExpirationDate),
		Status:        determineQuoteStatus(&oeHdrRec),
	}

	for _, oeLineRec := range oeHdrRec.OeLines {
		quote.Items = append(quote.Items, &domain.QuoteItem{
			ProductId:         strconv.Itoa(int(oeLineRec.InvMastUid)),
			ProductSn:         oeLineRec.InvMast.ItemId,
			ProductName:       oeLineRec.InvMast.ItemDesc,
			OrderedQuantity:   getOptionalValue(oeLineRec.QtyOrdered, 0),
			CustomerProductSn: oeLineRec.CustomerPartNumber,
			UnitType:          getOptionalValue(oeLineRec.UnitOfMeasure, ""),
			UnitPrice:         getOptionalValue(oeLineRec.UnitPrice, 0),
			TotalPrice:        getOptionalValue(oeLineRec.ExtendedPrice, 0),
		})
	}

	return &quote, nil

}
func (r *OrderRepository) CreateQuote(ctx context.Context, quote *CreateQuoteParams) (*string, error) {
	branchId, err := strconv.ParseFloat(quote.BranchId, 64)
	if err != nil {
		return nil, err
	}

	shipToRec, err := r.getShipTo(ctx, branchId)
	if err != nil {
		return nil, err
	}

	// TODO validate contact belongs to Customer
	addressRec, err := r.getAddress(ctx, shipToRec.ShipToId)
	if err != nil {
		return nil, err
	}

	customerRec, err := r.getCustomer(ctx, shipToRec.CustomerId)
	if err != nil {
		return nil, err
	}

	contactRec, err := r.getContact(ctx, quote.ContactId)
	if err != nil {
		return nil, err
	}
	oeHdrUid, err := r.getNextCounter(ctx, "oe_hdr")
	if err != nil {
		return nil, fmt.Errorf("failed to get oe_hdr_uid: %w", err)
	}
	orderNo, err := r.getNextCounter(ctx, "WO")
	if err != nil {
		return nil, fmt.Errorf("failed to get order_no: %w", err)
	}

	// Create new OeHdr record with eSTORE taker

	oeHdrNew := OeHdr{
		OeHdr: prophet.OeHdr{
			OeHdrUid:       oeHdrUid,
			OrderNo:        strconv.Itoa(int(orderNo)),
			CustomerId:     customerRec.CustomerId,
			ContactId:      &quote.ContactId,
			PickTicketType: customerRec.PickTicketType,
			PackingBasis:   addressRec.ShipToPackingBasis,
			Terms:          &customerRec.TermsId,
			Taker:          ptrTo("eSTORE"),
			AddressId:      &branchId,
			Ship2Name:      &addressRec.Name,
			Ship2Add1:      addressRec.MailAddress1,
			Ship2Add2:      addressRec.MailAddress2,
			Ship2Add3:      addressRec.MailAddress3,
			Ship2City:      addressRec.MailCity,
			Ship2State:     addressRec.MailState,
			Ship2Country:   addressRec.MailCountry,
			ShipToPhone:    contactRec.DirectPhone,
			CarrierId: convertOptionalValue(addressRec.CarrierId, func(v string) *float64 {
				carrierId, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil
				}
				return &carrierId
			}, nil),
			PromiseDate:      ptrTo(time.Now().AddDate(0, 0, 1)),
			OrderDate:        ptrTo(time.Now()),
			RequestedDate:    &quote.RequestDate,
			LastMaintainedBy: "admin",
			CompanyId:        ptrTo("MRS"),
			LocationId:       ptrTo(float64(1001)),
			SourceLocationId: ptrTo(float64(1001)),

			OrderType:    ptrTo(int32(706)),
			SourceCodeNo: 920,
			CodFlag:      ptrTo("N"),

			ProjectedOrder:        ptrTo("Y"),
			Completed:             ptrTo("N"),
			FobFlag:               ptrTo("N"),
			RmaFlag:               ptrTo("N"),
			ThirdPartyBillingFlag: ptrTo("S"),
			Approved:              ptrTo("N"),
			CancelFlag:            ptrTo("N"),
			FrontCounter:          ptrTo("N"),
			HandlingChargeReqFlag: ptrTo("N"),
			ValidationStatus:      ptrTo("OK"),
		},
	}

	if err := r.createOeHdr(ctx, &oeHdrNew); err != nil {
		return nil, err
	}

	if quote.Notes != "" {
		//TODO: Create a new OeHdrNotepad record using the notes
		noteId, err := r.getNextCounter(ctx, "")
		if err != nil {
			return nil, fmt.Errorf("failed to get oeHdrNotepad_uid: %w", err)
		}

		noteArea := &NoteArea{
			NoteArea: prophet.NoteArea{
				NoteId: float64(noteId),
			},
		}

		_, err = r.db.NewInsert().Model(&noteArea).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create oeHdrNotepad record: %w", err)
		}

	}

	if err := r.createOeHdrSalesRep(ctx, oeHdrNew.OrderNo, *customerRec.SalesrepId); err != nil {
		return nil, err
	}

	// create new QuoteHdr record

	quoteHdrUid, err := r.getNextCounter(ctx, "quote_hdr")
	if err != nil {
		return nil, fmt.Errorf("failed to get quote_hdr_uid: %w", err)
	}

	quoteHdrNew := prophet.QuoteHdr{
		LastMaintainedBy: "admin",
		QuoteHdrUid:      quoteHdrUid,
		OeHdrUid:         oeHdrUid,
		ExpirationDate:   ptrTo(time.Now().AddDate(1, 0, 0)),
	}

	_, err = r.db.NewInsert().Model(&quoteHdrNew).Exec(ctx)

	// create new OeLine record for each item
	for _, item := range quote.Items {

		productId, err := strconv.ParseFloat(item.ProductId, 64)
		if err != nil {
			return nil, err
		}
		invMastRec, err := r.getProductDetails(ctx, int32(productId))
		if err != nil {
			return nil, err
		}

		lineNo, err := r.getNextOeLineLineNo(ctx, oeHdrUid)
		if err != nil {
			return nil, err
		}

		locRec := invMastRec.InvLocs[0]

		oeLineUid, err := r.getNextCounter(ctx, "oe_line")

		oeLineNew := OeLine{
			OeLine: prophet.OeLine{
				OeLineUid:            oeLineUid,
				InvMastUid:           invMastRec.InvMastUid,
				LineNo:               float64(lineNo),
				LineSeqNo:            &lineNo,
				OrderNo:              oeHdrNew.OrderNo,
				OeHdrUid:             oeHdrUid,
				CustomerPartNumber:   invMastRec.ItemId,
				ExtendedDesc:         invMastRec.ExtendedDesc,
				SupplierId:           locRec.PrimarySupplierId,
				ProductGroupId:       locRec.ProductGroupId,
				SalesDiscountGroupId: locRec.SalesDiscountGroup,
				UnitOfMeasure:        ptrTo("EA"),   //TODO use query to get default uom
				UnitSize:             1,             //TODO use query to get uit size
				UnitQuantity:         item.Quantity, // TODO use unit size multipled by order quantity
				Complete:             ptrTo("N"),
				CancelFlag:           ptrTo("N"),
				CaptureUsage:         "Y",
				ShipLocId:            ptrTo(float64(1001)),
				Assembly:             ptrTo("N"), // TODO use query to determine if assembly item
				Scheduled:            ptrTo("N"),
				QtyOrdered:           &item.Quantity,
				CompanyNo:            "MRS",
				SourceLocId:          ptrTo(float64(1001)),
			},
		}
		_, err = r.db.NewInsert().Model(&oeLineNew).Exec(ctx)
		if err != nil {
			return nil, err
		}

	}

	return &oeHdrNew.OrderNo, nil
}

func (r *OrderRepository) ListOrders(ctx context.Context, params ListOrdersParams) ([]*domain.OrderSummary,
	int32, error) {
	var oeHdrRecs []*OeHdr

	branchId, err := strconv.ParseFloat(params.BranchId, 64)
	if err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.PageSize

	err = r.db.NewSelect().Model(&oeHdrRecs).Where(
		"address_id = ? and projected_order = ? and delete_flag = ? and rma_flag <> ?", branchId, "N",
		"N", "Y").Limit(params.PageSize).Offset(offset).Order("date_created DESC").Scan(ctx)

	if err != nil {
		return nil, 0, err
	}

	var orders []*domain.OrderSummary

	for _, oeHdrRec := range oeHdrRecs {

		orderSummary := &domain.OrderSummary{
			Id:            oeHdrRec.OrderNo,
			PurchaseOrder: getOptionalValue(oeHdrRec.PoNo, ""),
			Status:        determineOrderStatus(oeHdrRec),
			DateCreated:   oeHdrRec.DateCreated,
			DateRequested: getOptionalValue(oeHdrRec.RequestedDate, time.Now()),
			ContactId:     getOptionalValue(oeHdrRec.ContactId, ""),
			BranchId: convertOptionalValue(oeHdrRec.AddressId, func(v float64) string {
				return strconv.FormatFloat(v, 'f', 0, 64)
			}, ""),
		}

		if oeHdrRec.AddressId != nil {
			orderSummary.BranchId = strconv.FormatFloat(*oeHdrRec.AddressId, 'f', 0, 64)
		}

		orders = append(orders, orderSummary)

	}

	// Fetch the total count of records disregarding pagination
	totalRecordCount, err := r.db.NewSelect().
		Model(&oeHdrRecs).Where(
		"address_id = ? and projected_order = ? and delete_flag = ? and rma_flag <> ?", branchId, "N",
		"N", "Y").Count(ctx)

	return orders, int32(totalRecordCount), err
}

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	var oeHdrRec OeHdr

	err := r.db.NewSelect().Model(&oeHdrRec).Relation("OeLines", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Relation("InvMast")
	}).Where("order_no = ?", id).Where("projected_order = ? and delete_flag = ? and rma_flag <> ?", "N", "N", "Y").Scan(ctx)
	if err != nil {
		return nil, err
	}

	order := domain.Order{
		Id:                   oeHdrRec.OrderNo,
		PurchaseOrder:        getOptionalValue(oeHdrRec.PoNo, ""),
		Status:               determineOrderStatus(&oeHdrRec),
		Taker:                getOptionalValue(oeHdrRec.Taker, ""),
		DateCreated:          offSetToUtc(&oeHdrRec.DateCreated),
		DateRequested:        offSetToUtc(oeHdrRec.RequestedDate),
		DeliveryInstructions: getOptionalValue(oeHdrRec.DeliveryInstructions, ""),
		BranchId: convertOptionalValue(oeHdrRec.AddressId, func(v float64) string {
			return strconv.FormatFloat(v, 'f', 0, 64)
		}, ""),
	}

	if oeHdrRec.AddressId != nil {
		order.BranchId = strconv.FormatFloat(*oeHdrRec.AddressId, 'f', 0, 64)
	}

	order.ShippingAddress = domain.Address{
		Id: convertOptionalValue(oeHdrRec.AddressId, func(v float64) string {
			return strconv.FormatFloat(v, 'f', 0, 64)
		}, ""),
		Name:       getOptionalValue(oeHdrRec.Ship2Name, ""),
		LineOne:    getOptionalValue(oeHdrRec.Ship2Add1, ""),
		LineTwo:    getOptionalValue(oeHdrRec.Ship2Add2, ""),
		City:       getOptionalValue(oeHdrRec.Ship2City, ""),
		State:      getOptionalValue(oeHdrRec.Ship2State, ""),
		PostalCode: getOptionalValue(oeHdrRec.Ship2Zip, ""),
		Country:    getOptionalValue(oeHdrRec.Ship2Country, ""),
	}

	for _, oeLineRec := range oeHdrRec.OeLines {
		order.Items = append(order.Items, &domain.OrderItem{
			ProductId:         strconv.Itoa(int(oeLineRec.InvMastUid)),
			ProductSn:         oeLineRec.InvMast.ItemId,
			ProductName:       oeLineRec.InvMast.ItemDesc,
			CustomerProductSn: oeLineRec.CustomerPartNumber,
			OrderedQuantity:   getOptionalValue(oeLineRec.QtyOrdered, 0),
			UnitType:          getOptionalValue(oeLineRec.UnitOfMeasure, ""),
			UnitPrice:         getOptionalValue(oeLineRec.UnitPrice, 0),
			TotalPrice:        getOptionalValue(oeLineRec.ExtendedPrice, 0),
			ShippedQuantity:   getOptionalValue(oeLineRec.QtyInvoiced, 0),
		})
	}

	return &order, nil
}

func (r *OrderRepository) getNextCounter(ctx context.Context, counterName string) (int32, error) {
	var uid int32
	query := fmt.Sprintf(`DECLARE @uid int
		EXEC @uid = p21_get_counter '%s', 1
		SELECT @uid`, counterName)
	err := r.db.QueryRowContext(ctx, query).Scan(&uid)
	return uid, err
}

func (r *OrderRepository) getCustomer(ctx context.Context, customerId float64) (*Customer, error) {
	customerRec := Customer{}
	err := r.db.NewSelect().Model(&customerRec).Where("customer_id = ? AND company_id = ?", customerId, "MRS").Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &customerRec, nil
}

func (r *OrderRepository) getAddress(ctx context.Context, branchId float64) (*Address, error) {
	addressRec := Address{}
	err := r.db.NewSelect().Model(&addressRec).Where("id = ?", branchId).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &addressRec, nil
}

func (r *OrderRepository) getContact(ctx context.Context, id string) (*Contacts, error) {
	contactRec := Contacts{}
	err := r.db.NewSelect().Model(&contactRec).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &contactRec, nil
}

func (r *OrderRepository) createOeHdr(ctx context.Context, oeHdr *OeHdr) error {
	_, err := r.db.NewInsert().Model(oeHdr).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create OeHdr: %w", err)
	}
	return nil
}
func (r *OrderRepository) createOeHdrSalesRep(ctx context.Context, orderNo string, salesRepId string) error {
	oeHdrSalesRep := &OeHdrSalesrep{
		OeHdrSalesrep: prophet.OeHdrSalesrep{
			OrderNumber:      orderNo,
			SalesrepId:       salesRepId,
			LastMaintainedBy: "admin",
			CommissionSplit:  100,
			PrimarySalesrep:  ptrTo("Y"),
		},
	}

	_, err := r.db.NewInsert().Model(oeHdrSalesRep).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create oeHdrSalesRep: %w", err)
	}
	return nil
}

func (r *OrderRepository) getProductDetails(ctx context.Context, invMastUid int32) (*InvMast, error) {
	invMastRec := InvMast{}

	err := r.db.NewSelect().Model(&invMastRec).Relation("InvLocs", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("location_id = ? and company_id = ?", 1001, "MRS")
	}).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &invMastRec, nil

}

func (r *OrderRepository) getShipTo(ctx context.Context, id float64) (*prophet.ShipTo, error) {
	shipToRec := prophet.ShipTo{}
	err := r.db.NewSelect().Model(&shipToRec).Where("ship_to_id = ? and company_id = ?", id, "MRS").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shipToRec, nil

}

func (r *OrderRepository) getNextOeLineLineNo(ctx context.Context, oeHdrUid int32) (int32, error) {
	var lineNo *int32
	q := `SELECT MAX(line_no) FROM oe_line WHERE oe_hdr_uid = ?`

	err := r.db.QueryRowContext(ctx, q, oeHdrUid).Scan(&lineNo)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 1, nil
		default:
			return 1, err
		}
	}
	if lineNo == nil {
		return 1, nil
	}

	return *lineNo + 1, nil

}

func ptrTo[T any](v T) *T {
	return &v
}

func getOptionalValue[T comparable](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// convertOptionalValue converts a pointer value to another type using a converter function or returns a default value if nil.
func convertOptionalValue[T any, U any](value *T, converter func(T) U, defaultValue U) U {
	if value == nil {
		return defaultValue
	}
	return converter(*value)
}

func offSetToUtc(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	//TODO: load offset from database
	return t.Add(5 * time.Hour)
}

func determineOrderStatus(oeHdrRec *OeHdr) domain.OrderStatus {
	if getOptionalValue(oeHdrRec.CancelFlag, "N") == "Y" {
		return domain.OrderStatusCancelled
	} else if getOptionalValue(oeHdrRec.Completed, "N") == "Y" {
		return domain.OrderStatusCompleted
	} else if getOptionalValue(oeHdrRec.Approved, "N") == "Y" {
		return domain.OrderStatusApproved
	} else if getOptionalValue(oeHdrRec.Approved, "N") == "N" {
		return domain.OrderStatusPendingApproval
	} else {
		return domain.OrderStatusUnknown
	}
}

func determineQuoteStatus(oeHdrRec *OeHdr) domain.QuoteStatus {
	switch {
	case getOptionalValue(oeHdrRec.CancelFlag, "N") == "Y":
		return domain.QuoteStatusCancelled
	case oeHdrRec.QuoteHdr != nil &&
		time.Now().After(getOptionalValue(oeHdrRec.QuoteHdr.ExpirationDate,
			time.Now())) || oeHdrRec.QuoteHdr.CompleteFlag == "Y":
		return domain.QuoteStatusExpired
	case getOptionalValue(oeHdrRec.Approved, "N") == "Y":
		return domain.QuoteStatusApproved
	case getOptionalValue(oeHdrRec.Approved, "N") == "N":
		return domain.QuoteStatusPendingApproval
	default:
		return domain.QuoteStatusUnknown
	}
}
