package model

type StageXProcess struct {
	bun.BaseModel         `bun:"table:stage_x_process"`
	StageXProcessUid      int32     `bun:"stage_x_process_uid,type:int,pk"`
	ProcessUid            int32     `bun:"process_uid,type:int"`
	StageUid              int32     `bun:"stage_uid,type:int"`
	SequenceNo            int16     `bun:"sequence_no,type:smallint"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`
	StagePoDescription    string    `bun:"stage_po_description,type:varchar(255),nullzero"`
	StageCode             string    `bun:"stage_code,type:varchar(255)"`
	StageDescription      string    `bun:"stage_description,type:varchar(8000)"`
	StageWipAccountNumber string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"`
	EstimatedHours        float64   `bun:"estimated_hours,type:decimal(6,2)"`
	PoRequired            string    `bun:"po_required,type:char,default:('N')"`
	VendorId              float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`
	SupplierId            float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	DivisionId            float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	BuyerId               string    `bun:"buyer_id,type:varchar(16),nullzero"`
	Cost                  float64   `bun:"cost,type:decimal(19,9)"`
	CostType              int16     `bun:"cost_type,type:smallint,nullzero"`
	CostUom               string    `bun:"cost_uom,type:varchar(8),nullzero"`
	CostUnitSize          float64   `bun:"cost_unit_size,type:decimal(19,4),nullzero"`
	StageQtyUom           string    `bun:"stage_qty_uom,type:varchar(32),nullzero"`
	QtyUnitSize           float64   `bun:"qty_unit_size,type:decimal(19,4)"`
	AllowPartials         string    `bun:"allow_partials,type:char,default:('N')"`
	MinimumPoCost         float64   `bun:"minimum_po_cost,type:decimal(19,9),nullzero"`
	WipInvMastUid         int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`
	BackflushFlag         string    `bun:"backflush_flag,type:char,nullzero"`
	TravelerRequiredFlag  string    `bun:"traveler_required_flag,type:char,default:('N')"`
	ConsolidatableFlag    string    `bun:"consolidatable_flag,type:char,nullzero"`
	ContainerInvMastUid   int32     `bun:"container_inv_mast_uid,type:int,nullzero"`
	ContainerUom          string    `bun:"container_uom,type:varchar(8),nullzero"`
	ContainerFlag         string    `bun:"container_flag,type:char,nullzero"`
	CommentProcess        string    `bun:"comment_process,type:char,default:('N')"`
	Comment               string    `bun:"comment,type:varchar,nullzero"`
}
