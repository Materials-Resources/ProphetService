package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvMast struct {
	bun.BaseModel                 `bun:"table:inv_mast"`
	InvMastUid                    int32      `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	ItemId                        string     `bun:"item_id,type:varchar(40),unique"`                           // The id used to refer to the item in the application.
	ItemDesc                      string     `bun:"item_desc,type:varchar(40)"`                                // The item description.
	DeleteFlag                    string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	Weight                        *float64   `bun:"weight,type:decimal(19,12)"`                                // How much does the inventory item weigh?
	NetWeight                     *float64   `bun:"net_weight,type:decimal(19,12)"`                            // weight of an item excluding packaging or packing materials
	DateCreated                   time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified              time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy              string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Inactive                      string     `bun:"inactive,type:char(1)"`                                     // Is this item inactive?
	Cube                          *float64   `bun:"cube,type:decimal(19,12)"`                                  // Cubic volume of the item.
	DateInactive                  *time.Time `bun:"date_inactive,type:datetime"`                               // date item was made inactive, or no longer available
	DateReactive                  *time.Time `bun:"date_reactive,type:datetime"`                               // This column is unused.
	CatchWeightIndicator          string     `bun:"catch_weight_indicator,type:char(1)"`                       // This column is unused.
	PurchasingWeight              *float64   `bun:"purchasing_weight,type:decimal(19,12)"`                     // The purchasing weight of the item.
	ClassId1                      *string    `bun:"class_id1,type:varchar(255)"`                               // Item class 1
	ClassId2                      *string    `bun:"class_id2,type:varchar(255)"`                               // Item class 2
	ClassId3                      *string    `bun:"class_id3,type:varchar(255)"`                               // Item class 3
	ClassId4                      *string    `bun:"class_id4,type:varchar(255)"`                               // Item class 4
	ClassId5                      *string    `bun:"class_id5,type:varchar(255)"`                               // Item class 5
	UpcOrEan                      *string    `bun:"upc_or_ean,type:varchar(3)"`                                // This column is unused.
	UpcOrEanId                    *string    `bun:"upc_or_ean_id,type:varchar(15)"`                            // This column is unused.
	Serialized                    string     `bun:"serialized,type:char(1)"`                                   // When Y, indicates that the item is serialized.
	ProductType                   string     `bun:"product_type,type:char(1)"`                                 // Used for classifying the item as regular, lot bill, slab, etc.
	Hose                          string     `bun:"hose,type:char(1)"`                                         // This column is unused.
	Fitting                       string     `bun:"fitting,type:char(1)"`                                      // This column is unused.
	DLength                       *float64   `bun:"d_length,type:decimal(6,4)"`                                // This column is unused.
	CatalogItem                   string     `bun:"catalog_item,type:char(1)"`                                 // This column is unused.
	ShortCode                     *string    `bun:"short_code,type:varchar(30)"`                               // A short code used for retrieiving the item throughout the application.
	TrackLots                     string     `bun:"track_lots,type:char(1)"`                                   // When Y, indicates that the item tracks lots.
	Requisition                   string     `bun:"requisition,type:char(1),default:('N')"`                    // Is the note area an available display area for a requistion note?
	Price1                        *float64   `bun:"price1,type:decimal(19,9)"`                                 // Item price 1
	Price2                        *float64   `bun:"price2,type:decimal(19,9)"`                                 // Item price 2
	Price3                        *float64   `bun:"price3,type:decimal(19,9)"`                                 // Item price 3
	Price4                        *float64   `bun:"price4,type:decimal(19,9)"`                                 // Item price 4
	Price5                        *float64   `bun:"price5,type:decimal(19,9)"`                                 // Item price 5
	Price6                        *float64   `bun:"price6,type:decimal(19,9)"`                                 // Item price 6
	Price7                        *float64   `bun:"price7,type:decimal(19,9)"`                                 // Item price 7
	Price8                        *float64   `bun:"price8,type:decimal(19,9)"`                                 // Item price 8
	Price9                        *float64   `bun:"price9,type:decimal(19,9)"`                                 // Item price 9
	Price10                       *float64   `bun:"price10,type:decimal(19,9)"`                                // Item price 10
	DefaultProductGroup           *string    `bun:"default_product_group,type:varchar(8)"`                     // This column is unused.
	DefaultSalesDiscountGroup     *string    `bun:"default_sales_discount_group,type:varchar(8)"`              // Sales discount group of the item. Used for sales pricing.
	DefaultPurchaseDiscGroup      *string    `bun:"default_purchase_disc_group,type:varchar(8)"`               // Purchase discount group of the item. Used for purchase pricing.
	SalesPricingUnit              *string    `bun:"sales_pricing_unit,type:varchar(8)"`                        // Sales pricing unit of measure.
	SalesPricingUnitSize          *float64   `bun:"sales_pricing_unit_size,type:decimal(19,4)"`                // Size of the sales pricing unit.
	PurchasePricingUnit           *string    `bun:"purchase_pricing_unit,type:varchar(8)"`                     // Purchase pricing unit of measure.
	PurchasePricingUnitSize       *float64   `bun:"purchase_pricing_unit_size,type:decimal(19,4)"`             // Size of the purchase pricing unit.
	ExtendedDesc                  *string    `bun:"extended_desc,type:varchar(255)"`                           // Additional description for the item.
	CommissionClassId             *string    `bun:"commission_class_id,type:varchar(8)"`                       // What is the unique identifier for this item commission class?
	OtherChargeItem               *string    `bun:"other_charge_item,type:char(1)"`                            // To indicate whether a component is an other charge
	InvoiceType                   *string    `bun:"invoice_type,type:varchar(2)"`                              // What type of invoice can this ship to recieve?
	PickTicketType                *string    `bun:"pick_ticket_type,type:varchar(2)"`                          // This column is unused.
	DefaultSellingUnit            *string    `bun:"default_selling_unit,type:varchar(8)"`                      // default sales / order unit of measure
	UpdateViaPricingService       *string    `bun:"update_via_pricing_service,type:char(1)"`                   // Indicates whether the item is updated via pricing service.
	DefaultPurchasingUnit         *string    `bun:"default_purchasing_unit,type:varchar(8)"`                   // Default purchasing unit of measure.
	Hi                            *float64   `bun:"hi,type:decimal(6,0)"`                                      // The number of layers on a pallet - used for tracki
	Ti                            *float64   `bun:"ti,type:decimal(6,0)"`                                      // The number of packages per layer on a pallet - use
	ConvertQuantities             *string    `bun:"convert_quantities,type:char(1)"`                           // To allow items to automatically convert quantities
	HazMatFlag                    string     `bun:"haz_mat_flag,type:char(1),default:('N')"`                   // Is this material hazardous -  per United States shipping regulations?
	ClassCode                     *float64   `bun:"class_code,type:decimal(6,2)"`                              // This column is unused.
	ItemTermsDiscountPct          float64    `bun:"item_terms_discount_pct,type:decimal(19,2),default:(0)"`    // Terms discount amount for the item.
	TpcxStatus                    *int32     `bun:"tpcx_status,type:int"`                                      // When greater than zero indicates that the item is rationalized at the TPCx hub.
	DefaultTransferUnit           *string    `bun:"default_transfer_unit,type:varchar(8)"`                     // Default transfer unit of measure.
	Keywords                      *string    `bun:"keywords,type:text"`                                        // Keywords describing the item.
	FulltextTimestamp             *string    `bun:"fulltext_timestamp,type:timestamp"`                         // Used to determine when the full-text index needs to be updated.
	VendorConsigned               string     `bun:"vendor_consigned,type:char(1),default:('N')"`               // Indicates that the item is Vendor Consigned
	CreatedBy                     *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	Disposition                   *string    `bun:"disposition,type:char(1),default:('N')"`                      // default disposition for item not in stock during Order Entry time
	SourceTypeCd                  *int32     `bun:"source_type_cd,type:int"`                                     // Indicates how this item was created
	UseTagsFlag                   string     `bun:"use_tags_flag,type:char(1),default:('N')"`                    // Wheater this item uses tags or not.
	BaseUnit                      string     `bun:"base_unit,type:varchar(8)"`                                   // smallest unit of measure with unit size 1
	TagHoldClassUid               *int32     `bun:"tag_hold_class_uid,type:int"`                                 // Tag and Hold Type
	ShippableUnitFlag             string     `bun:"shippable_unit_flag,type:char(1),default:('N')"`              // Indicating an item is packaged in a shippable state or not for the Picking and Manifesting process with FASCOR.
	AutoAllocationFlag            *string    `bun:"auto_allocation_flag,type:char(1)"`                           // Indicate if the Item should automatically allocate to open backorder at the time of PO Receipts.
	ParkerProductCd               *string    `bun:"parker_product_cd,type:varchar(255)"`                         // Product Code assigned by Parker Hannifin
	ParkerDivisionCd              *string    `bun:"parker_division_cd,type:varchar(255)"`                        // Division Code assigned by Parker Hannifin
	IvaTaxableFlag                *string    `bun:"iva_taxable_flag,type:char(1)"`                               // IVA taxable item indicator
	RestrictedFlag                *string    `bun:"restricted_flag,type:char(1)"`                                // Custom, indicate the sale of an item is restricted
	SoldOutsideUsFlag             *string    `bun:"sold_outside_us_flag,type:char(1),default:('N')"`             // This will be a Y/N flag that will determine if the item may be sold outside of the US or not
	CurrencyId                    *float64   `bun:"currency_id,type:decimal(19,0)"`                              // Determines the associated item currency
	ServiceTermsDiscountPct       float64    `bun:"service_terms_discount_pct,type:decimal(19,2),default:((0))"` // Service Terms Discount Percent
	ServiceCommissionClassId      *string    `bun:"service_commission_class_id,type:varchar(8)"`                 // Service Commission class
	OverrideSpecificCosting       string     `bun:"override_specific_costing,type:char(1),default:('Y')"`        // Column to indicate if specific costing will be used or not used for 'S' disp items.
	ConfigurableFlag              *string    `bun:"configurable_flag,type:char(1),default:('N')"`                // Indicator field to indicate whether this item is a configurable AQNet Item
	LastPricingServiceDate        *time.Time `bun:"last_pricing_service_date,type:datetime"`                     // Specifies the last time this item was updated by Pricing Service
	CommodityCode                 *string    `bun:"commodity_code,type:varchar(255)"`                            // United Nations Standard Products and Services Code (required for level 3 credit card processing)
	DefaultPriceFamilyUid         *int32     `bun:"default_price_family_uid,type:int"`                           // Unique Identifier for price_family table for Price Family ID as it pertains to this item.
	SpaItemFlag                   *string    `bun:"spa_item_flag,type:char(1),default:('N')"`                    // Indicates this is an SPA item
	AvailForSchDeliveryFlag       *string    `bun:"avail_for_sch_delivery_flag,type:char(1)"`                    // Determines whether or not an item will be available for scheduling on the Degree Days and Calendar Based Tabs of Ship To Maintenance.
	ItemTypeCd                    *int32     `bun:"item_type_cd,type:int,default:((300))"`                       // Custom column indicates what item type is.  Possible values are: Fillable, Parent or None
	CustParentInvMastUid          *int32     `bun:"cust_parent_inv_mast_uid,type:int"`                           // Custom column indicates what parent item id is if the item is fillable.
	UnspscCode                    *string    `bun:"unspsc_code,type:varchar(255)"`                               // United Nations Standard Products and Services Code, global standard classification system
	DciCode                       *string    `bun:"dci_code,type:varchar(255)"`                                  // Standard for specifying items/services used heavily in in electrical and plumbing distribution
	EpaCertReqFlag                *string    `bun:"epa_cert_req_flag,type:char(1),default:('N')"`                // Indicates whether an EPA certification is required in order to sell the item
	AiaEnabledFlag                *string    `bun:"aia_enabled_flag,type:char(1)"`                               // Indicates whether or not this item is enabled for Advanced Inventory Allocation
	AiaRemnantQty                 *float64   `bun:"aia_remnant_qty,type:decimal(19,9)"`                          // An item is considered a remnant when its quantity on hand is less than or equal to this quantity (in SKUs); used for Advanced Inventory Allocation
	ManufacturerProgramTypeUid    *int32     `bun:"manufacturer_program_type_uid,type:int"`                      // link to a manufacturer rebate program type, if populated.
	ManufacturerProgramTypePct    *float64   `bun:"manufacturer_program_type_pct,type:decimal(19,2)"`            // Percentage of the manufacturer program rebate type.
	Length                        *float64   `bun:"length,type:decimal(19,9)"`                                   // Length dimension of this item.
	Width                         *float64   `bun:"width,type:decimal(19,9)"`                                    // Width dimension of this item.
	Height                        *float64   `bun:"height,type:decimal(19,9)"`                                   // Height dimension of this item.
	UseOcTaxRulesFlag             *string    `bun:"use_oc_tax_rules_flag,type:char(1),default:('N')"`            // Determines if an other charge item will use special tax rules.
	TallyFlag                     *string    `bun:"tally_flag,type:char(1)"`                                     // Custom column indicates if this is a Tally Item.
	OcPrintOnPickTicketFlag       *string    `bun:"oc_print_on_pick_ticket_flag,type:char(1)"`                   // For other charge items only, indicates whether this item should print in the items section of pick tickets.
	OcPrintOnInvoiceFlag          *string    `bun:"oc_print_on_invoice_flag,type:char(1)"`                       // For other charge items only, indicates whether this item should print in the items section of invoices.
	UseRevisionsFlag              *string    `bun:"use_revisions_flag,type:char(1),default:('N')"`               // Flag to indicate whether or not revisions will be tracked.
	SingleUseOrReusable           *string    `bun:"single_use_or_reusable,type:char(1)"`                         // Specifies if this is a single use or reusable item.
	UnitconvOverrideOeFlag        *string    `bun:"unitconv_override_oe_flag,type:char(1)"`                      // This column indicates whether the user can override unit conversion rules for this item in Order Entry - values Y, N
	UnitconvOverridePurcFlag      *string    `bun:"unitconv_override_purc_flag,type:char(1)"`                    // This column indicates whether the user can override unit conversion rules for this item in Purchasing - values Y, N
	ItemSalesTaxClass             *string    `bun:"item_sales_tax_class,type:varchar(255)"`                      // Item sales tax classification used by Vertex.
	LifoPoolItemClass             *string    `bun:"lifo_pool_item_class,type:varchar(8)"`                        // LIFO pool item code assigned to an item
	CountryOfOriginReqFlag        *string    `bun:"country_of_origin_req_flag,type:char(1)"`                     // Custom column to indicate if country of origin is required in PO receipts, RMA Receipts, Transfer Receipts, etc.
	NmfcHdrUid                    *int32     `bun:"nmfc_hdr_uid,type:int"`                                       // Unique Identifier for a NMFC code
	CatchLotWeightFlag            *string    `bun:"catch_lot_weight_flag,type:char(1),default:('N')"`            // Flag to determine whether to catch the weight at the lot level
	Ucc128PackTypeCd              *int32     `bun:"ucc128_pack_type_cd,type:int"`                                // Custom: Indicates pack type for ucc128.  Standard(3020) or Pick and Pack (3021)
	Ucc128StandardPackSize        *int32     `bun:"ucc128_standard_pack_size,type:int"`                          // Custom: Number of items per standard pack.
	RedemptionItemFlag            *string    `bun:"redemption_item_flag,type:char(1)"`                           // Rewards Redemtion Item Y/N
	RedemptionValue               *float64   `bun:"redemption_value,type:decimal(19,9)"`                         // Reward pints needed for item
	BoFillCanadianPurchaseFlag    *string    `bun:"bo_fill_canadian_purchase_flag,type:char(1),default:('N')"`   // This field indicates whether this item is a canadian purchase for backorder fulfillment purporse
	CountryOfOriginCode           *string    `bun:"country_of_origin_code,type:char(2)"`                         // Country of origin code, used in item maint.
	ApplyStateFuelSurchargeFlag   *string    `bun:"apply_state_fuel_surcharge_flag,type:char(1),default:('N')"`  // Flag to indicate if the item is a fuel surcharge item
	TypeOfSale                    *string    `bun:"type_of_sale,type:varchar(255)"`                              // Custom - F49889: Type of Sale code used to classify this item for Carrier rebates.
	CommissionCostMultiplier      *float64   `bun:"commission_cost_multiplier,type:decimal(19,9)"`               // column indicates the commission cost multiplier
	ItemLabelTypeCd               *int32     `bun:"item_label_type_cd,type:int"`                                 // indicate an Item as a Private Label item, Standard or Both
	DrumPickupFlag                *string    `bun:"drum_pickup_flag,type:char(1),default:('N')"`                 // Indicate if the item is a Drum Pickup item
	ItemNotes                     *string    `bun:"item_notes,type:varchar(255)"`                                // Contain notes that must be displayed on the custom tab in Order Entry
	RetailAccessories             *string    `bun:"retail_accessories,type:varchar(255)"`                        // Contain basic information on accessories
	HazmatOrmdFlag                string     `bun:"hazmat_ormd_flag,type:char(1),default:('N')"`                 // Indicate that small quantities of an otherwise hazardous material may be shipped without normal hazmat fees
	HazmatOrmdQty                 *int32     `bun:"hazmat_ormd_qty,type:int"`                                    // Contain small quantities
	FunctionalClass               *string    `bun:"functional_class,type:varchar(10)"`                           // Store values that will be determined by business rules
	SalesSetQuantity              *int32     `bun:"sales_set_quantity,type:int"`                                 // Quantity in sold sets
	SuppressFromWebFlag           string     `bun:"suppress_from_web_flag,type:char(1),default:('N')"`           // Determine which items are visible on web site and other “EDI” type customers
	VndrStock                     *int32     `bun:"vndr_stock,type:int"`                                         // User defined custom field indicating available stock qty for VNDR
	PtlxStock                     *int32     `bun:"ptlx_stock,type:int"`                                         // User defined custom field indicating available stock qty for PTLX
	PfStock                       *int32     `bun:"pf_stock,type:int"`                                           // User defined custom field indicating available stock qty for PF
	RoundPrices                   *string    `bun:"round_prices,type:char(1),default:('N')"`                     // Round prices calculated from Item Multipliers
	PumpoffItem                   *string    `bun:"pumpoff_item,type:varchar(40)"`                               // Reference to a pumpoff item code.
	PumpoffUom                    *string    `bun:"pumpoff_uom,type:varchar(8)"`                                 // The UOM used by a pumpoff item.
	ConocoDivisionClassId         *string    `bun:"conoco_division_class_id,type:varchar(8)"`                    // Used by Conoco integration to identify their items. Will be reported back on buyback exports.
	BuyGetClass                   *string    `bun:"buy_get_class,type:varchar(8)"`                               // The buy get class to group buy get items.
	DiscontinuedDate              *time.Time `bun:"discontinued_date,type:datetime"`                             // Custom column to indicate the item is discontinued
	LabelGroupItemFlag            *string    `bun:"label_group_item_flag,type:char(1)"`                          // (Custom) Indicates that this item is a special item used to track a list (group) of items that a customer has requested labels for in OE
	CylinderItemFlag              *string    `bun:"cylinder_item_flag,type:char(1)"`                             // Column flag to identify the item as a cylinder
	CoreItemFlag                  *string    `bun:"core_item_flag,type:char(1)"`                                 // Indicates if item is considered core
	GenericItemDesc               *string    `bun:"generic_item_desc,type:varchar(40)"`                          // Generic Item Description
	AllowCustomDescFlag           *string    `bun:"allow_custom_desc_flag,type:char(1),default:('N')"`           // Allow Custom Description Flag
	ManufacturerName              *string    `bun:"manufacturer_name,type:varchar(255)"`                         // Manufacturer name for this item.
	BrandName                     *string    `bun:"brand_name,type:varchar(255)"`                                // Brand for this item.
	PartNumber                    *string    `bun:"part_number,type:varchar(255)"`                               // Part Number for this item.
	DoNotAutoAllocateDesc         *string    `bun:"do_not_auto_allocate_desc,type:varchar(255)"`                 // (Custom) Free form description of why an item should not be automatically allocated during order entry or po receipts.  Null value indicates that normal baseline allocation should occur.
	ExcludeOrderLevelDiscountFlag *string    `bun:"exclude_order_level_discount_flag,type:char(1)"`              // Used as a flag to determine when to exclude item from order level discount at OE
	SampleItemFlag                *string    `bun:"sample_item_flag,type:char(1),default:('N')"`                 // Indicates if the Item is a sample item
	SampleInvMastUid              *int32     `bun:"sample_inv_mast_uid,type:int"`                                // Indicates the sample item uid for this item
	CoreParentItemId              *string    `bun:"core_parent_item_id,type:varchar(40)"`
	AttributeGroupUid             *int32     `bun:"attribute_group_uid,type:int"` // Attribute group that applies to this item.
	PartnerProgramLaborItemFlag   *string    `bun:"partner_program_labor_item_flag,type:char(1)"`
	BrandType                     *int32     `bun:"brand_type,type:int"`                                      // Brand Type
	IqsItemFlag                   *string    `bun:"iqs_item_flag,type:char(1)"`                               // Indicates that this item is used with the IQS quality and compliance 3rd party application.
	ImportAsDirectShipFlag        *string    `bun:"import_as_direct_ship_flag,type:char(1),default:('N')"`    // Flag to determine if the item will import as a direct ship for Order Imports
	DoNotFieldDestroyFlag         *string    `bun:"do_not_field_destroy_flag,type:char(1)"`                   // Custom: Indicates that an item cannot be flagged as a field destroy on a B2B RMA Request.
	InvMastUidDupUsage            *int32     `bun:"inv_mast_uid_dup_usage,type:int"`                          // Link to inv_mast record for duplicate item usage.
	RolledItemFlag                *string    `bun:"rolled_item_flag,type:char(1)"`                            // Custom: Indicates that item is a rolled item (i.e. carpet)
	ExcludeFromEdi832Flag         *string    `bun:"exclude_from_edi832_flag,type:char(1)"`                    // Flag to determine if an item should be included in the EDI Export 832
	OrderHoldClassId              *string    `bun:"order_hold_class_id,type:varchar(255)"`                    // Column to hold class ids separated by commas.
	DefaultDirectShipDispFlag     *string    `bun:"default_direct_ship_disp_flag,type:char(1),default:('N')"` // Used to indicate if the order line disposition should automatically default to N in various areas of the application
	BuyerId                       *string    `bun:"buyer_id,type:varchar(16)"`                                // Custom: Buyer ID associated with purchasing this item.
	TariffSurchargePct            *float64   `bun:"tariff_surcharge_pct,type:decimal(19,9)"`                  // The percentage of price and cost that will be applied as a surcharge to cover tariff costs when this item is shipped.
	TariffSurchargeInvMastUid     *int32     `bun:"tariff_surcharge_inv_mast_uid,type:int"`                   // The other charge item that will be used to apply the surcharge to a shipment when this item is shipped.
	PackingWeightTrackingOption   *string    `bun:"packing_weight_tracking_option,type:char(1)"`              // Determine if item level system will use the gross weight (weight of item plus packaging) or the Net weight (just the item). Valid value are G or N
	ExcludeFromAutoShortBuyFlag   *string    `bun:"exclude_from_auto_short_buy_flag,type:char(1)"`            // Custom: Indicates corresponding item should be excluded from the auto short buy process.
	ItemClassificationId          *string    `bun:"item_classification_id,type:varchar(255)"`                 // Custom: Item classification id associated with this item.
	HazmatFeeInvMastUid           *int32     `bun:"hazmat_fee_inv_mast_uid,type:int"`                         // Links the other charge item to be used as a hazmat fee for *this* item.
	BolClass                      *string    `bun:"bol_class,type:varchar(8)"`                                // BOL Class for item.
	RentalItemFlag                *string    `bun:"rental_item_flag,type:char(1),default:('N')"`              // Indicates this is a rentable item
	RentalItemUid                 *int32     `bun:"rental_item_uid,type:int"`                                 // UID assigned by rental
	DfltDimensionScale            *string    `bun:"dflt_dimension_scale,type:varchar(2)"`                     // Track default dimension scale for atomic lot item
	EccEnabledFlag                string     `bun:"ecc_enabled_flag,type:char(1),default:('N')"`              // Enable Item for ECC
	LocalItemFlag                 *string    `bun:"local_item_flag,type:char(1),default:('N')"`               // Local Item Indicator
}
