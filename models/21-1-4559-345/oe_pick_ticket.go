package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OePickTicket struct {
	bun.BaseModel            `bun:"table:oe_pick_ticket"`
	CompanyId                string     `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	PickTicketNo             float64    `bun:"pick_ticket_no,type:decimal(19,0),pk"`                          // Pick ticket number.
	OrderNo                  string     `bun:"order_no,type:varchar(8)"`                                      // What order does this invoice belong to?
	PrintDate                *time.Time `bun:"print_date,type:datetime"`                                      // The date the pick ticket was printed.
	CarrierId                *float64   `bun:"carrier_id,type:decimal(19,0)"`                                 // What is the id of this carrier (if any)?
	TrackingNo               *string    `bun:"tracking_no,type:varchar(40)"`                                  // The carrier tracking number assigned to the shipment.
	Instructions             *string    `bun:"instructions,type:varchar(255)"`                                // Special instructions regarding the shipment.
	FreightOut               *float64   `bun:"freight_out,type:decimal(19,4)"`                                // Freight cost incurred to ship the entire pick ticket to the customer
	FreightIn                *float64   `bun:"freight_in,type:decimal(19,4)"`                                 // Freight cost incurred to ship the item in to the warehouse
	ShipDate                 *time.Time `bun:"ship_date,type:datetime"`                                       // When was this order shipped?
	InvoiceNo                *float64   `bun:"invoice_no,type:decimal(19,0)"`                                 // Invoice number for pick ticket
	PickAndHold              string     `bun:"pick_and_hold,type:char(1)"`                                    // Indicates whether this is a pick and hold.
	DateCreated              time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified         time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy         string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	PrintedFlag              *string    `bun:"printed_flag,type:char(1),default:('Y')"`                       // Indicates whether the pick ticket has been printed.
	LocationId               float64    `bun:"location_id,type:decimal(19,0)"`                                // What is the unique location identifier for this ro
	DeleteFlag               string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DirectShipment           *string    `bun:"direct_shipment,type:char(1)"`                                  // Indicated whether this pick ticket is for a Direct Shipment
	FrontCounterTax          *string    `bun:"front_counter_tax,type:char(1)"`                                // indicates that front counter tax shall be charged
	TotalTax                 *float64   `bun:"total_tax,type:decimal(19,2)"`                                  // Total tax amount for this pick ticket
	FrontCounter             string     `bun:"front_counter,type:char(1)"`                                    // Indicates that the pick ticket was printed from Fr
	InvoiceIdWhenShipped     *string    `bun:"invoice_id_when_shipped,type:varchar(10)"`                      // Invoice Number this pick ticket is associated with
	Auxiliary                string     `bun:"auxiliary,type:char(1),default:('N')"`                          // Will an oe_line will be printed on an auxiliary pi
	InvoiceBatchUid          int32      `bun:"invoice_batch_uid,type:int,default:(1)"`                        // Unique identifier for the invoice batch.
	FreightCodeUid           *int32     `bun:"freight_code_uid,type:int"`                                     // Unique identifier for the freight code.
	DeliveryListStatus       int32      `bun:"delivery_list_status,type:int,default:(1024)"`                  // Indicates if this pick ticket is assigned to a Delivery List
	RetrievedByWms           string     `bun:"retrieved_by_wms,type:char(1),default:('N')"`
	Picker                   *string    `bun:"picker,type:varchar(255)"`                                // Who picked the order
	LineItemSortSequence     int32      `bun:"line_item_sort_sequence,type:int,default:(1)"`            // Indicates the sort sequence of the line items
	ConfirmableRowStatusFlag int32      `bun:"confirmable_row_status_flag,type:int,default:('1980')"`   // Indicates current picking status.
	WeightOverrideFlag       *string    `bun:"weight_override_flag,type:char(1)"`                       // Indicates whether the total weight of all items on the pick ticket will be overridden during freight calculation.
	OverrideValue            *float64   `bun:"override_value,type:decimal(19,9)"`                       // An overriding value to be treated as the total weight of the pick ticket durign freight calculation.
	CreatedBy                string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`    // User who created the record
	OePickTicketTypeCd       *int32     `bun:"oe_pick_ticket_type_cd,type:int"`                         // Type of pick ticket this is
	OutgoingFreightCost      *float64   `bun:"outgoing_freight_cost,type:decimal(19,9)"`                // User entered value used for calculation of the freight_out, currently only used by custom.
	QualityControl           *string    `bun:"quality_control,type:varchar(255)"`                       // User who Quality Controlled for this shipment
	PackingListPrintedFlag   *string    `bun:"packing_list_printed_flag,type:char(1)"`                  // The column is to indicate whether a packing list has been printed or not.
	DeaPickTicketFlag        *string    `bun:"dea_pick_ticket_flag,type:char(1)"`                       // This flag should be set when the pick ticket contains DEA items
	Packer                   *string    `bun:"packer,type:varchar(255)"`                                // Initials of the person that packed the order
	Checker                  *string    `bun:"checker,type:varchar(255)"`                               // Initials of the person that checked the order
	ReviewShipmentFlag       *string    `bun:"review_shipment_flag,type:char(1),default:('N')"`         // Indicateds whether a shipment has been reviewed but not invoiced.
	OriginalFreightOut       *float64   `bun:"original_freight_out,type:decimal(19,9)"`                 // Freight value orginally entered for a pick ticket if strategic pricing is enabled.
	ExportedToRouterFlag     *string    `bun:"exported_to_router_flag,type:char(1)"`                    // If Y then this pick ticket has already been exported to the routing software.
	PriceConfirmedFlag       *string    `bun:"price_confirmed_flag,type:char(1)"`                       // Indicates if the price has been confirmed.
	ShipConfirmedFlag        *string    `bun:"ship_confirmed_flag,type:char(1)"`                        // Indicates if the shipment has been confirmed.
	OrigSortOrder            *int32     `bun:"orig_sort_order,type:int"`                                // Custom column to store the original sort order of the pick ticket
	VicsBolNumber            *string    `bun:"vics_bol_number,type:varchar(17)"`                        // VICS Bill of Lading Number
	OutFreightCost           *float64   `bun:"out_freight_cost,type:decimal(19,9)"`                     // Used for custom functionality to store a secondary outgoing freight cost for reporting.
	FreightOutEditedFlag     *string    `bun:"freight_out_edited_flag,type:char(1)"`                    // Determines if the out freight has been manually edited.
	PrintPricesOnPackinglist *string    `bun:"print_prices_on_packinglist,type:char(1)"`                // Custom column to indicate if the price should be printed on the packing list at pick ticket level.  This value overrides the value from ship_to when printing packing list.
	PregeneratedInvoiceNo    *float64   `bun:"pregenerated_invoice_no,type:decimal(19,0)"`              // Invoice number that is generated at PT creation time and reserved to be used when this PT is invoiced.
	PrepaidFreightOut        *float64   `bun:"prepaid_freight_out,type:decimal(19,9)"`                  // Custom column for the portion of the outgoing freight to be charged to the customer that is prepaid.
	PrintCanadianB3FormsFlag string     `bun:"print_canadian_b3_forms_flag,type:char(1),default:('N')"` // Determines if Canadian B3 customs forms should be avalable to print for this shipment.
	ProNumber                *string    `bun:"pro_number,type:varchar(255)"`                            // Custom column to store the pro number
	PackingListPrintBy       *string    `bun:"packing_list_print_by,type:varchar(30)"`                  // Custom (F51767): user ID who printed the packing list associated with this pick ticket.
	PackingListPrintDate     *time.Time `bun:"packing_list_print_date,type:datetime"`                   // Custom (F51767): timestamp for when the packing list associated with this pick ticket was printed.
	SlabFlag                 string     `bun:"slab_flag,type:char(1),default:('N')"`                    // Determines if this PT contains slab items.
	ShippingAccount          *string    `bun:"shipping_account,type:varchar(255)"`                      // The shipping account AT&Twill send
	SidNo                    *string    `bun:"sid_no,type:varchar(255)"`                                // Identifier AT&T will be sending
	SentToAttDate            *time.Time `bun:"sent_to_att_date,type:datetime"`                          // Date when the pick ticket was last exported.
	UserConfirmedPickFlag    *string    `bun:"user_confirmed_pick_flag,type:char(1)"`                   // Indicates if the user marked the shipment picked and ready to be shipped.
	Scan                     *string    `bun:"scan,type:varchar(255)"`                                  // Used to store scan value from external process.
	ActualFedexFreightOut    *float64   `bun:"actual_fedex_freight_out,type:decimal(19,4)"`             // Value of the actual freight that FedEx is charging for the shipment
	DiffFedexChargeFlag      string     `bun:"diff_fedex_charge_flag,type:char(1),default:('N')"`       // Flag that is set when there is a difference between what is being charged to the user and what is being charged by FedEx.
	PickConfirmedFlag        *string    `bun:"pick_confirmed_flag,type:char(1)"`                        // Custom (F53615): determines if this ticket has been pick confirmed.
	SentToTrackaboutFlag     *string    `bun:"sent_to_trackabout_flag,type:char(1)"`                    // Used with TrackAbout 3rd party integration, indicates that this pick ticket was sent to TrackAbout
	ShippingRouteUid         *int32     `bun:"shipping_route_uid,type:int"`                             // Shipping route associated with the pick ticket
	BolNumber                *string    `bun:"bol_number,type:varchar(255)"`
	DriverId                 *string    `bun:"driver_id,type:varchar(16)"`      // Custom (F60010) Contact ID of the driver assigned to this shipment.
	SalesMarketGroupUid      *int32     `bun:"sales_market_group_uid,type:int"` // Foreign key to table Sales Market Group.
	ArnNumber                *string    `bun:"arn_number,type:varchar(255)"`    // ARN Number (Acquirer reference number)
	ScanPackUid              *int32     `bun:"scan_pack_uid,type:int"`          // Unique identifier for scan_pack table
	ReferenceNo              *string    `bun:"reference_no,type:varchar(255)"`  // Reference number. Added for Amazon big box functionality. Could be used for more generic purpose if needed.
	PickCompleteFlag         *string    `bun:"pick_complete_flag,type:char(1)"` // Custom: Indicates that a user has flagged this pick ticket as completely picked.
	RfnavStopNo              *int32     `bun:"rfnav_stop_no,type:int"`          // Custom (F73204): stop number returned from the RF Navigator system
	RfnavPickStatusCd        *int32     `bun:"rfnav_pick_status_cd,type:int"`   // For the RF Navigator WMS integration: specifies the current pick status.  Valid values contained in code group 2299.
}
