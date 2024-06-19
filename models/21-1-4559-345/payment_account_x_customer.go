package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentAccountXCustomer struct {
	bun.BaseModel              `bun:"table:payment_account_x_customer"`
	PaymentAccountXCustomerUid int32     `bun:"payment_account_x_customer_uid,type:int,pk"`                   // Unique identifier for the payment account by customer row
	PaymentAccountUid          int32     `bun:"payment_account_uid,type:int,unique"`                          // Unique identifier for the payment account information record
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),unique"`                        // System-generated ID that identifies customers.
	CompanyId                  string    `bun:"company_id,type:varchar(8),unique"`                            // Unique code that identifies a company.
	PaymentTypeId              float64   `bun:"payment_type_id,type:decimal(19,0),unique"`                    // System generated payment type identifier
	RowStatusFlag              int32     `bun:"row_status_flag,type:int"`                                     // Indicates the status of the record
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DefaultForCustomerFlag     string    `bun:"default_for_customer_flag,type:char(1),default:('N')"`         // Indicates the Customer Payment Account record is the default for the associated Company/Customer.
	CreditcardProcessorUid     int32     `bun:"creditcard_processor_uid,type:int,nullzero"`                   // Associates, where necessary, the Customer Payment Account with an electronic payment processor record.
	AutomaticPaymentFlag       string    `bun:"automatic_payment_flag,type:char(1),nullzero"`                 // Column flag to indicates an automatic payment process.
	AutomaticPaymentLimit      float64   `bun:"automatic_payment_limit,type:decimal(19,9),nullzero"`          // Indicates the limit amount when corresponding flag is selected.
	EpfMerchantAccountUid      int32     `bun:"epf_merchant_account_uid,type:int,nullzero"`                   // Associates, where necessary, the Customer Payment Account with an EPF merchant account record.
}
