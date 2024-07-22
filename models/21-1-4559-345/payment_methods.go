package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentMethods struct {
	bun.BaseModel     `bun:"table:payment_methods"`
	PaymentMethodId   string    `bun:"payment_method_id,type:varchar(2),pk"`                      // Payment Method Identifier
	PaymentMethodDesc string    `bun:"payment_method_desc,type:varchar(30)"`                      // A description of the method
	DeleteFlag        string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
