package model

type MetricsHierarchyExecutionLevel struct {
	bun.BaseModel                     `bun:"table:metrics_hierarchy_execution_level"`
	MetricsHierarchyExecutionLevelUid int32 `bun:"metrics_hierarchy_execution_level_uid,type:int,pk,identity"`
	MetricsUid                        int32 `bun:"metrics_uid,type:int"`
	MetricsHierarchyLevelUid          int32 `bun:"metrics_hierarchy_level_uid,type:int"`
	Active                            `bun:"active,type:bit"`
}
