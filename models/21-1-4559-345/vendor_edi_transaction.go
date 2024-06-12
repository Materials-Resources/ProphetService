package model

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorEdiTransaction struct {
	bun.BaseModel           `bun:"table:vendor_edi_transaction"`
	VendorEdiTransactionUid int32     `bun:"vendor_edi_transaction_uid,type:int,pk"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`
	VendorId                float64   `bun:"vendor_id,type:decimal(19,0)"`
	EdiTransaction          int32     `bun:"edi_transaction,type:int"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
