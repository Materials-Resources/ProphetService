package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ImpexpSource struct {
	bun.BaseModel    `bun:"table:impexp_source"`
	ImpexpSourceUid  int32      `bun:"impexp_source_uid,type:int,pk"`                             // Unique Identifier
	ImpexpSourceId   string     `bun:"impexp_source_id,type:varchar(50),unique"`                  // Text identifier of the type of import this record is for. Ex USERIMPORT, EDIIMPORT, B2B. etc.
	ImpexpSourceDesc string     `bun:"impexp_source_desc,type:varchar(50)"`                       // Friendly description for type of import, Ex User Import.
	DateCreated      *time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified *time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy *string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ImpexpType       *string    `bun:"impexp_type,type:char(1)"`                                  // Code to signify whether this record is for import, exporting or both.
}
