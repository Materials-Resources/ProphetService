package model

type PricePageItem struct {
	bun.BaseModel    `bun:"table:price_page_item"`
	PricePageItemUid int32     `bun:"price_page_item_uid,type:int,pk,identity"`
	PricePageUid     int32     `bun:"price_page_uid,type:int"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int"`
	EffectiveDate    time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
}
