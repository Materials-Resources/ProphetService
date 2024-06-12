package model

type CellRange struct {
	bun.BaseModel    `bun:"table:cell_range"`
	FinReportId      string    `bun:"fin_report_id,type:varchar(15),pk"`
	Cell             string    `bun:"cell,type:varchar(10),pk"`
	FromAccountNo    string    `bun:"from_account_no,type:varchar(32),pk"`
	ThruAccountNo    string    `bun:"thru_account_no,type:varchar(32),pk"`
	AddSubtract      string    `bun:"add_subtract,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
