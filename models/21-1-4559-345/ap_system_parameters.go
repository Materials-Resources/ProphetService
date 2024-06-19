package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApSystemParameters struct {
	bun.BaseModel              `bun:"table:ap_system_parameters"`
	VoucherEntryByItemApproval string    `bun:"voucher_entry_by_item_approval,type:char(1)"`
	VoucherEntryByAmtApproval  string    `bun:"voucher_entry_by_amt_approval,type:char(1)"`
	PrepayVouchersApproval     string    `bun:"prepay_vouchers_approval,type:char(1)"`
	CrDrMemoEntryApproval      string    `bun:"cr_dr_memo_entry_approval,type:char(1)"`
	PostRecurringVchrApproval  string    `bun:"post_recurring_vchr_approval,type:char(1)"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ValidateJobId              string    `bun:"validate_job_id,type:char(1),nullzero"`
	DisplayJobId               string    `bun:"display_job_id,type:char(1),nullzero"`
	ApSystemParameterUid       int32     `bun:"ap_system_parameter_uid,type:int,pk"`
}
