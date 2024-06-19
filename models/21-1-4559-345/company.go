package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Company struct {
	bun.BaseModel                     `bun:"table:company"`
	CompanyId                         string    `bun:"company_id,type:varchar(8),pk"`                                  // Unique code that identifies a company.
	CompanyName                       string    `bun:"company_name,type:varchar(40)"`                                  // Name of company
	FederalIdNo                       string    `bun:"federal_id_no,type:varchar(255),nullzero"`                       // A number assigned to your company by the government.  This number is used for the IRS 1099 form and for various other payroll purposes.
	FiscalYearEnd                     string    `bun:"fiscal_year_end,type:varchar(5),nullzero"`                       // The number of the month in which a company’s fiscal year ends.
	DateCreated                       time.Time `bun:"date_created,type:datetime"`                                     // Indicates the date/time this record was created.
	DateLastModified                  time.Time `bun:"date_last_modified,type:datetime"`                               // Indicates the date/time this record was last modified.
	LastMaintainedBy                  string    `bun:"last_maintained_by,type:varchar(30),default:(suser_name(null))"` // ID of the user who last maintained this record
	AddressId                         float64   `bun:"address_id,type:decimal(19,0)"`                                  // Identifier of the address for this company.
	DeleteFlag                        string    `bun:"delete_flag,type:char(1)"`                                       // Indicates whether this record is logically deleted
	RePriceLineItems                  string    `bun:"re_price_line_items,type:char(1),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	PricingByLocation                 string    `bun:"pricing_by_location,type:char(1),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	GlAccountLength                   string    `bun:"gl_account_length,type:varchar(32)"`                             // Determines the length designated for the General Ledger accounts.
	GlCapitalizedOverhead             string    `bun:"gl_capitalized_overhead,type:varchar(32),nullzero"`              // Account for tracking labor.
	GlFreightIn                       string    `bun:"gl_freight_in,type:varchar(32),nullzero"`                        // This column is no longer being used and is scheduled for removal in later version.
	NumberOfPeriods                   float64   `bun:"number_of_periods,type:decimal(3,0),nullzero"`                   // Number of fiscal year periods this company will use.
	RetainedEarningsAcctNo            string    `bun:"retained_earnings_acct_no,type:varchar(32),nullzero"`            // Balances that do not carry over will go into this account
	TermsTakenAccountNo               string    `bun:"terms_taken_account_no,type:varchar(32),nullzero"`               // The account where the amount of the discount awarded for payment of an invoice within the terms discount period is stored.
	HomeCurrencyId                    float64   `bun:"home_currency_id,type:decimal(19,0)"`                            // Indicates the home currency used for this company.
	GlGainOrLossAcct                  string    `bun:"gl_gain_or_loss_acct,type:varchar(32),nullzero"`                 // This column is no longer being used and is scheduled for removal in later version.
	San                               string    `bun:"san,type:varchar(15),nullzero"`                                  // Standard Address Number - for EDI.
	GroupControlNo                    string    `bun:"group_control_no,type:varchar(9),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	IntlSan                           string    `bun:"intl_san,type:varchar(15),nullzero"`                             // International Standard Address Number - for EDI.
	GstPaidAcct                       string    `bun:"gst_paid_acct,type:varchar(32),nullzero"`                        // This column is no longer being used and is scheduled for removal in later version.
	MaxInvoiceAllowance               float64   `bun:"max_invoice_allowance,type:decimal(19,4),nullzero"`              // The maximum amount that can be recorded as bad debt during one session of entering cash receipts for a customer.
	GstRegistrationNo                 string    `bun:"gst_registration_no,type:varchar(20),nullzero"`                  // This column is no longer being used and is scheduled for removal in later version.
	StageClearingAccountNo            string    `bun:"stage_clearing_account_no,type:varchar(32),nullzero"`            // Account used for tracking the work in process.
	TransmitterControlCode            string    `bun:"transmitter_control_code,type:varchar(5),nullzero"`              // This column is no longer being used and is scheduled for removal in later version.
	ForeignEntityIndicator            string    `bun:"foreign_entity_indicator,type:char(1),nullzero"`                 // Indicates that the foreign entity is enabled, 1099 forms printed for this company indicate that the company is not located in the United States.
	UseBranchAccounting               string    `bun:"use_branch_accounting,type:char(1),nullzero"`                    // Indicates if this company will use branch accounting or not.
	BranchSegmentInCoa                float64   `bun:"branch_segment_in_coa,type:decimal(2,0),nullzero"`               // Segment number in the account number where the branch id is stored.
	InvScan1                          float64   `bun:"inv_scan_1,type:decimal(1,0),nullzero"`                          // Indicates the first search order for an inventory scan.
	InvScan2                          float64   `bun:"inv_scan_2,type:decimal(1,0),nullzero"`                          // Indicates the second search order for an inventory scan.
	InvScan3                          float64   `bun:"inv_scan_3,type:decimal(1,0),nullzero"`                          // Indicates the third search order for an inventory scan.
	InvScan4                          float64   `bun:"inv_scan_4,type:decimal(1,0),nullzero"`                          // Indicates the fourth search order for an inventory scan.
	InvScan5                          float64   `bun:"inv_scan_5,type:decimal(1,0),nullzero"`                          // Indicates the fifth search order for an inventory scan.
	InvScan6                          float64   `bun:"inv_scan_6,type:decimal(1,0),nullzero"`                          // Indicates the sixth search order for an inventory scan.
	ArAgingPeriod1                    float64   `bun:"ar_aging_period1,type:decimal(3,0),nullzero"`                    // What is the first aging period for this company?
	ArAgingPeriod2                    float64   `bun:"ar_aging_period2,type:decimal(3,0),nullzero"`                    // What is the second aging period for this company?
	ArAgingPeriod3                    float64   `bun:"ar_aging_period3,type:decimal(3,0),nullzero"`                    // What is the third aging period for this company?
	JcFlag                            string    `bun:"jc_flag,type:char(1),nullzero"`                                  // Job Costing Flag
	InvReceiptsClearingAcct           string    `bun:"inv_receipts_clearing_acct,type:varchar(32),nullzero"`           // The account that stores the inventory value of goods that have been received against a PO while you are still waiting for the invoice for those goods.
	InterbranchRecPayableAcct         string    `bun:"interbranch_rec_payable_acct,type:varchar(32),nullzero"`         // An account that tracks payables and receivables made for other branches.
	BranchDebitEqualCredit            string    `bun:"branch_debit_equal_credit,type:char(1),nullzero"`                // If selected this will ensure that branches balance.
	FcPercentage                      float64   `bun:"fc_percentage,type:decimal(19,4),nullzero"`                      // Enter the default finance charge for all customers
	FcGraceDays                       float64   `bun:"fc_grace_days,type:decimal(19,4),nullzero"`                      // Enter the grace days past the net due date before
	MinimumFc                         float64   `bun:"minimum_fc,type:decimal(19,9),nullzero"`                         // Enter the minimum finance charge.
	MinimumFcToInvoice                float64   `bun:"minimum_fc_to_invoice,type:decimal(19,9),nullzero"`              // Enter the amount under which a finance charge invo
	WipCosAccountNumber               string    `bun:"wip_cos_account_number,type:varchar(32),nullzero"`               // Expense Account used for Multi-Stage Processing.
	CompoundFc                        string    `bun:"compound_fc,type:char(1),nullzero"`                              // Indicate whether you would like finance charges to
	FcAllowedAcct                     string    `bun:"fc_allowed_acct,type:varchar(32),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	FcInvoiceClass                    string    `bun:"fc_invoice_class,type:varchar(8),nullzero"`                      // Enter the invoice class for finance charges.
	GenerateFinanceCharges            string    `bun:"generate_finance_charges,type:char(1),nullzero"`                 // Indicate if you would like the system to calculate
	FcAllowedAccount                  string    `bun:"fc_allowed_account,type:varchar(32),nullzero"`                   // Enter the account to which finance charges should
	FcRevenueAccount                  string    `bun:"fc_revenue_account,type:varchar(32),nullzero"`                   // Enter the revenue account to be credited with fina
	UsingEncumbrances                 string    `bun:"using_encumbrances,type:char(1),nullzero"`                       // This column is no longer being used and is scheduled for removal in later version.
	CostToOrder                       float64   `bun:"cost_to_order,type:decimal(19,9),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	CarryingCost                      float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`                      // This column is no longer being used and is scheduled for removal in later version.
	UnitCost                          float64   `bun:"unit_cost,type:decimal(19,9),nullzero"`                          // This column is no longer being used and is scheduled for removal in later version.
	TrackApByBranch                   string    `bun:"track_ap_by_branch,type:char(1),nullzero"`                       // Allows you to maintain separate accounts payable for each branch in a company.
	TrackArByBranch                   string    `bun:"track_ar_by_branch,type:char(1),nullzero"`                       // Allows you to maintain separate accounts receivable for each branch in a company.
	TrackProdCapitalizedOh            string    `bun:"track_prod_capitalized_oh,type:char(1),nullzero"`                // Indicates whether capitalized overhead is tracked for labor.
	LogoPathFilename                  string    `bun:"logo_path_filename,type:varchar(255),nullzero"`                  // Allows the user to specify a logo bitmap file that
	CommissionExpenseAcct             string    `bun:"commission_expense_acct,type:varchar(32),nullzero"`              // Commission Expense Account number.
	AccrualAcct                       string    `bun:"accrual_acct,type:varchar(32),nullzero"`                         // This column is no longer being used and is scheduled for removal in later version.
	InvCostVarianceAcct               string    `bun:"inv_cost_variance_acct,type:varchar(32),nullzero"`               // The account that holds the difference between expected and actual inventory costs.
	InvFreightVarianceAcct            string    `bun:"inv_freight_variance_acct,type:varchar(32),nullzero"`            // The account that holds the difference between the freight amount entered during Inventory Receipt and that entered when the purchase order is converted to a voucher (i.e., when the vendor’s invoice has been received).
	InvQtyVarianceAcct                string    `bun:"inv_qty_variance_acct,type:varchar(32),nullzero"`                // The account that holds the difference between the inventory value received during Inventory Receipt and the inventory value when the purchase order is converted to a voucher (i.e., when the vendor’s invoice has been received).
	PostVarianceToAvgCost             string    `bun:"post_variance_to_avg_cost,type:char(1)"`                         // Determines whether item cost variances are posted to the moving average cost when POs are converted to vouchers.
	ProrateFreight                    string    `bun:"prorate_freight,type:char(1)"`                                   // This column is used to determine how to prorate a freight cost
	PostFreightToAvgCost              string    `bun:"post_freight_to_avg_cost,type:char(1)"`                          // This column will allow the user to determine whether freight variances are posted to the moving average cost when POs are converted to vouchers.
	FrlUseFrx                         string    `bun:"frl_use_frx,type:char(1),nullzero"`                              // Will FRX be used with this company?
	FrlAcctMask                       string    `bun:"frl_acct_mask,type:varchar(128),nullzero"`                       // What is the account mask -  for FRL purposes?
	FrlNumberOfSegments               int32     `bun:"frl_number_of_segments,type:int,nullzero"`                       // How many segments make up the account -  for FRL purposes?
	FrlNaturalSegment                 int32     `bun:"frl_natural_segment,type:int,nullzero"`                          // What is the natural segment -  for FRL purposes?
	CreateOrderBasedTransfers         string    `bun:"create_order_based_transfers,type:char(1)"`                      // Indicates whether an order-based transfer is generated when an order item’s Source and Ship Location IDs are different.
	MarkUpOnTransfers                 string    `bun:"mark_up_on_transfers,type:char(1)"`                              // Indicates whether the system will calculate a mark-up against the cost of an item at the Source Location when creating a transfer.
	MarkUpType                        string    `bun:"mark_up_type,type:char(1),nullzero"`                             // Indicates the source of the mark up of the order - (multiplier, percentage, value).
	MarkUpValue                       float64   `bun:"mark_up_value,type:decimal(19,9),nullzero"`                      // The amount by which the cost of the order item will increase when an order- based transfer is generated.
	SourceMarkUpAccount               string    `bun:"source_mark_up_account,type:varchar(32),nullzero"`               // The account that is used for the mark-up.
	MarkUpSource                      string    `bun:"mark_up_source,type:char(1),nullzero"`                           // Indicates the source cost of the mark-up when the Source Location and Ship Location IDs of an order item are different and an order-based transfer is generated.
	SourceOtherChargeAccount          string    `bun:"source_other_charge_account,type:varchar(32),nullzero"`          // Indicates if the source mark-up is for an other account (i.e. freight...)
	DestinationMarkUpAcct             string    `bun:"destination_mark_up_acct,type:varchar(32),nullzero"`             // Mark up account for the destination location in tr
	DestinationOtherChargeAcct        string    `bun:"destination_other_charge_acct,type:varchar(32),nullzero"`        // Other charge account for the destination location
	FrlLastTransferDate               string    `bun:"frl_last_transfer_date,type:smalldatetime,default:('1/1/1980')"` // This is the last time that the FRL Transfer process was run.  This will assist the FRX Transfer process in determining what records are new that have to be transfered.
	DefaultSourcePriceCd              int32     `bun:"default_source_price_cd,type:int,nullzero"`                      // The source price used in the default pricing calculation.  The default pricing calculation is used only when the system cannot find a price using the Pricing Libraries assigned to a customer.
	DefaultMultiplier                 float64   `bun:"default_multiplier,type:decimal(19,9),nullzero"`                 // This column is no longer being used and is scheduled for removal in later version.
	LandedCostClearingAcct            string    `bun:"landed_cost_clearing_acct,type:varchar(32),default:(null)"`      // The account that stores the total landed cost charges for a companys inventory.
	FrlMaxBatchId                     int32     `bun:"frl_max_batch_id,type:int,nullzero"`                             // The highest GL transaction number that has been processed by the FRx Transfer program.
	DefaultSalesLocationId            float64   `bun:"default_sales_location_id,type:decimal(19,0),nullzero"`          // This column is no longer being used and is scheduled for removal in later version.
	EdiExportPath                     string    `bun:"edi_export_path,type:varchar(255),nullzero"`                     // Path used when exported EDI transactions.
	FreightCodeUid                    int32     `bun:"freight_code_uid,type:int,nullzero"`                             // Indentifier for the default freight code.
	NumberOfDemandPeriods             int16     `bun:"number_of_demand_periods,type:smallint,default:(1)"`             // Number of fiscal periods in a year
	InvFreightExpenseAcct             string    `bun:"inv_freight_expense_acct,type:varchar(32),nullzero"`             // For a direct ship that does not charge freight to the customer, or when you pay freight for an Inventory Return, the freight amount posts to this account.
	RequisitionClearingAcct           string    `bun:"requisition_clearing_acct,type:varchar(32),nullzero"`            // The account that stores the inventory value of requisition goods that have been received against a PO while you are still waiting for the invoice for those goods.
	RestockingFeeAccountNo            string    `bun:"restocking_fee_account_no,type:varchar(32),nullzero"`            // This field is used to assign an expense account to which the restocking fees entered on the return will post.
	SignaturePath                     string    `bun:"signature_path,type:varchar(255),nullzero"`                      // The fully qualified path and filename where the signature bitmap file is stored after you synch your personal digital assistant with CommerceCenter after you have finished your deliveries.
	LotCostChangeAccountNo            string    `bun:"lot_cost_change_account_no,type:varchar(32),nullzero"`           // Changes in inventory value resulting from enabling or disabling Lot Costing are posted to this account for reporting purposes.
	RebateAccountNo                   string    `bun:"rebate_account_no,type:varchar(32),nullzero"`                    // Account used for rebates.
	RebateAllowanceAccountNo          string    `bun:"rebate_allowance_account_no,type:varchar(32),nullzero"`          // Allowed amount for rebates.
	AllowDiscountsOnPartialPmt        string    `bun:"allow_discounts_on_partial_pmt,type:char(1),default:('N')"`      // Enabling this option allows you to apply terms and discounts in Cash Receipts for invoices that are not paid in full.
	AutoApplyAvailDiscounts           string    `bun:"auto_apply_avail_discounts,type:char(1),default:('N')"`          // Indicates whether terms should always be taken.
	GracePeriodDays                   int32     `bun:"grace_period_days,type:int,default:(0)"`                         // Specifies the number of days added to a terms due date where the terms discount will still be offered.
	DisassemblyReasonId               string    `bun:"disassembly_reason_id,type:varchar(5),nullzero"`                 // Default reason for disassembly.
	TpcxMember                        string    `bun:"tpcx_member,type:char(1),default:('N')"`                         // Rhis column indicates that the company is a TPCx member.
	VendorConsignInvOffsetAcct        string    `bun:"vendor_consign_inv_offset_acct,type:varchar(32),nullzero"`       // Vendor Consign inventory Offset Acct
	FinanceChgInvDueDateCode          int32     `bun:"finance_chg_inv_due_date_code,type:int,default:(1248)"`          // Indicates whether the finance charge invoices due dates are calculated using the customers terms or whether the invoice is due immediately.
	RmaReceiptClearingAcct            string    `bun:"rma_receipt_clearing_acct,type:varchar(32),nullzero"`            // RMA receipt clearing account
	RmaAdjustmentReasonId             string    `bun:"rma_adjustment_reason_id,type:varchar(5),nullzero"`              // RMA inventory adjustment reason code
	RmaDefaultRtsLocationId           float64   `bun:"rma_default_rts_location_id,type:decimal(19,0),default:(null)"`  // RMA default return-to-stock location ID.
	RmaAutoReturnToStock              string    `bun:"rma_auto_return_to_stock,type:char(1),default:('N')"`            // Determines whether items are automtically returned to stock when confirming an RMA.
	InventoryAdjustmentReasonId       string    `bun:"inventory_adjustment_reason_id,type:varchar(5),nullzero"`        // Default inventory adjustment reason code
	ReturnedCheckFeeAcct              string    `bun:"returned_check_fee_acct,type:varchar(32),nullzero"`              // Account for any charges billed to the customer for returned checks.
	AllowableInvoiceVariance          float64   `bun:"allowable_invoice_variance,type:decimal(19,9),nullzero"`         // Allowable amount / percentage by which Imported AR Payments can vary from the Invoice amount.
	InvoiceVarianceTypeFlag           string    `bun:"invoice_variance_type_flag,type:char(1),default:('A')"`          // Whether the allowable invoice variance is specified in Amount or Percentage. Only possible values are A [Amount] and P [Percentage]
	CreatePrepaidInvoiceFlag          string    `bun:"create_prepaid_invoice_flag,type:char(1),default:('Y')"`         // Whether a prepaid invoice would be created if the Imported AR Payment is more than the Invoice amount.
	TagSuffix                         string    `bun:"tag_suffix,type:varchar(255),nullzero"`                          // Suffix to add to auto tag number for company
	DataIdentifierGroupUid            int32     `bun:"data_identifier_group_uid,type:int,nullzero"`                    // UID of the Data Identifier Group record to use for this company as its default
	RfScannerPrefix                   string    `bun:"rf_scanner_prefix,type:varchar(255),nullzero"`                   // Data prefixed by RF scanner device when barcode is scanned
	DefaultArpymtBankNo               float64   `bun:"default_arpymt_bank_no,type:decimal(19,0),nullzero"`             // Default Bank Number for AR Payment imports
	DefaultArpymtBranchId             string    `bun:"default_arpymt_branch_id,type:varchar(8),nullzero"`              // Default Branch ID for AR Payment imports
	CarrierInsuranceAcct              string    `bun:"carrier_insurance_acct,type:varchar(32),nullzero"`               // Account used for carrier insurance revenue.
	PostLcVarianceToAvgCost           string    `bun:"post_lc_variance_to_avg_cost,type:char(1),default:('N')"`        // Determines whether Landed Cost variances are posted to the moving average cost when Landed Cost is converted to vouchers.
	RestockFeePercentage              float64   `bun:"restock_fee_percentage,type:decimal(19,9),default:(0)"`          // The value is used to determine the restocking fee on an RMA
	RestockFeeAccount                 string    `bun:"restock_fee_account,type:varchar(32),nullzero"`                  // Indicated the revenue account that will be credited when a restocking fee is applied to an RMA
	CouponClearingAcctNo              string    `bun:"coupon_clearing_acct_no,type:varchar(32),nullzero"`              // Specifies the clearing (liability) account number to hold the total coupon value after purchase and before redemption.
	DisassemblyCostVarAcctNo          string    `bun:"disassembly_cost_var_acct_no,type:varchar(32),nullzero"`         // G/L Account for Disassembly Cost Variance
	CouponIssuedAcctNo                string    `bun:"coupon_issued_acct_no,type:varchar(32),nullzero"`                // a unbilled AR/Asset account
	JobIdRequiredFlag                 string    `bun:"job_id_required_flag,type:char(1),default:('N')"`                // If 'Y', job id will be required throughout the system for this company.
	AssemblyCostVarianceAcct          string    `bun:"assembly_cost_variance_acct,type:varchar(255),nullzero"`         // This will be an account that will be used to post a variance if the custom changes the cost on an assembly.
	TrackLostSalesFlag                string    `bun:"track_lost_sales_flag,type:char(1),default:('N')"`               // Track lost sales on transactions
	RestockingFeeTaxableFlag          string    `bun:"restocking_fee_taxable_flag,type:char(1),default:('Y')"`         // Indicates whether a restocking fee on an RMA is taxable or not.
	CompanyMacRoundingAcct            string    `bun:"company_mac_rounding_acct,type:varchar(32),nullzero"`            // The account used to store rounding errors when company MAC adjustments are applied
	SpecialInventoryAcct              string    `bun:"special_inventory_acct,type:varchar(32),nullzero"`               // Account number used to track inventory value of special disposition items.
	InvInVesselAcct                   string    `bun:"inv_in_vessel_acct,type:varchar(32),nullzero"`                   // Inventory on vessel GL account number
	InvInTransitClearingAcct          string    `bun:"inv_in_transit_clearing_acct,type:varchar(32),nullzero"`         // Inventory In-transit clearing account
	CommissionReceivableAcct          string    `bun:"commission_receivable_acct,type:varchar(32),nullzero"`           // The default commission receivable account when manfucturer rep is enabled.
	CommissionRevenueAcct             string    `bun:"commission_revenue_acct,type:varchar(32),nullzero"`              // The default commission revenue account when manfucturer rep is enabled.
	CommissionAllowanceAcct           string    `bun:"commission_allowance_acct,type:varchar(32),nullzero"`            // The default commission allowance account when manfucturer rep is enabled.
	LotBillClearingAcct               string    `bun:"lot_bill_clearing_acct,type:varchar(32),nullzero"`               // liability account that maintains the value of all lot bill shipments whose cost is determined at the header level
	AssemblyVarianceAcct              string    `bun:"assembly_variance_acct,type:varchar(32),nullzero"`               // When use standard cost for assemblies is enabled, this account holds the difference between the standard cost and actual cost of the assembly.
	GlAcctForAppliedLabor             string    `bun:"gl_acct_for_applied_labor,type:varchar(32),nullzero"`            // G/L account for applied labor
	GlAcctForLaborCos                 string    `bun:"gl_acct_for_labor_cos,type:varchar(32),nullzero"`                // G/L account for labor COS
	GlAcctForFreightClearing          string    `bun:"gl_acct_for_freight_clearing,type:varchar(32),nullzero"`         // G/L account for freight clearing
	CurrencyBankClearingAcct          string    `bun:"currency_bank_clearing_acct,type:varchar(32),nullzero"`          // Currency Bank Clearing Account that will be used in Cash Receipts - this will be used when Item Specific Currency System Setting is enabled
	PostCommissionAccrualsFlag        string    `bun:"post_commission_accruals_flag,type:char(1),default:('N')"`       // Determines whether to post commission accurals to the gl once the commission report is approved.
	UpsAccountNo                      string    `bun:"ups_account_no,type:varchar(255),nullzero"`                      // UPS account number that will be the default for all locations of this company
	UpsOltUserId                      string    `bun:"ups_olt_user_id,type:varchar(255),nullzero"`                     // UPS user id to access the UPS online tools. (Encrypted)
	UpsOltPassword                    string    `bun:"ups_olt_password,type:varchar(255),nullzero"`                    // UPS password to access the UPS online tools. (Encrypted)
	UpsOltAccessKey                   string    `bun:"ups_olt_access_key,type:varchar(255),nullzero"`                  // UPS access key for the online tools. (Encrypted)
	SpAllocFromStockVarAcct           string    `bun:"sp_alloc_from_stock_var_acct,type:varchar(32),nullzero"`         // Account used to store the difference between the cost used to relieve inventory, and the cost of goods sold for specials when special orders are filled out of stock instead of via the special's purchase order. Currently only used by custom.
	DunsNo                            string    `bun:"duns_no,type:varchar(255),nullzero"`                             // To be used for 832 export.
	GlAcctForLaborWip                 string    `bun:"gl_acct_for_labor_wip,type:varchar(32),nullzero"`                // The column is for Labor WIP account of a company
	DefaultFaxCoverUid                int32     `bun:"default_fax_cover_uid,type:int,nullzero"`                        // UID of the fax cover to be used for this company
	PredefinedCoaCd                   int32     `bun:"predefined_coa_cd,type:int,nullzero"`                            // Status of Predefined Chart of Accounts functionality in Company Maintenance (Not Enabled, Enabled & In Use).
	GlAcctLongtermContracts           string    `bun:"gl_acct_longterm_contracts,type:varchar(32),nullzero"`           // GL acct used to offset revenue during progress billing.
	ProdAutoTransferReasonId          string    `bun:"prod_auto_transfer_reason_id,type:varchar(5),nullzero"`          // Reason ID for inventory adjusments when auto-transferring in Production Order Pick Ticket Confirmation
	CouponAllowedAcctNo               string    `bun:"coupon_allowed_acct_no,type:varchar(32),nullzero"`               // Custom (F17388): GL account number to hold the difference between a coupon's value and the amount billed to the coupon's client.
	PedigreeContactId                 string    `bun:"pedigree_contact_id,type:varchar(16),nullzero"`                  // Indicates contact id designated as a pedigree contact for this company
	ValidateCommissionSplitFlag       string    `bun:"validate_commission_split_flag,type:char(1),default:('N')"`      // Indicates whether salesrep commission splits will be required to add up to 100%.
	DunsNumber                        string    `bun:"duns_number,type:varchar(255),nullzero"`                         // Dun and Bradstreet assigned identifier for a company.
	AcceptPymtsWithNoInvFlag          string    `bun:"accept_pymts_with_no_inv_flag,type:char(1),default:('N')"`       // Flag to indicate whether AR Payments Import would accept payments without any specified invoice number
	AcceptIncorrectInvFlag            string    `bun:"accept_incorrect_inv_flag,type:char(1),default:('N')"`           // Flag to indicate whether AR Payments Import would accept payments with incorrect invoice number
	AcceptPartialPaymentFlag          string    `bun:"accept_partial_payment_flag,type:char(1),default:('N')"`         // Flag to indicate whether AR Payments Import would accept partial payments
	BranchStart                       int16     `bun:"branch_start,type:smallint,nullzero"`                            // Starting position of branch (disregarding mask / dashes)
	BranchLength                      int16     `bun:"branch_length,type:smallint,nullzero"`                           // Length of branch segment
	AccountLength                     int16     `bun:"account_length,type:smallint,nullzero"`                          // Length of account (disregarding mask / dashes)
	VertexEnabled                     string    `bun:"vertex_enabled,type:char(1),default:('N')"`                      // Flag indicating whether this company use vertex integration
	DefaultSalesTaxPayableAcct        string    `bun:"default_sales_tax_payable_acct,type:varchar(32),nullzero"`       // Default Sales Tax Payable Account No
	PallRepNumber                     string    `bun:"pall_rep_number,type:varchar(5),nullzero"`                       // Holds a rep number for distributers that order from the vendor Pall
	IntrastatEnabledFlag              string    `bun:"intrastat_enabled_flag,type:char(1),default:('N')"`              // Determines if this company will use the Intrastat functionality.
	UseIntAddressFormatFlag           string    `bun:"use_int_address_format_flag,type:char(1),default:('N')"`         // Indicates whether the company should use international address formatting for addresses on invoices.
	UseCogsAcctForOciFlag             string    `bun:"use_cogs_acct_for_oci_flag,type:char(1),default:('N')"`          // Indicates that during Prod Order Proc, we post to cogs acct for other charge items
	RestockingFeeId                   string    `bun:"restocking_fee_id,type:varchar(40),nullzero"`                    // Custom column which is an other charge item setup in Item Maintenance window for restokcing fee
	UseQuoteRevisionFlag              string    `bun:"use_quote_revision_flag,type:char(1),default:('N')"`             // Custom column to enable quote revision functionality at the company level
	SummaryCommodityCode              string    `bun:"summary_commodity_code,type:varchar(255),nullzero"`              // Summary Commodity Code for Level III Processing
	CompanyTaxgroupByZipcodeFlag      string    `bun:"company_taxgroup_by_zipcode_flag,type:char(1),default:('N')"`    // Setting to determine if a company should use the tax group by zipcode functionality
	ServiceUserNumber                 string    `bun:"service_user_number,type:varchar(6),nullzero"`                   // We will be using this field in the header record of the file we send to the bank
	CfnGasCogsAcct                    string    `bun:"cfn_gas_cogs_acct,type:varchar(32),nullzero"`                    // column that stored a CFN Cost of Goods Sold Account
	CfnReceivableAcct                 string    `bun:"cfn_receivable_acct,type:varchar(32),nullzero"`                  // column that sored a CFN Receivable Account
	CfnGasRevenueAcct                 string    `bun:"cfn_gas_revenue_acct,type:varchar(32),nullzero"`                 // column that stored CFN Revenue Account
	RoundUnitPricesFlag               string    `bun:"round_unit_prices_flag,type:char(1),default:('N')"`              // Round Unit Prices to 2 decimals flag
	AvalaraEnabled                    string    `bun:"avalara_enabled,type:char(1),default:('N')"`                     // Flag to indicate whether this company is using avalara tax integration.
	IvaEnabled                        string    `bun:"iva_enabled,type:varchar(1),default:('N')"`                      // Whether the Company will use IVA (Mexico) functionality.
	TrackInvOnWaterBranchFlag         string    `bun:"track_inv_on_water_branch_flag,type:char(1),default:('Y')"`      // Track Inventory on Water by Branchf flag
	FreightWipAcct                    string    `bun:"freight_wip_acct,type:varchar(32),nullzero"`                     // G/L Account for Freight WIP
	PendingWarrantyAcct               string    `bun:"pending_warranty_acct,type:varchar(32),nullzero"`                // Account to be used when the Credit memo postings are made.
	DefaultCustomerTaxClass           string    `bun:"default_customer_tax_class,type:varchar(255),nullzero"`          // Related to third party tax features - indicates the customer tax class that should be used when creating a new ship to.
	CountrySpecificCd                 int32     `bun:"country_specific_cd,type:int,nullzero"`                          // Country Specific Functionality Code
	BorrowingRateToCarry              float64   `bun:"borrowing_rate_to_carry,type:decimal(19,4),default:((0))"`       // Rate to borrow funds to cover a cash shortfall
	FillOrKillReasonUid               int32     `bun:"fill_or_kill_reason_uid,type:int,nullzero"`                      // Lost Sales UID used when any order or order line is cancelled during import because of 'fill or kill' logic.
	FillAndKillReasonUid              int32     `bun:"fill_and_kill_reason_uid,type:int,nullzero"`                     // Lost Sales UID used when any order line is cancelled on an import when 'fill or kill' logic was not used.
	InvoiceOnlyLocationId             float64   `bun:"invoice_only_location_id,type:decimal(19,0),nullzero"`
	ExternalTaxCompanyCode            string    `bun:"external_tax_company_code,type:varchar(255),nullzero"` // For use with 3rd party tax integrations; the company code that should be used for tax requests when the P21 company ID can not be used directly.
	ArToApClearingAcct                string    `bun:"ar_to_ap_clearing_acct,type:varchar(32),nullzero"`
	CompanyUid                        int32     `bun:"company_uid,type:int,autoincrement,unique"` // Unique identifier for the table
	GlForeignCardlockAcct             string    `bun:"gl_foreign_cardlock_acct,type:varchar(32),nullzero"`
	BinTypeUid                        int32     `bun:"bin_type_uid,type:int,nullzero"`                                     // Custom column to indicate the bin type for automatically create Inventory Return.
	PriceRoundingFlag                 string    `bun:"price_rounding_flag,type:char(1),nullzero"`                          // Indicates if this company rounds prices.
	OverUnderRoundingAcct             string    `bun:"over_under_rounding_acct,type:varchar(255),nullzero"`                // The gl account used when rounding creates a discrepency between cash received and the actual invoice amount.
	FutureCostDefaultFromLocId        float64   `bun:"future_cost_default_from_loc_id,type:decimal(19,0),nullzero"`        // The location ID from which data will be retrieved for location-specific values (such a Standard Cost) for this company
	EditTrackingNumberFlag            string    `bun:"edit_tracking_number_flag,type:varchar(1),default:('N')"`            // Indicator to use the functionality to add multiple tracking no in OE and FC
	RebateCogsAccountNo               string    `bun:"rebate_cogs_account_no,type:varchar(32),nullzero"`                   // Custom account column for specifying a Rebate Cost Of Goods Sold Account
	FillableReasonId                  string    `bun:"fillable_reason_id,type:varchar(5),nullzero"`                        // Reason for fillable container adjustment
	CfnDieselCogsAcct                 string    `bun:"cfn_diesel_cogs_acct,type:varchar(32),nullzero"`                     // CFN Diesel Cost of Goods Sold Account
	CfnDieselRevenueAcct              string    `bun:"cfn_diesel_revenue_acct,type:varchar(32),nullzero"`                  // CFN Diesel Reveneue Account
	BusinessDaysPerYear               int16     `bun:"business_days_per_year,type:smallint,default:((250))"`               // Number of business days in a calendar year
	AveragePretaxProfit               float64   `bun:"average_pretax_profit,type:decimal(19,4),default:((0.075))"`         // Average profit before taxes, defaults to industry avg
	SbVendorCreditOffsetAcct          string    `bun:"sb_vendor_credit_offset_acct,type:varchar(32),nullzero"`             // ServiceBench Vendor Credit Offset Account for Account Receivable
	EnableMultiCurrencyPayments       string    `bun:"enable_multi_currency_payments,type:char(1),nullzero"`               // Indicates that this company allows payments to be taken in multiple currencies and enables the display currency functionality as well.
	CurrencyClearingAccountNo         string    `bun:"currency_clearing_account_no,type:varchar(255),nullzero"`            // Custom: Indicates the gl clearing account used for multi-currency transactions.
	BrokerEmailAddress                string    `bun:"broker_email_address,type:varchar(255),nullzero"`                    // Indicated the address to email the Pro Forma invoice information
	WarrantyRmaReceiptClearingAcct    string    `bun:"warranty_rma_receipt_clearing_acct,type:varchar(32),nullzero"`       // GL account that will be used for Inventory Return transactions.
	WarrInvReceiptsClearingAcct       string    `bun:"warr_inv_receipts_clearing_acct,type:varchar(32),nullzero"`          // Custom column for Warranty Inventory Return Inventory Receipts Clearing Account.
	TrafficManagerContactId           string    `bun:"traffic_manager_contact_id,type:varchar(16),nullzero"`               // Traffic Manager
	ProFormaSignaturePath             string    `bun:"pro_forma_signature_path,type:varchar(255),nullzero"`                // Pro Forma Signature
	ProFormaSignatureName             string    `bun:"pro_forma_signature_name,type:varchar(255),nullzero"`                // Pro Forma Name
	CostAndRevenueBasedOnSalesrepFlag string    `bun:"cost_and_revenue_based_on_salesrep_flag,type:char(1),default:('N')"` // Option to post cost and revenue based on sales rep affiliation
	TaCylinderRentalRevAcct           string    `bun:"ta_cylinder_rental_rev_acct,type:varchar(32),nullzero"`              // TrackAbout Cylinder Rental Revenue Account
	TaEquipmentRentalRevAcct          string    `bun:"ta_equipment_rental_rev_acct,type:varchar(32),nullzero"`             // TrackAbout Equipment Rental Revenue Account
	RemnantReasonId                   string    `bun:"remnant_reason_id,type:varchar(5),nullzero"`                         // Remnant Reason Id
	RemnantSourceCostCd               int32     `bun:"remnant_source_cost_cd,type:int,default:((3617))"`                   // Remnant Cost Source
	UseXfersrclocRevacctLcFlag        string    `bun:"use_xfersrcloc_revacct_lc_flag,type:char(1),nullzero"`               // Custom column to indicate if company will use Transfer Source Location Revenue Account for Landed Cost Amount.
	XfersrclocRevenueAcctForLc        string    `bun:"xfersrcloc_revenue_acct_for_lc,type:varchar(32),nullzero"`           // Custom column for transfer source location revenue account for landed cost amount
	UseDimensions                     string    `bun:"use_dimensions,type:char(1),default:('N')"`                          // Whether company wants to track various GL dimensions
	ProdOrderWipTrackingCd            int32     `bun:"prod_order_wip_tracking_cd,type:int,default:((3752))"`               // Determines when the Production Order WIP is recognized (3752 - Allocated Material, 3753 - Confirmed Pick Tickets)
	InventoryWipAccountNo             string    `bun:"inventory_wip_account_no,type:varchar(32),nullzero"`                 // This will be an account that will be used to post WIP for the material in the Production Order WIP Worksheet
	OtherChargeWipAccountNo           string    `bun:"other_charge_wip_account_no,type:varchar(32),nullzero"`              // This will be an account that will be used to post WIP for Other Charge Items in the Production Order WIP Worksheet
	ZoneCostExchangeAcct              string    `bun:"zone_cost_exchange_acct,type:varchar(32),nullzero"`                  // The gl account used for the difference between the two zone costs at destination location of the transfer. Custom.
	ZoneCostVarianceAcct              string    `bun:"zone_cost_variance_acct,type:varchar(32),nullzero"`                  // The gl account used for the difference between the two zone costs at the source location of the transfer. Custom.
	ShowVatPerInvLineFlag             string    `bun:"show_vat_per_inv_line_flag,type:char(1),default:('N')"`              // Flag to enable/disable VAT per invoice line functionality
	RetailDeliveryFeePayableAcct      string    `bun:"retail_delivery_fee_payable_acct,type:varchar(32),nullzero"`         // Payable Account No used for Avalara Colorado Retail Delivery Fee.
}
