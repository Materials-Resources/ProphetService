package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceColumn struct {
	bun.BaseModel           `bun:"table:pricing_service_column"`
	ColumnId                float64    `bun:"column_id,type:decimal(19,0),pk"`                               // Unique identifier for record
	ColumnName              string     `bun:"column_name,type:varchar(40)"`                                  // Database column name
	ColumnDesc              *string    `bun:"column_desc,type:varchar(255)"`                                 // Description of the column name
	AlternateColumnName     *string    `bun:"alternate_column_name,type:varchar(40)"`                        // Column name in other table
	InventoryDefaultsName   *string    `bun:"inventory_defaults_name,type:varchar(40)"`                      // Column name in the inventory (item) defaults table
	RequiredInd             *string    `bun:"required_ind,type:char(1)"`                                     // Indicates whether the column is required for a Build pricing service layout
	DefaultInd              *string    `bun:"default_ind,type:char(1)"`                                      // Indicates wehther the column default value is available in the item default table
	DeleteFlag              string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated             time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DataType                *string    `bun:"data_type,type:varchar(10)"`                                    // Data type of the database column
	ColumnSize              *float64   `bun:"column_size,type:decimal(19,0)"`                                // Size of the database column
	ValidationRule          string     `bun:"validation_rule,type:char(1),default:('N')"`                    // Indicates whether a business rule should be triggered for the column during the import process.
	TpcxColumnSequence      int32      `bun:"tpcx_column_sequence,type:int,default:(0)"`                     // Default sequence number in the pricing service layout.
	TpcxOnly                string     `bun:"tpcx_only,type:char(1),default:('N')"`                          // Indicates whether the column is to be made available in pricing service layout design.
	TpcxAbbrevColumnSeq     int32      `bun:"tpcx_abbrev_column_seq,type:int,default:(0)"`                   // TPCx Abbreviated format column sequence
	TpcxAbbrevColumnName    *string    `bun:"tpcx_abbrev_column_name,type:varchar(40)"`                      // TPCx Abbreviated format column name
	TpcxAbbrevDate          *time.Time `bun:"tpcx_abbrev_date,type:datetime"`                                // Date sequence defined for TPCx abbreviated format
	TpcxAbbrevAddForUpdate  string     `bun:"tpcx_abbrev_add_for_update,type:char(1),default:('N')"`         // Whether to add column for Update Layout
	ActLayoutOnly           string     `bun:"act_layout_only,type:char(1),default:('N')"`                    // Indicator for Activant Pricing Service layout
	ActLayoutColumnSeq      int32      `bun:"act_layout_column_seq,type:int,default:(0)"`                    // Column sequence number on Activant Pricing Service layout file
	ActLayoutColumnName     *string    `bun:"act_layout_column_name,type:varchar(40)"`                       // Activant Pricing Service layout file's column name
	ActLayoutDate           *time.Time `bun:"act_layout_date,type:datetime"`                                 // Activant Pricing Service layout date
	ActLayoutAddForUpdate   string     `bun:"act_layout_add_for_update,type:char(1),default:('N')"`          // Activant Pricing Service layout update flag
	AqnetLayoutOnly         string     `bun:"aqnet_layout_only,type:char(1),default:('N')"`                  // Indicates whether column applies only to AQNet layouts
	AqnetLayoutColumnSeq    int32      `bun:"aqnet_layout_column_seq,type:int,default:((0))"`                // Column sequence number in the import file
	AqnetLayoutColumnName   *string    `bun:"aqnet_layout_column_name,type:varchar(40)"`                     // Column name used by AQNet
	AqnetLayoutDate         *time.Time `bun:"aqnet_layout_date,type:datetime"`                               // Date the column_name was added to the support table
	AqnetLayoutAddForUpdate string     `bun:"aqnet_layout_add_for_update,type:char(1),default:('N')"`        // Whether column should be added to Update layout
	PricingTemplateDfltType string     `bun:"pricing_template_dflt_type,type:char(1),default:('N')"`         // Determines whether this column appears in Pricing Service Template Maintenance.
	ImportNewValuesFlag     string     `bun:"import_new_values_flag,type:char(1),default:('N')"`             // Determines if this column will import new values during a build/update.
}
