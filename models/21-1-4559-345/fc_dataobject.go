package model

type FcDataobject struct {
	bun.BaseModel     `bun:"table:fc_dataobject"`
	FcDataobjectUid   int32     `bun:"fc_dataobject_uid,type:int,pk"`
	Dataobject        string    `bun:"dataobject,type:varchar(255)"`
	PrimaryDataobject string    `bun:"primary_dataobject,type:varchar(255),nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
