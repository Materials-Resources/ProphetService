package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentTypeMx struct {
	bun.BaseModel             `bun:"table:payment_type_mx"`
	PaymentTypeMxUid          int32     `bun:"payment_type_mx_uid,type:int,autoincrement,scanonly,pk"`       // PK
	PaymentTypeMxId           string    `bun:"payment_type_mx_id,type:varchar(2)"`                           // Payment Type Id
	PaymentTypeMxDesc         string    `bun:"payment_type_mx_desc,type:varchar(255)"`                       // Payment type description
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	IsBanked                  string    `bun:"is_banked,type:char(1),nullzero"`                              // Flag to indicate if payment type is banked
	OperationNumber           string    `bun:"operation_number,type:varchar(255),nullzero"`                  // Operation Number
	RfcPayerAccount           string    `bun:"rfc_payer_account,type:varchar(255),nullzero"`                 // RFC of the Payer Account
	PatternPayerAccount       string    `bun:"pattern_payer_account,type:varchar(255),nullzero"`             // Pattern for Payer Account
	RfcBeneficiaryAccount     string    `bun:"rfc_beneficiary_account,type:varchar(255),nullzero"`           // RFC of the Beneficiary Account
	BeneficiaryAccount        string    `bun:"beneficiary_account,type:varchar(255),nullzero"`               // Beneficiary_Account
	PatternBeneficiaryAccount string    `bun:"pattern_beneficiary_account,type:varchar(255),nullzero"`       // Pattern for Beneficiary_Account
	ChainPaymentType          string    `bun:"chain_payment_type,type:varchar(255),nullzero"`                // Chain Payment Type
	ForeignBankName           string    `bun:"foreign_bank_name,type:varchar(255),nullzero"`                 // Foreign Bank Name
	VersionNo                 float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`                 // Version Number defined by SAT
	RevisionNo                string    `bun:"revision_no,type:varchar(255),default:('0')"`                  // Revision Number defined by SAT
	ValidFromDate             time.Time `bun:"valid_from_date,type:datetime,nullzero"`                       // Valid date from defined by SAT
	ValidUntilDate            time.Time `bun:"valid_until_date,type:datetime,nullzero"`                      // Valid date until defined by SAT
	PayerAccount              string    `bun:"payer_account,type:varchar(255),nullzero"`                     // Payer Account
}
