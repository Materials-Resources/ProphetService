package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierXLoc struct {
	bun.BaseModel              `bun:"table:inventory_supplier_x_loc"`
	InventorySupplierXLocUid   int32     `bun:"inventory_supplier_x_loc_uid,type:int,pk"`
	InventorySupplierUid       int32     `bun:"inventory_supplier_uid,type:int"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`
	PrimarySupplier            string    `bun:"primary_supplier,type:char,default:('N')"`
	AverageLeadTime            int16     `bun:"average_lead_time,type:smallint,default:(0)"`
	RowStatusFlag              int32     `bun:"row_status_flag,type:int"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LocListPrice               float64   `bun:"loc_list_price,type:decimal(19,9),default:(0)"`
	LocCost                    float64   `bun:"loc_cost,type:decimal(19,9),default:(0)"`
	LocContractNumber          string    `bun:"loc_contract_number,type:varchar(40),nullzero"`
	FixedLeadTime              int16     `bun:"fixed_lead_time,type:smallint,default:(0)"`
	DefaultDisposition         string    `bun:"default_disposition,type:char,nullzero"`
	VmiStatus                  int32     `bun:"vmi_status,type:int,nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OverrideBegDate            time.Time `bun:"override_beg_date,type:datetime,nullzero"`
	OverrideEndDate            time.Time `bun:"override_end_date,type:datetime,nullzero"`
	VmiLastSendDate            time.Time `bun:"vmi_last_send_date,type:datetime,nullzero"`
	OverrideVmiStatus          int32     `bun:"override_vmi_status,type:int,nullzero"`
	OverrideVmiStatusFlag      string    `bun:"override_vmi_status_flag,type:char,default:('N')"`
	KeyVmiIndicatorChangedFlag string    `bun:"key_vmi_indicator_changed_flag,type:char,default:('N')"`
	FutureCost                 float64   `bun:"future_cost,type:decimal(19,9),nullzero"`
	EffectiveDate              time.Time `bun:"effective_date,type:datetime,nullzero"`
	ManualLeadTime             int32     `bun:"manual_lead_time,type:int,nullzero"`
	FutureListPrice            float64   `bun:"future_list_price,type:decimal(19,9),nullzero"`
}
