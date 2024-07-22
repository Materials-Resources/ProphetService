package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardTransrequest struct {
	bun.BaseModel             `bun:"table:creditcard_transrequest"`
	CreditcardTransrequestUid int32     `bun:"creditcard_transrequest_uid,type:int,pk"`                     // UID of this table
	PaymentNumber             *float64  `bun:"payment_number,type:decimal(19,0)"`                           // Reference number to payment information
	ReqfileLocation           *string   `bun:"reqfile_location,type:varchar(255)"`                          // The path to where the request file was created
	RequesttypeId             int32     `bun:"requesttype_id,type:int"`                                     // ID of credit card transaction type.  The ID is from creditcard_transtype.
	RequestStatus             *string   `bun:"request_status,type:varchar(10)"`                             // Status of the credit card transaction request
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`              // Indicates the date/time this record was created.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"` // ID of the user who last maintained this record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`        // Indicates the date/time this record was last modified.
}
