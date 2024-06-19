package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type VoucherClass struct {
	bun.BaseModel           `bun:"table:voucher_class"`
	VoucherClass            string    `bun:"voucher_class,type:varchar(8),pk"`                          // A user-defined code that can be used to identify a group of vouchers.
	VoucherClassDesc        string    `bun:"voucher_class_desc,type:varchar(30)"`                       // What is this voucher class for?
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	FreightVoucherClassFlag string    `bun:"freight_voucher_class_flag,type:char(1),nullzero"`          // Indicates whether a voucher class will be used to input freight vouchers, used by custom only.
}
