package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type StageXProcess struct {
	bun.BaseModel         `bun:"table:stage_x_process"`
	StageXProcessUid      int32     `bun:"stage_x_process_uid,type:int,pk"`                               // Unique identifier for each relationship between a stage and a process.
	ProcessUid            int32     `bun:"process_uid,type:int"`                                          // What process belongs to this relationship to a transaction?
	StageUid              int32     `bun:"stage_uid,type:int"`                                            // FK to stage table.
	SequenceNo            int16     `bun:"sequence_no,type:smallint"`                                     // What sequence should the process be performed in -  for this stage?
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`                                 // Indicates current record status.
	StagePoDescription    string    `bun:"stage_po_description,type:varchar(255),nullzero"`
	StageCode             string    `bun:"stage_code,type:varchar(255)"`                       // Code used to identify the stage.
	StageDescription      string    `bun:"stage_description,type:varchar(8000)"`               // A brief description of the stage.
	StageWipAccountNumber string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"` // Work in process account to hold cost of materials in process.
	EstimatedHours        float64   `bun:"estimated_hours,type:decimal(6,2)"`                  // Estimated hours it should take to complete the stage.
	PoRequired            string    `bun:"po_required,type:char(1),default:('N')"`             // Is a Purchase Order required when moving material into the stage?
	VendorId              float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`              // Vendor for stages that require a po.
	SupplierId            float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`            // Supplier for stages that require a po.
	DivisionId            float64   `bun:"division_id,type:decimal(19,0),nullzero"`            // Division for stages that require a po.
	BuyerId               string    `bun:"buyer_id,type:varchar(16),nullzero"`                 // Buyer for stages that require a po.
	Cost                  float64   `bun:"cost,type:decimal(19,9)"`                            // Cost of the work associated with the stage.
	CostType              int16     `bun:"cost_type,type:smallint,nullzero"`                   // Unit to determine cost for stage (ie Weight, UOM)
	CostUom               string    `bun:"cost_uom,type:varchar(8),nullzero"`                  // The unit of measure associated with the cost.
	CostUnitSize          float64   `bun:"cost_unit_size,type:decimal(19,4),nullzero"`         // Size of the cost unit of measure
	StageQtyUom           string    `bun:"stage_qty_uom,type:varchar(32),nullzero"`            // Describes the unit of measure the material gets processed in. (Ex. Feet - Pounds)
	QtyUnitSize           float64   `bun:"qty_unit_size,type:decimal(19,4)"`                   // Size of the quantity unit of measure
	AllowPartials         string    `bun:"allow_partials,type:char(1),default:('N')"`          // Indicates whether partial quantities can be moved to a new stage.
	MinimumPoCost         float64   `bun:"minimum_po_cost,type:decimal(19,9),nullzero"`        // Minimum po cost for stages that require a po.
	WipInvMastUid         int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`                 // Column holds inv_mast_uid of the item at the point that a process is performed on the item.
	BackflushFlag         string    `bun:"backflush_flag,type:char(1),nullzero"`               // Determines whether the process will allow Backflushing functionality.
	TravelerRequiredFlag  string    `bun:"traveler_required_flag,type:char(1),default:('N')"`  // Indicates if the system should update the master row_status_flag to print pending when material is moved in.
	ConsolidatableFlag    string    `bun:"consolidatable_flag,type:char(1),nullzero"`
	ContainerInvMastUid   int32     `bun:"container_inv_mast_uid,type:int,nullzero"`   // Key to item table for fillable container item ID.
	ContainerUom          string    `bun:"container_uom,type:varchar(8),nullzero"`     // UOM to used for creating adjustment for fillable container used.
	ContainerFlag         string    `bun:"container_flag,type:char(1),nullzero"`       // Flag indicating this process step has a fillable container.
	CommentProcess        string    `bun:"comment_process,type:char(1),default:('N')"` // indicate if this process is comment process or not
	Comment               string    `bun:"comment,type:varchar(max),nullzero"`         // let user enter comment
}
