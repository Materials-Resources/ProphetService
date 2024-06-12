package model

import "github.com/uptrace/bun"

type MetricsXDataSource struct {
	bun.BaseModel         `bun:"table:metrics_x_data_source"`
	MetricsXDataSourceUid int32  `bun:"metrics_x_data_source_uid,type:int,pk,identity"`
	DataSourceName        string `bun:"data_source_name,type:nvarchar(255),nullzero"`
	MetricsUid            int32  `bun:"metrics_uid,type:int,nullzero"`
}
