package model

type MetricsSettings struct {
	bun.BaseModel    `bun:"table:metrics_settings"`
	SettingUid       int32 `bun:"setting_uid,type:int,pk,identity"`
	MajorKey         `bun:"major_key,type:nvarchar(50)"`
	MinorKey         `bun:"minor_key,type:nvarchar(50)"`
	Value            `bun:"value,type:nvarchar,nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        `bun:"created_by,type:nvarchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy `bun:"last_maintained_by,type:nvarchar(255),default:(suser_sname())"`
}
