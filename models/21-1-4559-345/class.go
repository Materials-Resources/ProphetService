package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Class struct {
	bun.BaseModel           `bun:"table:class"`
	ClassType               string    `bun:"class_type,type:varchar(2),pk"`                             // The type of the class.
	ClassNumber             float64   `bun:"class_number,type:decimal(3,0),pk"`                         // The number of the class.
	ClassId                 string    `bun:"class_id,type:varchar(255),pk"`                             // Identifier for the class.
	ClassDescription        string    `bun:"class_description,type:varchar(255)"`                       // The description of the class.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LogoPathFilename        string    `bun:"logo_path_filename,type:varchar(255),nullzero"`             // Logo file associated with a particular class
	ExportClassFlag         string    `bun:"export_class_flag,type:char(1),default:('N')"`              // Indicates whether this is an export class and will use an item's export description or not.
	FuelSurchargePercentage float64   `bun:"fuel_surcharge_percentage,type:decimal(19,4),nullzero"`     // A custom column which indicates the percentage that will be applied against a shipment when a shipment is confirmed.
	MaxFuelChargePerShip    float64   `bun:"max_fuel_charge_per_ship,type:decimal(19,4),nullzero"`      // A custom column which indicates the maximum fuel surcharge that will be applied against a shipment when a shipment is confirmed.
	ClassUid                int32     `bun:"class_uid,type:int"`
	AffinityFlag            string    `bun:"affinity_flag,type:char(1),nullzero"`                   // Indicates the Affinity for a specific Class 2. Possible values: A - Z
	HarmonizedCode          string    `bun:"harmonized_code,type:varchar(13),nullzero"`             // Classification number used to identify a type of item(s).
	AvailForCycleCountFlag  string    `bun:"avail_for_cycle_count_flag,type:char(1),default:('N')"` // Determines whether or not an item with this class will be available for cycle count.
}
