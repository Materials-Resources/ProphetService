package model

type PricePageDiscgrp struct {
	bun.BaseModel       `bun:"table:price_page_discgrp"`
	PricePageDiscgrpUid int32     `bun:"price_page_discgrp_uid,type:int,pk,identity"`
	PricePageUid        int32     `bun:"price_page_uid,type:int"`
	DiscountGroupId     string    `bun:"discount_group_id,type:varchar(8)"`
	EffectiveDate       time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate      time.Time `bun:"expiration_date,type:datetime"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
}
