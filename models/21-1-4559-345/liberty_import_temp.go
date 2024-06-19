package gen

import "github.com/uptrace/bun"

type LibertyImportTemp struct {
	bun.BaseModel `bun:"table:liberty_import_temp"`
	SupplierPart  string  `bun:"Supplier Part,type:varchar(max),nullzero"`
	ListPrice     float64 `bun:"List Price,type:float,nullzero"`
	PurchaseCost  float64 `bun:"Purchase Cost,type:float,nullzero"`
	Uom           int32   `bun:"UOM,type:int,nullzero"`
	CasePack      string  `bun:"CASE PACK,type:varchar(max),nullzero"`
	ProductGroup  string  `bun:"Product Group,type:varchar(max),nullzero"`
	Supplier      int32   `bun:"Supplier,type:int,nullzero"`
}
