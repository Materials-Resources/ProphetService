package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvHdrXInventoryReceipts struct {
	bun.BaseModel           `bun:"table:apinv_hdr_x_inventory_receipts"`
	ApinvHdrXInvReceiptsUid int32     `bun:"apinv_hdr_x_inv_receipts_uid,type:int,pk"`
	ReceiptNumber           float64   `bun:"receipt_number,type:decimal(19,0)"`
	VoucherNumber           string    `bun:"voucher_number,type:varchar(10),nullzero"`
	ExchangeRate            float64   `bun:"exchange_rate,type:decimal(19,6)"`
	ReceiptAmount           float64   `bun:"receipt_amount,type:decimal(19,4)"`
	ReceiptAmountForeign    float64   `bun:"receipt_amount_foreign,type:decimal(19,4)"`
	AmountPaid              float64   `bun:"amount_paid,type:decimal(19,4)"`
	AmountPaidForeign       float64   `bun:"amount_paid_foreign,type:decimal(19,4)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:('CURRENT_TIMESTAMP')"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:('CURRENT_TIMESTAMP')"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:('SYSTEM_USER')"`
}
