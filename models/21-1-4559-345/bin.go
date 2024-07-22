package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Bin struct {
	bun.BaseModel        `bun:"table:bin"`
	CompanyId            string     `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	LocationId           float64    `bun:"location_id,type:decimal(19,0),pk"`                         // Where was the used material located?
	BinId                string     `bun:"bin_id,type:varchar(10),pk"`                                // Enter a bin ID
	DeleteFlag           string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated          time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DateLastCounted      *time.Time `bun:"date_last_counted,type:datetime"`                           // When was the inventory at this location last counted?
	CreatedBy            string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // user who created bin record
	BinTypeUid           *int32     `bun:"bin_type_uid,type:int"`                                     // Foreign key to the bin_type table
	PickLockedFlag       string     `bun:"pick_locked_flag,type:char(1),default:('N')"`               // Indicate whether user can pick quantity from the bin
	PutLockedFlag        string     `bun:"put_locked_flag,type:char(1),default:('N')"`                // Indicate whether user can put quantity in the bin
	FullFlag             string     `bun:"full_flag,type:char(1),default:('N')"`                      // Indicate whether bin is full
	FrozenFlag           string     `bun:"frozen_flag,type:char(1),default:('N')"`                    // Indicate whether bin is frozen for counting
	MaxWeight            float64    `bun:"max_weight,type:decimal(19,9),default:(0)"`                 // Maximum weight the bin can hold
	CurrentWeight        float64    `bun:"current_weight,type:decimal(19,9),default:(0)"`             // Current weight in the bin
	CurrentVolume        float64    `bun:"current_volume,type:decimal(19,9),default:(0)"`             // Current volume in the bin
	PutawayZoneUid       *int32     `bun:"putaway_zone_uid,type:int"`                                 // foreing key to bin_zone_uid for putaway zone
	PutawayZoneSequence  int32      `bun:"putaway_zone_sequence,type:int,default:(0)"`                // Sequence bin has in putaway zone
	PickZoneUid          *int32     `bun:"pick_zone_uid,type:int"`                                    // Foreign key - bin_zone_uid for pick zone
	PickZoneSequence     int32      `bun:"pick_zone_sequence,type:int,default:(0)"`                   // Sequence bin has in pick zone
	WarehouseSequence    int32      `bun:"warehouse_sequence,type:int,default:(0)"`                   // Sequence used if no putaway/pick sequence specified
	BinLength            float64    `bun:"bin_length,type:decimal(19,9),default:(0)"`                 // Length of bin
	BinHeight            float64    `bun:"bin_height,type:decimal(19,9),default:(0)"`                 // Height of bin
	BinWidth             float64    `bun:"bin_width,type:decimal(19,9),default:(0)"`                  // Width of bin
	BinUid               int32      `bun:"bin_uid,type:int,unique"`                                   // Uniquer row ID for the table.
	RfBinFlag            string     `bun:"rf_bin_flag,type:char(1),default:('N')"`                    // These bins are RF devices, and hold what the RF scans.
	ConsolidationBinFlag string     `bun:"consolidation_bin_flag,type:char(1),default:('N')"`         // These bins are used to assemble orders, from the RF Group Order Picking process.
	RfTerminalUid        *int32     `bun:"rf_terminal_uid,type:int"`                                  // Pointer to the rf_terminal record that the bin record was created for
	MaxUniqueItems       *int32     `bun:"max_unique_items,type:int,default:(0)"`                     // The maximum number of unique item_ids that can be stored in this bin.
	StageBinFlag         string     `bun:"stage_bin_flag,type:char(1),default:('N')"`                 // Indicates that this bin is a staging area bin. This bin will be linked to a truck door bin.
	DoorBinFlag          string     `bun:"door_bin_flag,type:char(1),default:('N')"`                  // Indicates that this bin is a truck door bin.
	StageBinXDoorBinUid  *int32     `bun:"stage_bin_x_door_bin_uid,type:int"`                         // Bin that will indicate this is a truck door bin for the stage bin.
	RmaBinFlag           *string    `bun:"rma_bin_flag,type:char(1),default:('N')"`                   // Default RMA Bin for a Location
	MasterBinFlag        string     `bun:"master_bin_flag,type:char(1),default:('Y')"`                // Indicates if the bin is a master bin.
	MasterBinUid         *int32     `bun:"master_bin_uid,type:int"`                                   // Indicates the master bin for this bin.
	PickConfirmedFlag    string     `bun:"pick_confirmed_flag,type:char(1),default:('N')"`            // Custom (F53615): determines if a bin can be employed as a picking confirmed bin for shipping purposes.
	DateLastValidated    *time.Time `bun:"date_last_validated,type:datetime"`                         // Last time BIN was validated
	LastValidatedBy      *string    `bun:"last_validated_by,type:varchar(30)"`                        // who validated it last time?
	AllowedAgeRangeWeeks *int32     `bun:"allowed_age_range_weeks,type:int"`                          // Custom F67642 - Indicates the range of manufacture dates (in weeks) allowed for a given item to be deposited into this bin.
}
