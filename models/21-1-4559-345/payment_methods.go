package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentMethods struct {
	bun.BaseModel     `bun:"table:payment_methods"`
	PaymentMethodId   string    `bun:"payment_method_id,type:varchar(2),pk"`
	PaymentMethodDesc string    `bun:"payment_method_desc,type:varchar(30)"`
	DeleteFlag        string    `bun:"delete_flag,type:char"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
