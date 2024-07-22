package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentAccount struct {
	bun.BaseModel        `bun:"table:payment_account"`
	PaymentAccountUid    int32      `bun:"payment_account_uid,type:int,pk"`                              // Unique identifier for the payment account information record
	PaymentAcctId        *string    `bun:"payment_acct_id,type:varchar(255)"`                            // An identifier for the payment account typically used by external entities (the
	PaymentAcctDesc      *string    `bun:"payment_acct_desc,type:varchar(255)"`                          // Display name for the payment account
	PaymentAcctTypeCd    int32      `bun:"payment_acct_type_cd,type:int"`                                // Iindicates the payment account type (i.e.: credit card, check/EFT, PayPal, FirePay, wire transfer, etc.)
	AcctNo               *string    `bun:"acct_no,type:varchar(255)"`                                    // The actual or masked account number (i.e.: masked credit card number, the PayPal email address, bank ABA number, etc.)
	ExtendedAcctInfo     *string    `bun:"extended_acct_info,type:varchar(255)"`                         // Ancillary account information (for example, a bank routing number)
	AcctExpirationDate   *time.Time `bun:"acct_expiration_date,type:datetime"`                           // Datetime indicating the expiration date (nullable in the case of check/EFT, PayPal, etc)
	RowStatusFlag        int32      `bun:"row_status_flag,type:int"`                                     // Indicates status of the current record
	DateCreated          time.Time  `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string     `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RoutingNumber        *string    `bun:"routing_number,type:varchar(255)"`                             // EFT bank routing number.
	BankName             *string    `bun:"bank_name,type:varchar(255)"`                                  // EFT bank name.
	BankCity             *string    `bun:"bank_city,type:varchar(40)"`                                   // EFT bank city.
	BankState            *string    `bun:"bank_state,type:varchar(2)"`                                   // EFT bakc state.
	TaxpayerAccountId    *string    `bun:"taxpayer_account_id,type:varchar(20)"`                         // Taxpayer Identification Number for this EFT payment account
	NetworkTransactionId *string    `bun:"network_transaction_id,type:varchar(30)"`                      // [Credential on File] - Identifies the last value returned by the network (i.e.: Visa, MasterCard, etc.) for a Credential on File (CoF) transaction.
}
