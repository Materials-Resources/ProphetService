package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardValidation struct {
	bun.BaseModel           `bun:"table:creditcard_validation"`
	CreditcardValidationUid int32     `bun:"creditcard_validation_uid,type:int,autoincrement,identity,pk"` // Unique identifier for table.
	CardTypeCd              int32     `bun:"card_type_cd,type:int"`                                        // Credit Card type code
	ValidationExpression    string    `bun:"validation_expression,type:varchar(255)"`                      // Credit Card Number Validation Regular Expression
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
