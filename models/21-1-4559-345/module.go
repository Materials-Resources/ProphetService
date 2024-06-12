package model

type Module struct {
	bun.BaseModel     `bun:"table:module"`
	ModuleId          string    `bun:"module_id,type:varchar(8),pk"`
	ModuleDescription string    `bun:"module_description,type:varchar(40)"`
	FrameName         string    `bun:"frame_name,type:varchar(255),nullzero"`
	ClassName         string    `bun:"class_name,type:varchar(255),nullzero"`
	ModuleGroupNo     int32     `bun:"module_group_no,type:int,nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
