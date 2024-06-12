package model

type PaymentAccount struct {
	bun.BaseModel        `bun:"table:payment_account"`
	PaymentAccountUid    int32     `bun:"payment_account_uid,type:int,pk"`
	PaymentAcctId        string    `bun:"payment_acct_id,type:varchar(255),nullzero"`
	PaymentAcctDesc      string    `bun:"payment_acct_desc,type:varchar(255),nullzero"`
	PaymentAcctTypeCd    int32     `bun:"payment_acct_type_cd,type:int"`
	AcctNo               string    `bun:"acct_no,type:varchar(255),nullzero"`
	ExtendedAcctInfo     string    `bun:"extended_acct_info,type:varchar(255),nullzero"`
	AcctExpirationDate   time.Time `bun:"acct_expiration_date,type:datetime,nullzero"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RoutingNumber        string    `bun:"routing_number,type:varchar(255),nullzero"`
	BankName             string    `bun:"bank_name,type:varchar(255),nullzero"`
	BankCity             string    `bun:"bank_city,type:varchar(40),nullzero"`
	BankState            string    `bun:"bank_state,type:varchar(2),nullzero"`
	TaxpayerAccountId    string    `bun:"taxpayer_account_id,type:varchar(20),nullzero"`
	NetworkTransactionId string    `bun:"network_transaction_id,type:varchar(30),nullzero"`
}
