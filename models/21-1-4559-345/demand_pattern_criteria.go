package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandPatternCriteria struct {
	bun.BaseModel             `bun:"table:demand_pattern_criteria"`
	DemandPatternCriteriaUid  int32     `bun:"demand_pattern_criteria_uid,type:int,pk"`
	DemandPatternCriteriaId   string    `bun:"demand_pattern_criteria_id,type:varchar(255)"`
	CompanyId                 string    `bun:"company_id,type:varchar(255)"`
	Description               string    `bun:"description,type:varchar(255)"`
	FromLocationId            float64   `bun:"from_location_id,type:decimal(19,0)"`
	ToLocationId              float64   `bun:"to_location_id,type:decimal(19,0)"`
	FromItemId                string    `bun:"from_item_id,type:varchar(255)"`
	ToItemId                  string    `bun:"to_item_id,type:varchar(255)"`
	FromProductGroupId        string    `bun:"from_product_group_id,type:varchar(255)"`
	ToProductGroupId          string    `bun:"to_product_group_id,type:varchar(255)"`
	FromSupplierId            float64   `bun:"from_supplier_id,type:decimal(19,0)"`
	ToSupplierId              float64   `bun:"to_supplier_id,type:decimal(19,0)"`
	FromPurchaseClassId       string    `bun:"from_purchase_class_id,type:varchar(255)"`
	ToPurchaseClassId         string    `bun:"to_purchase_class_id,type:varchar(255)"`
	ConsiderLevelFlag         string    `bun:"consider_level_flag,type:char"`
	ConsiderTrendFlag         string    `bun:"consider_trend_flag,type:char"`
	ConsiderSeasonalFlag      string    `bun:"consider_seasonal_flag,type:char"`
	ConsiderSeasonalTrendFlag string    `bun:"consider_seasonal_trend_flag,type:char"`
	ConsiderSlowFlag          string    `bun:"consider_slow_flag,type:char"`
	ConsiderErraticFlag       string    `bun:"consider_erratic_flag,type:char"`
	LastEvaluationDate        time.Time `bun:"last_evaluation_date,type:datetime"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	WarnIfTrendExceedsPercent float64   `bun:"warn_if_trend_exceeds_percent,type:decimal(19,9)"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:((704))"`
	BuyerId                   string    `bun:"buyer_id,type:varchar(16),nullzero"`
	ConsiderSporadicFlag      string    `bun:"consider_sporadic_flag,type:char,default:('Y')"`
}
