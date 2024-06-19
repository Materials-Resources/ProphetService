package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLine struct {
	bun.BaseModel               `bun:"table:oe_line"`
	OrderNo                     string    `bun:"order_no,type:varchar(8),pk"`                                // What order does this invoice belong to?
	UnitPrice                   float64   `bun:"unit_price,type:decimal(19,9),nullzero"`                     // Price of the line item in sales units.
	QtyOrdered                  float64   `bun:"qty_ordered,type:decimal(19,9),nullzero"`                    // Quantity ordered.
	QtyPerAssembly              float64   `bun:"qty_per_assembly,type:decimal(19,9)"`                        // If an assembly component, the quantity per assembly.
	CompanyNo                   string    `bun:"company_no,type:varchar(8)"`                                 // Unique code that identifies a company.
	DeleteFlag                  string    `bun:"delete_flag,type:char(1)"`                                   // Indicates whether this record is logically deleted
	ManualPriceOveride          string    `bun:"manual_price_overide,type:char(1),nullzero"`                 // Indicates whether the price has been edited.
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`                        // ID of the user who last maintained this record
	ExtendedPrice               float64   `bun:"extended_price,type:decimal(19,4),nullzero"`                 // The total price of the order line item.
	SalesTax                    float64   `bun:"sales_tax,type:decimal(19,4),nullzero"`                      // What is the sales tax on this item?
	LineNo                      float64   `bun:"line_no,type:decimal(19,0),pk"`                              // What line is this row?
	Complete                    string    `bun:"complete,type:char(1),nullzero"`                             // When Y, this line item has been fully invoiced/canceled.
	UnitOfMeasure               string    `bun:"unit_of_measure,type:varchar(8),nullzero"`                   // What is the unit of measure for this row?
	BaseUtPrice                 float64   `bun:"base_ut_price,type:decimal(19,9),nullzero"`                  // Price before any multipliers / quantity breaks are applied.
	CalcType                    string    `bun:"calc_type,type:varchar(10),nullzero"`                        // Multiplier/markup/difference/etc. Used to describe the calc_value.
	CalcValue                   float64   `bun:"calc_value,type:decimal(19,4),nullzero"`                     // Amount to apply against the base_ut_price to yield a unit_price.
	Combinable                  string    `bun:"combinable,type:char(1),nullzero"`                           // Indicates that the item is eligible for combinable discounting.
	Disposition                 string    `bun:"disposition,type:char(1),nullzero"`                          // The status of unallocated material.
	QtyAllocated                float64   `bun:"qty_allocated,type:decimal(19,9),nullzero"`                  // The quantity allocated but not picked for the order line.
	QtyOnPickTickets            float64   `bun:"qty_on_pick_tickets,type:decimal(19,9),nullzero"`            // The quantity picked but not invoiced for the order line.
	QtyInvoiced                 float64   `bun:"qty_invoiced,type:decimal(19,9),nullzero"`                   // The quantity invoiced for the order line.
	ExpediteDate                time.Time `bun:"expedite_date,type:datetime,nullzero"`                       // The date cutoff for purchasing the material if unallocated.
	RequiredDate                time.Time `bun:"required_date,type:datetime,nullzero"`                       // When is this purchase order line item required by?
	SourceLocId                 float64   `bun:"source_loc_id,type:decimal(19,0),nullzero"`                  // What is the source location of this line item?
	ShipLocId                   float64   `bun:"ship_loc_id,type:decimal(19,0),nullzero"`                    // Where should this order be shipped to?
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                    // What supplier supplies material for this stage?
	ProductGroupId              string    `bun:"product_group_id,type:varchar(8),nullzero"`                  // The product group of the order line item.
	Assembly                    string    `bun:"assembly,type:char(1),nullzero"`                             // Indicates whether the order line is an assembly item.
	Scheduled                   string    `bun:"scheduled,type:char(1),nullzero"`                            // Indicates whether scheduled releases exist for the order line item.
	LotBill                     string    `bun:"lot_bill,type:char(1)"`                                      // Indicates whether the order line item is a lot bill header.
	ExtendedDesc                string    `bun:"extended_desc,type:varchar(255),nullzero"`                   // Additional description for the order line which defaults from the inventory master record.
	UnitSize                    float64   `bun:"unit_size,type:decimal(19,4)"`                               // The size of the sales unit of measure.
	UnitQuantity                float64   `bun:"unit_quantity,type:decimal(19,9)"`                           // The quantity ordered in terms of sales units.
	NextBreak                   float64   `bun:"next_break,type:decimal(19,9),nullzero"`                     // The quantity necessary to receive the next price break.
	NextUtPrice                 float64   `bun:"next_ut_price,type:decimal(19,9),nullzero"`                  // The unit price of the next price break.
	CustomerPartNumber          string    `bun:"customer_part_number,type:varchar(40)"`                      // The id used to order the item. Could be customer part number, alt code, etc.
	TaxItem                     string    `bun:"tax_item,type:char(1),nullzero"`                             // Indicates the item has sales tax.
	OkToInterchange             string    `bun:"ok_to_interchange,type:char(1),nullzero"`                    // Indicates that this item may be interchanged for a substitute at receiving.
	OtherCharge                 string    `bun:"other_charge,type:char(1)"`                                  // Indicates that the order line item is a charge rather than material.
	SalesCost                   float64   `bun:"sales_cost,type:decimal(19,9),nullzero"`                     // The cost based upon the inventory costing basis.
	CommissionCost              float64   `bun:"commission_cost,type:decimal(19,9),nullzero"`                // The cost to be used to calculate commissions.
	RequestedDownpayment        float64   `bun:"requested_downpayment,type:decimal(19,9),nullzero"`          // The downpayment amount  of the order line item.
	DownpaymentDate             time.Time `bun:"downpayment_date,type:datetime,nullzero"`                    // The date the downpayment was requested.
	WillCall                    string    `bun:"will_call,type:char(1),nullzero"`                            // Indicates that this material will be picked up rather than shipped.
	OtherCost                   float64   `bun:"other_cost,type:decimal(19,9),nullzero"`                     // The other cost of the order line item. Commonly used for rebates.
	DownpaymentRemaining        float64   `bun:"downpayment_remaining,type:decimal(19,9),nullzero"`          // Amount of downpayment remaining for a line item.
	CancelFlag                  string    `bun:"cancel_flag,type:char(1),nullzero"`                          // If this column is set to Y - then the line item has been canceled.
	CommissionCostEdited        string    `bun:"commission_cost_edited,type:char(1),nullzero"`               // If this column is set to Y -  then commission_cost s
	OtherCostEdited             string    `bun:"other_cost_edited,type:char(1),nullzero"`                    // If this column is set to Y -  then other_cost should
	PoCost                      float64   `bun:"po_cost,type:decimal(19,9),nullzero"`                        // The PO Cost if a non-stock, direct ship or special order line.
	QtyCanceled                 float64   `bun:"qty_canceled,type:decimal(19,9),nullzero"`                   // The quantity canceled of the order line item.
	PricingUnit                 string    `bun:"pricing_unit,type:varchar(8),nullzero"`                      // Maintains the pricing unit for the invoice line.
	PricingUnitSize             float64   `bun:"pricing_unit_size,type:decimal(19,4),nullzero"`              // Maintains the pricing unit size.
	QtyStaged                   float64   `bun:"qty_staged,type:decimal(19,9),nullzero"`                     // This column accumulates quantity currently picked
	ItemSource                  int16     `bun:"item_source,type:smallint"`                                  // this field will keep track of what item corresponds to the item id entered in order entry.
	OeHdrUid                    int32     `bun:"oe_hdr_uid,type:int"`                                        // The unique identifier of the associated order header.
	ParentOeLineUid             int32     `bun:"parent_oe_line_uid,type:int"`                                // What is the parent of this invoice line item -  if any?
	OeLineUid                   int32     `bun:"oe_line_uid,type:int"`                                       // Internal unique value to identify an order line.
	PricePageUid                int32     `bun:"price_page_uid,type:int,nullzero"`                           // The page used to price this order line.
	PricingOption               int32     `bun:"pricing_option,type:int"`                                    // Determines how the assembly item is priced to your customer.
	DetailType                  int32     `bun:"detail_type,type:int"`                                       // Indicates the line item type for assembly components and lot groups.
	UserLineNo                  string    `bun:"user_line_no,type:varchar(25),nullzero"`                     // This is a user defined line number. Used for custom software.
	ItemTermsDiscountPct        float64   `bun:"item_terms_discount_pct,type:decimal(19,2),default:(0)"`     // Terms amount of the order line item.
	DefaultInOe                 string    `bun:"default_in_oe,type:char(1),default:('N')"`                   // Indicates that this order line appears on a new order by default.
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int"`                                      // Unique identifier for the item id.
	FreightIn                   float64   `bun:"freight_in,type:decimal(19,4),default:(0)"`                  // In-bound freight amount.
	PeerOeLineUid               int32     `bun:"peer_oe_line_uid,type:int,default:(0)"`                      // Links order lines together by the unique identifier.
	PeerType                    int32     `bun:"peer_type,type:int,default:(0)"`                             // Reason why order lines are linked together ex. accessory items.
	SubstituteItem              string    `bun:"substitute_item,type:char(1),default:('N')"`                 // Indicates this item is a substitute.
	AllocateUsageToOriginalItem string    `bun:"allocate_usage_to_original_item,type:char(1),default:('N')"` // When Y usage will be accumulated for the original line item.
	OrderCostEdited             string    `bun:"order_cost_edited,type:char(1),default:('N')"`               // Indicates that the other cost column was edited.
	CaptureUsage                string    `bun:"capture_usage,type:char(1)"`                                 // Indicates if the lines will capture usage or not.
	PurchaseClassId             string    `bun:"purchase_class_id,type:varchar(8),nullzero"`                 // Purchase Class ID for this item.
	PickDate                    time.Time `bun:"pick_date,type:datetime,nullzero"`                           // Stores the date on which the OE line record would be picked
	ShippingRouteUid            int32     `bun:"shipping_route_uid,type:int,nullzero"`                       // Shipping route UID
	InvXrefUid                  int32     `bun:"inv_xref_uid,type:int,nullzero"`                             // Link to the unique identifier for customer part numbers
	HoseQtyNeeded               float64   `bun:"hose_qty_needed,type:decimal(19,9),default:(0)"`             // The quantity per assembly for hose and hose sleeve type assembly components.
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OriginalQtyOrdered          float64   `bun:"original_qty_ordered,type:decimal(19,9),nullzero"`      // Used for custom. Qty ordered before overreceipt of material.
	TagHoldClassUid             int32     `bun:"tag_hold_class_uid,type:int,nullzero"`                  // Tag and Hold Type
	CostCenter                  string    `bun:"cost_center,type:varchar(255),nullzero"`                // stores the value of cost center
	DispositionEditedFlag       string    `bun:"disposition_edited_flag,type:char(1),nullzero"`         // Flag indicating if the oe line disposition has been manually edited.  Valid values are Y, N and NULL.
	JobPriceHdrUid              int32     `bun:"job_price_hdr_uid,type:int,nullzero"`                   // Provides the link to the job based pricing data associated w/the line item
	DeaRestrictionFailed        string    `bun:"dea_restriction_failed,type:char(1),nullzero"`          // Indicates DEA item failed validations.
	RestockFeePercentage        float64   `bun:"restock_fee_percentage,type:decimal(19,9),nullzero"`    // Used to determine the restock fee for an item on an RMA
	CustomerConfiguredPrice     float64   `bun:"customer_configured_price,type:decimal(19,9),nullzero"` // Used by custom software to store different price data depending on implementation.
	CustomerPickedFlag          string    `bun:"customer_picked_flag,type:char(1),nullzero"`            // Allow the user to indicate that the item has been picked or brought to the front counter.
	JobPriceLineUid             int32     `bun:"job_price_line_uid,type:int,nullzero"`                  // Unique Identifier for Job Price Line table
	CustPoNo                    string    `bun:"cust_po_no,type:varchar(255),nullzero"`                 // Customer PO Number related to this OE Line record.
	JobPriceBinUid              int32     `bun:"job_price_bin_uid,type:int,nullzero"`                   // Unique Identifier for Job Price Bin table
	UseContractCost             string    `bun:"use_contract_cost,type:char(1),default:('Y')"`          // Y(es)/N(o) field to tell whether cost can come from contract.
	CostPricePageUid            int32     `bun:"cost_price_page_uid,type:int,nullzero"`                 // Unique Identifier of the Price Page from which cost was calculated.
	PackageTypeUid              int32     `bun:"package_type_uid,type:int,nullzero"`                    // Foreign key to the package_type table.  This is the pacakge_type that the user wants to use when looking for tags.
	TagPackageQtyPer            float64   `bun:"tag_package_qty_per,type:decimal(19,0),nullzero"`       // Holds the tag quantity per package.  This is the pacakge_type quantity that the user wants to use when looking for tags.
	FreightCodeUid              int32     `bun:"freight_code_uid,type:int,nullzero"`                    // Freight code for this line.
	AcknowledgementDate         time.Time `bun:"acknowledgement_date,type:datetime,nullzero"`           // According to the spec for Feature 22107 (custom), the aknowledgement date is a promised date
	UsedSpecificCostFlag        string    `bun:"used_specific_cost_flag,type:char(1),default:('N')"`    // Column will indicate if item used specific cost or not for 'S' disp items.
	OperationUid                int32     `bun:"operation_uid,type:int,nullzero"`                       // Indicates the line is for material (code: 1578) or for labor (code: 1577).
	ServiceLaborUid             int32     `bun:"service_labor_uid,type:int,nullzero"`                   // Unique identifier of labor associated with this assembly.
	QuoteOutcomeCd              int32     `bun:"quote_outcome_cd,type:int,nullzero"`                    // Provides the ability to define the outcome of a quote line as a win, loss or none.
	ItemCommissionClassId       string    `bun:"item_commission_class_id,type:varchar(8),nullzero"`
	FromAltPrevRequestsFlag     string    `bun:"from_alt_prev_requests_flag,type:char(1),nullzero"`         // Indicates if this line was added via the alternate previous requests window
	UdCost                      float64   `bun:"ud_cost,type:decimal(19,9),nullzero"`                       // User defined cost
	TankName                    string    `bun:"tank_name,type:varchar(255),nullzero"`                      // Specifies the associated calendar based delivery tank name.
	ItemBillToId                float64   `bun:"item_bill_to_id,type:decimal(19,0),nullzero"`               // Customer id that is the responsible billing party for this line item
	BuybackNo                   string    `bun:"buyback_no,type:varchar(255),nullzero"`                     // Buyback number for this line item
	RoutingStatusFlag           string    `bun:"routing_status_flag,type:char(1),nullzero"`                 // Character flag that will be used to determine the status of the routing flat
	SplitFlag                   string    `bun:"split_flag,type:char(1),nullzero"`                          // This flag will be informational and let the user know if unallocated qty will be split out
	AdditionalDescription       string    `bun:"additional_description,type:varchar(255),nullzero"`         // Custom -  free form field that is not validated - used for additional information to further define the item for the contractors.  For feature 23262.
	BillHoldFlag                string    `bun:"bill_hold_flag,type:char(1),nullzero"`                      // Used to indicate if the line is  to be processed as a  Bill and Hold
	CorePrice                   float64   `bun:"core_price,type:decimal(19,9),nullzero"`                    // Custom column for the extended price for core item
	Clock                       string    `bun:"clock,type:varchar(255),nullzero"`                          // Used to track specific custom line item data on sales order
	Cell                        string    `bun:"cell,type:varchar(255),nullzero"`                           // Used to track specific custom line item data on sales order
	EnvironmentalFee            float64   `bun:"environmental_fee,type:decimal(19,4),nullzero"`             // Environmental Fee related to the order line
	AdminFee                    float64   `bun:"admin_fee,type:decimal(19,4),nullzero"`                     // Admin Fee related to order line
	ExtDiscAmt                  float64   `bun:"ext_disc_amt,type:decimal(19,9),nullzero"`                  // Custom (F33700): extended discount amount.  Amount by which this line item has been discounted.
	SystemCalcUnitPrice         float64   `bun:"system_calc_unit_price,type:decimal(19,9),nullzero"`        // System calculated unit price
	Buyer                       string    `bun:"buyer,type:varchar(255),nullzero"`                          // Tracks the buyer for the order line
	Recipient                   string    `bun:"recipient,type:varchar(255),nullzero"`                      // Tracks the recipient for the order line
	StrategicUnitPrice          float64   `bun:"strategic_unit_price,type:decimal(19,9),nullzero"`          // Unit price as calculated from the strategic pricing library
	StrategicPricePageUid       int32     `bun:"strategic_price_page_uid,type:int,nullzero"`                // Indicate price page the stragegic unit price comes from
	VerifiedFlag                string    `bun:"verified_flag,type:char(1),nullzero"`                       // Indicates line has been verified - custom
	VerifiedCode                string    `bun:"verified_code,type:varchar(255),nullzero"`                  // The code entered for a verified line item. - custom
	CoreStatusCd                int32     `bun:"core_status_cd,type:int,nullzero"`                          // Strategic Price Core Status Code used to calculate the strategic price.  Core Status is based on the item.
	UsedStrategicPricingFlag    string    `bun:"used_strategic_pricing_flag,type:char(1),nullzero"`         // Indicates if this line was priced via strategic pricing in the past.
	LoanUid                     int32     `bun:"loan_uid,type:int,nullzero"`                                // Loan that the line is associated with
	LoanItemUid                 int32     `bun:"loan_item_uid,type:int,nullzero"`                           // Loan Item that the line is associated with
	OriginalUnitPrice           float64   `bun:"original_unit_price,type:decimal(19,9),default:((0))"`      // Original Unit Price of the Item
	OriginalFreightIn           float64   `bun:"original_freight_in,type:decimal(19,9),nullzero"`           // Freight value orginally entered for a line if strategic pricing is enabled.
	StrategicListCostValue      float64   `bun:"strategic_list_cost_value,type:decimal(19,9),nullzero"`     // Strategic List Price or Strategic Cost based on what is defined in the Strategic Pricing Library that is used.
	StrategicListCostCd         int32     `bun:"strategic_list_cost_cd,type:int,nullzero"`                  // Defines the type of the strategic_list_cost_value field.
	SalesMarketGroupUid         int32     `bun:"sales_market_group_uid,type:int,nullzero"`                  // Custom (F41628): FK to sales_market_group.sales_market_group_uid.  Link to associated sales marketing group.  Populated when a RMA line is linked to an invoice line.
	SpecialItemIndicator        int32     `bun:"special_item_indicator,type:int,nullzero"`                  // Indicator for special items - 1579 for non stockable, but buyable and sellable.
	AwardedDate                 time.Time `bun:"awarded_date,type:datetime,nullzero"`                       // The date the quote line was converted to an order line
	SampleFlag                  string    `bun:"sample_flag,type:char(1),nullzero"`                         // Indicates if this line was given to the customer as a sample product.
	RoutingAllocationChange     float64   `bun:"routing_allocation_change,type:decimal(19,9),nullzero"`     // This column is used to hold the allocated quantity that has been changed since the line was sent to Geocom.
	Price1                      float64   `bun:"price1,type:decimal(19,9),nullzero"`                        // Item Price 1 stored at time of order creation
	UnitPickFee                 float64   `bun:"unit_pick_fee,type:decimal(19,9),nullzero"`                 // Custom column for additional price needs to be added to unit_price
	OeLineAltCode               string    `bun:"oe_line_alt_code,type:varchar(40),nullzero"`                // Custom column to store the alternate code the user used when order an item in Order Entry window.
	RetailPrice                 float64   `bun:"retail_price,type:decimal(19,9),nullzero"`                  // Custom (F45532): holds the retail item price
	QuotePriceOeLineUid         int32     `bun:"quote_price_oe_line_uid,type:int,nullzero"`                 // Uid for quote line that price on this order was set from.
	CoreItemCost                float64   `bun:"core_item_cost,type:decimal(19,9),nullzero"`                // Custom column for core item supplier cost.
	ItemCommitmentDetailUid     int32     `bun:"item_commitment_detail_uid,type:int,nullzero"`              // Unique Identifier for a specific Item Commitment
	SuppressCustomCost          string    `bun:"suppress_custom_cost,type:char(1),nullzero"`                // Indicates whether custom commission calculation rules would be used for an order line.
	BuyListApprovalInitial      string    `bun:"buy_list_approval_initial,type:varchar(3),nullzero"`        // Custom column to store the initial for whoever overrides the buy list item.
	GeocomForcedSendFlag        string    `bun:"geocom_forced_send_flag,type:char(1),default:('N')"`        // Indicates that the line was forced to geocom manually
	GeocomForcedQuantity        float64   `bun:"geocom_forced_quantity,type:decimal(19,0),nullzero"`        // The quantity that was forced to geocom to route
	SuppressPickTicketIndicator int32     `bun:"suppress_pick_ticket_indicator,type:int,nullzero"`          // Custom: Indicates whether item should be suppressed from being picked by the pick ticket scan
	BeltingWidth                float64   `bun:"belting_width,type:decimal(19,9),nullzero"`                 // F46678: Belting functionality - width of the cut that is needed on the belt grid.
	BeltingLength               float64   `bun:"belting_length,type:decimal(19,9),nullzero"`                // F46678: Belting functionality - height of the cut that is needed on the belt grid.
	BeltingNumCuts              int32     `bun:"belting_num_cuts,type:int,nullzero"`                        // F46678: Belting functionality - number of cuts that are required at the belting_width and belting_height.
	CarrierId                   float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`                    // Stores the carrier for this order line.
	FreightInEditedFlag         string    `bun:"freight_in_edited_flag,type:char(1),nullzero"`              // Customer (46747(: determines if the freight_in column has been manually edited.
	SalesDiscountGroupId        string    `bun:"sales_discount_group_id,type:varchar(8),nullzero"`          // Sales Discount Group
	PriceFamilyUid              int32     `bun:"price_family_uid,type:int,nullzero"`                        // Price Family UID
	CostCarrierContractLineUid  int32     `bun:"cost_carrier_contract_line_uid,type:int,nullzero"`          // pecifies carrier contract line used to cost this line.  FK to carrier_contract_line
	PriceCarrierContractLineUid int32     `bun:"price_carrier_contract_line_uid,type:int,nullzero"`         // pecifies carrier contract line used to price this line.  FK to carrier_contract_line
	PumpOffFlag                 string    `bun:"pump_off_flag,type:char(1),default:('N')"`                  // Pegmost pump off flag
	CostCarrierContractZLineUid int32     `bun:"cost_carrier_contract_z_line_uid,type:int,nullzero"`        // Specifies carrier contract z quote line used to cost this line.  FK to carrier_contract_z_line
	TargetPrice                 float64   `bun:"target_price,type:decimal(19,9),nullzero"`                  // column that stored the Target Price (Contract Price – Rebate Amount)
	Edi830LastShipmentNo        string    `bun:"edi830_last_shipment_no,type:varchar(255),nullzero"`        // Custom: populated from the EDI 830 planning schedule import - the last shipment/invoice number received by the trading partner.
	Edi830LastReceiveQty        float64   `bun:"edi830_last_receive_qty,type:decimal(19,9),nullzero"`       // Custom: populated from the EDI 830 planning schedule import - the last shipment/invoice quantity received by the trading partner.
	Edi830LastReceiveDate       time.Time `bun:"edi830_last_receive_date,type:datetime,nullzero"`           // Custom: populated from the EDI 830 planning schedule import - the last date a shipment/invoice was received by the trading partner.
	FaspacYtdQtyInvoiced        float64   `bun:"faspac_ytd_qty_invoiced,type:decimal(19,9),nullzero"`       // Custom: holds the YTD quantity invoiced from the FASPAC system for conversions.
	CompetitorUid               int32     `bun:"competitor_uid,type:int,nullzero"`                          // Unique Identifier for Competitor
	PriceAdjNote                string    `bun:"price_adj_note,type:varchar(255),nullzero"`                 // Price Adjustment Note when line's final price is lower than the net price
	ImportedPrice               float64   `bun:"imported_price,type:decimal(19,9),nullzero"`                // Price of Imported Order through EDI.
	GlCode                      string    `bun:"gl_code,type:varchar(255),nullzero"`                        // order line item GL code associate with this record
	RfqIndicatorFlag            string    `bun:"rfq_indicator_flag,type:char(1),default:('Y')"`             // Flag indicating if the oe line has a RFQ Indicator.
	BuyGetRewardsFlag           string    `bun:"buy_get_rewards_flag,type:char(1),default:('N')"`           // Flag to determine if this line has an associated buy get reward
	BuyGetXRewardsProgramUid    int32     `bun:"buy_get_x_rewards_program_uid,type:int,nullzero"`           // Unique identifier to a buy_get_x_rewards_program record
	CarrierRebateCost           float64   `bun:"carrier_rebate_cost,type:decimal(19,9),nullzero"`           // Stored MAC -Rebate
	ExtendedDescLocation        string    `bun:"extended_desc_location,type:varchar(255),nullzero"`         // Additional description for the order line which defaults from the inventory master record in terms of the Location's or Company's Language
	ExtendedDescCustomer        string    `bun:"extended_desc_customer,type:varchar(255),nullzero"`         // Additional description for the order line which defaults from the inventory master record in terms of the Customer's Language.
	LineDiscount                float64   `bun:"line_discount,type:decimal(19,9),nullzero"`                 // Order Line Discount
	LineDiscountDescription     string    `bun:"line_discount_description,type:varchar(255),nullzero"`      // Order Line Discount Description
	UnitDistributorNet          float64   `bun:"unit_distributor_net,type:decimal(19,9),nullzero"`          // Distributor net value per sales unit that qualifies towards rebates.
	ExtendedDistributorNet      float64   `bun:"extended_distributor_net,type:decimal(19,4),nullzero"`      // Line Item's total Distributor Net value that qualifies towards rebates.
	EcoFeeAmount                float64   `bun:"eco_fee_amount,type:decimal(19,9),nullzero"`                // The total eco fee for a line
	BypassProdOrderProcessing   string    `bun:"bypass_prod_order_processing,type:char(1),nullzero"`        // Determine if we should bypass the prod order processing for this assembly item and treat it as a kit.
	GenericCustomDescription    string    `bun:"generic_custom_description,type:varchar(255),nullzero"`     // Stores generic description or custom description entered by user
	LineSeqNo                   int32     `bun:"line_seq_no,type:int,nullzero"`                             // Column will hold the sequence order in which items were entered in OE.
	OeBuyGetRewardsUid          int32     `bun:"oe_buy_get_rewards_uid,type:int,nullzero"`                  // Unique identifier for oe_buy_get_rewards
	RestrictedByAddressFlag     string    `bun:"restricted_by_address_flag,type:char(1),default:('N')"`     // Stored Y/N if exist a restricted class by ship to address
	SecondaryUnitPrice          float64   `bun:"secondary_unit_price,type:decimal(19,9),nullzero"`          // store a second set of prices for all items
	SecondaryExtendedPrice      float64   `bun:"secondary_extended_price,type:decimal(19,9),nullzero"`      // store a second set of prices for all items
	OriginalDisposition         string    `bun:"original_disposition,type:char(1),nullzero"`                // to store the very first disposition assigned to item
	SecondaryManualPriceOveride string    `bun:"secondary_manual_price_overide,type:char(1),default:('N')"` // Stored Y/N if manually secondary prices are changed
	ExtendedItemFlag            string    `bun:"extended_item_flag,type:char(1),nullzero"`                  // This flag marks if the qty_ordered of the assembly component, should be affected when the assembly item qty is changed in order entry.
	PrintDescOnFormsFlag        string    `bun:"print_desc_on_forms_flag,type:char(1),nullzero"`
	PrintPartNo                 string    `bun:"print_part_no,type:varchar(40),nullzero"`
	CourtesyCoreReturnFlag      string    `bun:"courtesy_core_return_flag,type:char(1),nullzero"`
	CurrentbinUid               int32     `bun:"currentbin_uid,type:int,nullzero"`
	QaStatus                    string    `bun:"qa_status,type:varchar(255),nullzero"`
	ServiceItemFlag             string    `bun:"service_item_flag,type:char(1),nullzero"`
	PrediscountUnitPrice        float64   `bun:"prediscount_unit_price,type:decimal(19,9),nullzero"`
	CustPercentageDisc          float64   `bun:"cust_percentage_disc,type:decimal(19,9),nullzero"`        // Contains the discount percentage applied to the unit price, defined on cust maint for feature 57709
	OriginalQtyAllocated        float64   `bun:"original_qty_allocated,type:decimal(19,9),default:((0))"` // Stores the quantity that was originally allocated (if any) when the hold for an order is removed
	PriceLockFlag               string    `bun:"price_lock_flag,type:char(1),nullzero"`                   // Flag value to indicate whether the price fields on this order line can be edited.
	ExtElectronicDiscAmt        float64   `bun:"ext_electronic_disc_amt,type:decimal(19,9),nullzero"`     // Extended Eletronic Discount Amount for the order line
	ServiceItemId               string    `bun:"service_item_id,type:varchar(255),nullzero"`              // Custom: Holds the item that is actually being serviced since inv_mast_uid will be an other charge. NOTE: This is not necessarily going to be a valid P21 Item ID.
	PoContractNumber            string    `bun:"po_contract_number,type:varchar(255),nullzero"`           // Custom: Indicates the contract number to be used on associated purchase order line items.
	ExcludeFromEdi844867Flag    string    `bun:"exclude_from_edi_844_867_flag,type:char(1),nullzero"`     // Flag used to exclude line from EDI exports 844 and 867
	AddToOpenPtFlag             string    `bun:"add_to_open_pt_flag,type:char(1),default:('N')"`          // Custom flag to identify which lines were added to an open pick ticket
	SplitToOeLineUid            int32     `bun:"split_to_oe_line_uid,type:int,nullzero"`                  // Contains the oe_line_uid of the new line that is created during Order Location Switch import if the current line has any qty on pick tickets or invoiced
	SplitFromOeLineUid          int32     `bun:"split_from_oe_line_uid,type:int,nullzero"`                // Contains the oe_line_uid of the existing/original line from which it was split
	FullRolledItemCd            int32     `bun:"full_rolled_item_cd,type:int,nullzero"`                   // Custom: Indicates whether the full rolled item is a factory full roll or a cut full roll.  Null indicates a non-roll or partial roll.
	RequiredTransferShipDate    time.Time `bun:"required_transfer_ship_date,type:datetime,nullzero"`      // Required Ship date for the Transfer
	ReasonId                    string    `bun:"reason_id,type:varchar(5),nullzero"`                      // Custom: Indicates reason for a sample item shipment.
	DamagedItemSerialNumber     string    `bun:"damaged_item_serial_number,type:varchar(40),nullzero"`    // Custom column to save damaged item serial number
	ExportParkerPrinterName     string    `bun:"export_parker_printer_name,type:varchar(255),nullzero"`   // Name of the printer that will be sent to Parker PTS Export
	ExportParkerLabelCopies     int32     `bun:"export_parker_label_copies,type:int,nullzero"`            // Number of copies to print. Used in Parker PTS Export
	CreateTrackaboutLeaseFlag   string    `bun:"create_trackabout_lease_flag,type:char(1),default:('N')"` // Indicator that the order line was added to a TrackAbout lease
	JobPriceShipControlNoUid    int32     `bun:"job_price_ship_control_no_uid,type:int,nullzero"`         // FK to job_price_ship_control_no table
	PricingMultiplierFromLot    float64   `bun:"pricing_multiplier_from_lot,type:decimal(19,9),nullzero"` // Pricing multiplier from lots assigned to the line to the price
	RentalFlag                  string    `bun:"rental_flag,type:char(1),default:('N')"`                  // Defines if an item line will be rented on an order
	RentalReturnQty             float64   `bun:"rental_return_qty,type:decimal(19,9),nullzero"`           // track return qty reported by essentials system.
	RentalPrice                 float64   `bun:"rental_price,type:decimal(19,9),nullzero"`                // Rental price for rental item
	QuoteConversionListPrice    float64   `bun:"quote_conversion_list_price,type:decimal(19,9),nullzero"` // Stores the supplier / location list price (or supplier list price if there is no list price at the supplier / location level) at the time the line was added to the order from a quote or Previous Request.
	QuotedPrice                 float64   `bun:"quoted_price,type:decimal(19,9),nullzero"`                // Stores the price from the original line (quote line or order line if added through Previous Requests).
	ConversionOrderCost         float64   `bun:"conversion_order_cost,type:decimal(19,9),nullzero"`       // Stores the order cost from the original line (quote line or order line if added through Previous Requests).
	DirectShipFreightAmount     float64   `bun:"direct_ship_freight_amount,type:decimal(19,4),nullzero"`  // Custom: Indicates the expected freight associated with line item (D disposition only).
	ReturnCylinderQuantity      int32     `bun:"return_cylinder_quantity,type:int,nullzero"`              // The return quantity that is associated with a trackabout empty cylinder line
	LinkedChargeParentOeLineUid int32     `bun:"linked_charge_parent_oe_line_uid,type:int,nullzero"`      // (Custom F73686) The oe_line_uid of the parent item that this other charge is linked to
	LinkedChargeScaleQtyFlag    string    `bun:"linked_charge_scale_qty_flag,type:char(1),nullzero"`      // (Custom F73686) Indicates that the other charge qty should be automatically set based on its linked parent qty
	SourceQuoteOeLineUid        int32     `bun:"source_quote_oe_line_uid,type:int,nullzero"`              // For an order converted from a quote using the QTO wizard, stores the oe_line_uid of the quote line from which this order line originated.
	RestrictedClassUid          int32     `bun:"restricted_class_uid,type:int,nullzero"`                  // store restricted class uid
}
