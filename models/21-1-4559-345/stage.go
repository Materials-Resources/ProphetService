package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Stage struct {
	bun.BaseModel         `bun:"table:stage"`
	StageUid              int32     `bun:"stage_uid,type:int,pk"`                                         // FK to stage table.
	StageCode             string    `bun:"stage_code,type:varchar(255)"`                                  // Human-readable short code used to identify a stage.
	StageDescription      string    `bun:"stage_description,type:varchar(8000)"`                          // A brief description of the stage.
	StageWipAccountNumber string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"`            // Work in process account to hold cost of materials in process.
	EstimatedHours        float64   `bun:"estimated_hours,type:decimal(6,2)"`                             // How many hours should it take to complete a stage?
	SupplierId            float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                       // What supplier supplies material for this stage?
	VendorId              float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`                         // What is the unique vendor identifier for this row?
	CompanyId             string    `bun:"company_id,type:varchar(8),nullzero"`                           // Unique code that identifies a company.
	Cost                  float64   `bun:"cost,type:decimal(19,9)"`                                       // Cost of the work associated with the stage.
	CostUom               string    `bun:"cost_uom,type:varchar(8),nullzero"`                             // The unit of measure associated with the cost.
	AllowPartials         string    `bun:"allow_partials,type:char(1)"`                                   // Indicates whether partial quantities can be moved to a new stage.
	PoRequired            string    `bun:"po_required,type:char(1)"`                                      // Is a Purchase Order required when moving material into a stage?
	MinimumPoCost         float64   `bun:"minimum_po_cost,type:decimal(19,9),nullzero"`                   // What is the minimum cost allowed before a Stage Purchase order is created?
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`                                 // Indicates current record status.
	BuyerId               string    `bun:"buyer_id,type:varchar(16),nullzero"`                            // What buy is associated with this stage?
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	StageQtyUom           string    `bun:"stage_qty_uom,type:varchar(32),nullzero"`                       // Describes the unit of measure the material gets processed in. (Ex. Feet -  Pounds)
	DivisionId            float64   `bun:"division_id,type:decimal(19,0),nullzero"`                       // What is the unique identifier for this division?
	CostType              int16     `bun:"cost_type,type:smallint,nullzero"`                              // Unit to determine cost for stage (ie Weight, UOM)
	QtyUnitSize           float64   `bun:"qty_unit_size,type:decimal(19,4)"`                              // Size of the quantity unit of measure
	CostUnitSize          float64   `bun:"cost_unit_size,type:decimal(19,4),nullzero"`                    // Size of the cost unit of measure
	AllLocationsFlag      string    `bun:"all_locations_flag,type:char(1),default:('Y')"`                 // Indicates whether the stage is for a single location or all locations in the company
	LocationId            float64   `bun:"location_id,type:decimal(19,0),nullzero"`                       // Location for the stage
	StagePoDesc           string    `bun:"stage_po_desc,type:varchar(255),nullzero"`                      // Provides a description of a stage for a specific process. Only applicable if stage.stage_po = Y.
	WipInvMastUid         int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`                            // Column holds inv_mast_uid of the item at the point that a process is performed on the item.
	BackflushFlag         string    `bun:"backflush_flag,type:char(1),nullzero"`                          // Determines whether the process will allow Backflushing functionality.
	TravelerRequiredFlag  string    `bun:"traveler_required_flag,type:char(1),default:('N')"`             // Indicates if the system should update the master row_status_flag to print pending when material is moved in.
	ConsolidatableFlag    string    `bun:"consolidatable_flag,type:char(1),default:('N')"`                // Consolidatable Flag
	ContainerInvMastUid   int32     `bun:"container_inv_mast_uid,type:int,nullzero"`                      // Key to item table for fillable container item ID
	ContainerUom          string    `bun:"container_uom,type:varchar(8),nullzero"`                        // UOM to use for creating adjustment for fillable container use
	ContainerFlag         string    `bun:"container_flag,type:char(1),nullzero"`                          // Flag indicating this process step has a fillable container
	CommentProcess        string    `bun:"comment_process,type:char(1),default:('N')"`                    // indicate if this process is comment process or not
	Comment               string    `bun:"comment,type:varchar(max),nullzero"`                            // let user enter comment
}
