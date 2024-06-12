package model

type PaymentAccountXCustomer struct {
	bun.BaseModel              `bun:"table:payment_account_x_customer"`
	PaymentAccountXCustomerUid int32     `bun:"payment_account_x_customer_uid,type:int,pk"`
	PaymentAccountUid          int32     `bun:"payment_account_uid,type:int"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0)"`
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`
	PaymentTypeId              float64   `bun:"payment_type_id,type:decimal(19,0)"`
	RowStatusFlag              int32     `bun:"row_status_flag,type:int"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DefaultForCustomerFlag     string    `bun:"default_for_customer_flag,type:char,default:('N')"`
	CreditcardProcessorUid     int32     `bun:"creditcard_processor_uid,type:int,nullzero"`
	AutomaticPaymentFlag       string    `bun:"automatic_payment_flag,type:char,nullzero"`
	AutomaticPaymentLimit      float64   `bun:"automatic_payment_limit,type:decimal(19,9),nullzero"`
	EpfMerchantAccountUid      int32     `bun:"epf_merchant_account_uid,type:int,nullzero"`
}
