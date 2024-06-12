package model

type MetricsHierarchyLevel struct {
	bun.BaseModel            `bun:"table:metrics_hierarchy_level"`
	MetricsHierarchyLevelUid int32 `bun:"metrics_hierarchy_level_uid,type:int,pk,identity"`
	Name                     `bun:"name,type:nvarchar(255)"`
	Active                   `bun:"active,type:bit"`
	ParentHierarchyLevelUid  int32 `bun:"parent_hierarchy_level_uid,type:int,nullzero"`
	ChildRelationship        `bun:"child_relationship,type:nvarchar,nullzero"`
	KeyField                 `bun:"key_field,type:nvarchar(255),nullzero"`
	TableName                `bun:"table_name,type:nvarchar(255),nullzero"`
	ScorecardEligible        `bun:"scorecard_eligible,type:bit,nullzero"`
}
