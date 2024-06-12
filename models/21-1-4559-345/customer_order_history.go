package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerOrderHistory struct {
	bun.BaseModel           `bun:"table:customer_order_history"`
	CustomerOrderHistoryUid int32     `bun:"customer_order_history_uid,type:int,pk,identity"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`
	CustomerId              float64   `bun:"customer_id,type:decimal(19,0)"`
	YearOrdered             int32     `bun:"year_ordered,type:int"`
	MonthOrdered            int32     `bun:"month_ordered,type:int"`
	ExtendedPrice           float64   `bun:"extended_price,type:decimal(19,4),default:(0)"`
	ExtendedCost            float64   `bun:"extended_cost,type:decimal(19,4),default:(0)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
