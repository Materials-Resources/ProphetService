package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountLocCriteria struct {
	bun.BaseModel            `bun:"table:cycle_count_loc_criteria"`
	CycleCountLocCriteriaUid int32     `bun:"cycle_count_loc_criteria_uid,type:int,pk"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	CountBy                  int32     `bun:"count_by,type:int"`
	CyclesPerYear            int16     `bun:"cycles_per_year,type:smallint"`
	CountDaysPerYear         int16     `bun:"count_days_per_year,type:smallint"`
	CountType                int32     `bun:"count_type,type:int"`
	IncludeNonStock          string    `bun:"include_non_stock,type:char"`
	IncludeRequisition       string    `bun:"include_requisition,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	CcIncludeNullAbcRankFlag string    `bun:"cc_include_null_abc_rank_flag,type:char,default:('Y')"`
	ExcludeZeroQohBinFlag    string    `bun:"exclude_zero_qoh_bin_flag,type:char,default:('N')"`
}
