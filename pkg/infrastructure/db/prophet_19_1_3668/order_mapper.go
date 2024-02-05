package prophet_19_1_3668

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
)

func ToDbOeHdr(ent *entities.ValidatedOrder) *model.OeHdr {
	var oeHdrModel = &model.OeHdr{
		OrderNo:       "",
		CustomerId:    0,
		OrderDate:     sql.NullTime{Time: time.Now(), Valid: true},
		Ship2Name:     sql.NullString{},
		Ship2Add1:     sql.NullString{},
		Ship2Add2:     sql.NullString{},
		Ship2City:     sql.NullString{},
		Ship2State:    sql.NullString{},
		Ship2Zip:      sql.NullString{},
		Ship2Country:  sql.NullString{},
		RequestedDate: sql.NullTime{},
		PoNo:          sql.NullString{String: ent.PurchaseOrder, Valid: true},
		Terms:         sql.NullString{},
		ShipToPhone:   sql.NullString{},
		DeleteFlag:    "N",
		Completed:     sql.NullString{String: "N", Valid: true},
		//TODO add CompanyId to config
		CompanyId:        sql.NullString{String: "MRS", Valid: true},
		DateCreated:      time.Now(),
		DateLastModified: time.Now(),
		LastMaintainedBy: "",
		CodFlag:          sql.NullString{String: "N", Valid: true},
		ProjectedOrder:   sql.NullString{},
		PoNoAppend:       sql.NullString{},
		//TODO add LocationId to config
		LocationId:            sql.NullFloat64{Float64: 1001, Valid: true},
		CarrierId:             sql.NullFloat64{},
		AddressId:             sql.NullFloat64{},
		ContactId:             sql.NullString{},
		HandlingChargeReqFlag: sql.NullString{String: "N", Valid: true},
		FobFlag:               sql.NullString{String: "Y"},
		RmaFlag:               sql.NullString{String: "N", Valid: true},
		Taker:                 sql.NullString{},
		ThirdPartyBillingFlag: sql.NullString{String: "S", Valid: true},
		Approved:              sql.NullString{String: "N", Valid: true},
		//TODO add SourceLocationId to config
		SourceLocationId:           sql.NullFloat64{Float64: 1001, Valid: true},
		PackingBasis:               sql.NullString{String: "Partial", Valid: true},
		DeliveryInstructions:       sql.NullString{},
		PickTicketType:             sql.NullString{},
		RequestedDownpayment:       sql.NullFloat64{Float64: 0.0000, Valid: true},
		DownpaymentInvoiced:        sql.NullString{String: "N", Valid: true},
		CancelFlag:                 sql.NullString{String: "N", Valid: true},
		WillCall:                   sql.NullString{String: "N", Valid: true},
		FrontCounter:               sql.NullString{String: "N", Valid: true},
		ValidationStatus:           sql.NullString{},
		OeHdrUid:                   0,
		SourceId:                   sql.NullString{},
		SourceCodeNo:               0,
		FreightCodeUid:             sql.NullInt32{},
		ShippingRouteUid:           sql.NullInt32{},
		ExcludeRebates:             "N",
		CaptureUsageDefault:        "Y",
		Ship2EmailAddress:          sql.NullString{},
		OrderType:                  sql.NullInt32{Int32: 706, Valid: true},
		InvoiceExchRateSourceCd:    sql.NullInt32{Int32: 222, Valid: true},
		RmaExpirationDate:          sql.NullTime{},
		CreatedBy:                  sql.NullString{},
		TagHoldCancelDate:          sql.NullTime{},
		RestockFeePercentage:       sql.NullFloat64{},
		ValidatedViaOpenOrdersFlag: sql.NullString{},
		OriginalPromiseDate:        sql.NullTime{},
		PromiseDate:                sql.NullTime{},
		DownpaymentPercentage:      sql.NullFloat64{},
		PrepaidInvoiceFlag:         sql.NullString{String: "N", Valid: true},
		LandedCostIncludedCd:       sql.NullInt32{Int32: 1423, Valid: true},
		ExpediteDate:               sql.NullTime{},
		PackingListSentFlag:        sql.NullInt32{},
		ApplyFuelSurchargeFlag:     sql.NullString{},
		ExcludeFromCreditLimitFlag: sql.NullString{String: "N", Valid: true},
		WebShopperId:               sql.NullInt32{},
		WebShopperEmail:            sql.NullString{},
		BlindShipFlag:              sql.NullString{String: "N", Valid: true},
		PtsLabelPrintFlag:          sql.NullString{String: "N", Valid: true},
		DaysEarly:                  sql.NullInt32{Int32: 0, Valid: true},
		TransitDays:                sql.NullInt32{Int32: 1, Valid: true},
		AdvancedBillingFlag:        sql.NullString{String: "N", Valid: true},
		AdvancedBillingPrintFlag:   sql.NullString{String: "N", Valid: true},
	}

	return oeHdrModel
}
func FromDBPickTicket(dbPickTicket *model.OePickTicket) (*entities.ValidatedPickTicket, error) {
	var shippingAddress = &entities.Address{
		Name:       dbPickTicket.OeHdr.Ship2Name.String,
		LineOne:    dbPickTicket.OeHdr.Ship2Add1.String,
		LineTwo:    dbPickTicket.OeHdr.Ship2Add2.String,
		City:       dbPickTicket.OeHdr.Ship2City.String,
		State:      dbPickTicket.OeHdr.Ship2State.String,
		PostalCode: dbPickTicket.OeHdr.Ship2Zip.String,
	}
	validatedShippingAddress, err := entities.NewValidatedAddress(shippingAddress)
	if err != nil {
		return nil, err
	}
	var e = &entities.PickTicket{
		ID:                 fmt.Sprintf("%.0f", dbPickTicket.PickTicketNo),
		ShippingAddress:    *validatedShippingAddress,
		OrderId:            dbPickTicket.OrderNo,
		OrderPurchaseOrder: dbPickTicket.OeHdr.PoNo.String,
		OrderContactName:   "",
	}
	return entities.NewValidatedPickTicket(e)
}

func FromDBOrder(dbOrder *model.OeHdr) (*entities.ValidatedOrder, error) {

	var shippingAddress = &entities.Address{
		Name:       dbOrder.Ship2Name.String,
		LineOne:    dbOrder.Ship2Add1.String,
		LineTwo:    dbOrder.Ship2Add2.String,
		City:       dbOrder.Ship2City.String,
		State:      dbOrder.Ship2State.String,
		PostalCode: dbOrder.Ship2Zip.String,
	}

	validatedShippingAddress, err := entities.NewValidatedAddress(shippingAddress)
	if err != nil {
		return nil, err
	}

	var customerContact = &entities.OrderCustomerContact{
		Id:          dbOrder.ContactId.String,
		Name:        dbOrder.Contact.FirstName,
		PhoneNumber: dbOrder.Contact.DirectPhone.String,
		Title:       dbOrder.Contact.Title.String,
	}

	validatedOrderCustomerContact, err := entities.NewValidatedOrderCustomerContact(customerContact)

	var o = &entities.Order{
		ShippingAddress: *validatedShippingAddress,
		PurchaseOrder:   dbOrder.PoNo.String,
		CustomerContact: *validatedOrderCustomerContact,
	}

	for _, item := range dbOrder.OeLineItems {
		var orderItem = &entities.OrderItem{
			ID:              strconv.Itoa(int(item.InvMastUid)),
			SN:              item.InvMast.ItemId,
			Name:            item.InvMast.ItemDesc,
			QuantityOrdered: item.UnitQuantity,
		}
		validatedOrderItem, err := entities.NewValidatedOrderItem(orderItem)
		if err != nil {
			//TODO handle error

			fmt.Println(err)
		}
		o.Items = append(o.Items, *validatedOrderItem)

	}

	o.ID = dbOrder.OrderNo
	return entities.NewValidatedOrder(o)
}
