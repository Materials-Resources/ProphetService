package model

import "github.com/uptrace/bun"

type P21DeletionSql struct {
	bun.BaseModel     `bun:"table:p21_deletion_sql"`
	P21DeletionSqlUid int32  `bun:"p21_deletion_sql_uid,type:int,identity"`
	DeletionType      int32  `bun:"deletion_type,type:int"`
	DeletionOrder     int32  `bun:"deletion_order,type:int"`
	DeletionDb        string `bun:"deletion_db,type:char"`
	DeletionAction    string `bun:"deletion_action,type:char"`
	DeletionSql       string `bun:"deletion_sql,type:nvarchar(4000)"`
	DeletionTable     string `bun:"deletion_table,type:varchar(256)"`
}
