package prophet

import "github.com/uptrace/bun"

type P21PriceEngineHierarchy struct {
	bun.BaseModel              `bun:"table:p21_price_engine_hierarchy"`
	P21PriceEngineHierarchyUid int32   `bun:"p21_price_engine_hierarchy_uid,type:int,autoincrement,identity,pk"` // Unique identifier for this record.
	Name                       *string `bun:"name,type:varchar(255)"`                                            // Name of the pricing method.
	HierarchyOrder             *int32  `bun:"hierarchy_order,type:int"`                                          // The sequence number of the pricing method.
	ClrClassName               *string `bun:"clr_class_name,type:varchar(255)"`                                  // The Class name that is used for this pricing method in the CLR assembly.
}
