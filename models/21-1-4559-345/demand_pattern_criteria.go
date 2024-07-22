package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandPatternCriteria struct {
	bun.BaseModel             `bun:"table:demand_pattern_criteria"`
	DemandPatternCriteriaUid  int32     `bun:"demand_pattern_criteria_uid,type:int,pk"`                      // Unique ID for the demand pattern criteria
	DemandPatternCriteriaId   string    `bun:"demand_pattern_criteria_id,type:varchar(255),unique"`          // User defined key which uniquely identifies a criteria.
	CompanyId                 string    `bun:"company_id,type:varchar(255)"`                                 // Company to run demand pattern identification for.
	Description               string    `bun:"description,type:varchar(255)"`                                // User specified description for the criteria
	FromLocationId            float64   `bun:"from_location_id,type:decimal(19,0)"`                          // Used to limit items by location_id
	ToLocationId              float64   `bun:"to_location_id,type:decimal(19,0)"`                            // Used to limit items by location_id
	FromItemId                string    `bun:"from_item_id,type:varchar(255)"`                               // Used to limit items by item_id
	ToItemId                  string    `bun:"to_item_id,type:varchar(255)"`                                 // Used to limit items by item_id
	FromProductGroupId        string    `bun:"from_product_group_id,type:varchar(255)"`                      // Used to limit items by product group
	ToProductGroupId          string    `bun:"to_product_group_id,type:varchar(255)"`                        // Used to limit items by product group
	FromSupplierId            float64   `bun:"from_supplier_id,type:decimal(19,0)"`                          // Used to limit items by supplier
	ToSupplierId              float64   `bun:"to_supplier_id,type:decimal(19,0)"`                            // Used to limit items by supplier
	FromPurchaseClassId       string    `bun:"from_purchase_class_id,type:varchar(255)"`                     // Used to limit items by ABC class
	ToPurchaseClassId         string    `bun:"to_purchase_class_id,type:varchar(255)"`                       // Used to limit items by ABC class
	ConsiderLevelFlag         string    `bun:"consider_level_flag,type:char(1)"`                             // Whether or not to include Level items in processing
	ConsiderTrendFlag         string    `bun:"consider_trend_flag,type:char(1)"`                             // Whether or not to include Trend items in processing
	ConsiderSeasonalFlag      string    `bun:"consider_seasonal_flag,type:char(1)"`                          // Whether or not to include Seasonal items in processing
	ConsiderSeasonalTrendFlag string    `bun:"consider_seasonal_trend_flag,type:char(1)"`                    // Whether or not to include Seasonal items with Trend in processing
	ConsiderSlowFlag          string    `bun:"consider_slow_flag,type:char(1)"`                              // Whether or not to include Slow moving items in processing
	ConsiderErraticFlag       string    `bun:"consider_erratic_flag,type:char(1)"`                           // Whether or not to include Erratic items in processing
	LastEvaluationDate        time.Time `bun:"last_evaluation_date,type:datetime"`                           // Do not include items which have had their demand pattern identified since this date.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	WarnIfTrendExceedsPercent float64   `bun:"warn_if_trend_exceeds_percent,type:decimal(19,9)"`             // Flag items identified as trend items when their average trend percent change is greater than this value.
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Current status of this criteria.
	BuyerId                   *string   `bun:"buyer_id,type:varchar(16)"`                                    // Custom F68622: Optional buyer ID criteria to limit items by buyer.
	ConsiderSporadicFlag      string    `bun:"consider_sporadic_flag,type:char(1),default:('Y')"`            // Sporadic Demand Flag
}
