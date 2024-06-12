package model

type PortalElement struct {
	bun.BaseModel     `bun:"table:portal_element"`
	PortalElementUid  int32     `bun:"portal_element_uid,type:int,pk"`
	PortalElementName string    `bun:"portal_element_name,type:varchar(255)"`
	Classname         string    `bun:"classname,type:varchar(255)"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	IconName          string    `bun:"icon_name,type:varchar(255),nullzero"`
	PortalCd          int32     `bun:"portal_cd,type:int,nullzero"`
	ReportMetadataUid int32     `bun:"report_metadata_uid,type:int,nullzero"`
}
