package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLine struct {
	bun.BaseModel                `bun:"table:prod_order_line"`
	ProdOrderNumber              float64    `bun:"prod_order_number,type:decimal(19,0),pk"`   // Production order number.
	LineNumber                   float64    `bun:"line_number,type:decimal(19,0),pk"`         // Line number
	QtyToMake                    float64    `bun:"qty_to_make,type:decimal(19,9)"`            // Quantity to make
	QtyCompleted                 float64    `bun:"qty_completed,type:decimal(19,9)"`          // Quantity completed.
	Completed                    string     `bun:"completed,type:char(1)"`                    // Indicates whether the line is completed.
	Cancel                       string     `bun:"cancel,type:char(1)"`                       // Indicates whether the line is canceled.
	Labor                        *float64   `bun:"labor,type:decimal(19,2)"`                  // The labor amount for the line.
	ShipPartialQuantity          string     `bun:"ship_partial_quantity,type:char(1)"`        // Indicates whether to ship less than the quantity to make.
	ShipIncompleteAssemblies     string     `bun:"ship_incomplete_assemblies,type:char(1)"`   // Indicates whether to ship an item that is not fully assembled.
	PrintCompntsOnPickTicket     string     `bun:"print_compnts_on_pick_ticket,type:char(1)"` // Indicates whether components should print on a pick ticket.
	PrintComponentsOnInvoice     string     `bun:"print_components_on_invoice,type:char(1)"`  // Indicates whether components should print on an invoice.
	DefaultDisposition           string     `bun:"default_disposition,type:char(1)"`
	OverallAssemblyLength        *float64   `bun:"overall_assembly_length,type:decimal(19,9)"`                    // Used for hose assemblies. Not currently used.
	DateCreated                  time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified             time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy             string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	UnitOfMeasure                *string    `bun:"unit_of_measure,type:varchar(8)"`                               // What is the unit of measure for this row?
	UnitSize                     float64    `bun:"unit_size,type:decimal(19,4)"`                                  // The size of the unit_of_measure.
	FreeForm                     string     `bun:"free_form,type:char(1)"`                                        // To indicate whether assembly is a free form assembly.
	CumAssemblyCost              *float64   `bun:"cum_assembly_cost,type:decimal(19,9)"`                          // Stores the accumulated assembly cost that has been processed.
	PrintIncompleteAssemblies    string     `bun:"print_incomplete_assemblies,type:char(1),default:('N')"`        // Determines how assemblies print on Production Order forms.
	InvMastUid                   int32      `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	AssemblyForStock             string     `bun:"assembly_for_stock,type:char(1),default:('N')"`                 // Indicates whether the assembly should be treated as a stock item after processing.
	ExtendedDesc                 *string    `bun:"extended_desc,type:varchar(255)"`                               // Extended description of the assembly item.
	CostToDisassemble            float64    `bun:"cost_to_disassemble,type:decimal(19,9),default:(0)"`            // Cost to disassemble this assembly.
	PrintReturnToStockComponents string     `bun:"print_return_to_stock_components,type:char(1),default:('Y')"`   // Indicates whether to print RTS components on forms.
	PriceReturnToStockComponents string     `bun:"price_return_to_stock_components,type:char(1),default:('Y')"`
	HoseAssemblyFlag             string     `bun:"hose_assembly_flag,type:char(1),default:('N')"`      // When (Y)es, signifies that this production order line is a hose/cable type assembly.
	HoseOverallLength            float64    `bun:"hose_overall_length,type:decimal(19,9),default:(0)"` // Overall length of a hose\cable type assembly including fittings\adaptors.
	HoseUom                      *string    `bun:"hose_uom,type:varchar(8)"`                           // FK to column unit_of_measure.unit_id - unit of measure chosen to describe the measurement of the overall length for this production order line.
	CreatedBy                    *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ExtdCompCost                 *float64   `bun:"extd_comp_cost,type:decimal(19,4)"`                       // A cost representing the sum of the components costs
	ExtdCompCostEditedFlag       *string    `bun:"extd_comp_cost_edited_flag,type:char(1)"`                 // This will be a Y/N field indicating whether or not the ext_comp_cost has been edited by the user
	DepositBin                   *string    `bun:"deposit_bin,type:varchar(10)"`                            // bin where all the components will put to when assemble
	QtyOnPickTickets             float64    `bun:"qty_on_pick_tickets,type:decimal(19,9),default:((0))"`    // Assembly item SKU quantities on production pick tickets for use in production order picking.
	QtyConfirmed                 float64    `bun:"qty_confirmed,type:decimal(19,9),default:((0))"`          // SKU qty on pick tickets that have been confirmed for use in production order processing.
	RequiredDate                 *time.Time `bun:"required_date,type:datetime"`                             // The date in which this production order line item required.
	ExpediteDate                 *time.Time `bun:"expedite_date,type:datetime"`                             // The date in which this assembly line item needs to be expedited
	EstimatedFreightCost         *float64   `bun:"estimated_freight_cost,type:decimal(19,9),default:((0))"` // Indicates estimated freight cost for assembly items- column added for custom feature
	ActualFreightCost            *float64   `bun:"actual_freight_cost,type:decimal(19,9),default:((0))"`    // Indicates actual freight cost for assembly items- column added for custom feature
	ManualFreightOverideFlag     *string    `bun:"manual_freight_overide_flag,type:char(1),default:('N')"`  // This will be set to Y if the user changed freight - column added for custom feature
	CommittedStartDate           *time.Time `bun:"committed_start_date,type:datetime"`                      // Date user commits to starting the assembly
	StartDate                    *time.Time `bun:"start_date,type:datetime"`                                // Date production order is printed
	EndDate                      *time.Time `bun:"end_date,type:datetime"`                                  // Actual date assembly is finished
	CommentDate                  *time.Time `bun:"comment_date,type:datetime"`                              // User-defined date
	ProdOrderLineUid             int32      `bun:"prod_order_line_uid,type:int,unique"`                     // Unique Identifier for prod_order_line table
	DesignCompleteFlag           *string    `bun:"design_complete_flag,type:char(1)"`                       // Indicates whether the design of the Assembly has been completed. CAD Drawings, Blue Prints, etc…
	EstimatedProcessCost         *float64   `bun:"estimated_process_cost,type:decimal(19,9)"`               // Estimated Secondary Processing Cost for the Assemblies.
	CountryOfOrigin              *string    `bun:"country_of_origin,type:varchar(255)"`                     // Custom column to store assembly components country of origin rolled up from all its components in Production Order Processing Window.
}
