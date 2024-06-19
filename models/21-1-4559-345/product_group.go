package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProductGroup struct {
	bun.BaseModel                `bun:"table:product_group"`
	CompanyId                    string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	ProductGroupId               string    `bun:"product_group_id,type:varchar(8),pk"`                           // Unique Identifier for this Group.
	ProductGroupDesc             string    `bun:"product_group_desc,type:varchar(80)"`                           // Description of this group.
	AssetAccountNo               string    `bun:"asset_account_no,type:varchar(80)"`                             // Default Asset account for this product group.
	RevenueAccountNo             string    `bun:"revenue_account_no,type:varchar(80),nullzero"`                  // The default revenue account number for this product group.
	CosAccountNo                 string    `bun:"cos_account_no,type:varchar(80),nullzero"`                      // Default COS account for this product group.
	DeleteFlag                   string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                  time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	ParkerProductCode            string    `bun:"parker_product_code,type:varchar(255),nullzero"`                // stores the parker product code for this group.
	TaxGroupId                   string    `bun:"tax_group_id,type:varchar(10),nullzero"`                        // Foreign key to the table tax_group_hdr
	CompletionLeadTime           int32     `bun:"completion_lead_time,type:int,nullzero"`                        // This custom column will store the completion lead days needed for this product group
	RequiredLeadTime             int32     `bun:"required_lead_time,type:int,nullzero"`                          // This custom column will store the required lead days needed for this product group
	ProductGroupUid              int32     `bun:"product_group_uid,type:int,autoincrement,unique,scanonly"`      // Unique internal ID number for this table.
	LandedCostAccountNo          string    `bun:"landed_cost_account_no,type:varchar(255),nullzero"`             // Landed Cost income account
	EnvironmentalFeeFlag         string    `bun:"environmental_fee_flag,type:char(1),default:('N')"`             // Indicates if the customer will be charged an environmental fee.
	AdminFeeFlag                 string    `bun:"admin_fee_flag,type:char(1),default:('N')"`                     // Indicates if this customer should be charged a handling/admin fee?
	IncludeInSizeAnalysisFlag    string    `bun:"include_in_size_analysis_flag,type:char(1),default:('Y')"`      // Set to Y if product group is to be included in Customer Size analysis
	RmaRevenueAccountNo          string    `bun:"rma_revenue_account_no,type:varchar(80),nullzero"`              // RMA Revenue Account
	PriceRoundingFlag            string    `bun:"price_rounding_flag,type:char(1),default:('N')"`                // Determines if items with this product group should try to round prices.
	PriceRoundingThreshold       float64   `bun:"price_rounding_threshold,type:decimal(19,9),nullzero"`          // The price amount above which to apply price rounding.
	ApplyMinMarginRulesFlag      string    `bun:"apply_min_margin_rules_flag,type:char(1),default:('N')"`        // Custom: Indicates if minimum margin rules should be applied to items in this product group.
	VendorRebateAccountNo        string    `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`            // A general ledger account that tracks the value of Vendor Rebates
	EnableLineProfitWarning      string    `bun:"enable_line_profit_warning,type:char(1),default:('N')"`         // Product group profit warning enabled
	MinimumOrderLineProfit       float64   `bun:"minimum_order_line_profit,type:decimal(19,9),nullzero"`         // Minimum profit
	MaximumOrderLineProfit       float64   `bun:"maximum_order_line_profit,type:decimal(19,9),nullzero"`         // Maximum profit
	OeSkipProfitCheckUnpriced    string    `bun:"oe_skip_profit_check_unpriced,type:char(1),default:('N')"`      // Skip unpriced
	OeSkipProfitCheckUncosted    string    `bun:"oe_skip_profit_check_uncosted,type:char(1),default:('N')"`      // Skip uncosted
	CompressorFlag               string    `bun:"compressor_flag,type:char(1),default:('N')"`                    // Indicate if the items in the Product Group are Compressors
	CardlockProductGroupTypeCd   int32     `bun:"cardlock_product_group_type_cd,type:int,nullzero"`              // The type of cardlock product type of a product group
	TaRentalRevenueAccountNo     string    `bun:"ta_rental_revenue_account_no,type:varchar(32),nullzero"`        // TrackAbout Rental Revenue Account for product group
	PercentReqToAllocTransfer    float64   `bun:"percent_req_to_alloc_transfer,type:decimal(19,9),nullzero"`     // Custom: Indicates that items with this corresponding product group will only be allocated to transfers when this percent of the quantity is available.
	FutureEffectiveDate          time.Time `bun:"future_effective_date,type:datetime,nullzero"`                  // Has the date from where the price rounding threshold can be updated.
	FuturePriceRoundingThreshold float64   `bun:"future_price_rounding_threshold,type:decimal(19,9),nullzero"`   // Contains the new price threshold that is going to be updated into
}
