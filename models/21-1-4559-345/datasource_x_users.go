package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DatasourceXUsers struct {
	bun.BaseModel       `bun:"table:datasource_x_users"`
	DatasourceXUsersUid int32     `bun:"datasource_x_users_uid,type:int,autoincrement,identity,pk"`    // Identifier
	DatasourceUid       int32     `bun:"datasource_uid,type:int"`                                      // Datasource Uid
	UsersId             string    `bun:"users_id,type:varchar(30)"`                                    // User Id
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
