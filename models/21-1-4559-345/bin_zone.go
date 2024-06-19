package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BinZone struct {
	bun.BaseModel            `bun:"table:bin_zone"`
	BinZoneUid               int32     `bun:"bin_zone_uid,type:int,pk"`                                     // Unique identifier for table
	LocationId               float64   `bun:"location_id,type:decimal(19,0),unique"`                        // Location
	BinZoneId                string    `bun:"bin_zone_id,type:varchar(255),unique"`                         // Name of zone
	ZoneDesc                 string    `bun:"zone_desc,type:varchar(255)"`                                  // Zone description
	ZoneSequence             int32     `bun:"zone_sequence,type:int"`                                       // The order in which zones are used to put away/pick stock
	ZoneTypeCd               int32     `bun:"zone_type_cd,type:int,unique"`                                 // Indicate whether this is a Put Zone or a Pick Zone
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                     // Indicate whether record is deleted or not
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	BinReplenishmentInterval int32     `bun:"bin_replenishment_interval,type:int,nullzero"`                 // Time to elapse (in minutes) before trying to send an alert or notify all Wireless users.
	BinReplenishmentUser     string    `bun:"bin_replenishment_user,type:varchar(30),nullzero"`             // User associated with doing bin replenishment for this zone.
	BypassFullQtyFlag        string    `bun:"bypass_full_qty_flag,type:char(1),default:('N')"`              // Exclude zone from full line/package allocations during pick ticket printing.
	AiaPointValue            float64   `bun:"aia_point_value,type:decimal(19,9),nullzero"`                  // The Advanced Inventory Allocation point value assigned to this bin zone
	AllocationUom            string    `bun:"allocation_uom,type:varchar(8),nullzero"`                      // Used if you are using alternate bin allocation strategy to determine order of bin allocation.
	PrinterName              string    `bun:"printer_name,type:varchar(255),nullzero"`                      // Custom column to specify printer path per Pick Zone/Location
	TagItemIdValidationFlag  string    `bun:"tag_item_id_validation_flag,type:char(1),nullzero"`            // Choose if this zone should perform item validations.
	DayPickZoneFlag          string    `bun:"day_pick_zone_flag,type:char(1),nullzero"`                     // Flag column to mark the zone and treat it as a day pick zone in WWMS
}
