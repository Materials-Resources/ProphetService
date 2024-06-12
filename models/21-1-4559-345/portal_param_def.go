package model

type PortalParamDef struct {
	bun.BaseModel     `bun:"table:portal_param_def"`
	PortalParamDefUid int32     `bun:"portal_param_def_uid,type:int,pk"`
	PortalElementUid  int32     `bun:"portal_element_uid,type:int"`
	SequenceNo        int32     `bun:"sequence_no,type:int"`
	ParameterName     string    `bun:"parameter_name,type:varchar(255)"`
	ParameterDesc     string    `bun:"parameter_desc,type:varchar(255)"`
	DatatypeCd        int32     `bun:"datatype_cd,type:int"`
	DefaultValue      string    `bun:"default_value,type:varchar(255)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
