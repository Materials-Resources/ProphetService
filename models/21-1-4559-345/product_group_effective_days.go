package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProductGroupEffectiveDays struct {
	bun.BaseModel             `bun:"table:product_group_effective_days"`
	ProdGroupEffectiveDaysUid int32     `bun:"prod_group_effective_days_uid,type:int,pk"`                    // this will be the primary key counter column
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8)"`                             // the product group id
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`                                   // this is the company id
	EffectiveDays             int32     `bun:"effective_days,type:int,default:((0))"`                        // number of days a job quote remains effective
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	JobEntryFlag              string    `bun:"job_entry_flag,type:char(1),default:('Y')"`                    // Flag indicating if price controls are used in Job Quote Entry
	OrderEntryFlag            string    `bun:"order_entry_flag,type:char(1),default:('Y')"`                  // Flag indicating if price controls are used in Order Entry
	JobMinGrossProfit         float64   `bun:"job_min_gross_profit,type:decimal(19,9),default:((0))"`        // Minimum gross profit for Job Quote Entry
	OrderMinGrossProfit       float64   `bun:"order_min_gross_profit,type:decimal(19,9),default:((0))"`      // Minimum gross profit for Order Entry
}
