package prophet

import "github.com/uptrace/bun"

type PoSystemParameters struct {
	bun.BaseModel        `bun:"table:po_system_parameters"`
	RoundingFactor       float64 `bun:"rounding_factor,type:decimal(19,4)"`
	PricingUomIsOrderUom *string `bun:"pricing_uom_is_order_uom,type:char(1)"`
	PoSystemParameterUid string  `bun:"po_system_parameter_uid,type:numeric(19,0),pk"`
}
