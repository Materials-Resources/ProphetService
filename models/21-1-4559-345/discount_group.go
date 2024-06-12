package model

type DiscountGroup struct {
	bun.BaseModel            `bun:"table:discount_group"`
	DiscountGroupId          string    `bun:"discount_group_id,type:varchar(8),pk"`
	DiscountGroupDescription string    `bun:"discount_group_description,type:varchar(40)"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ExtendedDesc             string    `bun:"extended_desc,type:varchar(255),nullzero"`
}
