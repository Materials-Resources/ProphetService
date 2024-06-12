package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Branch struct {
	bun.BaseModel             `bun:"table:branch"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`
	BranchId                  string    `bun:"branch_id,type:varchar(8),pk"`
	BranchDescription         string    `bun:"branch_description,type:varchar(40)"`
	PayableGroupId            string    `bun:"payable_group_id,type:varchar(255),nullzero"`
	ReceivableGroupId         string    `bun:"receivable_group_id,type:varchar(255),nullzero"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LogoPathFilename          string    `bun:"logo_path_filename,type:varchar(255),nullzero"`
	RemitToAddressId          float64   `bun:"remit_to_address_id,type:decimal(19,0),nullzero"`
	LaborBillbackAccountNo    string    `bun:"labor_billback_account_no,type:varchar(32),nullzero"`
	DefaultSalesrepId         string    `bun:"default_salesrep_id,type:varchar(255),nullzero"`
	DunsNumber                string    `bun:"duns_number,type:varchar(255),nullzero"`
	PreventAutoAssignLotsFlag string    `bun:"prevent_auto_assign_lots_flag,type:char,default:('N')"`
	BranchUid                 int32     `bun:"branch_uid,type:int,identity"`
	NetProfitConfigurationUid int32     `bun:"net_profit_configuration_uid,type:int,nullzero"`
	CompanyUid                int32     `bun:"company_uid,type:int"`
	CfdiTimezoneOffset        string    `bun:"cfdi_timezone_offset,type:varchar(255),nullzero"`
	EmailSubjectPrefix        string    `bun:"email_subject_prefix,type:varchar(255),nullzero"`
}
