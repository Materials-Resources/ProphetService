package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardProcessorXUsers struct {
	bun.BaseModel                `bun:"table:creditcard_processor_x_users"`
	CreditcardProcessorXUsersUid int32     `bun:"creditcard_processor_x_users_uid,type:int,pk"` // UID of this table
	UserId                       string    `bun:"user_id,type:varchar(30)"`                     // ID of a user who was assigned to a credit card processor.  The ID is from user table.
	ProcessorUid                 int32     `bun:"processor_uid,type:int"`                       // A credit card processor center UID, which is from creditcard_processor
	DateCreated                  time.Time `bun:"date_created,type:datetime"`                   // Indicates the date/time this record was created.
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30)"`          // ID of the user who last maintained this record
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`             // Indicates the date/time this record was last modified.
	RowStatusFlag                int32     `bun:"row_status_flag,type:int,default:(704)"`       // Indicates the status of each row in the table.
}
