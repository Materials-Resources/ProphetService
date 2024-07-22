package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLineComponent struct {
	bun.BaseModel             `bun:"table:prod_order_line_component"`
	ProdOrderLineComponentUid int32      `bun:"prod_order_line_component_uid,type:int,default:(0),pk"` // Unique identifier for the record.
	ProdOrderNumber           float64    `bun:"prod_order_number,type:decimal(19,0),unique"`           // Production order number
	LineNumber                float64    `bun:"line_number,type:decimal(19,0),unique"`                 // Assembly line number
	ComponentNumber           float64    `bun:"component_number,type:decimal(19,0),unique"`            // Component line number
	CompanyId                 string     `bun:"company_id,type:varchar(8)"`                            // Unique code that identifies a company.
	SourceLocationId          float64    `bun:"source_location_id,type:decimal(19,0)"`                 // Location from which the component was sourced.
	ExtendedDesc              *string    `bun:"extended_desc,type:varchar(255)"`                       // Extended description for this production order component.
	QtyRequested              float64    `bun:"qty_requested,type:decimal(19,9)"`                      // Component quantity requested in SKUs.
	QtyAllocated              float64    `bun:"qty_allocated,type:decimal(19,9)"`                      // Component quantity allocated in SKUs.
	QtyCanceled               *float64   `bun:"qty_canceled,type:decimal(19,9)"`                       // Component quantity canceled in SKUs.
	QtyUsed                   *float64   `bun:"qty_used,type:decimal(19,9)"`                           // Component quantity processed in SKUs.
	UnitOfMeasure             string     `bun:"unit_of_measure,type:varchar(8)"`                       // What is the unit of measure for this row?
	UnitSize                  float64    `bun:"unit_size,type:decimal(19,4)"`                          // Size of the unit_of_measure.
	Disposition               *string    `bun:"disposition,type:char(1)"`                              // Disposition of the unallocated quantity.
	SupplierId                float64    `bun:"supplier_id,type:decimal(19,0)"`                        // Supplier of the component
	DivisionId                float64    `bun:"division_id,type:decimal(19,0)"`                        // The division of the supplier.
	Taxable                   string     `bun:"taxable,type:char(1)"`                                  // Indicates whether the component is taxable.
	OtherCharge               string     `bun:"other_charge,type:char(1)"`                             // Indicates whether the component is an other charge item.
	UnitCost                  float64    `bun:"unit_cost,type:decimal(19,9)"`                          // Cost of the component.
	PoCost                    *float64   `bun:"po_cost,type:decimal(19,9)"`                            // PO cost of the component. Used for direct ship, special and non-stock components.
	DateCreated               time.Time  `bun:"date_created,type:datetime"`                            // Indicates the date/time this record was created.
	DateLastModified          time.Time  `bun:"date_last_modified,type:datetime"`                      // Indicates the date/time this record was last modified.
	LastMaintainedBy          string     `bun:"last_maintained_by,type:varchar(30)"`                   // ID of the user who last maintained this record
	UnitsRequested            float64    `bun:"units_requested,type:decimal(19,9)"`                    // Quantity requested in units.
	SubAssembly               string     `bun:"sub_assembly,type:char(1)"`                             // Indicates that this component is an assembly.
	QtyPerAssembly            float64    `bun:"qty_per_assembly,type:decimal(19,9)"`                   // Component quantity per assembly.
	Complete                  string     `bun:"complete,type:char(1)"`                                 // Indicates whether the component is completed.
	InvMastUid                int32      `bun:"inv_mast_uid,type:int"`                                 // Unique identifier for the item id.
	CompExtendedDesc          *string    `bun:"comp_extended_desc,type:varchar(255)"`                  // Extended description of the component.
	ComponentCutLength        float64    `bun:"component_cut_length,type:decimal(19,9),default:(0)"`   // SKU qty needed for hose\cable and hose sleeve type assembly components.
	CutLengthEditedFlag       string     `bun:"cut_length_edited_flag,type:char(1),default:('N')"`     // When enabled (Y), indicates that the component_cut_length has been manually edited.
	QtyNeeded                 float64    `bun:"qty_needed,type:decimal(19,9),default:(0)"`             // The quantity per assembly for hose and hose sleeve type components.
	UsedSpecificCostFlag      string     `bun:"used_specific_cost_flag,type:char(1),default:('N')"`    // Column will indicate if item used specific cost or not for 'S' disp items.
	QtyOnPickTickets          float64    `bun:"qty_on_pick_tickets,type:decimal(19,9),default:((0))"`  // store qty that on pick tickets
	QtyConfirmed              float64    `bun:"qty_confirmed,type:decimal(19,9),default:((0))"`        // SKU qty on pick tickets that have been confirmed for use in production order processing.
	RequiredDate              *time.Time `bun:"required_date,type:datetime"`                           // The date in which this production order line component is required.
	ExpediteDate              *time.Time `bun:"expedite_date,type:datetime"`                           // The date in which this production order line component needs to be expedited
	ServiceLaborUid           *int32     `bun:"service_labor_uid,type:int"`                            // Unique identifier of labor associated with this assembly.
	OperationUid              *int32     `bun:"operation_uid,type:int"`                                // Indicates if the line is for material (code: 1578) or for labor (code: 1577).
	InventoryCost             *float64   `bun:"inventory_cost,type:decimal(19,9)"`                     // An inventory cost of a component based on inventory costing basis
	NewCost                   *float64   `bun:"new_cost,type:decimal(19,9)"`                           // A user specified inventory cost of a component.  The default value is from inventory_cost
	RetrievedByWms            string     `bun:"retrieved_by_wms,type:char(1),default:('N')"`           // column to indicate if the production order component was retrieved by wms
	OriginalCompletionDate    *time.Time `bun:"original_completion_date,type:datetime"`                // Original expected completion date
	CompletionDate            *time.Time `bun:"completion_date,type:datetime"`                         // Current expected completion date
	BackflushFlag             string     `bun:"backflush_flag,type:char(1),default:('N')"`             // Indicate if the component item is a backflush component.
	QtyScrapped               *float64   `bun:"qty_scrapped,type:decimal(19,9)"`                       // This column holds the sku qty that has been scrapped.
	CountryOfOrigin           *string    `bun:"country_of_origin,type:varchar(255)"`                   // Custom column to store item's country of origin in Production Order Processing Window
	BeltingLength             *float64   `bun:"belting_length,type:decimal(19,9)"`                     // Length of the cut for belting components
	BeltingWidth              *float64   `bun:"belting_width,type:decimal(19,9)"`                      // Width of the cut for belting components
	UomConvertedFlag          *string    `bun:"uom_converted_flag,type:char(1)"`                       // Custom - Indicates whether the uom is converted for this component
	LooseShipFlag             *string    `bun:"loose_ship_flag,type:char(1)"`                          // Flag to indicate whether an assembly component on the production order is a loose ship item.
	ExtendedItemFlag          *string    `bun:"extended_item_flag,type:char(1)"`
	UnitCostEdited            string     `bun:"unit_cost_edited,type:char(1),default:('N')"` // indicated if unit_cost get manually edited
	LaborTypeCd               *int32     `bun:"labor_type_cd,type:int"`                      // Labor Type - Direct / Indirect
}
