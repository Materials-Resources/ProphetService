package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FormDestinationHierarchy struct {
	bun.BaseModel               `bun:"table:form_destination_hierarchy"`
	FormDestinationHierarchyUid int32     `bun:"form_destination_hierarchy_uid,type:int,pk"`                   // Unique record identifier
	FormDestinationUid          int32     `bun:"form_destination_uid,type:int"`                                // Identifies unique form/output/source combination
	SequenceNo                  int32     `bun:"sequence_no,type:int"`                                         // User defined sort order for sources
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`                                     // Indicates current record status
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
