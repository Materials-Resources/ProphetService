package model

import "github.com/uptrace/bun"

type CopyItemTable struct {
	bun.BaseModel `bun:"table:copy_item_table"`
	SequenceNo    int32  `bun:"sequence_no,type:int,pk"`
	TableName     string `bun:"table_name,type:varchar(40)"`
	TableUid      string `bun:"table_uid,type:varchar(40),nullzero"`
	KeyColumn     string `bun:"key_column,type:varchar(40)"`
	CounterId     string `bun:"counter_id,type:varchar(64),nullzero"`
}
