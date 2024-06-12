package model

type DiscountInstallment10005 struct {
	bun.BaseModel           `bun:"table:discount_installment_10005"`
	DiscountInstallmentUid  int32     `bun:"discount_installment_uid,type:int,pk"`
	DiscountInstallmentName string    `bun:"discount_installment_name,type:varchar(20)"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DiscountInstallmentCode string    `bun:"discount_installment_code,type:char(3)"`
}
