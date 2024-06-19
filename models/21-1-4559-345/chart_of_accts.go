package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ChartOfAccts struct {
	bun.BaseModel              `bun:"table:chart_of_accts"`
	AccountNo                  string    `bun:"account_no,type:varchar(32),pk"`                            // What account is this record for?
	CompanyNo                  string    `bun:"company_no,type:varchar(8),pk"`                             // Unique code that identifies a company.
	AccountDesc                string    `bun:"account_desc,type:varchar(40)"`                             // What is the account description?
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	AccountType                string    `bun:"account_type,type:char(1)"`                                 // What is the account type?
	BranchId                   string    `bun:"branch_id,type:varchar(8)"`                                 // Which branch is this account for?  This is needed to implement branch accounting.
	MaintainEncumbrances       string    `bun:"maintain_encumbrances,type:char(1),nullzero"`               // Does this account maintain encumbrances?
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	AccountInformation         string    `bun:"account_information,type:varchar(255),nullzero"`            // a field for extended information about the account
	GlReportDefaultPrintMethod int32     `bun:"gl_report_default_print_method,type:int,default:((1016))"`  // default printing type for GL reports
	ChartOfAcctsUid            int32     `bun:"chart_of_accts_uid,type:int,autoincrement,unique,scanonly"` // Unique identity column for chart_of_accts
	RecordTypeCd               int32     `bun:"record_type_cd,type:int,nullzero"`                          // Determine the type of record (User Defined, System Defined, etc...)
	DocLinkSmartFormFlag       string    `bun:"doc_link_smart_form_flag,type:char(1),nullzero"`            // Flag to indicate that the associated gl account is valid to use with doc-link smart forms.
	AccountDefinitionCd        int32     `bun:"account_definition_cd,type:int,nullzero"`                   // Determines whether the Account is an AR/AP Account
}
