package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentAccountXContact struct {
	bun.BaseModel             `bun:"table:payment_account_x_contact"`
	PaymentAccountXContactUid int32     `bun:"payment_account_x_contact_uid,type:int,pk"`
	PaymentAccountUid         int32     `bun:"payment_account_uid,type:int"`
	ContactId                 string    `bun:"contact_id,type:varchar(16)"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0)"`
	PaymentTypeId             float64   `bun:"payment_type_id,type:decimal(19,0)"`
	DefaultForContactFlag     string    `bun:"default_for_contact_flag,type:char,default:('N')"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CreditcardProcessorUid    int32     `bun:"creditcard_processor_uid,type:int,nullzero"`
	EpfMerchantAccountUid     int32     `bun:"epf_merchant_account_uid,type:int,nullzero"`
}
