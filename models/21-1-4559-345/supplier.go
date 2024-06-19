package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Supplier struct {
	bun.BaseModel                   `bun:"table:supplier"`
	SupplierId                      float64   `bun:"supplier_id,type:decimal(19,0),pk"`                         // Unique Identifier for this supplier.
	BuyerId                         string    `bun:"buyer_id,type:varchar(16),nullzero"`                        // What buyer is associated with this supplier?
	DaysEarly                       float64   `bun:"days_early,type:decimal(19,0),default:(0)"`                 // Indicates the maximum number of days (prior to the required date) that you will accept material from this vendor.
	DaysLate                        float64   `bun:"days_late,type:decimal(19,0),default:(0)"`                  // Indicates the maximum number of days (past the required dates) that you will accept material from this vendor
	DeleteFlag                      string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                     time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified                time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy                string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ControlValue                    string    `bun:"control_value,type:varchar(16),nullzero"`                   // Weight/Volume/Dollars
	TargetValue                     float64   `bun:"target_value,type:decimal(19,9),nullzero"`                  // Target amount (in Weight, Volume, Dollars, or Units) to purchase on a PO.
	ReviewCycle                     float64   `bun:"review_cycle,type:decimal(3,0),nullzero"`                   //  How often (in days) a supplier’s line will be reviewed by a buyer.
	DateOfLastReview                time.Time `bun:"date_of_last_review,type:datetime,nullzero"`                // When was the last time GPOR was ran for this supplier, defining requirements?
	DefaultCarrier                  float64   `bun:"default_carrier,type:decimal(19,0),nullzero"`               // Carrier most often used for shipments from this supplier
	DefaultShippingInstructions     string    `bun:"default_shipping_instructions,type:varchar(255),nullzero"`  // Shipping instructions for this supplier.
	AverageLeadTime                 float64   `bun:"average_lead_time,type:decimal(3,0),nullzero"`              // What is the average lead time for this supplier
	LeadTimeSource                  string    `bun:"lead_time_source,type:varchar(8),nullzero"`                 // Calculate lead time at the item or supplier level?
	LeadTimeSafetyFactor            float64   `bun:"lead_time_safety_factor,type:decimal(19,4),nullzero"`       // Added to lead time to insure timely delivery.
	SafetyStockDays                 float64   `bun:"safety_stock_days,type:decimal(19,4),nullzero"`             // The number of extra days supply of material that should be maintained for the items purchased during this session.
	Fct                             string    `bun:"fct,type:varchar(10),nullzero"`                             // No longer used
	CurrencyId                      float64   `bun:"currency_id,type:decimal(19,0),nullzero"`                   // What is the unique currency identifier for this supplier?
	SupplierName                    string    `bun:"supplier_name,type:varchar(255),nullzero"`                  // What is the name of this supplier?
	TransitDays                     int32     `bun:"transit_days,type:int,default:(1)"`                         // Used for quality tracking purposes. Provides Transit Days
	RmaRequired                     string    `bun:"rma_required,type:char(1),default:('N')"`                   // This field is available for you to select if the supplier requires a Return Material Authorization (RMA) number when you are returning goods to them
	CreatePoFromOe                  int32     `bun:"create_po_from_oe,type:int"`                                // 984 - Create directships, 985 - create all, 986 create none.
	TpcxRationalized                string    `bun:"tpcx_rationalized,type:char(1),default:('N')"`              // The column is vendor rationalization flag so CommerceCenter knows if there is an item being rationalized.  If a supplier has at least one item being rationalized within TPCx, the flag is set to Y.
	FreightControlValue             string    `bun:"freight_control_value,type:varchar(16),nullzero"`           // This supplements the standard control value and allows the distributor to maintain separate thresholds for things like minimum order amonts and free freight.
	FreightTargetValue              float64   `bun:"freight_target_value,type:decimal(19,9),default:(0)"`       // This supplements the standard target  value and allows the distributor to maintain separate thresholds for things like minimum order amonts and free freight.
	CreatedBy                       string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CreateVendorRfqCd               int32     `bun:"create_vendor_rfq_cd,type:int,default:(986)"`                    // Indicates where the user is allowed to create vendor rfq.
	VendorRfqDays                   int32     `bun:"vendor_rfq_days,type:int,default:(30)"`                          // The number of days before the RFQ expires.
	ShipDays                        int32     `bun:"ship_days,type:int,nullzero"`                                    // Number of days until receipt from the supplier ship date
	DeaCode                         string    `bun:"dea_code,type:char(1),nullzero"`                                 // A special code to identify DEA
	DeaNumber                       string    `bun:"dea_number,type:char(255),nullzero"`                             // Indicates DEA Number
	DddReport                       string    `bun:"ddd_report,type:char(1),nullzero"`                               // YesNo value field to indicate DEA should appear on report
	LegacyId                        string    `bun:"legacy_id,type:varchar(255),nullzero"`                           // column to hold supplier ID from legacy app
	ParkerDivisionCd                string    `bun:"parker_division_cd,type:varchar(10),nullzero"`                   // Item level custom column.
	ParkerSupplierFlag              string    `bun:"parker_supplier_flag,type:char(1),nullzero"`                     // Y/N flag to determine if this is a supplier of Parker Hannifin goods.
	CombineStockNsSpecialFlag       string    `bun:"combine_stock_ns_special_flag,type:char(1),default:('N')"`       // Combine Stock, Non-stock, and Specials flag
	SafetyStockType                 int32     `bun:"safety_stock_type,type:int,nullzero"`                            // Method used to determine safety stock
	ServiceLevelMeasure             int32     `bun:"service_level_measure,type:int,nullzero"`                        // Customer service level (stock out back order)
	ServiceLevelPctGoal             float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`             // Percent of service level goal for safety stock
	UseContainersFlag               string    `bun:"use_containers_flag,type:char(1),default:('N')"`                 // Indicates if this supplier uses container receipts
	ContainerVolume                 float64   `bun:"container_volume,type:decimal(19,9),default:((0))"`              // Container volume capacity default value
	ContainerWeight                 float64   `bun:"container_weight,type:decimal(19,9),default:((0))"`              // Container weight capacity default
	PoApprovalThresholdFlag         string    `bun:"po_approval_threshold_flag,type:char(1),nullzero"`               // Flag indicating whether or not the PO approval threshold should be ignored for the current user on POs for this supplier
	SupplierSystemCd                int32     `bun:"supplier_system_cd,type:int,nullzero"`                           // Custom: System used to communicate with this supplier
	PrintPoFlag                     string    `bun:"print_po_flag,type:char(1),nullzero"`                            // Indicates whether POs to this supplier will be printed by default. Supplier can have multiple default methods (print, fax, email, edi).
	FaxPoFlag                       string    `bun:"fax_po_flag,type:char(1),nullzero"`                              // Indicates whether POs to this supplier will be faxed by default. Supplier can have multiple default methods (print, fax, email, edi).
	EmailPoFlag                     string    `bun:"email_po_flag,type:char(1),nullzero"`                            // Indicates whether POs to this supplier will be emailed by default. Supplier can have multiple default methods (print, fax, email, edi).
	EdiPoFlag                       string    `bun:"edi_po_flag,type:char(1),nullzero"`                              // Indicates whether POs to this supplier will be sent via EDI by default. Supplier can have multiple default methods (print, fax, email, edi).
	RestrictPedigreeItem            string    `bun:"restrict_pedigree_item,type:char(1),default:('Y')"`              // Indicates if pedigree items are quarantined on receipt
	PallSupplierNumber              string    `bun:"pall_supplier_number,type:varchar(20),nullzero"`                 // Holds the vendor Pall's supplier number. This is used for reporting information to the vendor Pall.
	RestrictOtfItemFlag             string    `bun:"restrict_otf_item_flag,type:char(1),default:('N')"`              // (Custom) Indicates whether or not an on-the-fly item can be created for this supplier.
	SerialTrackingOption            string    `bun:"serial_tracking_option,type:char(1),default:('D')"`              // Custom (F47033): determines the serial number tracking option for this supplier.  Valid values are 'D' (Default to the system setting), 'I' (track inbound and outbound) and 'O' (track outbound only).
	TermsDiscountExcludedFlag       string    `bun:"terms_discount_excluded_flag,type:char(1),default:('N')"`        // Flag indication of whether supplier is eligible for customer defaulted terms discount.
	WeeklyTruckAllowedFlag          string    `bun:"weekly_truck_allowed_flag,type:char(1),nullzero"`                // Indicates that this supplier can make deliveries using weekly trucks
	RatelinxLocationId              int32     `bun:"ratelinx_location_id,type:int,nullzero"`                         // Holds the associated location ID from the RateLinx system..
	DefaultPriceExpirationDate      time.Time `bun:"default_price_expiration_date,type:datetime,nullzero"`           // Default price expiration date for all items for this supplier
	FreightFactor                   float64   `bun:"freight_factor,type:decimal(19,9),nullzero"`                     // Percentage of item price that will be used to default incoming freight for this supplier.
	MinFreightFactor                float64   `bun:"min_freight_factor,type:decimal(19,9),nullzero"`                 // Minimum incoming freight that should be defaulted (based on freight_factor) for this supplier.
	DealerWarrantyClaimsCd          int32     `bun:"dealer_warranty_claims_cd,type:int,nullzero"`                    // Determines if this supplier accepts dealer warranty claims.
	ParkerHannafinSupplierFlag      string    `bun:"parker_hannafin_supplier_flag,type:char(1),default:('N')"`       // Indicate that the supplier is Parker Hannafin
	ParkerAccountNo                 string    `bun:"parker_account_no,type:varchar(255),nullzero"`                   // Parker account number.
	BrandNameCd                     int16     `bun:"brand_name_cd,type:smallint,nullzero"`                           // Brand name code for the supplier, possible three values are 1399, 3059, 3060
	RcdSupplierFlag                 string    `bun:"rcd_supplier_flag,type:char(1),nullzero"`                        // Indicated whether this supplier is a RCD supplier
	SurchargeFlag                   string    `bun:"surcharge_flag,type:char(1),default:('N')"`                      // Indicates if Surcharge tab is enabled or not
	PoDefaultMethod                 int32     `bun:"po_default_method,type:int,nullzero"`                            // Purchase Order Default Delivery Method
	UseQtyRdyForContBuild           string    `bun:"use_qty_rdy_for_cont_build,type:char(1),default:('N')"`          // Determines if we use the ready quantity or open quantity during container building
	FidelitoneSupplierId            string    `bun:"fidelitone_supplier_id,type:varchar(255),nullzero"`              // Custom: Code used by Fidelitone to identify this supplier
	DefaultServicebenchSupplierFlag string    `bun:"default_servicebench_supplier_flag,type:char(1),default:('N')"`  // Checkbox called Default ServiceBench Supplier in Supplier Maintenance
	DefaultPoPackingBasis           string    `bun:"default_po_packing_basis,type:varchar(255),nullzero"`            // Default packing basis to be used on purchase orders for this supplier.
	DfltSupplierRedemptionMethod    int32     `bun:"dflt_supplier_redemption_method,type:int,nullzero"`              // Code representing the default supplier redemption method.
	SupplierRedemptionEmail         string    `bun:"supplier_redemption_email,type:varchar(255),nullzero"`           // Email address for supplier redemption.
	SupplierRedemptionFaxNumber     string    `bun:"supplier_redemption_fax_number,type:varchar(20),nullzero"`       // Fax number for supplier redemption.
	ExportOrderInfoFlag             string    `bun:"export_order_info_flag,type:char(1),nullzero"`                   // Y indicates customer order info is to be exported for POs
	MinimumProcessPoCost            float64   `bun:"minimum_process_po_cost,type:decimal(19,9),nullzero"`            // Minimum Process PO Cost
	BulkDiscountPercentage          float64   `bun:"bulk_discount_percentage,type:decimal(19,9),nullzero"`           // Custom: Discount percentage offered by supplier when qualified items are purchased in bulk.
	BulkContractNumberList          string    `bun:"bulk_contract_number_list,type:varchar(255),nullzero"`           // Custom: Comma delimited list of supplier contract numbers that qualify for a bulk discount.
	DoNotAutoCreateReturnsFlag      string    `bun:"do_not_auto_create_returns_flag,type:char(1),default:('N')"`     // If checked, will prevent the system from automatically creating inventory returns for this supplier
	CarrierShipMethodUid            int32     `bun:"carrier_ship_method_uid,type:int,nullzero"`                      // Carrier Ship Method Unique Identifier tied to the Carrier Ship Method UId from the carrier_ship_method table
	DoNotCreateUnapprovedPoFlag     string    `bun:"do_not_create_unapproved_po_flag,type:char(1),nullzero"`         // Custom column to indicate not to create unapproved po for 'RO and ERO' sales order priority backorders from oe.
	DeviationLookbackPds            int32     `bun:"deviation_lookback_pds,type:int,nullzero"`                       // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd              float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`              // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd              float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`              // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd            float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`            // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays              float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`              // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays              float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`              // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult           float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`           // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays          float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`         // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays          float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`         // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	ExclSendLinkedInfoToRfnavFlag   string    `bun:"excl_send_linked_info_to_rfnav_flag,type:char(1),default:('N')"` // Determines whether or not send linked po info to RF Navigator.
}
