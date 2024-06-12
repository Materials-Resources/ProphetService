package model

type CreditStatus struct {
	bun.BaseModel        `bun:"table:credit_status"`
	CreditStatusId       string    `bun:"credit_status_id,type:varchar(8)"`
	CreditStatusDesc     string    `bun:"credit_status_desc,type:varchar(40)"`
	OrderEntryAction     string    `bun:"order_entry_action,type:char"`
	InvoiceEntryAction   string    `bun:"invoice_entry_action,type:char"`
	ShippingAction       string    `bun:"shipping_action,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	PickTicketAction     string    `bun:"pick_ticket_action,type:char,nullzero"`
	ValidationAction     string    `bun:"validation_action,type:char,nullzero"`
	CreditStatusUid      int32     `bun:"credit_status_uid,type:int,pk"`
	RequireCcPaymentFlag string    `bun:"require_cc_payment_flag,type:char,default:('N')"`
	CcAcceptedFlag       string    `bun:"cc_accepted_flag,type:char,nullzero"`
}
