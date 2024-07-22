package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TransactionSet struct {
	bun.BaseModel      `bun:"table:transaction_set"`
	TransactionSetUid  int32     `bun:"transaction_set_uid,type:int,pk"`                           // Unique Identifier of record
	TransactionSetId   string    `bun:"transaction_set_id,type:varchar(32)"`                       // Identifies the transaction set, ex InvRetShip.
	TransactionSetDesc string    `bun:"transaction_set_desc,type:varchar(255)"`                    // Friendly description of transaction set, ex Inventory Return Shipping.
	DateCreated        time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Enablefor          *string   `bun:"enablefor,type:char(6)"`                                    // Whether the transaction set is enabled for import, export or both.
	AllowXml           *string   `bun:"allow_xml,type:char(1),default:('N')"`                      // Whether the XML data can be used.
	P21CodeNo          *int32    `bun:"p21_code_no,type:int"`                                      // P21_code table value used in customer or vendor edi transaction tables
	EdiTransactionId   *string   `bun:"edi_transaction_id,type:varchar(255)"`                      // X12 transaction code, 810, 820 etc.
}
