package model

import (
	"github.com/uptrace/bun"
	"time"
)

type SummaryPoLocDaily struct {
	bun.BaseModel        `bun:"table:summary_po_loc_daily"`
	SummaryPoLocDailyUid int32     `bun:"summary_po_loc_daily_uid,type:int,pk,identity"`
	CompanyId            string    `bun:"company_id,type:varchar(8)"`
	LocationId           float64   `bun:"location_id,type:decimal(19,0)"`
	SummaryDate          time.Time `bun:"summary_date,type:datetime"`
	SummaryGroupCd       int16     `bun:"summary_group_cd,type:smallint"`
	SummaryCd            int32     `bun:"summary_cd,type:int"`
	SummaryValue         float64   `bun:"summary_value,type:decimal(19,9)"`
	SummaryHdrCount      int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount     int32     `bun:"summary_line_count,type:int"`
}
