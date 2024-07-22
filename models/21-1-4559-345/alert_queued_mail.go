package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertQueuedMail struct {
	bun.BaseModel      `bun:"table:alert_queued_mail"`
	AlertQueuedMailUid int32     `bun:"alert_queued_mail_uid,type:int,pk"`      // Unique Identifier of record
	Subject            *string   `bun:"subject,type:varchar(8000)"`             // Email subject
	EmailTo            *string   `bun:"email_to,type:varchar(8000)"`            // Identifies who is the intended recipient of the email message.
	EmailBody          string    `bun:"email_body,type:text"`                   // The body of the email message.
	ReasonCd           int32     `bun:"reason_cd,type:int"`                     // Code that identifies the reason why this email message got queued. Possible codes from new code group "Queued Email Reasons".
	DateCreated        time.Time `bun:"date_created,type:datetime"`             // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`       // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`    // ID of the user who last maintained this record
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`               // Indicates current record status.
	MailitemId         *int32    `bun:"mailitem_id,type:int"`                   // Identifies the dbmail item
	SenderName         *string   `bun:"sender_name,type:varchar(255)"`          // Name of the alert's sender
	SenderEmailAddress *string   `bun:"sender_email_address,type:varchar(255)"` // Email address of the alert's sender
}
