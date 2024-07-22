package prophet

import "github.com/uptrace/bun"

type MetricsFilterSql struct {
	bun.BaseModel            `bun:"table:metrics_filter_sql"`
	MetricsFilterSqlUid      int32  `bun:"metrics_filter_sql_uid,type:int,autoincrement,identity,pk"`                                                                                                                                              // metrics_filter_sql_uid
	MetricsUid               int32  `bun:"metrics_uid,type:int"`                                                                                                                                                                                   // metrics_uid
	MetricsHierarchyLevelUid int32  `bun:"metrics_hierarchy_level_uid,type:int"`                                                                                                                                                                   // metrics_hierarchy_level_uid
	MetricFilterSql          string `bun:"metric_filter_sql,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // metric_filter_sql
}
