package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type GlSystemParameters struct {
	bun.BaseModel            `bun:"table:gl_system_parameters"`
	JournalEntryApproval     string    `bun:"journal_entry_approval,type:char(1)"`
	RepetitiveJeApproval     string    `bun:"repetitive_je_approval,type:char(1)"`
	ImportJeApproval         string    `bun:"import_je_approval,type:char(1)"`
	CoaMaskCompliance        string    `bun:"coa_mask_compliance,type:char(1)"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ValidateJobId            *string   `bun:"validate_job_id,type:char(1)"`
	DisplayJobId             *string   `bun:"display_job_id,type:char(1)"`
	TrackEncumbrances        string    `bun:"track_encumbrances,type:char(1)"`
	EncumbrancesBudgetColumn *float64  `bun:"encumbrances_budget_column,type:decimal(1,0)"`
	GlSystemParameterUid     string    `bun:"gl_system_parameter_uid,type:numeric(19,0),pk"`
}
