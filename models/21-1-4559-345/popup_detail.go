package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupDetail struct {
	bun.BaseModel                       `bun:"table:popup_detail"`
	PopupDetailUid                      int32     `bun:"popup_detail_uid,type:int,autoincrement,identity,pk"`            // Popup Detail Id
	PopupTitle                          string    `bun:"popup_title,type:varchar(40)"`                                   // Title
	PopupDesc                           string    `bun:"popup_desc,type:varchar(255)"`                                   // Description
	RestrictHdlr                        string    `bun:"restrict_hdlr,type:varchar(32)"`                                 // Restrict Handler
	PopupWidth                          int16     `bun:"popup_width,type:smallint"`                                      // Default popup window width
	PopupHeight                         int16     `bun:"popup_height,type:smallint"`                                     // Defaukt popup window height
	PageSize                            int16     `bun:"page_size,type:smallint"`                                        // Popup page size
	Literal                             int16     `bun:"literal,type:smallint"`                                          // Literal
	MinimumChars                        int16     `bun:"minimum_chars,type:smallint"`                                    // Minimum characters typed to show opoup
	DateCreated                         time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified                    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy                    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
	OnStartLoad                         bool      `bun:"on_start_load,type:bit,nullzero"`                                // Column to determine if the popup should load when its opened
	AutoFiltering                       bool      `bun:"auto_filtering,type:bit,nullzero"`                               // Column to determine if the partial key should filter as the user types in
	FocusOnGrid                         bool      `bun:"focus_on_grid,type:bit,nullzero"`                                // Column to determine if the popup grid should be focused when the popup is opened
	ColumnsResizable                    bool      `bun:"columns_resizable,type:bit,nullzero"`                            // Column to determine if the popup grid columns can be resized
	ExportEnabled                       bool      `bun:"export_enabled,type:bit,nullzero"`                               // If the popup's grid can be exported
	GroupingEnabled                     bool      `bun:"grouping_enabled,type:bit,nullzero"`                             // If the popup's grid can be grouped
	ColumnFilteringEnabled              bool      `bun:"column_filtering_enabled,type:bit,nullzero"`                     // If the columns filters are enabled
	CountRetrievalEnabled               bool      `bun:"count_retrieval_enabled,type:bit,nullzero"`                      // If the count of rows should be retrieved
	GenericSearchEnabled                bool      `bun:"generic_search_enabled,type:bit,nullzero"`                       // If the generic search feature is enabled
	MultipleRowSelectionEnabled         bool      `bun:"multiple_row_selection_enabled,type:bit,nullzero"`               // If multiple rows can be selected.
	UserConfigurationEnabled            bool      `bun:"user_configuration_enabled,type:bit,nullzero"`                   // If the popup's user configuration is enabled
	CoreFlag                            bool      `bun:"core_flag,type:bit,nullzero"`                                    // If the popup is a core popup
	DeleteFlag                          bool      `bun:"delete_flag,type:bit,nullzero"`                                  // If the popup is deleted
	InactiveFlag                        bool      `bun:"inactive_flag,type:bit,nullzero"`                                // If the popup is inactive
	AdditionalWhereEnabled              bool      `bun:"additional_where_enabled,type:bit,nullzero"`                     // Determines if the popup should apply additional conditions to its where
	PopupAttribute                      int32     `bun:"popup_attribute,type:int,nullzero"`                              // The attibute that defines the popup
	RankingEnabled                      bool      `bun:"ranking_enabled,type:bit,nullzero"`                              // Determines if the ranking feature is enabled
	Source                              string    `bun:"source,type:varchar(500),nullzero"`                              // Place where the definition of the popup has been extracted
	ExpandablePopup                     string    `bun:"expandable_popup,type:varchar(50),nullzero"`                     // Sets the column name of the popup to use in an expandable popup.
	CompSecEnabled                      bool      `bun:"comp_sec_enabled,type:bit,nullzero"`                             // Determines if the popup should apply company security its query
	OverridePopupDesc                   bool      `bun:"override_popup_desc,type:bit,default:((0))"`                     // flag to override description for Dynachange
	OverridePopupWidth                  bool      `bun:"override_popup_width,type:bit,default:((0))"`                    // flag to override width for Dynachange
	OverridePopupHeight                 bool      `bun:"override_popup_height,type:bit,default:((0))"`                   // flag to override height for Dynachange
	OverridePageSize                    bool      `bun:"override_page_size,type:bit,default:((0))"`                      // flag to override page size for Dynachange
	OverrideMinimumChars                bool      `bun:"override_minimum_chars,type:bit,default:((0))"`                  // flag to override minimum chars for Dynachange
	OverrideOnStartLoad                 bool      `bun:"override_on_start_load,type:bit,default:((0))"`                  // flag to override on start load for Dynachange
	OverrideAutoFiltering               bool      `bun:"override_auto_filtering,type:bit,default:((0))"`                 // flag to override auto filtering for Dynachange
	OverrideFocusOnGrid                 bool      `bun:"override_focus_on_grid,type:bit,default:((0))"`                  // flag to override description for Dynachange
	OverrideColumnsResizable            bool      `bun:"override_columns_resizable,type:bit,default:((0))"`              // flag to override columns resizable for Dynachange
	OverrideExportEnabled               bool      `bun:"override_export_enabled,type:bit,default:((0))"`                 // flag to override export for Dynachange
	OverrideGroupingEnabled             bool      `bun:"override_grouping_enabled,type:bit,default:((0))"`               // flag to override grouping for Dynachange
	OverrideColumnFilteringEnabled      bool      `bun:"override_column_filtering_enabled,type:bit,default:((0))"`       // flag to override column filtering for Dynachange
	OverrideCountRetrievalEnabled       bool      `bun:"override_count_retrieval_enabled,type:bit,default:((0))"`        // flag to override count retrieval for Dynachange
	OverrideGenericSearchEnabled        bool      `bun:"override_generic_search_enabled,type:bit,default:((0))"`         // flag to override generic search for Dynachange
	OverrideMultipleRowSelectionEnabled bool      `bun:"override_multiple_row_selection_enabled,type:bit,default:((0))"` // flag to override multiple row selection for Dynachange
	OverrideUserConfigurationEnabled    bool      `bun:"override_user_configuration_enabled,type:bit,default:((0))"`     // flag to override user configuration for Dynachange
	OverrideInactiveFlag                bool      `bun:"override_inactive_flag,type:bit,default:((0))"`                  // flag to override inactive for Dynachange
	OverrideAdditionalWhereEnabled      bool      `bun:"override_additional_where_enabled,type:bit,default:((0))"`       // flag to override additional where for Dynachange
	OverrideRankingEnabled              bool      `bun:"override_ranking_enabled,type:bit,default:((0))"`                // flag to override ranking for Dynachange
	OverrideSource                      bool      `bun:"override_source,type:bit,default:((0))"`                         // flag to override sourcefor Dynachange
	OverrideExpandablePopup             bool      `bun:"override_expandable_popup,type:bit,default:((0))"`               // flag to override expandable popup for Dynachange
	OverrideCompSecEnabled              bool      `bun:"override_comp_sec_enabled,type:bit,default:((0))"`               // flag to override Company Securrity for Dynachange
	PopupDetailUidParent                int32     `bun:"popup_detail_uid_parent,type:int,nullzero"`                      // Uid of the parent detail
	OverrideTitle                       bool      `bun:"override_title,type:bit,default:((0))"`                          // Check if the title of the child popup should take precedence
}
