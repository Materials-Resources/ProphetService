package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceHdrBudgetPrd struct {
	bun.BaseModel           `bun:"table:job_price_hdr_budget_prd"`
	JobPriceHdrBudgetPrdUid int32     `bun:"job_price_hdr_budget_prd_uid,type:int,pk"`                     // Unique internal ID number.
	JobPriceHdrUid          int32     `bun:"job_price_hdr_uid,type:int"`                                   // FK to column job_price_hdr.job_price_hdr_uid.  Link to associated job_price_hdr instance.
	Period                  float64   `bun:"period,type:decimal(3,0)"`                                     // Period number.
	YearForPeriod           float64   `bun:"year_for_period,type:decimal(4,0)"`                            // Year for the Period.
	BeginningDate           time.Time `bun:"beginning_date,type:datetime"`                                 // Date the period begins.
	EndingDate              time.Time `bun:"ending_date,type:datetime"`                                    // Date the period ends.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
