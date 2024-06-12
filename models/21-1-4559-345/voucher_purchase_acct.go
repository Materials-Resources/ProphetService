package model

import (
	"github.com/uptrace/bun"
	"time"
)

type VoucherPurchaseAcct struct {
	bun.BaseModel          `bun:"table:voucher_purchase_acct"`
	VoucherPurchaseAcctUid int32     `bun:"voucher_purchase_acct_uid,type:int,pk"`
	VoucherNo              string    `bun:"voucher_no,type:varchar(10)"`
	PurchaseAcctNo         string    `bun:"purchase_acct_no,type:varchar(32)"`
	PurchaseAmt            float64   `bun:"purchase_amt,type:decimal(19,2)"`
	PurchaseAmtDisplay     float64   `bun:"purchase_amt_display,type:decimal(19,2)"`
	PurchaseDesc           string    `bun:"purchase_desc,type:varchar(255),nullzero"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DisputedFlag           string    `bun:"disputed_flag,type:char,default:('N')"`
	IvaWithheldFlag        string    `bun:"iva_withheld_flag,type:char,nullzero"`
	IvaReceivedFlag        string    `bun:"iva_received_flag,type:char,nullzero"`
	TakeTerms              string    `bun:"take_terms,type:char,nullzero"`
	IvaReceivedPercent     float64   `bun:"iva_received_percent,type:decimal(19,2),nullzero"`
	IvaWithheldPercent     float64   `bun:"iva_withheld_percent,type:decimal(19,6),nullzero"`
	IvaReceivedEditedFlag  string    `bun:"iva_received_edited_flag,type:char,nullzero"`
	IvaWithheldEditedFlag  string    `bun:"iva_withheld_edited_flag,type:char,nullzero"`
}
