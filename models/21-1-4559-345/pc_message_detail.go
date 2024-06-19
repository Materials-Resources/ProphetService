package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PcMessageDetail struct {
	bun.BaseModel `bun:"table:pc_message_detail"`
	MessageSkey   int32     `bun:"message_skey,type:int,pk"`                 // Padlock table
	LanguageSkey  int32     `bun:"language_skey,type:int,pk"`                // Padlock table
	MessageTitle  string    `bun:"message_title,type:varchar(255),nullzero"` // Padlock table
	MessageText   string    `bun:"message_text,type:varchar(255),nullzero"`  // Padlock table
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`    // Padlock table
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`       // Padlock table
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`     // Padlock table
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`        // Padlock table
}
