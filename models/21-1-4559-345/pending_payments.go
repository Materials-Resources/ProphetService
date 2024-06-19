package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PendingPayments struct {
	bun.BaseModel          `bun:"table:pending_payments"`
	PendingPaymentUid      int32     `bun:"pending_payment_uid,type:int,pk"`                                // This is the system-generated unique identifier for a pending payment
	CompanyId              string    `bun:"company_id,type:varchar(8)"`                                     // Unique code that identifies a company.
	VoucherNo              string    `bun:"voucher_no,type:varchar(10)"`                                    // Voucher Number
	PendingAmount          float64   `bun:"pending_amount,type:decimal(19,4),default:(0)"`                  // How much is the pending payment for?
	TermsAmountToPay       float64   `bun:"terms_amount_to_pay,type:decimal(19,4),default:(0)"`             // Terms amount.
	PayFlag                string    `bun:"pay_flag,type:char(1),default:('Y')"`                            // Indicates if the voucher was marked for payment
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(suser_name(null))"` // ID of the user who last maintained this record
	IvaTermsAmountToPay    float64   `bun:"iva_terms_amount_to_pay,type:decimal(19,9),default:((0))"`       // IVA Terms Amount to be Paid
	IvaWithheldAmountToPay float64   `bun:"iva_withheld_amount_to_pay,type:decimal(19,9),default:((0))"`    // IVA Withheld Amount to be Paid
}
