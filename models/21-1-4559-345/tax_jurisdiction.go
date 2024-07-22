package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxJurisdiction struct {
	bun.BaseModel               `bun:"table:tax_jurisdiction"`
	JurisdictionId              string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the unique identifier for this tax jurisdiction?
	JurisdictionDesc            string    `bun:"jurisdiction_desc,type:varchar(40)"`                            // What is this tax jurisdiction for?
	TaxPercentage               float64   `bun:"tax_percentage,type:decimal(19,8)"`                             // Indicates the tax percentage associated with this jurisdiction.
	DeleteFlag                  string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	LineItemLimit               *float64  `bun:"line_item_limit,type:decimal(19,4)"`                            // The maximum amounts to which tax can be charged per line item in a specific jurisdiction.
	InvoiceLimit                *float64  `bun:"invoice_limit,type:decimal(19,4)"`                              // The maximum amounts to which tax can be charged per invoice in a specific jurisdiction.
	FreightInTaxable            *string   `bun:"freight_in_taxable,type:char(1)"`                               // incoming freight is taxable
	FreightOutTaxable           *string   `bun:"freight_out_taxable,type:char(1)"`                              // outgoing freight is taxable
	TaxIsTaxable                string    `bun:"tax_is_taxable,type:char(1)"`                                   // Determines whether the tax charged in a particular jurisdiction is also taxable.
	TaxAmountPerUnit            *float64  `bun:"tax_amount_per_unit,type:decimal(19,9)"`                        // A flat tax rate that is applied per unit.
	UnitOfMeasure               *string   `bun:"unit_of_measure,type:varchar(8)"`                               // What is the unit of measure for this row?
	AfterUnits                  *float64  `bun:"after_units,type:decimal(19,0)"`                                // Tax amount that applies after this number of units purchased per order
	AfterMaximumUnits           *float64  `bun:"after_maximum_units,type:decimal(19,0)"`                        // Tax amount that applies after this max units purchased per month
	SeeSchedule                 *string   `bun:"see_schedule,type:char(1)"`                                     // Indicates whether the tax jurisdiction is setup to use the schedule option.
	TaxOnGrossNetQtyInvoiced    *string   `bun:"tax_on_gross_net_qty_invoiced,type:char(1)"`                    // Taxed based on gross/net/units invoiced
	CalculateTaxMethod          string    `bun:"calculate_tax_method,type:char(1)"`                             // Determines how the system will generate a tax total. (either price or cost)
	LineItemLimitOrRange        string    `bun:"line_item_limit_or_range,type:char(1),default:('L')"`           // Line Item Limit Or Range
	MinLineItemRange            *float64  `bun:"min_line_item_range,type:decimal(19,4)"`                        // Min Line Item Range
	MaxLineItemRange            *float64  `bun:"max_line_item_range,type:decimal(19,4)"`                        // Max Line Item Range
	IncomingFreight             string    `bun:"incoming_freight,type:char(1),default:('Y')"`                   // Indicate if incoming freight is taxable on the tax jurisdiction.
	OutgoingFreight             string    `bun:"outgoing_freight,type:char(1),default:('Y')"`                   // Indicate if outgoing freight is taxable on the tax jurisdiction.
	TaxTermsTakenFlag           string    `bun:"tax_terms_taken_flag,type:char(1),default:('N')"`               // Not all companies will be allowed to take terms on taxes so this needs to by tax jurisdiction setting
	ReportCodeUid               *int32    `bun:"report_code_uid,type:int"`                                      // Allows a reporting code to be associated with the jurisdiction
	OtherChargeFreightTaxable   *string   `bun:"other_charge_freight_taxable,type:char(1),default:('N')"`       // Custom column to determine taxibility of other charge items.
	TaxBeforeAfterRestockFee    *string   `bun:"tax_before_after_restock_fee,type:varchar(255)"`                // This column will indicate in RMAs, if we should calculate tax after or before adding in the Restock Fee.
	OrderLimit                  *float64  `bun:"order_limit,type:decimal(19,4)"`                                // Indicate the order total limit that could be charged tax for the jurisdiction.
	VatFlag                     *string   `bun:"vat_flag,type:char(1)"`                                         // For VAT (value added tax) enabled users, determines if this is a VAT jurisdiction.
	EcoFeesTaxableFlag          string    `bun:"eco_fees_taxable_flag,type:char(1),default:('N')"`              // If set to 'Y', Eco Fees are considered taxable within the jurisdiction
	FreightIvaWithheldFlag      string    `bun:"freight_iva_withheld_flag,type:char(1),default:('N')"`          // Determined if this juris is being used for freight iva withheld_tax
	JurisdictionTaxTypeCd       *int32    `bun:"jurisdiction_tax_type_cd,type:int"`                             // Jurisdiction Tax Type
	RentalFlag                  *string   `bun:"rental_flag,type:char(1),default:('N')"`                        // Indicates it only applies to rental items
	FlatFeeAmount               float64   `bun:"flat_fee_amount,type:decimal(19,4),default:((0))"`              // A specific tax amount charged per taxable shipment.
	ApplyFlatFeeToLinkedRmaFlag string    `bun:"apply_flat_fee_to_linked_rma_flag,type:char(1),default:('N')"`  // Indicates whether the flat fee amount should be calculated for an RMA linked line.
}
