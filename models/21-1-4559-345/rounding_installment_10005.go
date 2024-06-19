package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RoundingInstallment10005 struct {
	bun.BaseModel           `bun:"table:rounding_installment_10005"`
	RoundingInstallmentUid  int32     `bun:"rounding_installment_uid,type:int,pk"`       // Unique internal record identifier - not visible to the user
	RoundingInstallmentName string    `bun:"rounding_installment_name,type:varchar(20)"` // Displayed roudning option (First, Last)
	DateCreated             time.Time `bun:"date_created,type:datetime"`                 // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`           // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30)"`        // ID of the user who last maintained this record
	RoundingInstallmentCode string    `bun:"rounding_installment_code,type:char(3)"`     // Internal code for rounding option
}
