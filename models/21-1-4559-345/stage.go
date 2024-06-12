package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Stage struct {
	bun.BaseModel         `bun:"table:stage"`
	StageUid              int32     `bun:"stage_uid,type:int,pk"`
	StageCode             string    `bun:"stage_code,type:varchar(255)"`
	StageDescription      string    `bun:"stage_description,type:varchar(8000)"`
	StageWipAccountNumber string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"`
	EstimatedHours        float64   `bun:"estimated_hours,type:decimal(6,2)"`
	SupplierId            float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	VendorId              float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`
	CompanyId             string    `bun:"company_id,type:varchar(8),nullzero"`
	Cost                  float64   `bun:"cost,type:decimal(19,9)"`
	CostUom               string    `bun:"cost_uom,type:varchar(8),nullzero"`
	AllowPartials         string    `bun:"allow_partials,type:char"`
	PoRequired            string    `bun:"po_required,type:char"`
	MinimumPoCost         float64   `bun:"minimum_po_cost,type:decimal(19,9),nullzero"`
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`
	BuyerId               string    `bun:"buyer_id,type:varchar(16),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	StageQtyUom           string    `bun:"stage_qty_uom,type:varchar(32),nullzero"`
	DivisionId            float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	CostType              int16     `bun:"cost_type,type:smallint,nullzero"`
	QtyUnitSize           float64   `bun:"qty_unit_size,type:decimal(19,4)"`
	CostUnitSize          float64   `bun:"cost_unit_size,type:decimal(19,4),nullzero"`
	AllLocationsFlag      string    `bun:"all_locations_flag,type:char,default:('Y')"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	StagePoDesc           string    `bun:"stage_po_desc,type:varchar(255),nullzero"`
	WipInvMastUid         int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`
	BackflushFlag         string    `bun:"backflush_flag,type:char,nullzero"`
	TravelerRequiredFlag  string    `bun:"traveler_required_flag,type:char,default:('N')"`
	ConsolidatableFlag    string    `bun:"consolidatable_flag,type:char,default:('N')"`
	ContainerInvMastUid   int32     `bun:"container_inv_mast_uid,type:int,nullzero"`
	ContainerUom          string    `bun:"container_uom,type:varchar(8),nullzero"`
	ContainerFlag         string    `bun:"container_flag,type:char,nullzero"`
	CommentProcess        string    `bun:"comment_process,type:char,default:('N')"`
	Comment               string    `bun:"comment,type:varchar,nullzero"`
}
