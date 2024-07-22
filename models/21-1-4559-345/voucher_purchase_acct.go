package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VoucherPurchaseAcct struct {
	bun.BaseModel          `bun:"table:voucher_purchase_acct"`
	VoucherPurchaseAcctUid int32     `bun:"voucher_purchase_acct_uid,type:int,pk"`                        // Unique identifier
	VoucherNo              string    `bun:"voucher_no,type:varchar(10)"`                                  // Associated voucher number
	PurchaseAcctNo         string    `bun:"purchase_acct_no,type:varchar(32)"`                            // Purchase account number
	PurchaseAmt            float64   `bun:"purchase_amt,type:decimal(19,2)"`                              // Amount in home currency
	PurchaseAmtDisplay     float64   `bun:"purchase_amt_display,type:decimal(19,2)"`                      // Amount if foreign currency
	PurchaseDesc           *string   `bun:"purchase_desc,type:varchar(255)"`                              // Description of transaction
	CompanyId              string    `bun:"company_id,type:varchar(8)"`                                   // Company ID
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DisputedFlag           string    `bun:"disputed_flag,type:char(1),default:('N')"`                     // Flag indicating whether the charge is disputed
	IvaWithheldFlag        *string   `bun:"iva_withheld_flag,type:char(1)"`                               // Indicate the record is for IVA Withheld
	IvaReceivedFlag        *string   `bun:"iva_received_flag,type:char(1)"`                               // Indicate teh record is for IVA received
	TakeTerms              *string   `bun:"take_terms,type:char(1)"`                                      // Flag  to indicate whether we want these charges to be included in the terms
	IvaReceivedPercent     *float64  `bun:"iva_received_percent,type:decimal(19,2)"`                      // IVA Received Percent
	IvaWithheldPercent     *float64  `bun:"iva_withheld_percent,type:decimal(19,6)"`                      // IVA Withheld Percent
	IvaReceivedEditedFlag  *string   `bun:"iva_received_edited_flag,type:char(1)"`                        // IVA Received Amount Edited
	IvaWithheldEditedFlag  *string   `bun:"iva_withheld_edited_flag,type:char(1)"`                        // IVA Withheld Amount Edited
}
