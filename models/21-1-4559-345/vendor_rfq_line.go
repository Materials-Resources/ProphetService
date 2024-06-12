package model

type VendorRfqLine struct {
	bun.BaseModel           `bun:"table:vendor_rfq_line"`
	VendorRfqLineUid        int32     `bun:"vendor_rfq_line_uid,type:int,pk"`
	PoLineUid               int32     `bun:"po_line_uid,type:int"`
	QtyConverted            float64   `bun:"qty_converted,type:decimal(19,9),default:(0)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DeliveryInfo            string    `bun:"delivery_info,type:varchar(255),nullzero"`
	Responded               int32     `bun:"responded,type:int,default:(1104)"`
	RfqanalysisCompleteFlag string    `bun:"rfqanalysis_complete_flag,type:char,nullzero"`
}
