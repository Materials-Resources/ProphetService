package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Branch struct {
	bun.BaseModel             `bun:"table:branch"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	BranchId                  string    `bun:"branch_id,type:varchar(8),pk"`                              // 8 character identifier for the branch
	BranchDescription         string    `bun:"branch_description,type:varchar(40)"`                       // What is the human-readable branch description?
	PayableGroupId            *string   `bun:"payable_group_id,type:varchar(255)"`                        // This column is no longer being used and is scheduled for removal in later version.
	ReceivableGroupId         *string   `bun:"receivable_group_id,type:varchar(255)"`                     // Identifier for receivable group.
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LogoPathFilename          *string   `bun:"logo_path_filename,type:varchar(255)"`                      // Allows the user to specify a logo bitmap file that
	RemitToAddressId          *float64  `bun:"remit_to_address_id,type:decimal(19,0)"`                    // Remit To Address information at the Branch level for printing on invoices
	LaborBillbackAccountNo    *string   `bun:"labor_billback_account_no,type:varchar(32)"`                // Default Labor Bill Back Account Number for locations.
	DefaultSalesrepId         *string   `bun:"default_salesrep_id,type:varchar(255)"`                     // Default Salesrep ID to be used when creating customers.
	DunsNumber                *string   `bun:"duns_number,type:varchar(255)"`                             // Dun and Bradstreet assigned identifier for a branch.
	PreventAutoAssignLotsFlag string    `bun:"prevent_auto_assign_lots_flag,type:char(1),default:('N')"`  // Indicates if automatic allocation of lots for allocated quantity should be overridden
	BranchUid                 int32     `bun:"branch_uid,type:int,autoincrement,unique,identity"`         // Unique identifier for the table
	NetProfitConfigurationUid *int32    `bun:"net_profit_configuration_uid,type:int"`                     // For CPA, the net profit calculation to be used for the branch
	CompanyUid                int32     `bun:"company_uid,type:int"`                                      // Unique identifier for company table. Used with CPA
	CfdiTimezoneOffset        *string   `bun:"cfdi_timezone_offset,type:varchar(255)"`                    // CFDI Time Zone Offset
	EmailSubjectPrefix        *string   `bun:"email_subject_prefix,type:varchar(255)"`                    // (Custom F80588) Prefix that will be added to the normal subject of form based email based on the branch of the transaction
}
