package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CashDrawerTransaction struct {
	bun.BaseModel            `bun:"table:cash_drawer_transaction"`
	CashDrawerTransactionUid int32     `bun:"cash_drawer_transaction_uid,type:int,pk"`                      // Uniquely Identifies each record in the table.
	CashDrawerId             string    `bun:"cash_drawer_id,type:varchar(8)"`                               // Identifier for drawer
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                   // Identifier for the company associated with this cash drawer.
	SequenceNo               float64   `bun:"sequence_no,type:decimal(19,0)"`                               // Current active history record.
	TransactionAmt           float64   `bun:"transaction_amt,type:decimal(19,4)"`                           // Transaction amount for the cash drawer transaction.
	TransactionAcctNo        string    `bun:"transaction_acct_no,type:varchar(32),nullzero"`                // Transaction account number - Account when making a withdrawal
	TransactionTypeCd        int32     `bun:"transaction_type_cd,type:int"`                                 // Remittance / Deposit / Withdrawal.
	CcUserId                 string    `bun:"cc_user_id,type:varchar(30),nullzero"`                         // User responsible for this transaction.
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10),nullzero"`                         // Invoice Number for transactions not tied to a remittance/order
	PaymentNumber            float64   `bun:"payment_number,type:decimal(19,0),nullzero"`                   // Associated Payment NUmber
}
