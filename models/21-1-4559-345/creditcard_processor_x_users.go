package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardProcessorXUsers struct {
	bun.BaseModel                `bun:"table:creditcard_processor_x_users"`
	CreditcardProcessorXUsersUid int32     `bun:"creditcard_processor_x_users_uid,type:int,pk"`
	UserId                       string    `bun:"user_id,type:varchar(30)"`
	ProcessorUid                 int32     `bun:"processor_uid,type:int"`
	DateCreated                  time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30)"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`
	RowStatusFlag                int32     `bun:"row_status_flag,type:int,default:(704)"`
}
