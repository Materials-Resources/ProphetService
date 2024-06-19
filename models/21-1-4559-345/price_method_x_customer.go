package gen

import "github.com/uptrace/bun"

type PriceMethodXCustomer struct {
	bun.BaseModel           `bun:"table:price_method_x_customer"`
	PriceMethodXCustomerUid int64  `bun:"price_method_x_customer_uid,type:bigint,autoincrement,identity"`
	PriceMethodUid          int64  `bun:"price_method_uid,type:bigint"`
	CompanyId               string `bun:"company_id,type:varchar(8)"`
	CustomerId              int32  `bun:"customer_id,type:int,nullzero"`
	ShipToId                int32  `bun:"ship_to_id,type:int,nullzero"`
}
