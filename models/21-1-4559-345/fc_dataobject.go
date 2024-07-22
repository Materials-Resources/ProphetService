package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FcDataobject struct {
	bun.BaseModel     `bun:"table:fc_dataobject"`
	FcDataobjectUid   int32     `bun:"fc_dataobject_uid,type:int,pk"`                                // Unique Identifier for fc_dataobject record.
	Dataobject        string    `bun:"dataobject,type:varchar(255)"`                                 // The name of a dataobject that has some type of field chooser edit associated with it.
	PrimaryDataobject *string   `bun:"primary_dataobject,type:varchar(255)"`                         // The dataobject for the primary datastore that shares data with the edited dataobject.
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
