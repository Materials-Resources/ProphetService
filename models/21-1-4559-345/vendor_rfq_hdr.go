package model

type VendorRfqHdr struct {
	bun.BaseModel    `bun:"table:vendor_rfq_hdr"`
	VendorRfqHdrUid  int32     `bun:"vendor_rfq_hdr_uid,type:int,pk"`
	PoHdrUid         int32     `bun:"po_hdr_uid,type:int"`
	ControlNo        int32     `bun:"control_no,type:int"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	Date840Sent      time.Time `bun:"date_840_sent,type:datetime,nullzero"`
	TypeCd           int32     `bun:"type_cd,type:int,nullzero"`
	ResponseDate     time.Time `bun:"response_date,type:datetime,default:(null)"`
}
