package model

type PopupDetail struct {
	bun.BaseModel                       `bun:"table:popup_detail"`
	PopupDetailUid                      int32     `bun:"popup_detail_uid,type:int,pk,identity"`
	PopupTitle                          string    `bun:"popup_title,type:varchar(40)"`
	PopupDesc                           string    `bun:"popup_desc,type:varchar(255)"`
	RestrictHdlr                        string    `bun:"restrict_hdlr,type:varchar(32)"`
	PopupWidth                          int16     `bun:"popup_width,type:smallint"`
	PopupHeight                         int16     `bun:"popup_height,type:smallint"`
	PageSize                            int16     `bun:"page_size,type:smallint"`
	Literal                             int16     `bun:"literal,type:smallint"`
	MinimumChars                        int16     `bun:"minimum_chars,type:smallint"`
	DateCreated                         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OnStartLoad                         `bun:"on_start_load,type:bit,nullzero"`
	AutoFiltering                       `bun:"auto_filtering,type:bit,nullzero"`
	FocusOnGrid                         `bun:"focus_on_grid,type:bit,nullzero"`
	ColumnsResizable                    `bun:"columns_resizable,type:bit,nullzero"`
	ExportEnabled                       `bun:"export_enabled,type:bit,nullzero"`
	GroupingEnabled                     `bun:"grouping_enabled,type:bit,nullzero"`
	ColumnFilteringEnabled              `bun:"column_filtering_enabled,type:bit,nullzero"`
	CountRetrievalEnabled               `bun:"count_retrieval_enabled,type:bit,nullzero"`
	GenericSearchEnabled                `bun:"generic_search_enabled,type:bit,nullzero"`
	MultipleRowSelectionEnabled         `bun:"multiple_row_selection_enabled,type:bit,nullzero"`
	UserConfigurationEnabled            `bun:"user_configuration_enabled,type:bit,nullzero"`
	CoreFlag                            `bun:"core_flag,type:bit,nullzero"`
	DeleteFlag                          `bun:"delete_flag,type:bit,nullzero"`
	InactiveFlag                        `bun:"inactive_flag,type:bit,nullzero"`
	AdditionalWhereEnabled              `bun:"additional_where_enabled,type:bit,nullzero"`
	PopupAttribute                      int32 `bun:"popup_attribute,type:int,nullzero"`
	RankingEnabled                      `bun:"ranking_enabled,type:bit,nullzero"`
	Source                              string `bun:"source,type:varchar(500),nullzero"`
	ExpandablePopup                     string `bun:"expandable_popup,type:varchar(50),nullzero"`
	CompSecEnabled                      `bun:"comp_sec_enabled,type:bit,nullzero"`
	OverridePopupDesc                   `bun:"override_popup_desc,type:bit,default:((0))"`
	OverridePopupWidth                  `bun:"override_popup_width,type:bit,default:((0))"`
	OverridePopupHeight                 `bun:"override_popup_height,type:bit,default:((0))"`
	OverridePageSize                    `bun:"override_page_size,type:bit,default:((0))"`
	OverrideMinimumChars                `bun:"override_minimum_chars,type:bit,default:((0))"`
	OverrideOnStartLoad                 `bun:"override_on_start_load,type:bit,default:((0))"`
	OverrideAutoFiltering               `bun:"override_auto_filtering,type:bit,default:((0))"`
	OverrideFocusOnGrid                 `bun:"override_focus_on_grid,type:bit,default:((0))"`
	OverrideColumnsResizable            `bun:"override_columns_resizable,type:bit,default:((0))"`
	OverrideExportEnabled               `bun:"override_export_enabled,type:bit,default:((0))"`
	OverrideGroupingEnabled             `bun:"override_grouping_enabled,type:bit,default:((0))"`
	OverrideColumnFilteringEnabled      `bun:"override_column_filtering_enabled,type:bit,default:((0))"`
	OverrideCountRetrievalEnabled       `bun:"override_count_retrieval_enabled,type:bit,default:((0))"`
	OverrideGenericSearchEnabled        `bun:"override_generic_search_enabled,type:bit,default:((0))"`
	OverrideMultipleRowSelectionEnabled `bun:"override_multiple_row_selection_enabled,type:bit,default:((0))"`
	OverrideUserConfigurationEnabled    `bun:"override_user_configuration_enabled,type:bit,default:((0))"`
	OverrideInactiveFlag                `bun:"override_inactive_flag,type:bit,default:((0))"`
	OverrideAdditionalWhereEnabled      `bun:"override_additional_where_enabled,type:bit,default:((0))"`
	OverrideRankingEnabled              `bun:"override_ranking_enabled,type:bit,default:((0))"`
	OverrideSource                      `bun:"override_source,type:bit,default:((0))"`
	OverrideExpandablePopup             `bun:"override_expandable_popup,type:bit,default:((0))"`
	OverrideCompSecEnabled              `bun:"override_comp_sec_enabled,type:bit,default:((0))"`
	PopupDetailUidParent                int32 `bun:"popup_detail_uid_parent,type:int,nullzero"`
	OverrideTitle                       `bun:"override_title,type:bit,default:((0))"`
}
