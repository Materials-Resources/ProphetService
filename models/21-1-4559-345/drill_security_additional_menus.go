package model

type DrillSecurityAdditionalMenus struct {
	bun.BaseModel        `bun:"table:drill_security_additional_menus"`
	DsAdditionalMenusUid int32     `bun:"ds_additional_menus_uid,type:int,identity"`
	BaseMenu             string    `bun:"base_menu,type:varchar(255)"`
	DuplicateMenu        string    `bun:"duplicate_menu,type:varchar(255)"`
	WindowName           string    `bun:"window_name,type:varchar(255)"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
