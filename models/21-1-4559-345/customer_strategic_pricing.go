package model

type CustomerStrategicPricing struct {
	bun.BaseModel               `bun:"table:customer_strategic_pricing"`
	CustomerStrategicPricingUid int32     `bun:"customer_strategic_pricing_uid,type:int,pk"`
	CustomerId                  float64   `bun:"customer_id,type:decimal(19,0)"`
	CompanyId                   string    `bun:"company_id,type:varchar(8)"`
	WarehouseSizeCd             int32     `bun:"warehouse_size_cd,type:int,nullzero"`
	WarehouseSizeUpdateDate     time.Time `bun:"warehouse_size_update_date,type:datetime,nullzero"`
	RetailSizeCd                int32     `bun:"retail_size_cd,type:int,nullzero"`
	RetailSizeUpdateDate        time.Time `bun:"retail_size_update_date,type:datetime,nullzero"`
	CustomerCategoryUid         int32     `bun:"customer_category_uid,type:int,nullzero"`
	ListPriceOptionCd           int32     `bun:"list_price_option_cd,type:int"`
	FreightChargeOptionCd       int32     `bun:"freight_charge_option_cd,type:int"`
	PricingCustomerId           float64   `bun:"pricing_customer_id,type:decimal(19,0),nullzero"`
	CustomerSensitivityCd       int32     `bun:"customer_sensitivity_cd,type:int"`
	CustSensitivityUpdateDate   time.Time `bun:"cust_sensitivity_update_date,type:datetime,nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
