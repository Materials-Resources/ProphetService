package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrSalesrep struct {
	bun.BaseModel              `bun:"table:invoice_hdr_salesrep"`
	InvoiceNumber              string    `bun:"invoice_number,type:varchar(10),unique"`                      // Indicates the invoice number.
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),unique"`                         // The sales representative associated with this invoice.
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`                                  // Unique code that identifies a company.
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`                       // The vendor associated with this invoice.
	CommissionAmount           float64   `bun:"commission_amount,type:decimal(19,4)"`                        // Amount of commission paid to the salesrep.
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                    // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                  // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                            // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`   // ID of the user who last maintained this record
	DatePaid                   time.Time `bun:"date_paid,type:datetime,nullzero"`                            // Indicates the date when the sales representative was paid for this invoice.
	CommissionPaid             float64   `bun:"commission_paid,type:decimal(19,4),nullzero"`                 // Indicates whether commission was paid to the sales representative for this invoice.
	PrimarySalesrep            string    `bun:"primary_salesrep,type:char(1),default:('N')"`                 // Indicates the primary salesrep.
	CommissionPercentage       float64   `bun:"commission_percentage,type:decimal(19,4),default:((100))"`    // Percentage of commission split among the salesreps per invoice
	SalesrepSourceCd           int32     `bun:"salesrep_source_cd,type:int,nullzero"`                        // Indicate where the salesrep came from
	EditedCommissionFlag       string    `bun:"edited_commission_flag,type:char(1),nullzero"`                // Indicates whether commission amount was edited by user
	PriorCommissionAmount      float64   `bun:"prior_commission_amount,type:decimal(19,2),nullzero"`         // Total commission amt from last commission saved
	SplitFlag                  string    `bun:"split_flag,type:char(1),nullzero"`                            // This column indicates this record was created during a split commission  for this salesrep and invoice
	InvoiceHdrSalesrepUid      int32     `bun:"invoice_hdr_salesrep_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the record
	CommissionScheduleUid      int32     `bun:"commission_schedule_uid,type:int,nullzero"`                   // Unique record for a schedule
	CommissionCost             float64   `bun:"commission_cost,type:decimal(19,4),nullzero"`                 // Total commission cost used when calculating commission
	ExtendedPrice              float64   `bun:"extended_price,type:decimal(19,4),nullzero"`                  // Total price used when calculating commissions
	CalculationFromSplitFlag   string    `bun:"calculation_from_split_flag,type:char(1),default:('N')"`      // Indicates whether this commission is a result of a split and the salesrep previously existed for this invoice
	ExcludeSplitValidationFlag string    `bun:"exclude_split_validation_flag,type:char(1),default:('Y')"`    // Excludes this salesrep from any validation that requires the salesreps' total commission_percentage to add up to 100%.
	PartialInvoiceOnlyFlag     string    `bun:"partial_invoice_only_flag,type:char(1),nullzero"`             // If 'Y' then this salesrep should only calculate commissions on those invoice lines to which they are assigned to in invoice_line_salesrep.
	ForfeitedAmount            float64   `bun:"forfeited_amount,type:decimal(19,9),nullzero"`                // Amount that was forfeited because of past due invoices.
	LinePercentageEditedFlag   string    `bun:"line_percentage_edited_flag,type:char(1),nullzero"`           // Indicates that there is an override of the commission percentage on oe_line_salesrep for the order for this invoice.
	CommissionOverridePercent  float64   `bun:"commission_override_percent,type:decimal(19,9),nullzero"`     // Percent to override the normal commission calculation for the associated salesrep.
	ExtendedPriceNoCndAdjust   float64   `bun:"extended_price_no_cnd_adjust,type:decimal(19,4),nullzero"`    // custom column that stores Extended Price prior to Canadian adjustment
}
