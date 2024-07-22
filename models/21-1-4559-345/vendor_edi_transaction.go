package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorEdiTransaction struct {
	bun.BaseModel           `bun:"table:vendor_edi_transaction"`
	VendorEdiTransactionUid int32     `bun:"vendor_edi_transaction_uid,type:int,pk"`                    // Unique Identifier
	CompanyId               string    `bun:"company_id,type:varchar(8),unique"`                         // Unique code that identifies a company.
	VendorId                float64   `bun:"vendor_id,type:decimal(19,0),unique"`                       // Vendor ID for this record.
	EdiTransaction          int32     `bun:"edi_transaction,type:int,unique"`                           // EDI transaction set for this record, ex 850 PO.
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
