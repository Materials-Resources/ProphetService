package model

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRule struct {
	bun.BaseModel              `bun:"table:business_rule"`
	BusinessRuleUid            int32     `bun:"business_rule_uid,type:int,pk,identity"`
	RuleTypeCd                 int32     `bun:"rule_type_cd,type:int"`
	RuleName                   string    `bun:"rule_name,type:varchar(255)"`
	FieldName                  string    `bun:"field_name,type:varchar(255),nullzero"`
	ClassName                  string    `bun:"class_name,type:varchar(255),nullzero"`
	ApplyDuringSaveFlag        string    `bun:"apply_during_save_flag,type:char,default:('N')"`
	ApplyGloballyFlag          string    `bun:"apply_globally_flag,type:char,default:('N')"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag              int32     `bun:"row_status_flag,type:int,default:((704))"`
	RunForAllFlag              string    `bun:"run_for_all_flag,type:char,default:('Y')"`
	WindowName                 string    `bun:"window_name,type:varchar(255),nullzero"`
	WindowTitle                string    `bun:"window_title,type:varchar(255),nullzero"`
	LicenseKey                 string    `bun:"license_key,type:varchar(255),nullzero"`
	ApplyToAllRowsFlag         string    `bun:"apply_to_all_rows_flag,type:char,default:('N')"`
	MultirowFlag               string    `bun:"multirow_flag,type:char,default:('N')"`
	BusinessRuleEventUid       int32     `bun:"business_rule_event_uid,type:int,nullzero"`
	RunTypeCd                  int32     `bun:"run_type_cd,type:int,default:((0))"`
	BusinessRuleEventKeyUid    int32     `bun:"business_rule_event_key_uid,type:int,nullzero"`
	InternalRuleFlag           string    `bun:"internal_rule_flag,type:char,default:('N')"`
	EnabledForVersionCd        int32     `bun:"enabled_for_version_cd,type:int,default:((1423))"`
	RulePageUrl                string    `bun:"rule_page_url,type:varchar(8000),nullzero"`
	ShowRulePageUrlDesktopFlag string    `bun:"show_rule_page_url_desktop_flag,type:char,default:('N')"`
	OrderEntryFlag             string    `bun:"order_entry_flag,type:char,nullzero"`
	FrontCounterFlag           string    `bun:"front_counter_flag,type:char,nullzero"`
	RmaEntryFlag               string    `bun:"rma_entry_flag,type:char,nullzero"`
}
