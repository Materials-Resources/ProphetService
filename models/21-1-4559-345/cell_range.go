package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CellRange struct {
	bun.BaseModel    `bun:"table:cell_range"`
	FinReportId      string    `bun:"fin_report_id,type:varchar(15),pk"`                         // What is the Financial Report ID?
	Cell             string    `bun:"cell,type:varchar(10),pk"`                                  // Which cell on the financial report does this range apply to?
	FromAccountNo    string    `bun:"from_account_no,type:varchar(32),pk"`                       // The first account in a range of accounts that will be added together.
	ThruAccountNo    string    `bun:"thru_account_no,type:varchar(32),pk"`                       // The last account in a range of accounts that will be added together.
	AddSubtract      string    `bun:"add_subtract,type:char(1)"`                                 // Indicates whether to add or subtract account amounts to calculate the value of a particular cell.
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
