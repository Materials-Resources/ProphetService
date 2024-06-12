package model

type Metrics struct {
	bun.BaseModel         `bun:"table:metrics"`
	MetricsUid            int32 `bun:"metrics_uid,type:int,pk,identity"`
	UniqueName            `bun:"unique_name,type:nvarchar(255),nullzero"`
	Name                  `bun:"name,type:nvarchar(255)"`
	Description           `bun:"description,type:nvarchar(255),nullzero"`
	AllocatorFlag         `bun:"allocator_flag,type:bit,default:((0))"`
	HigherBetterFlag      `bun:"higher_better_flag,type:bit,default:((0))"`
	RankableFlag          `bun:"rankable_flag,type:bit,default:((0))"`
	DirectCostFlag        `bun:"direct_cost_flag,type:bit,default:((0))"`
	ScorecardEligibleFlag `bun:"scorecard_eligible_flag,type:bit,default:((0))"`
	MetricSql             `bun:"metric_sql,type:nvarchar,nullzero"`
	MetricColumn          `bun:"metric_column,type:nvarchar(255),nullzero"`
	MetricTable           `bun:"metric_table,type:nvarchar(255),nullzero"`
	MetricColumnDatatype  `bun:"metric_column_datatype,type:nvarchar(255),nullzero"`
	MetricType            `bun:"metric_type,type:nvarchar(255),nullzero"`
	EnableCalculation     `bun:"enable_calculation,type:bit,default:((0))"`
	FormatMask            `bun:"format_mask,type:nvarchar(255),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             `bun:"created_by,type:nvarchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      `bun:"last_maintained_by,type:nvarchar(255),default:(suser_sname())"`
	ExecSeq               int32 `bun:"exec_seq,type:int,nullzero"`
}
