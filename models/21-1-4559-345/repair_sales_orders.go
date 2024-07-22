package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type RepairSalesOrders struct {
	bun.BaseModel        `bun:"table:repair_sales_orders"`
	RepairSalesOrdersUid int32     `bun:"repair_sales_orders_uid,type:int,autoincrement,identity,pk"`
	RepairSalesOrdersRun int32     `bun:"repair_sales_orders_run,type:int,unique"`
	OrderNo              string    `bun:"order_no,type:varchar(255),unique"`
	DateRun              time.Time `bun:"date_run,type:datetime"`
}
