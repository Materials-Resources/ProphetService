package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProcessXTransactionDetail struct {
	bun.BaseModel          `bun:"table:process_x_transaction_detail"`
	ProcessXTransDetailUid int32     `bun:"process_x_trans_detail_uid,type:int,pk"`                        // Stage that has a purchase order.
	ProcessXTransactionUid int32     `bun:"process_x_transaction_uid,type:int"`                            // FK to process_x_transaction table
	StageUid               int32     `bun:"stage_uid,type:int"`                                            // Unique -  internal identifier for a stage
	SequenceNo             int16     `bun:"sequence_no,type:smallint"`                                     // User defined sort order of the stages.
	StartDate              time.Time `bun:"start_date,type:datetime,nullzero"`                             // Date the stage was started.
	LagTimeHours           int16     `bun:"lag_time_hours,type:smallint"`                                  // Allows additional time to be added to the time it takes to complete a stage. Updates expected_date from process_x_transaction.
	QtyIn                  float64   `bun:"qty_in,type:decimal(19,9)"`                                     // Quantity currently in a stage.
	QtyOut                 float64   `bun:"qty_out,type:decimal(19,9)"`                                    // Quantity that has moved through the stage.
	QtyLost                float64   `bun:"qty_lost,type:decimal(19,9)"`                                   // Quantity lost performing this stage.
	QtyPartiallyReceived   float64   `bun:"qty_partially_received,type:decimal(19,9)"`                     // Quantity that was received from a Stage PO. Only applicable if stage.allow_partials =
	RowStatusFlag          int16     `bun:"row_status_flag,type:smallint"`                                 // Indicates current record status.
	Cost                   float64   `bun:"cost,type:decimal(19,9)"`                                       // Cost of the stage.
	CostUom                string    `bun:"cost_uom,type:varchar(8),nullzero"`                             // Describes the unit of measure of the cost of the stage. (ex. $10 per BOX)
	StageQtyUom            string    `bun:"stage_qty_uom,type:varchar(8),nullzero"`                        // Describes the unit of measure the material gets processed in. (Ex. Feet -  Pounds)
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	QtySplit               float64   `bun:"qty_split,type:decimal(19,9),default:(0)"`                      // Quantity that was split .
	CostUnitSize           float64   `bun:"cost_unit_size,type:decimal(19,4)"`                             // Unit size for the cost of the stage.
	QtyUnitSize            float64   `bun:"qty_unit_size,type:decimal(19,4)"`                              // Unit size the material gets processed in.
	CostType               int16     `bun:"cost_type,type:smallint,nullzero"`                              // Indicates whether the cost is determined by Weight of the raw item or by Unit of Measure for the raw item.
	PoRequired             string    `bun:"po_required,type:char(1),default:('N')"`                        // Is a Purchase Order required when moving material into a stage?
	AllowPartials          string    `bun:"allow_partials,type:char(1),default:('Y')"`                     // Indicates whether partial quantities can be moved to a new stage.
	StageWipAccountNumber  string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"`            // The GL account affected by the costs associated with a stage while work is in progress.
	StageCd                string    `bun:"stage_cd,type:varchar(255),nullzero"`                           // Code used to identify the stage.
	StageDesc              string    `bun:"stage_desc,type:varchar(8000),nullzero"`                        // A brief description of the stage.
	EstimatedHours         float64   `bun:"estimated_hours,type:decimal(6,2),nullzero"`                    // Estimated hours it should take to complete the stage.
	SupplierId             float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                       // Supplier for stages that require a po.
	VendorId               float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`                         // Vendor for stages that require a po.
	MinPoCost              float64   `bun:"min_po_cost,type:decimal(19,9),nullzero"`                       // Minimum po cost for stages that require a po.
	BuyerId                string    `bun:"buyer_id,type:varchar(16),nullzero"`                            // Buyer for stages that require a po.
	DivisionId             float64   `bun:"division_id,type:decimal(19,0),nullzero"`                       // Division for stages that require a po.
	StagePoDesc            string    `bun:"stage_po_desc,type:varchar(255),nullzero"`                      // Description to print on pos for stages that require a po.
	PurchasingWeight       float64   `bun:"purchasing_weight,type:decimal(19,9),nullzero"`                 // store purchasing weight value for the stage
	WipInvMastUid          int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`                            // Column holds inv_mast_uid of the item at the point that a process is performed on the item.
	BackflushFlag          string    `bun:"backflush_flag,type:char(1),nullzero"`                          // Determines whether the process will allow Backflushing functionality
	TravelerRequiredFlag   string    `bun:"traveler_required_flag,type:char(1),default:('N')"`             // Indicates if the system should update the master row_status_flag to print pending when material is moved in.
	PostedGlPoQty          float64   `bun:"posted_gl_po_qty,type:decimal(19,9),default:((0))"`             // column to hold quantity that was posted to gl for process po
	ConsolidatableFlag     string    `bun:"consolidatable_flag,type:char(1),nullzero"`                     // Indicates whether this process (stage) is eligible to be combined on a consolidated Process PO with processes from other process Transaction.
	ContainerInvMastUid    int32     `bun:"container_inv_mast_uid,type:int,nullzero"`                      // Key to item table for fillable container item ID
	ContainerUom           string    `bun:"container_uom,type:varchar(8),nullzero"`                        // UOM to use for creating adjustment for fillable container use
	ContainerFlag          string    `bun:"container_flag,type:char(1),nullzero"`                          // Flag indicating this process step has a fillable container
	CommentProcess         string    `bun:"comment_process,type:char(1),default:('N')"`                    // indicate if this process is comment process or not
	Comment                string    `bun:"comment,type:varchar(max),nullzero"`                            // let user enter comment
}
