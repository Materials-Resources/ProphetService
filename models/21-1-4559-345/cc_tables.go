package model

import "github.com/uptrace/bun"

type CcTables struct {
	bun.BaseModel    `bun:"table:cc_tables"`
	TableUid         int32  `bun:"table_uid,type:int,pk"`
	TableName        string `bun:"table_name,type:varchar(255)"`
	TableDescription string `bun:"table_description,type:varchar(255),nullzero"`
}
