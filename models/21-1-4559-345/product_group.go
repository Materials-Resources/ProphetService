package model

type ProductGroup struct {
	bun.BaseModel                `bun:"table:product_group"`
	CompanyId                    string    `bun:"company_id,type:varchar(8),pk"`
	ProductGroupId               string    `bun:"product_group_id,type:varchar(8),pk"`
	ProductGroupDesc             string    `bun:"product_group_desc,type:varchar(80)"`
	AssetAccountNo               string    `bun:"asset_account_no,type:varchar(80)"`
	RevenueAccountNo             string    `bun:"revenue_account_no,type:varchar(80),nullzero"`
	CosAccountNo                 string    `bun:"cos_account_no,type:varchar(80),nullzero"`
	DeleteFlag                   string    `bun:"delete_flag,type:char"`
	DateCreated                  time.Time `bun:"date_created,type:datetime"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ParkerProductCode            string    `bun:"parker_product_code,type:varchar(255),nullzero"`
	TaxGroupId                   string    `bun:"tax_group_id,type:varchar(10),nullzero"`
	CompletionLeadTime           int32     `bun:"completion_lead_time,type:int,nullzero"`
	RequiredLeadTime             int32     `bun:"required_lead_time,type:int,nullzero"`
	ProductGroupUid              int32     `bun:"product_group_uid,type:int,identity"`
	LandedCostAccountNo          string    `bun:"landed_cost_account_no,type:varchar(255),nullzero"`
	EnvironmentalFeeFlag         string    `bun:"environmental_fee_flag,type:char,default:('N')"`
	AdminFeeFlag                 string    `bun:"admin_fee_flag,type:char,default:('N')"`
	IncludeInSizeAnalysisFlag    string    `bun:"include_in_size_analysis_flag,type:char,default:('Y')"`
	RmaRevenueAccountNo          string    `bun:"rma_revenue_account_no,type:varchar(80),nullzero"`
	PriceRoundingFlag            string    `bun:"price_rounding_flag,type:char,default:('N')"`
	PriceRoundingThreshold       float64   `bun:"price_rounding_threshold,type:decimal(19,9),nullzero"`
	ApplyMinMarginRulesFlag      string    `bun:"apply_min_margin_rules_flag,type:char,default:('N')"`
	VendorRebateAccountNo        string    `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`
	EnableLineProfitWarning      string    `bun:"enable_line_profit_warning,type:char,default:('N')"`
	MinimumOrderLineProfit       float64   `bun:"minimum_order_line_profit,type:decimal(19,9),nullzero"`
	MaximumOrderLineProfit       float64   `bun:"maximum_order_line_profit,type:decimal(19,9),nullzero"`
	OeSkipProfitCheckUnpriced    string    `bun:"oe_skip_profit_check_unpriced,type:char,default:('N')"`
	OeSkipProfitCheckUncosted    string    `bun:"oe_skip_profit_check_uncosted,type:char,default:('N')"`
	CompressorFlag               string    `bun:"compressor_flag,type:char,default:('N')"`
	CardlockProductGroupTypeCd   int32     `bun:"cardlock_product_group_type_cd,type:int,nullzero"`
	TaRentalRevenueAccountNo     string    `bun:"ta_rental_revenue_account_no,type:varchar(32),nullzero"`
	PercentReqToAllocTransfer    float64   `bun:"percent_req_to_alloc_transfer,type:decimal(19,9),nullzero"`
	FutureEffectiveDate          time.Time `bun:"future_effective_date,type:datetime,nullzero"`
	FuturePriceRoundingThreshold float64   `bun:"future_price_rounding_threshold,type:decimal(19,9),nullzero"`
}
