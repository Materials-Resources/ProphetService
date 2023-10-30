package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvMast struct {
	bun.BaseModel `bun:"table:inv_mast"`

	InvMastUid                int32           `bun:"inv_mast_uid,type:int,pk"`
	ItemId                    string          `bun:"item_id,type:varchar(40),unique"`
	ItemDesc                  string          `bun:"item_desc,type:varchar(40)"`
	DeleteFlag                string          `bun:"delete_flag,type:char"`
	Weight                    sql.NullFloat64 `bun:"weight,type:decimal(19, 12)"`
	NetWeight                 sql.NullFloat64 `bun:"net_weight,type:decimal(19, 12)"`
	DateCreated               time.Time       `bun:"date_created,type:datetime"`
	DateLastModified          time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string          `bun:"last_maintained_by,type:varchar(30)"`
	Inactive                  string          `bun:"inactive,type:char"`
	CatchWeightIndicator      string          `bun:"catch_weight_indicator,type:char"`
	PurchasingWeight          sql.NullFloat64 `bun:"purchasing_weight,type:decimal(19, 12)"`
	Serialized                string          `bun:"serialized,type:char"`
	ProductType               string          `bun:"product_type,type:char"`
	Hose                      string          `bun:"hose,type:char"`
	Fitting                   string          `bun:"fitting,type:char"`
	CatalogItem               string          `bun:"catalog_item,type:char"`
	TrackLots                 string          `bun:"track_lots,type:char"`
	Requisition               string          `bun:"requisition,type:char,default:'N'"`
	Price1                    sql.NullFloat64 `bun:"price1,type:decimal(19, 9)"`
	Price2                    sql.NullFloat64 `bun:"price2,type:decimal(19, 9)"`
	Price3                    sql.NullFloat64 `bun:"price3,type:decimal(19, 9)"`
	Price4                    sql.NullFloat64 `bun:"price4,type:decimal(19, 9)"`
	Price5                    sql.NullFloat64 `bun:"price5,type:decimal(19, 9)"`
	Price6                    sql.NullFloat64 `bun:"price6,type:decimal(19, 9)"`
	Price7                    sql.NullFloat64 `bun:"price7,type:decimal(19, 9)"`
	Price8                    sql.NullFloat64 `bun:"price8,type:decimal(19, 9)"`
	Price9                    sql.NullFloat64 `bun:"price9,type:decimal(19, 9)"`
	Price10                   sql.NullFloat64 `bun:"price10,type:decimal(19, 9)"`
	DefaultSalesDiscountGroup sql.NullString  `bun:"default_sales_discount_group,type:varchar(8)"`
	DefaultPurchaseDiscGroup  sql.NullString  `bun:"default_purchase_disc_group,type:varchar(8)"`
	SalesPricingUnit          sql.NullString  `bun:"sales_pricing_unit,type:varchar(8)"`
	SalesPricingUnitSize      sql.NullFloat64 `bun:"sales_pricing_unit_size,type:decimal(19, 4)"`
	PurchasePricingUnit       sql.NullString  `bun:"purchase_pricing_unit,type:varchar(8)"`
	PurchasePricingUnitSize   sql.NullFloat64 `bun:"purchase_pricing_unit_size,type:decimal(19, 4)"`
	ExtendedDesc              sql.NullString  `bun:"extended_desc,type:varchar(255)"`
	DefaultSellingUnit        sql.NullString  `bun:"default_selling_unit,type:varchar(8)"`
	DefaultPurchasingUnit     sql.NullString  `bun:"default_purchasing_unit,type:varchar(8)"`
	BaseUnit                  string          `bun:"base_unit,type:varchar(8)"`

	InvLoc InvLoc `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

var _ bun.BeforeInsertHook = (*InvMast)(nil)

func (m *InvMast) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	m.DateCreated = time.Now()
	m.DateLastModified = time.Now()
	return nil
}

var _ bun.BeforeUpdateHook = (*InvMast)(nil)

func (m *InvMast) BeforeUpdate(ctx context.Context, query *bun.UpdateQuery) error {
	m.DateLastModified = time.Now()
	return nil
}
