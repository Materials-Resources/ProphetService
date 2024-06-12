package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxJurisdiction struct {
	bun.BaseModel               `bun:"table:tax_jurisdiction"`
	JurisdictionId              string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	JurisdictionDesc            string    `bun:"jurisdiction_desc,type:varchar(40)"`
	TaxPercentage               float64   `bun:"tax_percentage,type:decimal(19,8)"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	LineItemLimit               float64   `bun:"line_item_limit,type:decimal(19,4),nullzero"`
	InvoiceLimit                float64   `bun:"invoice_limit,type:decimal(19,4),nullzero"`
	FreightInTaxable            string    `bun:"freight_in_taxable,type:char,nullzero"`
	FreightOutTaxable           string    `bun:"freight_out_taxable,type:char,nullzero"`
	TaxIsTaxable                string    `bun:"tax_is_taxable,type:char"`
	TaxAmountPerUnit            float64   `bun:"tax_amount_per_unit,type:decimal(19,9),nullzero"`
	UnitOfMeasure               string    `bun:"unit_of_measure,type:varchar(8),nullzero"`
	AfterUnits                  float64   `bun:"after_units,type:decimal(19,0),nullzero"`
	AfterMaximumUnits           float64   `bun:"after_maximum_units,type:decimal(19,0),nullzero"`
	SeeSchedule                 string    `bun:"see_schedule,type:char,nullzero"`
	TaxOnGrossNetQtyInvoiced    string    `bun:"tax_on_gross_net_qty_invoiced,type:char,nullzero"`
	CalculateTaxMethod          string    `bun:"calculate_tax_method,type:char"`
	LineItemLimitOrRange        string    `bun:"line_item_limit_or_range,type:char,default:('L')"`
	MinLineItemRange            float64   `bun:"min_line_item_range,type:decimal(19,4),nullzero"`
	MaxLineItemRange            float64   `bun:"max_line_item_range,type:decimal(19,4),nullzero"`
	IncomingFreight             string    `bun:"incoming_freight,type:char,default:('Y')"`
	OutgoingFreight             string    `bun:"outgoing_freight,type:char,default:('Y')"`
	TaxTermsTakenFlag           string    `bun:"tax_terms_taken_flag,type:char,default:('N')"`
	ReportCodeUid               int32     `bun:"report_code_uid,type:int,nullzero"`
	OtherChargeFreightTaxable   string    `bun:"other_charge_freight_taxable,type:char,default:('N')"`
	TaxBeforeAfterRestockFee    string    `bun:"tax_before_after_restock_fee,type:varchar(255),nullzero"`
	OrderLimit                  float64   `bun:"order_limit,type:decimal(19,4),nullzero"`
	VatFlag                     string    `bun:"vat_flag,type:char,nullzero"`
	EcoFeesTaxableFlag          string    `bun:"eco_fees_taxable_flag,type:char,default:('N')"`
	FreightIvaWithheldFlag      string    `bun:"freight_iva_withheld_flag,type:char,default:('N')"`
	JurisdictionTaxTypeCd       int32     `bun:"jurisdiction_tax_type_cd,type:int,nullzero"`
	RentalFlag                  string    `bun:"rental_flag,type:char,default:('N')"`
	FlatFeeAmount               float64   `bun:"flat_fee_amount,type:decimal(19,4),default:((0))"`
	ApplyFlatFeeToLinkedRmaFlag string    `bun:"apply_flat_fee_to_linked_rma_flag,type:char,default:('N')"`
}
