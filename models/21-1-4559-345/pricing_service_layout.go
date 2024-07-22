package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceLayout struct {
	bun.BaseModel           `bun:"table:pricing_service_layout"`
	LayoutId                float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Unique identifier for a pricing service layout definition
	LayoutType              *string   `bun:"layout_type,type:varchar(10)"`                              // Code indicating whether the layout is intended to build new items, update existing ones, or a combination
	LayoutDesc              *string   `bun:"layout_desc,type:varchar(50)"`                              // User-defined description of the layout definition
	LayoutKey               *string   `bun:"layout_key,type:varchar(10)"`                               // Which item table column should be used to identify whether the item already exists
	LayoutPath              *string   `bun:"layout_path,type:varchar(255)"`                             // The path or the import source file
	LayoutFilename          *string   `bun:"layout_filename,type:varchar(50)"`                          // The name of the import source file
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LayoutTabletype         *string   `bun:"layout_tabletype,type:varchar(10)"`                         // Used in the ODBC driver definition, this code identifies whether the import file is fixed-format or tab-delimited.
	LayoutDelimiter         *string   `bun:"layout_delimiter,type:char(1)"`                             // Used in the ODBC driver definition.
	LayoutFirstlinenames    *string   `bun:"layout_firstlinenames,type:char(1)"`                        // Used in the ODBC driver interface, this code identifies whether the first record is column names or data
	CompanyId               string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	LocationId              float64   `bun:"location_id,type:decimal(19,0)"`                            // Determines within which location the pricing service changes should be applied
	AddLocation             *string   `bun:"add_location,type:char(1)"`                                 // Add new location to existing items (Y/N)
	AddSupplier             *string   `bun:"add_supplier,type:char(1)"`                                 // Add new supplier to existing items (Y/N)
	TpcxFormatCd            int32     `bun:"tpcx_format_cd,type:int,default:(0)"`                       // Identifies layout as using TPCx abbreviated format, full format, or not TPCx
	PurchaseTransferGroupId *string   `bun:"purchase_transfer_group_id,type:varchar(8)"`                // purchase transfer group for multiple locations
	PricingTemplateId       *string   `bun:"pricing_template_id,type:varchar(255)"`                     // The Pricing Template that this layout will use to default values.
	CatalogStatusFlag       string    `bun:"catalog_status_flag,type:char(1),default:('N')"`            // Indicates whether item_catalog definition has been changed
	DataSourceCd            int32     `bun:"data_source_cd,type:int,default:((2564))"`                  // Source of pricing service import: file or catalog
	DataDestinationCd       int32     `bun:"data_destination_cd,type:int,default:((2566))"`             // Target of pricing service import: catalog, items, both
	BuildInBothFlag         string    `bun:"build_in_both_flag,type:char(1),default:('N')"`             // Indicates whether new items added to both destinations
	ApplyValueListFirstFlag string    `bun:"apply_value_list_first_flag,type:char(1),default:('Y')"`    // Indicates whethe value lists are applied first or before extension rules, etc.
	TempDataStorageCd       int32     `bun:"temp_data_storage_cd,type:int,default:((2570))"`            // Temporary Data Storage Type
}
