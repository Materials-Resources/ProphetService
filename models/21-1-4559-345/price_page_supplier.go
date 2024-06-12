package model

type PricePageSupplier struct {
	bun.BaseModel        `bun:"table:price_page_supplier"`
	PricePageSupplierUid int32     `bun:"price_page_supplier_uid,type:int,pk,identity"`
	PricePageUid         int32     `bun:"price_page_uid,type:int"`
	SupplierId           float64   `bun:"supplier_id,type:decimal(19,0)"`
	EffectiveDate        time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate       time.Time `bun:"expiration_date,type:datetime"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`
}
