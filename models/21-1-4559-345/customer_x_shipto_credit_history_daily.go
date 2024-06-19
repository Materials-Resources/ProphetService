package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerXShiptoCreditHistoryDaily struct {
	bun.BaseModel                     `bun:"table:customer_x_shipto_credit_history_daily"`
	CustomerXShiptoCdtHistoryDailyUid int32     `bun:"customer_x_shipto_cdt_history_daily_uid,type:int,autoincrement,identity,pk"` // Unique identifier
	CompanyId                         string    `bun:"company_id,type:varchar(8)"`                                                 // The company id associated with the customer
	CustomerId                        float64   `bun:"customer_id,type:decimal(19,0)"`                                             // The customer identifier
	ShipToId                          float64   `bun:"ship_to_id,type:decimal(19,0)"`                                              // The ShipTo identifier
	DateInvoiced                      time.Time `bun:"date_invoiced,type:datetime"`                                                // Date of the invoice/payment
	HighCreditUsed                    float64   `bun:"high_credit_used,type:decimal(19,2),nullzero"`                               // Highest credit used for the day invoiced.
	AmountPaid                        float64   `bun:"amount_paid,type:decimal(19,2),nullzero"`                                    // Total amount paid for the day
	SumDaysXPayment                   float64   `bun:"sum_days_x_payment,type:decimal(19,4),nullzero"`                             // Sum of days * pymt amounts. Used for fast/slow calc.
	InvoicedSales                     float64   `bun:"invoiced_sales,type:decimal(19,2)"`                                          // Invoiced sales for the day invoiced.
	InvoicedCost                      float64   `bun:"invoiced_cost,type:decimal(19,4),nullzero"`                                  // Total cost amount invoiced for the specified day.
	InvoicedOtherCost                 float64   `bun:"invoiced_other_cost,type:decimal(19,4),default:((0))"`                       // Total other cost amount invoiced for the specified month.
	InvoicedCommissionCost            float64   `bun:"invoiced_commission_cost,type:decimal(19,4),default:((0))"`                  // Total commission cost amount invoiced for the specified month.
	FreightBilled                     float64   `bun:"freight_billed,type:decimal(19,4),default:((0))"`                            // Total billed freight for the specified month.
	FreightUnbilled                   float64   `bun:"freight_unbilled,type:decimal(19,4),default:((0))"`                          // Total unbilled freight for the specified month.
	AvgFastSlowDaysAmountPaid         float64   `bun:"avg_fast_slow_days_amount_paid,type:decimal(19,4),nullzero"`                 // avg_fast_slow_days_amount_paid
	DateCreated                       time.Time `bun:"date_created,type:datetime,default:(getdate())"`                             // Date and time the record was originally created
	CreatedBy                         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                       // User who created the record
	DateLastModified                  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                       // Date and time the record was modified
	LastMaintainedBy                  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`               // User who last changed the record
}
