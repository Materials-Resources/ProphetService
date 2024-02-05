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
	LocationId                 float64         `bun:"location_id,type:decimal(19,0)"`
	PrimarySupplier            string          `bun:"primary_supplier,type:char,default:('N')"`
	AverageLeadTime            int16           `bun:"average_lead_time,type:smallint,default:(0)"`
	RowStatusFlag              int32           `bun:"row_status_flag,type:int"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LocListPrice               sql.NullFloat64 `bun:"loc_list_price,type:decimal(19,9),default:(0)"`
	LocCost                    sql.NullFloat64 `bun:"loc_cost,type:decimal(19,9),default:(0)"`
	LocContractNumber          sql.NullString  `bun:"loc_contract_number,type:varchar(40),nullzero"`
	FixedLeadTime              sql.NullInt16   `bun:"fixed_lead_time,type:smallint,default:(0)"`
	DefaultDisposition         sql.NullString  `bun:"default_disposition,type:char,nullzero"`
	VmiStatus                  sql.NullInt32   `bun:"vmi_status,type:int,nullzero"`
	CreatedBy                  sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OverrideBegDate            sql.NullTime    `bun:"override_beg_date,type:datetime,nullzero"`
	OverrideEndDate            sql.NullTime    `bun:"override_end_date,type:datetime,nullzero"`
	VmiLastSendDate            sql.NullTime    `bun:"vmi_last_send_date,type:datetime,nullzero"`
	OverrideVmiStatus          sql.NullInt32   `bun:"override_vmi_status,type:int,nullzero"`
	OverrideVmiStatusFlag      string          `bun:"override_vmi_status_flag,type:char,default:('N')"`
	KeyVmiIndicatorChangedFlag string          `bun:"key_vmi_indicator_changed_flag,type:char,default:('N')"`
	FutureCost                 sql.NullFloat64 `bun:"future_cost,type:decimal(19,9),nullzero"`
	EffectiveDate              sql.NullTime    `bun:"effective_date,type:datetime,nullzero"`
	ManualLeadTime             sql.NullInt32   `bun:"manual_lead_time,type:int,nullzero"`
	FutureListPrice            sql.NullFloat64 `bun:"future_list_price,type:decimal(19,9),nullzero"`

	InventorySupplier InventorySupplier `bun:"rel:has-one,join:inventory_supplier_uid=inventory_supplier_uid"`
}
