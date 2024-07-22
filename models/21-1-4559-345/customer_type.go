package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerType struct {
	bun.BaseModel    `bun:"table:customer_type"`
	CustomerTypeUid  int32     `bun:"customer_type_uid,type:int,pk"`                               // System generated unique ID column
	CustomerType     string    `bun:"customer_type,type:varchar(40)"`                              // The various type of customer are assigned different customer types
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`                               // Identifies the current status of the record
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`              // Date when record was created
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`        // Date when record was last modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"` // User who last modified the record
}
