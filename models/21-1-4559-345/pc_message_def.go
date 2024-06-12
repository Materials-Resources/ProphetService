package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcMessageDef struct {
	bun.BaseModel     `bun:"table:pc_message_def"`
	MessageSkey       int32     `bun:"message_skey,type:int,pk"`
	MessageGroup      string    `bun:"message_group,type:char(8)"`
	MessageId         string    `bun:"message_id,type:varchar(20)"`
	SeverityLevel     int32     `bun:"severity_level,type:int,nullzero"`
	LoggingInd        int32     `bun:"logging_ind,type:int,nullzero"`
	DialogClassname   string    `bun:"dialog_classname,type:varchar(20),nullzero"`
	ResponseClassname string    `bun:"response_classname,type:varchar(40),nullzero"`
	ResponseDefault   int32     `bun:"response_default,type:int,nullzero"`
	Description       string    `bun:"description,type:varchar(100),nullzero"`
	CreateUserId      string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate        time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId       string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate         time.Time `bun:"maint_date,type:datetime,nullzero"`
}
