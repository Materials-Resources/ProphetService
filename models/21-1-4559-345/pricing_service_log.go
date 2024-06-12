package model

type PricingServiceLog struct {
	bun.BaseModel     `bun:"table:pricing_service_log"`
	LayoutId          float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LogId             float64   `bun:"log_id,type:decimal(19,0),pk"`
	StartDatetime     time.Time `bun:"start_datetime,type:datetime,nullzero"`
	EndDatetime       time.Time `bun:"end_datetime,type:datetime,nullzero"`
	FileRecordCount   int32     `bun:"file_record_count,type:int,nullzero"`
	ErrorReport       string    `bun:"error_report,type:varchar(255),nullzero"`
	DetailReport      string    `bun:"detail_report,type:varchar(255),nullzero"`
	SummaryReport     string    `bun:"summary_report,type:varchar(255),nullzero"`
	ErrorRecords      string    `bun:"error_records,type:varchar(255),nullzero"`
	DeleteFlag        string    `bun:"delete_flag,type:char"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	UpdateReport      string    `bun:"update_report,type:varchar(255),nullzero"`
	PerformanceReport string    `bun:"performance_report,type:varchar(255),nullzero"`
}
