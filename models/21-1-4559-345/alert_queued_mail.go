package model

type AlertQueuedMail struct {
	bun.BaseModel      `bun:"table:alert_queued_mail"`
	AlertQueuedMailUid int32     `bun:"alert_queued_mail_uid,type:int,pk"`
	Subject            string    `bun:"subject,type:varchar(8000),nullzero"`
	EmailTo            string    `bun:"email_to,type:varchar(8000),nullzero"`
	EmailBody          string    `bun:"email_body,type:text(2147483647)"`
	ReasonCd           int32     `bun:"reason_cd,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`
	MailitemId         int32     `bun:"mailitem_id,type:int,nullzero"`
	SenderName         string    `bun:"sender_name,type:varchar(255),nullzero"`
	SenderEmailAddress string    `bun:"sender_email_address,type:varchar(255),nullzero"`
}
