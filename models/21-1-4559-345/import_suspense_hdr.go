package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseHdr struct {
	bun.BaseModel        `bun:"table:import_suspense_hdr"`
	ImportSuspenseHdrUid int32     `bun:"import_suspense_hdr_uid,type:int,pk"`                       // Unique identifier for each record
	ImpexpSourceId       string    `bun:"impexp_source_id,type:varchar(50)"`                         // Identifies the import type - EDI Import, User Import, Warehouse Automation, etc.
	TransactionSetId     string    `bun:"transaction_set_id,type:varchar(32)"`                       // Identifies the type of import - Sales Order, Inventory Receipts, etc.
	MasterFilePath       string    `bun:"master_file_path,type:varchar(255)"`                        // The path of the master/header file - for scheduled imports, this will be the polling path
	MasterFileName       string    `bun:"master_file_name,type:varchar(255)"`                        // Name of the master/header file without the path
	RequiredForImport    string    `bun:"required_for_import,type:varchar(255)"`                     // List of tables for which files are required for this import - used within the application for enabling options
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ScheduledImport      string    `bun:"scheduled_import,type:char(1),default:('N')"`               // Y/N indicator of whether the original import file was processed by the scheduled import service manager
	ImpexpSourceDesc     *string   `bun:"impexp_source_desc,type:varchar(255)"`                      // Description of the import type
	TransactionSetDesc   *string   `bun:"transaction_set_desc,type:varchar(255)"`                    // Description of the transaction set
	FileFormatCd         *int32    `bun:"file_format_cd,type:int,default:(1012)"`                    // Identifies the file format of the original import file - tab-delimited, xml document, fixed-format
	TransactionSusPath   *string   `bun:"transaction_sus_path,type:varchar(200)"`                    // Holds the suspended file path
}
