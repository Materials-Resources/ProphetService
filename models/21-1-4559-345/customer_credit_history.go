package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerCreditHistory struct {
	bun.BaseModel             `bun:"table:customer_credit_history"`
	CustomerCreditHistoryUid  int32     `bun:"customer_credit_history_uid,type:int,autoincrement,pk"`        // Unique identifier
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`                                   // The company id associated with the customer
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0)"`                               // The customer identifier
	YearInvoiced              int32     `bun:"year_invoiced,type:int"`                                       // Year of the invoice/payment
	MonthInvoiced             int32     `bun:"month_invoiced,type:int"`                                      // Month of the invoice/payment
	HighCreditUsed            float64   `bun:"high_credit_used,type:decimal(19,2),default:(0)"`              // Highest credit used for the year/month invoiced.
	AmountPaid                float64   `bun:"amount_paid,type:decimal(19,2),default:(0)"`                   // Total amount paid for the year/month
	AvgFastSlowDays           float64   `bun:"avg_fast_slow_days,type:decimal(19,4),default:(0)"`            // Average fast/slow days
	SumDaysXPayment           float64   `bun:"sum_days_x_payment,type:decimal(19,4),default:(0)"`            // Sum of days * pymt amounts. Used for fast/slow calc.
	InvoicedSales             float64   `bun:"invoiced_sales,type:decimal(19,2),default:(0)"`                // Invoiced sales for the year/month invoiced.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	InvoicedCost              float64   `bun:"invoiced_cost,type:decimal(19,4),default:(0)"`                 // Total cost amount invoiced for the specified month.
	InvoicedOtherCharges      float64   `bun:"invoiced_other_charges,type:decimal(19,2),default:((0))"`      // The total amont of invoiced other charge for this month/year
	InvoicedOtherCost         float64   `bun:"invoiced_other_cost,type:decimal(19,9),default:((0))"`         // Total other cost amount invoiced for the specified month.
	InvoicedCommissionCost    float64   `bun:"invoiced_commission_cost,type:decimal(19,9),default:((0))"`    // Total commission cost amount invoiced for the specified month.
	FreightBilled             float64   `bun:"freight_billed,type:decimal(19,4),default:((0))"`              // Total billed freight for the specified month.
	FreightUnbilled           float64   `bun:"freight_unbilled,type:decimal(19,4),default:((0))"`            // Total unbilled freight for the specified month.
	AvgFastSlowDaysAmountPaid float64   `bun:"avg_fast_slow_days_amount_paid,type:decimal(19,4),default:((0))"`
	MerchandiseInvoicedSales  float64   `bun:"merchandise_invoiced_sales,type:decimal(19,2),nullzero"` // Custom column indicate the invoiced sales amount for merchandise item only.
	NoOfMerchandiseInvoices   int32     `bun:"no_of_merchandise_invoices,type:int,nullzero"`           // Custom column indicate the number of invoices for the time period used in calculating the value of the new “Merchandise Sales” colum
}
