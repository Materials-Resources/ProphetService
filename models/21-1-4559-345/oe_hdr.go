package model

type OeHdr struct {
	bun.BaseModel                  `bun:"table:oe_hdr"`
	OrderNo                        string    `bun:"order_no,type:varchar(8),pk"`
	CustomerId                     float64   `bun:"customer_id,type:decimal(19,0)"`
	OrderDate                      time.Time `bun:"order_date,type:datetime,nullzero"`
	Ship2Name                      string    `bun:"ship2_name,type:varchar(50),nullzero"`
	Ship2Add1                      string    `bun:"ship2_add1,type:varchar(50),nullzero"`
	Ship2Add2                      string    `bun:"ship2_add2,type:varchar(50),nullzero"`
	Ship2City                      string    `bun:"ship2_city,type:varchar(50),nullzero"`
	Ship2State                     string    `bun:"ship2_state,type:varchar(50),nullzero"`
	Ship2Zip                       string    `bun:"ship2_zip,type:varchar(10),nullzero"`
	Ship2Country                   string    `bun:"ship2_country,type:varchar(50),nullzero"`
	RequestedDate                  time.Time `bun:"requested_date,type:datetime,nullzero"`
	PoNo                           string    `bun:"po_no,type:varchar(50),nullzero"`
	Terms                          string    `bun:"terms,type:varchar(2),nullzero"`
	ShipToPhone                    string    `bun:"ship_to_phone,type:varchar(20),nullzero"`
	DeleteFlag                     string    `bun:"delete_flag,type:char"`
	Completed                      string    `bun:"completed,type:char,nullzero"`
	CompanyId                      string    `bun:"company_id,type:varchar(8),nullzero"`
	DateCreated                    time.Time `bun:"date_created,type:datetime"`
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30)"`
	CodFlag                        string    `bun:"cod_flag,type:char,nullzero"`
	GrossMargin                    float64   `bun:"gross_margin,type:decimal(19,9),nullzero"`
	ProjectedOrder                 string    `bun:"projected_order,type:char,nullzero"`
	PoNoAppend                     string    `bun:"po_no_append,type:varchar(20),nullzero"`
	LocationId                     float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	CarrierId                      float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`
	AddressId                      float64   `bun:"address_id,type:decimal(19,0),nullzero"`
	ContactId                      string    `bun:"contact_id,type:varchar(16),nullzero"`
	CorpAddressId                  float64   `bun:"corp_address_id,type:decimal(19,0),nullzero"`
	HandlingChargeReqFlag          string    `bun:"handling_charge_req_flag,type:char,nullzero"`
	PaymentMethod                  string    `bun:"payment_method,type:varchar(8),nullzero"`
	FobFlag                        string    `bun:"fob_flag,type:char,nullzero"`
	Class1id                       string    `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                       string    `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                       string    `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                       string    `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                       string    `bun:"class_5id,type:varchar(255),nullzero"`
	RmaFlag                        string    `bun:"rma_flag,type:char,nullzero"`
	Taker                          string    `bun:"taker,type:varchar(30),nullzero"`
	JobName                        string    `bun:"job_name,type:varchar(40),nullzero"`
	ThirdPartyBillingFlag          string    `bun:"third_party_billing_flag,type:char,nullzero"`
	Approved                       string    `bun:"approved,type:char,nullzero"`
	SourceLocationId               float64   `bun:"source_location_id,type:decimal(19,0),nullzero"`
	PackingBasis                   string    `bun:"packing_basis,type:varchar(16),nullzero"`
	DeliveryInstructions           string    `bun:"delivery_instructions,type:varchar(255),nullzero"`
	PickTicketType                 string    `bun:"pick_ticket_type,type:varchar(2),nullzero"`
	RequestedDownpayment           float64   `bun:"requested_downpayment,type:decimal(19,4),nullzero"`
	DownpaymentInvoiced            string    `bun:"downpayment_invoiced,type:char,nullzero"`
	CancelFlag                     string    `bun:"cancel_flag,type:char,nullzero"`
	WillCall                       string    `bun:"will_call,type:char,nullzero"`
	FrontCounter                   string    `bun:"front_counter,type:char,nullzero"`
	ValidationStatus               string    `bun:"validation_status,type:varchar(8),nullzero"`
	OeHdrUid                       int32     `bun:"oe_hdr_uid,type:int"`
	SourceId                       string    `bun:"source_id,type:varchar(8),nullzero"`
	SourceCodeNo                   int32     `bun:"source_code_no,type:int"`
	CreditCardHold                 int32     `bun:"credit_card_hold,type:int,nullzero"`
	InvoiceBatchUid                int32     `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid                 int32     `bun:"freight_code_uid,type:int,nullzero"`
	FreightOut                     float64   `bun:"freight_out,type:decimal(19,4),default:(0)"`
	ShippingRouteUid               int32     `bun:"shipping_route_uid,type:int,nullzero"`
	ExcludeRebates                 string    `bun:"exclude_rebates,type:char"`
	CaptureUsageDefault            string    `bun:"capture_usage_default,type:char"`
	JobPriceHdrUid                 int32     `bun:"job_price_hdr_uid,type:int,nullzero"`
	FrontCounterRma                string    `bun:"front_counter_rma,type:varchar,default:('N')"`
	Ship2EmailAddress              string    `bun:"ship2_email_address,type:varchar(255),nullzero"`
	OrderType                      int32     `bun:"order_type,type:int,nullzero"`
	Taxable                        string    `bun:"taxable,type:char,nullzero"`
	ProfitPercent                  float64   `bun:"profit_percent,type:decimal(19,9),default:(0.00)"`
	OrderCostBasis                 int32     `bun:"order_cost_basis,type:int,default:(1)"`
	CurrencyLineUid                int32     `bun:"currency_line_uid,type:int,nullzero"`
	InvoiceExchRateSourceCd        int32     `bun:"invoice_exch_rate_source_cd,type:int,nullzero"`
	RmaExpirationDate              time.Time `bun:"rma_expiration_date,type:datetime,default:(getdate())"`
	CreatedBy                      string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	InvoiceNo                      string    `bun:"invoice_no,type:varchar(10),nullzero"`
	BillToContactId                string    `bun:"bill_to_contact_id,type:varchar(16),nullzero"`
	TagHoldCancelDate              time.Time `bun:"tag_hold_cancel_date,type:datetime,nullzero"`
	CostCenterTrackingOption       int32     `bun:"cost_center_tracking_option,type:int,nullzero"`
	SkipProfitExceptionCheck       string    `bun:"skip_profit_exception_check,type:char,default:('N')"`
	ConsBackorderProcessingFlag    string    `bun:"cons_backorder_processing_flag,type:char,nullzero"`
	RestockFeePercentage           float64   `bun:"restock_fee_percentage,type:decimal(19,9),nullzero"`
	DateOrderCompleted             time.Time `bun:"date_order_completed,type:datetime,nullzero"`
	ValidatedViaOpenOrdersFlag     string    `bun:"validated_via_open_orders_flag,type:char,default:('N')"`
	AcknowledgementDate            time.Time `bun:"acknowledgement_date,type:datetime,nullzero"`
	ProductGroupCostBasis          int32     `bun:"product_group_cost_basis,type:int,nullzero"`
	OriginalPromiseDate            time.Time `bun:"original_promise_date,type:datetime,nullzero"`
	PromiseDate                    time.Time `bun:"promise_date,type:datetime,nullzero"`
	RequestedShipDate              time.Time `bun:"requested_ship_date,type:datetime,nullzero"`
	ExpectedCompletionDate         time.Time `bun:"expected_completion_date,type:datetime,nullzero"`
	JobControlFlag                 string    `bun:"job_control_flag,type:char,nullzero"`
	ApplyBuilderAllowanceFlag      string    `bun:"apply_builder_allowance_flag,type:char,default:('N')"`
	MerchandiseCreditFlag          string    `bun:"merchandise_credit_flag,type:char,default:('N')"`
	ReqPymtUponReleaseFlag         string    `bun:"req_pymt_upon_release_flag,type:char,default:('N')"`
	PmDate                         time.Time `bun:"pm_date,type:datetime,nullzero"`
	DownpaymentPercentage          float64   `bun:"downpayment_percentage,type:decimal(19,2),nullzero"`
	PrepaidInvoiceFlag             string    `bun:"prepaid_invoice_flag,type:char,nullzero"`
	LandedCostIncludedCd           int32     `bun:"landed_cost_included_cd,type:int,nullzero"`
	ExportedFlag                   string    `bun:"exported_flag,type:char,nullzero"`
	WebReferenceNo                 string    `bun:"web_reference_no,type:varchar(255),nullzero"`
	UpsCode                        string    `bun:"ups_code,type:varchar(40),nullzero"`
	DefaultPricingCd               int32     `bun:"default_pricing_cd,type:int,nullzero"`
	FreightChargeByMileHdrUid      int32     `bun:"freight_charge_by_mile_hdr_uid,type:int,nullzero"`
	FreightMileageAmt              float64   `bun:"freight_mileage_amt,type:decimal(19,2),nullzero"`
	SupplierOrderNo                string    `bun:"supplier_order_no,type:varchar(255),nullzero"`
	SupplierReleaseNo              string    `bun:"supplier_release_no,type:varchar(255),nullzero"`
	RoutedEtaDate                  time.Time `bun:"routed_eta_date,type:datetime,nullzero"`
	RouteOverrideDate              time.Time `bun:"route_override_date,type:datetime,nullzero"`
	ExpediteDate                   time.Time `bun:"expedite_date,type:datetime,nullzero"`
	OrderOpenStartDate             time.Time `bun:"order_open_start_date,type:datetime,nullzero"`
	EnvironmentalFeeFlag           string    `bun:"environmental_fee_flag,type:char,default:('N')"`
	AdminFeeFlag                   string    `bun:"admin_fee_flag,type:char,default:('N')"`
	PromiseDateExtendedDesc        string    `bun:"promise_date_extended_desc,type:varchar(8000),nullzero"`
	PromiseDateEditedDate          time.Time `bun:"promise_date_edited_date,type:datetime,nullzero"`
	OrderDiscType                  int32     `bun:"order_disc_type,type:int,nullzero"`
	OrderDiscFactor                float64   `bun:"order_disc_factor,type:decimal(19,9),nullzero"`
	OrderAckPrintedFlag            string    `bun:"order_ack_printed_flag,type:char,nullzero"`
	PackingListSentFlag            int32     `bun:"packing_list_sent_flag,type:int,nullzero"`
	RmaDeliveryListStatus          int32     `bun:"rma_delivery_list_status,type:int,default:((1024))"`
	StrategicLibraryUid            int32     `bun:"strategic_library_uid,type:int,nullzero"`
	StrategicLibraryOriginalUid    int32     `bun:"strategic_library_original_uid,type:int,nullzero"`
	SendPartialOrderFlag           string    `bun:"send_partial_order_flag,type:char,default:('N')"`
	SupplierStatus                 string    `bun:"supplier_status,type:char,nullzero"`
	OrderPriorityUid               int32     `bun:"order_priority_uid,type:int,nullzero"`
	B2bUpsFreightAmount            float64   `bun:"b2b_ups_freight_amount,type:decimal(19,9),nullzero"`
	UserDefinedDate                time.Time `bun:"user_defined_date,type:datetime,nullzero"`
	BlindAddressingFlag            string    `bun:"blind_addressing_flag,type:char,default:('N')"`
	ReplaceCompanyNameFlag         string    `bun:"replace_company_name_flag,type:char,default:('N')"`
	UseVendorItemTermsFlag         string    `bun:"use_vendor_item_terms_flag,type:char,nullzero"`
	FirstPackingListNo             float64   `bun:"first_packing_list_no,type:decimal(19,0),nullzero"`
	SalesMarketGroupUid            int32     `bun:"sales_market_group_uid,type:int,nullzero"`
	ShipConfirmedFlag              string    `bun:"ship_confirmed_flag,type:char,nullzero"`
	PriceConfirmedFlag             string    `bun:"price_confirmed_flag,type:char,nullzero"`
	QuotedFreightOut               float64   `bun:"quoted_freight_out,type:decimal(19,9),nullzero"`
	GeocomAppend                   int32     `bun:"geocom_append,type:int,nullzero"`
	ApplyFuelSurchargeFlag         string    `bun:"apply_fuel_surcharge_flag,type:char,nullzero"`
	OverrideFreightCodeFlag        string    `bun:"override_freight_code_flag,type:char,nullzero"`
	ServiceOrderPriorityUid        int32     `bun:"service_order_priority_uid,type:int,nullzero"`
	ExcludeFromCreditLimitFlag     string    `bun:"exclude_from_credit_limit_flag,type:char,nullzero"`
	PackingListFilename            string    `bun:"packing_list_filename,type:char,nullzero"`
	OverrideMinOrderChargeFlag     string    `bun:"override_min_order_charge_flag,type:char,nullzero"`
	PlacedByName                   string    `bun:"placed_by_name,type:varchar(255),nullzero"`
	ApprovedForArFlag              string    `bun:"approved_for_ar_flag,type:char,nullzero"`
	SecondRouteOverrideDate        time.Time `bun:"second_route_override_date,type:datetime,nullzero"`
	Subject                        string    `bun:"subject,type:varchar(255),nullzero"`
	ServicePlanUid                 int32     `bun:"service_plan_uid,type:int,nullzero"`
	CentralFaxNumber               string    `bun:"central_fax_number,type:varchar(20),nullzero"`
	Url                            string    `bun:"url,type:varchar(255),nullzero"`
	OriginalPackingBasis           string    `bun:"original_packing_basis,type:varchar(16),nullzero"`
	FreightOutEditedFlag           string    `bun:"freight_out_edited_flag,type:char,nullzero"`
	OverrideContactEmailFlag       string    `bun:"override_contact_email_flag,type:char,default:('N')"`
	OverrideEmailAddress           string    `bun:"override_email_address,type:varchar(255),nullzero"`
	PrintPricesOnPackinglist       string    `bun:"print_prices_on_packinglist,type:char,nullzero"`
	ImportSource                   string    `bun:"import_source,type:varchar(255),nullzero"`
	QuoteLockedFlag                string    `bun:"quote_locked_flag,type:char,nullzero"`
	WebShopperId                   int32     `bun:"web_shopper_id,type:int,nullzero"`
	WebShopperEmail                string    `bun:"web_shopper_email,type:varchar(255),nullzero"`
	RecalcScheduledDsPrice         string    `bun:"recalc_scheduled_ds_price,type:char,nullzero"`
	OrderTypeCust                  string    `bun:"order_type_cust,type:varchar(255),nullzero"`
	EngineerId                     string    `bun:"engineer_id,type:varchar(255),nullzero"`
	CarrierContractHdrUid          int32     `bun:"carrier_contract_hdr_uid,type:int,nullzero"`
	OrderAckPrintPricesFlag        string    `bun:"order_ack_print_prices_flag,type:char,nullzero"`
	FreightOutEstimate             float64   `bun:"freight_out_estimate,type:decimal(19,9),nullzero"`
	FreightChargeEstimate          float64   `bun:"freight_charge_estimate,type:decimal(19,9),nullzero"`
	DocumentCaptureDate            time.Time `bun:"document_capture_date,type:datetime,nullzero"`
	LimitMaxShipmentsPerOrder      float64   `bun:"limit_max_shipments_per_order,type:decimal(19,4),nullzero"`
	ReasonCreditMemoCodeUid        int32     `bun:"reason_credit_memo_code_uid,type:int,nullzero"`
	SourceCreditMemoCodeUid        int32     `bun:"source_credit_memo_code_uid,type:int,nullzero"`
	NetBillingFlag                 string    `bun:"net_billing_flag,type:char,default:('N')"`
	NetBillingEdited               string    `bun:"net_billing_edited,type:char,default:('N')"`
	SchedOrderDiscFlag             string    `bun:"sched_order_disc_flag,type:char,nullzero"`
	WillCallDiscFlag               string    `bun:"will_call_disc_flag,type:char,nullzero"`
	SingleOrderDiscFlag            string    `bun:"single_order_disc_flag,type:char,nullzero"`
	VolumeDiscFlag                 string    `bun:"volume_disc_flag,type:char,nullzero"`
	ServiceInvoiceNo               string    `bun:"service_invoice_no,type:varchar(10),nullzero"`
	Ship2Add3                      string    `bun:"ship2_add3,type:varchar(50),nullzero"`
	BillToId                       float64   `bun:"bill_to_id,type:decimal(19,0),nullzero"`
	ElectronicOrderFlag            string    `bun:"electronic_order_flag,type:char,default:('N')"`
	CodWithoutRemittanceFlag       string    `bun:"cod_without_remittance_flag,type:char,nullzero"`
	SoTotalPartsPriceOnInvoice     string    `bun:"so_total_parts_price_on_invoice,type:char,nullzero"`
	SoTotalLaborPriceOnInvoice     string    `bun:"so_total_labor_price_on_invoice,type:char,nullzero"`
	InquiryFlag                    string    `bun:"inquiry_flag,type:char,nullzero"`
	HoldInvoiceFlag                string    `bun:"hold_invoice_flag,type:char,default:('N')"`
	DoNotExportToPtsFlag           string    `bun:"do_not_export_to_pts_flag,type:char,default:('N')"`
	GenericDescOnPackinglistFlag   string    `bun:"generic_desc_on_packinglist_flag,type:char,nullzero"`
	GenericDescOnInvoiceAckFlag    string    `bun:"generic_desc_on_invoice_ack_flag,type:char,nullzero"`
	WillCallNotificationFlag       string    `bun:"will_call_notification_flag,type:char,nullzero"`
	PrebillingDate                 time.Time `bun:"prebilling_date,type:datetime,nullzero"`
	BlindShipFlag                  string    `bun:"blind_ship_flag,type:char,nullzero"`
	ArchitectId                    float64   `bun:"architect_id,type:decimal(19,0),nullzero"`
	BuilderId                      float64   `bun:"builder_id,type:decimal(19,0),nullzero"`
	ContractorId                   float64   `bun:"contractor_id,type:decimal(19,0),nullzero"`
	DesignerId                     float64   `bun:"designer_id,type:decimal(19,0),nullzero"`
	HomeownerId                    float64   `bun:"homeowner_id,type:decimal(19,0),nullzero"`
	QuoteType                      int32     `bun:"quote_type,type:int,nullzero"`
	PtsLabelPrintFlag              string    `bun:"pts_label_print_flag,type:char,nullzero"`
	DiscountItemId                 string    `bun:"discount_item_id,type:varchar(255),nullzero"`
	CustomerDiscPct                float64   `bun:"customer_disc_pct,type:decimal(19,9),nullzero"`
	ConsolidatePerOrderFlag        string    `bun:"consolidate_per_order_flag,type:char,nullzero"`
	DfltUpsShipViaCode             int32     `bun:"dflt_ups_ship_via_code,type:int,nullzero"`
	DfltFedexServiceTypeUid        int32     `bun:"dflt_fedex_service_type_uid,type:int,nullzero"`
	InvoiceOnlyFlag                string    `bun:"invoice_only_flag,type:char,nullzero"`
	Ship2Latitude                  string    `bun:"ship2_latitude,type:varchar(20),nullzero"`
	Ship2Longitude                 string    `bun:"ship2_longitude,type:varchar(20),nullzero"`
	BypassHandlingChargeFlag       string    `bun:"bypass_handling_charge_flag,type:char,nullzero"`
	FreightChargeUid               int32     `bun:"freight_charge_uid,type:int,nullzero"`
	LumpSumBillingFlag             string    `bun:"lump_sum_billing_flag,type:char,nullzero"`
	FieldDestroyFlag               string    `bun:"field_destroy_flag,type:char,nullzero"`
	WarrantyReturnOrderFlag        string    `bun:"warranty_return_order_flag,type:char,default:('N')"`
	AllowAutoApplyOrigInvFlag      string    `bun:"allow_auto_apply_orig_inv_flag,type:char,nullzero"`
	DaysEarly                      int32     `bun:"days_early,type:int,nullzero"`
	TransitDays                    int32     `bun:"transit_days,type:int,nullzero"`
	DisplayInvoiceExchRateSourceCd int32     `bun:"display_invoice_exch_rate_source_cd,type:int,nullzero"`
	NatlAcctCustomerUid            int32     `bun:"natl_acct_customer_uid,type:int,nullzero"`
	SlaRequestedDate               time.Time `bun:"sla_requested_date,type:datetime,nullzero"`
	EdiManualOverrideFlag          string    `bun:"edi_manual_override_flag,type:char,default:('N')"`
	RevisioningEnabledFlag         string    `bun:"revisioning_enabled_flag,type:char,nullzero"`
	QuoteSubmittedFlag             string    `bun:"quote_submitted_flag,type:char,nullzero"`
	WarrantyRmaFlag                string    `bun:"warranty_rma_flag,type:char,default:('N')"`
	GenerateProFormaInvoicesFlag   string    `bun:"generate_pro_forma_invoices_flag,type:char,nullzero"`
	AsbDeliveryMethodUid           int32     `bun:"asb_delivery_method_uid,type:int,nullzero"`
	DeliveryContactId              string    `bun:"delivery_contact_id,type:varchar(16),nullzero"`
	AdvancedBillingFlag            string    `bun:"advanced_billing_flag,type:char,nullzero"`
	AdvancedBillingPrintFlag       string    `bun:"advanced_billing_print_flag,type:char,nullzero"`
	EwingJobNo                     int32     `bun:"ewing_job_no,type:int,nullzero"`
	PricingLocId                   int32     `bun:"pricing_loc_id,type:int,nullzero"`
	RentalTransactionNo            string    `bun:"rental_transaction_no,type:varchar(255),nullzero"`
	RentalQuoteNo                  string    `bun:"rental_quote_no,type:varchar(255),nullzero"`
	RentalTransactionStatusCd      int32     `bun:"rental_transaction_status_cd,type:int,nullzero"`
	RentalReturnDate               time.Time `bun:"rental_return_date,type:datetime,nullzero"`
	MafSurchargeOverride           string    `bun:"maf_surcharge_override,type:varchar,default:('N')"`
	MafSurchargeReasonCode         string    `bun:"maf_surcharge_reason_code,type:varchar(255),nullzero"`
	OracleCarrierId                float64   `bun:"oracle_carrier_id,type:decimal(19,9),nullzero"`
	RentalBillingFlag              string    `bun:"rental_billing_flag,type:char,default:('U')"`
	InsideSales                    string    `bun:"inside_sales,type:varchar(30),nullzero"`
	Ship2Url                       string    `bun:"ship2_url,type:varchar(255),nullzero"`
	GlDimensionProjectNo           string    `bun:"gl_dimension_project_no,type:varchar(255),nullzero"`
	RentalStartDate                time.Time `bun:"rental_start_date,type:datetime,nullzero"`
}
