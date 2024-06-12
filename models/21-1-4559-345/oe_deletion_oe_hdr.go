package model

type OeDeletionOeHdr struct {
	bun.BaseModel `bun:"table:oe_deletion_oe_hdr"`
	OrderNo       string `bun:"order_no,type:varchar(8)"`
	DeleteFlag    string `bun:"delete_flag,type:char"`
	Completed     string `bun:"completed,type:char"`
	Type          string `bun:"type,type:char"`
}
