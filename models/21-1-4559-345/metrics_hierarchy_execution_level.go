package gen

import "github.com/uptrace/bun"

type MetricsHierarchyExecutionLevel struct {
	bun.BaseModel                     `bun:"table:metrics_hierarchy_execution_level"`
	MetricsHierarchyExecutionLevelUid int32 `bun:"metrics_hierarchy_execution_level_uid,type:int,autoincrement,pk"` // metrics_hierarchy_execution_level_uid
	MetricsUid                        int32 `bun:"metrics_uid,type:int"`                                            // metrics_uid
	MetricsHierarchyLevelUid          int32 `bun:"metrics_hierarchy_level_uid,type:int"`                            // metrics_hierarchy_level_uid
	Active                            bool  `bun:"active,type:bit"`                                                 // active
}
