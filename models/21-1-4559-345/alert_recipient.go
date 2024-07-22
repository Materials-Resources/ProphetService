package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertRecipient struct {
	bun.BaseModel     `bun:"table:alert_recipient"`
	AlertRecipientUid int32     `bun:"alert_recipient_uid,type:int,pk"`                           // Unique Identifier of record
	AlertMessageUid   int32     `bun:"alert_message_uid,type:int,unique"`                         // Unique identifier from alert_message table
	AlertEmailName    *string   `bun:"alert_email_name,type:varchar(255)"`                        // Name that identifies the email address for alert.
	AlertEmailAddress string    `bun:"alert_email_address,type:varchar(255),unique"`              // Email address for this alert
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	RecordTypeCd      int32     `bun:"record_type_cd,type:int,unique"`                            // his column identifies what type of record we are dealing with. Ex. a "Reply" email address or a "TO" email address
	RecipientTypeCd   int32     `bun:"recipient_type_cd,type:int"`                                // Code that identifies a recipient as a To,CC, BCC, Etc.
}
