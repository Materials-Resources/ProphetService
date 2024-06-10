package models

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

var _ bun.BeforeAppendModelHook = (*ProductGroup)(nil)

type ProductGroup struct {
	bun.BaseModel                `bun:"table:product_group"`
	CompanyId                    string          `bun:"company_id,type:varchar(8),pk"`
	ProductGroupId               string          `bun:"product_group_id,type:varchar(8),pk"`
	ProductGroupDesc             string          `bun:"product_group_desc,type:varchar(80)"`
	AssetAccountNo               string          `bun:"asset_account_no,type:varchar(80)"`
	RevenueAccountNo             sql.NullString  `bun:"revenue_account_no,type:varchar(80),nullzero"`
	CosAccountNo                 sql.NullString  `bun:"cos_account_no,type:varchar(80),nullzero"`
	DeleteFlag                   string          `bun:"delete_flag,type:char"`
	DateCreated                  time.Time       `bun:"date_created,type:datetime"`
	DateLastModified             time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy             string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ParkerProductCode            sql.NullString  `bun:"parker_product_code,type:varchar(255),nullzero"`
	TaxGroupId                   sql.NullString  `bun:"tax_group_id,type:varchar(10),nullzero"`
	CompletionLeadTime           sql.NullInt32   `bun:"completion_lead_time,type:int,nullzero"`
	RequiredLeadTime             sql.NullInt32   `bun:"required_lead_time,type:int,nullzero"`
	ProductGroupUid              int32           `bun:"product_group_uid,type:int,identity,scanonly"`
	LandedCostAccountNo          sql.NullString  `bun:"landed_cost_account_no,type:varchar(255),nullzero"`
	EnvironmentalFeeFlag         string          `bun:"environmental_fee_flag,type:char,default:('N')"`
	AdminFeeFlag                 string          `bun:"admin_fee_flag,type:char,default:('N')"`
	IncludeInSizeAnalysisFlag    sql.NullString  `bun:"include_in_size_analysis_flag,type:char,default:('Y')"`
	RmaRevenueAccountNo          sql.NullString  `bun:"rma_revenue_account_no,type:varchar(80),nullzero"`
	PriceRoundingFlag            string          `bun:"price_rounding_flag,type:char,default:('N')"`
	PriceRoundingThreshold       sql.NullFloat64 `bun:"price_rounding_threshold,type:decimal(19,9),nullzero"`
	ApplyMinMarginRulesFlag      string          `bun:"apply_min_margin_rules_flag,type:char,default:('N')"`
	VendorRebateAccountNo        sql.NullString  `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`
	EnableLineProfitWarning      string          `bun:"enable_line_profit_warning,type:char,default:('N')"`
	MinimumOrderLineProfit       sql.NullFloat64 `bun:"minimum_order_line_profit,type:decimal(19,9),nullzero"`
	MaximumOrderLineProfit       sql.NullFloat64 `bun:"maximum_order_line_profit,type:decimal(19,9),nullzero"`
	OeSkipProfitCheckUnpriced    string          `bun:"oe_skip_profit_check_unpriced,type:char,default:('N')"`
	OeSkipProfitCheckUncosted    string          `bun:"oe_skip_profit_check_uncosted,type:char,default:('N')"`
	CompressorFlag               string          `bun:"compressor_flag,type:char,default:('N')"`
	CardlockProductGroupTypeCd   sql.NullInt32   `bun:"cardlock_product_group_type_cd,type:int,nullzero"`
	TaRentalRevenueAccountNo     sql.NullString  `bun:"ta_rental_revenue_account_no,type:varchar(32),nullzero"`
	PercentReqToAllocTransfer    sql.NullFloat64 `bun:"percent_req_to_alloc_transfer,type:decimal(19,9),nullzero"`
	FutureEffectiveDate          sql.NullTime    `bun:"future_effective_date,type:datetime,nullzero"`
	FuturePriceRoundingThreshold sql.NullFloat64 `bun:"future_price_rounding_threshold,type:decimal(19,9),nullzero"`
}

func (p *ProductGroup) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		p.DateCreated = time.Now()
		p.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		p.DateLastModified = time.Now()
	}
	return nil
}

func (p *ProductGroup) setDefaults() {
	p.CompanyId = "MRS"
	p.AssetAccountNo = "121001"
	p.DeleteFlag = "N"
	p.RevenueAccountNo = sql.NullString{String: "401001", Valid: true}
	p.CosAccountNo = sql.NullString{String: "501001", Valid: true}
	p.LastMaintainedBy = "admin"
}

type ProductGroupModel struct {
	bun bun.IDB
}

// GetById returns a ProductGroup by the given product_group_id
func (m *ProductGroupModel) GetById(ctx context.Context, id string) (*ProductGroup, error) {
	var productGroup ProductGroup
	err := m.bun.NewSelect().Model(&productGroup).Where("product_group_id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &productGroup, nil
}

func (m *ProductGroupModel) Get(ctx context.Context, uid int32) (*ProductGroup, error) {
	var productGroup ProductGroup
	err := m.bun.NewSelect().Model(&productGroup).Where("product_group_uid = ?", uid).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &productGroup, nil

}

// Insert creates a new ProductGroup for the given data
func (m *ProductGroupModel) Insert(ctx context.Context, productGroup *ProductGroup) error {

	productGroup.setDefaults()

	_, err := m.bun.NewInsert().Model(productGroup).Exec(ctx)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "mssql: Violation of PRIMARY KEY constraint 'pk_product_group'."):
			return ErrDuplicateRecord
		}
		return err
	}

	return nil
}

// Update updates the data for ProductGroup
func (m *ProductGroupModel) Update(ctx context.Context, productGroup *ProductGroup) error {
	_, err := m.bun.NewUpdate().Model(productGroup).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductGroupModel) GetAll(ctx context.Context) ([]ProductGroup, error) {
	var productGroups []ProductGroup
	err := m.bun.NewSelect().Model(&productGroups).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return productGroups, nil

}
