package gen

import "github.com/uptrace/bun"

type Area struct {
	bun.BaseModel              `bun:"table:area"`
	Area                       string `bun:"area,type:varchar(40),pk"`                                 // Name of the area
	Customer                   string `bun:"customer,type:char(1)"`                                    // Is the note area an available display area for a customer note?
	Supplier                   string `bun:"supplier,type:char(1)"`                                    // Is the note area an available display area for a supplier note?
	Item                       string `bun:"item,type:char(1)"`                                        // Is the note area an available display area for a item note?
	CustomerPartNo             string `bun:"customer_part_no,type:char(1)"`                            // Is the note area an available display area for a customer_part_no note?
	Quote                      string `bun:"quote,type:char(1)"`                                       // Is the note area an available display area for a quote note?
	PurchaseOrder              string `bun:"purchase_order,type:char(1)"`                              // Is the note area an available display area for a purchase_order note?
	VendorReturn               string `bun:"vendor_return,type:char(1)"`                               // Is the note area an available display area for a vendor_return note?
	Requisition                string `bun:"requisition,type:char(1)"`                                 // Is the note area an available display area for a requistion note?
	InventoryTransfer          string `bun:"inventory_transfer,type:char(1)"`                          // Is the note area an available display area for a inventory_transfer note?
	SalesOrder                 string `bun:"sales_order,type:char(1)"`                                 // Is the note area an available display area for a sales_order note?
	JournalEntry               string `bun:"journal_entry,type:char(1),nullzero"`                      // Is the note area an available display area for a journal_entry note?
	PoReceipts                 string `bun:"po_receipts,type:char(1),default:('N')"`                   // Is the note area an available display area for a po_receipts note?
	Invoice                    string `bun:"invoice,type:char(1),nullzero"`                            // This column will indicate which areas in the system display invoice notes.
	JobPricing                 string `bun:"job_pricing,type:char(1),default:('N')"`                   // Ties a job to a note and display the note in rder Entry, RMA, Front Counter & Job Based Pricing windows.
	JobSite                    string `bun:"job_site,type:char(1),default:('N')"`                      // Whether this area can be selected in the Job Site window
	WorkOrder                  string `bun:"work_order,type:char(1),default:('N')"`                    // Whether this area can be selected in the Work Order window
	Vendor                     string `bun:"vendor,type:char(1),default:('N')"`                        // Is the note area an available display area for a vendor note?
	ConsignmentReplenishOeFlag string `bun:"consignment_replenish_oe_flag,type:char(1),default:('N')"` // Is the note area an available display area for a consignment replenishment order note?
	ShipTo                     string `bun:"ship_to,type:char(1),default:('N')"`                       // allows note to be shown in ship to
	RoutingMaintenance         string `bun:"routing_maintenance,type:char(1),default:('N')"`           // Determines should the notes be displayed in Pre Defined Routing Maintenance.
	ProcessMaintenance         string `bun:"process_maintenance,type:char(1),default:('N')"`           // Determines should the notes be displayed in Process Maintenance.
	ProcessTransaction         string `bun:"process_transaction,type:char(1),default:('N')"`           // Determines should the notes be displayed in Process Transaction.
	InvoiceLine                string `bun:"invoice_line,type:char(1),default:('N')"`                  // Is the note area an available display area for a invoice_line note
	ServiceItem                string `bun:"service_item,type:char(1),default:('N')"`                  // This column will indicate which areas in the system display service item notes.
	JobControl                 string `bun:"job_control,type:char(1),default:('N')"`                   // Indicates which areas are available to display job control notes.
	WwmsAreaFlag               string `bun:"wwms_area_flag,type:char(1),default:('N')"`                // Indicates a note area that shows up in wireless windows.
	SalesOrderLineFlag         string `bun:"sales_order_line_flag,type:char(1),default:('N')"`         // Determines if this area will be available for sales order line notes.
	ServiceOrder               string `bun:"service_order,type:char(1),default:('N')"`                 // area type for service order header notes.
}
