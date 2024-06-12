package model

type ProdOrdLineCompntXOeLine struct {
	bun.BaseModel             `bun:"table:prod_ord_line_compnt_x_oe_line"`
	ProdOrderLineComponentUid int32     `bun:"prod_order_line_component_uid,type:int,pk"`
	OeLineUid                 int32     `bun:"oe_line_uid,type:int,pk"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
