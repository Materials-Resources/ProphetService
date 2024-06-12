package model

type DwobjectDependency struct {
	bun.BaseModel         `bun:"table:dwobject_dependency"`
	DwobjectDependencyUid int32     `bun:"dwobject_dependency_uid,type:int,pk,identity"`
	PrimaryObjectName     string    `bun:"primary_object_name,type:varchar(255)"`
	SecondaryObjectName   string    `bun:"secondary_object_name,type:varchar(255)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
