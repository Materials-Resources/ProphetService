package gen

import "github.com/uptrace/bun"

type SalesPricingOption struct {
	bun.BaseModel `bun:"table:sales_pricing_option"`
	PricingOption string `bun:"pricing_option,type:varchar(255)"`
	ConfigId      int32  `bun:"config_id,type:int"`
}
