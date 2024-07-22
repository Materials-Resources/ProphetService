package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionRuleDetail struct {
	bun.BaseModel    `bun:"table:commission_rule_detail"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	CommissionRuleId string    `bun:"commission_rule_id,type:varchar(8),pk"`                     // Commission Rule ID
	BreakNumber      float64   `bun:"break_number,type:decimal(4,0),pk"`                         // Break Number (used for sorting the breaks)
	BreakValue       float64   `bun:"break_value,type:decimal(19,9)"`                            // Break Value - amount at which the break changes
	BreakType        string    `bun:"break_type,type:char(1)"`                                   // Break Type - computed or an exact amount
	Value            float64   `bun:"value,type:decimal(19,9)"`                                  // Value (for computation if Computed -  or exact if Ex
	Factor           *string   `bun:"factor,type:char(1)"`                                       // Whether use multiplier or percentage to calculate commission
	Source           *string   `bun:"source,type:char(1)"`                                       // What is the source for this repeating journal entry?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
