package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type Order struct {
	bun.BaseModel        `bun:"table:oe_hdr"`
	OrderNo              string    `bun:"order_no,pk"`
	Approved             string    `bun:"approved,nullzero"`
	ContactId            string    `bun:"contact_id,nullzero"`
	CompanyId            string    `bun:"company_id"`
	Completed            string    `bun:"completed,nullzero"`
	CustomerId           float64   `bun:"customer_id"`
	DeleteFlag           string    `bun:"delete_flag"`
	DeliveryInstructions string    `bun:"delivery_instructions,nullzero"`
	DateCreated          time.Time `bun:"date_Created"`
	DateModified         time.Time `bun:"date_last_modified"`
	LastMaintainedBy     string    `bun:"last_maintained_by"`
	OrderDate            time.Time `bun:"order_date,nullzero"`
	PoNo                 string    `bun:"po_no,notnull,nullzero"`
	ShipToName           string    `bun:"ship2_name,nullzero"`
	ShipToAdd1           string    `bun:"ship2_add1,nullzero"`
	ShipToAdd2           string    `bun:"ship2_add2,nullzero"`
	ShipToCity           string    `bun:"ship2_city,nullzero"`
	ShipToState          string    `bun:"ship2_state,nullzero"`
	ShipToZip            string    `bun:"ship2_zip,nullzero"`
	ShipToCountry        string    `bun:"ship2_country,nullzero"`
	OeHdrUid             int32     `bun:"oe_hdr_uid,unique"`
	SourceCodeNo         int32     `bun:"source_code_no"`
	InvoiceBatchUid      int32     `bun:"invoice_batch_uid,default:1"`
	FreightOut           float64   `bun:"freight_out,default:0"`
	ExcludeRebates       string    `bun:"exclude_rebates"`
	CaptureUsageDefault  string    `bun:"capture_usage_default"`
	FrontCounterRma      string    `bun:"front_counter_rma,default:N"`
	ProfitPercent        float64   `bun:"profit_percent,default:0"`
	OrderCostBasis       int32     `bun:"order_cost_basis,default:1"`

	Customer Customer `bun:"rel:has-one,join:customer_id=customer_id,join_on:company_id=company_id"`

	OrderLine []*OrderLine `bun:"rel:has-many,join:order_no=order_no"`
}

type OrderLine struct {
	bun.BaseModel `bun:"table:oe_line"`

	CustomerPartId string          `bun:"customer_part_number"`
	OrderNo        string          `bun:"order_no,pk"`
	LineNo         float32         `bun:"line_no,pk"`
	DeleteFlag     string          `bun:"delete_flag"`
	InvMastUid     int32           `bun:"inv_mast_uid"`
	RequiredDate   sql.NullTime    `bun:"required_date"`
	ExpediteDate   sql.NullTime    `bun:"expedite_date"`
	QtyOrdered     float64         `bun:"qty_ordered"`
	ExtendedPrice  sql.NullFloat64 `bun:"extended_price"`
	UnitOfMeasure  sql.NullString  `bun:"unit_of_measure"`
}
