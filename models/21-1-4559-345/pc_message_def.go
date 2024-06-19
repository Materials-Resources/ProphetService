package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PcMessageDef struct {
	bun.BaseModel     `bun:"table:pc_message_def"`
	MessageSkey       int32     `bun:"message_skey,type:int,pk"`                     // Padlock table
	MessageGroup      string    `bun:"message_group,type:char(8),unique"`            // Padlock table
	MessageId         string    `bun:"message_id,type:varchar(20),unique"`           // Padlock table
	SeverityLevel     int32     `bun:"severity_level,type:int,nullzero"`             // Padlock table
	LoggingInd        int32     `bun:"logging_ind,type:int,nullzero"`                // Padlock table
	DialogClassname   string    `bun:"dialog_classname,type:varchar(20),nullzero"`   // Padlock table
	ResponseClassname string    `bun:"response_classname,type:varchar(40),nullzero"` // Padlock table
	ResponseDefault   int32     `bun:"response_default,type:int,nullzero"`           // Padlock table
	Description       string    `bun:"description,type:varchar(100),nullzero"`       // Padlock table
	CreateUserId      string    `bun:"create_user_id,type:char(10),nullzero"`        // Padlock table
	CreateDate        time.Time `bun:"create_date,type:datetime,nullzero"`           // Padlock table
	MaintUserId       string    `bun:"maint_user_id,type:char(10),nullzero"`         // Padlock table
	MaintDate         time.Time `bun:"maint_date,type:datetime,nullzero"`            // Padlock table
}
