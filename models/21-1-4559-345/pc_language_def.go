package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcLanguageDef struct {
	bun.BaseModel   `bun:"table:pc_language_def"`
	LanguageSkey    int32     `bun:"language_skey,type:int,pk"`
	LanguageId      string    `bun:"language_id,type:varchar(20)"`
	LanguageComment string    `bun:"language_comment,type:varchar(100),nullzero"`
	CreateUserId    string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate      time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId     string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate       time.Time `bun:"maint_date,type:datetime,nullzero"`
}
