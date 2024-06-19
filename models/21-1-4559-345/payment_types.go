package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentTypes struct {
	bun.BaseModel            `bun:"table:payment_types"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	PaymentMethodId          string    `bun:"payment_method_id,type:varchar(2)"`                         // Payment Method Identifier
	PaymentTypeId            float64   `bun:"payment_type_id,type:decimal(19,0),pk"`                     // System generated payment type identifier
	PaymentTypeDesc          string    `bun:"payment_type_desc,type:varchar(30)"`                        // A descriptive field
	AccountNumber            string    `bun:"account_number,type:varchar(32),nullzero"`                  // Enter a valid account number
	Discount                 float64   `bun:"discount,type:decimal(4,2),nullzero"`                       // Discouint rate charged by credit card company
	FlatRate                 float64   `bun:"flat_rate,type:decimal(4,2),nullzero"`                      // flat rate per transaction charged by credit card c
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	WildcardAccount          string    `bun:"wildcard_account,type:char(1),default:('N')"`               // Determines whether branches are automatically wildcarded on the GL account
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	UseCheckVerificationFlag string    `bun:"use_check_verification_flag,type:char(1),default:('N')"`    // Indicates which check payment types will be submitted to protobase for verification
	CardTypeCd               int32     `bun:"card_type_cd,type:int,nullzero"`                            // Specific Credit Card types used by Realex
	CashDiscountEligibleFlag string    `bun:"cash_discount_eligible_flag,type:char(1),default:('N')"`    // Indicates if payment type is eligible for cash discount
	ConvenienceFeePercentage float64   `bun:"convenience_fee_percentage,type:decimal(19,9),nullzero"`    // Convenience Fee percentage for credit card type payments
	ServiceVoucher           string    `bun:"service_voucher,type:char(1),nullzero"`                     // Indicate that the payment is via a service voucher
	BilltrustFlag            string    `bun:"billtrust_flag,type:char(1),nullzero"`                      // Custom: Indicates if this payment type is a billtrust payment.
	PaypalFlag               string    `bun:"paypal_flag,type:char(1),nullzero"`                         // PayPal indicator
	DisplayInRmaFlag         string    `bun:"display_in_rma_flag,type:char(1),default:('Y')"`            // Display in RMAs Flag
	BankNoOverride           float64   `bun:"bank_no_override,type:decimal(19,0),nullzero"`              // Allows user to specify a bank that is always used in Bank reconciliation process for this cash/check payment type
}
