package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertMessage struct {
	bun.BaseModel          `bun:"table:alert_message"`
	AlertMessageUid        int32     `bun:"alert_message_uid,type:int,pk"`                             // Unique Identifier of record
	Subject                string    `bun:"subject,type:varchar(255)"`                                 // Subject of message
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	AlertImplementationUid *int32    `bun:"alert_implementation_uid,type:int,unique"`                  // Unique identifier from alert_implementation record
	Header                 *string   `bun:"header,type:text"`                                          // Header of message
	LineItem               *string   `bun:"line_item,type:text"`                                       // Line item of message
	Footer                 *string   `bun:"footer,type:text"`                                          // Footer of message
	Attachment             *string   `bun:"attachment,type:varchar(255)"`                              // String that holds the UNC path of an attachment
	SenderName             *string   `bun:"sender_name,type:varchar(255)"`                             // Name of the alert's sender
	SenderEmailAddress     *string   `bun:"sender_email_address,type:varchar(255)"`                    // Email address of the alert's sender
}
