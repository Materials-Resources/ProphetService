package gen

import "github.com/uptrace/bun"

type PriceMethod struct {
	bun.BaseModel  `bun:"table:price_method"`
	PriceMethodUid int64   `bun:"price_method_uid,type:bigint,autoincrement,identity"`
	PricePagePrice bool    `bun:"price_page_price,type:bit,default:((0))"`
	CustomerPrice  bool    `bun:"customer_price,type:bit,default:((0))"`
	ContractPrice  bool    `bun:"contract_price,type:bit,default:((0))"`
	PriceMethod    string  `bun:"price_method,type:varchar(255)"`
	SourcePriceCd  int32   `bun:"source_price_cd,type:int,nullzero"`
	Multiplier     float64 `bun:"multiplier,type:decimal(19,9),nullzero"`
	Price          float64 `bun:"price,type:decimal(19,9),nullzero"`
}
