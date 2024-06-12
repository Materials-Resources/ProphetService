package model

type InvXref struct {
	bun.BaseModel             `bun:"table:inv_xref"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0),pk"`
	TheirItemId               string    `bun:"their_item_id,type:varchar(40),pk"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	InvMastUid                int32     `bun:"inv_mast_uid,type:int"`
	DefaultSourceLocationId   float64   `bun:"default_source_location_id,type:decimal(19,0),nullzero"`
	ConsignmentProcessing     int32     `bun:"consignment_processing,type:int,default:(2)"`
	InvXrefUid                int32     `bun:"inv_xref_uid,type:int"`
	CustomerAccountNumber     string    `bun:"customer_account_number,type:varchar(16),nullzero"`
	BaselinePrice             float64   `bun:"baseline_price,type:decimal(19,9),nullzero"`
	Requestor                 string    `bun:"requestor,type:varchar(25),nullzero"`
	Department                string    `bun:"department,type:varchar(16),nullzero"`
	Category                  string    `bun:"category,type:varchar(16),nullzero"`
	ContractNumber            string    `bun:"contract_number,type:varchar(16),nullzero"`
	Miscellaneous1            string    `bun:"miscellaneous1,type:varchar(16),nullzero"`
	Miscellaneous2            string    `bun:"miscellaneous2,type:varchar(16),nullzero"`
	Miscellaneous3            string    `bun:"miscellaneous3,type:varchar(16),nullzero"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TheirBinId                string    `bun:"their_bin_id,type:varchar(10),nullzero"`
	VmiFlag                   string    `bun:"vmi_flag,type:char,nullzero"`
	LastReceiptDate           time.Time `bun:"last_receipt_date,type:datetime,nullzero"`
	LastPoNo                  string    `bun:"last_po_no,type:varchar(255),nullzero"`
	OnHandQty                 float64   `bun:"on_hand_qty,type:decimal(19,9),nullzero"`
	MinQty                    float64   `bun:"min_qty,type:decimal(19,9),nullzero"`
	MaxQty                    float64   `bun:"max_qty,type:decimal(19,9),nullzero"`
	MinUpdateOn852ImportFlag  string    `bun:"min_update_on_852_import_flag,type:char,nullzero"`
	MaxUpdateOn852ImportFlag  string    `bun:"max_update_on_852_import_flag,type:char,nullzero"`
	DirectShipDispositionFlag string    `bun:"direct_ship_disposition_flag,type:char,nullzero"`
	LastPricePaid             float64   `bun:"last_price_paid,type:decimal(19,9),nullzero"`
	ExtendedDescription       string    `bun:"extended_description,type:varchar(255),nullzero"`
	ItemRevisionUid           int32     `bun:"item_revision_uid,type:int,nullzero"`
	AvgUnitSellPrice          float64   `bun:"avg_unit_sell_price,type:decimal(19,9),nullzero"`
	ItemLabelUom              string    `bun:"item_label_uom,type:varchar(8),nullzero"`
	ItemLabelUnitSize         float64   `bun:"item_label_unit_size,type:decimal(19,4),nullzero"`
	ItemLabelActionFlag       string    `bun:"item_label_action_flag,type:char,nullzero"`
	ItemLabelText             string    `bun:"item_label_text,type:varchar(255),nullzero"`
	ItemLabelCalcTypeFlag     string    `bun:"item_label_calc_type_flag,type:char,nullzero"`
	ItemLabelMultiplier       float64   `bun:"item_label_multiplier,type:decimal(19,9),nullzero"`
}
