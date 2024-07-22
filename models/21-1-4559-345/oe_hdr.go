package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdr struct {
	bun.BaseModel               `bun:"table:oe_hdr"`
	OrderNo                     string     `bun:"order_no,type:varchar(8),pk"`                     // What order does this invoice belong to?
	CustomerId                  float64    `bun:"customer_id,type:decimal(19,0)"`                  // Customer paying invoice -  remitter
	OrderDate                   *time.Time `bun:"order_date,type:datetime"`                        // The date the order was taken.
	Ship2Name                   *string    `bun:"ship2_name,type:varchar(50)"`                     // What organization should this order be shipped to?
	Ship2Add1                   *string    `bun:"ship2_add1,type:varchar(50)"`                     // Ship-to address 1
	Ship2Add2                   *string    `bun:"ship2_add2,type:varchar(50)"`                     // What is the second line of the address that this order should be shipped to?
	Ship2City                   *string    `bun:"ship2_city,type:varchar(50)"`                     // What is the city of the ship to address?
	Ship2State                  *string    `bun:"ship2_state,type:varchar(50)"`                    // What state should this order be shipped to?
	Ship2Zip                    *string    `bun:"ship2_zip,type:varchar(10)"`                      // What postal code should this order be shipped to?
	Ship2Country                *string    `bun:"ship2_country,type:varchar(50)"`                  // What country should this order be shipped to?
	RequestedDate               *time.Time `bun:"requested_date,type:datetime"`                    // The date by which the entire order is needed.
	PoNo                        *string    `bun:"po_no,type:varchar(50)"`                          // Customer PO Number
	Terms                       *string    `bun:"terms,type:varchar(2)"`                           // Default terms id for all invoices associated with this order.
	ShipToPhone                 *string    `bun:"ship_to_phone,type:varchar(20)"`                  // What is the phone number for the location that this will be shipped to?
	DeleteFlag                  string     `bun:"delete_flag,type:char(1)"`                        // Indicates whether this record is logically deleted
	Completed                   *string    `bun:"completed,type:char(1)"`                          // Indicates whether the purchase order or transfer i
	CompanyId                   *string    `bun:"company_id,type:varchar(8)"`                      // Unique code that identifies a company.
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                      // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30)"`             // ID of the user who last maintained this record
	CodFlag                     *string    `bun:"cod_flag,type:char(1)"`                           // Should this order be shipped COD?
	GrossMargin                 *float64   `bun:"gross_margin,type:decimal(19,9)"`                 // Profit percentage of the order.
	ProjectedOrder              *string    `bun:"projected_order,type:char(1)"`                    // If Y then a quote.
	PoNoAppend                  *string    `bun:"po_no_append,type:varchar(20)"`                   // An automatically generated appendage to a customer po number for uniqueness.
	LocationId                  *float64   `bun:"location_id,type:decimal(19,0)"`                  // Where was the used material located?
	CarrierId                   *float64   `bun:"carrier_id,type:decimal(19,0)"`                   // What is the id of this carrier (if any)?
	AddressId                   *float64   `bun:"address_id,type:decimal(19,0)"`                   // What is the address of this contact?
	ContactId                   *string    `bun:"contact_id,type:varchar(16)"`                     // Contact associated with the order.
	CorpAddressId               *float64   `bun:"corp_address_id,type:decimal(19,0)"`              // The corporate address id for this order.
	HandlingChargeReqFlag       *string    `bun:"handling_charge_req_flag,type:char(1)"`           // Should shipping and handling be charged?
	PaymentMethod               *string    `bun:"payment_method,type:varchar(8)"`                  // This column is unused.
	FobFlag                     *string    `bun:"fob_flag,type:char(1)"`                           // free on board value - when freight costs become responsibility of customer
	Class1Id                    *string    `bun:"class_1id,type:varchar(255)"`                     // Order class 1
	Class2Id                    *string    `bun:"class_2id,type:varchar(255)"`                     // Order class 2
	Class3Id                    *string    `bun:"class_3id,type:varchar(255)"`                     // Order class 3
	Class4Id                    *string    `bun:"class_4id,type:varchar(255)"`                     // Order class 4
	Class5Id                    *string    `bun:"class_5id,type:varchar(255)"`                     // Order class 5
	RmaFlag                     *string    `bun:"rma_flag,type:char(1)"`                           // Is this an RMA?
	Taker                       *string    `bun:"taker,type:varchar(30)"`                          // The order taker.
	JobName                     *string    `bun:"job_name,type:varchar(40)"`                       // User defined column for job tracking.
	ThirdPartyBillingFlag       *string    `bun:"third_party_billing_flag,type:char(1)"`           // Is the customer responsible for billing?
	Approved                    *string    `bun:"approved,type:char(1)"`                           // Indicates whether the order is approved.
	SourceLocationId            *float64   `bun:"source_location_id,type:decimal(19,0)"`           // The default source location for the order lines.
	PackingBasis                *string    `bun:"packing_basis,type:varchar(16)"`                  // Indicates if and how to handle multiple shipments.
	DeliveryInstructions        *string    `bun:"delivery_instructions,type:varchar(255)"`         // Delivery instructions that will print on the pick ticket.
	PickTicketType              *string    `bun:"pick_ticket_type,type:varchar(2)"`                // Indicates whether the pick ticket will be priced or unpriced.
	RequestedDownpayment        *float64   `bun:"requested_downpayment,type:decimal(19,4)"`        // Total downpayment.
	DownpaymentInvoiced         *string    `bun:"downpayment_invoiced,type:char(1)"`               // The amount of the downpayment that has been invoiced so far.
	CancelFlag                  *string    `bun:"cancel_flag,type:char(1)"`                        // When Y, the order has been canceled.
	WillCall                    *string    `bun:"will_call,type:char(1)"`                          // Default for oe_line.will_call.
	FrontCounter                *string    `bun:"front_counter,type:char(1)"`                      // Indicates that the order was taken at the counter. Used for taxing purposes.
	ValidationStatus            *string    `bun:"validation_status,type:varchar(8)"`               // This column holds the credit status of an order.
	OeHdrUid                    int32      `bun:"oe_hdr_uid,type:int,unique"`                      // Unique identifier for the record.
	SourceId                    *string    `bun:"source_id,type:varchar(8)"`                       // External ID associated with the order creation ex. web reference number
	SourceCodeNo                int32      `bun:"source_code_no,type:int"`                         // Method of creating the order ex. B2BSeller
	CreditCardHold              *int32     `bun:"credit_card_hold,type:int"`                       // Indicates the credit card status of an order.
	InvoiceBatchUid             int32      `bun:"invoice_batch_uid,type:int,default:(1)"`          // Unique identifier for the invoice batch.
	FreightCodeUid              *int32     `bun:"freight_code_uid,type:int"`                       // Unique identifier for the freight code.
	FreightOut                  float64    `bun:"freight_out,type:decimal(19,4),default:(0)"`      // Amount of outbound freight.
	ShippingRouteUid            *int32     `bun:"shipping_route_uid,type:int"`                     // Unique identifier for the shipping route.
	ExcludeRebates              string     `bun:"exclude_rebates,type:char(1)"`                    // Column will instruct the system whether to calculate order item rebates -  for a particular order -  at invoice time.
	CaptureUsageDefault         string     `bun:"capture_usage_default,type:char(1)"`              // The default value for oe_lines pertaining to capture usage
	JobPriceHdrUid              *int32     `bun:"job_price_hdr_uid,type:int"`                      // Provides the link to the job based pricing data associated w/the order.
	FrontCounterRma             string     `bun:"front_counter_rma,type:varchar(1),default:('N')"` // Indicate this Order is front counter rma or not.
	Ship2EmailAddress           *string    `bun:"ship2_email_address,type:varchar(255)"`           // Email address associated w\Ship_to address
	OrderType                   *int32     `bun:"order_type,type:int"`
	Taxable                     *string    `bun:"taxable,type:char(1)"`                                  // Specifies order taxable status
	ProfitPercent               float64    `bun:"profit_percent,type:decimal(19,9),default:(0.00)"`      // This will allow anyone looking at the order what was profit percent when creating/saving the order
	OrderCostBasis              int32      `bun:"order_cost_basis,type:int,default:(1)"`                 // This will allow anyone looking at the order what order cost basis was used when creating/saving the order
	CurrencyLineUid             *int32     `bun:"currency_line_uid,type:int"`                            // Identifies which currency_line rcd relates to this order.
	InvoiceExchRateSourceCd     *int32     `bun:"invoice_exch_rate_source_cd,type:int"`                  // Indicate which exchange rate will be used when the invoice is created. Order Exchange Rate or current exchange rate.
	RmaExpirationDate           *time.Time `bun:"rma_expiration_date,type:datetime,default:(getdate())"` // Date before which customer must return all material designated on this RMA.
	CreatedBy                   *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	InvoiceNo                   *string    `bun:"invoice_no,type:varchar(10)"`                               // FK to invoice_hdr.invoice_no.  For RMAs, specifies the invoice linked to this record.
	BillToContactId             *string    `bun:"bill_to_contact_id,type:varchar(16)"`                       // Bill to Contact associated with the order.
	TagHoldCancelDate           *time.Time `bun:"tag_hold_cancel_date,type:datetime"`                        // Tag and Hold Cancel Date
	CostCenterTrackingOption    *int32     `bun:"cost_center_tracking_option,type:int"`                      // Determines the cost center tracking option when the order was created.
	SkipProfitExceptionCheck    string     `bun:"skip_profit_exception_check,type:char(1),default:('N')"`    // Indicate whether system should skip the profit exception rule check
	ConsBackorderProcessingFlag *string    `bun:"cons_backorder_processing_flag,type:char(1)"`               // indicate if the order is eligible for consolidated backorder processing
	RestockFeePercentage        *float64   `bun:"restock_fee_percentage,type:decimal(19,9)"`                 // Used to determine the restock fee on an RMA
	DateOrderCompleted          *time.Time `bun:"date_order_completed,type:datetime"`                        // The date an order is marked complete Y
	ValidatedViaOpenOrdersFlag  *string    `bun:"validated_via_open_orders_flag,type:char(1),default:('N')"` // This flag will indicate whether of not an order has been updated via AR Validate Open Orders.
	AcknowledgementDate         *time.Time `bun:"acknowledgement_date,type:datetime"`                        // According to the spec for Feature 22107, the aknowledgement date is a promised date
	ProductGroupCostBasis       *int32     `bun:"product_group_cost_basis,type:int"`                         // Cost basis for product group profit percent
	OriginalPromiseDate         *time.Time `bun:"original_promise_date,type:datetime"`                       // Original promise date
	PromiseDate                 *time.Time `bun:"promise_date,type:datetime"`                                // Current promise date
	RequestedShipDate           *time.Time `bun:"requested_ship_date,type:datetime"`                         // Custom (F23038): date that this order is requested to ship
	ExpectedCompletionDate      *time.Time `bun:"expected_completion_date,type:datetime"`                    // Date service order is expected to be completed
	JobControlFlag              *string    `bun:"job_control_flag,type:char(1)"`                             // Indicates whether job_name references a true job_control_hdr record.
	ApplyBuilderAllowanceFlag   string     `bun:"apply_builder_allowance_flag,type:char(1),default:('N')"`   // Indicates whether builders allowance applies to the order.
	MerchandiseCreditFlag       string     `bun:"merchandise_credit_flag,type:char(1),default:('N')"`        // Flag indicates if RMA will create a Merchandise Credit instead of an invoice
	ReqPymtUponReleaseFlag      string     `bun:"req_pymt_upon_release_flag,type:char(1),default:('N')"`     // Indicates whether pymt is required upon release of items.
	PmDate                      *time.Time `bun:"pm_date,type:datetime"`                                     // Date that the preventative maintenace for the service items on the generated service order should take place.
	DownpaymentPercentage       *float64   `bun:"downpayment_percentage,type:decimal(19,2)"`                 // Percentage of amount back ordered that was requested as downpayment.
	PrepaidInvoiceFlag          *string    `bun:"prepaid_invoice_flag,type:char(1)"`                         // Indicates whether the order was prepaid.
	LandedCostIncludedCd        *int32     `bun:"landed_cost_included_cd,type:int"`                          // Determine if landed cost should be included in the sales unit price for each line.
	ExportedFlag                *string    `bun:"exported_flag,type:char(1)"`                                // Indicates if the order has been exported through the Export Sales Order/Quote Request window.
	WebReferenceNo              *string    `bun:"web_reference_no,type:varchar(255)"`                        // Stores the web reference number associated with the order (for orders imported via B2B)
	UpsCode                     *string    `bun:"ups_code,type:varchar(40)"`                                 // Custom: UPS account number associated with customer order.
	DefaultPricingCd            *int32     `bun:"default_pricing_cd,type:int"`                               /*
		Indicates how to handle service order pricing - 1707 - By Parts
		1708 - By Service Item Using Parts.
	*/
	FreightChargeByMileHdrUid      *int32     `bun:"freight_charge_by_mile_hdr_uid,type:int"`            // Unique identifier to freight charge by mile table that is associated with this order headerrecord.
	FreightMileageAmt              *float64   `bun:"freight_mileage_amt,type:decimal(19,2)"`             // Freight mileage associated with this order header record.
	SupplierOrderNo                *string    `bun:"supplier_order_no,type:varchar(255)"`                // Custom: Order number used by supplier when they order from distributor
	SupplierReleaseNo              *string    `bun:"supplier_release_no,type:varchar(255)"`              // Custom: Release number for the order that was received from a supplier
	RoutedEtaDate                  *time.Time `bun:"routed_eta_date,type:datetime"`                      // ETA of an order that has already been put through geocom routing.
	RouteOverrideDate              *time.Time `bun:"route_override_date,type:datetime"`                  // Data on the order header used during a geocom export
	ExpediteDate                   *time.Time `bun:"expedite_date,type:datetime"`                        // Default item expedite date
	OrderOpenStartDate             *time.Time `bun:"order_open_start_date,type:datetime"`                // Custom column to store the date when saving an order as marked with validation status COD or Hold.
	EnvironmentalFeeFlag           *string    `bun:"environmental_fee_flag,type:char(1),default:('N')"`  // Indicates if the customer will be charged an environmental fee.
	AdminFeeFlag                   *string    `bun:"admin_fee_flag,type:char(1),default:('N')"`          // Indicates if this customer should be charged a handling/admin fee
	PromiseDateExtendedDesc        *string    `bun:"promise_date_extended_desc,type:varchar(8000)"`      // Notes for hdr promise date
	PromiseDateEditedDate          *time.Time `bun:"promise_date_edited_date,type:datetime"`             // Date a promise date was edited
	OrderDiscType                  *int32     `bun:"order_disc_type,type:int"`                           // Custom (F33700): order discount type.  Possible values are 230 (percentage) and 714 (dollar amount).
	OrderDiscFactor                *float64   `bun:"order_disc_factor,type:decimal(19,9)"`               // Custom (F33700): order discount factor.  Will be a dollar amount or percentage based on the order_disc_type column.
	OrderAckPrintedFlag            *string    `bun:"order_ack_printed_flag,type:char(1)"`                // The column is to indicate whether an order acknowledgement has been printed or not.
	PackingListSentFlag            *int32     `bun:"packing_list_sent_flag,type:int"`                    // Indicates how the Packing List would be sent for this Order.
	RmaDeliveryListStatus          *int32     `bun:"rma_delivery_list_status,type:int,default:((1024))"` // Indicates whether a RMA has been assigned to a Delivery List
	StrategicLibraryUid            *int32     `bun:"strategic_library_uid,type:int"`                     // Strategic Pricing Library used to calculate strategic prices on oe_lines
	StrategicLibraryOriginalUid    *int32     `bun:"strategic_library_original_uid,type:int"`            // Original strategic pricing library for the order
	SendPartialOrderFlag           *string    `bun:"send_partial_order_flag,type:char(1),default:('N')"` // Geocom setting to send an order even if all the lines are not ready to route
	SupplierStatus                 *string    `bun:"supplier_status,type:char(1)"`                       // Custom (F36566): Determines the status that the order has with the supplier - U (Unapproved), A (Approved), or R (Rejected).
	OrderPriorityUid               *int32     `bun:"order_priority_uid,type:int"`
	B2bUpsFreightAmount            *float64   `bun:"b2b_ups_freight_amount,type:decimal(19,9)"`            // UPS freight amount that was imported with the order from B2B.
	UserDefinedDate                *time.Time `bun:"user_defined_date,type:datetime"`                      // A date field that the user defines on the order.
	BlindAddressingFlag            *string    `bun:"blind_addressing_flag,type:char(1),default:('N')"`     // Flag indicating if the ship to location uses Blind Addressing logic when printing Picket Tickets and Packing Lists.
	ReplaceCompanyNameFlag         *string    `bun:"replace_company_name_flag,type:char(1),default:('N')"` // Flag indicating if just the company name is replaced when printing Pick Tickets and Packing Lists.
	UseVendorItemTermsFlag         *string    `bun:"use_vendor_item_terms_flag,type:char(1)"`              // If 'Y' and system setting line_item_terms is 'customer terms or vendor\items', line terms will use vendor\item terms. If 'N', use customer terms always for this customer.
	FirstPackingListNo             *float64   `bun:"first_packing_list_no,type:decimal(19,0)"`             // Holds the pick ticket number of the first packing list that was printed.
	SalesMarketGroupUid            *int32     `bun:"sales_market_group_uid,type:int"`                      // Foreign key to table Sales Market Group.
	ShipConfirmedFlag              *string    `bun:"ship_confirmed_flag,type:char(1)"`                     // Indicates if the shipment has been confirmed.
	PriceConfirmedFlag             *string    `bun:"price_confirmed_flag,type:char(1)"`                    // Indicates if the price has been confirmed.
	QuotedFreightOut               *float64   `bun:"quoted_freight_out,type:decimal(19,9)"`                // The originally quoted outgoing freight amount.
	GeocomAppend                   *int32     `bun:"geocom_append,type:int"`
	ApplyFuelSurchargeFlag         *string    `bun:"apply_fuel_surcharge_flag,type:char(1)"`      // A custom column to save apply_fuel_surcharge value on the order
	OverrideFreightCodeFlag        *string    `bun:"override_freight_code_flag,type:char(1)"`     // Custom column to indicate if freight code in Order Entry has been overriden
	ServiceOrderPriorityUid        *int32     `bun:"service_order_priority_uid,type:int"`         // Unique Identifier of the priority of this service order.
	ExcludeFromCreditLimitFlag     *string    `bun:"exclude_from_credit_limit_flag,type:char(1)"` // Custom column to exculde this order from affecting the credit limit
	PackingListFilename            *string    `bun:"packing_list_filename,type:char(1)"`          // Filename specified for Packing List form.
	OverrideMinOrderChargeFlag     *string    `bun:"override_min_order_charge_flag,type:char(1)"` // Indicates that any minimum order charge amount applied to the order should be zero.
	PlacedByName                   *string    `bun:"placed_by_name,type:varchar(255)"`            // Order Placed By Name
	ApprovedForArFlag              *string    `bun:"approved_for_ar_flag,type:char(1)"`           // Indicates that this order is approved for review by AR.
	SecondRouteOverrideDate        *time.Time `bun:"second_route_override_date,type:datetime"`    // Second Geocom Override Date
	Subject                        *string    `bun:"subject,type:varchar(255)"`                   // Custom: F46649 - Used to store subject for service orders - defaulted from tasks.
	ServicePlanUid                 *int32     `bun:"service_plan_uid,type:int"`
	CentralFaxNumber               *string    `bun:"central_fax_number,type:varchar(20)"`                    // Custom field to store modified main fax number of Ship to stored on the Order.
	Url                            *string    `bun:"url,type:varchar(255)"`                                  // Custom field to store modified website address stored on the Order.
	OriginalPackingBasis           *string    `bun:"original_packing_basis,type:varchar(16)"`                // Stores original packing basis when packing basis is placed on hold.
	FreightOutEditedFlag           *string    `bun:"freight_out_edited_flag,type:char(1)"`                   // Customer F46747): determines if the out freight has been manually edited.
	OverrideContactEmailFlag       *string    `bun:"override_contact_email_flag,type:char(1),default:('N')"` // Indicates whether the email address entered in Order Entry should override the contact email for any emailed invoice related to the order.
	OverrideEmailAddress           *string    `bun:"override_email_address,type:varchar(255)"`               // Email address for use when an invoice for the order is sent via email.
	PrintPricesOnPackinglist       *string    `bun:"print_prices_on_packinglist,type:char(1)"`               // Custom column to indicate if the price should be printed on the packing list at order level.  This value overrides the value from ship_to when printing packing list.
	ImportSource                   *string    `bun:"import_source,type:varchar(255)"`                        // Determines the source of an imported order.  Example data might be EDI830, manual, B2B, etc.
	QuoteLockedFlag                *string    `bun:"quote_locked_flag,type:char(1)"`                         // Indicates if the quote can be changed or if it is locked.
	WebShopperId                   *int32     `bun:"web_shopper_id,type:int"`                                // Web shopper ID of order placer
	WebShopperEmail                *string    `bun:"web_shopper_email,type:varchar(255)"`                    // Email address of web shopper
	RecalcScheduledDsPrice         *string    `bun:"recalc_scheduled_ds_price,type:char(1)"`                 // This column indicates whether price has to be recalculated for scheduled direct ship lines for this order or not.
	OrderTypeCust                  *string    `bun:"order_type_cust,type:varchar(255)"`                      // Used to define custom order types
	EngineerId                     *string    `bun:"engineer_id,type:varchar(255)"`                          // Allow the entry of a contact to represent an engineer
	CarrierContractHdrUid          *int32     `bun:"carrier_contract_hdr_uid,type:int"`                      // Indicates the carrier contract in effect for J-Quotes.  FK to carrier_contract_hdr.
	OrderAckPrintPricesFlag        *string    `bun:"order_ack_print_prices_flag,type:char(1)"`               // Field used to determine if the order ack will print prices.
	FreightOutEstimate             *float64   `bun:"freight_out_estimate,type:decimal(19,9)"`                // Custom (F49614): holds the RateLinx out-freight estimate
	FreightChargeEstimate          *float64   `bun:"freight_charge_estimate,type:decimal(19,9)"`             // Custom (F49614): holds the estimated charge associated with the freight_out_estimate.
	DocumentCaptureDate            *time.Time `bun:"document_capture_date,type:datetime"`                    // Time of capture to document management system, used to detect order ack and quotation edits
	LimitMaxShipmentsPerOrder      *float64   `bun:"limit_max_shipments_per_order,type:decimal(19,4)"`       // Maximum Shipments allowed per order
	ReasonCreditMemoCodeUid        *int32     `bun:"reason_credit_memo_code_uid,type:int"`                   // Reason Credit Memo Code associated with the record.
	SourceCreditMemoCodeUid        *int32     `bun:"source_credit_memo_code_uid,type:int"`                   // Source Credit Memo Code associated with the record.
	NetBillingFlag                 string     `bun:"net_billing_flag,type:char(1),default:('N')"`            // Enable Net Billing custom functionality
	NetBillingEdited               string     `bun:"net_billing_edited,type:char(1),default:('N')"`          // Enable Net Billing custom functionality (control field - not visible)
	SchedOrderDiscFlag             *string    `bun:"sched_order_disc_flag,type:char(1)"`                     // Scheduled Order Discount Item field
	WillCallDiscFlag               *string    `bun:"will_call_disc_flag,type:char(1)"`                       // Will Call Discount Item field
	SingleOrderDiscFlag            *string    `bun:"single_order_disc_flag,type:char(1)"`                    // Single Order Discount field
	VolumeDiscFlag                 *string    `bun:"volume_disc_flag,type:char(1)"`                          // Volume Discount field
	ServiceInvoiceNo               *string    `bun:"service_invoice_no,type:varchar(10)"`                    // Indicates Service Order Invoice Number
	Ship2Add3                      *string    `bun:"ship2_add3,type:varchar(50)"`                            // Ship to address line 3
	BillToId                       *float64   `bun:"bill_to_id,type:decimal(19,0)"`                          // Indicates the ID of the finance company that is paying for the order.
	ElectronicOrderFlag            *string    `bun:"electronic_order_flag,type:char(1),default:('N')"`       // Flag to to identify orders submitted electronically or via fax versus those that were phoned in
	CodWithoutRemittanceFlag       *string    `bun:"cod_without_remittance_flag,type:char(1)"`               // Indicates whether to allow a COD order to be approved without a remittance.
	SoTotalPartsPriceOnInvoice     *string    `bun:"so_total_parts_price_on_invoice,type:char(1)"`           // Indicates Service Order Invoice shows total part price or not
	SoTotalLaborPriceOnInvoice     *string    `bun:"so_total_labor_price_on_invoice,type:char(1)"`           // Indicates Service Order Invoice shows total labor price or not
	InquiryFlag                    *string    `bun:"inquiry_flag,type:char(1)"`                              // Determines if this quote is an Inquiry transaction.
	HoldInvoiceFlag                string     `bun:"hold_invoice_flag,type:char(1),default:('N')"`           // Determines if pick tickets associated with this order require special processing.
	DoNotExportToPtsFlag           *string    `bun:"do_not_export_to_pts_flag,type:char(1),default:('N')"`   // Flag which stop send transactions to PTS
	GenericDescOnPackinglistFlag   *string    `bun:"generic_desc_on_packinglist_flag,type:char(1)"`          // Indicates if generic description will be printed on packing lists instead of regular item description
	GenericDescOnInvoiceAckFlag    *string    `bun:"generic_desc_on_invoice_ack_flag,type:char(1)"`          // Indicates if generic description will be printed on invoices and acknowledgements instead of regular item description
	WillCallNotificationFlag       *string    `bun:"will_call_notification_flag,type:char(1)"`               // Determines if email notifications should be sent out for this order
	PrebillingDate                 *time.Time `bun:"prebilling_date,type:datetime"`                          // Prebilling Date
	BlindShipFlag                  *string    `bun:"blind_ship_flag,type:char(1)"`                           // Custom column to indicate the order will be shipped to the customer’s customer as if the shipment was coming from the customer, rather than from the distributor.
	ArchitectId                    *float64   `bun:"architect_id,type:decimal(19,0)"`                        // Used in Builders Selection Sheet. Valid IDs from customer table.
	BuilderId                      *float64   `bun:"builder_id,type:decimal(19,0)"`                          // Used in Builders Selection Sheet. Valid IDs from customer table.
	ContractorId                   *float64   `bun:"contractor_id,type:decimal(19,0)"`                       // Used in Builders Selection Sheet. Valid IDs from customer table.
	DesignerId                     *float64   `bun:"designer_id,type:decimal(19,0)"`                         // Used in Builders Selection Sheet. Valid IDs from customer table.
	HomeownerId                    *float64   `bun:"homeowner_id,type:decimal(19,0)"`                        // Used in Builders Selection Sheet. Valid IDs from customer table.
	QuoteType                      *int32     `bun:"quote_type,type:int"`                                    // Stored Quote Type; None, Regular Quote, Builders Selection Sheet, Sample Checkout
	PtsLabelPrintFlag              *string    `bun:"pts_label_print_flag,type:char(1)"`                      // Flag to save the PTS Label Print checkbox
	DiscountItemId                 *string    `bun:"discount_item_id,type:varchar(255)"`
	CustomerDiscPct                *float64   `bun:"customer_disc_pct,type:decimal(19,9)"`
	ConsolidatePerOrderFlag        *string    `bun:"consolidate_per_order_flag,type:char(1)"`
	DfltUpsShipViaCode             *int32     `bun:"dflt_ups_ship_via_code,type:int"`      // Order default for ups code
	DfltFedexServiceTypeUid        *int32     `bun:"dflt_fedex_service_type_uid,type:int"` // Order default for fedex code
	InvoiceOnlyFlag                *string    `bun:"invoice_only_flag,type:char(1)"`
	Ship2Latitude                  *string    `bun:"ship2_latitude,type:varchar(20)"`  // Latitude for ship to address for Avalara
	Ship2Longitude                 *string    `bun:"ship2_longitude,type:varchar(20)"` // Longitude for ship to address for Avalara
	BypassHandlingChargeFlag       *string    `bun:"bypass_handling_charge_flag,type:char(1)"`
	FreightChargeUid               *int32     `bun:"freight_charge_uid,type:int"`                           // Freight charge to override the customer freight code
	LumpSumBillingFlag             *string    `bun:"lump_sum_billing_flag,type:char(1)"`                    // supress unit & extended price on invoice printing
	FieldDestroyFlag               *string    `bun:"field_destroy_flag,type:char(1)"`                       // RMA Field Destroy
	WarrantyReturnOrderFlag        *string    `bun:"warranty_return_order_flag,type:char(1),default:('N')"` // Warranty return order flag
	AllowAutoApplyOrigInvFlag      *string    `bun:"allow_auto_apply_orig_inv_flag,type:char(1)"`           // When selected, this will prevent the RMA from being automatically applied to the original invoice as a Credit Memo.  The RMA will work like a normal RMA with linked lines.
	DaysEarly                      *int32     `bun:"days_early,type:int"`                                   // Days early for pick dates calculation.
	TransitDays                    *int32     `bun:"transit_days,type:int"`                                 // Transit days for pick dates calculation.
	DisplayInvoiceExchRateSourceCd *int32     `bun:"display_invoice_exch_rate_source_cd,type:int"`          // Indicates the source of the display exchange rate used during invoice creation (order or current).
	NatlAcctCustomerUid            *int32     `bun:"natl_acct_customer_uid,type:int"`                       // Custom (F67104): for customers eligible for national account pricing, determines the national account customer ID that will be employed for pricing
	SlaRequestedDate               *time.Time `bun:"sla_requested_date,type:datetime"`                      // Custom: Original required date calculated using the ship-to or customer's Service Level Agreement.
	EdiManualOverrideFlag          *string    `bun:"edi_manual_override_flag,type:char(1),default:('N')"`   // Custom - Indicates whether to override sending EDI 810/856 on a per order basis.
	RevisioningEnabledFlag         *string    `bun:"revisioning_enabled_flag,type:char(1)"`                 // Custom column to indicate if the order is revision enabled
	QuoteSubmittedFlag             *string    `bun:"quote_submitted_flag,type:char(1)"`                     // Custom column to indicate the quote has been submitted to the customer
	WarrantyRmaFlag                *string    `bun:"warranty_rma_flag,type:char(1),default:('N')"`          // When enabled P21 application will utilize the new “Warranty RMA Receipt Clearing Account” (instead of the existing RMA Receipt Clearing Account)
	GenerateProFormaInvoicesFlag   *string    `bun:"generate_pro_forma_invoices_flag,type:char(1)"`         // Export pro forma invoice in shipping
	AsbDeliveryMethodUid           *int32     `bun:"asb_delivery_method_uid,type:int"`                      // Custom: Indicates auto short buy delivery method associated with this order.
	DeliveryContactId              *string    `bun:"delivery_contact_id,type:varchar(16)"`                  // Used to track the name, address and telephone number of the person who actually be receiving the order form customer side
	AdvancedBillingFlag            *string    `bun:"advanced_billing_flag,type:char(1)"`                    // Indicates if this order was processed through advanced billing.
	AdvancedBillingPrintFlag       *string    `bun:"advanced_billing_print_flag,type:char(1)"`              // Indicates if this advanced bill order has been printed.
	EwingJobNo                     *int32     `bun:"ewing_job_no,type:int"`                                 // Job No exclusive for Ewing
	PricingLocId                   *int32     `bun:"pricing_loc_id,type:int"`                               // (Custom F70122) Pricing Location ID
	RentalTransactionNo            *string    `bun:"rental_transaction_no,type:varchar(255)"`               // No to identify if an order is linked to a Rental transaction.
	RentalQuoteNo                  *string    `bun:"rental_quote_no,type:varchar(255)"`                     // Stores the quote no for the rental transaction related to this order
	RentalTransactionStatusCd      *int32     `bun:"rental_transaction_status_cd,type:int"`                 // id to define each of the statuses from rentals orders
	RentalReturnDate               *time.Time `bun:"rental_return_date,type:datetime"`                      // Defines the time the order items will be rented
	MafSurchargeOverride           string     `bun:"maf_surcharge_override,type:varchar(1),default:('N')"`  // Custom (F71325): Exclude order from MAF calculations
	MafSurchargeReasonCode         *string    `bun:"maf_surcharge_reason_code,type:varchar(255)"`           // Custom (F71325): Reason for exclusion from MAF calculations
	OracleCarrierId                *float64   `bun:"oracle_carrier_id,type:decimal(19,9)"`                  // Carrier ID for Trane used for Oracle and AWS processing
	RentalBillingFlag              *string    `bun:"rental_billing_flag,type:char(1),default:('U')"`        // Flag to indicate rental billing choice - U = Upfront , A = Arrears
	InsideSales                    *string    `bun:"inside_sales,type:varchar(30)"`                         // Sales representative who made the sale.
	Ship2Url                       *string    `bun:"ship2_url,type:varchar(255)"`                           // Store the URL for the order
	GlDimensionProjectNo           *string    `bun:"gl_dimension_project_no,type:varchar(255)"`             // GL Dimension - Project Number
	RentalStartDate                *time.Time `bun:"rental_start_date,type:datetime"`                       // Start Date for a rental order
}
