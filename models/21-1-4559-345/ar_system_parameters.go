package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ArSystemParameters struct {
	bun.BaseModel               `bun:"table:ar_system_parameters"`
	InvoiceEntryApproval        string    `bun:"invoice_entry_approval,type:char(1)"`
	CrDrMemoEntryApproval       string    `bun:"cr_dr_memo_entry_approval,type:char(1)"`
	PrepaidInvoiceEntryApproval string    `bun:"prepaid_invoice_entry_approval,type:char(1)"`
	ImportInvoicesApproval      string    `bun:"import_invoices_approval,type:char(1)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`
	ValidateJobId               string    `bun:"validate_job_id,type:char(1),nullzero"`
	DisplayJobId                string    `bun:"display_job_id,type:char(1),nullzero"`
	ArIntercompanyFunctionality string    `bun:"ar_intercompany_functionality,type:char(1),nullzero"`
	NumInvoiceCopies            float64   `bun:"num_invoice_copies,type:decimal(2,0),nullzero"`
	FinanceChargeDesc           string    `bun:"finance_charge_desc,type:varchar(40),nullzero"`
	CommissionMultiplier        string    `bun:"commission_multiplier,type:varchar(2),nullzero"`
	CommissionBasedOn           string    `bun:"commission_based_on,type:char(2),nullzero"`
	PostCommissionToGl          string    `bun:"post_commission_to_gl,type:char(1),nullzero"`
	ReverseFinanceCharges       string    `bun:"reverse_finance_charges,type:char(1),nullzero"`
	CreditLimitAction           string    `bun:"credit_limit_action,type:char(1),nullzero"`
	ArSystemParameterUid        int32     `bun:"ar_system_parameter_uid,type:int,pk"`
}
