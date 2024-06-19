package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AssemblyHdr struct {
	bun.BaseModel                `bun:"table:assembly_hdr"`
	Backflush                    string    `bun:"backflush,type:char(1)"`                                        // This column is not currently used.
	AssemblyDays                 float64   `bun:"assembly_days,type:decimal(4,1)"`                               // The number of days anticipated to build the assembly.
	DisassemblyDays              float64   `bun:"disassembly_days,type:decimal(4,1)"`                            // Anticipated number of days to disassemble.
	Labor                        float64   `bun:"labor,type:decimal(19,9)"`                                      // Set the standard cost of labor for assembly items in this field.
	LaborCostMultiplier          float64   `bun:"labor_cost_multiplier,type:decimal(19,9)"`                      // Much like a price multiplier can be applied to the base list cost of items in your inventory for certain customers, the labor cost multiplier may be applied to you base labor cost to modify the actual price.
	Class1                       string    `bun:"class1,type:varchar(255),nullzero"`                             // What is the assembly class?
	Class2                       string    `bun:"class2,type:varchar(255),nullzero"`                             // What is the assembly class?
	Class3                       string    `bun:"class3,type:varchar(255),nullzero"`                             // What is the assembly class?
	Class4                       string    `bun:"class4,type:varchar(255),nullzero"`                             // What is the assembly class?
	Class5                       string    `bun:"class5,type:varchar(255),nullzero"`                             // What is the assembly class?
	ShipPartialQuantity          string    `bun:"ship_partial_quantity,type:char(1)"`                            // Indicates if partial quantity can be shipped.
	AllowOeAddComponents         string    `bun:"allow_oe_add_components,type:char(1)"`                          // Determines whether new components can be added in order entry.
	PrintCompntsOnPickTicket     string    `bun:"print_compnts_on_pick_ticket,type:char(1)"`                     // Determines whether components should print on a pick ticket.
	PrintComponentsOnInvoice     string    `bun:"print_components_on_invoice,type:char(1)"`                      // Determines whether components should print on an invoice.
	DefaultDisposition           string    `bun:"default_disposition,type:char(1)"`                              // The disposition setting (e.g., Backorder, Direct Ship, etc.) automatically assigned to items ordered by a specific customer when the order cannot be filled.
	ProductionOrderProcessing    string    `bun:"production_order_processing,type:char(1)"`                      // Determines whether this is an assembly that you build (Y) or a kit (N).
	OverallAssemblyLength        float64   `bun:"overall_assembly_length,type:decimal(19,4)"`                    // This column is not currently used.
	DeleteFlag                   string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                  time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	ShipIncompleteAssemblies     string    `bun:"ship_incomplete_assemblies,type:char(1)"`                       // Indicates whether assemblies not completely build should be shipped.
	AutoCreateProdOrder          string    `bun:"auto_create_prod_order,type:char(1),default:('Y')"`             // Indicates whether production orders should be created in order entry.
	PricingOption                int16     `bun:"pricing_option,type:tinyint,default:(3)"`                       // Indicates whether the assembly is priced by the header or components.
	PrintIncompleteAssemblies    string    `bun:"print_incomplete_assemblies,type:char(1),default:('N')"`        // Determines how assemblies print on Production Order forms.
	InvMastUid                   int32     `bun:"inv_mast_uid,type:int,pk"`                                      // Unique identifier for the item id.
	AssemblyForStock             string    `bun:"assembly_for_stock,type:char(1),default:('N')"`                 // Determines whether, once built, the assembly should be treated as a regular stock item.
	AssemblyUsageAccumulation    string    `bun:"assembly_usage_accumulation,type:char(1),default:('A')"`        // Determines whether usage is accululated when the production order is completed.
	CalcComponentServiceLevel    string    `bun:"calc_component_service_level,type:char(1),default:('N')"`       // Indicates whether the service level should be calculated for components when processing this assembly.
	ExtendedDesc                 string    `bun:"extended_desc,type:varchar(255),nullzero"`                      // Extended description.
	AllowDisassembly             string    `bun:"allow_disassembly,type:char(1),default:('N')"`                  // Determines whether the assembly can be disassembled.
	CostOfDisassembly            float64   `bun:"cost_of_disassembly,type:decimal(19,9),default:(0)"`            // The cost associated with disassembly.
	ProrateCostBy                string    `bun:"prorate_cost_by,type:char(1),default:('A')"`                    // Determines how to spread the disassembly cost among the components.
	PrintReturnToStockComponents string    `bun:"print_return_to_stock_components,type:char(1),default:('Y')"`   // Determines if return to stock (RTS) components print on invoices and pick tickets.  RTS components are components that have been removed from an assembly or kit and returned to stock.
	PriceReturnToStockComponents string    `bun:"price_return_to_stock_components,type:char(1),default:('Y')"`   // Determines if the price of the return to stock (RTS) component is included in the assembly item price if you have assembly pricing set to pricing by components or by assembly using components.
	HoseAssemblyFlag             string    `bun:"hose_assembly_flag,type:char(1),default:('N')"`                 // A flag to indicate an assembly is a hose assembly
	HoseOverallLength            float64   `bun:"hose_overall_length,type:decimal(19,9),default:(0)"`            // Length of a hose
	HoseUnitOfMeasure            string    `bun:"hose_unit_of_measure,type:varchar(8),nullzero"`                 // Unit of measure for the hose length
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	Type                         int32     `bun:"type,type:int,nullzero"`                                      // Custom: F25838 - identifies the assembly type (regular, master or pro-rated).  See the code_p21 table for associated values.
	ProcessUid                   int32     `bun:"process_uid,type:int,nullzero"`                               // Pre-defined Routing Associated with the Assembly
	PrintTravelerWithProdorder   string    `bun:"print_traveler_with_prodorder,type:varchar(1),default:('N')"` // Indicates if traveler should be printed when first production order is printed in production order entry
	DoNotBulkEditFlag            string    `bun:"do_not_bulk_edit_flag,type:char(1),default:('N')"`            // flag to indicate if the assembly will appear in the list
	AutoExpandFlag               string    `bun:"auto_expand_flag,type:char(1),nullzero"`                      // Indicates if kit should be auto-expanded
	BypassOeProdOrderProcessing  string    `bun:"bypass_oe_prod_order_processing,type:char(1),nullzero"`       // Custom: Default value for the bypass prod order option in OE.
	AssemblyClassUid             int32     `bun:"assembly_class_uid,type:int,nullzero"`
	PackBy                       string    `bun:"pack_by,type:char(1),default:('H')"`                       // Packing by header or components
	MinutesToMake                int32     `bun:"minutes_to_make,type:int"`                                 // Minutes to make assembly
	IncludeComponentsOnEdi856    string    `bun:"include_components_on_edi_856,type:char(1),default:('N')"` // include component in export.
}
