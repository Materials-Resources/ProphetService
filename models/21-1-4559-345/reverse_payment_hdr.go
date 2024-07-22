package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ReversePaymentHdr struct {
	bun.BaseModel          `bun:"table:reverse_payment_hdr"`
	ReversePaymentHdrUid   int32     `bun:"reverse_payment_hdr_uid,type:int,pk"`                          // Unique ID for Reverse Payment Header table
	CompanyId              string    `bun:"company_id,type:varchar(8)"`                                   // Company for which the Cash Receipt is reversed
	Period                 float64   `bun:"period,type:decimal(3,0)"`                                     // Period in which the postings are applied
	YearForPeriod          float64   `bun:"year_for_period,type:decimal(4,0)"`                            // Year in which the postings are applied
	ReversePaymentApproved string    `bun:"reverse_payment_approved,type:char(1),default:('N')"`          // Whether the Cash Receipt Reversal has been approved
	DateReversed           time.Time `bun:"date_reversed,type:datetime"`                                  // Date on which the Cash Receipt is reversed
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
