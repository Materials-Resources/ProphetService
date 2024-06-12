package model

type ProdOrdLinePo struct {
	bun.BaseModel    `bun:"table:prod_ord_line_po"`
	ProdOrderNumber  float64   `bun:"prod_order_number,type:decimal(19,0),pk"`
	LineNumber       float64   `bun:"line_number,type:decimal(19,0),pk"`
	ComponentNumber  float64   `bun:"component_number,type:decimal(19,0),pk"`
	PoNumber         float64   `bun:"po_number,type:decimal(19,0),pk"`
	PoLineNumber     float64   `bun:"po_line_number,type:decimal(19,0),pk"`
	Quantity         float64   `bun:"quantity,type:decimal(19,9)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint,default:(1)"`
	ConnectionType   string    `bun:"connection_type,type:char,pk,default:('P')"`
}
