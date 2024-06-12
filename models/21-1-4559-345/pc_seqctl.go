package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcSeqctl struct {
	bun.BaseModel   `bun:"table:pc_seqctl"`
	SeqId           string    `bun:"seq_id,type:varchar(18),pk"`
	CurrentSkey     int32     `bun:"current_skey,type:int,nullzero"`
	SeqctlComment   string    `bun:"seqctl_comment,type:varchar(100),nullzero"`
	TableName       string    `bun:"table_name,type:varchar(18),nullzero"`
	ColumnName      string    `bun:"column_name,type:varchar(18),nullzero"`
	DistributedFlag string    `bun:"distributed_flag,type:varchar,nullzero"`
	CreateUserId    string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate      time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId     string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate       time.Time `bun:"maint_date,type:datetime,nullzero"`
}
