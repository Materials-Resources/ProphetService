package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerCreditHistory struct {
	bun.BaseModel             `bun:"table:customer_credit_history"`
	CustomerCreditHistoryUid  int32     `bun:"customer_credit_history_uid,type:int,pk,identity"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0)"`
	YearInvoiced              int32     `bun:"year_invoiced,type:int"`
	MonthInvoiced             int32     `bun:"month_invoiced,type:int"`
	HighCreditUsed            float64   `bun:"high_credit_used,type:decimal(19,2),default:(0)"`
	AmountPaid                float64   `bun:"amount_paid,type:decimal(19,2),default:(0)"`
	AvgFastSlowDays           float64   `bun:"avg_fast_slow_days,type:decimal(19,4),default:(0)"`
	SumDaysXPayment           float64   `bun:"sum_days_x_payment,type:decimal(19,4),default:(0)"`
	InvoicedSales             float64   `bun:"invoiced_sales,type:decimal(19,2),default:(0)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	InvoicedCost              float64   `bun:"invoiced_cost,type:decimal(19,4),default:(0)"`
	InvoicedOtherCharges      float64   `bun:"invoiced_other_charges,type:decimal(19,2),default:((0))"`
	InvoicedOtherCost         float64   `bun:"invoiced_other_cost,type:decimal(19,9),default:((0))"`
	InvoicedCommissionCost    float64   `bun:"invoiced_commission_cost,type:decimal(19,9),default:((0))"`
	FreightBilled             float64   `bun:"freight_billed,type:decimal(19,4),default:((0))"`
	FreightUnbilled           float64   `bun:"freight_unbilled,type:decimal(19,4),default:((0))"`
	AvgFastSlowDaysAmountPaid float64   `bun:"avg_fast_slow_days_amount_paid,type:decimal(19,4),default:((0))"`
	MerchandiseInvoicedSales  float64   `bun:"merchandise_invoiced_sales,type:decimal(19,2),nullzero"`
	NoOfMerchandiseInvoices   int32     `bun:"no_of_merchandise_invoices,type:int,nullzero"`
}
