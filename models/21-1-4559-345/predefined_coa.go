package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PredefinedCoa struct {
	bun.BaseModel      `bun:"table:predefined_coa"`
	PredefinedCoaUid   int32     `bun:"predefined_coa_uid,type:int,autoincrement,identity,pk"`        // Unique Identifier for the record.
	AccountNumber      string    `bun:"account_number,type:varchar(10),unique"`                       // Account Number
	AccountDescription string    `bun:"account_description,type:varchar(255)"`                        // Account Description
	AccountType        string    `bun:"account_type,type:char(1)"`                                    // Account Type (Asset, Expense, Liability, etc...)
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
