package model

type RmaReceiptLine struct {
	bun.BaseModel      `bun:"table:rma_receipt_line"`
	RmaReceiptLineUid  int32     `bun:"rma_receipt_line_uid,type:int,pk"`
	RmaReceiptHdrUid   int32     `bun:"rma_receipt_hdr_uid,type:int"`
	OeLineUid          int32     `bun:"oe_line_uid,type:int"`
	ReceiptLineNo      int32     `bun:"receipt_line_no,type:int"`
	QtyReceived        float64   `bun:"qty_received,type:decimal(19,9),default:(0)"`
	QtyToStock         float64   `bun:"qty_to_stock,type:decimal(19,9),default:(0)"`
	QtyToAdjust        float64   `bun:"qty_to_adjust,type:decimal(19,9),default:(0)"`
	QtyToSupplier      float64   `bun:"qty_to_supplier,type:decimal(19,9),default:(0)"`
	FreightIn          float64   `bun:"freight_in,type:decimal(19,9),default:(0)"`
	ReasonAdjustmentId string    `bun:"reason_adjustment_id,type:varchar(5),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:(704)"`
	ParentOeLineUid    int32     `bun:"parent_oe_line_uid,type:int,default:(null)"`
	UnitOfMeasure      string    `bun:"unit_of_measure,type:varchar(8),nullzero"`
	UnitSize           float64   `bun:"unit_size,type:decimal(19,9),nullzero"`
	QtyToTransfer      float64   `bun:"qty_to_transfer,type:decimal(19,9),nullzero"`
	CoreItemCost       float64   `bun:"core_item_cost,type:decimal(19,9),nullzero"`
	CountryOfOrigin    string    `bun:"country_of_origin,type:varchar(255),nullzero"`
}
