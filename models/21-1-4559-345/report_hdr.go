package model

type ReportHdr struct {
	bun.BaseModel        `bun:"table:report_hdr"`
	ReportHdrUid         int32     `bun:"report_hdr_uid,type:int,pk,identity"`
	ReportName           string    `bun:"report_name,type:varchar(255)"`
	ReportSourcePath     string    `bun:"report_source_path,type:varchar(255)"`
	PostPurgeFlag        string    `bun:"post_purge_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	NumberOfReportCopies int32     `bun:"number_of_report_copies,type:int,default:((1))"`
	ReportCollate        int32     `bun:"report_collate,type:int,default:((1))"`
}
