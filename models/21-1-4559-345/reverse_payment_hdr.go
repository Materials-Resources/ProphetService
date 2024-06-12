package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReversePaymentHdr struct {
	bun.BaseModel          `bun:"table:reverse_payment_hdr"`
	ReversePaymentHdrUid   int32     `bun:"reverse_payment_hdr_uid,type:int,pk"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	Period                 float64   `bun:"period,type:decimal(3,0)"`
	YearForPeriod          float64   `bun:"year_for_period,type:decimal(4,0)"`
	ReversePaymentApproved string    `bun:"reverse_payment_approved,type:char,default:('N')"`
	DateReversed           time.Time `bun:"date_reversed,type:datetime"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
