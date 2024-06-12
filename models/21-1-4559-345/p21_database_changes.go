package model

import "github.com/uptrace/bun"

type P21DatabaseChanges struct {
	bun.BaseModel         `bun:"table:p21_database_changes"`
	P21DatabaseChangesUid int32  `bun:"p21_database_changes_uid,type:int,identity"`
	Change                string `bun:"change,type:varchar(50),nullzero"`
	ObjectType            string `bun:"object_type,type:varchar(255),nullzero"`
	ObjectName            string `bun:"object_name,type:varchar(255),nullzero"`
	TableName             string `bun:"table_name,type:varchar(255),nullzero"`
	Version               string `bun:"version,type:varchar(255),nullzero"`
}
