package model

type InstallmentFrequency10005 struct {
	bun.BaseModel            `bun:"table:installment_frequency_10005"`
	InstallmentFrequencyUid  int32     `bun:"installment_frequency_uid,type:int,pk"`
	InstallmentFrequencyName string    `bun:"installment_frequency_name,type:varchar(15)"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
	InstallmentFrequencyCode string    `bun:"installment_frequency_code,type:char(3)"`
}
