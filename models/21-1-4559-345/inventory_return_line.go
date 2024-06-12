package model

type InventoryReturnLine struct {
	bun.BaseModel          `bun:"table:inventory_return_line"`
	InventoryReturnLineUid int32     `bun:"inventory_return_line_uid,type:int,pk"`
	InventoryReturnHdrUid  int32     `bun:"inventory_return_hdr_uid,type:int"`
	LineNumber             int32     `bun:"line_number,type:int"`
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`
	ExtendedDescription    string    `bun:"extended_description,type:varchar(255),nullzero"`
	QuantityUom            string    `bun:"quantity_uom,type:varchar(8)"`
	QuantityUnitSize       float64   `bun:"quantity_unit_size,type:decimal(19,4)"`
	UnitQuantity           float64   `bun:"unit_quantity,type:decimal(19,9),default:(0)"`
	QtyToReturn            float64   `bun:"qty_to_return,type:decimal(19,9),default:(0)"`
	QtyPicked              float64   `bun:"qty_picked,type:decimal(19,9),default:(0)"`
	QtyReturned            float64   `bun:"qty_returned,type:decimal(19,9),default:(0)"`
	QtyVouched             float64   `bun:"qty_vouched,type:decimal(19,9),default:(0)"`
	UnitPrice              float64   `bun:"unit_price,type:decimal(19,9),default:(0)"`
	UnitPriceDisplay       float64   `bun:"unit_price_display,type:decimal(19,9),default:(0)"`
	PriceUom               string    `bun:"price_uom,type:varchar(8)"`
	PriceUnitSize          float64   `bun:"price_unit_size,type:decimal(19,4)"`
	ExtendedPrice          float64   `bun:"extended_price,type:decimal(19,2),default:(0)"`
	PriceEdited            string    `bun:"price_edited,type:char,default:('N')"`
	ReturnDate             time.Time `bun:"return_date,type:datetime,default:(getdate())"`
	ExpectedDate           time.Time `bun:"expected_date,type:datetime,default:(getdate())"`
	SupplierPartNumber     string    `bun:"supplier_part_number,type:varchar(255),nullzero"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:(702)"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	RetrievedByWms         string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	CoreItemCost           float64   `bun:"core_item_cost,type:decimal(19,9),nullzero"`
}
