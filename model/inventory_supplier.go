package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InventorySupplier struct {
	bun.BaseModel `bun:"table:inventory_supplier"`

	InvMastUid           int32           `bun:"inv_mast_uid,type:int"`
	SupplierId           float64         `bun:"supplier_id,type:decimal(19)"`
	DivisionId           float64         `bun:"division_id,type:decimal(19)"`
	LeadTimeDays         float64         `bun:"lead_time_days,type:decimal(4)"`
	UpcCode              sql.NullString  `bun:"upc_code,type:varchar(14)"`
	CheckDigit           sql.NullFloat64 `bun:"check_digit,type:decimal(1)"`
	DeleteFlag           string          `bun:"delete_flag,type:char"`
	DateCreated          time.Time       `bun:"date_created,type:datetime"`
	DateLastModified     time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string          `bun:"last_maintained_by,type:varchar(30)"`
	SupplierPartNo       sql.NullString  `bun:"supplier_part_no,type:varchar(40)"`
	ListPrice            float64         `bun:"list_price,type:decimal(19, 9)"`
	Cost                 float64         `bun:"cost,type:decimal(19, 9)"`
	BackhaulType         string          `bun:"backhaul_type,type:char"`
	InventorySupplierUid int32           `bun:"inventory_supplier_uid,type:int,pk"`

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
