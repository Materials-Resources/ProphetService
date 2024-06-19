package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FinReport struct {
	bun.BaseModel       `bun:"table:fin_report"`
	FinReportId         string    `bun:"fin_report_id,type:varchar(15),pk"`                         // ID of a financial report.
	FinReportDesc       string    `bun:"fin_report_desc,type:varchar(40)"`                          // Description of a financial report ID.
	Application         string    `bun:"application,type:varchar(15),nullzero"`                     // The spreadsheet application used to create a particular report.
	Worksheet           string    `bun:"worksheet,type:varchar(32),nullzero"`                       // The name of the worksheet for the finanacial report.
	DateCreated         time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	PathAndFileName     string    `bun:"path_and_file_name,type:varchar(255),nullzero"`             // The path and file name of the financial worksheet.
	DeleteFlag          string    `bun:"delete_flag,type:char(1),default:('N')"`                    // Indicates whether the financial statement is deleted.
	RecordTypeCd        int32     `bun:"record_type_cd,type:int,nullzero"`                          // Determine the type of record (User Defined, System Defined, Not In Use, etc...)
	FinReportUid        int32     `bun:"fin_report_uid,type:int,autoincrement,unique,scanonly"`     // Unique identity column for fin_report
	SourceTypeCd        int32     `bun:"source_type_cd,type:int,nullzero"`                          // Source type code for fin_report.
	StatementTypeCd     int32     `bun:"statement_type_cd,type:int,default:((2137))"`               // The statement type code for this record.
	InitialNumberOfRows int32     `bun:"initial_number_of_rows,type:int,nullzero"`                  // This value determines how many initial rows are added to express setup if no previous rows exist.
	ReportTypeCd        int32     `bun:"report_type_cd,type:int,nullzero"`                          // The report type of this record.
}
