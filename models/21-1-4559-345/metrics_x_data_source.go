package prophet

import "github.com/uptrace/bun"

type MetricsXDataSource struct {
	bun.BaseModel         `bun:"table:metrics_x_data_source"`
	MetricsXDataSourceUid int32   `bun:"metrics_x_data_source_uid,type:int,autoincrement,identity,pk"`                                                                                                                                          // metrics_x_data_source_uid
	DataSourceName        *string `bun:"data_source_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // data_source_name
	MetricsUid            *int32  `bun:"metrics_uid,type:int"`                                                                                                                                                                                  // metrics_uid
}
