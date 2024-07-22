package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransaction struct {
	bun.BaseModel           `bun:"table:dct_transaction"`
	DctTransactionUid       int32     `bun:"dct_transaction_uid,type:int,pk"`                              // unique identifier for record
	TransactionId           string    `bun:"transaction_id,type:varchar(255)"`                             // import transaction code
	TransactionDesc         string    `bun:"transaction_desc,type:varchar(255)"`                           // import transaction description
	BaselineTransactionFlag string    `bun:"baseline_transaction_flag,type:char(1),default:('N')"`         // indicates transaction is baseline or custom
	ProcessTypeCd           int32     `bun:"process_type_cd,type:int,default:(1657)"`                      // processing done within business objects or window
	ImportTypeCd            int32     `bun:"import_type_cd,type:int,default:(1649)"`                       // import or update import
	RowStatusFlag           int32     `bun:"row_status_flag,type:int,default:(704)"`                       // record status
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	MenuclickedName         *string   `bun:"menuclicked_name,type:varchar(255)"`                           // Identifies the menu item the user clicks to open associated import window
	LegacyExportFlag        string    `bun:"legacy_export_flag,type:char(1),default:('N')"`                // Indicates whether any legacy solutin exports data for this Data Express transaction
}
