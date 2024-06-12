package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PoLine struct {
	bun.BaseModel            `bun:"table:po_line"`
	PoNo                     float64   `bun:"po_no,type:decimal(19,0),pk"`
	QtyOrdered               float64   `bun:"qty_ordered,type:decimal(19,9)"`
	QtyReceived              float64   `bun:"qty_received,type:decimal(19,9)"`
	ReceivedDate             time.Time `bun:"received_date,type:datetime,nullzero"`
	UnitPrice                float64   `bun:"unit_price,type:decimal(19,6)"`
	CompanyNo                string    `bun:"company_no,type:varchar(8)"`
	MfgPartNo                string    `bun:"mfg_part_no,type:varchar(20),nullzero"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateDue                  time.Time `bun:"date_due,type:datetime,nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	NextDueInPoCost          float64   `bun:"next_due_in_po_cost,type:decimal(19,9),nullzero"`
	Complete                 string    `bun:"complete,type:char"`
	VouchCompleted           string    `bun:"vouch_completed,type:char"`
	CancelFlag               string    `bun:"cancel_flag,type:char,nullzero"`
	InBoundCurryId           float64   `bun:"in_bound_curry_id,type:decimal(19,0),nullzero"`
	AccountNo                string    `bun:"account_no,type:varchar(32),nullzero"`
	QtyToVouch               float64   `bun:"qty_to_vouch,type:decimal(19,9),nullzero"`
	ClosedFlag               string    `bun:"closed_flag,type:char,nullzero"`
	ItemDescription          string    `bun:"item_description,type:varchar(40),nullzero"`
	UnitOfMeasure            string    `bun:"unit_of_measure,type:varchar(8),nullzero"`
	UnitSize                 float64   `bun:"unit_size,type:decimal(19,4)"`
	UnitQuantity             float64   `bun:"unit_quantity,type:decimal(19,9)"`
	LineNo                   float64   `bun:"line_no,type:decimal(19,0),pk"`
	PricingBookId            string    `bun:"pricing_book_id,type:varchar(8),nullzero"`
	PricingBookItemId        string    `bun:"pricing_book_item_id,type:varchar(40),nullzero"`
	PricingBookSupplierId    float64   `bun:"pricing_book_supplier_id,type:decimal(19,0),nullzero"`
	PricingBookDiscGrpId     string    `bun:"pricing_book_disc_grp_id,type:varchar(8),nullzero"`
	PricingBookEffectiveDate time.Time `bun:"pricing_book_effective_date,type:datetime,nullzero"`
	Combinable               string    `bun:"combinable,type:char,nullzero"`
	CalcType                 string    `bun:"calc_type,type:varchar(10),nullzero"`
	CalcValue                float64   `bun:"calc_value,type:decimal(19,9),nullzero"`
	RequiredDate             time.Time `bun:"required_date,type:datetime,nullzero"`
	NextBreak                float64   `bun:"next_break,type:decimal(19,9),nullzero"`
	NextUtPrice              float64   `bun:"next_ut_price,type:decimal(19,9),nullzero"`
	BaseUtPrice              float64   `bun:"base_ut_price,type:decimal(19,4),default:(0)"`
	PriceEdit                string    `bun:"price_edit,type:char,nullzero"`
	NewItem                  string    `bun:"new_item,type:char,nullzero"`
	QuantityChanged          string    `bun:"quantity_changed,type:char,nullzero"`
	PricingUnit              string    `bun:"pricing_unit,type:varchar(8),nullzero"`
	PricingUnitSize          float64   `bun:"pricing_unit_size,type:decimal(19,4),nullzero"`
	ExtendedDesc             string    `bun:"extended_desc,type:varchar(255),nullzero"`
	UnitPriceDisplay         float64   `bun:"unit_price_display,type:decimal(19,6)"`
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`
	ExcludeFromLeadTime      string    `bun:"exclude_from_lead_time,type:char,default:('N')"`
	SourceType               int32     `bun:"source_type,type:int,nullzero"`
	ExpDateUpdates           int32     `bun:"exp_date_updates,type:int,nullzero"`
	PoLineUid                int32     `bun:"po_line_uid,type:int"`
	EdiNewStatus             string    `bun:"edi_new_status,type:char,default:('N')"`
	LineType                 string    `bun:"line_type,type:char,nullzero"`
	ContractNumber           string    `bun:"contract_number,type:varchar(40),nullzero"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ParentPoLineNo           int32     `bun:"parent_po_line_no,type:int,nullzero"`
	SupplierShipDate         time.Time `bun:"supplier_ship_date,type:datetime,nullzero"`
	EnteredAsCode            string    `bun:"entered_as_code,type:varchar(40),nullzero"`
	GporRunUid               int32     `bun:"gpor_run_uid,type:int,nullzero"`
	PurchasePricingPageUid   int32     `bun:"purchase_pricing_page_uid,type:int,nullzero"`
	ExpediteFlag             string    `bun:"expedite_flag,type:char,nullzero"`
	OriginalUnitPriceDisplay float64   `bun:"original_unit_price_display,type:decimal(19,9),nullzero"`
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	ExpediteNotes            string    `bun:"expedite_notes,type:varchar(8000),nullzero"`
	ExpediteFollowupFlag     string    `bun:"expedite_followup_flag,type:char,default:('N')"`
	DesiredReceiptLocationId float64   `bun:"desired_receipt_location_id,type:decimal(19,0),nullzero"`
	CountryOfOrigin          string    `bun:"country_of_origin,type:char(8),nullzero"`
	AcknowledgedDate         time.Time `bun:"acknowledged_date,type:datetime,nullzero"`
	B3Qty                    float64   `bun:"b3_qty,type:decimal(19,9),default:((0))"`
	BulkBuyFlag              string    `bun:"bulk_buy_flag,type:char,nullzero"`
	CadPurchaseCost          float64   `bun:"cad_purchase_cost,type:decimal(19,9),nullzero"`
	QtyReady                 float64   `bun:"qty_ready,type:decimal(19,9),nullzero"`
	QtyReadyUnitSize         float64   `bun:"qty_ready_unit_size,type:decimal(19,9),nullzero"`
	QtyReadyUom              string    `bun:"qty_ready_uom,type:varchar(8),nullzero"`
	UnitQtyReady             float64   `bun:"unit_qty_ready,type:decimal(19,9),nullzero"`
	ListPriceMultiplier      float64   `bun:"list_price_multiplier,type:decimal(19,9),nullzero"`
	CarrierStatus            string    `bun:"carrier_status,type:varchar(20),nullzero"`
	ExpectedShipDate         time.Time `bun:"expected_ship_date,type:datetime,nullzero"`
	DateDueLastModified      time.Time `bun:"date_due_last_modified,type:datetime,nullzero"`
	Acknowledged             string    `bun:"acknowledged,type:char,nullzero"`
}
