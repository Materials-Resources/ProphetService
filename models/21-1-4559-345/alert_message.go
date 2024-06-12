package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertMessage struct {
	bun.BaseModel          `bun:"table:alert_message"`
	AlertMessageUid        int32     `bun:"alert_message_uid,type:int,pk"`
	Subject                string    `bun:"subject,type:varchar(255)"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AlertImplementationUid int32     `bun:"alert_implementation_uid,type:int,nullzero"`
	Header                 string    `bun:"header,type:text(2147483647),nullzero"`
	LineItem               string    `bun:"line_item,type:text(2147483647),nullzero"`
	Footer                 string    `bun:"footer,type:text(2147483647),nullzero"`
	Attachment             string    `bun:"attachment,type:varchar(255),nullzero"`
	SenderName             string    `bun:"sender_name,type:varchar(255),nullzero"`
	SenderEmailAddress     string    `bun:"sender_email_address,type:varchar(255),nullzero"`
}
