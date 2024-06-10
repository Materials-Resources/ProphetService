package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type OePickTicket struct {
	bun.BaseModel            `bun:"table:oe_pick_ticket"`
	CompanyId                string          `bun:"company_id,type:varchar(8)"`
	PickTicketNo             float64         `bun:"pick_ticket_no,type:decimal(19,0),pk"`
	OrderNo                  string          `bun:"order_no,type:varchar(8)"`
	PrintDate                sql.NullTime    `bun:"print_date,type:datetime,nullzero"`
	CarrierId                sql.NullFloat64 `bun:"carrier_id,type:decimal(19,0),nullzero"`
	TrackingNo               sql.NullString  `bun:"tracking_no,type:varchar(40),nullzero"`
	Instructions             sql.NullString  `bun:"instructions,type:varchar(255),nullzero"`
	FreightOut               sql.NullFloat64 `bun:"freight_out,type:decimal(19,4),nullzero"`
	FreightIn                sql.NullFloat64 `bun:"freight_in,type:decimal(19,4),nullzero"`
	ShipDate                 sql.NullTime    `bun:"ship_date,type:datetime,nullzero"`
	InvoiceNo                sql.NullFloat64 `bun:"invoice_no,type:decimal(19,0),nullzero"`
	PickAndHold              string          `bun:"pick_and_hold,type:char"`
	DateCreated              time.Time       `bun:"date_created,type:datetime"`
	DateLastModified         time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrintedFlag              sql.NullString  `bun:"printed_flag,type:char,default:('Y')"`
	LocationId               float64         `bun:"location_id,type:decimal(19,0)"`
	DeleteFlag               string          `bun:"delete_flag,type:char"`
	DirectShipment           sql.NullString  `bun:"direct_shipment,type:char,nullzero"`
	FrontCounterTax          sql.NullString  `bun:"front_counter_tax,type:char,nullzero"`
	TotalTax                 sql.NullFloat64 `bun:"total_tax,type:decimal(19,2),nullzero"`
	FrontCounter             string          `bun:"front_counter,type:char"`
	InvoiceIdWhenShipped     sql.NullString  `bun:"invoice_id_when_shipped,type:varchar(10),nullzero"`
	Auxiliary                string          `bun:"auxiliary,type:char,default:('N')"`
	InvoiceBatchUid          int32           `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid           sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	DeliveryListStatus       int32           `bun:"delivery_list_status,type:int,default:(1024)"`
	RetrievedByWms           string          `bun:"retrieved_by_wms,type:char,default:('N')"`
	Picker                   sql.NullString  `bun:"picker,type:varchar(255),nullzero"`
	LineItemSortSequence     int32           `bun:"line_item_sort_sequence,type:int,default:(1)"`
	ConfirmableRowStatusFlag int32           `bun:"confirmable_row_status_flag,type:int,default:('1980')"`
	WeightOverrideFlag       sql.NullString  `bun:"weight_override_flag,type:char,nullzero"`
	OverrideValue            sql.NullFloat64 `bun:"override_value,type:decimal(19,9),nullzero"`
	CreatedBy                string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OePickTicketTypeCd       sql.NullInt32   `bun:"oe_pick_ticket_type_cd,type:int,nullzero"`
	OutgoingFreightCost      sql.NullFloat64 `bun:"outgoing_freight_cost,type:decimal(19,9),nullzero"`
	QualityControl           sql.NullString  `bun:"quality_control,type:varchar(255),nullzero"`
	PackingListPrintedFlag   sql.NullString  `bun:"packing_list_printed_flag,type:char,nullzero"`
	DeaPickTicketFlag        sql.NullString  `bun:"dea_pick_ticket_flag,type:char,nullzero"`
	Packer                   sql.NullString  `bun:"packer,type:varchar(255),nullzero"`
	Checker                  sql.NullString  `bun:"checker,type:varchar(255),nullzero"`
	ReviewShipmentFlag       sql.NullString  `bun:"review_shipment_flag,type:char,default:('N')"`
	OriginalFreightOut       sql.NullFloat64 `bun:"original_freight_out,type:decimal(19,9),nullzero"`
	ExportedToRouterFlag     sql.NullString  `bun:"exported_to_router_flag,type:char,nullzero"`
	PriceConfirmedFlag       sql.NullString  `bun:"price_confirmed_flag,type:char,nullzero"`
	ShipConfirmedFlag        sql.NullString  `bun:"ship_confirmed_flag,type:char,nullzero"`
	OrigSortOrder            sql.NullInt32   `bun:"orig_sort_order,type:int,nullzero"`
	VicsBolNumber            sql.NullString  `bun:"vics_bol_number,type:varchar(17),nullzero"`
	OutFreightCost           sql.NullFloat64 `bun:"out_freight_cost,type:decimal(19,9),nullzero"`
	FreightOutEditedFlag     sql.NullString  `bun:"freight_out_edited_flag,type:char,nullzero"`
	PrintPricesOnPackinglist sql.NullString  `bun:"print_prices_on_packinglist,type:char,nullzero"`
	PregeneratedInvoiceNo    sql.NullFloat64 `bun:"pregenerated_invoice_no,type:decimal(19,0),nullzero"`
	PrepaidFreightOut        sql.NullFloat64 `bun:"prepaid_freight_out,type:decimal(19,9),nullzero"`
	PrintCanadianB3FormsFlag string          `bun:"print_canadian_b3_forms_flag,type:char,default:('N')"`
	ProNumber                sql.NullString  `bun:"pro_number,type:varchar(255),nullzero"`
	PackingListPrintBy       sql.NullString  `bun:"packing_list_print_by,type:varchar(30),nullzero"`
	PackingListPrintDate     sql.NullTime    `bun:"packing_list_print_date,type:datetime,nullzero"`
	SlabFlag                 string          `bun:"slab_flag,type:char,default:('N')"`
	ShippingAccount          sql.NullString  `bun:"shipping_account,type:varchar(255),nullzero"`
	SidNo                    sql.NullString  `bun:"sid_no,type:varchar(255),nullzero"`
	SentToAttDate            sql.NullTime    `bun:"sent_to_att_date,type:datetime,nullzero"`
	UserConfirmedPickFlag    sql.NullString  `bun:"user_confirmed_pick_flag,type:char,nullzero"`
	Scan                     sql.NullString  `bun:"scan,type:varchar(255),nullzero"`
	ActualFedexFreightOut    sql.NullFloat64 `bun:"actual_fedex_freight_out,type:decimal(19,4),nullzero"`
	DiffFedexChargeFlag      string          `bun:"diff_fedex_charge_flag,type:char,default:('N')"`
	PickConfirmedFlag        sql.NullString  `bun:"pick_confirmed_flag,type:char,nullzero"`
	SentToTrackaboutFlag     sql.NullString  `bun:"sent_to_trackabout_flag,type:char,nullzero"`
	ShippingRouteUid         sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	BolNumber                sql.NullString  `bun:"bol_number,type:varchar(255),nullzero"`
	DriverId                 sql.NullString  `bun:"driver_id,type:varchar(16),nullzero"`
	SalesMarketGroupUid      sql.NullInt32   `bun:"sales_market_group_uid,type:int,nullzero"`
	ArnNumber                sql.NullString  `bun:"arn_number,type:varchar(255),nullzero"`
	ScanPackUid              sql.NullInt32   `bun:"scan_pack_uid,type:int,nullzero"`
	ReferenceNo              sql.NullString  `bun:"reference_no,type:varchar(255),nullzero"`
	PickCompleteFlag         sql.NullString  `bun:"pick_complete_flag,type:char,nullzero"`
	RfnavStopNo              sql.NullInt32   `bun:"rfnav_stop_no,type:int,nullzero"`
	RfnavPickStatusCd        sql.NullInt32   `bun:"rfnav_pick_status_cd,type:int,nullzero"`

	OeHdr   *OeHdr   `bun:"rel:belongs-to,join:order_no=order_no"`
	Carrier *Address `bun:"rel:has-one,join:carrier_id=id,join_on:carrier_flag='Y'"`
}

type OePickTicketModel struct {
	bun bun.IDB
}

func (m OePickTicketModel) Get(ctx context.Context, pickTicketNo float64) (*OePickTicket, error) {
	var pickTicket OePickTicket
	err := m.bun.NewSelect().Model(&pickTicket).
		Relation("OeHdr").
		Where("pick_ticket_no = ?", pickTicketNo).Scan(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &pickTicket, nil
}

type OePickTicketGetAllParams struct {
	OrderNo *[]string
}

func (m OePickTicketModel) GetAll(ctx context.Context, params OePickTicketGetAllParams) ([]*OePickTicket, error) {
	var pickTickets []*OePickTicket
	query := m.bun.NewSelect().Model(&pickTickets)

	m.applyCarrierRelation(query)

	if params.OrderNo != nil {
		query.Where("order_no IN (?)", bun.In(*params.OrderNo))
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return pickTickets, nil
}

func (m OePickTicketModel) applyCarrierRelation(q *bun.SelectQuery) {
	q.Relation("Carrier", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Column("name")
	})

}
