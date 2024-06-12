package model

type ApinvLine struct {
	bun.BaseModel         `bun:"table:apinv_line"`
	VoucherNo             string    `bun:"voucher_no,type:varchar(10)"`
	Quantity              float64   `bun:"quantity,type:decimal(19,9)"`
	ItemId                string    `bun:"item_id,type:varchar(40),nullzero"`
	UnitPrice             float64   `bun:"unit_price,type:decimal(19,9)"`
	PurchaseAmount        float64   `bun:"purchase_amount,type:decimal(19,4)"`
	PurchaseAccount       string    `bun:"purchase_account,type:varchar(32)"`
	Description           string    `bun:"description,type:varchar(30),nullzero"`
	TypeId                string    `bun:"type_id,type:varchar(8),nullzero"`
	Ten99Type             float64   `bun:"ten99_type,type:decimal(19,0)"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	CompanyNo             string    `bun:"company_no,type:varchar(8)"`
	HomeCurrencyAmt       float64   `bun:"home_currency_amt,type:decimal(19,4),nullzero"`
	ProjectId             string    `bun:"project_id,type:varchar(32),nullzero"`
	JobId                 string    `bun:"job_id,type:varchar(32),nullzero"`
	BranchId              string    `bun:"branch_id,type:varchar(8),nullzero"`
	UnitPriceDisplay      float64   `bun:"unit_price_display,type:decimal(19,9)"`
	PurchaseAmountDisplay float64   `bun:"purchase_amount_display,type:decimal(19,4)"`
	ApinvLineUid          int32     `bun:"apinv_line_uid,type:int,pk,default:(0)"`
	InvMastUid            int32     `bun:"inv_mast_uid,type:int,nullzero"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DisputedFlag          string    `bun:"disputed_flag,type:char,default:('N')"`
	DisputedAmt           float64   `bun:"disputed_amt,type:decimal(19,4),default:(0)"`
	OriginalDisputedAmt   float64   `bun:"original_disputed_amt,type:decimal(19,4),nullzero"`
	ApinvLineTypeCd       int32     `bun:"apinv_line_type_cd,type:int,nullzero"`
	ParentApinvLineUid    int32     `bun:"parent_apinv_line_uid,type:int,nullzero"`
	PoLineUid             int32     `bun:"po_line_uid,type:int,nullzero"`
	SourceNo              float64   `bun:"source_no,type:decimal(19,9),nullzero"`
	SequenceNumber        int32     `bun:"sequence_number,type:int,nullzero"`
	PayableSourceUid      int32     `bun:"payable_source_uid,type:int,nullzero"`
	FreightAmount         float64   `bun:"freight_amount,type:decimal(19,4),nullzero"`
	FreightAmountDisplay  float64   `bun:"freight_amount_display,type:decimal(19,4),nullzero"`
	GlDimenTypeUid        int32     `bun:"gl_dimen_type_uid,type:int,nullzero"`
	GlDimensionId         string    `bun:"gl_dimension_id,type:varchar(255),nullzero"`
	GlDimensionDesc       string    `bun:"gl_dimension_desc,type:varchar(255),nullzero"`
	PercentAllocated      float64   `bun:"percent_allocated,type:decimal(19,4),nullzero"`
}
