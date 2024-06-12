package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentTypeMx struct {
	bun.BaseModel             `bun:"table:payment_type_mx"`
	PaymentTypeMxUid          int32     `bun:"payment_type_mx_uid,type:int,pk,identity"`
	PaymentTypeMxId           string    `bun:"payment_type_mx_id,type:varchar(2)"`
	PaymentTypeMxDesc         string    `bun:"payment_type_mx_desc,type:varchar(255)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	IsBanked                  string    `bun:"is_banked,type:char,nullzero"`
	OperationNumber           string    `bun:"operation_number,type:varchar(255),nullzero"`
	RfcPayerAccount           string    `bun:"rfc_payer_account,type:varchar(255),nullzero"`
	PatternPayerAccount       string    `bun:"pattern_payer_account,type:varchar(255),nullzero"`
	RfcBeneficiaryAccount     string    `bun:"rfc_beneficiary_account,type:varchar(255),nullzero"`
	BeneficiaryAccount        string    `bun:"beneficiary_account,type:varchar(255),nullzero"`
	PatternBeneficiaryAccount string    `bun:"pattern_beneficiary_account,type:varchar(255),nullzero"`
	ChainPaymentType          string    `bun:"chain_payment_type,type:varchar(255),nullzero"`
	ForeignBankName           string    `bun:"foreign_bank_name,type:varchar(255),nullzero"`
	VersionNo                 float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`
	RevisionNo                string    `bun:"revision_no,type:varchar(255),default:('0')"`
	ValidFromDate             time.Time `bun:"valid_from_date,type:datetime,nullzero"`
	ValidUntilDate            time.Time `bun:"valid_until_date,type:datetime,nullzero"`
	PayerAccount              string    `bun:"payer_account,type:varchar(255),nullzero"`
}
