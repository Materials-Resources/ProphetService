package model

import (
	"github.com/uptrace/bun"
	"time"
)

type RoundingInstallment10005 struct {
	bun.BaseModel           `bun:"table:rounding_installment_10005"`
	RoundingInstallmentUid  int32     `bun:"rounding_installment_uid,type:int,pk"`
	RoundingInstallmentName string    `bun:"rounding_installment_name,type:varchar(20)"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30)"`
	RoundingInstallmentCode string    `bun:"rounding_installment_code,type:char(3)"`
}
