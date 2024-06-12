package model

type FreightquotePkgType struct {
	bun.BaseModel           `bun:"table:freightquote_pkg_type"`
	FreightquotePkgTypeUid  int32     `bun:"freightquote_pkg_type_uid,type:int,pk,identity"`
	FreightquotePkgTypeId   string    `bun:"freightquote_pkg_type_id,type:varchar(40)"`
	FreightquotePkgTypeDesc string    `bun:"freightquote_pkg_type_desc,type:varchar(255),nullzero"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
