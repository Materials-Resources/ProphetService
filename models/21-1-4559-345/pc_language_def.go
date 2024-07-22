package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PcLanguageDef struct {
	bun.BaseModel   `bun:"table:pc_language_def"`
	LanguageSkey    int32      `bun:"language_skey,type:int,pk"`           // Padlock table
	LanguageId      string     `bun:"language_id,type:varchar(20),unique"` // Padlock table
	LanguageComment *string    `bun:"language_comment,type:varchar(100)"`  // Padlock table
	CreateUserId    *string    `bun:"create_user_id,type:char(10)"`        // Padlock table
	CreateDate      *time.Time `bun:"create_date,type:datetime"`           // Padlock table
	MaintUserId     *string    `bun:"maint_user_id,type:char(10)"`         // Padlock table
	MaintDate       *time.Time `bun:"maint_date,type:datetime"`            // Padlock table
}
