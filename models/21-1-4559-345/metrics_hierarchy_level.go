package model

import "github.com/uptrace/bun"

type MetricsHierarchyLevel struct {
	bun.BaseModel            `bun:"table:metrics_hierarchy_level"`
	MetricsHierarchyLevelUid int32  `bun:"metrics_hierarchy_level_uid,type:int,pk,identity"`
	Name                     string `bun:"name,type:nvarchar(255)"`
	Active                   bool   `bun:"active,type:bit"`
	ParentHierarchyLevelUid  int32  `bun:"parent_hierarchy_level_uid,type:int,nullzero"`
	ChildRelationship        string `bun:"child_relationship,type:nvarchar,nullzero"`
	KeyField                 string `bun:"key_field,type:nvarchar(255),nullzero"`
	TableName                string `bun:"table_name,type:nvarchar(255),nullzero"`
	ScorecardEligible        bool   `bun:"scorecard_eligible,type:bit,nullzero"`
}
