package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Datastream struct {
	bun.BaseModel    `bun:"table:datastream"`
	DatastreamUid    int32     `bun:"datastream_uid,type:int,autoincrement,pk"`                     // Unique identifier for rows.
	DatastreamName   string    `bun:"datastream_name,type:varchar(255),unique"`                     // The name of the data stream data object.
	DatastreamDesc   string    `bun:"datastream_desc,type:varchar(255),nullzero"`                   // Description explaining how and where the data stream is used.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
