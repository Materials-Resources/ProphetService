package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun/schema"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

type InvLoc struct {
	bun.BaseModel                   `bun:"table:inv_loc"`
	LocationId                      float64         `bun:"location_id,type:decimal(19,0),pk"`
	QtyOnHand                       sql.NullFloat64 `bun:"qty_on_hand,type:decimal(19,9),default:(0)"`
	QtyInProcess                    float64         `bun:"qty_in_process,type:decimal(19,9),default:(0)"`
	DateCreated                     time.Time       `bun:"date_created,type:datetime"`
	DateLastModified                time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy                string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CompanyId                       string          `bun:"company_id,type:varchar(8),pk"`
	LastRecPo                       sql.NullFloat64 `bun:"last_rec_po,type:decimal(19,9),nullzero"`
	LastRecPoWithDisc               sql.NullFloat64 `bun:"last_rec_po_with_disc,type:decimal(19,9),nullzero"`
	GlAccountNo                     sql.NullString  `bun:"gl_account_no,type:varchar(32),nullzero"`
	PurchOrTransfer                 sql.NullString  `bun:"purch_or_transfer,type:char,nullzero"`
	NextDueInPoCost                 sql.NullFloat64 `bun:"next_due_in_po_cost,type:decimal(19,9),nullzero"`
	NextDueInPoDate                 sql.NullTime    `bun:"next_due_in_po_date,type:datetime,nullzero"`
	RevenueAccountNo                sql.NullString  `bun:"revenue_account_no,type:varchar(32),nullzero"`
	CosAccountNo                    sql.NullString  `bun:"cos_account_no,type:varchar(32),nullzero"`
	Sellable                        sql.NullString  `bun:"sellable,type:char,nullzero"`
	MovingAverageCost               sql.NullFloat64 `bun:"moving_average_cost,type:decimal(19,9),nullzero"`
	StandardCost                    sql.NullFloat64 `bun:"standard_cost,type:decimal(19,9),nullzero"`
	ProtectedStockQty               sql.NullFloat64 `bun:"protected_stock_qty,type:decimal(19,9),nullzero"`
	InvMin                          sql.NullFloat64 `bun:"inv_min,type:decimal(19,9),nullzero"`
	InvMax                          sql.NullFloat64 `bun:"inv_max,type:decimal(19,9),nullzero"`
	SafetyStock                     sql.NullFloat64 `bun:"safety_stock,type:decimal(19,9),nullzero"`
	Stockable                       sql.NullString  `bun:"stockable,type:char,nullzero"`
	ReplenishmentLocation           sql.NullFloat64 `bun:"replenishment_location,type:decimal(19,0),nullzero"`
	MonthsInSeason                  sql.NullFloat64 `bun:"months_in_season,type:decimal(2,0),nullzero"`
	AverageMonthlyUsage             sql.NullFloat64 `bun:"average_monthly_usage,type:decimal(19,9),nullzero"`
	ProductGroupId                  sql.NullString  `bun:"product_group_id,type:varchar(8),nullzero"`
	PurchaseClass                   sql.NullString  `bun:"purchase_class,type:varchar(8),nullzero"`
	CycleCountingCategory           sql.NullFloat64 `bun:"cycle_counting_category,type:decimal(19,0),nullzero"`
	PurchaseDiscountGroup           sql.NullString  `bun:"purchase_discount_group,type:varchar(8),nullzero"`
	SalesDiscountGroup              sql.NullString  `bun:"sales_discount_group,type:varchar(8),nullzero"`
	NoCharge                        sql.NullString  `bun:"no_charge,type:char,nullzero"`
	Price1                          sql.NullFloat64 `bun:"price1,type:decimal(19,9),nullzero"`
	Price2                          sql.NullFloat64 `bun:"price2,type:decimal(19,9),nullzero"`
	Price3                          sql.NullFloat64 `bun:"price3,type:decimal(19,9),nullzero"`
	Price4                          sql.NullFloat64 `bun:"price4,type:decimal(19,9),nullzero"`
	Price5                          sql.NullFloat64 `bun:"price5,type:decimal(19,9),nullzero"`
	Price6                          sql.NullFloat64 `bun:"price6,type:decimal(19,9),nullzero"`
	Price7                          sql.NullFloat64 `bun:"price7,type:decimal(19,9),nullzero"`
	Price8                          sql.NullFloat64 `bun:"price8,type:decimal(19,9),nullzero"`
	Price9                          sql.NullFloat64 `bun:"price9,type:decimal(19,9),nullzero"`
	Price10                         sql.NullFloat64 `bun:"price10,type:decimal(19,9),nullzero"`
	ReplenishmentMethod             sql.NullString  `bun:"replenishment_method,type:varchar(8),nullzero"`
	OrderQuantity                   sql.NullFloat64 `bun:"order_quantity,type:decimal(19,9),default:(0)"`
	QtyAllocated                    sql.NullFloat64 `bun:"qty_allocated,type:decimal(19,9),default:(0)"`
	QtyBackordered                  sql.NullFloat64 `bun:"qty_backordered,type:decimal(19,9),default:(0)"`
	QtyInTransit                    sql.NullFloat64 `bun:"qty_in_transit,type:decimal(19,9),default:(0)"`
	TrackBins                       sql.NullString  `bun:"track_bins,type:char,nullzero"`
	DefaultInOe                     sql.NullString  `bun:"default_in_oe,type:char,nullzero"`
	PeriodFirstStocked              sql.NullFloat64 `bun:"period_first_stocked,type:decimal(3,0),nullzero"`
	YearFirstStocked                sql.NullFloat64 `bun:"year_first_stocked,type:decimal(4,0),nullzero"`
	UsageLock                       sql.NullString  `bun:"usage_lock,type:char,nullzero"`
	UsageLockPeriod                 sql.NullFloat64 `bun:"usage_lock_period,type:decimal(3,0),nullzero"`
	UsageLockYear                   sql.NullFloat64 `bun:"usage_lock_year,type:decimal(4,0),nullzero"`
	PrimaryBin                      sql.NullString  `bun:"primary_bin,type:varchar(10),nullzero"`
	TaxGroupId                      sql.NullString  `bun:"tax_group_id,type:varchar(10),nullzero"`
	QtyReservedDueIn                sql.NullFloat64 `bun:"qty_reserved_due_in,type:decimal(19,9),default:(0)"`
	DateLastCounted                 sql.NullString  `bun:"date_last_counted,type:smalldatetime,nullzero"`
	InvMastUid                      int32           `bun:"inv_mast_uid,type:int,pk"`
	Requisition                     string          `bun:"requisition,type:char,default:('N')"`
	DeadstockFlag                   string          `bun:"deadstock_flag,type:char,default:('N')"`
	LotBinIntegration               sql.NullString  `bun:"lot_bin_integration,type:char,nullzero"`
	AddToCycleCount                 string          `bun:"add_to_cycle_count,type:char,default:('N')"`
	DefaultShipment                 int32           `bun:"default_shipment,type:int,default:(1138)"`
	PrimarySupplierId               sql.NullFloat64 `bun:"primary_supplier_id,type:decimal(19,0),nullzero"`
	OnObtFlag                       sql.NullString  `bun:"on_obt_flag,type:char,default:('N')"`
	OnReleaseScheduleFlag           sql.NullString  `bun:"on_release_schedule_flag,type:char,default:('N')"`
	OnBackorderFlag                 sql.NullString  `bun:"on_backorder_flag,type:char,default:('N')"`
	LastSaleDate                    time.Time       `bun:"last_sale_date,type:datetime,default:('01/01/1990')"`
	LastPurchaseDate                sql.NullTime    `bun:"last_purchase_date,type:datetime,nullzero"`
	Buy                             string          `bun:"buy,type:char,default:('Y')"`
	Make                            string          `bun:"make,type:char,default:('N')"`
	CreatedBy                       sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	PeriodsToSupplyMin              sql.NullInt16   `bun:"periods_to_supply_min,type:tinyint,default:(0)"`
	PeriodsToSupplyMax              sql.NullInt16   `bun:"periods_to_supply_max,type:tinyint,default:(0)"`
	PutawayRank                     sql.NullString  `bun:"putaway_rank,type:varchar(8),nullzero"`
	InvLastChangedDate              sql.NullTime    `bun:"inv_last_changed_date,type:datetime,nullzero"`
	IvaTaxableFlag                  sql.NullString  `bun:"iva_taxable_flag,type:char,nullzero"`
	ItemPutawayAttributeUid         sql.NullInt32   `bun:"item_putaway_attribute_uid,type:int,nullzero"`
	SplittableFlag                  sql.NullString  `bun:"splittable_flag,type:char,nullzero"`
	MinReplenishmentQty             sql.NullFloat64 `bun:"min_replenishment_qty,type:decimal(19,9),nullzero"`
	Discontinued                    string          `bun:"discontinued,type:char,default:('N')"`
	AllowDsDiscontinuedItems        string          `bun:"allow_ds_discontinued_items,type:char,default:('N')"`
	AllowSpDiscontinuedItems        string          `bun:"allow_sp_discontinued_items,type:char,default:('N')"`
	SafetyStockType                 sql.NullInt32   `bun:"safety_stock_type,type:int,nullzero"`
	ServiceLevelMeasure             sql.NullInt32   `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal             sql.NullFloat64 `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	DemandPatternCd                 sql.NullInt32   `bun:"demand_pattern_cd,type:int,nullzero"`
	DemandPatternEvaluationDate     sql.NullTime    `bun:"demand_pattern_evaluation_date,type:datetime,nullzero"`
	DemandForecastFormulaUid        sql.NullInt32   `bun:"demand_forecast_formula_uid,type:int,nullzero"`
	PatternManuallyEditedFlag       sql.NullString  `bun:"pattern_manually_edited_flag,type:char,nullzero"`
	DemandPatternBehaviorCd         sql.NullInt32   `bun:"demand_pattern_behavior_cd,type:int,nullzero"`
	PatternLikeInvMastUid           sql.NullInt32   `bun:"pattern_like_inv_mast_uid,type:int,nullzero"`
	PatternLikeLocationId           sql.NullFloat64 `bun:"pattern_like_location_id,type:decimal(19,0),nullzero"`
	BehavesLikeLockFlag             sql.NullString  `bun:"behaves_like_lock_flag,type:char,nullzero"`
	BehavesLikeLockPeriod           sql.NullInt32   `bun:"behaves_like_lock_period,type:int,nullzero"`
	BehavesLikeLockYear             sql.NullInt32   `bun:"behaves_like_lock_year,type:int,nullzero"`
	LandedCostAccountNo             sql.NullString  `bun:"landed_cost_account_no,type:varchar(255),nullzero"`
	PriceFamilyUid                  sql.NullInt32   `bun:"price_family_uid,type:int,nullzero"`
	DeleteFlag                      sql.NullString  `bun:"delete_flag,type:char,nullzero"`
	CardlockProductId               sql.NullString  `bun:"cardlock_product_id,type:varchar(255),nullzero"`
	DefaultSellingUnit              sql.NullString  `bun:"default_selling_unit,type:varchar(8),nullzero"`
	LocItemTypeCd                   sql.NullInt32   `bun:"loc_item_type_cd,type:int,nullzero"`
	LocCustParentInvMastUid         sql.NullInt32   `bun:"loc_cust_parent_inv_mast_uid,type:int,nullzero"`
	MainBulkLocationFlag            sql.NullString  `bun:"main_bulk_location_flag,type:char,nullzero"`
	DfltSourceLocFlag               sql.NullString  `bun:"dflt_source_loc_flag,type:char,nullzero"`
	CycleCountFlag                  sql.NullString  `bun:"cycle_count_flag,type:char,nullzero"`
	PromotionalFlag                 sql.NullString  `bun:"promotional_flag,type:char,nullzero"`
	RmaRevenueAccountNo             sql.NullString  `bun:"rma_revenue_account_no,type:varchar(32),nullzero"`
	FutureStandardCost              sql.NullFloat64 `bun:"future_standard_cost,type:decimal(19,9),nullzero"`
	EffectiveDate                   sql.NullTime    `bun:"effective_date,type:datetime,nullzero"`
	AssemblyPromptOption            sql.NullInt32   `bun:"assembly_prompt_option,type:int,nullzero"`
	RestrictedFlag                  sql.NullString  `bun:"restricted_flag,type:char,nullzero"`
	DrpItemFlag                     string          `bun:"drp_item_flag,type:char,default:('N')"`
	SavedDemandForecastFormulaUid   sql.NullInt32   `bun:"saved_demand_forecast_formula_uid,type:int,nullzero"`
	DemandFormulaComputedYearPeriod sql.NullInt32   `bun:"demand_formula_computed_year_period,type:int,nullzero"`
	DemandPatternComputedYearPeriod sql.NullInt32   `bun:"demand_pattern_computed_year_period,type:int,nullzero"`
	YearPeriodLastForecast          sql.NullInt32   `bun:"year_period_last_forecast,type:int,nullzero"`
	VendorRebateAccountNo           sql.NullString  `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`
	TransferUsagePercent            sql.NullFloat64 `bun:"transfer_usage_percent,type:decimal(19,9),nullzero"`
	ManufacturerReqStockFlag        string          `bun:"manufacturer_req_stock_flag,type:char,default:('N')"`
	FcBin                           sql.NullString  `bun:"fc_bin,type:varchar(10),nullzero"`
	PumpoffItem                     sql.NullString  `bun:"pumpoff_item,type:varchar(40),nullzero"`
	PumpoffUom                      sql.NullString  `bun:"pumpoff_uom,type:varchar(8),nullzero"`
	OnFutureOrderFlag               sql.NullString  `bun:"on_future_order_flag,type:char,default:('N')"`
	OnFutureProdOrderFlag           sql.NullString  `bun:"on_future_prod_order_flag,type:char,default:('N')"`
	MaxLiability                    sql.NullFloat64 `bun:"max_liability,type:decimal(19,9),nullzero"`
	MaxTransferQty                  sql.NullFloat64 `bun:"max_transfer_qty,type:decimal(19,9),nullzero"`
	AltTaxGroupId                   sql.NullString  `bun:"alt_tax_group_id,type:varchar(10),nullzero"`
	ReturnableFlag                  sql.NullString  `bun:"returnable_flag,type:char,default:('Y')"`
	IqsItemFlag                     sql.NullString  `bun:"iqs_item_flag,type:char,nullzero"`
	SlabTrackBinsFlag               sql.NullString  `bun:"slab_track_bins_flag,type:char,nullzero"`
	Edi832DiscontinuedSentFlag      sql.NullString  `bun:"edi_832_discontinued_sent_flag,type:char,nullzero"`
	BuyerId                         sql.NullString  `bun:"buyer_id,type:varchar(16),nullzero"`
	TaRentalRevenueAccountNo        sql.NullString  `bun:"ta_rental_revenue_account_no,type:varchar(32),nullzero"`
	CycleCountNo                    sql.NullFloat64 `bun:"cycle_count_no,type:decimal(19,0),nullzero"`
	CycleCountQty                   sql.NullFloat64 `bun:"cycle_count_qty,type:decimal(19,9),nullzero"`
	TaRentalTaxGroupId              sql.NullString  `bun:"ta_rental_tax_group_id,type:varchar(10),nullzero"`
	TransferConversionUom           sql.NullString  `bun:"transfer_conversion_uom,type:varchar(8),nullzero"`
	EccEnabledFlag                  string          `bun:"ecc_enabled_flag,type:char,default:('N')"`
	InvMastUidDupUsage              sql.NullInt32   `bun:"inv_mast_uid_dup_usage,type:int,nullzero"`
	TransferOrderPoint              sql.NullFloat64 `bun:"transfer_order_point,type:decimal(19,9),nullzero"`
	TransferOrderQty                sql.NullFloat64 `bun:"transfer_order_qty,type:decimal(19,9),nullzero"`
	OsmiItemFlag                    sql.NullString  `bun:"osmi_item_flag,type:char,default:('N')"`
	LocalItemFlag                   sql.NullString  `bun:"local_item_flag,type:char,default:('N')"`
	CriticalItemFlag                sql.NullString  `bun:"critical_item_flag,type:char,nullzero"`
	DeviationLookbackPds            sql.NullInt32   `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd              sql.NullFloat64 `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd              sql.NullFloat64 `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd            sql.NullFloat64 `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays              sql.NullFloat64 `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays              sql.NullFloat64 `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult           sql.NullFloat64 `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays          sql.NullFloat64 `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays          sql.NullFloat64 `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

var _ bun.BeforeAppendModelHook = (*InvLoc)(nil)

func (m *InvLoc) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type InvLocModel struct {
	bun bun.IDB
}

func (m *InvLocModel) Get(ctx context.Context, locationId float64, companyId string, invMastUid int32) (
	*InvLoc, error) {
	var invLoc InvLoc
	err := m.bun.NewSelect().Model(&invLoc).Relation("InvMast").Where(
		"inv_loc.location_id = ? AND inv_loc.company_id = ? AND inv_loc.inv_mast_uid = ? ", locationId, companyId,
		invMastUid).Scan(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &invLoc, nil
}

// GetByProductGroupId selects InvLoc by product_group_id and returns records
func (m *InvLocModel) GetByProductGroupId(ctx context.Context, productGroupId string) ([]InvLoc, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var invLocs []InvLoc

	err := m.bun.NewSelect().Model(&invLocs).Where(
		"product_group_id = ?",
		productGroupId).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return invLocs, nil

}

// GetByInvMastUid selects InvLoc by inv_mast_uid and returns associated InvMast and ProductGroup
func (m *InvLocModel) GetByInvMastUid(ctx context.Context, locationIds []float64, uid int32) (
	[]InvLoc,
	error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var invLocs []InvLoc
	err := m.bun.NewSelect().Model(&invLocs).Relation("InvMast").Where(
		"inv_loc.inv_mast_uid = ?",
		uid).Where("inv_loc.location_id IN (?)", bun.In(locationIds)).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return invLocs, nil
}

// GetByInvMastUids selects InvLoc by inv_mast_uid and returns associated InvMast and ProductGroup
func (m *InvLocModel) GetByInvMastUids(ctx context.Context, invMastUids []int32) ([]*InvLoc, error) {
	var invLocs []*InvLoc

	err := m.bun.NewSelect().Model(&invLocs).Relation("InvMast").Where(
		"inv_loc.inv_mast_uid IN (?)", bun.In(invMastUids)).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invLocs, nil
}

// GetAll selects all records from inv_loc table
func (m *InvLocModel) GetAll(
	ctx context.Context, locationId float64, deleteFlag DeleteFlag, filter Filters) (
	[]InvLoc, Metadata, error,
) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var invLocs []InvLoc
	err := m.bun.NewSelect().Model(&invLocs).Relation("InvMast").Where(
		"inv_loc.inv_mast_uid > ?", filter.cursor()).Where(
		"location_id = ?", locationId).Limit(int(filter.limit())).Order("inv_loc.inv_mast_uid ASC").Scan(ctx)
	if err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(1, invLocs[len(invLocs)-1].InvMastUid, filter.cursor())
	return invLocs, metadata, nil
}

// Update updates record with provided invLoc if matching version is found
func (m *InvLocModel) Update(ctx context.Context, invLoc *InvLoc) error {
	_, err := m.bun.NewUpdate().Model(invLoc).WherePK().Exec(ctx)
	if err != nil {
		switch {
		case strings.Contains(
			err.Error(),
			"The UPDATE statement conflicted with the FOREIGN KEY constraint \"fk_inv_loc_product_group\""):
			return errors.New("the provided product group was not found")
		}
		return err
	}
	return nil

}

func (m *InvLocModel) Delete(ctx context.Context, invLoc *InvLoc) error {
	_, err := m.bun.NewDelete().Model(invLoc).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
