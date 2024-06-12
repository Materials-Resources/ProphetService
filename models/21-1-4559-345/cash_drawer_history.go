package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CashDrawerHistory struct {
	bun.BaseModel           `bun:"table:cash_drawer_history"`
	CashDrawerId            string    `bun:"cash_drawer_id,type:varchar(8),pk"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`
	SequenceNumber          float64   `bun:"sequence_number,type:decimal(19,0),pk"`
	UserId                  string    `bun:"user_id,type:varchar(30),default:(user_name())"`
	StartingBalance         float64   `bun:"starting_balance,type:decimal(19,4)"`
	ClosingBalance          float64   `bun:"closing_balance,type:decimal(19,4),nullzero"`
	WithdrawTotal           float64   `bun:"withdraw_total,type:decimal(19,4),nullzero"`
	DepositTotal            float64   `bun:"deposit_total,type:decimal(19,4),nullzero"`
	DepositNumber           string    `bun:"deposit_number,type:varchar(40),nullzero"`
	BankNo                  float64   `bun:"bank_no,type:decimal(19,0)"`
	CashOnHandAccountNumber string    `bun:"cash_on_hand_account_number,type:varchar(32)"`
	DateOpened              time.Time `bun:"date_opened,type:datetime"`
	DateClosed              time.Time `bun:"date_closed,type:datetime,nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	TotalCashCardLoads      float64   `bun:"total_cash_card_loads,type:decimal(19,4),nullzero"`
}
