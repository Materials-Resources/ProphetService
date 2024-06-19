package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseFkerror struct {
	bun.BaseModel            `bun:"table:import_suspense_fkerror"`
	ImportSuspenseFkerrorUid int32     `bun:"import_suspense_fkerror_uid,type:int,pk"`                   // unique ideintifier for each record
	ImportSuspenseHdrUid     int32     `bun:"import_suspense_hdr_uid,type:int"`                          // link to the import_suspense_hdr table
	FkTablename              string    `bun:"fk_tablename,type:varchar(255)"`                            // name of the table in which the foreign key criteria was not found
	FkError                  string    `bun:"fk_error,type:text"`                                        // criteria used for foreign key validation, usually the SQL where clause
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
