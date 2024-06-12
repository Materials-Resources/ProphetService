package model

type SystemAlerts struct {
	bun.BaseModel     `bun:"table:system_alerts"`
	SystemAlertsUid   int32     `bun:"system_alerts_uid,type:int,pk"`
	UserId            string    `bun:"user_id,type:varchar(30)"`
	SystemAlertTypeCd int32     `bun:"system_alert_type_cd,type:int"`
	Subject           string    `bun:"subject,type:varchar(255)"`
	Notes             string    `bun:"notes,type:varchar(255)"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`
}
