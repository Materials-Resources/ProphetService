package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerCall struct {
	bun.BaseModel       `bun:"table:customer_call"`
	CustomerCallUid     int32     `bun:"customer_call_uid,type:int,pk"`               // Unique Identifier of customer_call table.
	CompanyId           string    `bun:"company_id,type:varchar(8)"`                  // Unique code that identifies a company.
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`              // Customer paying invoice -  remitter
	TopicUid            int32     `bun:"topic_uid,type:int"`                          // Unique Identifier of topic table.
	ContactId           string    `bun:"contact_id,type:varchar(16)"`                 // What contact deals with this sales representative.
	NumberOfTimesCalled int16     `bun:"number_of_times_called,type:smallint"`        // Number of times the customer has been contacted.
	PromisedDate        time.Time `bun:"promised_date,type:datetime,nullzero"`        // Date customer promised to send a payment.
	PromisedAmount      float64   `bun:"promised_amount,type:decimal(19,4),nullzero"` // Payment amount customer promised to send.
	CallBackDate        time.Time `bun:"call_back_date,type:datetime"`                // Date to follow up with the customer.
	Notes               string    `bun:"notes,type:text"`                             // Notes regarding the call.
	RowStatusFlag       int16     `bun:"row_status_flag,type:smallint"`               // Indicates current record status.
	DateCreated         time.Time `bun:"date_created,type:datetime"`                  // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`            // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`         // ID of the user who last maintained this record
}
