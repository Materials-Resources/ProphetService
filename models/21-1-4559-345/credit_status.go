package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditStatus struct {
	bun.BaseModel        `bun:"table:credit_status"`
	CreditStatusId       string    `bun:"credit_status_id,type:varchar(8),unique"`                   // Unique identifier for this credit status.
	CreditStatusDesc     string    `bun:"credit_status_desc,type:varchar(40)"`                       // A description of the credit status.
	OrderEntryAction     string    `bun:"order_entry_action,type:char(1)"`                           // An option that effects order entry transactions for a customer with this Credit Status ID.
	InvoiceEntryAction   string    `bun:"invoice_entry_action,type:char(1)"`                         // An option that effects invoice entry transactions for a customer with this Credit Status ID.
	ShippingAction       string    `bun:"shipping_action,type:char(1)"`                              // An option that effects shipping transactions for a customer with this Credit Status ID.
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	PickTicketAction     string    `bun:"pick_ticket_action,type:char(1),nullzero"`                  // An option that effects pick ticket transactions for a customer with this Credit Status ID.
	ValidationAction     string    `bun:"validation_action,type:char(1),nullzero"`                   // Determines the effect of a customers credit status on their orders.
	CreditStatusUid      int32     `bun:"credit_status_uid,type:int,pk"`                             // Unique key for credit status record
	RequireCcPaymentFlag string    `bun:"require_cc_payment_flag,type:char(1),default:('N')"`        // This is a flag that indicates whether or not a credit card payment is require in order entry
	CcAcceptedFlag       string    `bun:"cc_accepted_flag,type:char(1),nullzero"`                    // Credit Card Accepted Credit Status
}
