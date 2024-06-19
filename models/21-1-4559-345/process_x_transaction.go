package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProcessXTransaction struct {
	bun.BaseModel                  `bun:"table:process_x_transaction"`
	ProcessXTransactionUid         int32     `bun:"process_x_transaction_uid,type:int,pk"`                         // FK to process_x_transaction table
	TransactionType                string    `bun:"transaction_type,type:varchar(5),unique"`                       // What kind of transaction started the process?
	TransactionNo                  int32     `bun:"transaction_no,type:int,unique"`                                // The document number of the transaction that started the process (i.e. Order Number -  PO Number)
	TransactionLineNo              int16     `bun:"transaction_line_no,type:smallint,unique"`                      // What was the line number of the transaction that started the process?
	ProcessUid                     int32     `bun:"process_uid,type:int,nullzero"`                                 // What process belongs to this relationship to a transaction?
	RawInvMastUid                  int32     `bun:"raw_inv_mast_uid,type:int,unique"`                              // The FK to the inv_mast record containing the item being processed.
	FinishedInvMastUid             int32     `bun:"finished_inv_mast_uid,type:int,unique"`                         // What item is being producted by this relatinship between a process and a transaction?
	BeginDate                      time.Time `bun:"begin_date,type:datetime"`                                      // The date to begin posting this voucher.
	EstimatedDate                  time.Time `bun:"estimated_date,type:datetime"`                                  // What was the original estimate for a completion date for this relationship between a process and a transaction?
	ExpectedDate                   time.Time `bun:"expected_date,type:datetime"`                                   // What is the expected date that this relatinship between a process and a transaction should be completed?
	RawQtyRequested                float64   `bun:"raw_qty_requested,type:decimal(19,9)"`                          // Quantity of raw material needed to start the process. Stored in eaches.
	RawQtyAllocated                float64   `bun:"raw_qty_allocated,type:decimal(19,9)"`                          // Quantity of raw material allocated to the process. Decremeted as material is completely processed.
	LocationId                     float64   `bun:"location_id,type:decimal(19,0)"`                                // What location started the process?
	DateCreated                    time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	RowStatusFlag                  int16     `bun:"row_status_flag,type:smallint"`                                 // Indicates current record status.
	AccumulatedCostOfProcessing    float64   `bun:"accumulated_cost_of_processing,type:decimal(19,4)"`             // Accumulated cost of process.
	ParentProcessXTransUid         int32     `bun:"parent_process_x_trans_uid,type:int,unique"`                    // The unique identifier for the parent process.
	Disposition                    string    `bun:"disposition,type:char(1)"`                                      // The disposition of the unallocated raw material.
	QtyCompleted                   float64   `bun:"qty_completed,type:decimal(19,9)"`                              // Quantity for which processing is complete.
	ExtendedDescription            string    `bun:"extended_description,type:text,nullzero"`                       // Extended description of the process.
	ProcessCd                      string    `bun:"process_cd,type:varchar(255),nullzero"`                         // Code used to identify this process.
	ProcessName                    string    `bun:"process_name,type:varchar(255),nullzero"`                       // Name of the process.
	ProcessDesc                    string    `bun:"process_desc,type:varchar(255),nullzero"`                       // Extended description for this process.
	RawYieldQty                    float64   `bun:"raw_yield_qty,type:decimal(19,9),nullzero"`                     // The number of raw items with which you are starting a process.
	FinishedYieldQty               float64   `bun:"finished_yield_qty,type:decimal(19,9),nullzero"`                // The number of finished items you have after a process has been completed.
	RawYieldUom                    string    `bun:"raw_yield_uom,type:varchar(8),nullzero"`                        // The unit of measure for the raw item.
	FinishedYieldUom               string    `bun:"finished_yield_uom,type:varchar(8),nullzero"`                   // The unit of measure for the finished item.
	FinishedQtyCompleted           float64   `bun:"finished_qty_completed,type:decimal(19,9)"`                     // Created qty of finished item
	ReportPrintedFlag              string    `bun:"report_printed_flag,type:char(1)"`                              // Whether the traveler report has been printed
	CompletedDate                  time.Time `bun:"completed_date,type:datetime,nullzero"`                         // Completion date for process trasaction
	RetrievedByWms                 string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`                   // Determines whether or not this record has been retrieved by pathguide
	CaptureUsageFlag               string    `bun:"capture_usage_flag,type:char(1),default:('N')"`                 // Column indicates if usage will be captured for raw items. If set to Y usage will be captured. If set to N usage will not be captured.
	ConfirmableRowStatusFlag       int32     `bun:"confirmable_row_status_flag,type:int,default:((1980))"`         // Indicates current picking status.
	CountryOfOrigin                string    `bun:"country_of_origin,type:varchar(255),nullzero"`                  // Custom column to store finished item's country of origin in Secondary Processing Window
	ProcessPoWithoutAllocationFlag string    `bun:"process_po_without_allocation_flag,type:char(1),nullzero"`      // Allows the user to process first po without allocating to the transaction
	QwoNo                          string    `bun:"qwo_no,type:varchar(255),nullzero"`                             // QWO number. Used to track QA/QC receipt processes in both P21 and WA Solution
}
