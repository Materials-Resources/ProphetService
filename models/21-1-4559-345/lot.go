package model

type Lot struct {
	bun.BaseModel               `bun:"table:lot"`
	CompanyId                   string    `bun:"company_id,type:varchar(8),pk"`
	LocationId                  float64   `bun:"location_id,type:decimal(19,0),pk"`
	Lot                         string    `bun:"lot,type:varchar(40),pk"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	QtyOnHand                   float64   `bun:"qty_on_hand,type:decimal(19,9),nullzero"`
	QtyAllocated                float64   `bun:"qty_allocated,type:decimal(19,9),nullzero"`
	Comment                     string    `bun:"comment,type:varchar(255),nullzero"`
	CreationDate                time.Time `bun:"creation_date,type:datetime,default:(getdate())"`
	SkuCost                     float64   `bun:"sku_cost,type:decimal(19,9),default:(0.0000)"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,pk"`
	UserDefined1                string    `bun:"user_defined1,type:varchar(30),nullzero"`
	UserDefined2                string    `bun:"user_defined2,type:varchar(30),nullzero"`
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	LotUid                      int32     `bun:"lot_uid,type:int"`
	NoOfLots                    int32     `bun:"no_of_lots,type:int,default:(1)"`
	CertificationDate           time.Time `bun:"certification_date,type:datetime,nullzero"`
	SupplierCertificationDate   time.Time `bun:"supplier_certification_date,type:datetime,nullzero"`
	RecertificationDate         time.Time `bun:"recertification_date,type:datetime,nullzero"`
	SupplierRecertificationDate time.Time `bun:"supplier_recertification_date,type:datetime,nullzero"`
	QualityCdUid                int32     `bun:"quality_cd_uid,type:int,nullzero"`
	ExpirationDate              time.Time `bun:"expiration_date,type:datetime,nullzero"`
	ShelfLifeOriginDate         time.Time `bun:"shelf_life_origin_date,type:datetime,nullzero"`
	ShelfLifeExpDate            time.Time `bun:"shelf_life_exp_date,type:datetime,nullzero"`
	ShelfLifeCalcExpDate        time.Time `bun:"shelf_life_calc_exp_date,type:datetime,nullzero"`
	RetrievedByWms              string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	ItemRevisionUid             int32     `bun:"item_revision_uid,type:int,nullzero"`
	DimensionKey                int32     `bun:"dimension_key,type:int,nullzero"`
	BeltRemnantFlag             string    `bun:"belt_remnant_flag,type:char,default:('N')"`
	SkuWeight                   float64   `bun:"sku_weight,type:decimal(19,9),nullzero"`
	BeltPurchasingFlag          string    `bun:"belt_purchasing_flag,type:char,default:('N')"`
	BeltPurchasingLengthFlag    string    `bun:"belt_purchasing_length_flag,type:char,default:('N')"`
	BeltPurchasingWidthFlag     string    `bun:"belt_purchasing_width_flag,type:char,default:('N')"`
	FullRollUnitQty             float64   `bun:"full_roll_unit_qty,type:decimal(19,9),nullzero"`
	FullRollUnitSize            float64   `bun:"full_roll_unit_size,type:decimal(19,4),nullzero"`
	FullRollUom                 string    `bun:"full_roll_uom,type:varchar(255),nullzero"`
	SwisslogItemId              string    `bun:"swisslog_item_id,type:varchar(255),nullzero"`
	PricingMultiplier           float64   `bun:"pricing_multiplier,type:decimal(19,9),default:((1.00))"`
	ParentLotUid                int32     `bun:"parent_lot_uid,type:int,nullzero"`
	MaxAvailableDimKey          int32     `bun:"max_available_dim_key,type:int,nullzero"`
	StepCutFlag                 string    `bun:"step_cut_flag,type:char,default:('N')"`
	AtomicLotLength             float64   `bun:"atomic_lot_length,type:decimal(19,9),nullzero"`
	AtomicLotHeight             float64   `bun:"atomic_lot_height,type:decimal(19,9),nullzero"`
	AtomicLotDimensionScale     string    `bun:"atomic_lot_dimension_scale,type:varchar(2),nullzero"`
	AtomicLotUsableArea         string    `bun:"atomic_lot_usable_area,type:varchar(255),nullzero"`
	AtomicLotBaseId             string    `bun:"atomic_lot_base_id,type:varchar(35),nullzero"`
}
