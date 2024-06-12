package model

import "github.com/uptrace/bun"

type P21PriceEngineHierarchy struct {
	bun.BaseModel              `bun:"table:p21_price_engine_hierarchy"`
	P21PriceEngineHierarchyUid int32  `bun:"p21_price_engine_hierarchy_uid,type:int,pk,identity"`
	Name                       string `bun:"name,type:varchar(255),nullzero"`
	HierarchyOrder             int32  `bun:"hierarchy_order,type:int,nullzero"`
	ClrClassName               string `bun:"clr_class_name,type:varchar(255),nullzero"`
}
