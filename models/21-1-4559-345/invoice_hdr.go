package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdr struct {
	bun.BaseModel              `bun:"table:invoice_hdr"`
	InvoiceNo                  string     `bun:"invoice_no,type:varchar(10),pk"`                            // Unique number assigned to each invoice record.
	OrderNo                    *string    `bun:"order_no,type:varchar(8)"`                                  // Order number associated with this invoice.
	OrderDate                  *time.Time `bun:"order_date,type:datetime"`                                  // When was this order taken?
	InvoiceDate                time.Time  `bun:"invoice_date,type:datetime"`                                // Indicates the date on the invoice.
	CustomerId                 string     `bun:"customer_id,type:varchar(12)"`                              // Customer on invoice .
	Bill2Name                  *string    `bun:"bill2_name,type:varchar(50)"`                               // The name of the address that the bill should be sent to.
	Bill2Contact               *string    `bun:"bill2_contact,type:varchar(30)"`                            // The contact for the address that the bill should be sent to.
	Bill2Address1              *string    `bun:"bill2_address1,type:varchar(50)"`                           // The first address line of the address that the bill should be sent to.
	Bill2Address2              *string    `bun:"bill2_address2,type:varchar(50)"`                           // The second address line of the address that the bill should be sent to.
	Bill2City                  *string    `bun:"bill2_city,type:varchar(50)"`                               // The city of the address that the bill should be sent to.
	Bill2State                 *string    `bun:"bill2_state,type:varchar(50)"`                              // The state of the address that the bill should be sent to.
	Bill2PostalCode            *string    `bun:"bill2_postal_code,type:varchar(10)"`                        // The postal code of the address that the bill should be sent to.
	Ship2Name                  *string    `bun:"ship2_name,type:varchar(50)"`                               // Name of the ship to.
	Ship2Contact               *string    `bun:"ship2_contact,type:varchar(30)"`                            // The contact for the ship to address.
	Ship2Address1              *string    `bun:"ship2_address1,type:varchar(50)"`                           // The first line of the ship to address.
	Ship2Address2              *string    `bun:"ship2_address2,type:varchar(50)"`                           // The second line of the ship to address.
	Ship2City                  *string    `bun:"ship2_city,type:varchar(50)"`                               // The ship tos city.
	Ship2State                 *string    `bun:"ship2_state,type:varchar(50)"`                              // The ship tos state.
	Ship2PostalCode            *string    `bun:"ship2_postal_code,type:varchar(10)"`                        // The ship tos postal code.
	CarrierName                *string    `bun:"carrier_name,type:varchar(50)"`                             // The name of the carrier.
	Fob                        *string    `bun:"fob,type:varchar(20)"`                                      // Point in the delivery process at which freight costs and liability become the responsibility of the customer.
	TermsDesc                  *string    `bun:"terms_desc,type:varchar(20)"`                               // Description of the terms.
	PoNo                       *string    `bun:"po_no,type:varchar(255)"`                                   // PO number associated with the invoice.
	SalesrepId                 *string    `bun:"salesrep_id,type:varchar(16)"`                              // What sales representative does this this invoice_line to sales representative relationship refer to?
	SalesrepName               *string    `bun:"salesrep_name,type:varchar(40)"`                            // What sales representative is responsible for this invoice? Why names sometimes null even when the  id points to a salesrep with a  name?:[select salesrep_id -  salesrep_name -  count(*) from invoice_hdr group by salesrep_id -  salesrep_name]
	Brokerage                  *float64   `bun:"brokerage,type:decimal(19,4)"`                              // Amount of brokerage cost on the invoice.
	Freight                    *float64   `bun:"freight,type:decimal(19,4)"`                                // Amount of freight cost on the invoice.
	ArAccountNo                *string    `bun:"ar_account_no,type:varchar(32)"`                            // AR account number on this invoice.
	GlFreightAccountNo         *string    `bun:"gl_freight_account_no,type:varchar(32)"`                    // GL account number used to track freight amount.
	GlBrokerageAccountNo       *string    `bun:"gl_brokerage_account_no,type:varchar(32)"`                  // GL account number used to track brokerage amount.
	BrokerageAmt               *float64   `bun:"brokerage_amt,type:decimal(19,4)"`                          // Brokerage amount on the invoice.
	Period                     *float64   `bun:"period,type:decimal(3,0)"`                                  // The period in which this invoice belongs.
	YearForPeriod              *float64   `bun:"year_for_period,type:decimal(4,0)"`                         // The year in which this invoice belongs.
	StoreNo                    *string    `bun:"store_no,type:varchar(5)"`                                  // This column is no longer being used and is scheduled for removal in later version.
	InvoiceType                *string    `bun:"invoice_type,type:varchar(2)"`                              // Indicates the type of invoice (ie - Invoice, Debit, Downpayment)
	ShipToId                   *float64   `bun:"ship_to_id,type:decimal(19,0)"`                             // Ship To associated with the invoice.
	ShipDate                   *time.Time `bun:"ship_date,type:datetime"`                                   // Shipping date for this invoice.
	TotalAmount                float64    `bun:"total_amount,type:decimal(19,4),default:(0)"`               // Total amount on the invoice.
	AmountPaid                 float64    `bun:"amount_paid,type:decimal(19,4),default:(0)"`                // Amount paid on the invoice.
	TermsTaken                 float64    `bun:"terms_taken,type:decimal(19,4),default:(0)"`                // Amount of terms taken on the invoice.
	Allowed                    float64    `bun:"allowed,type:decimal(19,4),default:(0)"`                    // Allowed amount on the invoice.
	PaidInFullFlag             *string    `bun:"paid_in_full_flag,type:char(1),default:('N')"`              // Indicates whether the invoice is paid in full.
	PaidByCheckNo              *float64   `bun:"paid_by_check_no,type:decimal(19,4)"`                       // Indicates the check number which paid the invoice.
	DatePaid                   *time.Time `bun:"date_paid,type:datetime"`                                   // Indicates the date the invoice was paid.
	PrintFlag                  *string    `bun:"print_flag,type:char(1)"`                                   // Indicates whether the invoice was printed.
	PrintDate                  *time.Time `bun:"print_date,type:datetime"`                                  // The date the invoice was printed.
	CompanyNo                  string     `bun:"company_no,type:varchar(8)"`                                // Unique code that identifies a company.
	CustomerIdNumber           float64    `bun:"customer_id_number,type:decimal(19,0)"`                     // The customer on the invoice.
	DateCreated                time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Printed                    *string    `bun:"printed,type:char(1)"`                                      // Indicates whether the invoice was printed.
	PrintedDate                *time.Time `bun:"printed_date,type:datetime"`                                // The date the invoice was printed.
	CorpAddressId              *float64   `bun:"corp_address_id,type:decimal(19,0)"`                        // Indicates the corporate address for this invoice.
	ShippingCost               *float64   `bun:"shipping_cost,type:decimal(19,4)"`                          // The shipping cost for this invoice.
	Bill2Country               *string    `bun:"bill2_country,type:varchar(50)"`                            // The country of the billing address that this invoice applies to.
	Ship2Country               *string    `bun:"ship2_country,type:varchar(50)"`                            // The country of the ship to address that this invoice applies to.
	InvoiceReferenceNo         string     `bun:"invoice_reference_no,type:varchar(10)"`                     // The invoice number that this invoice is referencing.
	InvoiceAdjustmentType      *string    `bun:"invoice_adjustment_type,type:char(1)"`                      // The type of invoice.
	InvoiceDesc                *string    `bun:"invoice_desc,type:varchar(255)"`                            // A description of the invoice.
	MemoAmount                 float64    `bun:"memo_amount,type:decimal(19,4),default:(0)"`                // The memo amount on the invoice.
	BadDebtAmount              float64    `bun:"bad_debt_amount,type:decimal(19,4),default:(0)"`            // The bad debt amount on the invoice.
	InvoiceClass               *string    `bun:"invoice_class,type:varchar(8)"`                             // Invoice class.
	PeriodFullyPaid            *float64   `bun:"period_fully_paid,type:decimal(3,0)"`                       // Indicates the period the invoice was fully paid.
	YearFullyPaid              *float64   `bun:"year_fully_paid,type:decimal(4,0)"`                         // Indicates the year the invoice was fully paid.
	Approved                   *string    `bun:"approved,type:char(1)"`                                     // This indicates whether or not the voucher is approved.
	FcThruDate                 *time.Time `bun:"fc_thru_date,type:datetime"`                                // The last time a finance charge invoice was generated for a late invoice.  The Finance Charge Through Date is used if an additional invoice is generated for the same late invoice.
	CumulativeFc               *float64   `bun:"cumulative_fc,type:decimal(19,4)"`                          // The total amount of finance charges charged against a particular invoice.
	NetDueDate                 *time.Time `bun:"net_due_date,type:datetime"`                                // This is the net due date.
	TermsDueDate               *time.Time `bun:"terms_due_date,type:datetime"`                              // This is the terms due date.
	TermsId                    *string    `bun:"terms_id,type:varchar(2)"`                                  // Indicates the terms associated with the invoice.
	BranchId                   string     `bun:"branch_id,type:varchar(8)"`                                 // Indicates the branch of the invoice.
	DisputedFlag               *string    `bun:"disputed_flag,type:char(1)"`                                // Indicates whether the invoice is disputed.
	StatementPeriod            *float64   `bun:"statement_period,type:decimal(3,0)"`                        // Indicates the statement period.
	StatementYear              *float64   `bun:"statement_year,type:decimal(4,0)"`                          // Indicates the statement year.
	OtherChargeAmount          float64    `bun:"other_charge_amount,type:decimal(19,4),default:(0)"`        // What is the total of other charge items of the inv
	TaxAmount                  float64    `bun:"tax_amount,type:decimal(19,4),default:(0)"`                 // What is the total of the tax amount of the invoice
	OriginalDocumentType       string     `bun:"original_document_type,type:char(1)"`                       // Where was the invoice was generated from?
	Consolidated               string     `bun:"consolidated,type:char(1)"`                                 // Is this invoice a consolated invoice?
	SoldToAhUid                int32      `bun:"sold_to_ah_uid,type:int"`                                   // What is the address history entry that this invoice was sold to?
	SoldToCustomerId           float64    `bun:"sold_to_customer_id,type:decimal(19,0)"`                    // What customer was this invoice sold to?
	ShipToPhone                *string    `bun:"ship_to_phone,type:varchar(20)"`                            // What is the phone number for the location that this will be shipped to?
	InvoiceBatchUid            int32      `bun:"invoice_batch_uid,type:int,default:(1)"`                    // Identifies the invoice batch number to which this invoice belongs.
	FreightCodeUid             *int32     `bun:"freight_code_uid,type:int"`                                 // Identifies the freight code for the invoice.
	ShippingRouteUid           *int32     `bun:"shipping_route_uid,type:int"`                               // Identifies the shipping route for the invoice.
	TransmissionMethod         *int32     `bun:"transmission_method,type:int"`                              // The EDI tranmission method that is used for this invoice.
	TermsAmount                float64    `bun:"terms_amount,type:decimal(19,4),default:(0)"`               // The calculated terms amount for this invoice.
	SalesLocationId            *float64   `bun:"sales_location_id,type:decimal(19,0)"`                      // The location that receives credit for making a sale.
	SourceTypeCd               *int32     `bun:"source_type_cd,type:int"`                                   // Identifies the source where this invoice record was created.
	Ship2EmailAddress          *string    `bun:"ship2_email_address,type:varchar(255)"`                     // email address associated w/ship-to address
	CurrencyLineUid            *int32     `bun:"currency_line_uid,type:int"`                                // Identifies which currency_line rcd relates to this invoice.
	CreatedBy                  *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TaxTermsTaken              float64    `bun:"tax_terms_taken,type:decimal(19,9),default:((0.00))"` // The amount of the terms against taxes that has already been taken
	CarrierInsuranceAmt        *float64   `bun:"carrier_insurance_amt,type:decimal(19,9)"`            // Insurance based on value of shipped goods.
	InvNoDisplay               *string    `bun:"inv_no_display,type:varchar(255)"`                    // Invoice number display - allows for prefixed invoice numbers
	IvaExemptionNumber         *string    `bun:"iva_exemption_number,type:varchar(30)"`               // IVA Tax Exemption Number for the customer on the invoice
	IvaTaxableFlag             *string    `bun:"iva_taxable_flag,type:char(1)"`                       // Indicate if the customer on the invoice computes IVA tax.
	JobId                      *string    `bun:"job_id,type:varchar(32)"`                             // Indentifier for the job.
	InvoicePaidPeriod          *int32     `bun:"invoice_paid_period,type:int"`
	InvoicePeriod              *int32     `bun:"invoice_period,type:int"`
	CreditMemoForTermsFlag     *string    `bun:"credit_memo_for_terms_flag,type:char(1)"`            // This flag will be used to determine if the current invoice was created to deal with terms taken on another inovoice
	RecordTypeCd               *int32     `bun:"record_type_cd,type:int"`                            // Determines the type of record.
	CommissionCost             *float64   `bun:"commission_cost,type:decimal(19,9),default:((0))"`   // The header level commission cost.  Used to calculate commissions if commission is paid by the total invoice.
	ReceiverName               *string    `bun:"receiver_name,type:varchar(255)"`                    // Receiver name for Pegmost Export
	CartonCount                *int32     `bun:"carton_count,type:int,default:(NULL)"`               // Number of cartons associated with the invoice.
	StrategicFreightIn         *float64   `bun:"strategic_freight_in,type:decimal(19,9)"`            // Amount of Strategic Freight In that was calculated to be charged on this invoice.  This is for reporting purposes only.
	StrategicFreightOut        *float64   `bun:"strategic_freight_out,type:decimal(19,9)"`           // Amount of Strategic Freight Out that was calculated to be charged on this invoice.  This is for reporting purposes only.
	AffiliatedTrainingCenter   *float64   `bun:"affiliated_training_center,type:decimal(19,0)"`      // Training center associated with the customer
	ExternalReferenceNo        *string    `bun:"external_reference_no,type:varchar(255)"`            // Customer defined usage - no validatation
	TotalFreightchargeWeight   *float64   `bun:"total_freightcharge_weight,type:decimal(19,4)"`      // Sum of the extended weight (Item Maintenance Weight/Volume Tabs Gross weight value) of all items on the shipment except those that are listed on the free freight tab in the customer record (suppliers and items).
	CardlockConsInvoiceFlag    *string    `bun:"cardlock_cons_invoice_flag,type:char(1)"`            // Indicates if this invoice was consolidated (Y) from cardlock invoices, and will print with the appropriate form.
	RemoveFromOpenDefRevWindow *string    `bun:"remove_from_open_def_rev_window,type:char(1)"`       // Indicates that this downpayment invoice should not show in the Open Deferred Revenue Window
	RecordTypeReferenceNo      *string    `bun:"record_type_reference_no,type:varchar(10)"`          // Generic column whose value corresponds to the description in record_type_cd.
	RebillInvoiceReasonUid     *int32     `bun:"rebill_invoice_reason_uid,type:int"`                 // Populated only on rebill invoices.  Indicates reason for credit and rebill.
	EdiOrderPrintedFlag        *string    `bun:"edi_order_printed_flag,type:char(1)"`                // Custom column to indicate the EDI order invoice gets physically printed for the first time in Invoice Report window.
	DownpaymentApplied         *float64   `bun:"downpayment_applied,type:decimal(19,2)"`             // Represents the DP amt applied for a shipping invoice. And the DP amount for a DP invoice or credit/debit memo that references a DP invoice.
	ReverseRedemptionFlag      *string    `bun:"reverse_redemption_flag,type:char(1),default:('N')"` // A flag (Y, N) to indicate a reversal of the redemption invoice.
	SentToCarrierDate          *time.Time `bun:"sent_to_carrier_date,type:datetime"`                 // Date that this invoice was exported to Carrier.
	TaxTermsAmt                float64    `bun:"tax_terms_amt,type:decimal(19,9),default:((0))"`
	TaxAmountPaid              float64    `bun:"tax_amount_paid,type:decimal(19,9),default:((0))"`
	ReasonCreditMemoCodeUid    *int32     `bun:"reason_credit_memo_code_uid,type:int"`                 // Reason Credit Memo Code associated with the record.
	SourceCreditMemoCodeUid    *int32     `bun:"source_credit_memo_code_uid,type:int"`                 // Source Credit Memo Code associated with the record.
	ExternalClaimId            *string    `bun:"external_claim_id,type:varchar(255)"`                  // Claim No field used in AR Credit/Debit memo window
	Bill2Address3              *string    `bun:"bill2_address3,type:varchar(50)"`                      // Billing address line 3
	Ship2Address3              *string    `bun:"ship2_address3,type:varchar(50)"`                      // Shipping address line 3
	EftExportedDate            *time.Time `bun:"eft_exported_date,type:datetime"`                      // Date that this invoice was exported for ACH payment to the bank.
	CourtesyInvoiceFlag        *string    `bun:"courtesy_invoice_flag,type:char(1),default:('N')"`     // Flag if a courtesy past due invoice has already been sent
	IvaWithheldAmount          float64    `bun:"iva_withheld_amount,type:decimal(19,2),default:((0))"` // Stores freight IVA withheld amount
	CarrierId                  *float64   `bun:"carrier_id,type:decimal(19,0)"`                        // The numeric ID of the carrier for the invoice.
	FreeOutFreightMinMetFlag   *string    `bun:"free_out_freight_min_met_flag,type:char(1),default:('N')"`
	RegNumOfPackages           *string    `bun:"reg_num_of_packages,type:varchar(255)"`
	Ship2Latitude              *string    `bun:"ship2_latitude,type:varchar(20)"`  // Latitude for ship to address for Avalara.
	Ship2Longitude             *string    `bun:"ship2_longitude,type:varchar(20)"` // Longitude for ship to address for Avalara.
	DeliveryNo                 *int32     `bun:"delivery_no,type:int"`
	AllowedAmtReasonId         *string    `bun:"allowed_amt_reason_id,type:varchar(255)"`
	TotalServiceAmountRedeemed *float64   `bun:"total_service_amount_redeemed,type:decimal(19,9)"`     // The total service amount redeemed against the invoice from service vouchers
	CustomerClaimNo            *string    `bun:"customer_claim_no,type:varchar(255)"`                  // Non-accrued customer claims number.
	ClaimType                  *string    `bun:"claim_type,type:varchar(255)"`                         // Claim type for customer_claim.
	ClaimCoopFlag              *string    `bun:"claim_coop_flag,type:char(1)"`                         // Coop flag for customer_claim.
	BillToSupplierFlag         *string    `bun:"bill_to_supplier_flag,type:char(1)"`                   // Check if this memo is later billed to supplier.
	CustPriceProtVarAcctNo     *string    `bun:"cust_price_prot_var_acct_no,type:varchar(32)"`         // Variance account for customer price protection memos.
	ConsolidatedTypeFlag       *string    `bun:"consolidated_type_flag,type:char(1)"`                  // Consolidated Type Flag. (S = Single, W = Waiting to be consolidated)
	MerchandiseInvoiceFlag     *string    `bun:"merchandise_invoice_flag,type:char(1)"`                // Custom colum to indicate the invoice has merchandise items
	SalesMarketGroupUid        *int32     `bun:"sales_market_group_uid,type:int"`                      // Foreign key to table Sales Market Group.
	ActualFreightCost          *float64   `bun:"actual_freight_cost,type:decimal(19,9)"`               // Save actual freight cost from shipping and direct ship confirmation so the user can access it for reporting
	EcoFeeAmt                  *float64   `bun:"eco_fee_amt,type:decimal(19,9)"`                       // The total of the eco fee amount of the invoice
	FrghtDiscTaken             float64    `bun:"frght_disc_taken,type:decimal(19,4),default:((0.00))"` // Amount of freight discount taken on the invoice.
	RentalInvoiceNo            *string    `bun:"rental_invoice_no,type:varchar(10)"`                   // Invoice number for Rentals
}
