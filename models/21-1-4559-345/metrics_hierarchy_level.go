package prophet

import "github.com/uptrace/bun"

type MetricsHierarchyLevel struct {
	bun.BaseModel            `bun:"table:metrics_hierarchy_level"`
	MetricsHierarchyLevelUid int32   `bun:"metrics_hierarchy_level_uid,type:int,autoincrement,identity,pk"`                                                                                                                                          // metrics_hierarchy_level_uid
	Name                     string  `bun:"name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`               // name
	Active                   bool    `bun:"active,type:bit"`                                                                                                                                                                                         // active
	ParentHierarchyLevelUid  *int32  `bun:"parent_hierarchy_level_uid,type:int"`                                                                                                                                                                     // parent_hierarchy_level_uid
	ChildRelationship        *string `bun:"child_relationship,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // child_relationship
	KeyField                 *string `bun:"key_field,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`          // key_field
	TableName                *string `bun:"table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`         // table_name
	ScorecardEligible        *bool   `bun:"scorecard_eligible,type:bit"`                                                                                                                                                                             // scorecard_eligible
}
