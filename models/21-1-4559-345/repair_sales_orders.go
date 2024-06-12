package model

type RepairSalesOrders struct {
	bun.BaseModel        `bun:"table:repair_sales_orders"`
	RepairSalesOrdersUid int32     `bun:"repair_sales_orders_uid,type:int,pk,identity"`
	RepairSalesOrdersRun int32     `bun:"repair_sales_orders_run,type:int"`
	OrderNo              string    `bun:"order_no,type:varchar(255)"`
	DateRun              time.Time `bun:"date_run,type:datetime"`
}
