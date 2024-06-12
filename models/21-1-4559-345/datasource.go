package model

type Datasource struct {
	bun.BaseModel     `bun:"table:datasource"`
	DatasourceUid     int32     `bun:"datasource_uid,type:int,pk,identity"`
	DatasourceDesc    string    `bun:"datasource_desc,type:varchar(255)"`
	DatasourceObject  string    `bun:"datasource_object,type:varchar(255)"`
	ObjectTypeCd      int32     `bun:"object_type_cd,type:int"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	EnableForAllUsers string    `bun:"enable_for_all_users,type:char,default:('Y')"`
}
