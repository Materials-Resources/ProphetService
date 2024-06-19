package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledImportMaster struct {
	bun.BaseModel            `bun:"table:scheduled_import_master"`
	ScheduledImportMasterUid int32     `bun:"scheduled_import_master_uid,type:int,pk"`                   // Unique Identifier
	ImpexpSourceUid          int32     `bun:"impexp_source_uid,type:int"`                                // Pointer to the record that describes the type of import. User, EDI, B2B, etc.
	TransactionSetUid        int32     `bun:"transaction_set_uid,type:int"`                              // Pointer to the transaction record.
	PollingPath              string    `bun:"polling_path,type:varchar(200)"`                            // Fully qualified path the application will look for import files.
	TransactionLogPath       string    `bun:"transaction_log_path,type:varchar(200)"`                    // Fully qualified path the application will write the import log information.
	TransactionSumPath       string    `bun:"transaction_sum_path,type:varchar(200)"`                    // Fully qualified path the application will write the import summary information.
	TransactionSusPath       string    `bun:"transaction_sus_path,type:varchar(200)"`                    // Fully qualified path the application will write the suspended import files.
	TransactionErrPath       string    `bun:"transaction_err_path,type:varchar(200)"`                    // Fully qualified path the application will write the import error information.
	Active                   string    `bun:"active,type:char(1)"`                                       // Whether the import type is active or not, tells the system to look for the files or not.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	FileFormatCd             int32     `bun:"file_format_cd,type:int,default:(1012)"`                    // Code identifying the format the data is in, XML, tab-delimited.
	XmlDocumentUid           int32     `bun:"xml_document_uid,type:int,nullzero"`                        // If XML, pointer to the document record defining the data set.
	FileLockingFlag          string    `bun:"file_locking_flag,type:char(1),default:('N')"`              // Determines whether this import type uses file locking
}
