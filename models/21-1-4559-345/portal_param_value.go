package model

type PortalParamValue struct {
	bun.BaseModel       `bun:"table:portal_param_value"`
	PortalParamValueUid int32     `bun:"portal_param_value_uid,type:int,pk,identity"`
	PortalAssignmentUid int32     `bun:"portal_assignment_uid,type:int"`
	PortalParamDefUid   int32     `bun:"portal_param_def_uid,type:int"`
	Value               string    `bun:"value,type:varchar(255)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
