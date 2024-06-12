package model

type ProcessXTransaction struct {
	bun.BaseModel                  `bun:"table:process_x_transaction"`
	ProcessXTransactionUid         int32     `bun:"process_x_transaction_uid,type:int,pk"`
	TransactionType                string    `bun:"transaction_type,type:varchar(5)"`
	TransactionNo                  int32     `bun:"transaction_no,type:int"`
	TransactionLineNo              int16     `bun:"transaction_line_no,type:smallint"`
	ProcessUid                     int32     `bun:"process_uid,type:int,nullzero"`
	RawInvMastUid                  int32     `bun:"raw_inv_mast_uid,type:int"`
	FinishedInvMastUid             int32     `bun:"finished_inv_mast_uid,type:int"`
	BeginDate                      time.Time `bun:"begin_date,type:datetime"`
	EstimatedDate                  time.Time `bun:"estimated_date,type:datetime"`
	ExpectedDate                   time.Time `bun:"expected_date,type:datetime"`
	RawQtyRequested                float64   `bun:"raw_qty_requested,type:decimal(19,9)"`
	RawQtyAllocated                float64   `bun:"raw_qty_allocated,type:decimal(19,9)"`
	LocationId                     float64   `bun:"location_id,type:decimal(19,0)"`
	DateCreated                    time.Time `bun:"date_created,type:datetime"`
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	RowStatusFlag                  int16     `bun:"row_status_flag,type:smallint"`
	AccumulatedCostOfProcessing    float64   `bun:"accumulated_cost_of_processing,type:decimal(19,4)"`
	ParentProcessXTransUid         int32     `bun:"parent_process_x_trans_uid,type:int"`
	Disposition                    string    `bun:"disposition,type:char"`
	QtyCompleted                   float64   `bun:"qty_completed,type:decimal(19,9)"`
	ExtendedDescription            string    `bun:"extended_description,type:text(2147483647),nullzero"`
	ProcessCd                      string    `bun:"process_cd,type:varchar(255),nullzero"`
	ProcessName                    string    `bun:"process_name,type:varchar(255),nullzero"`
	ProcessDesc                    string    `bun:"process_desc,type:varchar(255),nullzero"`
	RawYieldQty                    float64   `bun:"raw_yield_qty,type:decimal(19,9),nullzero"`
	FinishedYieldQty               float64   `bun:"finished_yield_qty,type:decimal(19,9),nullzero"`
	RawYieldUom                    string    `bun:"raw_yield_uom,type:varchar(8),nullzero"`
	FinishedYieldUom               string    `bun:"finished_yield_uom,type:varchar(8),nullzero"`
	FinishedQtyCompleted           float64   `bun:"finished_qty_completed,type:decimal(19,9)"`
	ReportPrintedFlag              string    `bun:"report_printed_flag,type:char"`
	CompletedDate                  time.Time `bun:"completed_date,type:datetime,nullzero"`
	RetrievedByWms                 string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	CaptureUsageFlag               string    `bun:"capture_usage_flag,type:char,default:('N')"`
	ConfirmableRowStatusFlag       int32     `bun:"confirmable_row_status_flag,type:int,default:((1980))"`
	CountryOfOrigin                string    `bun:"country_of_origin,type:varchar(255),nullzero"`
	ProcessPoWithoutAllocationFlag string    `bun:"process_po_without_allocation_flag,type:char,nullzero"`
	QwoNo                          string    `bun:"qwo_no,type:varchar(255),nullzero"`
}
