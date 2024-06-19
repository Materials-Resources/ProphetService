package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Datasource struct {
	bun.BaseModel     `bun:"table:datasource"`
	DatasourceUid     int32     `bun:"datasource_uid,type:int,autoincrement,pk"`                     // Unique identifier for the table.
	DatasourceDesc    string    `bun:"datasource_desc,type:varchar(255)"`                            // Description of the data source
	DatasourceObject  string    `bun:"datasource_object,type:varchar(255)"`                          // Object that SQLcome from
	ObjectTypeCd      int32     `bun:"object_type_cd,type:int"`                                      // Indicates where the object comes from (ie Native, XML).
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	EnableForAllUsers string    `bun:"enable_for_all_users,type:char(1),default:('Y')"`              // When set to Y, all users can create reports from this data source
}
