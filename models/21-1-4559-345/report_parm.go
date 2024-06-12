package model

type ReportParm struct {
	bun.BaseModel    `bun:"table:report_parm"`
	ReportParmUid    int32     `bun:"report_parm_uid,type:int,pk,identity"`
	ReportHdrUid     int32     `bun:"report_hdr_uid,type:int"`
	ReportLevel      int32     `bun:"report_level,type:int"`
	ParmSeqNo        int32     `bun:"parm_seq_no,type:int"`
	ParmName         string    `bun:"parm_name,type:varchar(255)"`
	ParmValue        string    `bun:"parm_value,type:varchar"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ParmReportName   string    `bun:"parm_report_name,type:varchar(255),default:('')"`
}
