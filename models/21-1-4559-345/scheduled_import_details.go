package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledImportDetails struct {
	bun.BaseModel             `bun:"table:scheduled_import_details"`
	ScheduledImportDetailsUid int32     `bun:"scheduled_import_details_uid,type:int,pk"`                  // Unique Identifier.
	ScheduledImportDefUid     int32     `bun:"scheduled_import_def_uid,type:int"`                         // Pointer to the record that describes the data subset for the transaction set.
	Fileprefix                string    `bun:"fileprefix,type:varchar(20),nullzero"`                      // File prefix that the scheduled import service will use to identify the files data and transaction set.
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
