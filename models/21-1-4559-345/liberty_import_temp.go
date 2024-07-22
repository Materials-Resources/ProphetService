package prophet

import "github.com/uptrace/bun"

type LibertyImportTemp struct {
	bun.BaseModel `bun:"table:liberty_import_temp"`
	SupplierPart  *string  `bun:"Supplier Part,type:varchar(max)"`
	ListPrice     *float64 `bun:"List Price,type:float"`
	PurchaseCost  *float64 `bun:"Purchase Cost,type:float"`
	Uom           *int32   `bun:"UOM,type:int"`
	CasePack      *string  `bun:"CASE PACK,type:varchar(max)"`
	ProductGroup  *string  `bun:"Product Group,type:varchar(max)"`
	Supplier      *int32   `bun:"Supplier,type:int"`
}
