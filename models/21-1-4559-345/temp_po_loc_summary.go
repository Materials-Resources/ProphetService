package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TempPoLocSummary struct {
	bun.BaseModel    `bun:"table:temp_po_loc_summary"`
	TplsId           int32     `bun:"tpls_id,type:int,identity"`
	CompanyId        string    `bun:"company_id,type:varchar(8)"`
	LocationId       float64   `bun:"location_id,type:decimal(19,0)"`
	SummaryDate      time.Time `bun:"summary_date,type:datetime"`
	SummaryGroupCode int16     `bun:"summary_group_code,type:smallint"`
	SummaryCd        int32     `bun:"summary_cd,type:int"`
	SummaryValue     float64   `bun:"summary_value,type:decimal(19,9)"`
	SummaryHdrCount  int32     `bun:"summary_hdr_count,type:int"`
	SummaryLineCount int32     `bun:"summary_line_count,type:int"`
	Processed        string    `bun:"processed,type:char"`
	Guid             string    `bun:"guid,type:uniqueidentifier,nullzero"`
}
