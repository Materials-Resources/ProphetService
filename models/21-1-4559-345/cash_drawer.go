package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CashDrawer struct {
	bun.BaseModel           `bun:"table:cash_drawer"`
	CashDrawerId            string    `bun:"cash_drawer_id,type:varchar(8),pk"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`
	CashDrawerDescription   string    `bun:"cash_drawer_description,type:varchar(30)"`
	CurrentSequenceNo       float64   `bun:"current_sequence_no,type:decimal(19,0)"`
	OpeningBalance          float64   `bun:"opening_balance,type:decimal(19,4),nullzero"`
	Withdrawals             float64   `bun:"withdrawals,type:decimal(19,4),nullzero"`
	Deposits                float64   `bun:"deposits,type:decimal(19,4),nullzero"`
	CurrentBalance          float64   `bun:"current_balance,type:decimal(19,4)"`
	DrawerOpen              string    `bun:"drawer_open,type:char"`
	BankNo                  float64   `bun:"bank_no,type:decimal(19,0),nullzero"`
	CashOnHandAccountNumber string    `bun:"cash_on_hand_account_number,type:varchar(32)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CashCardLoad            float64   `bun:"cash_card_load,type:decimal(19,4),default:(0)"`
	CashDrawerUid           int32     `bun:"cash_drawer_uid,type:int,identity"`
	LocIdForBranchConflict  float64   `bun:"loc_id_for_branch_conflict,type:decimal(19,0),nullzero"`
}
