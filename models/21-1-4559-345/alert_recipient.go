package model

type AlertRecipient struct {
	bun.BaseModel     `bun:"table:alert_recipient"`
	AlertRecipientUid int32     `bun:"alert_recipient_uid,type:int,pk"`
	AlertMessageUid   int32     `bun:"alert_message_uid,type:int"`
	AlertEmailName    string    `bun:"alert_email_name,type:varchar(255),nullzero"`
	AlertEmailAddress string    `bun:"alert_email_address,type:varchar(255)"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	RecordTypeCd      int32     `bun:"record_type_cd,type:int"`
	RecipientTypeCd   int32     `bun:"recipient_type_cd,type:int"`
}
