package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRule struct {
	bun.BaseModel              `bun:"table:business_rule"`
	BusinessRuleUid            int32     `bun:"business_rule_uid,type:int,autoincrement,scanonly,pk"`         // Unique identifier
	RuleTypeCd                 int32     `bun:"rule_type_cd,type:int"`                                        // The type of the rule
	RuleName                   string    `bun:"rule_name,type:varchar(255)"`                                  // The name of the rule
	FieldName                  string    `bun:"field_name,type:varchar(255),nullzero"`                        // The field for which the rule applies
	ClassName                  string    `bun:"class_name,type:varchar(255),nullzero"`                        // The class for which the rule applies
	ApplyDuringSaveFlag        string    `bun:"apply_during_save_flag,type:char(1),default:('N')"`            // Does this rule get applied during the save?
	ApplyGloballyFlag          string    `bun:"apply_globally_flag,type:char(1),default:('N')"`               // Does this rule get applied throughout the entire system?
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag              int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Status of the Business Rule
	RunForAllFlag              string    `bun:"run_for_all_flag,type:char(1),default:('Y')"`                  // Run the business rule for all users?
	WindowName                 string    `bun:"window_name,type:varchar(255),nullzero"`                       // Name of the window containing the rule
	WindowTitle                string    `bun:"window_title,type:varchar(255),nullzero"`                      // Title of the window containing the rule
	LicenseKey                 string    `bun:"license_key,type:varchar(255),nullzero"`                       // Key for enabling this rule if DC Rules is not licensed
	ApplyToAllRowsFlag         string    `bun:"apply_to_all_rows_flag,type:char(1),default:('N')"`            // Apply business rules to all rows in datastore
	MultirowFlag               string    `bun:"multirow_flag,type:char(1),default:('N')"`                     // Indicates whether this business rule applies to multiple rows or a single row [Y - multirow/ N]
	BusinessRuleEventUid       int32     `bun:"business_rule_event_uid,type:int,nullzero"`                    // Determines the event from which this rule will be triggered.
	RunTypeCd                  int32     `bun:"run_type_cd,type:int,default:((0))"`                           // Applies to On Event rules.  Determines run type:  Synchronous or Asynchronous
	BusinessRuleEventKeyUid    int32     `bun:"business_rule_event_key_uid,type:int,nullzero"`                // UID to event/key record that determines the specific instance of an event that this rule applies to.
	InternalRuleFlag           string    `bun:"internal_rule_flag,type:char(1),default:('N')"`                // Identify whether this rule is used in P21 for internal purposes
	EnabledForVersionCd        int32     `bun:"enabled_for_version_cd,type:int,default:((1423))"`             // Indicates whether rule is used in desktop, web, or both.
	RulePageUrl                string    `bun:"rule_page_url,type:varchar(8000),nullzero"`                    // URL for the ASP.Net page to be launched by a web visual business rule.
	ShowRulePageUrlDesktopFlag string    `bun:"show_rule_page_url_desktop_flag,type:char(1),default:('N')"`   // Indicates whether the web visual rule URL should render in a browser control in desktop P21.
	OrderEntryFlag             string    `bun:"order_entry_flag,type:char(1),nullzero"`                       // Indicate if the rule for Order Entry window be executed or not
	FrontCounterFlag           string    `bun:"front_counter_flag,type:char(1),nullzero"`                     // Indicate if the rule for Front Counter window be executed or not
	RmaEntryFlag               string    `bun:"rma_entry_flag,type:char(1),nullzero"`                         // Indicate if the rule for RMA Entry window be executed or not
}
