package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceCustShiptoOrdlim struct {
	bun.BaseModel               `bun:"table:job_price_cust_shipto_ordlim"`
	JobPriceCustshiptoordlimUid int32     `bun:"job_price_custshiptoordlim_uid,type:int,autoincrement,identity,pk"` // Unique internal ID number.
	JobPriceCustShiptoUid       int32     `bun:"job_price_cust_shipto_uid,type:int"`                                // FK to column job_price_customer_shipto.job_price_cust_shipto_uid.  Link to associated customer/ship-to instance.
	JobPriceHdrBudgetPrdUid     int32     `bun:"job_price_hdr_budget_prd_uid,type:int"`                             // FK to column job_price_hdr_budget_prd.job_price_hdr_budget_prd_uid.  Link to associated budget period instance.
	OrderLimit                  int32     `bun:"order_limit,type:int,default:((0))"`                                // Maximum number of orders that can be associated w/this contract for the specified customer/ship-to/period.
	OrderLimitUsed              int32     `bun:"order_limit_used,type:int"`                                         // Number of orders that have been associated w/this contract for the specified customer/ship-to/period.
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
