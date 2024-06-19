package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type TempSalesSummary struct {
	bun.BaseModel         `bun:"table:temp_sales_summary"`
	TolsUid               int32     `bun:"tols_uid,type:int,autoincrement"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`
	SummaryGroupCode      int16     `bun:"summary_group_code,type:smallint"`
	Code                  int32     `bun:"code,type:int"`
	SummaryValue          float64   `bun:"summary_value,type:decimal(19,4)"`
	SummaryHdrCount       int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount      int32     `bun:"summary_line_count,type:int"`
	SummarySalesCost      float64   `bun:"summary_sales_cost,type:decimal(19,4)"`
	SummaryCommissionCost float64   `bun:"summary_commission_cost,type:decimal(19,4)"`
	SummaryOtherCost      float64   `bun:"summary_other_cost,type:decimal(19,4)"`
	OrderDate             time.Time `bun:"order_date,type:datetime"`
	Processed             string    `bun:"processed,type:char(1)"`
	Guid                  string    `bun:"guid,type:uniqueidentifier,nullzero"`
}
