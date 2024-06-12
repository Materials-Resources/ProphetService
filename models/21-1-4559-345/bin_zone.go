package model

import (
	"github.com/uptrace/bun"
	"time"
)

type BinZone struct {
	bun.BaseModel            `bun:"table:bin_zone"`
	BinZoneUid               int32     `bun:"bin_zone_uid,type:int,pk"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	BinZoneId                string    `bun:"bin_zone_id,type:varchar(255)"`
	ZoneDesc                 string    `bun:"zone_desc,type:varchar(255)"`
	ZoneSequence             int32     `bun:"zone_sequence,type:int"`
	ZoneTypeCd               int32     `bun:"zone_type_cd,type:int"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	BinReplenishmentInterval int32     `bun:"bin_replenishment_interval,type:int,nullzero"`
	BinReplenishmentUser     string    `bun:"bin_replenishment_user,type:varchar(30),nullzero"`
	BypassFullQtyFlag        string    `bun:"bypass_full_qty_flag,type:char,default:('N')"`
	AiaPointValue            float64   `bun:"aia_point_value,type:decimal(19,9),nullzero"`
	AllocationUom            string    `bun:"allocation_uom,type:varchar(8),nullzero"`
	PrinterName              string    `bun:"printer_name,type:varchar(255),nullzero"`
	TagItemIdValidationFlag  string    `bun:"tag_item_id_validation_flag,type:char,nullzero"`
	DayPickZoneFlag          string    `bun:"day_pick_zone_flag,type:char,nullzero"`
}
