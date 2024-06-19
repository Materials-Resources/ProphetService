package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerOrderHistoryDaily struct {
	bun.BaseModel              `bun:"table:customer_order_history_daily"`
	CustomerOdrHistoryDailyUid int32     `bun:"customer_odr_history_daily_uid,type:int,autoincrement,pk"`     // Unique identifier
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`                                   // The company id associated with the customer
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0)"`                               // The customer identifier
	DateOrdered                time.Time `bun:"date_ordered,type:datetime"`                                   // Date of the orders
	ExtendedPrice              float64   `bun:"extended_price,type:decimal(19,4)"`                            // The price * quantity value
	ExtendedCost               float64   `bun:"extended_cost,type:decimal(19,4)"`                             // The price * cost value
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
