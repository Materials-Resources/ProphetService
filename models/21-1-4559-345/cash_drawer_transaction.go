package model

type CashDrawerTransaction struct {
	bun.BaseModel            `bun:"table:cash_drawer_transaction"`
	CashDrawerTransactionUid int32     `bun:"cash_drawer_transaction_uid,type:int,pk"`
	CashDrawerId             string    `bun:"cash_drawer_id,type:varchar(8)"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	SequenceNo               float64   `bun:"sequence_no,type:decimal(19,0)"`
	TransactionAmt           float64   `bun:"transaction_amt,type:decimal(19,4)"`
	TransactionAcctNo        string    `bun:"transaction_acct_no,type:varchar(32),nullzero"`
	TransactionTypeCd        int32     `bun:"transaction_type_cd,type:int"`
	CcUserId                 string    `bun:"cc_user_id,type:varchar(30),nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10),nullzero"`
	PaymentNumber            float64   `bun:"payment_number,type:decimal(19,0),nullzero"`
}
