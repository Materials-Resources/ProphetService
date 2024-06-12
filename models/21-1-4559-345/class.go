package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Class struct {
	bun.BaseModel           `bun:"table:class"`
	ClassType               string    `bun:"class_type,type:varchar(2),pk"`
	ClassNumber             float64   `bun:"class_number,type:decimal(3,0),pk"`
	ClassId                 string    `bun:"class_id,type:varchar(255),pk"`
	ClassDescription        string    `bun:"class_description,type:varchar(255)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LogoPathFilename        string    `bun:"logo_path_filename,type:varchar(255),nullzero"`
	ExportClassFlag         string    `bun:"export_class_flag,type:char,default:('N')"`
	FuelSurchargePercentage float64   `bun:"fuel_surcharge_percentage,type:decimal(19,4),nullzero"`
	MaxFuelChargePerShip    float64   `bun:"max_fuel_charge_per_ship,type:decimal(19,4),nullzero"`
	ClassUid                int32     `bun:"class_uid,type:int"`
	AffinityFlag            string    `bun:"affinity_flag,type:char,nullzero"`
	HarmonizedCode          string    `bun:"harmonized_code,type:varchar(13),nullzero"`
	AvailForCycleCountFlag  string    `bun:"avail_for_cycle_count_flag,type:char,default:('N')"`
}
