package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcMessageDetail struct {
	bun.BaseModel `bun:"table:pc_message_detail"`
	MessageSkey   int32     `bun:"message_skey,type:int,pk"`
	LanguageSkey  int32     `bun:"language_skey,type:int,pk"`
	MessageTitle  string    `bun:"message_title,type:varchar(255),nullzero"`
	MessageText   string    `bun:"message_text,type:varchar(255),nullzero"`
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`
}
