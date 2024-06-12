package model

type CustomerCreditHistoryDaily struct {
	bun.BaseModel              `bun:"table:customer_credit_history_daily"`
	CustomerCdtHistoryDailyUid int32     `bun:"customer_cdt_history_daily_uid,type:int,pk,identity"`
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0)"`
	DateInvoiced               time.Time `bun:"date_invoiced,type:datetime"`
	HighCreditUsed             float64   `bun:"high_credit_used,type:decimal(19,2),nullzero"`
	AmountPaid                 float64   `bun:"amount_paid,type:decimal(19,2),nullzero"`
	SumDaysXPayment            float64   `bun:"sum_days_x_payment,type:decimal(19,4),nullzero"`
	InvoicedSales              float64   `bun:"invoiced_sales,type:decimal(19,2),nullzero"`
	InvoicedCost               float64   `bun:"invoiced_cost,type:decimal(19,4),nullzero"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	InvoicedOtherCost          float64   `bun:"invoiced_other_cost,type:decimal(19,9),default:((0))"`
	InvoicedCommissionCost     float64   `bun:"invoiced_commission_cost,type:decimal(19,9),default:((0))"`
	FreightBilled              float64   `bun:"freight_billed,type:decimal(19,4),default:((0))"`
	FreightUnbilled            float64   `bun:"freight_unbilled,type:decimal(19,4),default:((0))"`
	AvgFastSlowDaysAmountPaid  float64   `bun:"avg_fast_slow_days_amount_paid,type:decimal(19,4),nullzero"`
}
