package model

type ReportCodeP21 struct {
	bun.BaseModel    `bun:"table:report_code_p21"`
	ReportCodeP21Uid int32     `bun:"report_code_p21_uid,type:int,pk,identity"`
	CodeNo           int32     `bun:"code_no,type:int"`
	LanguageId       string    `bun:"language_id,type:varchar(8)"`
	CodeDescription  string    `bun:"code_description,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
