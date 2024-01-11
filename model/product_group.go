package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ProductGroup struct {
	bun.BaseModel                `bun:"table:product_group"`
	CompanyId                    string          `bun:"company_id,type:varchar(8),pk"`
	ProductGroupId               string          `bun:"product_group_id,type:varchar(8),pk"`
	ProductGroupDesc             string          `bun:"product_group_desc,type:varchar(40)"`
	AssetAccountNo               string          `bun:"asset_account_no,type:varchar(80)"`
	RevenueAccountNo             sql.NullString  `bun:"revenue_account_no,type:varchar(80)"`
	CosAccountNo                 sql.NullString  `bun:"cos_account_no,type:varchar(80)"`
	DeleteFlag                   string          `bun:"delete_flag,type:char"`
	DateCreated                  time.Time       `bun:"date_created,type:datetime"`
	DateLastModified             time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy             string          `bun:"last_maintained_by,type:varchar(30),default"`
	ParkerProductCode            sql.NullString  `bun:"parker_product_code,type:varchar(255)"`
	TaxGroupId                   sql.NullString  `bun:"tax_group_id,type:varchar(10)"`
	CompletionLeadTime           sql.NullInt32   `bun:"completion_lead_time,type:int"`
	RequiredLeadTime             sql.NullInt32   `bun:"required_lead_time,type:int"`
	ProductGroupUid              int32           `bun:"product_group_uid,type:int,autoincrement"`
	LandedCostAccountNo          sql.NullString  `bun:"landed_cost_account_no,type:varchar(255)"`
	EnvironmentalFeeFlag         string          `bun:"environmental_fee_flag,type:char,default"`
	AdminFeeFlag                 string          `bun:"admin_fee_flag,type:char,default"`
	IncludeInSizeAnalysisFlag    sql.NullString  `bun:"include_in_size_analysis_flag,type:char,default"`
	RmaRevenueAccountNo          sql.NullString  `bun:"rma_revenue_account_no,type:varchar(80)"`
	PriceRoundingFlag            string          `bun:"price_rounding_flag,type:char,default"`
	PriceRoundingThreshold       sql.NullFloat64 `bun:"price_rounding_threshold,type:decimal(19,9)"`
	ApplyMinMarginRulesFlag      string          `bun:"apply_min_margin_rules_flag,type:char,default"`
	VendorRebateAccountNo        sql.NullString  `bun:"vendor_rebate_account_no,type:varchar(32)"`
	EnableLineProfitWarning      string          `bun:"enable_line_profit_warning,type:char,default"`
	MinimumOrderLineProfit       sql.NullFloat64 `bun:"minimum_order_line_profit,type:decimal(19,9)"`
	MaximumOrderLineProfit       sql.NullFloat64 `bun:"maximum_order_line_profit,type:decimal(19,9)"`
	OeSkipProfitCheckUnpriced    string          `bun:"oe_skip_profit_check_unpriced,type:char,default"`
	OeSkipProfitCheckUncosted    string          `bun:"oe_skip_profit_check_uncosted,type:char,default"`
	CompressorFlag               string          `bun:"compressor_flag,type:char,default"`
	CardlockProductGroupTypeCd   sql.NullInt32   `bun:"cardlock_product_group_type_cd,type:int"`
	TaRentalRevenueAccountNo     sql.NullString  `bun:"ta_rental_revenue_account_no,type:varchar(32)"`
	PercentReqToAllocTransfer    sql.NullFloat64 `bun:"percent_req_to_alloc_transfer,type:decimal(19,9)"`
	FutureEffectiveDate          sql.NullTime    `bun:"future_effective_date,type:datetime"`
	FuturePriceRoundingThreshold sql.NullFloat64 `bun:"future_price_rounding_threshold,type:decimal(19,9)"`
}
