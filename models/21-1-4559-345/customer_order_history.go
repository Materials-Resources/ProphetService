package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerOrderHistory struct {
	bun.BaseModel           `bun:"table:customer_order_history"`
	CustomerOrderHistoryUid int32     `bun:"customer_order_history_uid,type:int,autoincrement,pk"` // Unique Identifier
	CompanyId               string    `bun:"company_id,type:varchar(8)"`                           // The company id associated with the customer
	CustomerId              float64   `bun:"customer_id,type:decimal(19,0)"`                       // The customer identifier
	YearOrdered             int32     `bun:"year_ordered,type:int"`                                // Year of the orders
	MonthOrdered            int32     `bun:"month_ordered,type:int"`                               // Month of the orders
	ExtendedPrice           float64   `bun:"extended_price,type:decimal(19,4),default:(0)"`        // The price * quantity value
	ExtendedCost            float64   `bun:"extended_cost,type:decimal(19,4),default:(0)"`         // The price * cost value
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`       // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"` // User who created the record
}
