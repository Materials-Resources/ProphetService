package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PcSeqctl struct {
	bun.BaseModel   `bun:"table:pc_seqctl"`
	SeqId           string     `bun:"seq_id,type:varchar(18),pk"`       // Padlock table
	CurrentSkey     *int32     `bun:"current_skey,type:int"`            // Padlock table
	SeqctlComment   *string    `bun:"seqctl_comment,type:varchar(100)"` // Padlock table
	TableName       *string    `bun:"table_name,type:varchar(18)"`      // Padlock table
	ColumnName      *string    `bun:"column_name,type:varchar(18)"`     // Padlock table
	DistributedFlag *string    `bun:"distributed_flag,type:varchar(1)"` // Padlock table
	CreateUserId    *string    `bun:"create_user_id,type:char(10)"`     // Padlock table
	CreateDate      *time.Time `bun:"create_date,type:datetime"`        // Padlock table
	MaintUserId     *string    `bun:"maint_user_id,type:char(10)"`      // Padlock table
	MaintDate       *time.Time `bun:"maint_date,type:datetime"`         // Padlock table
}
