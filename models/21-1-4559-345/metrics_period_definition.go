package model

import "github.com/uptrace/bun"

type MetricsPeriodDefinition struct {
	bun.BaseModel              `bun:"table:metrics_period_definition"`
	MetricsPeriodDefinitionUid int32  `bun:"metrics_period_definition_uid,type:int,pk,identity"`
	PeriodDefinition           string `bun:"period_definition,type:char(4)"`
	Active                     bool   `bun:"active,type:bit"`
	Description                string `bun:"description,type:nvarchar(255)"`
}
