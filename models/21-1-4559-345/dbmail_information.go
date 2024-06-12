package model

type DbmailInformation struct {
	bun.BaseModel            `bun:"table:dbmail_information"`
	DbmailInformationUid     int32     `bun:"dbmail_information_uid,type:int,pk,identity"`
	MailitemId               int32     `bun:"mailitem_id,type:int"`
	UseSystemAlert           string    `bun:"use_system_alert,type:char,nullzero"`
	CallFromWithinMailBo     string    `bun:"call_from_within_mail_bo,type:char,nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AppliedToAlertQueuedMail string    `bun:"applied_to_alert_queued_mail,type:char,default:('N')"`
}
