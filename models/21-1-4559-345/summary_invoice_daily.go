package model

type SummaryInvoiceDaily struct {
	bun.BaseModel          `bun:"table:summary_invoice_daily"`
	SummaryInvoiceDailyUid int32     `bun:"summary_invoice_daily_uid,type:int,pk,identity"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	BranchId               string    `bun:"branch_id,type:varchar(8),nullzero"`
	SummaryDate            time.Time `bun:"summary_date,type:datetime"`
	SummaryGroupCd         int16     `bun:"summary_group_cd,type:smallint"`
	SummaryCd              int32     `bun:"summary_cd,type:int"`
	SummaryValue           float64   `bun:"summary_value,type:decimal(19,4)"`
	SummaryHdrCount        int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount       int32     `bun:"summary_line_count,type:int"`
	SummaryAmountPaid      float64   `bun:"summary_amount_paid,type:decimal(19,4)"`
	SummaryShippingCost    float64   `bun:"summary_shipping_cost,type:decimal(19,4)"`
}
