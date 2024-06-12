package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Metrics struct {
	bun.BaseModel         `bun:"table:metrics"`
	MetricsUid            int32     `bun:"metrics_uid,type:int,pk,identity"`
	UniqueName            string    `bun:"unique_name,type:nvarchar(255),nullzero"`
	Name                  string    `bun:"name,type:nvarchar(255)"`
	Description           string    `bun:"description,type:nvarchar(255),nullzero"`
	AllocatorFlag         bool      `bun:"allocator_flag,type:bit,default:((0))"`
	HigherBetterFlag      bool      `bun:"higher_better_flag,type:bit,default:((0))"`
	RankableFlag          bool      `bun:"rankable_flag,type:bit,default:((0))"`
	DirectCostFlag        bool      `bun:"direct_cost_flag,type:bit,default:((0))"`
	ScorecardEligibleFlag bool      `bun:"scorecard_eligible_flag,type:bit,default:((0))"`
	MetricSql             string    `bun:"metric_sql,type:nvarchar,nullzero"`
	MetricColumn          string    `bun:"metric_column,type:nvarchar(255),nullzero"`
	MetricTable           string    `bun:"metric_table,type:nvarchar(255),nullzero"`
	MetricColumnDatatype  string    `bun:"metric_column_datatype,type:nvarchar(255),nullzero"`
	MetricType            string    `bun:"metric_type,type:nvarchar(255),nullzero"`
	EnableCalculation     bool      `bun:"enable_calculation,type:bit,default:((0))"`
	FormatMask            string    `bun:"format_mask,type:nvarchar(255),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:nvarchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:nvarchar(255),default:(suser_sname())"`
	ExecSeq               int32     `bun:"exec_seq,type:int,nullzero"`
}
