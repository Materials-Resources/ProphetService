package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CashDrawer struct {
	bun.BaseModel           `bun:"table:cash_drawer"`
	CashDrawerId            string    `bun:"cash_drawer_id,type:varchar(8),pk"`                         // Drawer identifier within a company.
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	CashDrawerDescription   string    `bun:"cash_drawer_description,type:varchar(30)"`                  // Description for the cash drawer
	CurrentSequenceNo       float64   `bun:"current_sequence_no,type:decimal(19,0)"`                    // Current active history record.
	OpeningBalance          *float64  `bun:"opening_balance,type:decimal(19,4)"`                        // Opening balance of drawer when it was opened
	Withdrawals             *float64  `bun:"withdrawals,type:decimal(19,4)"`                            // Sum of amount withdrawn from the cash drawer
	Deposits                *float64  `bun:"deposits,type:decimal(19,4)"`                               // Sum of deposits into the cash drawer
	CurrentBalance          float64   `bun:"current_balance,type:decimal(19,4)"`                        // Current balance -  includes remittances
	DrawerOpen              string    `bun:"drawer_open,type:char(1)"`                                  // Indicates whether the drawer is open or not.
	BankNo                  *float64  `bun:"bank_no,type:decimal(19,0)"`                                // Enter a valid bank number
	CashOnHandAccountNumber string    `bun:"cash_on_hand_account_number,type:varchar(32)"`              // Default cash on hand account number
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CashCardLoad            *float64  `bun:"cash_card_load,type:decimal(19,4),default:(0)"`             // Amount of Cash Cash transactions for the cash drawer
	CashDrawerUid           int32     `bun:"cash_drawer_uid,type:int,autoincrement,identity"`           // Unique identifier for the cash drawer
	LocIdForBranchConflict  *float64  `bun:"loc_id_for_branch_conflict,type:decimal(19,0)"`             // Location id to get the default branch if we are not able to determine a single branch from the users on cash drawer in non interactive mode.
}
