package model

type AuditTrailSupport struct {
	bun.BaseModel         `bun:"table:audit_trail_support"`
	AuditTrailSupportUid  int32     `bun:"audit_trail_support_uid,type:int,pk"`
	SourceAreaCd          int32     `bun:"source_area_cd,type:int"`
	SourceTableName       string    `bun:"source_table_name,type:varchar(255)"`
	CompanyColumnName     string    `bun:"company_column_name,type:varchar(255),nullzero"`
	LocationColumnName    string    `bun:"location_column_name,type:varchar(255),nullzero"`
	CompletedColumnName   string    `bun:"completed_column_name,type:varchar(255),nullzero"`
	LineNoDisplayFlag     string    `bun:"line_no_display_flag,type:char,default:('N')"`
	InvMastUidDisplayFlag string    `bun:"inv_mast_uid_display_flag,type:char,default:('N')"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
