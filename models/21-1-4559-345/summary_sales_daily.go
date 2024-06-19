package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type SummarySalesDaily struct {
	bun.BaseModel         `bun:"table:summary_sales_daily"`
	SummarySalesDailyUid  int32     `bun:"summary_sales_daily_uid,type:int,autoincrement,pk"`
	CompanyId             string    `bun:"company_id,type:varchar(8),unique"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0),unique"`
	SummaryDate           time.Time `bun:"summary_date,type:datetime,unique"`
	SummaryGroupCd        int16     `bun:"summary_group_cd,type:smallint,unique"`
	SummaryCd             int32     `bun:"summary_cd,type:int,unique"`
	SummaryValue          float64   `bun:"summary_value,type:decimal(19,4)"`
	SummaryHdrCount       int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount      int32     `bun:"summary_line_count,type:int"`
	SummarySalesCost      float64   `bun:"summary_sales_cost,type:decimal(19,4)"`
	SummaryCommissionCost float64   `bun:"summary_commission_cost,type:decimal(19,4)"`
	SummaryOtherCost      float64   `bun:"summary_other_cost,type:decimal(19,4)"`
}
