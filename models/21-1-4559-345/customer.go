package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Customer struct {
	bun.BaseModel                        `bun:"table:customer"`
	CustomerId                           float64    `bun:"customer_id,type:decimal(19,0),pk"`                             // System-generated ID that identifies customers.
	CompanyId                            string     `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	TaxableFlag                          string     `bun:"taxable_flag,type:char(1)"`                                     // Indicates whether or not sales tax is charged for the customer.
	ResaleCertificate                    *string    `bun:"resale_certificate,type:varchar(40)"`                           // ID issued to companies exempt from state sales tax.
	CreditLimit                          *float64   `bun:"credit_limit,type:decimal(22,4)"`                               // Maximum allowed amount of a customers outstanding debt at any given time.
	CreditLimitUsed                      *float64   `bun:"credit_limit_used,type:decimal(22,4)"`                          // System-generated value indicating how much a customer has currently applied to their credit limit.
	TermsId                              string     `bun:"terms_id,type:varchar(2)"`                                      // A system-generated number that identifies a set of payment terms.
	SalesrepId                           *string    `bun:"salesrep_id,type:varchar(16)"`                                  // The sales representative for this customer.
	ArAccountNo                          *string    `bun:"ar_account_no,type:varchar(32)"`                                // Account number representing an Accounts Receivable account.
	RevenueAccountNo                     *string    `bun:"revenue_account_no,type:varchar(32)"`                           // Account that shows the gross increase in your companys income.
	CosAccountNo                         *string    `bun:"cos_account_no,type:varchar(32)"`                               // Cost of Goods Sold account that tracks the cost of products sold.
	AllowedAccountNo                     *string    `bun:"allowed_account_no,type:varchar(32)"`                           // Account consisting of excess amounts that your company writes off from an invoice.
	TermsAccountNo                       *string    `bun:"terms_account_no,type:varchar(32)"`                             // Account consisting of discounts taken against invoices based upon the terms agreement.
	DeleteFlag                           string     `bun:"delete_flag,type:char(1),default:('N')"`                        // Indicates whether this record is logically deleted
	DateCreated                          time.Time  `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified                     time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy                     string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CustomerIdString                     string     `bun:"customer_id_string,type:varchar(22)"`                           // This is a string representation of the numeric customer id value. This exists only for performance reasons.
	FreightAccountNo                     *string    `bun:"freight_account_no,type:varchar(32)"`                           // Account that posts the freight amount when an invoice is created for this customer.
	BrokerageAccountNo                   *string    `bun:"brokerage_account_no,type:varchar(32)"`                         // Account that keeps track of your brokerage costs.
	Class1Id                             *string    `bun:"class_1id,type:varchar(255)"`                                   // A user-defined code that identifies a group of customers.
	Class2Id                             *string    `bun:"class_2id,type:varchar(255)"`                                   // A user-defined code that identifies a group of customers.
	Class3Id                             *string    `bun:"class_3id,type:varchar(255)"`                                   // A user-defined code that identifies a group of customers.
	Class4Id                             *string    `bun:"class_4id,type:varchar(255)"`                                   // A user-defined code that identifies a group of customers.
	Class5Id                             *string    `bun:"class_5id,type:varchar(255)"`                                   // A user-defined code that identifies a group of customers.
	TradePercentDisc                     *float64   `bun:"trade_percent_disc,type:decimal(22,4)"`                         // This column is unused.
	AcceptPartialOrders                  *string    `bun:"accept_partial_orders,type:char(1)"`                            // Can this customer accept partial orders?
	AcceptableWaitTime                   *float64   `bun:"acceptable_wait_time,type:decimal(22,4)"`                       // Maximum number of days a customer agrees to wait for ordered material.
	PriceFileId                          *float64   `bun:"price_file_id,type:decimal(22,4)"`                              // This column is unused.
	EdiOrPaper                           *string    `bun:"edi_or_paper,type:char(1)"`                                     // Does this customer use EDI or paper?
	SecurityInfo                         *string    `bun:"security_info,type:varchar(10)"`                                // This column is unused.
	InterchgReceiverId                   *string    `bun:"interchg_receiver_id,type:varchar(15)"`                         // Standard Address Number assigned to a contact if they use EDI to transmit or receive documents.
	IntlSan                              *string    `bun:"intl_san,type:varchar(15)"`                                     // Number assigned by the Accredited Standards Committee X12 to each trading partner outside of the US who is transmitting/receiving EDI documents.
	CreditCardNo                         *string    `bun:"credit_card_no,type:varchar(24)"`                               // Numeric code that appears on the customers credit card.
	CreditCardType                       *string    `bun:"credit_card_type,type:varchar(10)"`                             // The type of credit card the customer is using.
	CreditCardName                       *string    `bun:"credit_card_name,type:varchar(41)"`                             // The name of a person or organization, as it appears on the credit card.
	CreditCardExpirationDate             *time.Time `bun:"credit_card_expiration_date,type:datetime"`                     // The month and year a credit card expires.
	LastCheckNumber                      *string    `bun:"last_check_number,type:varchar(40)"`                            // Last check number received from the customer.
	LastCheckAmount                      *float64   `bun:"last_check_amount,type:decimal(22,4)"`                          // Last check amount received from the customer.
	LastCheckDate                        *time.Time `bun:"last_check_date,type:datetime"`                                 // Date of the last check received from the customer.
	ReceivableGroupId                    *string    `bun:"receivable_group_id,type:varchar(8)"`                           // Reserved for future development.
	CreditLimitPerOrder                  *float64   `bun:"credit_limit_per_order,type:decimal(19,4)"`                     // Maximum allowed amount a customer can spend for each order.
	OpenItemBalanceForward               *string    `bun:"open_item_balance_forward,type:char(1)"`                        // Indicates the type of statement the customer receives.
	OrderAcknowledgments                 string     `bun:"order_acknowledgments,type:char(1),default:('Y')"`              // Indicates whether the customer requires order acknowledgements.
	AcceptInterchangeableItems           *string    `bun:"accept_interchangeable_items,type:char(1)"`                     // Shows whether a customer agrees to accept interchangeable items if the items they order are out of stock.
	BillToContactId                      *string    `bun:"bill_to_contact_id,type:varchar(16)"`                           // The recipient of invoices generated for a specific customer.
	GenerateFinanceCharges               *string    `bun:"generate_finance_charges,type:char(1)"`                         // Indicate if you would like the system to calculate
	LimitMaxShipmentsPerOrder            *float64   `bun:"limit_max_shipments_per_order,type:decimal(19,4)"`              // Maximum number of shipments a customer wants to receive per order.
	CreditLimitCheckAtShipment           *string    `bun:"credit_limit_check_at_shipment,type:char(1)"`                   // Indicates whether the system checks a customers credit when ordered material is shipped.
	GenerateCustomerStatements           string     `bun:"generate_customer_statements,type:char(1),default:('Y')"`       // Indicates whether statements can be generated for a particular customer.
	SicCode                              *float64   `bun:"sic_code,type:decimal(6,0)"`                                    // Industry standard four-digit code that identifies a companys industry.
	MinimumOrderDollarAmount             *float64   `bun:"minimum_order_dollar_amount,type:decimal(19,4)"`                // Indicates the customers minimum required order amount.
	FederalExemptionNumber               *string    `bun:"federal_exemption_number,type:varchar(40)"`                     // ID issued to companies exempt from federal excise tax.
	OtherExemptionNumber                 *string    `bun:"other_exemption_number,type:varchar(40)"`                       // Tax exemption number other than one for State Sales Tax, State Excise Tax, or Federal Excise Tax.
	FloorPlanAccount                     string     `bun:"floor_plan_account,type:char(1)"`                               // Account that is used for Floor Plan.
	DefaultDisposition                   string     `bun:"default_disposition,type:char(1)"`                              // The disposition setting automatically assigned to items ordered by this customer when the order cannot be filled.
	FcPercentage                         float64    `bun:"fc_percentage,type:decimal(19,4),default:(0)"`                  // The finance charge percentage to be charged.
	FcGraceDays                          float64    `bun:"fc_grace_days,type:decimal(19,0),default:(0)"`                  // Enter the grace days past the net due date before finance charges are calculated.
	MinimumFinanceCharge                 float64    `bun:"minimum_finance_charge,type:decimal(19,4),default:(0)"`         // Indicates the minimum amount of finanace charge.
	FcCycle                              *string    `bun:"fc_cycle,type:char(1)"`                                         // The frequency in which you charge a customer penalties for lack of payment (i.e., monthly).
	LastFcDate                           *time.Time `bun:"last_fc_date,type:datetime"`                                    // Indicates the date when finance charges for this customer were last calculated.
	StatementFrequencyId                 *string    `bun:"statement_frequency_id,type:varchar(8)"`                        // This column is unused.
	HighestCreditLimitUsed               *float64   `bun:"highest_credit_limit_used,type:decimal(19,4)"`                  // Displays the most credit a customer has ever used against their set credit limit.
	BilledOnGrossNetQty                  *string    `bun:"billed_on_gross_net_qty,type:char(1)"`                          // Reserved for future development.
	StateExciseTaxExemptionNo            *string    `bun:"state_excise_tax_exemption_no,type:varchar(40)"`                // ID issued to companies exempt from state excise tax.
	CreditStatus                         *string    `bun:"credit_status,type:varchar(8)"`                                 // Customers rating on how well they pay their invoices.
	DeferredRevenueAccountNo             *string    `bun:"deferred_revenue_account_no,type:varchar(32)"`                  // The default account number to post deferred revenue for a new customer
	OverrideRevenueByItem                *string    `bun:"override_revenue_by_item,type:char(1)"`                         // Indicates whether revenue is tracked by item for this customer.
	PickTicketType                       *string    `bun:"pick_ticket_type,type:varchar(2)"`                              // Indicates what type of pick ticket prints for a specific ship to.
	FinanceChgRevenueAccountNo           *string    `bun:"finance_chg_revenue_account_no,type:varchar(32)"`               // Account that shows the gross increase in your companys income from finance charges.
	FinanceChargeShipToId                *float64   `bun:"finance_charge_ship_to_id,type:decimal(19,0)"`                  // The ship-to address that will display on finance charge invoices.
	CustomerName                         *string    `bun:"customer_name,type:varchar(255)"`                               // Description of customer ID.
	UseConsolidatedInvoicing             string     `bun:"use_consolidated_invoicing,type:char(1),default:('N')"`         // Indicates whether this customer uses consolidated invoicing.
	CiForCompleteOrdersOnly              string     `bun:"ci_for_complete_orders_only,type:char(1),default:('N')"`        // Used to track -  at the customer level -  information about consolidated invoicing -this column will determine if we should only consolidate invoices when Orders are marked complete.
	CiPrintDetail                        string     `bun:"ci_print_detail,type:char(1),default:('N')"`                    // Used to track -  at the customer level -  information about consolidated invoicing - this column will determine if the Invoices should default to print with detail or no detail.
	InvoicePrintQty                      float64    `bun:"invoice_print_qty,type:decimal(5,0),default:(0)"`               // Indicates how many copies of an invoice should print at one time for this customer.
	FobRequiredFlag                      string     `bun:"fob_required_flag,type:char(1),default:('Y')"`                  // Does the customer pay shipping charges?
	CodRequiredFlag                      string     `bun:"cod_required_flag,type:char(1),default:('N')"`                  // Indicates whether this customer requires Cash On Delivery.
	PricingMethodCd                      *int16     `bun:"pricing_method_cd,type:tinyint"`                                // Indicates what price the program will use when calculating order totals.
	SourcePriceCd                        *int32     `bun:"source_price_cd,type:int"`                                      // A price upon which a sales structure is based.
	Multiplier                           *float64   `bun:"multiplier,type:decimal(19,9)"`                                 // A discount factor that is muliplied by Source Price.
	PoNoRequired                         string     `bun:"po_no_required,type:char(1),default:('N')"`                     // Indicates whether a PO number is required in Order Entry before saving the order.
	ConsolidatedShipToId                 *float64   `bun:"consolidated_ship_to_id,type:decimal(19,0)"`                    // Indicates ship to ID for consolidated invoices.
	InvoiceBatchUid                      int32      `bun:"invoice_batch_uid,type:int,default:(1)"`                        // Unique ID relating to a four-digit code the system uses to identify invoice batches.
	CustomerStatementHistoryUid          *int32     `bun:"customer_statement_history_uid,type:int"`                       // Unique ID for customer statements.
	TradingPartnerName                   *string    `bun:"trading_partner_name,type:varchar(255)"`                        // Unique code for each supplier and customer with whom you do business via EDI.
	DefaultRebateLocationId              *float64   `bun:"default_rebate_location_id,type:decimal(19,0)"`                 // Location used for rebate tracking.
	LotBillSummaryOnInvoices             string     `bun:"lot_bill_summary_on_invoices,type:char(1),default:('Y')"`       // Indicates whether lot IDs will appear on invoices.
	PrintZeroDollarCustomers             string     `bun:"print_zero_dollar_customers,type:char(1),default:('N')"`        // This column is unused.
	GenerateStatementsBy                 int16      `bun:"generate_statements_by,type:smallint,default:(1005)"`           // Indicates if the customer generates customer statements using the customer ID or the corporate ID.
	PendingPaymentAccountNo              *string    `bun:"pending_payment_account_no,type:varchar(32)"`                   // Account used to track disputed zero dollar invoices.
	PrintPricesOnPackinglist             string     `bun:"print_prices_on_packinglist,type:char(1),default:('N')"`        // Indicates whether prices print on packing lists.
	JobPricing                           string     `bun:"job_pricing,type:char(1),default:('N')"`
	RecordType                           *int16     `bun:"record_type,type:smallint"`                                    // This column indicates whether the record is for a Customer, Prospect or for some other type.
	RecordSource                         *string    `bun:"record_source,type:varchar(255)"`                              // Indicates the source of this record
	SlxContactid                         *string    `bun:"slx_contactid,type:varchar(255)"`                              // Sales Logix contact id.
	StatementBatchUid                    int32      `bun:"statement_batch_uid,type:int,default:(1)"`                     // Allows Customer Statements to be generated in batches.
	AllowAdvanceBilling                  *string    `bun:"allow_advance_billing,type:char(1),default:('N')"`             // Allow advance billing
	AdvanceBillAccountNo                 *string    `bun:"advance_bill_account_no,type:varchar(32)"`                     // Advance Bill Account No
	IncludeNonAllocOnPickTix             int32      `bun:"include_non_alloc_on_pick_tix,type:int,default:(1277)"`        // Include non-allocated items on Pick Tickets
	ExcludeCanceldFromPickTix            string     `bun:"exclude_canceld_from_pick_tix,type:varchar(1),default:('N')"`  // Exclude items with C disposition from Pick Tickets
	ExcludeHoldFromPickTix               string     `bun:"exclude_hold_from_pick_tix,type:varchar(1),default:('N')"`     // Exclude items with H disposition from Pick Tickets
	ExcludeCanceldFromPackList           string     `bun:"exclude_canceld_from_pack_list,type:varchar(1),default:('N')"` // Exclude items with C disposition from Packing Lists
	ExcludeHoldFromPackList              string     `bun:"exclude_hold_from_pack_list,type:varchar(1),default:('N')"`    // Exclude items with H disposition from Packing Lists
	ExcludeCanceldFromOrderAck           string     `bun:"exclude_canceld_from_order_ack,type:varchar(1),default:('N')"` // Exclude items with C disposition from Order Acknowledgements
	ExcludeHoldFromOrderAck              string     `bun:"exclude_hold_from_order_ack,type:varchar(1),default:('N')"`    // Exclude items with H disposition from Order Acknowledgements
	DispAddlInfoOnInvcFlag               *string    `bun:"disp_addl_info_on_invc_flag,type:char(1)"`                     // Determines if additional info should be shown in the invoice datastream
	RemitToAddressId                     *float64   `bun:"remit_to_address_id,type:decimal(19,0)"`                       // Remit To Address ID
	DefaultBranchId                      *string    `bun:"default_branch_id,type:varchar(8)"`                            // Branch with which the customer is associated
	CustPartNoGroupHdrUid                *int32     `bun:"cust_part_no_group_hdr_uid,type:int,default:(0)"`              // Customer Group UID
	CurrencyId                           float64    `bun:"currency_id,type:decimal(19,0)"`                               // Currency associated with the customer.
	OverrideProfitLimit                  string     `bun:"override_profit_limit,type:char(1),default:('N')"`             // Indicates whether the customer profit settings should be used.
	MinimumOrderLineProfit               *float64   `bun:"minimum_order_line_profit,type:decimal(19,2)"`                 // The minimum desired profit percentage for an order line item.
	MaximumOrderLineProfit               *float64   `bun:"maximum_order_line_profit,type:decimal(19,2)"`                 // The maximum desired profit percentage for an order line item.
	MinimumOrderProfit                   *float64   `bun:"minimum_order_profit,type:decimal(19,2)"`                      // The minimum desired profit percentage for an order.
	MaximumOrderProfit                   *float64   `bun:"maximum_order_profit,type:decimal(19,2)"`                      // The maximum desired profit percentage for an order.
	DateAcctOpened                       *time.Time `bun:"date_acct_opened,type:datetime"`                               // Date the customer account became active.
	CreatedBy                            *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	IncludeNonAllocOnPackList            int32      `bun:"include_non_alloc_on_pack_list,type:int"`                    // When to include non-allocated items on the packing list
	PrintPackinglistInShipping           string     `bun:"print_packinglist_in_shipping,type:char(1),default:('N')"`   // Whether to print the packing list
	SourceTypeCd                         *int32     `bun:"source_type_cd,type:int"`                                    // Where the customer was made.
	PricingMethodCdWeb                   *int32     `bun:"pricing_method_cd_web,type:int"`                             // Indicates what method the program will use when calculating web based order totals.
	SourcePriceCdWeb                     *int32     `bun:"source_price_cd_web,type:int"`                               // The price column, used in conjunction w/the multiplier, used to calculate the order total for a web based order.
	MultiplierWeb                        *float64   `bun:"multiplier_web,type:decimal(19,9)"`                          // The multiplier, used in conjunction w/the source_price_cd_web, used to calculate the order total for a web based order.
	DefaultOrdersToWillCall              string     `bun:"default_orders_to_will_call,type:char(1),default:('N')"`     // When Y-es, specifies that the order will call flag will default to Y-es
	SendUcc128info                       *string    `bun:"send_ucc128info,type:char(1)"`                               // Whether to send UCC - 128 Information with 856 transaction
	MfrNo                                *string    `bun:"mfr_no,type:varchar(255)"`                                   // Manufacturer Number used to generate the unique UCC - 128 barcode for each label
	AlwaysUseJobPrice                    string     `bun:"always_use_job_price,type:char(1),default:('N')"`            // Flag to indicate use always use Job/Contract pricing
	AllowNonJobItem                      string     `bun:"allow_non_job_item,type:char(1),default:('N')"`              // If selected (by Checkbox) then user may add items not specified on the Job/Contract to the sales order.
	PromptForNonJobItem                  string     `bun:"prompt_for_non_job_item,type:char(1),default:('N')"`         // If selected, the user will be prompted when item added to an order that is not defined in Job/Contract
	AllowExceedJobQty                    string     `bun:"allow_exceed_job_qty,type:char(1),default:('N')"`            // If selected, will allow the item to continue being priced at the defined pricing even when it exceeds the defined quantity.
	PrintLotAttribOnInvoice              string     `bun:"print_lot_attrib_on_invoice,type:char(1),default:('N')"`     // indicate if should print determine lot attribute on invoice for this customer
	PrintLotAttribOnPacklist             string     `bun:"print_lot_attrib_on_packlist,type:char(1),default:('N')"`    // indicate if should print determine lot attribute on packing list
	OverTolerancePercentage              *int32     `bun:"over_tolerance_percentage,type:int"`                         // Percentage of over shipment a customer is willing to take.
	UnderTolerancePercentage             *int32     `bun:"under_tolerance_percentage,type:int"`                        // Percentage of uder shipment a customer is willing to take.
	CustomerTypeCd                       int32      `bun:"customer_type_cd,type:int,default:(1203)"`                   // Customer type identifier
	LeadSourceId                         *string    `bun:"lead_source_id,type:varchar(8)"`                             // Default lead source id value
	DataIdentifierGroupUid               *int32     `bun:"data_identifier_group_uid,type:int"`                         // Data Identifier Group UID for this customer (typically used in the generation of sales-order related documents requiring data identifiers)
	PassportCustomerId                   *string    `bun:"passport_customer_id,type:varchar(255)"`                     // Allen Bradley assigned value employed during an EDI 844 export.
	CostCenterTrackingOption             *int32     `bun:"cost_center_tracking_option,type:int"`                       // Determines the cost center tracking options. Values from code table.
	FascorWmsPricingOption               int32      `bun:"fascor_wms_pricing_option,type:int,default:(200)"`           // Specifies the value to be passed in the sku_retail_price field when pick tickets are exported to the FASCOR WMS app.
	AllowItemLevelContractFlag           string     `bun:"allow_item_level_contract_flag,type:char(1),default:('N')"`  // Allow user to select contracts on the item level in order entry
	ShipToCreditLimit                    float64    `bun:"ship_to_credit_limit,type:decimal(19,9),default:(0)"`        // This will hold the default credit limit for new/otf ship tos
	LegacyId                             *string    `bun:"legacy_id,type:varchar(255)"`                                // column to hold customer ID from legacy app
	ParkerCustomerCd                     *string    `bun:"parker_customer_cd,type:varchar(10)"`                        // Field needed for Parker Rebate reporting requirements.
	WebEnabledFlag                       string     `bun:"web_enabled_flag,type:char(1),default:('N')"`                // Column that indicates whether customer is in web DB
	AllowLineItemFreightFlag             string     `bun:"allow_line_item_freight_flag,type:char(1),default:('N')"`    // Indicates if this customer allows freight on a per item level
	FreightChargeUid                     *int32     `bun:"freight_charge_uid,type:int"`                                // Freight charge associated with this customer
	SendDsc856Flag                       *string    `bun:"send_dsc856_flag,type:char(1),default:('N')"`                // Flag indicating whether to send 856 upon Direct Ship Confirmation
	PassThroughOptionFlag                *string    `bun:"pass_through_option_flag,type:char(1)"`                      // This flag will be used to determine if this customer will use pass through functionality
	EnableBudgetCodesFlag                *string    `bun:"enable_budget_codes_flag,type:char(1)"`                      // Custom (F23038): determines if the customer employs job pricing budget codes
	EnableAggregateBudgetsFlag           *string    `bun:"enable_aggregate_budgets_flag,type:char(1)"`                 // Custom (F23038): determines if the customer employs aggregate budgets
	MaxOrdersPerBudget                   *int32     `bun:"max_orders_per_budget,type:int"`                             // Custom (F23038): determines functionality when the max numbers of order for a budget has been exceeded.
	ServiceTermsId                       *string    `bun:"service_terms_id,type:varchar(2)"`                           // Terms for service orders
	DownpaymentPercentage                float64    `bun:"downpayment_percentage,type:decimal(19,2),default:((0.00))"` // Percentage of amount ordered that will be requested as downpayment.
	ReqPymtUponReleaseOfItems            string     `bun:"req_pymt_upon_release_of_items,type:char(1),default:('N')"`  // Indicates whether full pymt is required when ordered items are invoiced.
	IncludeDpSummaryOnInvoices           string     `bun:"include_dp_summary_on_invoices,type:char(1),default:('N')"`  // Indicates whether printed invoices should include a downpayment summary.
	JobNumberRequiredFlag                string     `bun:"job_number_required_flag,type:char(1),default:('N')"`        // Indicates whether orders and RMAs entered for this customer require a job number.
	UseLastMarginPricingFlag             string     `bun:"use_last_margin_pricing_flag,type:char(1),default:('N')"`    // Indicates if the customer using last margin pricing
	InvoiceCompCostCdTier1               int32      `bun:"invoice_comp_cost_cd_tier1,type:int,default:((2110))"`       // First cost attempted to be used to compare invoice profit vs order profit
	DistrictId                           *float64   `bun:"district_id,type:decimal(19,0)"`                             // Allows Customers to be assigned to a District ID.
	PreferredFlag                        string     `bun:"preferred_flag,type:char(1),default:('N')"`                  // Indicates a Preferred customer
	ApplyFuelSurchargesCd                *int32     `bun:"apply_fuel_surcharges_cd,type:int"`                          // A custom column contains code of Yes, No, or Rep.  If its Yes, surcharges will be calculated based on info in class maint.  If its No, no surcharge will be applied.  If Rep, a surcharge will be calculated based on Sales Reps info.
	AppliedFuelchargesToDsFlag           *string    `bun:"applied_fuelcharges_to_ds_flag,type:char(1)"`                // A custom column which indicates whether a surcharge will be calculated and applied to invoices via Direct Ship Confirmation.
	SuppressZeroDollarFlag               *string    `bun:"suppress_zero_dollar_flag,type:char(1)"`                     // A custom column which indicates whether the system will suppress zero dollar fuel surcharge lines from appearing on any forms.
	InvoiceCompCostCdTier2               int32      `bun:"invoice_comp_cost_cd_tier2,type:int,default:((2111))"`       // Second cost attempted to be used to compare invoice profit vs order profit
	InvoiceCompCostCdTier3               int32      `bun:"invoice_comp_cost_cd_tier3,type:int,default:((1195))"`       // Third cost attempted to be used to compare invoice profit vs order profit
	SignatureRequiredFlag                *string    `bun:"signature_required_flag,type:char(1),default:('N')"`         // Indicates if the signature is required for the customer.
	PromiseDateBuffer                    *int32     `bun:"promise_date_buffer,type:int"`                               // Number of buffer days added to the promise date
	DaysUntilQuoteExpires                *int32     `bun:"days_until_quote_expires,type:int"`                          // Number of days until a new quote expires
	ReclaimDiscountOnMemosFlag           *string    `bun:"reclaim_discount_on_memos_flag,type:char(1)"`                // Indicates whether for RMAs the system should reduce the amount of the customers credit based on terms offered on the linked invoice.
	OrderSurchargeUid                    *int32     `bun:"order_surcharge_uid,type:int"`                               // foreign key to table order_surcharge
	SplitCoresFlag                       *string    `bun:"split_cores_flag,type:char(1)"`                              // This custom column indicates the core elements to be printed on a separate invoice.
	SalesrepAssignedDate                 *time.Time `bun:"salesrep_assigned_date,type:datetime"`                       // The date that the sales rep was assigned to the customer
	ClockCellTrackingFlag                *string    `bun:"clock_cell_tracking_flag,type:char(1),default:('N')"`        // Enable or disable tracking clock/cell
	EnvironmentalFeeFlag                 string     `bun:"environmental_fee_flag,type:char(1),default:('N')"`          // Indicates if the customer will be charged an environmental fee.
	AdminFeeFlag                         string     `bun:"admin_fee_flag,type:char(1),default:('N')"`                  // Indicates if this customer should be charged a handling/admin fee
	EpaCertOnFileFlag                    *string    `bun:"epa_cert_on_file_flag,type:char(1),default:('N')"`           // Indicates whether the customer has an EPA certification record and can buy EPA certification required tems
	OrderDiscType                        *int32     `bun:"order_disc_type,type:int"`                                   // Custom (F33700): order discount type. Possible values are 230 (percentage) and 714 (dollar amount).
	OrderDiscFactor                      *float64   `bun:"order_disc_factor,type:decimal(19,9)"`                       // Custom (F33700): order discount factor. Will be a dollar amount or percentage based on the order_disc_type column.
	UseVendorContractsFlag               *string    `bun:"use_vendor_contracts_flag,type:char(1),default:('N')"`       // The column is to indicate whether a vendor contract is used to determine a cost/rebate
	PedigreeCustomer                     string     `bun:"pedigree_customer,type:char(1),default:('N')"`               // Indicates customer interested in pedigree reports
	NationalAccountFlag                  *string    `bun:"national_account_flag,type:char(1)"`                         // Determines whether this customer is a national account or not.
	FreightMarkupMultiplier              *float64   `bun:"freight_markup_multiplier,type:decimal(19,9)"`               // Custom Column to indicate the multiplier used in freight calculation
	OrderPriorityUid                     *int32     `bun:"order_priority_uid,type:int"`
	ManufacturerDistributorNo            *string    `bun:"manufacturer_distributor_no,type:varchar(255)"`                // Manufacturer assigned customer number for processing rebates.
	ManufacturerProgramTypePct           *float64   `bun:"manufacturer_program_type_pct,type:decimal(19,2)"`             // Percentage of the manufacturer program rebate type.
	ManufacturerRebateLoc                *string    `bun:"manufacturer_rebate_loc,type:varchar(255)"`                    // Manufacturere assigned rebate location for this customer.
	FloorPlanTaxableFlag                 *string    `bun:"floor_plan_taxable_flag,type:char(1)"`                         // If this is a floor plan customer, this column indicates whether floor plans are taxable for them. Values Y - Yes, N - No
	UseVendorItemTermsFlag               *string    `bun:"use_vendor_item_terms_flag,type:char(1)"`                      // If 'Y' and system setting line_item_terms is 'customer terms or vendor\items', line terms will use vendor\item terms. If 'N', use customer terms always for this customer.
	RentalUpdateFlag                     *string    `bun:"rental_update_flag,type:char(1)"`                              // Determines if this customer has been sent to an external rental system.
	ForeignCurrencyGuaranteeType         *string    `bun:"foreign_currency_guarantee_type,type:char(1)"`                 // This column overides the system setting fc_guarantee_amount_type.
	SalesMarketGroupUid                  *int32     `bun:"sales_market_group_uid,type:int"`                              // Foreign key to table Sales Market Group.
	SpecialLabelingFlag                  string     `bun:"special_labeling_flag,type:char(1),default:('N')"`             // Indicates if the customer has special labeling requirements.
	SpecialPackagingFlag                 string     `bun:"special_packaging_flag,type:char(1),default:('N')"`            // Indicates if the customer has special packaging requirements.
	SfdcAccountId                        *string    `bun:"sfdc_account_id,type:varchar(255)"`                            // Salesforce.com account id - added by the import
	SfdcCreateDate                       *time.Time `bun:"sfdc_create_date,type:datetime"`                               // Date the record was created in Salesforce.com - added by import.
	SfdcUpdateDate                       *time.Time `bun:"sfdc_update_date,type:datetime"`                               // Date the record was updated in Salesforce.com - added by import.
	UseSysUpsHandlingChrgFlag            string     `bun:"use_sys_ups_handling_chrg_flag,type:varchar(1),default:('Y')"` // Indicates whether the customer uses the system setting level UPS OnLine fixed handling charge.
	UpsHandlingCharge                    float64    `bun:"ups_handling_charge,type:decimal(19,2),default:((0.00))"`      // Value of handling charge used for UPSOnLine.
	PrintPickFeeFlag                     *string    `bun:"print_pick_fee_flag,type:char(1)"`                             // Custom column to indicate if the pick fee needs to be printed as a separate line in OrdAck, PT and Invoice Form
	Zoom360Grade                         *string    `bun:"zoom360_grade,type:varchar(255)"`                              // Customer grade calculated by Zoom 360 ADS product
	Zoom360GradeExpirationDate           *time.Time `bun:"zoom360_grade_expiration_date,type:datetime"`                  // Date that Zoom 360 grade expires and is no longer valid
	PriceLabelFlag                       *string    `bun:"price_label_flag,type:char(1),default:('N')"`                  // This field indicates whether price labels will be generated during the parting process
	PrintCompleteOrdersFlag              *string    `bun:"print_complete_orders_flag,type:char(1)"`                      // Flag  to determine if the system will print a consolidated invoice when the final non-cancelled pick ticket for an order is confirmed. Possilbe values Y - Yes, N - No.
	BuyListHdrUid                        *int32     `bun:"buy_list_hdr_uid,type:int"`                                    // Custom column for a customer to tie to buy list hdr table for what can be buy what cannot.
	DaysOverdueForCreditHold             *int32     `bun:"days_overdue_for_credit_hold,type:int"`                        // No of days a customer is allowed before putting a credit hold
	GlCodeListHdrUid                     *int32     `bun:"gl_code_list_hdr_uid,type:int"`                                // Unique identifier for a GL Code List
	CmiCustomerNote                      *string    `bun:"cmi_customer_note,type:varchar(8000)"`                         // Custom note for a customer in CMI
	RmaRevenueAccountNo                  *string    `bun:"rma_revenue_account_no,type:varchar(32)"`                      // Default value for rma revenue account.
	SoPoNoRequiredFlag                   string     `bun:"so_po_no_required_flag,type:char(1),default:('N')"`            // Flag for PM order creation if PO number is empty string then check this flag to allow empty string PO number or not
	ArBatchType                          string     `bun:"ar_batch_type,type:char(1),default:('C')"`                     // Defines whether to use the customer invoice batch or a system default invoice batch in Order Entry.
	UseIntAddressFormatFlag              *string    `bun:"use_int_address_format_flag,type:char(1),default:('N')"`       // Indicates whether international address formatting should be used for addresses on invoices.
	GlCodeOverrideFlag                   *string    `bun:"gl_code_override_flag,type:char(1),default:('N')"`             // Flag to override the GL Code List ID for the customer
	Ucc128FormFilename                   *string    `bun:"ucc128_form_filename,type:varchar(255)"`                       // (Custom) Filename for the customer specific UCC128 label form
	LowestAcrossLibrariesFlag            string     `bun:"lowest_across_libraries_flag,type:char(1),default:('N')"`      // Determines if we will apply the lowest of pricing across all libraries defined for this customer.
	PriceRoundingFlag                    *string    `bun:"price_rounding_flag,type:char(1),default:('N')"`               // Indicates whether price rounding will occur for this customer.
	DealerWrrtyClaimsAccountNo           *string    `bun:"dealer_wrrty_claims_account_no,type:varchar(32)"`              // Dealer warranty claims account
	IncludeAgingInfoOnInvoiceFlag        *string    `bun:"include_aging_info_on_invoice_flag,type:char(1)"`              // Indicates if Aging information should be included on invoices forms when printing
	DecimalPrecisionPrice                *int32     `bun:"decimal_precision_price,type:int"`                             // Pricing decimal precision for this customer
	PrivateLabelFlag                     *string    `bun:"private_label_flag,type:char(1)"`                              // Indicates whether the customer is able to order Private Label items
	CfnCostGoodsSoldAccount              *string    `bun:"cfn_cost_goods_sold_account,type:varchar(32)"`                 // column that stored a CFN Cost of Goods Sold Account
	CfnReceivableAccount                 *string    `bun:"cfn_receivable_account,type:varchar(32)"`                      // column that sored a CFN Receivable Account
	CfnRevenueAccount                    *string    `bun:"cfn_revenue_account,type:varchar(32)"`                         // column that stored CFN Revenue Account
	ServicebenchServicerNo               *string    `bun:"servicebench_servicer_no,type:varchar(255)"`                   // Custom column to be used in Service Bench Parts Verification.
	ApplyConvenienceFeeFlag              *string    `bun:"apply_convenience_fee_flag,type:char(1),default:('N')"`        // Flag to determine whether the customer gets charged a convenience fee for credit card
	DefaultKitMarkupPercent              *float64   `bun:"default_kit_markup_percent,type:decimal(19,9)"`                // Component markup percent that will be used for the customer if no product/discount group specific markup has been defined for the component being priced.
	FiscalYearStartMonth                 *int16     `bun:"fiscal_year_start_month,type:smallint"`                        // Custom (F51314): specifies the fiscal year start month
	SalesTaxPayableAccountNo             *string    `bun:"sales_tax_payable_account_no,type:varchar(255)"`               // Sales Tax Payable account no for use with third part tax services
	CustomerTaxClass                     *string    `bun:"customer_tax_class,type:varchar(255)"`                         // Customer Tax Class used in conjunction with third party tax services
	ConsBackordersFlag                   *string    `bun:"cons_backorders_flag,type:char(1)"`                            // Indicates whether backorders can be consolidated for this customer
	ImportPriceTolerance                 *float64   `bun:"import_price_tolerance,type:decimal(19,9),default:((0))"`      // Import Price Tolerance
	ServicebenchAltServicerNo            *string    `bun:"servicebench_alt_servicer_no,type:varchar(255)"`               // Alternate number used in conjunction with ServiceBench claim import to identify a customer
	AlwaysUseEftPaymentFlag              *string    `bun:"always_use_eft_payment_flag,type:char(1)"`                     // Indicates whether the invoices created for this customer will be marked for EFT export.
	EftValidationSentFlag                *string    `bun:"eft_validation_sent_flag,type:char(1)"`                        // Indicates whether this customer has sent its eft transaction file to the bank for valiation.
	ScheduledOrderDiscount               *float64   `bun:"scheduled_order_discount,type:decimal(19,9)"`                  // Discount when the customer has a weekly truck or other regular delivery
	WillCallDiscount                     *float64   `bun:"will_call_discount,type:decimal(19,9)"`                        // This discount occurs when the customer picks up the order.
	FreightThreshold                     *float64   `bun:"freight_threshold,type:decimal(19,9)"`                         // Freight threshold amount
	PartnerProgramUid                    *int32     `bun:"partner_program_uid,type:int"`                                 // Partner Program associated with the customer.
	PlaceOrdersOnHold                    *string    `bun:"place_orders_on_hold,type:char(1)"`                            // Custom Column to indicate if order hold class should be added by default for the customer
	OrderHoldClassId                     *string    `bun:"order_hold_class_id,type:varchar(255)"`                        // Custom column for default hold class put on Order
	DuplicatePoNo                        *string    `bun:"duplicate_po_no,type:char(1)"`                                 // Allow to have duplicate PO numbers while importing.
	FormsOutputPrintFlag                 *string    `bun:"forms_output_print_flag,type:char(1)"`                         // Flag to indicate this customer wants forms to be printed by default.
	FormsOutputEmailFlag                 *string    `bun:"forms_output_email_flag,type:char(1)"`                         // Flag to indicate this customer wants forms to be emailed by default.
	FormsOutputFaxFlag                   *string    `bun:"forms_output_fax_flag,type:char(1)"`                           // Flag to indicate this customer wants forms to be faxed by default.
	CustomerConsolidationMethod          *int32     `bun:"customer_consolidation_method,type:int"`                       // Method this customer uses to generate consolidated invoices.
	ConsolidatePerOrder                  *string    `bun:"consolidate_per_order,type:char(1)"`                           // Automatically consolidate invoices based on order completion
	Edi846pasReferenceNo                 *int32     `bun:"edi846pas_reference_no,type:int"`                              // Custom column to indicate how many times edi846pas has sent to the customer
	PurchasingContactId                  *string    `bun:"purchasing_contact_id,type:varchar(16)"`                       // Purchasing contact ID
	CustomerDiscPct                      *float64   `bun:"customer_disc_pct,type:decimal(19,9)"`                         // Customer discount percent. Used in conjunction with system setting customer_discount_item_id.
	HoldInvoiceFlag                      string     `bun:"hold_invoice_flag,type:char(1),default:('N')"`                 // Determines the default for the order header hold invoice flag.
	GenerateInvoicesBy                   *int16     `bun:"generate_invoices_by,type:smallint,default:((1005))"`          // Indicates if the customer generates invoices using the customer ID or the corporate ID.
	VendorId                             *float64   `bun:"vendor_id,type:decimal(19,0)"`                                 // cross-reference will be made between customers and vendors
	CourtesyPastDueFlag                  string     `bun:"courtesy_past_due_flag,type:char(1),default:('Y')"`            // Flag a customer to be eligible for receiving courtesy past due invoices
	GenericDescOnPackinglistFlag         *string    `bun:"generic_desc_on_packinglist_flag,type:char(1)"`                // Determines if Packing List prints will use Generic Item description instead of the original Item description
	GenericDescOnInvoiceAckFlag          *string    `bun:"generic_desc_on_invoice_ack_flag,type:char(1)"`                // Determines if Order/Quote Acknowledgements and Invoices prints will use Generic Item description instead of the original Item description
	TrackCustomerPackageFlag             *string    `bun:"track_customer_package_flag,type:char(1)"`                     // Indicates whether we allow the customer package tracking process.
	AutoSubstituteItemsFlag              *string    `bun:"auto_substitute_items_flag,type:char(1)"`                      // Setting to see if customer automatically change an item to a substitute.
	FillOrKillFlag                       *string    `bun:"fill_or_kill_flag,type:char(1)"`                               // Setting to determine whether this customer wants to cancel orders or order lines that cannot be fully allocated during order import. [possible values H, L, N, null]
	PreventDcSwitchingFlag               *string    `bun:"prevent_dc_switching_flag,type:char(1),default:('N')"`         // Prevent DC Switching
	BypassDefItemUomFlag                 *string    `bun:"bypass_def_item_uom_flag,type:char(1),default:('N')"`
	DiscountItemId                       *string    `bun:"discount_item_id,type:varchar(255)"`
	DeliveryChargeInvMastUid             *int32     `bun:"delivery_charge_inv_mast_uid,type:int"`
	IndividualInvoiceFlag                *string    `bun:"individual_invoice_flag,type:char(1)"`
	CustomerUid                          int32      `bun:"customer_uid,type:int,autoincrement,unique,identity"` // Unique identifier for the table
	MultCustPartNoGroupFlag              *string    `bun:"mult_cust_part_no_group_flag,type:char(1)"`           // Flag that determines if the customer should use a customer part no group or all groups that the customer is a member of.
	MarketingAllowancePercentage         *float64   `bun:"marketing_allowance_percentage,type:decimal(19,9)"`   // Marketing allowance discount percentage.
	AdAllowancePercentage                *float64   `bun:"ad_allowance_percentage,type:decimal(19,9)"`          // Ad allowance discount percentage.
	DiscountPercentage                   *float64   `bun:"discount_percentage,type:decimal(19,9)"`
	DiscountGlAccountNo                  *string    `bun:"discount_gl_account_no,type:varchar(32)"`
	ValidateOeImportPricingFlag          *string    `bun:"validate_oe_import_pricing_flag,type:char(1)"`
	RecordUsageActualLocFlag             *string    `bun:"record_usage_actual_loc_flag,type:char(1)"`                         // Record the invoice line usage at the actual location.
	StmtFormsOutputPrintFlag             *string    `bun:"stmt_forms_output_print_flag,type:char(1)"`                         // Statement Forms Delivery Methods - PRINT
	StmtFormsOutputEmailFlag             *string    `bun:"stmt_forms_output_email_flag,type:char(1)"`                         // Statement Forms Delivery Methods - EMAIL
	StmtFormsOutputFaxFlag               *string    `bun:"stmt_forms_output_fax_flag,type:char(1)"`                           // Statement Forms Delivery Methods - FAX
	GenPastDueStmtFlag                   *string    `bun:"gen_past_due_stmt_flag,type:char(1)"`                               // Generate Past Due Statement Checkbox
	PdsFormsOutputPrintFlag              *string    `bun:"pds_forms_output_print_flag,type:char(1)"`                          // Past Due Statement Forms Delivery Methods - PRINT
	PdsFormsOutputEmailFlag              *string    `bun:"pds_forms_output_email_flag,type:char(1)"`                          // Past Due Statement Forms Delivery Methods - EMAIL
	PdsFormsOutputFaxFlag                *string    `bun:"pds_forms_output_fax_flag,type:char(1)"`                            // Past Due Statement Forms Delivery Methods - FAX
	ShippingRecalcPriceFlag              *string    `bun:"shipping_recalc_price_flag,type:char(1),default:('N')"`             // Flag to recalculate the price during shipping for the customer
	CalculateCanadianCommFlag            *string    `bun:"calculate_canadian_comm_flag,type:char(1)"`                         // Indicate if customer requires to calculate commissions as canadian.
	ElectronicOrderDiscountIgnoreOverdue *string    `bun:"electronic_order_discount_ignore_overdue,type:char(1)"`             // Customer setting to add Electronic Order Discount wether or not an invoice is over 90 days overdue.
	UsePalletLabels                      *string    `bun:"use_pallet_labels,type:varchar(1)"`                                 // Determines whether cartons and SSCC labels need to be created for pick tickets for the associated customer.
	ConsolidatedAsnFlag                  *string    `bun:"consolidated_asn_flag,type:char(1),default:('N')"`                  // Indicates if the customer consolidates ASNs.
	StartConsolidationDate               *time.Time `bun:"start_consolidation_date,type:datetime"`                            // Start Consolidation Date
	AllowAutoApplyOrigInv                *string    `bun:"allow_auto_apply_orig_inv,type:char(1)"`                            // Allow Auto Apply Original Invoice on RMA
	DrumDepositFlag                      *string    `bun:"drum_deposit_flag,type:char(1)"`                                    // Indicates if this bill to customer should always have drum deposits charged regardless of ship to flag.
	MultiplierPriceLibrary               *float64   `bun:"multiplier_price_library,type:decimal(19,9)"`                       // A discount factor that is muliplied by calculated price using pricing libraries. Custom.
	CompleteLotsFlag                     *string    `bun:"complete_lots_flag,type:char(1)"`                                   // Custom column indicates if the customer requires “Complete Lots”.
	WeightPerBox                         *float64   `bun:"weight_per_box,type:decimal(19,9)"`                                 // Weight Per Box column for the carton label to print
	DoNotChgOeLocToReplenFlag            *string    `bun:"do_not_chg_oe_loc_to_replen_flag,type:char(1)"`                     // Custom: Flag to prevent changing this customers OE source and ship location to replenishment location if cannot allocated complete.
	NatlAcctPricingEligibleFlag          string     `bun:"natl_acct_pricing_eligible_flag,type:char(1),default:('N')"`        // Custom (F67104): determines if this customer is eligible for national account pricing
	ServiceLevelAgreementUid             *int32     `bun:"service_level_agreement_uid,type:int"`                              // Custom: Service level agreement associated with customer record.
	DistributorCode                      *string    `bun:"distributor_code,type:varchar(255)"`                                // Custom (F64358) Informational field used when exporting data. Indicates the customers code for the distributor in external systems. (This is generally safe to repurpose for other features.)
	ReqSuppOrderNoForDispatch            *string    `bun:"req_supp_order_no_for_dispatch,type:char(1)"`                       // Used for custom, indicates whether to require a Supplier Order No value in order to create pick tickets for customer orders in the dispatch window.
	GlAppliedLaborAcct                   *string    `bun:"gl_applied_labor_acct,type:varchar(32)"`                            // G/L Account for Applied Labor on Customer level
	GlLaborCosAcct                       *string    `bun:"gl_labor_cos_acct,type:varchar(32)"`                                // G/L Account for Labor COS on Customer level
	GlServiceCosAcct                     *string    `bun:"gl_service_cos_acct,type:varchar(32)"`                              // G/L Account for Service COS on Customer level
	GlServiceRevenueAcct                 *string    `bun:"gl_service_revenue_acct,type:varchar(32)"`                          // G/L Account for Service Revenue on Customer level
	FreightDiscountAcct                  *string    `bun:"freight_discount_acct,type:varchar(32)"`                            // The account to capture the freight discount.
	WtyClaimApprCreditThreshold          *float64   `bun:"wty_claim_appr_credit_threshold,type:decimal(19,9)"`                // Quantity to issue approved AR credit memos upon creation of the Warranty Claim Entry
	PrntCartonLabelAfterFinalPkgFlag     *string    `bun:"prnt_carton_label_after_final_pkg_flag,type:char(1),default:('N')"` // Indicates if print Carton label after finalizing package for this customer
	PrntShippingLblAfterFinalPkgFlag     *string    `bun:"prnt_shipping_lbl_after_final_pkg_flag,type:char(1),default:('N')"` // Indicates if print shipping label after finalizing package for this customer
	PrntUcc128LabelAfterFinalPkgFlag     *string    `bun:"prnt_ucc128_label_after_final_pkg_flag,type:char(1),default:('N')"` // Indicates ifpPrint GS1-128 (ucc) label after finalizing package (must have customer set to UCC128) for this customer
	UseScanPackFlag                      string     `bun:"use_scan_pack_flag,type:char(1),default:('Y')"`                     // Custom: Indicates that this customer always uses scan and pack (regardless of send_ucc128info setting).
	ExcludeFromExternalTaxCalc           *string    `bun:"exclude_from_external_tax_calc,type:char(1)"`                       // Flag to indicate to not tax from external 3rd party package.
	AdvancedBillingFlag                  *string    `bun:"advanced_billing_flag,type:char(1)"`                                // Indicate if this customer allows advanced billing.
	ConsolidatedInvoicingTypeCd          *int32     `bun:"consolidated_invoicing_type_cd,type:int"`                           // Decides what type of consolidated invoicing to use
	AsnByJobContractLineFlag             *string    `bun:"asn_by_job_contract_line_flag,type:char(1)"`                        // Custom: Indicates corresponding customer receives consolidated 856 ASNS by Job/Contract Line
	Edi832LastFullRunDate                *time.Time `bun:"edi832_last_full_run_date,type:datetime"`                           // Date time when the EDI 832 export was run fully for this customer. (Custom)
	OverrideFcPriceConversionFlag        *string    `bun:"override_fc_price_conversion_flag,type:char(1)"`                    // Custom (F76244): Prevents foreign currency conversion of prices for this customer.  Relies on a special 1:1 exchange rate being setup.
	OverrideFcReceiptDefaults            string     `bun:"override_fc_receipt_defaults,type:char(1),default:('N')"`           // Overrides system settings defaults for Front Counter Receipt
	FcReceiptDefaultPrint                string     `bun:"fc_receipt_default_print,type:char(1),default:('N')"`               // Default value for Print in Receipt groupbox (Front Counter)
	FcReceiptDefaultEmail                string     `bun:"fc_receipt_default_email,type:char(1),default:('N')"`               // Default value for Email in Receipt groupbox (Front Counter)
	FcReceiptDefaultUnpriced             string     `bun:"fc_receipt_default_unpriced,type:char(1),default:('N')"`            // Default value for Unpriced in Receipt groupbox (Front Counter)
	FcReceiptDefaultSkinny               string     `bun:"fc_receipt_default_skinny,type:char(1),default:('N')"`              // Default value for Skinny in Receipt groupbox (Front Counter)
	LegalName                            *string    `bun:"legal_name,type:varchar(255)"`                                      // Extended customer name related with the CSF functionality
}
