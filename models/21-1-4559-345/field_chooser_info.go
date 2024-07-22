package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FieldChooserInfo struct {
	bun.BaseModel         `bun:"table:field_chooser_info"`
	FieldChooserInfoUid   int32     `bun:"field_chooser_info_uid,type:int,autoincrement,identity,pk"`    // UID for table
	Area                  int32     `bun:"area,type:int,unique"`                                         // Which area this is being used in - Order Lines, order header, etc.
	FieldName             string    `bun:"field_name,type:varchar(255),unique"`                          // Actual name of the field
	DisplayName           *string   `bun:"display_name,type:varchar(255)"`                               // Display name of the field
	Editable              *string   `bun:"editable,type:char(1)"`                                        // Set to Y if editable, N if not editable
	FieldTypeCd           *int32    `bun:"field_type_cd,type:int"`                                       // Whether this field should be a Checkbox, dddw, Edit, Editmask, RadioButtons
	FieldInfo             *string   `bun:"field_info,type:varchar(255)"`                                 // Information related to the field type - editmask, code group for the dddw
	ExtraText             *string   `bun:"extra_text,type:varchar(255)"`                                 // Required Column Indicator, calendar indicator, other information
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ConfigurationId       *int32    `bun:"configuration_id,type:int,unique"`                             // Customer configuration id for customization.
	AutoPopulateFlag      string    `bun:"auto_populate_flag,type:char(1),default:('N')"`                // Indicate whether app should auto-populate value
	ValueSourceTable      *string   `bun:"value_source_table,type:varchar(255)"`                         // Table from which auto-population will get the value
	ValueSourceColumn     *string   `bun:"value_source_column,type:varchar(255)"`                        // Column from which auto-population will get the value
	ValueSourceExpression *string   `bun:"value_source_expression,type:varchar(8000)"`                   // SQL expression for computed values in auto-population
}
