package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseSettings struct {
	bun.BaseModel    `bun:"table:import_suspense_settings"`
	ImpexpSourceId   string    `bun:"impexp_source_id,type:varchar(50),pk"`                      // type of import
	TransactionSetId string    `bun:"transaction_set_id,type:varchar(32),pk"`                    // transaction set of import
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(705)"`                    // Indicates current record status.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
