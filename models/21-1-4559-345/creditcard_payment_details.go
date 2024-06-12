package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardPaymentDetails struct {
	bun.BaseModel                   `bun:"table:creditcard_payment_details"`
	CreditcardPaymentDetailsUid     int32     `bun:"creditcard_payment_details_uid,type:int,pk"`
	PaymentNumber                   float64   `bun:"payment_number,type:decimal(19,0)"`
	FirstName                       string    `bun:"first_name,type:varchar(255),nullzero"`
	LastName                        string    `bun:"last_name,type:varchar(255),nullzero"`
	StreetAddress1                  string    `bun:"street_address1,type:varchar(255),nullzero"`
	StreetAddress2                  string    `bun:"street_address2,type:varchar(255),nullzero"`
	City                            string    `bun:"city,type:varchar(25),nullzero"`
	State                           string    `bun:"state,type:char(2),nullzero"`
	ZipCode                         string    `bun:"zip_code,type:varchar(10),nullzero"`
	OriginalAuthAmount              float64   `bun:"original_auth_amount,type:decimal(19,4),nullzero"`
	AuthAmount                      float64   `bun:"auth_amount,type:decimal(19,4),nullzero"`
	TransSettled                    string    `bun:"trans_settled,type:char"`
	Taker                           string    `bun:"taker,type:varchar(30)"`
	DateCreated                     time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	LastMaintainedBy                string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"`
	DateLastModified                time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	BatchNumber                     string    `bun:"batch_number,type:varchar(20),nullzero"`
	RetrievalRefNumber              string    `bun:"retrieval_ref_number,type:varchar(20),nullzero"`
	CustomerVerificationValue       string    `bun:"customer_verification_value,type:varchar(10),nullzero"`
	UserSpecifiedAmount             float64   `bun:"user_specified_amount,type:decimal(19,4),default:(0)"`
	SwitchIssueNumber               string    `bun:"switch_issue_number,type:varchar(2),nullzero"`
	RealexReferenceNumber           string    `bun:"realex_reference_number,type:varchar(255),nullzero"`
	RealexOrderId                   string    `bun:"realex_order_id,type:varchar(255),nullzero"`
	PaymentAccountId                string    `bun:"payment_account_id,type:varchar(255),nullzero"`
	AddressEditedFlag               string    `bun:"address_edited_flag,type:char,default:('N')"`
	CommercialCardResponseCode      string    `bun:"commercial_card_response_code,type:varchar(30),nullzero"`
	CardBrand                       string    `bun:"card_brand,type:varchar(30),nullzero"`
	MarketCodeValue                 string    `bun:"market_code_value,type:varchar(255),nullzero"`
	SwipedTransactionFlag           string    `bun:"swiped_transaction_flag,type:char,default:('N')"`
	EcommerceTransactionFlag        string    `bun:"ecommerce_transaction_flag,type:char,default:('N')"`
	TransactionCardholderLocationCd int32     `bun:"transaction_cardholder_location_cd,type:int,default:((2815))"`
	StreetAddress3                  string    `bun:"street_address3,type:varchar(50),nullzero"`
}
