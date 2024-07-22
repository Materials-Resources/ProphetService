package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Payment1099Detail struct {
	bun.BaseModel        `bun:"table:payment_1099_detail"`
	CompanyId            string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	VendorId             float64   `bun:"vendor_id,type:decimal(19,0)"`                              // What is the unique vendor identifier for this row?
	Year                 float64   `bun:"year,type:decimal(4,0)"`                                    // Year of the payment
	BankNo               float64   `bun:"bank_no,type:decimal(19,0)"`                                // Default bank number
	CheckNo              string    `bun:"check_no,type:varchar(255)"`                                // What was the check number?
	Ten99Type            float64   `bun:"ten99_type,type:decimal(19,0)"`                             // Enter the 1099 type
	Amount               float64   `bun:"amount,type:decimal(19,4)"`                                 // payment amount.
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Payment1099DetailUid int32     `bun:"payment_1099_detail_uid,type:int,pk"`                       // unique identifier of the table.
}
