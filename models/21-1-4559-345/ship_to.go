package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ShipTo struct {
	bun.BaseModel                 `bun:"table:ship_to"`
	CompanyId                     string     `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	ShipToId                      float64    `bun:"ship_to_id,type:decimal(19,0),pk"`                              // What is the ship_to location for this ship to jurisdiction?
	CustomerId                    float64    `bun:"customer_id,type:decimal(19,0)"`                                // Customer ID
	DefaultBranch                 string     `bun:"default_branch,type:varchar(8)"`                                // What is the default branch for this ship-to?
	FederalExemptionNumber        *string    `bun:"federal_exemption_number,type:varchar(40)"`                     // Exemption number for federal tax.
	StateExemptionNumber          *string    `bun:"state_exemption_number,type:varchar(40)"`                       // Exemption number for state tax.
	OtherExemptionNumber          *string    `bun:"other_exemption_number,type:varchar(40)"`                       // Exemption number for other tax.
	PriceFileId                   *float64   `bun:"price_file_id,type:decimal(19,0)"`                              // This column is unused.
	PreferredLocationId           *float64   `bun:"preferred_location_id,type:decimal(19,0)"`                      // Location ID that will be used to source material in Order Entry.
	DefaultShipTime               *time.Time `bun:"default_ship_time,type:datetime"`                               // When should a shipment normally go out?
	DefaultCarrierId              *float64   `bun:"default_carrier_id,type:decimal(19,0)"`                         // What carrier is normally used?
	Fob                           *string    `bun:"fob,type:varchar(20)"`                                          // Free On Board -- point in the delivery process at which freight costs and liability become the responsibility of the customer
	DeliveryInstructions          *string    `bun:"delivery_instructions,type:varchar(255)"`                       // Default delivery instructions.
	MorningBegDelivery            *time.Time `bun:"morning_beg_delivery,type:datetime"`                            // When should morning deliveries begin?
	MorningEndDelivery            *time.Time `bun:"morning_end_delivery,type:datetime"`                            // When should morning deliveries end?
	EveningBegDelivery            *time.Time `bun:"evening_beg_delivery,type:datetime"`                            // When should evening deliveries begin?
	EveningEndDelivery            *time.Time `bun:"evening_end_delivery,type:datetime"`                            // When are evening deliveries over?
	AcceptPartialOrders           string     `bun:"accept_partial_orders,type:char(1)"`                            // Can this address accept partial orders?
	AcceptableWaitTime            *float64   `bun:"acceptable_wait_time,type:decimal(3,0)"`                        // How long can the customer wait for the order.
	Class1Id                      *string    `bun:"class1_id,type:varchar(255)"`                                   // Ship to class 1
	Class2Id                      *string    `bun:"class2_id,type:varchar(255)"`                                   // Ship to class 2
	Class3Id                      *string    `bun:"class3_id,type:varchar(255)"`                                   // Ship to class 3
	Class4Id                      *string    `bun:"class4_id,type:varchar(255)"`                                   // Ship to class 4
	Class5Id                      *string    `bun:"class5_id,type:varchar(255)"`                                   // Ship to class 5
	PackingBasis                  *string    `bun:"packing_basis,type:varchar(16)"`                                // The packing basis for the ship-to.
	DateCreated                   time.Time  `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified              time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy              string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DeleteFlag                    string     `bun:"delete_flag,type:char(1),default:('N')"`                        // Indicates whether this record is logically deleted
	TermsId                       *string    `bun:"terms_id,type:varchar(2)"`                                      // The default terms id.
	InvoiceType                   *string    `bun:"invoice_type,type:varchar(2)"`                                  // What type of invoice is this invoice?
	BilledOnGrossNetQty           *string    `bun:"billed_on_gross_net_qty,type:char(1)"`                          // Indicates whether ship to will be billed on gross or net quantity
	StateExciseTaxExemptionNo     *string    `bun:"state_excise_tax_exemption_no,type:varchar(40)"`                // Exemption number for state excise tax.
	PickTicketType                *string    `bun:"pick_ticket_type,type:varchar(2)"`                              // Determines whether the pick ticket will be priced or unpriced.
	TaxGroupId                    *string    `bun:"tax_group_id,type:varchar(10)"`                                 // Associates a location to a tax group.
	HandlingChargeReqFlag         string     `bun:"handling_charge_req_flag,type:char(1),default:('N')"`           // Is a handling charge applicable to this shipment?
	ThirdPartyBillingFlag         string     `bun:"third_party_billing_flag,type:char(1),default:('S')"`           // Is a third party responsible for paying the shipping charges for this shipment?
	DaysEarly                     *int32     `bun:"days_early,type:int,default:(0)"`                               // indicates the maximum number of days (prior to the required date) that ship-to will accept material.
	DaysLate                      *int32     `bun:"days_late,type:int,default:(0)"`                                // indicates the maximum number of days (past the required dates) the ship-to will accept material.
	TransitDays                   *int32     `bun:"transit_days,type:int,default:(1)"`                             // Number of days material will be in transit.
	FreightCodeUid                *int32     `bun:"freight_code_uid,type:int"`                                     // Unique identifier for the outbound freight code.
	SignatureRequired             string     `bun:"signature_required,type:char(1),default:('N')"`                 // Indicates whether the signature is required for the PDA / delivery feature
	ShippingRouteUid              *int32     `bun:"shipping_route_uid,type:int"`                                   // Unique identifier for the shipping route.
	PdaOrderEntry                 string     `bun:"pda_order_entry,type:char(1),default:('N')"`                    // Define whether a ship to will be used for PDA order entry.
	PdaOelistCriteriaUid          *int32     `bun:"pda_oelist_criteria_uid,type:int"`                              // Foreign key back to Primary Key on pda_oelist_criteria
	ExcludeHoldFromPickTix        string     `bun:"exclude_hold_from_pick_tix,type:varchar(1),default:('N')"`      // Exclude items with H disposition from Pick Tickets
	ExcludeCanceldFromPackList    string     `bun:"exclude_canceld_from_pack_list,type:varchar(1),default:('N')"`  // Exclude items with C disposition from Packing Lists
	ExcludeHoldFromPackList       string     `bun:"exclude_hold_from_pack_list,type:varchar(1),default:('N')"`     // Exclude items with H disposition from Packing Lists
	ExcludeCanceldFromOrderAck    string     `bun:"exclude_canceld_from_order_ack,type:varchar(1),default:('N')"`  // Exclude items with C disposition from Order Acknowledgements
	ExcludeHoldFromOrderAck       string     `bun:"exclude_hold_from_order_ack,type:varchar(1),default:('N')"`     // Exclude items with H disposition from Order Acknowledgements
	IncludeNonAllocOnPickTix      int32      `bun:"include_non_alloc_on_pick_tix,type:int,default:(1277)"`         // Determines if non-allocated items will print on pick tickets
	ExcludeCanceldFromPickTix     string     `bun:"exclude_canceld_from_pick_tix,type:char(1),default:('N')"`      // Determines whether canceled items will print on pick tickets
	CreatedBy                     *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	IncludeNonAllocOnPackList     int32      `bun:"include_non_alloc_on_pack_list,type:int"`                     // When to include non-allocated items on the packing list
	PrintPackinglistInShipping    string     `bun:"print_packinglist_in_shipping,type:char(1),default:('N')"`    // Whether to print the packing list in Shipping
	PrintPricesOnPackinglist      string     `bun:"print_prices_on_packinglist,type:char(1),default:('Y')"`      // Whether to create a priced or unpriced packing list
	SourceTypeCd                  *int32     `bun:"source_type_cd,type:int"`                                     // Indicates how this rcd was created
	DataIdentifierGroupUid        *int32     `bun:"data_identifier_group_uid,type:int"`                          // Data Identifier Group UID for this ship-to (typically used in the generation of sales-order related documents requiring data identifiers)
	FedexFreightMarkup            *float64   `bun:"fedex_freight_markup,type:decimal(19,9)"`                     // Holds the FedEx freight markup amount.
	CreditStatus                  *string    `bun:"credit_status,type:varchar(8)"`                               // this will hold a credit status id it will be a foreign key to the
	CreditLimit                   *float64   `bun:"credit_limit,type:decimal(19,2)"`                             // this will hold the credit limit for this ship to, a monetary value
	CreditLimitUsed               *float64   `bun:"credit_limit_used,type:decimal(19,2)"`                        // this will hold the amount of the ship to's credit limit that has been used
	DefaultCustomerPoNo           *string    `bun:"default_customer_po_no,type:varchar(255)"`                    // What is the default po no for this ship-to?
	BudgetCodeApprovalFlag        *string    `bun:"budget_code_approval_flag,type:char(1)"`                      // Custom (F23038): determines if an order is placed on hold when a job pricing budget code's budgeted amount is exceeded when an item is added to an order.
	ServiceTermsId                *string    `bun:"service_terms_id,type:varchar(2)"`                            // Terms for service orders
	DefaultFreightMarkupPct       float64    `bun:"default_freight_markup_pct,type:decimal(19,4),default:((0))"` // Default freight markup percentage (variable rate handling charge) for shipping.
	FreightChargeByMileHdrUid     *int32     `bun:"freight_charge_by_mile_hdr_uid,type:int"`                     // Unique identifier to freight charge by mile table that is associated with this ship-to record.
	FreightMileageAmt             *float64   `bun:"freight_mileage_amt,type:decimal(19,2)"`                      // Freight mileage associated with this ship-to record.
	DegreeDaysFlag                *string    `bun:"degree_days_flag,type:char(1)"`                               // If checked, the Degree Day Tab will be enabled for the Ship To
	CalendarDaysFlag              *string    `bun:"calendar_days_flag,type:char(1)"`                             // If checked, the Calendar Based Tab will be enabled for the Ship To
	HeatAndHotWaterCd             *int32     `bun:"heat_and_hot_water_cd,type:int"`                              // Code to indicate to track Heat/Heat and Hot Water Degree Days.
	DrumDepositExemptFlag         *string    `bun:"drum_deposit_exempt_flag,type:char(1),default:('N')"`         // This column is set in Ship To Maintenace on the Ship To tab.  If this field is set ON, the drum deposit will be zero.
	TaxableFlag                   *string    `bun:"taxable_flag,type:char(1)"`                                   // Custom column to indicate if ship to id is taxable or not
	BillHoldFlag                  *string    `bun:"bill_hold_flag,type:char(1)"`                                 // If used as the the default for the Bill and Hold flag in the order header
	DisplayPoNoFlag               *string    `bun:"display_po_no_flag,type:char(1)"`                             // This column will determine if the PO number is to be input on the IPAQ hand held device.
	PoNoRequiredFlag              *string    `bun:"po_no_required_flag,type:char(1)"`                            // This column will determine if the PO number is a required field on the IPAQ hand held device.
	DisplayBadgeFlag              *string    `bun:"display_badge_flag,type:char(1)"`                             // This column will determine if the users Badge number is to be input on the IPAQ hand held device.
	BadgeRequiredFlag             *string    `bun:"badge_required_flag,type:char(1)"`                            // This column will determine if the users Badge number is a required field on the IPAQ hand held device.
	DelSchSourceLocation          *float64   `bun:"del_sch_source_location,type:decimal(19,0)"`                  // Default Source Location for Delivery Scheduler.
	ServiceSourceLocation         *float64   `bun:"service_source_location,type:decimal(19,0)"`                  // Default Source Location for Service Orders.
	DistributorAccountId          *string    `bun:"distributor_account_id,type:varchar(255)"`                    // Custom: Account for the distributor on the supplier system by ship to.
	CardlockCalcPriceFlag         *string    `bun:"cardlock_calc_price_flag,type:char(1),default:('N')"`         // Determines if price should be calculated for Cardlock Orders instead of using Cardlock price
	CourtesyAddressId             *float64   `bun:"courtesy_address_id,type:decimal(19,0)"`                      // Address ID to be used in EDI 867 processing for feature 32160, sub ref 1.
	DfoaShipToId                  *string    `bun:"dfoa_ship_to_id,type:varchar(255)"`                           // Custom: Alternate, user defined ship to id.
	ErouterTranType               *string    `bun:"erouter_tran_type,type:char(1)"`                              // Custom: Determines the type of ERouter transactions that this ship to is authorized for - V or D.
	CardlockDiscount              *float64   `bun:"cardlock_discount,type:decimal(19,9)"`                        // Dollar amount to discount orders imported from fuel island sales order import.
	OrderPriorityUid              *int32     `bun:"order_priority_uid,type:int"`
	DeliveryZoneUid               *int32     `bun:"delivery_zone_uid,type:int"`                                   // Foreign key to table Delivery_Zone.
	PlantCode                     *string    `bun:"plant_code,type:varchar(255)"`                                 // Custom: Plant Code that is used by Shell (eRouter Integration) for this Ship To.
	FreightChargeUid              *int32     `bun:"freight_charge_uid,type:int"`                                  // Unique identifier for freight charge
	UpsRoadnetAcctTypeId          *string    `bun:"ups_roadnet_acct_type_id,type:varchar(255)"`                   // UPS Roadnet specific: account type ID
	UpsRoadnetDeliveryDays        *string    `bun:"ups_roadnet_delivery_days,type:varchar(255)"`                  // UPS Roadnet specific: days on which deliveries are made to this ship-to.  May be a combination of the following: 'M'onday, 'T'uesday, 'W'ednesday, thu'R'sday, 'F'riday, 'S'aturday, s'U'nday (e.g. - MWR for Monday, Wednesday and Thursday delivery).
	UpsRoadnetZoneId              *string    `bun:"ups_roadnet_zone_id,type:varchar(255)"`                        // UPS Roadnet specific: zone ID
	UseCustUpsHandlngChrgFlag     string     `bun:"use_cust_ups_handlng_chrg_flag,type:varchar(1),default:('Y')"` // Indicates whether the ship to uses the customer UPS handling charge.
	UpsHandlingCharge             float64    `bun:"ups_handling_charge,type:decimal(19,2),default:((0.00))"`      // Value of handling charge used for UPSOnLine.
	SalesTaxPayableAccountNo      *string    `bun:"sales_tax_payable_account_no,type:varchar(255)"`               // GL account for external tax integration
	VertexTaxableFlag             *string    `bun:"vertex_taxable_flag,type:char(1)"`                             // Indicates if the ship to is taxable by Vertex
	CustomerTaxClass              *string    `bun:"customer_tax_class,type:varchar(255)"`                         // Customer tax class used by Vertex
	TaxExemptReasonUid            *int32     `bun:"tax_exempt_reason_uid,type:int"`                               // The unique identifier for a Ship To's tax exemption reason.
	ServiceCenterUid              *int32     `bun:"service_center_uid,type:int"`                                  // Unique Identifier of a service center that will be the default service center in service order entry.
	DunsNumber                    *string    `bun:"duns_number,type:varchar(20)"`                                 // Holds the DUNS Number of this ship to. Used for reporting data to the vendor Pall.
	SmallTruckFlag                *string    `bun:"small_truck_flag,type:char(1)"`                                // Indicates whether the Ship To is a small truck
	DeliveryTimeOffset            *int32     `bun:"delivery_time_offset,type:int,default:((0))"`                  // Offset Time to accommodate different time zones.
	TermsOfDeliveryCd             *int32     `bun:"terms_of_delivery_cd,type:int"`                                // Code to identify the terms of delivery of exported goods (CFR, CIF, CIP, CPT, DDP, DAF, DDU, DEQ, DES, EXW, FOB, FAS, FCA, XXX)
	ModeOfTransportCd             *int32     `bun:"mode_of_transport_cd,type:int"`                                // Code to identify the mode used when the good was exported (1, 2,3 ,4, 5, 7, 8, 9)
	UpsRoadnetExcludeFlag         string     `bun:"ups_roadnet_exclude_flag,type:char(1),default:('N')"`          // UPS Roadnet specific: determines if this ship-to is excluded from Roadnet data exports.
	DfoaSoldToId                  *string    `bun:"dfoa_sold_to_id,type:varchar(255)"`                            // Oil eRouter/eServe integration: Shell defined sold to id.
	InvoiceBatchUid               *int32     `bun:"invoice_batch_uid,type:int"`                                   // System generated unique identifier for invoice batches.
	DateAcctOpened                *time.Time `bun:"date_acct_opened,type:datetime"`                               // Date the ship to account became active.
	CourtesyContractShipToId      *float64   `bun:"courtesy_contract_ship_to_id,type:decimal(19,0)"`              // Courtesy contract Ship To ID.
	UpsThirdPartyBillingNo        *string    `bun:"ups_third_party_billing_no,type:varchar(255)"`                 // For UPS ConnectShip integration; specifies the 3rd party freight billing account number sent to the UPS ConnectShip system.
	ReqLotDocWithInvoiceFlag      *string    `bun:"req_lot_doc_with_invoice_flag,type:char(1)"`                   // Require linked lot documentation to print with invoices
	ReqLotDocWithPackinglistFlag  *string    `bun:"req_lot_doc_with_packinglist_flag,type:char(1)"`               // Require linked lot documentation to print with packing lists
	LimitOnlineWarrantiesFlag     *string    `bun:"limit_online_warranties_flag,type:char(1),default:('N')"`      // Limit Online Warranties to Ship To Purchases Only
	RemoteMargin                  *float64   `bun:"remote_margin,type:decimal(19,9)"`                             // column that set an amount to adjust the prices by
	CfnResaleFlag                 string     `bun:"cfn_resale_flag,type:char(1),default:('N')"`                   // column that determine something else about ship_to/customer
	SeparateInvoiceFlag           string     `bun:"separate_invoice_flag,type:char(1),default:('N')"`             // Separate invoice flag
	DoNotAutoInvoiceFlag          string     `bun:"do_not_auto_invoice_flag,type:char(1),default:('N')"`          // Do not auto-invoice flag
	SicCode                       *float64   `bun:"sic_code,type:decimal(6,0)"`                                   // Industry standard four-digit code that identifies a companys industry.
	ValvolineNumber               *string    `bun:"valvoline_number,type:varchar(255)"`                           // Valvoline Account Number
	ConocoShipToId                *string    `bun:"conoco_ship_to_id,type:varchar(255)"`                          // Conoco Ship To
	ConocoSoldToId                *string    `bun:"conoco_sold_to_id,type:varchar(255)"`                          // Conoco Sold To
	SalesMarketGroupUid           *int32     `bun:"sales_market_group_uid,type:int"`                              // Key of sales_market_group table
	OrderByBinPickTicketFlag      *string    `bun:"order_by_bin_pick_ticket_flag,type:char(1)"`                   // Indicates when set bins as the sorting criteria
	AltTaxRateEligibleFlag        *string    `bun:"alt_tax_rate_eligible_flag,type:char(1)"`
	ResidentialAddressFlag        *string    `bun:"residential_address_flag,type:char(1)"`
	ExcludePrintPackinglistInWwms *string    `bun:"exclude_print_packinglist_in_wwms,type:char(1)"`        // Determines whether (unconfirmed) packing lists will print for this ship to when picks for them are completed in WWMS.
	RecordUsageActualLocFlag      *string    `bun:"record_usage_actual_loc_flag,type:char(1)"`             // Record the invoice line usage at the actual location.
	ShipToLinkingId               *int32     `bun:"ship_to_linking_id,type:int"`                           // Custom (F65729): ID used to link multiple ship tos that need to kept in sync for multi-currency functionality where duplicate customers are created to handle transactions in different currencies.
	ServiceLevelAgreementUid      *int32     `bun:"service_level_agreement_uid,type:int"`                  // Custom: Service level agreement associated with ship to record.
	ScanAndPackFlag               string     `bun:"scan_and_pack_flag,type:char(1),default:('N')"`         // Indicated whether the customer/ship_to will have their goods use scan and pack.
	LotAgeRestrictionMonths       *int32     `bun:"lot_age_restriction_months,type:int"`                   // The maximum age of a lot, in terms of months, that this ship to is willing to accept
	FreightForwardAddressId       *float64   `bun:"freight_forward_address_id,type:decimal(19,0)"`         // ID for the address used as this ship tos freight forwarder address.
	AsbDeliveryMethodUid          *int32     `bun:"asb_delivery_method_uid,type:int"`                      // Custom: Default auto short by delivery method for corresponding ship to.
	StopNo                        *int32     `bun:"stop_no,type:int"`                                      // Default Stop Number for Delivery Lists for each Ship To ID
	AdvancedBillingFlag           *string    `bun:"advanced_billing_flag,type:char(1)"`                    // Indicates if this ship to should default to advanced billing in order entry.
	TrackaboutBillByShipTo        *string    `bun:"trackabout_bill_by_ship_to,type:char(1),default:('N')"` // Flag to determine if the ship to is set to be billed separately in TrackAbout
	FreeFreightDays               *int32     `bun:"free_freight_days,type:int"`                            // Custom (F70762): Determines the days selected where freight group charges do not apply.  Numerical representation of the selected days (1-Sunday/2-Monday/4-Tuesday/8-Wednesday/16-Thursday/32-Friday/64-Saturday).
	ProtectedFlag                 *string    `bun:"protected_flag,type:char(1),default:('N')"`             // Indicates the ship to is protected and the system will not check the customer’s credit limit, only ship to
}
