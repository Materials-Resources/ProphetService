package model

type PopupField struct {
	bun.BaseModel                    `bun:"table:popup_field"`
	PopupFieldUid                    int32  `bun:"popup_field_uid,type:int,pk,identity"`
	FieldName                        string `bun:"field_name,type:varchar(99)"`
	Header                           string `bun:"header,type:varchar(50),nullzero"`
	FieldType                        int16  `bun:"field_type,type:smallint"`
	FieldWidth                       int16  `bun:"field_width,type:smallint"`
	Mask                             string `bun:"mask,type:varchar(99),nullzero"`
	ReturnTo                         string `bun:"return_to,type:varchar(50),nullzero"`
	OnGrid                           `bun:"on_grid,type:bit,nullzero"`
	AsParameter                      `bun:"as_parameter,type:bit,nullzero"`
	Filter                           `bun:"filter,type:bit,nullzero"`
	Enabled                          `bun:"enabled,type:bit,nullzero"`
	PopupToShow                      string    `bun:"popup_to_show,type:varchar(50),nullzero"`
	DateCreated                      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                 time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                 string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	PopupDetailUid                   int32     `bun:"popup_detail_uid,type:int"`
	ParameterCodeGroup               string    `bun:"parameter_code_group,type:varchar(255),nullzero"`
	ParameterDefaultValue            string    `bun:"parameter_default_value,type:varchar(255),nullzero"`
	ParameterHeader                  string    `bun:"parameter_header,type:varchar(50),nullzero"`
	ParameterWidth                   int32     `bun:"parameter_width,type:int,nullzero"`
	ParameterEnabled                 `bun:"parameter_enabled,type:bit,nullzero"`
	ParameterFilter                  `bun:"parameter_filter,type:bit,nullzero"`
	ParameterType                    int32  `bun:"parameter_type,type:int,nullzero"`
	ParameterLocationX               int32  `bun:"parameter_location_x,type:int,nullzero"`
	ParameterLocationY               int32  `bun:"parameter_location_y,type:int,nullzero"`
	Rank                             int32  `bun:"rank,type:int,nullzero"`
	ParameterHiddenCondition         string `bun:"parameter_hidden_condition,type:varchar(1000),nullzero"`
	AliasDefinition                  string `bun:"alias_definition,type:varchar(5000),nullzero"`
	FieldExpression                  string `bun:"field_expression,type:varchar(5000),nullzero"`
	DataType                         string `bun:"data_type,type:varchar(255),nullzero"`
	Hidden                           `bun:"hidden,type:bit,default:((0))"`
	IsPinned                         `bun:"is_pinned,type:bit,default:((0))"`
	ParameterServerFilter            `bun:"parameter_server_filter,type:bit,default:((0))"`
	Position                         int32 `bun:"position,type:int,default:((0))"`
	OverrideReturnTo                 `bun:"override_return_to,type:bit,default:((0))"`
	OverrideOnGrid                   `bun:"override_on_grid,type:bit,default:((0))"`
	OverrideAsParameter              `bun:"override_as_parameter,type:bit,default:((0))"`
	OverrideFilter                   `bun:"override_filter,type:bit,default:((0))"`
	OverrideEnabled                  `bun:"override_enabled,type:bit,default:((0))"`
	OverridePopupToShow              `bun:"override_popup_to_show,type:bit,default:((0))"`
	OverrideParameterHeader          `bun:"override_parameter_header,type:bit,default:((0))"`
	OverrideParameterWidth           `bun:"override_parameter_width,type:bit,default:((0))"`
	OverrideParameterEnabled         `bun:"override_parameter_enabled,type:bit,default:((0))"`
	OverrideParameterFilter          `bun:"override_parameter_filter,type:bit,default:((0))"`
	OverrideParameterType            `bun:"override_parameter_type,type:bit,default:((0))"`
	OverrideParameterLocationX       `bun:"override_parameter_location_x,type:bit,default:((0))"`
	OverrideParameterLocationY       `bun:"override_parameter_location_y,type:bit,default:((0))"`
	OverrideParameterCodeGroup       `bun:"override_parameter_code_group,type:bit,default:((0))"`
	OverrideParameterDefaultValue    `bun:"override_parameter_default_value,type:bit,default:((0))"`
	OverrideRank                     `bun:"override_rank,type:bit,default:((0))"`
	OverrideParameterHiddenCondition `bun:"override_parameter_hidden_condition,type:bit,default:((0))"`
	OverrideAliasDefinition          `bun:"override_alias_definition,type:bit,default:((0))"`
	OverrideFieldExpression          `bun:"override_field_expression,type:bit,default:((0))"`
	OverrideDataType                 `bun:"override_data_type,type:bit,default:((0))"`
	OverrideHidden                   `bun:"override_hidden,type:bit,default:((0))"`
	OverrideIsPinned                 `bun:"override_is_pinned,type:bit,default:((0))"`
	OverrideParameterServerFilter    `bun:"override_parameter_server_filter,type:bit,default:((0))"`
	OverridePosition                 `bun:"override_position,type:bit,default:((0))"`
	OverrideHeader                   `bun:"override_header,type:bit,default:((0))"`
	OverrideFieldType                `bun:"override_field_type,type:bit,default:((0))"`
	OverrideFieldWidth               `bun:"override_field_width,type:bit,default:((0))"`
	OverrideMask                     `bun:"override_mask,type:bit,default:((0))"`
	OverrideFieldName                `bun:"override_field_name,type:bit,default:((0))"`
	PopupFieldUidParent              int32 `bun:"popup_field_uid_parent,type:int,nullzero"`
	PopupDetailUidParent             int32 `bun:"popup_detail_uid_parent,type:int,nullzero"`
	DeleteFlag                       `bun:"delete_flag,type:bit,default:((0))"`
}
