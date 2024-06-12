package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardValidation struct {
	bun.BaseModel           `bun:"table:creditcard_validation"`
	CreditcardValidationUid int32     `bun:"creditcard_validation_uid,type:int,pk,identity"`
	CardTypeCd              int32     `bun:"card_type_cd,type:int"`
	ValidationExpression    string    `bun:"validation_expression,type:varchar(255)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
