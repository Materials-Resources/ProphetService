package model

type PricePageTypeXCompany struct {
	bun.BaseModel            `bun:"table:price_page_type_x_company"`
	PricePageTypeXCompanyUid int32     `bun:"price_page_type_x_company_uid,type:int,pk"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	PricePageTypeCd          int16     `bun:"price_page_type_cd,type:smallint"`
	SequenceNumber           int16     `bun:"sequence_number,type:tinyint"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
}
