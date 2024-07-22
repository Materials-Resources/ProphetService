package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CashDrawerHistory struct {
	bun.BaseModel           `bun:"table:cash_drawer_history"`
	CashDrawerId            string     `bun:"cash_drawer_id,type:varchar(8),pk"`                         // Drawer identifier within a company.
	CompanyId               string     `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	SequenceNumber          float64    `bun:"sequence_number,type:decimal(19,0),pk"`                     // Indicates the sequence in which to process the loc
	UserId                  string     `bun:"user_id,type:varchar(30),default:(user_name())"`            // This column is unused.
	StartingBalance         float64    `bun:"starting_balance,type:decimal(19,4)"`                       // Balance when drawer opened
	ClosingBalance          *float64   `bun:"closing_balance,type:decimal(19,4)"`                        // Balance when drawer closed
	WithdrawTotal           *float64   `bun:"withdraw_total,type:decimal(19,4)"`                         // Total amount of money withdrawn from the drawer
	DepositTotal            *float64   `bun:"deposit_total,type:decimal(19,4)"`                          // Total amount of money deposited to the drawer
	DepositNumber           *string    `bun:"deposit_number,type:varchar(40)"`                           // Deposit number for payment -  applied when cash draw
	BankNo                  float64    `bun:"bank_no,type:decimal(19,0)"`                                // Default bank number
	CashOnHandAccountNumber string     `bun:"cash_on_hand_account_number,type:varchar(32)"`              // Cash on hand account money is held in until it is
	DateOpened              time.Time  `bun:"date_opened,type:datetime"`                                 // Date drawer is opened
	DateClosed              *time.Time `bun:"date_closed,type:datetime"`                                 // Date drawer is closed
	DateCreated             time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	TotalCashCardLoads      *float64   `bun:"total_cash_card_loads,type:decimal(19,4)"`                  // Amount of Cash Card transactions for a cash drawer once the drawer is closed.
}
