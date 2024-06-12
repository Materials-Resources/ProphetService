package model

type OeLinePo struct {
	bun.BaseModel           `bun:"table:oe_line_po"`
	OrderNumber             string    `bun:"order_number,type:varchar(8),pk"`
	LineNumber              float64   `bun:"line_number,type:decimal(19,0),pk"`
	PoNo                    float64   `bun:"po_no,type:decimal(19,0),pk"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	QuantityOnPo            float64   `bun:"quantity_on_po,type:decimal(19,9),nullzero"`
	PoLineNumber            float64   `bun:"po_line_number,type:decimal(19,0),pk"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	CancelFlag              string    `bun:"cancel_flag,type:char,nullzero"`
	Completed               string    `bun:"completed,type:char"`
	ConnectionType          string    `bun:"connection_type,type:char,pk"`
	OriginalOeLineFreightIn float64   `bun:"original_oe_line_freight_in,type:decimal(19,9),nullzero"`
}
