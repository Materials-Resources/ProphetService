package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EmailNotificationToken struct {
	bun.BaseModel             `bun:"table:email_notification_token"`
	EmailNotificationTokenUid int32     `bun:"email_notification_token_uid,type:int,autoincrement,identity"` // Unique identifier for the record
	TokenName                 string    `bun:"token_name,type:varchar(255)"`                                 // Name of the token so we can match on it, e.g. <customer_id>
	TokenDescription          string    `bun:"token_description,type:varchar(255)"`                          // Description of token so user can identify it, e.g. Customer ID
	TokenArea                 int32     `bun:"token_area,type:int"`                                          // The area the token should be visible
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DataTypeCd                int32     `bun:"data_type_cd,type:int,default:((1255))"`                       // Specifies what data type the column is
}
