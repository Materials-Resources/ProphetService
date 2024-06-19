package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvHdrXInventoryReceipts struct {
	bun.BaseModel           `bun:"table:apinv_hdr_x_inventory_receipts"`
	ApinvHdrXInvReceiptsUid int32     `bun:"apinv_hdr_x_inv_receipts_uid,type:int,pk"`                       // Unique identifier for table apinv_hdr_x_inventory_receipts
	ReceiptNumber           float64   `bun:"receipt_number,type:decimal(19,0),unique"`                       // Inventory receipt number
	VoucherNumber           string    `bun:"voucher_number,type:varchar(10),unique,nullzero"`                // Voucher number created by the receipt number
	ExchangeRate            float64   `bun:"exchange_rate,type:decimal(19,6)"`                               // Exchange rate calculated at the time of the inventory receipt
	ReceiptAmount           float64   `bun:"receipt_amount,type:decimal(19,4)"`                              // Inventory receipt amount used to create the voucher
	ReceiptAmountForeign    float64   `bun:"receipt_amount_foreign,type:decimal(19,4)"`                      // Inventory receipt amount used to create the voucher (This amount is in terms of the vendor or supplier currency)
	AmountPaid              float64   `bun:"amount_paid,type:decimal(19,4)"`                                 // How much of the receipt amount has been paid
	AmountPaidForeign       float64   `bun:"amount_paid_foreign,type:decimal(19,4)"`                         // How much of the receipt amount has been paid (This amount is in terms of the vendor or supplier currency)
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                                       // Indicates current record status.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:('CURRENT_TIMESTAMP')"`       // Date and time the record was originally created
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:('CURRENT_TIMESTAMP')"` // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:('SYSTEM_USER')"`    // User who last changed the record
}
