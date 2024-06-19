package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupField struct {
	bun.BaseModel                    `bun:"table:popup_field"`
	PopupFieldUid                    int32     `bun:"popup_field_uid,type:int,autoincrement,identity,pk"`           // Popup Field Id
	FieldName                        string    `bun:"field_name,type:varchar(99)"`                                  // Field name
	Header                           string    `bun:"header,type:varchar(50),nullzero"`                             // Header field name on window
	FieldType                        int16     `bun:"field_type,type:smallint"`                                     // Type
	FieldWidth                       int16     `bun:"field_width,type:smallint"`                                    // Field width
	Mask                             string    `bun:"mask,type:varchar(99),nullzero"`                               // Field mask
	ReturnTo                         string    `bun:"return_to,type:varchar(50),nullzero"`                          // Return to
	OnGrid                           bool      `bun:"on_grid,type:bit,nullzero"`                                    // Is field showed in grid
	AsParameter                      bool      `bun:"as_parameter,type:bit,nullzero"`                               // Filled with parameter value
	Filter                           bool      `bun:"filter,type:bit,nullzero"`                                     // Can use this filed to filter
	Enabled                          bool      `bun:"enabled,type:bit,nullzero"`                                    // Is enabled
	PopupToShow                      string    `bun:"popup_to_show,type:varchar(50),nullzero"`                      // Has secondary popup
	DateCreated                      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified                 time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy                 string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	PopupDetailUid                   int32     `bun:"popup_detail_uid,type:int"`                                    // Popup Index Id
	ParameterCodeGroup               string    `bun:"parameter_code_group,type:varchar(255),nullzero"`              // Code used by the parameter
	ParameterDefaultValue            string    `bun:"parameter_default_value,type:varchar(255),nullzero"`           // Default value of the parameter
	ParameterHeader                  string    `bun:"parameter_header,type:varchar(50),nullzero"`                   // Header of the parameter control
	ParameterWidth                   int32     `bun:"parameter_width,type:int,nullzero"`                            // Width of the parameter control
	ParameterEnabled                 bool      `bun:"parameter_enabled,type:bit,nullzero"`                          // If the parameter control is enabled
	ParameterFilter                  bool      `bun:"parameter_filter,type:bit,nullzero"`                           // If the parameter control should apply filters
	ParameterType                    int32     `bun:"parameter_type,type:int,nullzero"`                             // The type of the parameter control
	ParameterLocationX               int32     `bun:"parameter_location_x,type:int,nullzero"`                       // X axis of the location of a parameter
	ParameterLocationY               int32     `bun:"parameter_location_y,type:int,nullzero"`                       // Y axis of the location of a parameter
	Rank                             int32     `bun:"rank,type:int,nullzero"`                                       // Determines the rank value of the column
	ParameterHiddenCondition         string    `bun:"parameter_hidden_condition,type:varchar(1000),nullzero"`       // Sets a condition on which a parameter controls must be hidden
	AliasDefinition                  string    `bun:"alias_definition,type:varchar(5000),nullzero"`                 // The definition of a column alias associated to this popup field
	FieldExpression                  string    `bun:"field_expression,type:varchar(5000),nullzero"`                 // Stores the expression of a column  used in run time
	DataType                         string    `bun:"data_type,type:varchar(255),nullzero"`                         // Type of the data of the column that is related to the field.
	Hidden                           bool      `bun:"hidden,type:bit,default:((0))"`                                // If is enabled the field is hidden
	IsPinned                         bool      `bun:"is_pinned,type:bit,default:((0))"`                             // Saves if a column must be pinned as default
	ParameterServerFilter            bool      `bun:"parameter_server_filter,type:bit,default:((0))"`               // Determines if the parameter should be applied on the server side
	Position                         int32     `bun:"position,type:int,default:((0))"`                              // Determinates the position of the column in the grid
	OverrideReturnTo                 bool      `bun:"override_return_to,type:bit,default:((0))"`                    // flag to override return to for Dynachange
	OverrideOnGrid                   bool      `bun:"override_on_grid,type:bit,default:((0))"`                      // flag to override on grid for Dynachange
	OverrideAsParameter              bool      `bun:"override_as_parameter,type:bit,default:((0))"`                 // flag to override as parameter for Dynachange
	OverrideFilter                   bool      `bun:"override_filter,type:bit,default:((0))"`                       // flag to override filter for Dynachange
	OverrideEnabled                  bool      `bun:"override_enabled,type:bit,default:((0))"`                      // flag to override enabled for Dynachange
	OverridePopupToShow              bool      `bun:"override_popup_to_show,type:bit,default:((0))"`                // flag to override popup to show for Dynachange
	OverrideParameterHeader          bool      `bun:"override_parameter_header,type:bit,default:((0))"`             // flag to override parameter header for Dynachange
	OverrideParameterWidth           bool      `bun:"override_parameter_width,type:bit,default:((0))"`              // flag to override parameter width for Dynachange
	OverrideParameterEnabled         bool      `bun:"override_parameter_enabled,type:bit,default:((0))"`            // flag to override parameter enabled for Dynachange
	OverrideParameterFilter          bool      `bun:"override_parameter_filter,type:bit,default:((0))"`             // flag to override parameter filter for Dynachange
	OverrideParameterType            bool      `bun:"override_parameter_type,type:bit,default:((0))"`               // flag to override parameter type for Dynachange
	OverrideParameterLocationX       bool      `bun:"override_parameter_location_x,type:bit,default:((0))"`         // flag to override parameter location x for Dynachange
	OverrideParameterLocationY       bool      `bun:"override_parameter_location_y,type:bit,default:((0))"`         // flag to override parameter location y for Dynachange
	OverrideParameterCodeGroup       bool      `bun:"override_parameter_code_group,type:bit,default:((0))"`         // flag to override parameter code group for Dynachange
	OverrideParameterDefaultValue    bool      `bun:"override_parameter_default_value,type:bit,default:((0))"`      // flag to override parameter default value for Dynachange
	OverrideRank                     bool      `bun:"override_rank,type:bit,default:((0))"`                         // flag to override rank for Dynachange
	OverrideParameterHiddenCondition bool      `bun:"override_parameter_hidden_condition,type:bit,default:((0))"`   // flag to override parameter hidden condition for Dynachange
	OverrideAliasDefinition          bool      `bun:"override_alias_definition,type:bit,default:((0))"`             // flag to override alias definition for Dynachange
	OverrideFieldExpression          bool      `bun:"override_field_expression,type:bit,default:((0))"`             // flag to override field expression for Dynachange
	OverrideDataType                 bool      `bun:"override_data_type,type:bit,default:((0))"`                    // flag to override data type for Dynachange
	OverrideHidden                   bool      `bun:"override_hidden,type:bit,default:((0))"`                       // flag to override hidden for Dynachange
	OverrideIsPinned                 bool      `bun:"override_is_pinned,type:bit,default:((0))"`                    // flag to override is pinned for Dynachange
	OverrideParameterServerFilter    bool      `bun:"override_parameter_server_filter,type:bit,default:((0))"`      // flag to override parameter server filter for Dynachange
	OverridePosition                 bool      `bun:"override_position,type:bit,default:((0))"`                     // flag to override position for Dynachange
	OverrideHeader                   bool      `bun:"override_header,type:bit,default:((0))"`                       // flag to override header for Dynachange
	OverrideFieldType                bool      `bun:"override_field_type,type:bit,default:((0))"`                   // flag to override field type for Dynachange
	OverrideFieldWidth               bool      `bun:"override_field_width,type:bit,default:((0))"`                  // flag to override field width for Dynachange
	OverrideMask                     bool      `bun:"override_mask,type:bit,default:((0))"`                         // flag to override mask for Dynachange
	OverrideFieldName                bool      `bun:"override_field_name,type:bit,default:((0))"`                   // flag to override field name for Dynachange
	PopupFieldUidParent              int32     `bun:"popup_field_uid_parent,type:int,nullzero"`                     // Popup field Uid of the parent to manage the additive
	PopupDetailUidParent             int32     `bun:"popup_detail_uid_parent,type:int,nullzero"`                    // Popup detail Uid of the parent to manage the additive
	DeleteFlag                       bool      `bun:"delete_flag,type:bit,default:((0))"`                           // Flag that determinated if the current field is deleted
}
