package model

type OeHdrNotDelete struct {
	bun.BaseModel `bun:"table:oe_hdr_not_delete"`
	OrderNo       string `bun:"order_no,type:varchar(8)"`
	Reason        string `bun:"reason,type:varchar(255)"`
}
