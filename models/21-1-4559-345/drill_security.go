package model

type DrillSecurity struct {
	bun.BaseModel    `bun:"table:drill_security"`
	UsersId          string    `bun:"users_id,type:varchar(30)"`
	MenuName         string    `bun:"menu_name,type:varchar(500)"`
	WindowName       string    `bun:"window_name,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
