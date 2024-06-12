package model

type VoucherClass struct {
	bun.BaseModel           `bun:"table:voucher_class"`
	VoucherClass            string    `bun:"voucher_class,type:varchar(8),pk"`
	VoucherClassDesc        string    `bun:"voucher_class_desc,type:varchar(30)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	FreightVoucherClassFlag string    `bun:"freight_voucher_class_flag,type:char,nullzero"`
}
