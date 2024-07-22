package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Process struct {
	bun.BaseModel           `bun:"table:process"`
	ProcessUid              int32     `bun:"process_uid,type:int,pk"`                                       // What process does this predefined item refer to?
	ProcessCode             string    `bun:"process_code,type:varchar(255)"`                                // What is the short  code or abbreviation used to identify  this process?
	ProcessName             string    `bun:"process_name,type:varchar(255)"`                                // What is the name of the process?
	ProcessDescription      *string   `bun:"process_description,type:varchar(255)"`                         // What is the extended description for this process?
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	RowStatusFlag           int16     `bun:"row_status_flag,type:smallint"`                                 // Indicates current record status.
	CompanyId               *string   `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	LocationId              *float64  `bun:"location_id,type:decimal(19,0)"`                                // Location for the process
	AllLocationsFlag        string    `bun:"all_locations_flag,type:char(1),default:('Y')"`                 // Indicates whether the process  is for a single location or all locations in the company
	ReceiptProcessFlag      string    `bun:"receipt_process_flag,type:char(1),default:('N')"`               // Indicates whether the process is a generic QA/QC process that can be auto-started for any item at material receipts.
	RawInvMastUid           *int32    `bun:"raw_inv_mast_uid,type:int"`                                     // Unique identifier for the raw item used in this process
	FinishedInvMastUid      *int32    `bun:"finished_inv_mast_uid,type:int"`                                // Unique identifier for the finished item created through this process
	RawYieldQty             *float64  `bun:"raw_yield_qty,type:decimal(19,9)"`                              // The number of raw items with which you are starting a process.
	FinishedYieldQty        *float64  `bun:"finished_yield_qty,type:decimal(19,9)"`                         // The number of finished items you have after a process has been completed.
	RawYieldUom             *string   `bun:"raw_yield_uom,type:varchar(8)"`                                 // The unit of measure for the raw item.
	FinishedYieldUom        *string   `bun:"finished_yield_uom,type:varchar(8)"`                            // The unit of measure for the finished item.
	CaptureUsageFlag        string    `bun:"capture_usage_flag,type:char(1),default:('N')"`                 // Column indicates if usage will be captured for raw items. If set to Y usage will be captured. If set to N usage will not be captured.
	RawItemRevisionUid      *int32    `bun:"raw_item_revision_uid,type:int"`                                // Column holds item revision uid of raw item
	FinishedItemRevisionUid *int32    `bun:"finished_item_revision_uid,type:int"`                           // Column holds item revision uid of finished item
	ProductionOrderFlag     *string   `bun:"production_order_flag,type:char(1)"`                            // Pre-Defined Routing specific of Production Order process (On the Fly created in Assembly Maintenance)
	OneOffFlag              *string   `bun:"one_off_flag,type:char(1)"`                                     // Pre-Defined Routing specific of Production Order process (On the Fly created in Assembly Maintenance) that will only be used for the assembly associated with
}
