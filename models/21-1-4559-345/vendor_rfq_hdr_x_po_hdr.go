package model

type VendorRfqHdrXPoHdr struct {
	bun.BaseModel         `bun:"table:vendor_rfq_hdr_x_po_hdr"`
	VendorRfqHdrXPoHdrUid int32     `bun:"vendor_rfq_hdr_x_po_hdr_uid,type:int,pk"`
	PoHdrUid              int32     `bun:"po_hdr_uid,type:int"`
	VendorRfqHdrUid       int32     `bun:"vendor_rfq_hdr_uid,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
