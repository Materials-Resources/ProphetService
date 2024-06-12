package model

type UsersXCashDrawer struct {
	bun.BaseModel    `bun:"table:users_x_cash_drawer"`
	UserId           string    `bun:"user_id,type:varchar(30),pk"`
	CashDrawerId     string    `bun:"cash_drawer_id,type:varchar(8),pk"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
