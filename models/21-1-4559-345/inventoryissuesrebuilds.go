package model

type Inventoryissuesrebuilds struct {
	bun.BaseModel    `bun:"table:inventoryissuesrebuilds"`
	Rebuilduid       int32     `bun:"rebuilduid,type:int,pk,identity"`
	RebuildSp        string    `bun:"rebuild_sp,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RebuildEnabled   `bun:"rebuild_enabled,type:bit,nullzero"`
}
