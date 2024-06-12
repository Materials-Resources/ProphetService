package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentTypes struct {
	bun.BaseModel            `bun:"table:payment_types"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	PaymentMethodId          string    `bun:"payment_method_id,type:varchar(2)"`
	PaymentTypeId            float64   `bun:"payment_type_id,type:decimal(19,0),pk"`
	PaymentTypeDesc          string    `bun:"payment_type_desc,type:varchar(30)"`
	AccountNumber            string    `bun:"account_number,type:varchar(32),nullzero"`
	Discount                 float64   `bun:"discount,type:decimal(4,2),nullzero"`
	FlatRate                 float64   `bun:"flat_rate,type:decimal(4,2),nullzero"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	WildcardAccount          string    `bun:"wildcard_account,type:char,default:('N')"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	UseCheckVerificationFlag string    `bun:"use_check_verification_flag,type:char,default:('N')"`
	CardTypeCd               int32     `bun:"card_type_cd,type:int,nullzero"`
	CashDiscountEligibleFlag string    `bun:"cash_discount_eligible_flag,type:char,default:('N')"`
	ConvenienceFeePercentage float64   `bun:"convenience_fee_percentage,type:decimal(19,9),nullzero"`
	ServiceVoucher           string    `bun:"service_voucher,type:char,nullzero"`
	BilltrustFlag            string    `bun:"billtrust_flag,type:char,nullzero"`
	PaypalFlag               string    `bun:"paypal_flag,type:char,nullzero"`
	DisplayInRmaFlag         string    `bun:"display_in_rma_flag,type:char,default:('Y')"`
	BankNoOverride           float64   `bun:"bank_no_override,type:decimal(19,0),nullzero"`
}
