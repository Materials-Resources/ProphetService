package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EmailLog struct {
	bun.BaseModel          `bun:"table:email_log"`
	EmailLogUid            int32     `bun:"email_log_uid,type:int,autoincrement,scanonly,pk"`             // Email Log Unique Identifier
	CompanyId              string    `bun:"company_id,type:varchar(8),nullzero"`                          // Company ID
	TransactionType        string    `bun:"transaction_type,type:varchar(255),nullzero"`                  // Transaction Type
	TransactionNumber      string    `bun:"transaction_number,type:varchar(255),nullzero"`                // Transaction Number or ID
	TransactionRecipientId float64   `bun:"transaction_recipient_id,type:decimal(19,0),nullzero"`         // Recipent ID (Customer ID or Vendor ID)
	Subject                string    `bun:"subject,type:varchar(8000),nullzero"`                          // Emal Subject
	EmailTo                string    `bun:"email_to,type:varchar(8000),nullzero"`                         // Email To
	EmailBody              string    `bun:"email_body,type:text,nullzero"`                                // Email Body
	SenderName             string    `bun:"sender_name,type:varchar(255)"`                                // Email Sender Name
	SenderEmailAddress     string    `bun:"sender_email_address,type:varchar(255)"`                       // Email Sender Address
	EmailCc                string    `bun:"email_cc,type:varchar(8000),nullzero"`                         // Email CC
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	EmailStatusCd          int32     `bun:"email_status_cd,type:int,nullzero"`                            // Emal Status Code
	ErrorText              string    `bun:"error_text,type:varchar(8000),nullzero"`                       // Error Text
	AttachmentFileNames    string    `bun:"attachment_file_names,type:varchar(8000),nullzero"`            // Listing the file name(s) of any attachments.
	EmailBcc               string    `bun:"email_bcc,type:varchar(8000),nullzero"`                        // Bcc address that this email sends to
}
