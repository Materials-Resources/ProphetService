package gen

import "github.com/uptrace/bun"

type MetricsPeriodDefinition struct {
	bun.BaseModel              `bun:"table:metrics_period_definition"`
	MetricsPeriodDefinitionUid int32  `bun:"metrics_period_definition_uid,type:int,autoincrement,pk"`                                                                                                                                          // metrics_period_definition_uid
	PeriodDefinition           string `bun:"period_definition,type:char(4)"`                                                                                                                                                                   // period_definition
	Active                     bool   `bun:"active,type:bit"`                                                                                                                                                                                  // active
	Description                string `bun:"description,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // description
}
