package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentAccountXContact struct {
	bun.BaseModel             `bun:"table:payment_account_x_contact"`
	PaymentAccountXContactUid int32     `bun:"payment_account_x_contact_uid,type:int,pk"` // Unique identifier for the payment account by contact row
	PaymentAccountUid         int32     `bun:"payment_account_uid,type:int,unique"`       // Unique identifier for the payment account information record
	ContactId                 string    `bun:"contact_id,type:varchar(16),unique"`        // The id associated with the contact
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0)"`
	PaymentTypeId             float64   `bun:"payment_type_id,type:decimal(19,0),unique"` // System generated payment type identifier
	DefaultForContactFlag     string    `bun:"default_for_contact_flag,type:char(1),default:('N')"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`                                     // Indicates status of the current record
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	CreditcardProcessorUid    *int32    `bun:"creditcard_processor_uid,type:int"`                            // Associates, where necessary, the Contact Payment Account with an electronic payment processor record.
	EpfMerchantAccountUid     *int32    `bun:"epf_merchant_account_uid,type:int"`                            // Associates, where necessary, the Contact Payment Account with an EPF merchant account record.
}
