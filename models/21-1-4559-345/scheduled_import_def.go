package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledImportDef struct {
	bun.BaseModel         `bun:"table:scheduled_import_def"`
	ScheduledImportDefUid int32     `bun:"scheduled_import_def_uid,type:int,pk"`                      // Unique Identifier
	ImpexpSourceUid       int32     `bun:"impexp_source_uid,type:int"`                                // Pointer to the type of import record.
	TransactionSetUid     int32     `bun:"transaction_set_uid,type:int"`                              // Pointer to the transaction type record.
	DbtableId             string    `bun:"dbtable_id,type:varchar(40)"`                               // Table that will be updated as part of the import of the transaction type, ex oe_hdr for Sales Order Import.
	DbtableDesc           string    `bun:"dbtable_desc,type:varchar(40)"`                             // Friendly description for table ID, ex Order Header.
	Req                   string    `bun:"req,type:char(1)"`                                          // Whether the data for this table is required for this transaction type.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	FileFormatCd          int32     `bun:"file_format_cd,type:int,default:(1012)"`                    // Code describing the format of the data, ex XML or tab delimited.
	DefaultFileprefix     string    `bun:"default_fileprefix,type:varchar(255),nullzero"`             // Default code assigned to an import file
}
