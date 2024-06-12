package model

import "github.com/uptrace/bun"

type MetricsFilterSql struct {
	bun.BaseModel            `bun:"table:metrics_filter_sql"`
	MetricsFilterSqlUid      int32  `bun:"metrics_filter_sql_uid,type:int,pk,identity"`
	MetricsUid               int32  `bun:"metrics_uid,type:int"`
	MetricsHierarchyLevelUid int32  `bun:"metrics_hierarchy_level_uid,type:int"`
	MetricFilterSql          string `bun:"metric_filter_sql,type:nvarchar"`
}
