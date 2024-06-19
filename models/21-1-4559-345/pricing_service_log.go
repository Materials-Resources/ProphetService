package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceLog struct {
	bun.BaseModel     `bun:"table:pricing_service_log"`
	LayoutId          float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	LogId             float64   `bun:"log_id,type:decimal(19,0),pk"`                              // System-generated key to uniquely identify log record
	StartDatetime     time.Time `bun:"start_datetime,type:datetime,nullzero"`                     // when report started -  for this execution of the layout
	EndDatetime       time.Time `bun:"end_datetime,type:datetime,nullzero"`                       // When report ended -  for this execution of the layout
	FileRecordCount   int32     `bun:"file_record_count,type:int,nullzero"`                       // Number of records in the input file to the pricing service import
	ErrorReport       string    `bun:"error_report,type:varchar(255),nullzero"`                   // Location of error report file
	DetailReport      string    `bun:"detail_report,type:varchar(255),nullzero"`                  // Location of detail report file
	SummaryReport     string    `bun:"summary_report,type:varchar(255),nullzero"`                 // Location of summary report file
	ErrorRecords      string    `bun:"error_records,type:varchar(255),nullzero"`                  // Location of error records
	DeleteFlag        string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	UpdateReport      string    `bun:"update_report,type:varchar(255),nullzero"`                  // Location of price/cost update report file
	PerformanceReport string    `bun:"performance_report,type:varchar(255),nullzero"`             // Column identifies the optional performance report file name
}
