package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DbmailInformation struct {
	bun.BaseModel            `bun:"table:dbmail_information"`
	DbmailInformationUid     int32     `bun:"dbmail_information_uid,type:int,autoincrement,pk"`             // Unique identifier
	MailitemId               int32     `bun:"mailitem_id,type:int"`                                         // Identifies the dbmail item
	UseSystemAlert           string    `bun:"use_system_alert,type:char(1),nullzero"`                       // Determines whether or not to use the system alert
	CallFromWithinMailBo     string    `bun:"call_from_within_mail_bo,type:char(1),nullzero"`               // Determines whether or not the call was from within BO
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AppliedToAlertQueuedMail string    `bun:"applied_to_alert_queued_mail,type:char(1),default:('N')"`      // Determines whether or not this row has been applied to the alert_queued_mail table.
}
