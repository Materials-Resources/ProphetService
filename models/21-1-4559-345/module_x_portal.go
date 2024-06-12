package model

type ModuleXPortal struct {
	bun.BaseModel    `bun:"table:module_x_portal"`
	ModuleXPortalUid int32     `bun:"module_x_portal_uid,type:int,pk,identity"`
	ModuleCd         int32     `bun:"module_cd,type:int"`
	PortalCd         int32     `bun:"portal_cd,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
