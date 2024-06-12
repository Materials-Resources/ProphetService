package model

import "github.com/uptrace/bun"

type Area struct {
	bun.BaseModel              `bun:"table:area"`
	Area                       string `bun:"area,type:varchar(40),pk"`
	Customer                   string `bun:"customer,type:char"`
	Supplier                   string `bun:"supplier,type:char"`
	Item                       string `bun:"item,type:char"`
	CustomerPartNo             string `bun:"customer_part_no,type:char"`
	Quote                      string `bun:"quote,type:char"`
	PurchaseOrder              string `bun:"purchase_order,type:char"`
	VendorReturn               string `bun:"vendor_return,type:char"`
	Requisition                string `bun:"requisition,type:char"`
	InventoryTransfer          string `bun:"inventory_transfer,type:char"`
	SalesOrder                 string `bun:"sales_order,type:char"`
	JournalEntry               string `bun:"journal_entry,type:char,nullzero"`
	PoReceipts                 string `bun:"po_receipts,type:char,default:('N')"`
	Invoice                    string `bun:"invoice,type:char,nullzero"`
	JobPricing                 string `bun:"job_pricing,type:char,default:('N')"`
	JobSite                    string `bun:"job_site,type:char,default:('N')"`
	WorkOrder                  string `bun:"work_order,type:char,default:('N')"`
	Vendor                     string `bun:"vendor,type:char,default:('N')"`
	ConsignmentReplenishOeFlag string `bun:"consignment_replenish_oe_flag,type:char,default:('N')"`
	ShipTo                     string `bun:"ship_to,type:char,default:('N')"`
	RoutingMaintenance         string `bun:"routing_maintenance,type:char,default:('N')"`
	ProcessMaintenance         string `bun:"process_maintenance,type:char,default:('N')"`
	ProcessTransaction         string `bun:"process_transaction,type:char,default:('N')"`
	InvoiceLine                string `bun:"invoice_line,type:char,default:('N')"`
	ServiceItem                string `bun:"service_item,type:char,default:('N')"`
	JobControl                 string `bun:"job_control,type:char,default:('N')"`
	WwmsAreaFlag               string `bun:"wwms_area_flag,type:char,default:('N')"`
	SalesOrderLineFlag         string `bun:"sales_order_line_flag,type:char,default:('N')"`
	ServiceOrder               string `bun:"service_order,type:char,default:('N')"`
}
