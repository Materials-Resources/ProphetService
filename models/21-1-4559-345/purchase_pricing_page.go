package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePricingPage struct {
	bun.BaseModel          `bun:"table:purchase_pricing_page"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	PricingBookId          string    `bun:"pricing_book_id,type:varchar(8)"`                           // Indicates the associated purchase pricing book.
	SupplierId             float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                   // What supplier is participating in this relationship?
	DiscountGroupId        string    `bun:"discount_group_id,type:varchar(8),nullzero"`                // What is the unique identifier for this discount group?
	MajorGroupId           string    `bun:"major_group_id,type:varchar(8),nullzero"`                   // When the totaling method is major_group, this column can be used to combine items from different pricing pages.
	PricingDescription     string    `bun:"pricing_description,type:varchar(40),nullzero"`             // Description of the purchase pricing page
	PricingMethod          string    `bun:"pricing_method,type:varchar(16)"`                           // Indicates what base price should be used (source, price, or none)
	SourcePrice            string    `bun:"source_price,type:varchar(40),nullzero"`                    // If pricing method is Source, this indicates the price to be used (supplier list or supplie cost)
	Price                  float64   `bun:"price,type:decimal(19,9),nullzero"`                         // If pricing method is Price, this is the base price used by this pricing page
	EffectiveDate          time.Time `bun:"effective_date,type:datetime"`                              // When does this purchase pricing page go into effect
	ExpirationDate         time.Time `bun:"expiration_date,type:datetime"`                             // When does this page expire?
	ContractNumber         string    `bun:"contract_number,type:varchar(40),nullzero"`                 // The contract number for this page
	CalculationType        string    `bun:"calculation_type,type:varchar(16)"`                         // Indicates how the calculation value is used to calculate a price (Difference, Multiplier, Mark-up, Percentage)
	CalculationValue1      float64   `bun:"calculation_value1,type:decimal(19,9)"`                     // First calculation value
	CalculationValue2      float64   `bun:"calculation_value2,type:decimal(19,9)"`                     // Second calculation value
	CalculationValue3      float64   `bun:"calculation_value3,type:decimal(19,9)"`                     // Third calculation value
	CalculationValue4      float64   `bun:"calculation_value4,type:decimal(19,9)"`                     // Fourth calculation value
	CalculationValue5      float64   `bun:"calculation_value5,type:decimal(19,9)"`                     // Fifth calculation value
	CalculationValue6      float64   `bun:"calculation_value6,type:decimal(19,9)"`                     // Sixth calculation value
	CalculationValue7      float64   `bun:"calculation_value7,type:decimal(19,9)"`                     // Seventh calculation value
	CalculationValue8      float64   `bun:"calculation_value8,type:decimal(19,9)"`                     // Eighth calculation value
	CalculationValue9      float64   `bun:"calculation_value9,type:decimal(19,9)"`                     // Ninth calculation value
	CalculationValue10     float64   `bun:"calculation_value10,type:decimal(19,9)"`                    // Tenth calculation value
	CalculationValue11     float64   `bun:"calculation_value11,type:decimal(19,9)"`                    // Eleventh calculation value
	CalculationValue12     float64   `bun:"calculation_value12,type:decimal(19,9)"`                    // Twelveth calculation value
	CalculationValue13     float64   `bun:"calculation_value13,type:decimal(19,9)"`                    // Thirteenth calculation value
	CalculationValue14     float64   `bun:"calculation_value14,type:decimal(19,9)"`                    // Fourteenth calculation value
	CalculationValue15     float64   `bun:"calculation_value15,type:decimal(19,9)"`                    // Fifteenth calculation value
	Break1                 float64   `bun:"break1,type:decimal(19,9)"`                                 // If qty is less than this break, then use calculation_value1
	Break2                 float64   `bun:"break2,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value2
	Break3                 float64   `bun:"break3,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value3
	Break4                 float64   `bun:"break4,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value4
	Break5                 float64   `bun:"break5,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value5
	Break6                 float64   `bun:"break6,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value6
	Break7                 float64   `bun:"break7,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value7
	Break8                 float64   `bun:"break8,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value8
	Break9                 float64   `bun:"break9,type:decimal(19,9)"`                                 // Else if qty is less than this break, then use calculation_value9
	Break10                float64   `bun:"break10,type:decimal(19,9)"`                                // Else if qty is less than this break, then use calculation_value10
	Break11                float64   `bun:"break11,type:decimal(19,9)"`                                // Else if qty is less than this break, then use calculation_value11
	Break12                float64   `bun:"break12,type:decimal(19,9)"`                                // Else if qty is less than this break, then use calculation_value12
	Break13                float64   `bun:"break13,type:decimal(19,9)"`                                // Else if qty is less than this break, then use calculation_value13
	Break14                float64   `bun:"break14,type:decimal(19,9)"`                                // Else if qty is less than this break, then use calculation_value14
	Uom1                   string    `bun:"uom1,type:varchar(8),nullzero"`                             // UOM for break 1
	Uom2                   string    `bun:"uom2,type:varchar(8),nullzero"`                             // UOM for break 2
	Uom3                   string    `bun:"uom3,type:varchar(8),nullzero"`                             // UOM for break 3
	Uom4                   string    `bun:"uom4,type:varchar(8),nullzero"`                             // UOM for break 4
	Uom5                   string    `bun:"uom5,type:varchar(8),nullzero"`                             // UOM for break 5
	Uom6                   string    `bun:"uom6,type:varchar(8),nullzero"`                             // UOM for break 6
	Uom7                   string    `bun:"uom7,type:varchar(8),nullzero"`                             // UOM for break 7
	Uom8                   string    `bun:"uom8,type:varchar(8),nullzero"`                             // UOM for break 8
	Uom9                   string    `bun:"uom9,type:varchar(8),nullzero"`                             // UOM for break 9
	Uom10                  string    `bun:"uom10,type:varchar(8),nullzero"`                            // UOM for break 10
	Uom11                  string    `bun:"uom11,type:varchar(8),nullzero"`                            // UOM for break 11
	Uom12                  string    `bun:"uom12,type:varchar(8),nullzero"`                            // UOM for break 12
	Uom13                  string    `bun:"uom13,type:varchar(8),nullzero"`                            // UOM for break 13
	Uom14                  string    `bun:"uom14,type:varchar(8),nullzero"`                            // UOM for break 14
	TotalingMethod         string    `bun:"totaling_method,type:varchar(40)"`                          // Determines how items priced with this page can be combined to possibly change the price
	TotalingBasis          string    `bun:"totaling_basis,type:varchar(40)"`                           // Determines what quantities are combined when calculating a break (weight, UOM, price, etc.)
	DeleteFlag             string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CurrencyId             float64   `bun:"currency_id,type:decimal(19,0),nullzero"`                   // What is the unique currency identifier for this ro
	InvMastUid             int32     `bun:"inv_mast_uid,type:int,nullzero"`                            // Unique identifier for the item id.
	PurchasePricingPageUid int32     `bun:"purchase_pricing_page_uid,type:int,pk"`                     // Unique UID for this table
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	LocationId             float64   `bun:"location_id,type:decimal(19,0),nullzero"`       // column indicates the location of the pricing page
	AllLocationsFlag       string    `bun:"all_locations_flag,type:char(1),default:('Y')"` // Allow to create purchase pricing page without location id
}
