package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Metrics struct {
	bun.BaseModel         `bun:"table:metrics"`
	MetricsUid            int32     `bun:"metrics_uid,type:int,autoincrement,pk"`                                                                                                                                                                                           // metrics_uid
	UniqueName            string    `bun:"unique_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                       // unique_name
	Name                  string    `bun:"name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                                       // name
	Description           string    `bun:"description,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                       // description
	AllocatorFlag         bool      `bun:"allocator_flag,type:bit,default:((0))"`                                                                                                                                                                                           // allocator_flag
	HigherBetterFlag      bool      `bun:"higher_better_flag,type:bit,default:((0))"`                                                                                                                                                                                       // higher_better_flag
	RankableFlag          bool      `bun:"rankable_flag,type:bit,default:((0))"`                                                                                                                                                                                            // rankable_flag
	DirectCostFlag        bool      `bun:"direct_cost_flag,type:bit,default:((0))"`                                                                                                                                                                                         // direct_cost_flag
	ScorecardEligibleFlag bool      `bun:"scorecard_eligible_flag,type:bit,default:((0))"`                                                                                                                                                                                  // scorecard_eligible_flag
	MetricSql             string    `bun:"metric_sql,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                        // metric_sql
	MetricColumn          string    `bun:"metric_column,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                     // metric_column
	MetricTable           string    `bun:"metric_table,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                      // metric_table
	MetricColumnDatatype  string    `bun:"metric_column_datatype,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`            // metric_column_datatype
	MetricType            string    `bun:"metric_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                       // metric_type
	EnableCalculation     bool      `bun:"enable_calculation,type:bit,default:((0))"`                                                                                                                                                                                       // enable_calculation
	FormatMask            string    `bun:"format_mask,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                       // format_mask
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                                                  // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                                            // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:(suser_sname())"` // User who last changed the record
	ExecSeq               int32     `bun:"exec_seq,type:int,nullzero"`                                                                                                                                                                                                      // sequence number
}
