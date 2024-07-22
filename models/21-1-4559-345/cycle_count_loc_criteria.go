package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountLocCriteria struct {
	bun.BaseModel            `bun:"table:cycle_count_loc_criteria"`
	CycleCountLocCriteriaUid int32     `bun:"cycle_count_loc_criteria_uid,type:int,pk"`                 // Unique identifier for cycle count location criteria records
	LocationId               float64   `bun:"location_id,type:decimal(19,0),unique"`                    // Where was the used material located?
	CountBy                  int32     `bun:"count_by,type:int"`                                        // Criteria to count by (item/bin)
	CyclesPerYear            int16     `bun:"cycles_per_year,type:smallint"`                            // Number of cycles in a year
	CountDaysPerYear         int16     `bun:"count_days_per_year,type:smallint"`                        // Number of count days in a year
	CountType                int32     `bun:"count_type,type:int"`                                      // Type of count (available/on hand)
	IncludeNonStock          string    `bun:"include_non_stock,type:char(1)"`                           // Include non stock items in this cycle count
	IncludeRequisition       string    `bun:"include_requisition,type:char(1)"`                         // Include requisition items in this cycle count
	DateCreated              time.Time `bun:"date_created,type:datetime"`                               // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                         // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`                      // ID of the user who last maintained this record
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                 // Indicates current record status.
	CcIncludeNullAbcRankFlag string    `bun:"cc_include_null_abc_rank_flag,type:char(1),default:('Y')"` // column to determine if null abc class or put away rank will show up in cycle count
	ExcludeZeroQohBinFlag    string    `bun:"exclude_zero_qoh_bin_flag,type:char(1),default:('N')"`
}
