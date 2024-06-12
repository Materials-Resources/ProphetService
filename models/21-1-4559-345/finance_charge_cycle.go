package model

type FinanceChargeCycle struct {
	bun.BaseModel    `bun:"table:finance_charge_cycle"`
	FcCycleId        string    `bun:"fc_cycle_id,type:char,pk"`
	FcCycleDesc      string    `bun:"fc_cycle_desc,type:varchar(40)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
