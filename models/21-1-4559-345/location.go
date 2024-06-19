package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Location struct {
	bun.BaseModel                  `bun:"table:location"`
	CompanyId                      string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	LocationId                     float64   `bun:"location_id,type:decimal(19,0),pk"`                         // What is the unique location identifier for this row
	DefaultBranchId                string    `bun:"default_branch_id,type:varchar(8),nullzero"`                // Default branch of the company for the location.
	DeleteFlag                     string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                    time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ShippingGroupId                string    `bun:"shipping_group_id,type:varchar(8),nullzero"`                // This column is unused.
	TaxGroupId                     string    `bun:"tax_group_id,type:varchar(10),nullzero"`                    // Associates a location to a tax group.
	LocationName                   string    `bun:"location_name,type:varchar(255),nullzero"`                  // Indicates the location name corresponding to the location
	LotBinIntegration              string    `bun:"lot_bin_integration,type:char(1),nullzero"`                 // Does this location use lot/bin integration?
	LocationType                   int32     `bun:"location_type,type:int"`                                    // Location Type
	NoOfPeriodsToSupply            int16     `bun:"no_of_periods_to_supply,type:tinyint,nullzero"`             // this is the number of periods to review when periods supply is chosen as the available to transfer option
	AvailableToTransferCalcCd      int32     `bun:"available_to_transfer_calc_cd,type:int,nullzero"`           // Integer value that represents the algorithm used to calculate the available to transfer quantity from GPOR
	UseTagsFlag                    string    `bun:"use_tags_flag,type:char(1),default:('N')"`                  // Flags if using tags at location or not.
	TagPrefix                      string    `bun:"tag_prefix,type:varchar(255),nullzero"`                     // Tag prefix for location
	RfEnabledFlag                  string    `bun:"rf_enabled_flag,type:char(1),default:('N')"`                // RF management enabled at this location
	FedexLocAcctNo                 string    `bun:"fedex_loc_acct_no,type:varchar(255),nullzero"`              // Holds a FedEx account number specific to the location used to provide integration FedEx.
	FedexMeterNo                   string    `bun:"fedex_meter_no,type:varchar(255),nullzero"`                 // The meter number that fedex assigns to a location based on the account number
	DeaCode                        string    `bun:"dea_code,type:char(1),nullzero"`                            // Char field to indicate DEA code.
	DeaNumber                      string    `bun:"dea_number,type:varchar(255),nullzero"`                     // This will indicate the DEA Number associated with the distributor’s particular location.
	WwmsShowBinAllocWarning        string    `bun:"wwms_show_bin_alloc_warning,type:char(1),default:('Y')"`    // Tells system whether to show warning in CC when WWMS bin allocations are open
	ParkerDistributorNo            string    `bun:"parker_distributor_no,type:varchar(255),nullzero"`          // The number assigned by Parker Hannafin to this locaiton.
	RfFrontcounterpickingFlag      string    `bun:"rf_frontcounterpicking_flag,type:char(1),nullzero"`         // Indicate if the location supports Wireless Front Counter Picking
	DefaultOeSourceLocId           float64   `bun:"default_oe_source_loc_id,type:decimal(19,0),nullzero"`      // Default source location id for OrderEntry which used to override this location
	RegionUid                      int32     `bun:"region_uid,type:int,nullzero"`                              // Ties the location to an entry in the region table
	ConsignedShipToId              float64   `bun:"consigned_ship_to_id,type:decimal(19,0),nullzero"`          // The ship to location associated with the consigned inventory location
	PoReceiptsBinUid               int32     `bun:"po_receipts_bin_uid,type:int,nullzero"`                     // Unique identifier of bin that will be defaulted for PO receipts at this location.
	TransferReceiptsBinUid         int32     `bun:"transfer_receipts_bin_uid,type:int,nullzero"`               // Unique identifier of bin that will be defaulted for transfer receipts at this location.
	AutoconfirmProdPickTicket      string    `bun:"autoconfirm_prod_pick_ticket,type:char(1)"`                 // whether or not the system should try to automatically confirm production pick tickets at this location after creating them
	PrintProdPickTicket            string    `bun:"print_prod_pick_ticket,type:char(1),default:('N')"`         // whether or not the system should print production pick tickets at this location after creating them.
	UpsAccountNo                   string    `bun:"ups_account_no,type:varchar(255),nullzero"`                 // UPS OnLine® shipper account number
	UpsPickupTypeCd                int32     `bun:"ups_pickup_type_cd,type:int,nullzero"`                      // UPS OnLine® pickup type code
	UpsCustomerTypeCd              int32     `bun:"ups_customer_type_cd,type:int,nullzero"`                    // UPS OnLine® customer type code (Customer Classification)
	GeneralManagerSalesrepId       string    `bun:"general_manager_salesrep_id,type:varchar(16),nullzero"`     // Salesrep ID which is typically the general manager at this location who earns commission on all orders from this sales location.
	InvGroupHdrUid                 int32     `bun:"inv_group_hdr_uid,type:int,nullzero"`                       // Inventory mgmt group assigned to this location.
	DefaultFaxCoverUid             int32     `bun:"default_fax_cover_uid,type:int,nullzero"`                   // UID of the fax cover to be used for this location
	AssociatedLocationId           float64   `bun:"associated_location_id,type:decimal(19,0),nullzero"`        // Associated Inventory Location
	UpsOltAccessKey                string    `bun:"ups_olt_access_key,type:varchar(255),nullzero"`             // UPS access key for the OnLine® Tools. (Encrypted)
	UpsOltPassword                 string    `bun:"ups_olt_password,type:varchar(255),nullzero"`               // UPS password to access the UPS OnLine® Tools. (Encrypted)
	UpsOltUserId                   string    `bun:"ups_olt_user_id,type:varchar(255),nullzero"`                // UPS user id to access the UPS OnLine® Tools. (Encrypted)
	AiaPromptFlag                  string    `bun:"aia_prompt_flag,type:char(1),nullzero"`                     // Indicates whether or not this location will show a prompt to accept the results of Advanced Inventory Allocation
	CardlockAcctNo                 string    `bun:"cardlock_acct_no,type:varchar(255),nullzero"`               // Account number in Cardlock system used to identify location
	PedigreeContactId              string    `bun:"pedigree_contact_id,type:varchar(16),nullzero"`             // Indicates contact id designated as a pedigree contact for this location
	StrategicRetailLocationFlag    string    `bun:"strategic_retail_location_flag,type:char(1),default:('N')"` // When set, location is considered a retail location for strategic pricing.  When not set, it is considered a warehouse location
	WwmsDefaultPrimaryBinPo        string    `bun:"wwms_default_primary_bin_po,type:char(1),default:('N')"`    // Defaul Primary Bin in PO Receipts
	WwmsDefaultPrimaryBinTrans     string    `bun:"wwms_default_primary_bin_trans,type:char(1),default:('N')"` // Default Primary Bin in Transfer Receipts
	ManufacturerRebateLoc          string    `bun:"manufacturer_rebate_loc,type:varchar(255),nullzero"`        // Manufacturer assigned location number for processing rebates.
	TrackCustomerPackageFlag       string    `bun:"track_customer_package_flag,type:char(1),default:('N')"`    // Flag to indicate whether to track customer packages.
	WeightPerBox                   float64   `bun:"weight_per_box,type:decimal(19,9),default:((0))"`           // Weight amount per box.  Used to estimate total number of boxes need for shipment.
	RoadnetDoNotRouteFlag          string    `bun:"roadnet_do_not_route_flag,type:char(1),nullzero"`           // Determines if UPS Roadnet has been disabled for this location.
	DistributionCenter             string    `bun:"distribution_center,type:char(1),default:('N')"`
	NonstockAvailableToTransfer    int32     `bun:"nonstock_available_to_transfer,type:int,nullzero"`            // Integer value that represents the algorithm used to calculate the available to transfer quantity from GPOR for non-stock items.
	NonstockNoPeriodsToSupply      int32     `bun:"nonstock_no_periods_to_supply,type:int,nullzero"`             // Number of periods to supply for nonstock items.
	ServerTimeZoneOffset           int32     `bun:"server_time_zone_offset,type:int,nullzero"`                   // Indicates a positive or negative offset number of hours for a location from the timezone of a company wide central server.
	SalesTaxPayableGlAccount       string    `bun:"sales_tax_payable_gl_account,type:varchar(32),nullzero"`      // Sales Tax Payable Account No
	ExternalTaxBackupPercent       float64   `bun:"external_tax_backup_percent,type:decimal(19,9),nullzero"`     // The tax percentage to be calculated and charged when a 3rd party tax system is down.
	RegionFlag                     string    `bun:"region_flag,type:char(1),nullzero"`                           // Custom column to indicate which region this location record belongs to.
	UseBinsForTransScheduleFlag    string    `bun:"use_bins_for_trans_schedule_flag,type:char(1),default:('N')"` // Flag to indicate if location uses transfer scheduling.
	UsePalletsFlag                 string    `bun:"use_pallets_flag,type:char(1),default:('N')"`                 // Flag that indicates that this location will use pallets for transfers.
	UsePalletTypesFlag             string    `bun:"use_pallet_types_flag,type:char(1),default:('N')"`            // Flag that indicates that this location will use pallet types for transfers.
	VoucherClass                   string    `bun:"voucher_class,type:varchar(8),nullzero"`                      // Voucher class to be assigned to vouchers created as a result of sales or use of vendor consigned inventory.
	AllowMultipleAssembliesFlag    string    `bun:"allow_multiple_assemblies_flag,type:char(1),nullzero"`        // Indicates whether the location allows multiple assemblies per production order.
	UndershipmentAllocationFlag    string    `bun:"undershipment_allocation_flag,type:char(1),default:('N')"`    // Indicates whether the material remains allocated when undershipped for will call orders
	RatelinxLocationId             int32     `bun:"ratelinx_location_id,type:int,nullzero"`                      // Mapped location id in RateLinx
	CenteronDefaultTaker           string    `bun:"centeron_default_taker,type:varchar(30),nullzero"`            // Default Taker for a Centeron Order for this location
	StartPickSeqTagBulk            int32     `bun:"start_pick_seq_tag_bulk,type:int,nullzero"`                   // Allocation by tags will not use sequence < this for bulk allocation step
	HighVelocityLevel              float64   `bun:"high_velocity_level,type:decimal(19,9),nullzero"`             // Location specific value for high velocity level
	MidVelocityLevel               float64   `bun:"mid_velocity_level,type:decimal(19,9),nullzero"`              // Location specific value for mid velocity level
	TaxBySourceLocFlag             string    `bun:"tax_by_source_loc_flag,type:char(1),nullzero"`                // Indicates whether oe line with this location as source uses the tax group of the source location instead of ship to location
	ShippingLocationFlag           string    `bun:"shipping_location_flag,type:char(1),nullzero"`                // Custom: Indicates location can be used for shipping
	GstRegistrationNo              string    `bun:"gst_registration_no,type:varchar(20),nullzero"`               // GST Registration Number
	WwmsDefaultDepositBin          string    `bun:"wwms_default_deposit_bin,type:varchar(10),nullzero"`          // Default deposit bin for all WWMS deposits in this location if the flag to use it is set.
	QuotebuilderDefaultTaker       string    `bun:"quotebuilder_default_taker,type:varchar(30),nullzero"`        // Default Taker for a Quotebuilder Order for this location
	ExcludeOtfItemsFlag            string    `bun:"exclude_otf_items_flag,type:char(1),default:('N')"`           // Exclude for OTF Items
	ExtTaxCompanyCode              string    `bun:"ext_tax_company_code,type:varchar(255),nullzero"`             // (Custom) Used in conjunction with 3rd party taxing. Indicates what company code should be sent to the 3rd party system when taxing for a given sales location.
	FcaBinUid                      int32     `bun:"fca_bin_uid,type:int,nullzero"`                               // Front Counter Bin UID
	FcaReasonId                    string    `bun:"fca_reason_id,type:varchar(5),nullzero"`                      // Front Counter Adjustment Reason
	PreventAutoAssignLotsFlag      string    `bun:"prevent_auto_assign_lots_flag,type:char(1),default:('N')"`    // Indicates if automatic allocation of lots for allocated quantity should be overridden
	FedexLabelSizeCd               int32     `bun:"fedex_label_size_cd,type:int,nullzero"`                       // Code for the Fedex label size
	FedexLabelOrientationCd        int32     `bun:"fedex_label_orientation_cd,type:int,nullzero"`                // Code for the Fedex label orientation
	FedexLabelFormatCd             int32     `bun:"fedex_label_format_cd,type:int,nullzero"`                     // Fedex label format type
	PreventAutoTsCreationFlag      string    `bun:"prevent_auto_ts_creation_flag,type:char(1),default:('N')"`    // This flag maks the location from where no transfer should be created automatically after allocating items related to a pending BO,
	TrReceiptsUsePrimaryBinCd      int32     `bun:"tr_receipts_use_primary_bin_cd,type:int,default:((3406))"`    // Custom: Determine if primary bin should be used as default when receiving transfers in this location.
	DamagedBinUid                  int32     `bun:"damaged_bin_uid,type:int,nullzero"`
	NostockBinUid                  int32     `bun:"nostock_bin_uid,type:int,nullzero"`
	ConsignmentCustomerId          float64   `bun:"consignment_customer_id,type:decimal(19,0),nullzero"`
	RollupLocationId               float64   `bun:"rollup_location_id,type:decimal(19,0),nullzero"`
	PrintPackinglistInWwms         string    `bun:"print_packinglist_in_wwms,type:char(1),nullzero"`                                                                                                                                                               // Indicates whether picks processed in WWMS for this location should print packing lists
	PrintTransferPackinglistInWwms string    `bun:"print_transfer_packinglist_in_wwms,type:char(1),nullzero"`                                                                                                                                                      // Indicates whether transfer picks processed in WWMS for this location should print packing lists
	PickProblemBinUid              int32     `bun:"pick_problem_bin_uid,type:int,nullzero"`                                                                                                                                                                        // Unique identifier of bin that will be used to ensure items not found in the bin allocated can be handled without on the fly inventory adjustments.
	RfTransferPackListPrinter      string    `bun:"rf_transfer_pack_list_printer,type:varchar(255),nullzero"`                                                                                                                                                      // Printer to be used in WWMS for transfer packing lists.
	PrintUcc128WithPtFlag          string    `bun:"print_ucc128_with_pt_flag,type:char(1),default:('N')"`                                                                                                                                                          // Indicates whether to print UCC-128 labels with PTs.
	AutoAllocateInPoImportFlag     string    `bun:"auto_allocate_in_po_import_flag,type:char(1),nullzero"`                                                                                                                                                         // Custom: Indicates whether this records location should auto allocate material when received thru the custom po receipts import.
	DefaultFieldDestroyBinUid      int32     `bun:"default_field_destroy_bin_uid,type:int,nullzero"`                                                                                                                                                               // Unique id of putable bin set as default when using RMA field destroy functionality
	DefaultFieldDestroyReasonId    string    `bun:"default_field_destroy_reason_id,type:varchar(255),nullzero"`                                                                                                                                                    // Inventory adjustment reason code set as default when using RMA field destroy functionality
	WarrantyReturnInvMastUid       int32     `bun:"warranty_return_inv_mast_uid,type:int,nullzero"`                                                                                                                                                                // used to indentify warranty return assembly
	WarrantyReturnBinUid           int32     `bun:"warranty_return_bin_uid,type:int,nullzero"`                                                                                                                                                                     // used to indentify warranty return bin
	UseDqRoutingFlag               string    `bun:"use_dq_routing_flag,type:char(1),nullzero"`                                                                                                                                                                     // Flag if location uses DQ routing.
	UseCatchWeightProcessingFlag   string    `bun:"use_catch_weight_processing_flag,type:char(1),nullzero"`                                                                                                                                                        // Flag for allowing/restricting catch weight processing at location level
	RmaReceiptsUsePrimaryBin       string    `bun:"rma_receipts_use_primary_bin,type:char(1),nullzero"`                                                                                                                                                            // This value will be used instead of rma_receipts_use_primary_bin system setting
	FedexDefaultPhoneNumber        string    `bun:"fedex_default_phone_number,type:varchar(20),nullzero"`                                                                                                                                                          // Default phone used when there isnt one specified on the Ship To in OE
	SkipParkerExportInOe           string    `bun:"skip_parker_export_in_oe,type:char(1),default:('N')"`                                                                                                                                                           // Flag to skip sending Parker PTS info in OE
	AdjustFoundItemsFlag           string    `bun:"adjust_found_items_flag,type:char(1),default:('N')"`                                                                                                                                                            // Determines if the functionality to adjust found items via the Physical Count window is enabled
	RmaDefaultWarrantyBinUid       int32     `bun:"rma_default_warranty_bin_uid,type:int,nullzero"`                                                                                                                                                                // Custom column to to allow the user to pick a valid bin for the location to be the default warranty bin for RMA receipt.
	DamagedFlag                    string    `bun:"damaged_flag,type:char(1),nullzero"`                                                                                                                                                                            // Custom column indicates the location is for damaged item
	CardlockNexusLocation          string    `bun:"cardlock_nexus_location,type:char(1),nullzero"`                                                                                                                                                                 // Check Y for the location being used as the Company/taxing location for the Cardlock(CFN) functionality.
	FedexSmartpostHubCd            int32     `bun:"fedex_smartpost_hub_cd,type:int,nullzero"`                                                                                                                                                                      // Fedex SmartPost service Hub id code.
	DatagateSupplierNo             int32     `bun:"datagate_supplier_no,type:int,nullzero"`                                                                                                                                                                        // For the Ford Datagate extraction (export), specifies the Datagate supplier number exported in section 102 of the data.  This field should never contain a number larger than 99999 as the export field size is only 5 char's.
	DatagateInvSeqNo               int32     `bun:"datagate_inv_seq_no,type:int,nullzero"`                                                                                                                                                                         // File tracking (sequence) number for the Datagate inventory export.
	DatagateSalesSeqNo             int32     `bun:"datagate_sales_seq_no,type:int,nullzero"`                                                                                                                                                                       // File tracking (sequence) number for the Datagate sales export.
	PriorityLocationId             float64   `bun:"priority_location_id,type:decimal(19,0),nullzero"`                                                                                                                                                              // Custom column to specify an existing location as a priority location
	ExtTaxFreightBkupTaxExempt     string    `bun:"ext_tax_freight_bkup_tax_exempt,type:char(1),default:('N')"`                                                                                                                                                    // 3rd party taxing freight backup tax exemption
	SporadicItemThresholdPeriods   int32     `bun:"sporadic_item_threshold_periods,type:int,nullzero"`                                                                                                                                                             // Represents the number of zero usage periods in the past year that determines if the item is sproradic
	RfnavigatorUrl                 string    `bun:"rfnavigator_url,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"` // Location specific endpoint URL for RF Navigator system.
	RfnavigatorUserName            string    `bun:"rfnavigator_user_name,type:varchar(255),nullzero"`                                                                                                                                                              // User name for communication with RF Navigator.
	RfnavigatorPassword            string    `bun:"rfnavigator_password,type:varchar(255),nullzero"`                                                                                                                                                               // Password for communication with RF Navigator.
	UpdateVelocitySettings         string    `bun:"update_velocity_settings,type:varchar(1),default:('N')"`
	TrackaboutLocationTid          int32     `bun:"trackabout_location_tid,type:int,nullzero"`          // TrackAbout internal id for a P21 location
	RfnavigatorDropLocation        string    `bun:"rfnavigator_drop_location,type:varchar(8),nullzero"` // Staging/Drop Location
	RfnavigatorStartUnitNo         int32     `bun:"rfnavigator_start_unit_no,type:int,nullzero"`        // Starting unit number for transfer returns
}
