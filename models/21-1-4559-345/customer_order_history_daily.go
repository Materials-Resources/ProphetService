package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerOrderHistoryDaily struct {
	bun.BaseModel              `bun:"table:customer_order_history_daily"`
	CustomerOdrHistoryDailyUid int32     `bun:"customer_odr_history_daily_uid,type:int,pk,identity"`
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0)"`
	DateOrdered                time.Time `bun:"date_ordered,type:datetime"`
	ExtendedPrice              float64   `bun:"extended_price,type:decimal(19,4)"`
	ExtendedCost               float64   `bun:"extended_cost,type:decimal(19,4)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
