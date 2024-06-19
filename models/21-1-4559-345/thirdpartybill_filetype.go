package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ThirdpartybillFiletype struct {
	bun.BaseModel             `bun:"table:thirdpartybill_filetype"`
	ThirdpartybillFiletypeUid int32     `bun:"thirdpartybill_filetype_uid,type:int,autoincrement,scanonly,pk"` // Unique record identifier
	FiletypePrefix            string    `bun:"filetype_prefix,type:varchar(255)"`                              // P21 output filename prefix
	Description               string    `bun:"description,type:varchar(255)"`                                  // P21 output file description
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`                                       // Indicates current record status
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
