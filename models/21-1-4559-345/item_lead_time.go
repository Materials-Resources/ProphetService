package model

type ItemLeadTime struct {
	bun.BaseModel       `bun:"table:item_lead_time"`
	ItemLeadTimeKey     int32     `bun:"item_lead_time_key,type:int,pk"`
	LocationId          float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	SupplierId          float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	ReceiptNumber       float64   `bun:"receipt_number,type:decimal(19,0),nullzero"`
	ReceiptDate         time.Time `bun:"receipt_date,type:datetime"`
	PoNo                float64   `bun:"po_no,type:decimal(19,0),nullzero"`
	PoDate              time.Time `bun:"po_date,type:datetime"`
	ExcludeFromLeadTime string    `bun:"exclude_from_lead_time,type:char"`
	ItemOrSupplier      string    `bun:"item_or_supplier,type:char"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	InvMastUid          int32     `bun:"inv_mast_uid,type:int,nullzero"`
	ActualLeadDays      int16     `bun:"actual_lead_days,type:smallint,default:(0)"`
	EstimatedLeadDays   int16     `bun:"estimated_lead_days,type:smallint,default:(0)"`
	Edited              string    `bun:"edited,type:char,default:('N')"`
}
