package model

import (
	"github.com/uptrace/bun"
	"time"
)

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
	OnStartLoad                         bool      `bun:"on_start_load,type:bit,nullzero"`
	AutoFiltering                       bool      `bun:"auto_filtering,type:bit,nullzero"`
	FocusOnGrid                         bool      `bun:"focus_on_grid,type:bit,nullzero"`
	ColumnsResizable                    bool      `bun:"columns_resizable,type:bit,nullzero"`
	ExportEnabled                       bool      `bun:"export_enabled,type:bit,nullzero"`
	GroupingEnabled                     bool      `bun:"grouping_enabled,type:bit,nullzero"`
	ColumnFilteringEnabled              bool      `bun:"column_filtering_enabled,type:bit,nullzero"`
	CountRetrievalEnabled               bool      `bun:"count_retrieval_enabled,type:bit,nullzero"`
	GenericSearchEnabled                bool      `bun:"generic_search_enabled,type:bit,nullzero"`
	MultipleRowSelectionEnabled         bool      `bun:"multiple_row_selection_enabled,type:bit,nullzero"`
	UserConfigurationEnabled            bool      `bun:"user_configuration_enabled,type:bit,nullzero"`
	CoreFlag                            bool      `bun:"core_flag,type:bit,nullzero"`
	DeleteFlag                          bool      `bun:"delete_flag,type:bit,nullzero"`
	InactiveFlag                        bool      `bun:"inactive_flag,type:bit,nullzero"`
	AdditionalWhereEnabled              bool      `bun:"additional_where_enabled,type:bit,nullzero"`
	PopupAttribute                      int32     `bun:"popup_attribute,type:int,nullzero"`
	RankingEnabled                      bool      `bun:"ranking_enabled,type:bit,nullzero"`
	Source                              string    `bun:"source,type:varchar(500),nullzero"`
	ExpandablePopup                     string    `bun:"expandable_popup,type:varchar(50),nullzero"`
	CompSecEnabled                      bool      `bun:"comp_sec_enabled,type:bit,nullzero"`
	OverridePopupDesc                   bool      `bun:"override_popup_desc,type:bit,default:((0))"`
	OverridePopupWidth                  bool      `bun:"override_popup_width,type:bit,default:((0))"`
	OverridePopupHeight                 bool      `bun:"override_popup_height,type:bit,default:((0))"`
	OverridePageSize                    bool      `bun:"override_page_size,type:bit,default:((0))"`
	OverrideMinimumChars                bool      `bun:"override_minimum_chars,type:bit,default:((0))"`
	OverrideOnStartLoad                 bool      `bun:"override_on_start_load,type:bit,default:((0))"`
	OverrideAutoFiltering               bool      `bun:"override_auto_filtering,type:bit,default:((0))"`
	OverrideFocusOnGrid                 bool      `bun:"override_focus_on_grid,type:bit,default:((0))"`
	OverrideColumnsResizable            bool      `bun:"override_columns_resizable,type:bit,default:((0))"`
	OverrideExportEnabled               bool      `bun:"override_export_enabled,type:bit,default:((0))"`
	OverrideGroupingEnabled             bool      `bun:"override_grouping_enabled,type:bit,default:((0))"`
	OverrideColumnFilteringEnabled      bool      `bun:"override_column_filtering_enabled,type:bit,default:((0))"`
	OverrideCountRetrievalEnabled       bool      `bun:"override_count_retrieval_enabled,type:bit,default:((0))"`
	OverrideGenericSearchEnabled        bool      `bun:"override_generic_search_enabled,type:bit,default:((0))"`
	OverrideMultipleRowSelectionEnabled bool      `bun:"override_multiple_row_selection_enabled,type:bit,default:((0))"`
	OverrideUserConfigurationEnabled    bool      `bun:"override_user_configuration_enabled,type:bit,default:((0))"`
	OverrideInactiveFlag                bool      `bun:"override_inactive_flag,type:bit,default:((0))"`
	OverrideAdditionalWhereEnabled      bool      `bun:"override_additional_where_enabled,type:bit,default:((0))"`
	OverrideRankingEnabled              bool      `bun:"override_ranking_enabled,type:bit,default:((0))"`
	OverrideSource                      bool      `bun:"override_source,type:bit,default:((0))"`
	OverrideExpandablePopup             bool      `bun:"override_expandable_popup,type:bit,default:((0))"`
	OverrideCompSecEnabled              bool      `bun:"override_comp_sec_enabled,type:bit,default:((0))"`
	PopupDetailUidParent                int32     `bun:"popup_detail_uid_parent,type:int,nullzero"`
	OverrideTitle                       bool      `bun:"override_title,type:bit,default:((0))"`
}
