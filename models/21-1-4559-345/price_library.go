package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceLibrary struct {
	bun.BaseModel                 `bun:"table:price_library"`
	PriceLibraryUid               int32     `bun:"price_library_uid,type:int,pk"`            // Internal record identifier
	PriceLibraryId                string    `bun:"price_library_id,type:varchar(20),unique"` // User-defined Sales Pricing Library identifier
	Description                   string    `bun:"description,type:varchar(255)"`            // Sales Pricing Library Description
	CategoryCd                    int16     `bun:"category_cd,type:smallint"`                // Indicates that this is a Sales Pricing Library
	TypeCd                        int16     `bun:"type_cd,type:smallint"`                    // Indicates the type of this Sales Pricing Library (First Of, Lowest Of, Multipler, etc.)
	SourcePriceCd                 *int16    `bun:"source_price_cd,type:smallint"`            // Code value for the Source Price if the Sales Pricing Library is of type Multiplier
	Multiplier                    *float64  `bun:"multiplier,type:decimal(19,9)"`            // Multiplier value if the Sales Pricing Library is of type Multiplier
	RowStatusFlag                 int16     `bun:"row_status_flag,type:smallint"`            // Indicates current record status.
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime"`         // Indicates the date/time this record was last modified.
	DateCreated                   time.Time `bun:"date_created,type:datetime"`               // Indicates the date/time this record was created.
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30)"`      // ID of the user who last maintained this record
	CreatedBy                     *string   `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TerminalId                    *float64  `bun:"terminal_id,type:decimal(19,0)"`                          // (Custom) The terminal ID used to retrieve this price for this item
	CompanyId                     *string   `bun:"company_id,type:varchar(255)"`                            // Company that this library is for.
	LibraryOnContract             string    `bun:"library_on_contract,type:char(1),default:('N')"`          // ON/OFF contract pricing for Multipler Pricing Libraries.
	StrategicPriceLibraryFlag     string    `bun:"strategic_price_library_flag,type:char(1),default:('N')"` // Library is considered a strategic pricing library if this is set; otherwise it is considered to be a standard library
	CustomerSizeCd                *int32    `bun:"customer_size_cd,type:int"`                               // Strategic Pricing Customer Size - Very Tiny, Tiny, Small, Medium, Large, or Huge
	CustomerCategoryUid           *int32    `bun:"customer_category_uid,type:int"`                          // Strategic Pricing Customer Category
	ExcludeFromCustMultiplierFlag *string   `bun:"exclude_from_cust_multiplier_flag,type:char(1)"`          // Indicate the new pricing multiplier in Customer Maintenance should NOT be applied against this library. Custom field.
}
