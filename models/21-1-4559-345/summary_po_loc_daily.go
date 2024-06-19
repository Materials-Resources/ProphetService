package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type SummaryPoLocDaily struct {
	bun.BaseModel        `bun:"table:summary_po_loc_daily"`
	SummaryPoLocDailyUid int32     `bun:"summary_po_loc_daily_uid,type:int,autoincrement,pk"`
	CompanyId            string    `bun:"company_id,type:varchar(8),unique"`
	LocationId           float64   `bun:"location_id,type:decimal(19,0),unique"`
	SummaryDate          time.Time `bun:"summary_date,type:datetime,unique"`
	SummaryGroupCd       int16     `bun:"summary_group_cd,type:smallint,unique"`
	SummaryCd            int32     `bun:"summary_cd,type:int,unique"`
	SummaryValue         float64   `bun:"summary_value,type:decimal(19,9)"`
	SummaryHdrCount      int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount     int32     `bun:"summary_line_count,type:int"`
}
