package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InventorySupplierXLoc struct {
	bun.BaseModel `bun:"table:inventory_supplier_x_loc"`

	InventorySupplierXLocUid   int32           `bun:"inventory_supplier_x_loc_uid,type:int,pk"`
	InventorySupplierUid       int32           `bun:"inventory_supplier_uid,type:int"`
	LocationId                 float64         `bun:"location_id,type:decimal(19)"`
	PrimarySupplier            string          `bun:"primary_supplier,type:char,default:'N'"`
	AverageLeadTime            int32           `bun:"average_lead_time,type:smallint,default:0"`
	RowStatusFlag              int32           `bun:"row_status_flag,type:int"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30)"`
	LocListPrice               sql.NullFloat64 `bun:"loc_list_price,type:decimal(19, 9),default:0"`
	LocCost                    sql.NullFloat64 `bun:"loc_cost,type:decimal(19, 9),default:0"`
	FixedLeadTime              sql.NullInt32   `bun:"fixed_lead_time,type:smallint,default:0"`
	VmiStatus                  sql.NullInt32   `bun:"vmi_status,type:int"`
	CreatedBy                  sql.NullString  `bun:"created_by,type:varchar(255)"`
	OverrideVmiStatusFlag      string          `bun:"override_vmi_status_flag,type:char"`
	KeyVmiIndicatorChangedFlag string          `bun:"key_vmi_indicator_changed_flag,type:char"`

	InventorySupplier InventorySupplier `bun:"rel:has-one,"`
}
