package model

type Bin struct {
	bun.BaseModel        `bun:"table:bin"`
	CompanyId            string    `bun:"company_id,type:varchar(8),pk"`
	LocationId           float64   `bun:"location_id,type:decimal(19,0),pk"`
	BinId                string    `bun:"bin_id,type:varchar(10),pk"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DateLastCounted      time.Time `bun:"date_last_counted,type:datetime,nullzero"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	BinTypeUid           int32     `bun:"bin_type_uid,type:int,nullzero"`
	PickLockedFlag       string    `bun:"pick_locked_flag,type:char,default:('N')"`
	PutLockedFlag        string    `bun:"put_locked_flag,type:char,default:('N')"`
	FullFlag             string    `bun:"full_flag,type:char,default:('N')"`
	FrozenFlag           string    `bun:"frozen_flag,type:char,default:('N')"`
	MaxWeight            float64   `bun:"max_weight,type:decimal(19,9),default:(0)"`
	CurrentWeight        float64   `bun:"current_weight,type:decimal(19,9),default:(0)"`
	CurrentVolume        float64   `bun:"current_volume,type:decimal(19,9),default:(0)"`
	PutawayZoneUid       int32     `bun:"putaway_zone_uid,type:int,nullzero"`
	PutawayZoneSequence  int32     `bun:"putaway_zone_sequence,type:int,default:(0)"`
	PickZoneUid          int32     `bun:"pick_zone_uid,type:int,nullzero"`
	PickZoneSequence     int32     `bun:"pick_zone_sequence,type:int,default:(0)"`
	WarehouseSequence    int32     `bun:"warehouse_sequence,type:int,default:(0)"`
	BinLength            float64   `bun:"bin_length,type:decimal(19,9),default:(0)"`
	BinHeight            float64   `bun:"bin_height,type:decimal(19,9),default:(0)"`
	BinWidth             float64   `bun:"bin_width,type:decimal(19,9),default:(0)"`
	BinUid               int32     `bun:"bin_uid,type:int"`
	RfBinFlag            string    `bun:"rf_bin_flag,type:char,default:('N')"`
	ConsolidationBinFlag string    `bun:"consolidation_bin_flag,type:char,default:('N')"`
	RfTerminalUid        int32     `bun:"rf_terminal_uid,type:int,nullzero"`
	MaxUniqueItems       int32     `bun:"max_unique_items,type:int,default:(0)"`
	StageBinFlag         string    `bun:"stage_bin_flag,type:char,default:('N')"`
	DoorBinFlag          string    `bun:"door_bin_flag,type:char,default:('N')"`
	StageBinXDoorBinUid  int32     `bun:"stage_bin_x_door_bin_uid,type:int,nullzero"`
	RmaBinFlag           string    `bun:"rma_bin_flag,type:char,default:('N')"`
	MasterBinFlag        string    `bun:"master_bin_flag,type:char,default:('Y')"`
	MasterBinUid         int32     `bun:"master_bin_uid,type:int,nullzero"`
	PickConfirmedFlag    string    `bun:"pick_confirmed_flag,type:char,default:('N')"`
	DateLastValidated    time.Time `bun:"date_last_validated,type:datetime,nullzero"`
	LastValidatedBy      string    `bun:"last_validated_by,type:varchar(30),nullzero"`
	AllowedAgeRangeWeeks int32     `bun:"allowed_age_range_weeks,type:int,nullzero"`
}
