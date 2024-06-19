package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FormDestination struct {
	bun.BaseModel      `bun:"table:form_destination"`
	FormDestinationUid int32     `bun:"form_destination_uid,type:int,pk"`                             // Unique record identifier
	FormCd             int32     `bun:"form_cd,type:int"`                                             // Unique code that identifies a particular form
	OutputCd           int32     `bun:"output_cd,type:int"`                                           // Unique code that identifies a particular output type
	SourceCd           int32     `bun:"source_cd,type:int"`                                           // Unique code that identifies a particular address source (email address,fax number,etc)
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`                                     // Indicates current record status
}
