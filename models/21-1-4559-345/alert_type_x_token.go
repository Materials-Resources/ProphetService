package model

type AlertTypeXToken struct {
	bun.BaseModel      `bun:"table:alert_type_x_token"`
	AlertTypeXTokenUid int32     `bun:"alert_type_x_token_uid,type:int,pk"`
	AlertTypeUid       int32     `bun:"alert_type_uid,type:int"`
	TokenUid           int32     `bun:"token_uid,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
