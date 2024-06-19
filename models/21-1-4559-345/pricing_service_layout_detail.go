package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceLayoutDetail struct {
	bun.BaseModel       `bun:"table:pricing_service_layout_detail"`
	LayoutId            float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout, identifies with which layout the column information belongs
	LayoutDetailId      float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`                    // Unique identifier within the pricing service layout for each column definition
	ColumnName          string    `bun:"column_name,type:varchar(50),nullzero"`                     // Name of the database column to be updated
	StartPosition       float64   `bun:"start_position,type:decimal(19,0),nullzero"`                // Within fixed-format layouts, this indicates the offset of the column data
	DataPrecision       float64   `bun:"data_precision,type:decimal(19,0),nullzero"`                // Within fixed-formation layouts, this indicates the length of the data; not used for tab-delimited layouts
	DataType            string    `bun:"data_type,type:varchar(10),nullzero"`                       // Code for the type of data: Alphanumeric, numeric.
	DecimalScale        float64   `bun:"decimal_scale,type:decimal(19,0),nullzero"`                 // For numeric columns, the number of decimal places in the source data
	AcsType             string    `bun:"acs_type,type:varchar(10),nullzero"`                        // Not used
	ExtensionInd        string    `bun:"extension_ind,type:char(1),nullzero"`                       // Indicator of whether an extension rule applies to this column
	ValueInd            string    `bun:"value_ind,type:char(1),nullzero"`                           // Indicator of whether a value rule applies to this column
	ValueId             float64   `bun:"value_id,type:decimal(19,0),nullzero"`                      // Foreign key to pricing_service_value table
	ConversionInd       string    `bun:"conversion_ind,type:char(1),nullzero"`                      // Indicator of whether a conversion rule applies to this column
	FilterInd           string    `bun:"filter_ind,type:char(1),nullzero"`                          // Indicator of whether a filter rule applies to this column
	DeleteFlag          string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated         time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DataLength          float64   `bun:"data_length,type:decimal(19,0),nullzero"`                   // The length of the source data for this column
	RequiredInd         string    `bun:"required_ind,type:char(1),nullzero"`                        // Indicator of whether the column is required for the current layout type
	DefaultInd          string    `bun:"default_ind,type:char(1),nullzero"`                         // Indicator of whether the import process should use the item default value for this column
	NewItemsOnly        string    `bun:"new_items_only,type:char(1),nullzero"`                      // Indicator of whether the import process should apply the value to new items only
	DateMask            string    `bun:"date_mask,type:varchar(255),nullzero"`                      // Format of a date field in a pricing service source file
	StartPositionDate   time.Time `bun:"start_position_date,type:datetime,default:(getdate())"`     // Date when the column's start position changed
	SetByOption         string    `bun:"set_by_option,type:char(1),nullzero"`                       // Indicator whether to update column using Activant Pricing Service Layout option
	ImportNewValuesFlag string    `bun:"import_new_values_flag,type:char(1),nullzero"`              // Determines if this column will import new values during a build/update.
	ItemCatalogDefUid   int32     `bun:"item_catalog_def_uid,type:int,nullzero"`                    // Identifies item_catalog column which is the source of the import data
}
