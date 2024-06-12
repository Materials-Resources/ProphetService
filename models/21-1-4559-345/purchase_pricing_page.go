package model

type PurchasePricingPage struct {
	bun.BaseModel          `bun:"table:purchase_pricing_page"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	PricingBookId          string    `bun:"pricing_book_id,type:varchar(8)"`
	SupplierId             float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	DiscountGroupId        string    `bun:"discount_group_id,type:varchar(8),nullzero"`
	MajorGroupId           string    `bun:"major_group_id,type:varchar(8),nullzero"`
	PricingDescription     string    `bun:"pricing_description,type:varchar(40),nullzero"`
	PricingMethod          string    `bun:"pricing_method,type:varchar(16)"`
	SourcePrice            string    `bun:"source_price,type:varchar(40),nullzero"`
	Price                  float64   `bun:"price,type:decimal(19,9),nullzero"`
	EffectiveDate          time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate         time.Time `bun:"expiration_date,type:datetime"`
	ContractNumber         string    `bun:"contract_number,type:varchar(40),nullzero"`
	CalculationType        string    `bun:"calculation_type,type:varchar(16)"`
	CalculationValue1      float64   `bun:"calculation_value1,type:decimal(19,9)"`
	CalculationValue2      float64   `bun:"calculation_value2,type:decimal(19,9)"`
	CalculationValue3      float64   `bun:"calculation_value3,type:decimal(19,9)"`
	CalculationValue4      float64   `bun:"calculation_value4,type:decimal(19,9)"`
	CalculationValue5      float64   `bun:"calculation_value5,type:decimal(19,9)"`
	CalculationValue6      float64   `bun:"calculation_value6,type:decimal(19,9)"`
	CalculationValue7      float64   `bun:"calculation_value7,type:decimal(19,9)"`
	CalculationValue8      float64   `bun:"calculation_value8,type:decimal(19,9)"`
	CalculationValue9      float64   `bun:"calculation_value9,type:decimal(19,9)"`
	CalculationValue10     float64   `bun:"calculation_value10,type:decimal(19,9)"`
	CalculationValue11     float64   `bun:"calculation_value11,type:decimal(19,9)"`
	CalculationValue12     float64   `bun:"calculation_value12,type:decimal(19,9)"`
	CalculationValue13     float64   `bun:"calculation_value13,type:decimal(19,9)"`
	CalculationValue14     float64   `bun:"calculation_value14,type:decimal(19,9)"`
	CalculationValue15     float64   `bun:"calculation_value15,type:decimal(19,9)"`
	Break1                 float64   `bun:"break1,type:decimal(19,9)"`
	Break2                 float64   `bun:"break2,type:decimal(19,9)"`
	Break3                 float64   `bun:"break3,type:decimal(19,9)"`
	Break4                 float64   `bun:"break4,type:decimal(19,9)"`
	Break5                 float64   `bun:"break5,type:decimal(19,9)"`
	Break6                 float64   `bun:"break6,type:decimal(19,9)"`
	Break7                 float64   `bun:"break7,type:decimal(19,9)"`
	Break8                 float64   `bun:"break8,type:decimal(19,9)"`
	Break9                 float64   `bun:"break9,type:decimal(19,9)"`
	Break10                float64   `bun:"break10,type:decimal(19,9)"`
	Break11                float64   `bun:"break11,type:decimal(19,9)"`
	Break12                float64   `bun:"break12,type:decimal(19,9)"`
	Break13                float64   `bun:"break13,type:decimal(19,9)"`
	Break14                float64   `bun:"break14,type:decimal(19,9)"`
	Uom1                   string    `bun:"uom1,type:varchar(8),nullzero"`
	Uom2                   string    `bun:"uom2,type:varchar(8),nullzero"`
	Uom3                   string    `bun:"uom3,type:varchar(8),nullzero"`
	Uom4                   string    `bun:"uom4,type:varchar(8),nullzero"`
	Uom5                   string    `bun:"uom5,type:varchar(8),nullzero"`
	Uom6                   string    `bun:"uom6,type:varchar(8),nullzero"`
	Uom7                   string    `bun:"uom7,type:varchar(8),nullzero"`
	Uom8                   string    `bun:"uom8,type:varchar(8),nullzero"`
	Uom9                   string    `bun:"uom9,type:varchar(8),nullzero"`
	Uom10                  string    `bun:"uom10,type:varchar(8),nullzero"`
	Uom11                  string    `bun:"uom11,type:varchar(8),nullzero"`
	Uom12                  string    `bun:"uom12,type:varchar(8),nullzero"`
	Uom13                  string    `bun:"uom13,type:varchar(8),nullzero"`
	Uom14                  string    `bun:"uom14,type:varchar(8),nullzero"`
	TotalingMethod         string    `bun:"totaling_method,type:varchar(40)"`
	TotalingBasis          string    `bun:"totaling_basis,type:varchar(40)"`
	DeleteFlag             string    `bun:"delete_flag,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CurrencyId             float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	InvMastUid             int32     `bun:"inv_mast_uid,type:int,nullzero"`
	PurchasePricingPageUid int32     `bun:"purchase_pricing_page_uid,type:int,pk"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	LocationId             float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	AllLocationsFlag       string    `bun:"all_locations_flag,type:char,default:('Y')"`
}
