package model

type ScheduledJobNotifications struct {
	bun.BaseModel    `bun:"table:scheduled_job_notifications"`
	NotificationUid  int32 `bun:"notification_uid,type:int,pk,identity"`
	NotificationType `bun:"notification_type,type:nvarchar(255)"`
	NotificationDesc `bun:"notification_desc,type:nvarchar(255)"`
	NotificationName `bun:"notification_name,type:nvarchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
}
