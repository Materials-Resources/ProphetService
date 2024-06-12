package model

type MetricsPeriodDefinition struct {
	bun.BaseModel              `bun:"table:metrics_period_definition"`
	MetricsPeriodDefinitionUid int32  `bun:"metrics_period_definition_uid,type:int,pk,identity"`
	PeriodDefinition           string `bun:"period_definition,type:char(4)"`
	Active                     `bun:"active,type:bit"`
	Description                `bun:"description,type:nvarchar(255)"`
}
