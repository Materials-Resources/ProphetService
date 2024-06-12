package model

import (
	"github.com/uptrace/bun"
	"time"
)

type EmailLog struct {
	bun.BaseModel          `bun:"table:email_log"`
	EmailLogUid            int32     `bun:"email_log_uid,type:int,pk,identity"`
	CompanyId              string    `bun:"company_id,type:varchar(8),nullzero"`
	TransactionType        string    `bun:"transaction_type,type:varchar(255),nullzero"`
	TransactionNumber      string    `bun:"transaction_number,type:varchar(255),nullzero"`
	TransactionRecipientId float64   `bun:"transaction_recipient_id,type:decimal(19,0),nullzero"`
	Subject                string    `bun:"subject,type:varchar(8000),nullzero"`
	EmailTo                string    `bun:"email_to,type:varchar(8000),nullzero"`
	EmailBody              string    `bun:"email_body,type:text(2147483647),nullzero"`
	SenderName             string    `bun:"sender_name,type:varchar(255)"`
	SenderEmailAddress     string    `bun:"sender_email_address,type:varchar(255)"`
	EmailCc                string    `bun:"email_cc,type:varchar(8000),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	EmailStatusCd          int32     `bun:"email_status_cd,type:int,nullzero"`
	ErrorText              string    `bun:"error_text,type:varchar(8000),nullzero"`
	AttachmentFileNames    string    `bun:"attachment_file_names,type:varchar(8000),nullzero"`
	EmailBcc               string    `bun:"email_bcc,type:varchar(8000),nullzero"`
}
