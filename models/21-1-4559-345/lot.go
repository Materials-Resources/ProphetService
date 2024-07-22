package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Lot struct {
	bun.BaseModel               `bun:"table:lot"`
	CompanyId                   string     `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	LocationId                  float64    `bun:"location_id,type:decimal(19,0),pk"`                             // Where was the used material located?
	Lot                         string     `bun:"lot,type:varchar(40),pk"`                                       // Lot number
	DeleteFlag                  string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	QtyOnHand                   *float64   `bun:"qty_on_hand,type:decimal(19,9)"`                                // The lot quantity on-hand.
	QtyAllocated                *float64   `bun:"qty_allocated,type:decimal(19,9)"`                              // The lot quantity allocated.
	Comment                     *string    `bun:"comment,type:varchar(255)"`                                     // Comments about the lot.
	CreationDate                time.Time  `bun:"creation_date,type:datetime,default:(getdate())"`               // Date lot was created.
	SkuCost                     float64    `bun:"sku_cost,type:decimal(19,9),default:(0.0000)"`                  // Lot cost.
	InvMastUid                  int32      `bun:"inv_mast_uid,type:int,pk"`                                      // Unique identifier for the item id.
	UserDefined1                *string    `bun:"user_defined1,type:varchar(30)"`                                // User defined column.
	UserDefined2                *string    `bun:"user_defined2,type:varchar(30)"`                                // User defined column.
	SupplierId                  *float64   `bun:"supplier_id,type:decimal(19,0)"`
	LotUid                      int32      `bun:"lot_uid,type:int,unique"`
	NoOfLots                    int32      `bun:"no_of_lots,type:int,default:(1)"`                        // indicate the number of manufactores lots included whin the receipt lot
	CertificationDate           *time.Time `bun:"certification_date,type:datetime"`                       // lot certification date
	SupplierCertificationDate   *time.Time `bun:"supplier_certification_date,type:datetime"`              // supplier certification date
	RecertificationDate         *time.Time `bun:"recertification_date,type:datetime"`                     // re-certification date
	SupplierRecertificationDate *time.Time `bun:"supplier_recertification_date,type:datetime"`            // supplier re-certification date
	QualityCdUid                *int32     `bun:"quality_cd_uid,type:int"`                                // Uniquely identify the quality code
	ExpirationDate              *time.Time `bun:"expiration_date,type:datetime"`                          // indicate lot expiration date
	ShelfLifeOriginDate         *time.Time `bun:"shelf_life_origin_date,type:datetime"`                   // Actual date shelf life began
	ShelfLifeExpDate            *time.Time `bun:"shelf_life_exp_date,type:datetime"`                      // Shelf life expiration date: the date the product becomes unsuitable to sell
	ShelfLifeCalcExpDate        *time.Time `bun:"shelf_life_calc_exp_date,type:datetime"`                 // Shelf life calculated expiration date: expiration date calculated from origin date and supplier/item settings.
	RetrievedByWms              string     `bun:"retrieved_by_wms,type:char(1),default:('N')"`            // Indicates if the lot has been retrieved by wms
	ItemRevisionUid             *int32     `bun:"item_revision_uid,type:int"`                             // holds information about item/revision.
	DimensionKey                *int32     `bun:"dimension_key,type:int"`                                 // Currently used for belting functionality only. Holds the key to the dimension record for this lot. FK to dimensions.
	BeltRemnantFlag             string     `bun:"belt_remnant_flag,type:char(1),default:('N')"`           // Indicates if this lot is a belt remnant
	SkuWeight                   *float64   `bun:"sku_weight,type:decimal(19,9)"`                          // The weight per SKU in the lot.
	BeltPurchasingFlag          *string    `bun:"belt_purchasing_flag,type:char(1),default:('N')"`        // Flag to determine if a lot doesn't meet the purchasing requirements
	BeltPurchasingLengthFlag    *string    `bun:"belt_purchasing_length_flag,type:char(1),default:('N')"` // Flag to determine if the lot doesn't meet the purchasing length
	BeltPurchasingWidthFlag     *string    `bun:"belt_purchasing_width_flag,type:char(1),default:('N')"`  // Flag to determine if the lot doesn't meet the purchasing width
	FullRollUnitQty             *float64   `bun:"full_roll_unit_qty,type:decimal(19,9)"`                  // Custom: Represents the unit quantity of a full (uncut) roll of the corresponding lot as it was received into P21.
	FullRollUnitSize            *float64   `bun:"full_roll_unit_size,type:decimal(19,4)"`                 // Custom: Represents the unit size of the full (uncut) roll that was received into P21.
	FullRollUom                 *string    `bun:"full_roll_uom,type:varchar(255)"`                        // Custom: Represents the unit of measure of the full (uncut) roll that was received into P21.
	SwisslogItemId              *string    `bun:"swisslog_item_id,type:varchar(255)"`                     // The Item ID that was used to induct this lot into Swisslog. This Item ID will be used when generating pick requests to take it out of Swisslog.
	PricingMultiplier           *float64   `bun:"pricing_multiplier,type:decimal(19,9),default:((1.00))"` // Sales pricing multiplier that will be applied to items from the lot.
	ParentLotUid                *int32     `bun:"parent_lot_uid,type:int"`                                // For lots associated with belt items, link to the lot that this lot was cut from.
	MaxAvailableDimKey          *int32     `bun:"max_available_dim_key,type:int"`                         // For belting functionality.  FK to dimensions.dimension_key.  Estimated maximum cut dimensions available on this belt.
	StepCutFlag                 string     `bun:"step_cut_flag,type:char(1),default:('N')"`               // For belting functionality.  Determines if the maximum dimension area (as defined by column max_available_dim_key) contains cuts within it.
	AtomicLotLength             *float64   `bun:"atomic_lot_length,type:decimal(19,9)"`                   // The length of the atomic lot item.
	AtomicLotHeight             *float64   `bun:"atomic_lot_height,type:decimal(19,9)"`                   // The width of the atomic lot item .
	AtomicLotDimensionScale     *string    `bun:"atomic_lot_dimension_scale,type:varchar(2)"`             // The dimension scale (ft, in, m, cm) that the length and width are defined in.
	AtomicLotUsableArea         *string    `bun:"atomic_lot_usable_area,type:varchar(255)"`               // Comment field.  The usable area of the atomic lot.
	AtomicLotBaseId             *string    `bun:"atomic_lot_base_id,type:varchar(35)"`                    // The base lot ID or prefix from which a sequence of atomic lot items was generated.
}
