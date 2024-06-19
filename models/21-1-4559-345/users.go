package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Users struct {
	bun.BaseModel                     `bun:"table:users"`
	Id                                string    `bun:"id,type:varchar(30),pk"`                                        // What is the unique identifier for this counter?
	Name                              string    `bun:"name,type:varchar(40)"`                                         // This is the name of the user.
	Password                          string    `bun:"password,type:varchar(40),nullzero"`                            // What is the user
	DateCreated                       time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified                  time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy                  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`     // ID of the user who last maintained this record
	Active                            string    `bun:"active,type:char(1)"`                                           // This column is unused.
	DeleteFlag                        string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	LanguageId                        string    `bun:"language_id,type:varchar(8),default:(1)"`                       // What is the unique identifier for this language?
	DefaultCompany                    string    `bun:"default_company,type:varchar(8),nullzero"`                      // What is the default company for this user?
	DefaultBranch                     string    `bun:"default_branch,type:varchar(8),nullzero"`                       // In selected areas, branch id will default to this value.
	DefaultQuoteOrder                 string    `bun:"default_quote_order,type:char(1),default:('N')"`                // Indicates whether to default OE window to a quote or order.
	CreateCustomers                   string    `bun:"create_customers,type:char(1),nullzero"`                        // Can this user create customers?
	CreateShipTos                     string    `bun:"create_ship_tos,type:char(1),nullzero"`                         // Can this user create Ship To data?
	DefaultLocationId                 float64   `bun:"default_location_id,type:decimal(19,0),nullzero"`               // In selected areas, location id will default to this value.
	PromptBeforeClearing              string    `bun:"prompt_before_clearing,type:char(1),nullzero"`                  // Should this user be prompted before clearing?
	DefaultCustomerId                 float64   `bun:"default_customer_id,type:decimal(19,0),nullzero"`               // What is the default customer for this user?
	RemoteUser                        string    `bun:"remote_user,type:char(1)"`                                      // This was used for a prototype on-site package.
	CompanySecurity                   string    `bun:"company_security,type:char(1)"`                                 // Indicates that company security is used.
	WindowSecurity                    string    `bun:"window_security,type:char(1)"`                                  // This column is unused.
	SystemSecurity                    string    `bun:"system_security,type:char(1)"`                                  // This column is unused.
	CreateItemsAtOe                   string    `bun:"create_items_at_oe,type:char(1),nullzero"`                      // Can this user create items while in Order Entry?
	DefaultApplication                string    `bun:"default_application,type:varchar(40),nullzero"`                 // This is the module that will be opened after logging in.
	ScriptPath                        string    `bun:"script_path,type:varchar(255),nullzero"`                        // This is the destination folder for CommerceCenter script. Used primarily by Jetform.
	CreatePoFromOe                    int32     `bun:"create_po_from_oe,type:int"`                                    // Indicates whether a user can generate purchase orders from order entry.
	RoleUid                           int32     `bun:"role_uid,type:int"`                                             // DynaChange Designer access role.
	DesignerRights                    string    `bun:"designer_rights,type:varchar(8),default:('USER')"`              // Indicates the permissions to DynaChange allowed to a particular user.
	ContactId                         string    `bun:"contact_id,type:varchar(16),nullzero"`                          // What contact deals with this sales representative?
	PrintPreviewZoomPercentage        int16     `bun:"print_preview_zoom_percentage,type:smallint"`                   // Default zoom percentage for print preview.
	ClassId                           string    `bun:"class_id,type:varchar(8),nullzero"`                             // Identifier for the class.
	AutoGenerateTransferInOe          string    `bun:"auto_generate_transfer_in_oe,type:char(1),default:('N')"`       // Indicates whether a user can generate transfers from order entry.
	EmailAddress                      string    `bun:"email_address,type:varchar(255),nullzero"`                      // What is the email address for this contact?
	ReceiveSystemAlerts               string    `bun:"receive_system_alerts,type:char(1),default:('N')"`              // Indicates whether a user receives alerts.
	SalesrepId                        string    `bun:"salesrep_id,type:varchar(16),nullzero"`                         // If the user is a Sales Representative, this column stores the Sales Rep ID for the user.
	ReceiveImportFailureAlert         string    `bun:"receive_import_failure_alert,type:char(1),default:('N')"`       // Enabled users will receive an email alert on a scheduled import failure.
	DefaultToAdvancedSearch           string    `bun:"default_to_advanced_search,type:char(1),default:('N')"`         // indicates if user will default to the advanced search mode of the item popup window
	AdjQtyAvailableAtOe               int32     `bun:"adj_qty_available_at_oe,type:int,default:(1391)"`               // A column that sets whether a user is allowed to do inventory adjustments at OE
	LinkStockItemToPo                 string    `bun:"link_stock_item_to_po,type:char(1),default:('N')"`              // Indicates if a user has rights to link/un-link stock items to/from Purchase Order from outside the purchasing module in areas such as Order Entry, Production Order, MSP, Etc.
	HistoricalCostInSalesHist         string    `bun:"historical_cost_in_sales_hist,type:char(1),default:('N')"`      // Will allow distributer to enable the user to view historical cost information in OE
	OrderCostBasisOrderCost           string    `bun:"order_cost_basis_order_cost,type:char(1),default:('Y')"`        // IF Y, Order Cost is available as an option in OE totals tab
	OrderCostBasisCommCost            string    `bun:"order_cost_basis_comm_cost,type:char(1),default:('Y')"`         // IF Y, Commission Cost is available as an option in OE totals tab
	OrderCostBasisOtherCost           string    `bun:"order_cost_basis_other_cost,type:char(1),default:('Y')"`        // IF Y, Other Cost is available as an option in OE totals tab
	DefaultCostingBasis               int32     `bun:"default_costing_basis,type:int,default:(1)"`                    // One of the three values: Order, Commission Or Other cost
	TimeZoneCd                        int32     `bun:"time_zone_cd,type:int,nullzero"`                                // Default time zone for the user
	CurrentCashdrawerCounter          int32     `bun:"current_cashdrawer_counter,type:int,default:(0)"`               // Current counter reserved for cash_drawer_transaction
	MaxCashdrawerCounter              int32     `bun:"max_cashdrawer_counter,type:int,default:(0)"`                   // Max counter reserved for cash drawer transaction table
	CurrentOeLineCounter              int32     `bun:"current_oe_line_counter,type:int,default:(0)"`                  // Current counter for oe_line.oe_line_uid
	MaxOeLineCounter                  int32     `bun:"max_oe_line_counter,type:int,default:(0)"`                      // Max counter reserved for oe_line.oe_line_uid
	CurrentInvoiceLineCounter         int32     `bun:"current_invoice_line_counter,type:int,default:(0)"`             // Current counter reserved for invoice_line.invoice_line_Uid
	MaxInvoiceLineCounter             int32     `bun:"max_invoice_line_counter,type:int,default:(0)"`                 // Max counter reserved for invoice_line.invoice_line_uid
	TaskVisibilityCd                  int32     `bun:"task_visibility_cd,type:int,default:(1418)"`                    // Indicates tasks potentially viewable by user
	CreateVendorRfqCd                 int32     `bun:"create_vendor_rfq_cd,type:int,default:(986)"`                   // Indicates where the user is allowed to create vendor RFQs.
	UpdateCostFromRfqReceipt          string    `bun:"update_cost_from_rfq_receipt,type:char(1),default:('N')"`       // Indicates whether the user is enabled to update the supplier cost upon receipt of an RFQ.
	DisplayTransactionTasks           string    `bun:"display_transaction_tasks,type:char(1),default:('Y')"`          // Display tasks related to transactions when retrieved
	AllowNonstockTbo                  string    `bun:"allow_nonstock_tbo,type:char(1),default:('N')"`                 // Indicate whether user is allowed to create nonstock transfer backorders
	CurrentInvTranCounter             int32     `bun:"current_inv_tran_counter,type:int,default:(0)"`                 // The next value to use for transaction_number on inv_tran
	MaxInvTranCounter                 int32     `bun:"max_inv_tran_counter,type:int,default:(0)"`                     // The last counter we can use for transaction_number on inv_tran before reserving a new block from the counter table.
	CurrentGlCounter                  int32     `bun:"current_gl_counter,type:int,default:(0)"`                       // The counter value last used for the GL for this user
	MaxGlCounter                      int32     `bun:"max_gl_counter,type:int,default:(0)"`                           // The last value in the reserved block of GL values for this user
	CreateContractFromOe              string    `bun:"create_contract_from_oe,type:char(1),default:('N')"`            // Can this user create contract while in Order Entry
	MgmtAllowBranchEdit               string    `bun:"mgmt_allow_branch_edit,type:char(1),default:('Y')"`             // Denotes whether the brach can be edited in the Mgmt Summary
	MgmtUseDefaultBranch              string    `bun:"mgmt_use_default_branch,type:char(1),default:('N')"`            // Indicates that mgmt summary should default the view to the user default brach
	OePriceLibraryOverrideFlag        string    `bun:"oe_price_library_override_flag,type:char(1),default:('N')"`     // Custom: Whether the user may override price libraries in OE
	DefaultLabelPrinter               string    `bun:"default_label_printer,type:varchar(255),nullzero"`              // Label printer this user uses.
	UpdateProspectsOnlyFlag           string    `bun:"update_prospects_only_flag,type:char(1),default:('N')"`         // User will only be able to create and edit prospects if flag is Y
	DefaultItemSearch                 int32     `bun:"default_item_search,type:int,default:(1935)"`                   // Indicate the type of item the user wants to filter when doing a search item.
	SuppressManualAdjAllocFlag        string    `bun:"suppress_manual_adj_alloc_flag,type:char(1),default:('N')"`     // Indicates whether manual adjustment warning in PO and Transfer Receipts is suppressed
	ShippingControlFlag               string    `bun:"shipping_control_flag,type:char(1),default:('N')"`              // This will determine if the user will be allowed to confirm overships and/or underships
	CnvrtProspectToCustomerOe         string    `bun:"cnvrt_prospect_to_customer_oe,type:char(1),default:('N')"`      // Determines whether or not the user can convert prospects to customers in Order Entry
	AutoDisplayRooms                  string    `bun:"auto_display_rooms,type:char(1),default:('N')"`                 // user setting to show the rooms popup automatically thoughout the system
	DisplayPurchasepriceBreaks        string    `bun:"display_purchaseprice_breaks,type:char(1),default:('N')"`       // Determines whether or not to display Purchase Pricing breaks for the user
	DefaultSalesrepOnOrder            string    `bun:"default_salesrep_on_order,type:varchar(16),nullzero"`           // This salesrep will be defaulted on orders taken by this user.
	DefaultRepOnCommRptFlag           string    `bun:"default_rep_on_comm_rpt_flag,type:char(1),default:('N')"`       // Indicates whether the salesrep will be defaulted on the commission report criteria.
	OrderValidationPassword           string    `bun:"order_validation_password,type:varchar(255),nullzero"`          // Password column to validate hold orders
	BypassCheckVerifyPassword         string    `bun:"bypass_check_verify_password,type:varchar(255),nullzero"`       // Identifies a user who is able to bypass verifying a credit card payment
	MachineName                       string    `bun:"machine_name,type:varchar(255),nullzero"`                       // The hostname() of the machine this user connects with
	SalesSupervisorFlag               string    `bun:"sales_supervisor_flag,type:char(1),default:('N')"`              // Indicates whether the user is a sales supervisor
	VacationEndDate                   time.Time `bun:"vacation_end_date,type:datetime,nullzero"`                      // The last day of a users scheduled vacation
	DoeSalesrepId                     string    `bun:"doe_salesrep_id,type:varchar(16),nullzero"`                     // The salesrep_id that this user logs orders with in disconnected mode
	DoeUserFlag                       string    `bun:"doe_user_flag,type:char(1),default:('N')"`                      // Indicates that this account is a DOE user
	VacationEndDateModFlag            string    `bun:"vacation_end_date_mod_flag,type:char(1),default:('N')"`         // Flag that indicates whether the vacation_end_date was just modified for this user.  This is currently used by the alert.
	UsersUid                          int32     `bun:"users_uid,type:int,autoincrement"`                              // Unique identifier on the users table.
	InvGroupHdrUid                    int32     `bun:"inv_group_hdr_uid,type:int,nullzero"`                           // Inventory mgmt group assigned to this user.
	UsePoApprovalThresholdFlag        string    `bun:"use_po_approval_threshold_flag,type:char(1),default:('N')"`     // Flag to determine whether or not this user should dollar threshold to limit the approval of POs
	PoApprovalThreshold               float64   `bun:"po_approval_threshold,type:decimal(19,9),nullzero"`             // The dollar amount that limits the approval of POs
	PrevRequestsSearch                string    `bun:"prev_requests_search,type:char(1),default:('N')"`               // Default OE item search to use previous requests popup
	ViewCostOnOeLine                  string    `bun:"view_cost_on_oe_line,type:varchar(1),default:('N')"`            // Determines whether cost and profit is shown on the order line.
	ActiveDirectoryRole               string    `bun:"active_directory_role,type:varchar(255),nullzero"`              // Role kept in Active Directory and updated when the user logs into P21
	UpdateBranchOeReportsFlag         string    `bun:"update_branch_oe_reports_flag,type:char(1),nullzero"`           // The column is to indicate whether user can update branch id for sales order report and wbw report.
	OeAllowShipmentEdits              string    `bun:"oe_allow_shipment_edits,type:char(1),nullzero"`                 // The column is to indicate whether a user can edit a shipment in Shipping window after packing list has been printed
	OeAllowPackingListReprint         string    `bun:"oe_allow_packing_list_reprint,type:char(1),nullzero"`           // The column is to indicate whether a user can reprint a packing list in Shipping window after original packing list has been printed
	SalesrepSecurityType              string    `bun:"salesrep_security_type,type:char(1),nullzero"`                  // Salesrep Security Restriction level
	UserSalesrepId                    string    `bun:"user_salesrep_id,type:varchar(16),nullzero"`                    // Salesrep ID associated with user for security purposes
	DefaultAsTakerFlag                string    `bun:"default_as_taker_flag,type:char(1),default:('N')"`              // Default user as taker or require taker id login.
	SearchItemCatalogFlag             string    `bun:"search_item_catalog_flag,type:char(1),default:('N')"`           // Allow user to search from item_catalog
	ClearItemCatalogFlag              string    `bun:"clear_item_catalog_flag,type:char(1),default:('N')"`            // Allow user to  clear item_catalog
	MakeItemsSellableInOeFlag         string    `bun:"make_items_sellable_in_oe_flag,type:char(1),default:('N')"`     // Indicates whether the user will be able to make items sellable at a new location from OE
	AddItemLocationsInOeFlag          string    `bun:"add_item_locations_in_oe_flag,type:char(1),default:('N')"`      // Indicates whether the user will be able to add an item to a new location from OE
	ConfirmDeaPtFlag                  string    `bun:"confirm_dea_pt_flag,type:char(1),nullzero"`                     // Control whether or not a user can confirm DEA pick tickets
	StrategicPricingRoleUid           int32     `bun:"strategic_pricing_role_uid,type:int,nullzero"`                  // The strategic_pricing_role_uid for this users.
	CashDrawerId                      string    `bun:"cash_drawer_id,type:varchar(8),nullzero"`                       // Custom column for the user to default a cash drawer.
	DfltWwmsFormsPrinter              string    `bun:"dflt_wwms_forms_printer,type:varchar(255),nullzero"`            // Custom (F40090): default WWMS forms printer - user's default printer to print unconfirmed packing lists
	ShowComponentsInSalesHist         string    `bun:"show_components_in_sales_hist,type:char(1),default:('N')"`      // Determines whether components show in certain views of sales history
	FieldChooserPermCd                int32     `bun:"field_chooser_perm_cd,type:int,default:((300))"`                // Control user permissions for adding columns to field chooser and the database.
	CustomerProfitRoleUid             int32     `bun:"customer_profit_role_uid,type:int,nullzero"`                    // Customer Profitability Dashboard Role
	AllowPostToClosedGlPeriod         string    `bun:"allow_post_to_closed_gl_period,type:char(1),nullzero"`          // A flag to determine whether a user can post to a closed GL period
	DefaultSendToOutlookFlag          string    `bun:"default_send_to_outlook_flag,type:char(1),default:('N')"`       // Flag to determine if the user should default to send task to outlook
	DisplayOpenQuoteLinesCd           int32     `bun:"display_open_quote_lines_cd,type:int,nullzero"`                 // Allows the user to price their order lines from previous quotes in order entry.
	CreateItemsAtSoe                  string    `bun:"create_items_at_soe,type:char(1),nullzero"`                     // Can this user create items while in Service Order Entry?
	RestrictedClassOverrideFlag       string    `bun:"restricted_class_override_flag,type:char(1),default:('N')"`     // This column indicates whether the user can override restricted class options - values Y, N
	AllowShipToEditInOeFlag           string    `bun:"allow_ship_to_edit_in_oe_flag,type:char(1),default:('N')"`      // Custom flag indicating whether the user is allowed to edit existing Ship To's in Order Entry.
	ShowCcPaymentAcctOnSave           int32     `bun:"show_cc_payment_acct_on_save,type:int,nullzero"`                // Display Element Payment Account Response on th eselected areas
	ShowArCcFailed                    string    `bun:"show_ar_cc_failed,type:char(1),nullzero"`                       // Display Process to A/R option when Credit Card process fails
	OverrideCustPalletWarnFlag        string    `bun:"override_cust_pallet_warn_flag,type:char(1),default:('N')"`     // Designates the users ability to override customer pallet warning during truck complete routine
	LogoPathFilename                  string    `bun:"logo_path_filename,type:varchar(255),nullzero"`                 // Logo path definable at the user level.
	UserReportPath                    string    `bun:"user_report_path,type:varchar(255),nullzero"`                   // User level custom forms path.
	AllowShipmentConfirmation         string    `bun:"allow_shipment_confirmation,type:char(1),default:('N')"`        // Flag to determine whether the user is allowed to do confirmation in Shipping.
	ViewVendorTypeCd                  int32     `bun:"view_vendor_type_cd,type:int,default:((2981))"`                 // Indicates whether user is restricted in VSM to only viewing vendors of type Inventory or whether user can view both Inventory and Expense type vendors.
	OverRedeemRewardsFlag             string    `bun:"over_redeem_rewards_flag,type:char(1),default:('N')"`           // Rewards program functionality: determines if this user can redemm more rewards points than a customer has accumulated.
	ExtendedItemInfoPorgFlag          string    `bun:"extended_item_info_porg_flag,type:char(1),nullzero"`            // Indicates whether to show extended item information in PORG for a user.
	FedexLabelPrinterPath             string    `bun:"fedex_label_printer_path,type:varchar(255),nullzero"`           // Fedex Label printer for a user
	DefaultItemSearchInImi            string    `bun:"default_item_search_in_imi,type:char(1),default:('N')"`         // Provide the ability to control whether or not the IMI item popups search all locations independently in IMI.
	PoThresholdCurrencyId             int32     `bun:"po_threshold_currency_id,type:int,nullzero"`                    // Currency ID for PO Approval Threshold
	AddCustomerPartNumberInOe         string    `bun:"add_customer_part_number_in_oe,type:varchar(1),default:('N')"`  // Add Customer Part Number in OE
	Override754Flag                   string    `bun:"override_754_flag,type:char(1),default:('N')"`                  // Allow shipping confirmation without 754 info
	AllowAddLaborToCompletedPo        string    `bun:"allow_add_labor_to_completed_po,type:char(1),default:('N')"`    // Allow adding new labor to a completed production order.
	AllowPostprintEditLaborEst        string    `bun:"allow_postprint_edit_labor_est,type:char(1),default:('N')"`     // Indicates whether this user is allowed to edit the estimated labor on a production order after the production order form has been printed.
	DoNotExportCarrierPoFlag          string    `bun:"do_not_export_carrier_po_flag,type:char(1),default:('N')"`      // Determines if a po can not be exported to Carrier.
	NetworkName                       string    `bun:"network_name,type:varchar(30),nullzero"`                        // Stores alternative login name
	OrderCostBasisRebateCost          string    `bun:"order_cost_basis_rebate_cost,type:char(1),default:('N')"`       // Use carrier rebate cost
	ConvertQuoteToOrderPromptFlag     string    `bun:"convert_quote_to_order_prompt_flag,type:char(1),default:('Y')"` // Indicates whether this user gets prompted with making quote to an order in order entry.
	DefaultPickTicketPrinter          string    `bun:"default_pick_ticket_printer,type:varchar(255),nullzero"`        // Default printer for pick tickets in FCOE
	DefaultInvoicePrinter             string    `bun:"default_invoice_printer,type:varchar(255),nullzero"`            // Default printer for invoices in FCOE
	UserSignaturePath                 string    `bun:"user_signature_path,type:varchar(255),nullzero"`
	DisableItemCategoryTreeByDescFlag string    `bun:"disable_item_category_tree_by_desc_flag,type:char(1),default:('N')"`
	ImiDefaultUserLocation            string    `bun:"imi_default_user_location,type:char(1),default:('N')"`      // set location id in Item master inquiry default to location id from user
	DfltUcc128LabelPrinter            string    `bun:"dflt_ucc128_label_printer,type:varchar(255),nullzero"`      // Custom (F63490): the default printer for container (UCC128) labels for the shipping confirmation import
	AllowEdi855ManualOrdersFlag       string    `bun:"allow_edi_855_manual_orders_flag,type:char(1),nullzero"`    // Setting to enable users for sending edi 855s for manual sales orders
	EmailSignature                    string    `bun:"email_signature,type:varchar(max),nullzero"`                // signature default for the user used in email
	OeChangeCustomerWithItems         string    `bun:"oe_change_customer_with_items,type:char(1),default:('N')"`  // Change customer ID in OE after items are added to order
	AllowEditLaborTimeEntry           string    `bun:"allow_edit_labor_time_entry,type:char(1),default:('N')"`    // Whether the user can edit Labor Time Entry
	DefaultBssCustomerId              float64   `bun:"default_bss_customer_id,type:decimal(19,0),nullzero"`       // Custom: Indicates the customer id to which OE should default when associated user’s OE defaults to a Builder Selection Sheet.
	RebuildDrillSecurityFlag          string    `bun:"rebuild_drill_security_flag,type:char(1),default:('N')"`    // Determine if security needs to be rebuilt (new updates, dc menu changes).
	CheckCreationRole                 int32     `bun:"check_creation_role,type:int,nullzero"`                     // Role that describes the permission that the user has to interact with the
	ShowCcOnSaveCashReceipts          string    `bun:"show_cc_on_save_cash_receipts,type:char(1),default:('N')"`  // Flag to determine if token creation can happen at save time
	P21ClientLogPath                  string    `bun:"p21_client_log_path,type:varchar(255),nullzero"`            // Allows user to set a path for application error logging
	EnterpriseSearchUrl               string    `bun:"enterprise_search_url,type:varchar(255),nullzero"`          // URL path for Enterprise Search
	AllowReportCreationInStudio       int32     `bun:"allow_report_creation_in_studio,type:int,default:((1418))"` // Indicates whether users can create a report and who for
}
