package model

type FinReport struct {
	bun.BaseModel       `bun:"table:fin_report"`
	FinReportId         string    `bun:"fin_report_id,type:varchar(15),pk"`
	FinReportDesc       string    `bun:"fin_report_desc,type:varchar(40)"`
	Application         string    `bun:"application,type:varchar(15),nullzero"`
	Worksheet           string    `bun:"worksheet,type:varchar(32),nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	PathAndFileName     string    `bun:"path_and_file_name,type:varchar(255),nullzero"`
	DeleteFlag          string    `bun:"delete_flag,type:char,default:('N')"`
	RecordTypeCd        int32     `bun:"record_type_cd,type:int,nullzero"`
	FinReportUid        int32     `bun:"fin_report_uid,type:int,identity"`
	SourceTypeCd        int32     `bun:"source_type_cd,type:int,nullzero"`
	StatementTypeCd     int32     `bun:"statement_type_cd,type:int,default:((2137))"`
	InitialNumberOfRows int32     `bun:"initial_number_of_rows,type:int,nullzero"`
	ReportTypeCd        int32     `bun:"report_type_cd,type:int,nullzero"`
}
