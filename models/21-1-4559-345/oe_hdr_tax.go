package model

type OeHdrTax struct {
	bun.BaseModel    `bun:"table:oe_hdr_tax"`
	OrderNo          string    `bun:"order_no,type:varchar(8),pk"`
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	Taxable          string    `bun:"taxable,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	RmaLinkedTax     string    `bun:"rma_linked_tax,type:char,nullzero"`
}
