package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerStrategicPricing struct {
	bun.BaseModel               `bun:"table:customer_strategic_pricing"`
	CustomerStrategicPricingUid int32     `bun:"customer_strategic_pricing_uid,type:int,pk"`                   // Unique identifier for the table
	CustomerId                  float64   `bun:"customer_id,type:decimal(19,0),unique"`                        // Customer for this strategic pricing.
	CompanyId                   string    `bun:"company_id,type:varchar(8),unique"`                            // Company the customer relates to.
	WarehouseSizeCd             int32     `bun:"warehouse_size_cd,type:int,nullzero"`                          // Warehouse size: Very Tiny/Tiny/Small/Medium/Large/Huge
	WarehouseSizeUpdateDate     time.Time `bun:"warehouse_size_update_date,type:datetime,nullzero"`            // Date warehouse size was last updated
	RetailSizeCd                int32     `bun:"retail_size_cd,type:int,nullzero"`                             // Retail size: Very Tiny/Tiny/Small/Medium/Large/Huge
	RetailSizeUpdateDate        time.Time `bun:"retail_size_update_date,type:datetime,nullzero"`               // Date retail size was last updated
	CustomerCategoryUid         int32     `bun:"customer_category_uid,type:int,nullzero"`                      // Indicates the customer category
	ListPriceOptionCd           int32     `bun:"list_price_option_cd,type:int"`                                // Strategic/Supplier - Use Strategic or Supplier list price/cost.
	FreightChargeOptionCd       int32     `bun:"freight_charge_option_cd,type:int"`                            // Indicates either if Strategic Freight Charge or Actual Freight Cost is used
	PricingCustomerId           float64   `bun:"pricing_customer_id,type:decimal(19,0),nullzero"`              // The pricing customer/company is the customer used to determine pricing information.  The user can group customers who all get the same price.
	CustomerSensitivityCd       int32     `bun:"customer_sensitivity_cd,type:int"`                             // Customer Sensitivity to Prices: Very Low/Low/Medium/High/Very High
	CustSensitivityUpdateDate   time.Time `bun:"cust_sensitivity_update_date,type:datetime,nullzero"`          // Date customer sensitivity was last updated
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
