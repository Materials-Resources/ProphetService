package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type UsersXCashDrawer struct {
	bun.BaseModel    `bun:"table:users_x_cash_drawer"`
	UserId           string    `bun:"user_id,type:varchar(30),pk"`                                  // User associated with the cash drawer.
	CashDrawerId     string    `bun:"cash_drawer_id,type:varchar(8),pk"`                            // Identifier for the cash drawer.
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                                // Identifier for the company associated with this cash drawer.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
