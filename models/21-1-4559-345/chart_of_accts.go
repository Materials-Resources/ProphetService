package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ChartOfAccts struct {
	bun.BaseModel              `bun:"table:chart_of_accts"`
	AccountNo                  string    `bun:"account_no,type:varchar(32),pk"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8),pk"`
	AccountDesc                string    `bun:"account_desc,type:varchar(40)"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AccountType                string    `bun:"account_type,type:char"`
	BranchId                   string    `bun:"branch_id,type:varchar(8)"`
	MaintainEncumbrances       string    `bun:"maintain_encumbrances,type:char,nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	AccountInformation         string    `bun:"account_information,type:varchar(255),nullzero"`
	GlReportDefaultPrintMethod int32     `bun:"gl_report_default_print_method,type:int,default:((1016))"`
	ChartOfAcctsUid            int32     `bun:"chart_of_accts_uid,type:int,identity"`
	RecordTypeCd               int32     `bun:"record_type_cd,type:int,nullzero"`
	DocLinkSmartFormFlag       string    `bun:"doc_link_smart_form_flag,type:char,nullzero"`
	AccountDefinitionCd        int32     `bun:"account_definition_cd,type:int,nullzero"`
}
