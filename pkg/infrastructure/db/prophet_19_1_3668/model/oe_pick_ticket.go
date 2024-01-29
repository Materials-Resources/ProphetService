package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type OePickTicket struct {
	bun.BaseModel            `bun:"table:oe_pick_ticket"`
	CompanyId                string          `bun:"company_id,type:varchar(8)"`
	PickTicketNo             float64         `bun:"pick_ticket_no,type:decimal(19,0),pk"`
	OrderNo                  string          `bun:"order_no,type:varchar(8)"`
	PrintDate                sql.NullTime    `bun:"print_date,type:datetime"`
	CarrierId                sql.NullFloat64 `bun:"carrier_id,type:decimal(19,0)"`
	TrackingNo               sql.NullString  `bun:"tracking_no,type:varchar(40)"`
	Instructions             sql.NullString  `bun:"instructions,type:varchar(255)"`
	FreightOut               sql.NullFloat64 `bun:"freight_out,type:decimal(19,4)"`
	FreightIn                sql.NullFloat64 `bun:"freight_in,type:decimal(19,4)"`
	ShipDate                 sql.NullTime    `bun:"ship_date,type:datetime"`
	InvoiceNo                sql.NullFloat64 `bun:"invoice_no,type:decimal(19,0)"`
	PickAndHold              string          `bun:"pick_and_hold,type:char"`
	DateCreated              time.Time       `bun:"date_created,type:datetime"`
	DateLastModified         time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrintedFlag              sql.NullString  `bun:"printed_flag,type:char,default:('Y')"`
	LocationId               float64         `bun:"location_id,type:decimal(19,0)"`
	DeleteFlag               string          `bun:"delete_flag,type:char"`
	DirectShipment           sql.NullString  `bun:"direct_shipment,type:char"`
	FrontCounterTax          sql.NullString  `bun:"front_counter_tax,type:char"`
	TotalTax                 sql.NullFloat64 `bun:"total_tax,type:decimal(19,2)"`
	FrontCounter             string          `bun:"front_counter,type:char"`
	InvoiceIdWhenShipped     sql.NullString  `bun:"invoice_id_when_shipped,type:varchar(10)"`
	Auxiliary                string          `bun:"auxiliary,type:char,default:('N')"`
	InvoiceBatchUid          int32           `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid           sql.NullInt32   `bun:"freight_code_uid,type:int"`
	DeliveryListStatus       int32           `bun:"delivery_list_status,type:int,default:(1024)"`
	RetrievedByWms           string          `bun:"retrieved_by_wms,type:char,default:('N')"`
	Picker                   sql.NullString  `bun:"picker,type:varchar(255)"`
	LineItemSortSequence     int32           `bun:"line_item_sort_sequence,type:int,default:(1)"`
	ConfirmableRowStatusFlag int32           `bun:"confirmable_row_status_flag,type:int,default:('1980')"`
	WeightOverrideFlag       sql.NullString  `bun:"weight_override_flag,type:char"`
	OverrideValue            sql.NullFloat64 `bun:"override_value,type:decimal(19,9)"`
	CreatedBy                string          `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OePickTicketTypeCd       sql.NullInt32   `bun:"oe_pick_ticket_type_cd,type:int"`
	OutgoingFreightCost      sql.NullFloat64 `bun:"outgoing_freight_cost,type:decimal(19,9)"`
	QualityControl           sql.NullString  `bun:"quality_control,type:varchar(255)"`
	PackingListPrintedFlag   sql.NullString  `bun:"packing_list_printed_flag,type:char"`
	DeaPickTicketFlag        sql.NullString  `bun:"dea_pick_ticket_flag,type:char"`
	Packer                   sql.NullString  `bun:"packer,type:varchar(255)"`
	Checker                  sql.NullString  `bun:"checker,type:varchar(255)"`
	ReviewShipmentFlag       sql.NullString  `bun:"review_shipment_flag,type:char,default:('N')"`
	OriginalFreightOut       sql.NullFloat64 `bun:"original_freight_out,type:decimal(19,9)"`
	ExportedToRouterFlag     sql.NullString  `bun:"exported_to_router_flag,type:char"`
	PriceConfirmedFlag       sql.NullString  `bun:"price_confirmed_flag,type:char"`
	ShipConfirmedFlag        sql.NullString  `bun:"ship_confirmed_flag,type:char"`
	OrigSortOrder            sql.NullInt32   `bun:"orig_sort_order,type:int"`
	VicsBolNumber            sql.NullString  `bun:"vics_bol_number,type:varchar(17)"`
	OutFreightCost           sql.NullFloat64 `bun:"out_freight_cost,type:decimal(19,9)"`
	FreightOutEditedFlag     sql.NullString  `bun:"freight_out_edited_flag,type:char"`
	PrintPricesOnPackinglist sql.NullString  `bun:"print_prices_on_packinglist,type:char"`
	PregeneratedInvoiceNo    sql.NullFloat64 `bun:"pregenerated_invoice_no,type:decimal(19,0)"`
	PrepaidFreightOut        sql.NullFloat64 `bun:"prepaid_freight_out,type:decimal(19,9)"`
	PrintCanadianB3FormsFlag string          `bun:"print_canadian_b3_forms_flag,type:char,default:('N')"`
	ProNumber                sql.NullString  `bun:"pro_number,type:varchar(255)"`
	PackingListPrintBy       sql.NullString  `bun:"packing_list_print_by,type:varchar(30)"`
	PackingListPrintDate     sql.NullTime    `bun:"packing_list_print_date,type:datetime"`
	SlabFlag                 string          `bun:"slab_flag,type:char,default:('N')"`
	ShippingAccount          sql.NullString  `bun:"shipping_account,type:varchar(255)"`
	SidNo                    sql.NullString  `bun:"sid_no,type:varchar(255)"`
	SentToAttDate            sql.NullTime    `bun:"sent_to_att_date,type:datetime"`
	UserConfirmedPickFlag    sql.NullString  `bun:"user_confirmed_pick_flag,type:char"`
	Scan                     sql.NullString  `bun:"scan,type:varchar(255)"`
	ActualFedexFreightOut    sql.NullFloat64 `bun:"actual_fedex_freight_out,type:decimal(19,4)"`
	DiffFedexChargeFlag      string          `bun:"diff_fedex_charge_flag,type:char,default:('N')"`
	PickConfirmedFlag        sql.NullString  `bun:"pick_confirmed_flag,type:char"`
	SentToTrackaboutFlag     sql.NullString  `bun:"sent_to_trackabout_flag,type:char"`
	ShippingRouteUid         sql.NullInt32   `bun:"shipping_route_uid,type:int"`
	BolNumber                sql.NullString  `bun:"bol_number,type:varchar(255)"`
	DriverId                 sql.NullString  `bun:"driver_id,type:varchar(16)"`
	SalesMarketGroupUid      sql.NullInt32   `bun:"sales_market_group_uid,type:int"`
	ArnNumber                sql.NullString  `bun:"arn_number,type:varchar(255)"`
	ScanPackUid              sql.NullInt32   `bun:"scan_pack_uid,type:int"`
	ReferenceNo              sql.NullString  `bun:"reference_no,type:varchar(255)"`

	OeHdr OeHdr `bun:"rel:has-one,join:order_no=order_no"`
}
