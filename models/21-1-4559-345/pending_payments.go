package model

type PendingPayments struct {
	bun.BaseModel          `bun:"table:pending_payments"`
	PendingPaymentUid      int32     `bun:"pending_payment_uid,type:int,pk"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	VoucherNo              string    `bun:"voucher_no,type:varchar(10)"`
	PendingAmount          float64   `bun:"pending_amount,type:decimal(19,4),default:(0)"`
	TermsAmountToPay       float64   `bun:"terms_amount_to_pay,type:decimal(19,4),default:(0)"`
	PayFlag                string    `bun:"pay_flag,type:char,default:('Y')"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(suser_name(null))"`
	IvaTermsAmountToPay    float64   `bun:"iva_terms_amount_to_pay,type:decimal(19,9),default:((0))"`
	IvaWithheldAmountToPay float64   `bun:"iva_withheld_amount_to_pay,type:decimal(19,9),default:((0))"`
}
