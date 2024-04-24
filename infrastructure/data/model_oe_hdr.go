package data

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun/schema"
	"time"

	"github.com/uptrace/bun"
)

type OeHdr struct {
	bun.BaseModel                  `bun:"table:oe_hdr"`
	OrderNo                        string          `bun:"order_no,type:varchar(8),pk"`
	CustomerId                     float64         `bun:"customer_id,type:decimal(19,0)"`
	OrderDate                      sql.NullTime    `bun:"order_date,type:datetime,nullzero"`
	Ship2Name                      sql.NullString  `bun:"ship2_name,type:varchar(50),nullzero"`
	Ship2Add1                      sql.NullString  `bun:"ship2_add1,type:varchar(50),nullzero"`
	Ship2Add2                      sql.NullString  `bun:"ship2_add2,type:varchar(50),nullzero"`
	Ship2City                      sql.NullString  `bun:"ship2_city,type:varchar(50),nullzero"`
	Ship2State                     sql.NullString  `bun:"ship2_state,type:varchar(50),nullzero"`
	Ship2Zip                       sql.NullString  `bun:"ship2_zip,type:varchar(10),nullzero"`
	Ship2Country                   sql.NullString  `bun:"ship2_country,type:varchar(50),nullzero"`
	RequestedDate                  sql.NullTime    `bun:"requested_date,type:datetime,nullzero"`
	PoNo                           sql.NullString  `bun:"po_no,type:varchar(50),nullzero"`
	Terms                          sql.NullString  `bun:"terms,type:varchar(2),nullzero"`
	ShipToPhone                    sql.NullString  `bun:"ship_to_phone,type:varchar(20),nullzero"`
	DeleteFlag                     string          `bun:"delete_flag,type:char"`
	Completed                      sql.NullString  `bun:"completed,type:char,nullzero"`
	CompanyId                      sql.NullString  `bun:"company_id,type:varchar(8),nullzero"`
	DateCreated                    time.Time       `bun:"date_created,type:datetime"`
	DateLastModified               time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy               string          `bun:"last_maintained_by,type:varchar(30)"`
	CodFlag                        sql.NullString  `bun:"cod_flag,type:char,nullzero"`
	GrossMargin                    sql.NullFloat64 `bun:"gross_margin,type:decimal(19,9),nullzero"`
	ProjectedOrder                 sql.NullString  `bun:"projected_order,type:char,nullzero"`
	PoNoAppend                     sql.NullString  `bun:"po_no_append,type:varchar(20),nullzero"`
	LocationId                     sql.NullFloat64 `bun:"location_id,type:decimal(19,0),nullzero"`
	CarrierId                      sql.NullFloat64 `bun:"carrier_id,type:decimal(19,0),nullzero"`
	AddressId                      sql.NullFloat64 `bun:"address_id,type:decimal(19,0),nullzero"`
	ContactId                      sql.NullString  `bun:"contact_id,type:varchar(16),nullzero"`
	CorpAddressId                  sql.NullFloat64 `bun:"corp_address_id,type:decimal(19,0),nullzero"`
	HandlingChargeReqFlag          sql.NullString  `bun:"handling_charge_req_flag,type:char,nullzero"`
	PaymentMethod                  sql.NullString  `bun:"payment_method,type:varchar(8),nullzero"`
	FobFlag                        sql.NullString  `bun:"fob_flag,type:char,nullzero"`
	Class1id                       sql.NullString  `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                       sql.NullString  `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                       sql.NullString  `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                       sql.NullString  `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                       sql.NullString  `bun:"class_5id,type:varchar(255),nullzero"`
	RmaFlag                        sql.NullString  `bun:"rma_flag,type:char,nullzero"`
	Taker                          sql.NullString  `bun:"taker,type:varchar(30),nullzero"`
	JobName                        sql.NullString  `bun:"job_name,type:varchar(40),nullzero"`
	ThirdPartyBillingFlag          sql.NullString  `bun:"third_party_billing_flag,type:char,nullzero"`
	Approved                       sql.NullString  `bun:"approved,type:char,nullzero"`
	SourceLocationId               sql.NullFloat64 `bun:"source_location_id,type:decimal(19,0),nullzero"`
	PackingBasis                   sql.NullString  `bun:"packing_basis,type:varchar(16),nullzero"`
	DeliveryInstructions           sql.NullString  `bun:"delivery_instructions,type:varchar(255),nullzero"`
	PickTicketType                 sql.NullString  `bun:"pick_ticket_type,type:varchar(2),nullzero"`
	RequestedDownpayment           sql.NullFloat64 `bun:"requested_downpayment,type:decimal(19,4),nullzero"`
	DownpaymentInvoiced            sql.NullString  `bun:"downpayment_invoiced,type:char,nullzero"`
	CancelFlag                     sql.NullString  `bun:"cancel_flag,type:char,nullzero"`
	WillCall                       sql.NullString  `bun:"will_call,type:char,nullzero"`
	FrontCounter                   sql.NullString  `bun:"front_counter,type:char,nullzero"`
	ValidationStatus               sql.NullString  `bun:"validation_status,type:varchar(8),nullzero"`
	OeHdrUid                       int32           `bun:"oe_hdr_uid,type:int"`
	SourceId                       sql.NullString  `bun:"source_id,type:varchar(8),nullzero"`
	SourceCodeNo                   int32           `bun:"source_code_no,type:int"`
	CreditCardHold                 sql.NullInt32   `bun:"credit_card_hold,type:int,nullzero"`
	InvoiceBatchUid                int32           `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid                 sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	FreightOut                     float64         `bun:"freight_out,type:decimal(19,4),default:(0)"`
	ShippingRouteUid               sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	ExcludeRebates                 string          `bun:"exclude_rebates,type:char"`
	CaptureUsageDefault            string          `bun:"capture_usage_default,type:char"`
	JobPriceHdrUid                 sql.NullInt32   `bun:"job_price_hdr_uid,type:int,nullzero"`
	FrontCounterRma                string          `bun:"front_counter_rma,type:varchar,default:('N')"`
	Ship2EmailAddress              sql.NullString  `bun:"ship2_email_address,type:varchar(255),nullzero"`
	OrderType                      sql.NullInt32   `bun:"order_type,type:int,nullzero"`
	Taxable                        sql.NullString  `bun:"taxable,type:char,nullzero"`
	ProfitPercent                  float64         `bun:"profit_percent,type:decimal(19,9),default:(0.00)"`
	OrderCostBasis                 int32           `bun:"order_cost_basis,type:int,default:(1)"`
	CurrencyLineUid                sql.NullInt32   `bun:"currency_line_uid,type:int,nullzero"`
	InvoiceExchRateSourceCd        sql.NullInt32   `bun:"invoice_exch_rate_source_cd,type:int,nullzero"`
	RmaExpirationDate              sql.NullTime    `bun:"rma_expiration_date,type:datetime,default:(getdate())"`
	CreatedBy                      sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	InvoiceNo                      sql.NullString  `bun:"invoice_no,type:varchar(10),nullzero"`
	BillToContactId                sql.NullString  `bun:"bill_to_contact_id,type:varchar(16),nullzero"`
	TagHoldCancelDate              sql.NullTime    `bun:"tag_hold_cancel_date,type:datetime,nullzero"`
	CostCenterTrackingOption       sql.NullInt32   `bun:"cost_center_tracking_option,type:int,nullzero"`
	SkipProfitExceptionCheck       string          `bun:"skip_profit_exception_check,type:char,default:('N')"`
	ConsBackorderProcessingFlag    sql.NullString  `bun:"cons_backorder_processing_flag,type:char,nullzero"`
	RestockFeePercentage           sql.NullFloat64 `bun:"restock_fee_percentage,type:decimal(19,9),nullzero"`
	DateOrderCompleted             sql.NullTime    `bun:"date_order_completed,type:datetime,nullzero"`
	ValidatedViaOpenOrdersFlag     sql.NullString  `bun:"validated_via_open_orders_flag,type:char,default:('N')"`
	AcknowledgementDate            sql.NullTime    `bun:"acknowledgement_date,type:datetime,nullzero"`
	ProductGroupCostBasis          sql.NullInt32   `bun:"product_group_cost_basis,type:int,nullzero"`
	OriginalPromiseDate            sql.NullTime    `bun:"original_promise_date,type:datetime,nullzero"`
	PromiseDate                    sql.NullTime    `bun:"promise_date,type:datetime,nullzero"`
	RequestedShipDate              sql.NullTime    `bun:"requested_ship_date,type:datetime,nullzero"`
	ExpectedCompletionDate         sql.NullTime    `bun:"expected_completion_date,type:datetime,nullzero"`
	JobControlFlag                 sql.NullString  `bun:"job_control_flag,type:char,nullzero"`
	ApplyBuilderAllowanceFlag      string          `bun:"apply_builder_allowance_flag,type:char,default:('N')"`
	MerchandiseCreditFlag          string          `bun:"merchandise_credit_flag,type:char,default:('N')"`
	ReqPymtUponReleaseFlag         string          `bun:"req_pymt_upon_release_flag,type:char,default:('N')"`
	PmDate                         sql.NullTime    `bun:"pm_date,type:datetime,nullzero"`
	DownpaymentPercentage          sql.NullFloat64 `bun:"downpayment_percentage,type:decimal(19,2),nullzero"`
	PrepaidInvoiceFlag             sql.NullString  `bun:"prepaid_invoice_flag,type:char,nullzero"`
	LandedCostIncludedCd           sql.NullInt32   `bun:"landed_cost_included_cd,type:int,nullzero"`
	ExportedFlag                   sql.NullString  `bun:"exported_flag,type:char,nullzero"`
	WebReferenceNo                 sql.NullString  `bun:"web_reference_no,type:varchar(255),nullzero"`
	UpsCode                        sql.NullString  `bun:"ups_code,type:varchar(40),nullzero"`
	DefaultPricingCd               sql.NullInt32   `bun:"default_pricing_cd,type:int,nullzero"`
	FreightChargeByMileHdrUid      sql.NullInt32   `bun:"freight_charge_by_mile_hdr_uid,type:int,nullzero"`
	FreightMileageAmt              sql.NullFloat64 `bun:"freight_mileage_amt,type:decimal(19,2),nullzero"`
	SupplierOrderNo                sql.NullString  `bun:"supplier_order_no,type:varchar(255),nullzero"`
	SupplierReleaseNo              sql.NullString  `bun:"supplier_release_no,type:varchar(255),nullzero"`
	RoutedEtaDate                  sql.NullTime    `bun:"routed_eta_date,type:datetime,nullzero"`
	RouteOverrideDate              sql.NullTime    `bun:"route_override_date,type:datetime,nullzero"`
	ExpediteDate                   sql.NullTime    `bun:"expedite_date,type:datetime,nullzero"`
	OrderOpenStartDate             sql.NullTime    `bun:"order_open_start_date,type:datetime,nullzero"`
	EnvironmentalFeeFlag           sql.NullString  `bun:"environmental_fee_flag,type:char,default:('N')"`
	AdminFeeFlag                   sql.NullString  `bun:"admin_fee_flag,type:char,default:('N')"`
	PromiseDateExtendedDesc        sql.NullString  `bun:"promise_date_extended_desc,type:varchar(8000),nullzero"`
	PromiseDateEditedDate          sql.NullTime    `bun:"promise_date_edited_date,type:datetime,nullzero"`
	OrderDiscType                  sql.NullInt32   `bun:"order_disc_type,type:int,nullzero"`
	OrderDiscFactor                sql.NullFloat64 `bun:"order_disc_factor,type:decimal(19,9),nullzero"`
	OrderAckPrintedFlag            sql.NullString  `bun:"order_ack_printed_flag,type:char,nullzero"`
	PackingListSentFlag            sql.NullInt32   `bun:"packing_list_sent_flag,type:int,nullzero"`
	RmaDeliveryListStatus          sql.NullInt32   `bun:"rma_delivery_list_status,type:int,default:((1024))"`
	StrategicLibraryUid            sql.NullInt32   `bun:"strategic_library_uid,type:int,nullzero"`
	StrategicLibraryOriginalUid    sql.NullInt32   `bun:"strategic_library_original_uid,type:int,nullzero"`
	SendPartialOrderFlag           sql.NullString  `bun:"send_partial_order_flag,type:char,default:('N')"`
	SupplierStatus                 sql.NullString  `bun:"supplier_status,type:char,nullzero"`
	OrderPriorityUid               sql.NullInt32   `bun:"order_priority_uid,type:int,nullzero"`
	B2bUpsFreightAmount            sql.NullFloat64 `bun:"b2b_ups_freight_amount,type:decimal(19,9),nullzero"`
	UserDefinedDate                sql.NullTime    `bun:"user_defined_date,type:datetime,nullzero"`
	BlindAddressingFlag            sql.NullString  `bun:"blind_addressing_flag,type:char,default:('N')"`
	ReplaceCompanyNameFlag         sql.NullString  `bun:"replace_company_name_flag,type:char,default:('N')"`
	UseVendorItemTermsFlag         sql.NullString  `bun:"use_vendor_item_terms_flag,type:char,nullzero"`
	FirstPackingListNo             sql.NullFloat64 `bun:"first_packing_list_no,type:decimal(19,0),nullzero"`
	SalesMarketGroupUid            sql.NullInt32   `bun:"sales_market_group_uid,type:int,nullzero"`
	ShipConfirmedFlag              sql.NullString  `bun:"ship_confirmed_flag,type:char,nullzero"`
	PriceConfirmedFlag             sql.NullString  `bun:"price_confirmed_flag,type:char,nullzero"`
	QuotedFreightOut               sql.NullFloat64 `bun:"quoted_freight_out,type:decimal(19,9),nullzero"`
	GeocomAppend                   sql.NullInt32   `bun:"geocom_append,type:int,nullzero"`
	ApplyFuelSurchargeFlag         sql.NullString  `bun:"apply_fuel_surcharge_flag,type:char,nullzero"`
	OverrideFreightCodeFlag        sql.NullString  `bun:"override_freight_code_flag,type:char,nullzero"`
	ServiceOrderPriorityUid        sql.NullInt32   `bun:"service_order_priority_uid,type:int,nullzero"`
	ExcludeFromCreditLimitFlag     sql.NullString  `bun:"exclude_from_credit_limit_flag,type:char,nullzero"`
	PackingListFilename            sql.NullString  `bun:"packing_list_filename,type:char,nullzero"`
	OverrideMinOrderChargeFlag     sql.NullString  `bun:"override_min_order_charge_flag,type:char,nullzero"`
	PlacedByName                   sql.NullString  `bun:"placed_by_name,type:varchar(255),nullzero"`
	ApprovedForArFlag              sql.NullString  `bun:"approved_for_ar_flag,type:char,nullzero"`
	SecondRouteOverrideDate        sql.NullTime    `bun:"second_route_override_date,type:datetime,nullzero"`
	Subject                        sql.NullString  `bun:"subject,type:varchar(255),nullzero"`
	ServicePlanUid                 sql.NullInt32   `bun:"service_plan_uid,type:int,nullzero"`
	CentralFaxNumber               sql.NullString  `bun:"central_fax_number,type:varchar(20),nullzero"`
	Url                            sql.NullString  `bun:"url,type:varchar(255),nullzero"`
	OriginalPackingBasis           sql.NullString  `bun:"original_packing_basis,type:varchar(16),nullzero"`
	FreightOutEditedFlag           sql.NullString  `bun:"freight_out_edited_flag,type:char,nullzero"`
	OverrideContactEmailFlag       sql.NullString  `bun:"override_contact_email_flag,type:char,default:('N')"`
	OverrideEmailAddress           sql.NullString  `bun:"override_email_address,type:varchar(255),nullzero"`
	PrintPricesOnPackinglist       sql.NullString  `bun:"print_prices_on_packinglist,type:char,nullzero"`
	ImportSource                   sql.NullString  `bun:"import_source,type:varchar(255),nullzero"`
	QuoteLockedFlag                sql.NullString  `bun:"quote_locked_flag,type:char,nullzero"`
	WebShopperId                   sql.NullInt32   `bun:"web_shopper_id,type:int,nullzero"`
	WebShopperEmail                sql.NullString  `bun:"web_shopper_email,type:varchar(255),nullzero"`
	RecalcScheduledDsPrice         sql.NullString  `bun:"recalc_scheduled_ds_price,type:char,nullzero"`
	OrderTypeCust                  sql.NullString  `bun:"order_type_cust,type:varchar(255),nullzero"`
	EngineerId                     sql.NullString  `bun:"engineer_id,type:varchar(255),nullzero"`
	CarrierContractHdrUid          sql.NullInt32   `bun:"carrier_contract_hdr_uid,type:int,nullzero"`
	OrderAckPrintPricesFlag        sql.NullString  `bun:"order_ack_print_prices_flag,type:char,nullzero"`
	FreightOutEstimate             sql.NullFloat64 `bun:"freight_out_estimate,type:decimal(19,9),nullzero"`
	FreightChargeEstimate          sql.NullFloat64 `bun:"freight_charge_estimate,type:decimal(19,9),nullzero"`
	DocumentCaptureDate            sql.NullTime    `bun:"document_capture_date,type:datetime,nullzero"`
	LimitMaxShipmentsPerOrder      sql.NullFloat64 `bun:"limit_max_shipments_per_order,type:decimal(19,4),nullzero"`
	ReasonCreditMemoCodeUid        sql.NullInt32   `bun:"reason_credit_memo_code_uid,type:int,nullzero"`
	SourceCreditMemoCodeUid        sql.NullInt32   `bun:"source_credit_memo_code_uid,type:int,nullzero"`
	NetBillingFlag                 string          `bun:"net_billing_flag,type:char,default:('N')"`
	NetBillingEdited               string          `bun:"net_billing_edited,type:char,default:('N')"`
	SchedOrderDiscFlag             sql.NullString  `bun:"sched_order_disc_flag,type:char,nullzero"`
	WillCallDiscFlag               sql.NullString  `bun:"will_call_disc_flag,type:char,nullzero"`
	SingleOrderDiscFlag            sql.NullString  `bun:"single_order_disc_flag,type:char,nullzero"`
	VolumeDiscFlag                 sql.NullString  `bun:"volume_disc_flag,type:char,nullzero"`
	ServiceInvoiceNo               sql.NullString  `bun:"service_invoice_no,type:varchar(10),nullzero"`
	Ship2Add3                      sql.NullString  `bun:"ship2_add3,type:varchar(50),nullzero"`
	BillToId                       sql.NullFloat64 `bun:"bill_to_id,type:decimal(19,0),nullzero"`
	ElectronicOrderFlag            sql.NullString  `bun:"electronic_order_flag,type:char,default:('N')"`
	CodWithoutRemittanceFlag       sql.NullString  `bun:"cod_without_remittance_flag,type:char,nullzero"`
	SoTotalPartsPriceOnInvoice     sql.NullString  `bun:"so_total_parts_price_on_invoice,type:char,nullzero"`
	SoTotalLaborPriceOnInvoice     sql.NullString  `bun:"so_total_labor_price_on_invoice,type:char,nullzero"`
	InquiryFlag                    sql.NullString  `bun:"inquiry_flag,type:char,nullzero"`
	HoldInvoiceFlag                string          `bun:"hold_invoice_flag,type:char,default:('N')"`
	DoNotExportToPtsFlag           sql.NullString  `bun:"do_not_export_to_pts_flag,type:char,default:('N')"`
	GenericDescOnPackinglistFlag   sql.NullString  `bun:"generic_desc_on_packinglist_flag,type:char,nullzero"`
	GenericDescOnInvoiceAckFlag    sql.NullString  `bun:"generic_desc_on_invoice_ack_flag,type:char,nullzero"`
	WillCallNotificationFlag       sql.NullString  `bun:"will_call_notification_flag,type:char,nullzero"`
	PrebillingDate                 sql.NullTime    `bun:"prebilling_date,type:datetime,nullzero"`
	BlindShipFlag                  sql.NullString  `bun:"blind_ship_flag,type:char,nullzero"`
	ArchitectId                    sql.NullFloat64 `bun:"architect_id,type:decimal(19,0),nullzero"`
	BuilderId                      sql.NullFloat64 `bun:"builder_id,type:decimal(19,0),nullzero"`
	ContractorId                   sql.NullFloat64 `bun:"contractor_id,type:decimal(19,0),nullzero"`
	DesignerId                     sql.NullFloat64 `bun:"designer_id,type:decimal(19,0),nullzero"`
	HomeownerId                    sql.NullFloat64 `bun:"homeowner_id,type:decimal(19,0),nullzero"`
	QuoteType                      sql.NullInt32   `bun:"quote_type,type:int,nullzero"`
	PtsLabelPrintFlag              sql.NullString  `bun:"pts_label_print_flag,type:char,nullzero"`
	DiscountItemId                 sql.NullString  `bun:"discount_item_id,type:varchar(255),nullzero"`
	CustomerDiscPct                sql.NullFloat64 `bun:"customer_disc_pct,type:decimal(19,9),nullzero"`
	ConsolidatePerOrderFlag        sql.NullString  `bun:"consolidate_per_order_flag,type:char,nullzero"`
	DfltUpsShipViaCode             sql.NullInt32   `bun:"dflt_ups_ship_via_code,type:int,nullzero"`
	DfltFedexServiceTypeUid        sql.NullInt32   `bun:"dflt_fedex_service_type_uid,type:int,nullzero"`
	InvoiceOnlyFlag                sql.NullString  `bun:"invoice_only_flag,type:char,nullzero"`
	Ship2Latitude                  sql.NullString  `bun:"ship2_latitude,type:varchar(20),nullzero"`
	Ship2Longitude                 sql.NullString  `bun:"ship2_longitude,type:varchar(20),nullzero"`
	BypassHandlingChargeFlag       sql.NullString  `bun:"bypass_handling_charge_flag,type:char,nullzero"`
	FreightChargeUid               sql.NullInt32   `bun:"freight_charge_uid,type:int,nullzero"`
	LumpSumBillingFlag             sql.NullString  `bun:"lump_sum_billing_flag,type:char,nullzero"`
	FieldDestroyFlag               sql.NullString  `bun:"field_destroy_flag,type:char,nullzero"`
	WarrantyReturnOrderFlag        sql.NullString  `bun:"warranty_return_order_flag,type:char,default:('N')"`
	AllowAutoApplyOrigInvFlag      sql.NullString  `bun:"allow_auto_apply_orig_inv_flag,type:char,nullzero"`
	DaysEarly                      sql.NullInt32   `bun:"days_early,type:int,nullzero"`
	TransitDays                    sql.NullInt32   `bun:"transit_days,type:int,nullzero"`
	DisplayInvoiceExchRateSourceCd sql.NullInt32   `bun:"display_invoice_exch_rate_source_cd,type:int,nullzero"`
	NatlAcctCustomerUid            sql.NullInt32   `bun:"natl_acct_customer_uid,type:int,nullzero"`
	SlaRequestedDate               sql.NullTime    `bun:"sla_requested_date,type:datetime,nullzero"`
	EdiManualOverrideFlag          sql.NullString  `bun:"edi_manual_override_flag,type:char,default:('N')"`
	RevisioningEnabledFlag         sql.NullString  `bun:"revisioning_enabled_flag,type:char,nullzero"`
	QuoteSubmittedFlag             sql.NullString  `bun:"quote_submitted_flag,type:char,nullzero"`
	WarrantyRmaFlag                sql.NullString  `bun:"warranty_rma_flag,type:char,default:('N')"`
	GenerateProFormaInvoicesFlag   sql.NullString  `bun:"generate_pro_forma_invoices_flag,type:char,nullzero"`
	AsbDeliveryMethodUid           sql.NullInt32   `bun:"asb_delivery_method_uid,type:int,nullzero"`
	DeliveryContactId              sql.NullString  `bun:"delivery_contact_id,type:varchar(16),nullzero"`
	AdvancedBillingFlag            sql.NullString  `bun:"advanced_billing_flag,type:char,nullzero"`
	AdvancedBillingPrintFlag       sql.NullString  `bun:"advanced_billing_print_flag,type:char,nullzero"`
	EwingJobNo                     sql.NullInt32   `bun:"ewing_job_no,type:int,nullzero"`
	PricingLocId                   sql.NullInt32   `bun:"pricing_loc_id,type:int,nullzero"`
	RentalTransactionNo            sql.NullString  `bun:"rental_transaction_no,type:varchar(255),nullzero"`
	RentalQuoteNo                  sql.NullString  `bun:"rental_quote_no,type:varchar(255),nullzero"`
	RentalTransactionStatusCd      sql.NullInt32   `bun:"rental_transaction_status_cd,type:int,nullzero"`
	RentalReturnDate               sql.NullTime    `bun:"rental_return_date,type:datetime,nullzero"`
	MafSurchargeOverride           string          `bun:"maf_surcharge_override,type:varchar,default:('N')"`
	MafSurchargeReasonCode         sql.NullString  `bun:"maf_surcharge_reason_code,type:varchar(255),nullzero"`
	OracleCarrierId                sql.NullFloat64 `bun:"oracle_carrier_id,type:decimal(19,9),nullzero"`
	RentalBillingFlag              sql.NullString  `bun:"rental_billing_flag,type:char,default:('U')"`
	InsideSales                    sql.NullString  `bun:"inside_sales,type:varchar(30),nullzero"`
	Ship2Url                       sql.NullString  `bun:"ship2_url,type:varchar(255),nullzero"`
	GlDimensionProjectNo           sql.NullString  `bun:"gl_dimension_project_no,type:varchar(255),nullzero"`
	RentalStartDate                sql.NullTime    `bun:"rental_start_date,type:datetime,nullzero"`

	OeLines []OeLine `bun:"rel:has-many,join:order_no=order_no"`
	Contact Contacts `bun:"rel:has-one,join:contact_id=id"`
}

var _ bun.BeforeAppendModelHook = (*OeHdr)(nil)

func (m *OeHdr) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OeHdrModel struct {
	bun bun.IDB
}

type CreateOeHdrParams struct {
	CustomerId     float64
	Ship2Name      string
	Ship2Add1      string
	Ship2Add2      string
	Ship2City      string
	Ship2State     string
	Ship2Zip       string
	Ship2Country   string
	ShipToPhone    string
	PoNo           string
	ProjectedOrder string
}

// Create creates a new order.
func (m *OeHdrModel) Create(
	ctx context.Context, params CreateOeHdrParams) (*OeHdr, error) {
	oeHdr := &OeHdr{
		CustomerId:     params.CustomerId,
		Ship2Name:      sql.NullString{String: params.Ship2Name, Valid: true},
		Ship2Add1:      sql.NullString{String: params.Ship2Add1, Valid: true},
		Ship2Add2:      sql.NullString{String: params.Ship2Add2, Valid: true},
		Ship2City:      sql.NullString{String: params.Ship2City, Valid: true},
		Ship2State:     sql.NullString{String: params.Ship2State, Valid: true},
		Ship2Zip:       sql.NullString{String: params.Ship2Zip, Valid: true},
		Ship2Country:   sql.NullString{String: params.Ship2Country, Valid: true},
		ProjectedOrder: sql.NullString{String: params.ProjectedOrder, Valid: true},
		PoNo:           sql.NullString{String: params.PoNo, Valid: true},
		ShipToPhone:    sql.NullString{String: params.ShipToPhone, Valid: true},
		ContactId:      sql.NullString{String: "2563", Valid: true},
		Taker:          sql.NullString{String: "admin", Valid: true},

		PromiseDate:           sql.NullTime{Time: time.Now(), Valid: true},
		OrderDate:             sql.NullTime{Time: time.Now(), Valid: true},
		RequestedDate:         sql.NullTime{Time: time.Now(), Valid: true},
		ValidationStatus:      sql.NullString{String: "OK", Valid: true},
		Terms:                 sql.NullString{String: "30", Valid: true},
		DeleteFlag:            "N",
		SourceCodeNo:          920,
		Completed:             sql.NullString{String: "N", Valid: true},
		CompanyId:             sql.NullString{String: "MRS", Valid: true},
		LocationId:            sql.NullFloat64{Float64: 1001, Valid: true},
		CarrierId:             sql.NullFloat64{Float64: 100008, Valid: true},
		AddressId:             sql.NullFloat64{Float64: 1001, Valid: true},
		FobFlag:               sql.NullString{String: "F", Valid: true},
		RmaFlag:               sql.NullString{String: "N", Valid: true},
		ThirdPartyBillingFlag: sql.NullString{String: "S", Valid: true},
		Approved:              sql.NullString{String: "N", Valid: true},
		SourceLocationId:      sql.NullFloat64{Float64: 1001, Valid: true},
		PackingBasis:          sql.NullString{String: "Partial", Valid: true},
		DeliveryInstructions:  sql.NullString{String: "Leave on porch", Valid: true},
		PickTicketType:        sql.NullString{String: "UT", Valid: true},
		CancelFlag:            sql.NullString{String: "N", Valid: true},
		FrontCounter:          sql.NullString{String: "N", Valid: true},
	}

	err := m.generateOeHdrUid(ctx, &oeHdr.OeHdrUid)
	err = m.generateOrderNo(ctx, &oeHdr.OrderNo)
	if err != nil {
		return nil, err
	}
	_, err = m.bun.NewInsert().Model(oeHdr).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return oeHdr, nil
}

func (m *OeHdrModel) Insert(ctx context.Context, oeHdr *OeHdr) error {
	return nil
}

// Get returns the order by the given order number.
func (m *OeHdrModel) Get(ctx context.Context, orderNo string) (*OeHdr, error) {
	var oeHdr OeHdr
	err := m.bun.NewSelect().Model(&oeHdr).Relation("OeLines").Where("order_no = ?", orderNo).Scan(ctx)
	if err != nil {
		return nil, err

	}
	return &oeHdr, nil
}

// Update updates the given order.
func (m *OeHdrModel) Update(ctx context.Context, oeHdr *OeHdr) error {
	return nil
}

// Delete soft deletes the order by the given order number.
func (m *OeHdrModel) Delete(ctx context.Context, orderNo string) error {
	return nil
}

// GetByCustomerId returns all orders for a given customer Id.
func (m *OeHdrModel) GetByCustomerId(ctx context.Context, customerId float64) ([]*OeHdr, error) {
	var oeHdrs []*OeHdr
	count, err := m.bun.NewSelect().Model(&oeHdrs).Relation("OeLines").Where(
		"customer_id = ?", customerId).ScanAndCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, ErrRecordNotFound
	}

	return oeHdrs, nil
}

// generateOeHdrUid generates a new OeJdrUid for oeHdr.
func (m *OeHdrModel) generateOeHdrUid(ctx context.Context, uid *int32) error {

	query := `DECLARE @uid int
			EXEC @uid = p21_get_counter 'oe_hdr', 1
			SELECT @uid`

	err := m.bun.QueryRowContext(ctx, query).Scan(uid)
	if err != nil {
		return err
	}
	return nil
}

func (m *OeHdrModel) generateOrderNo(ctx context.Context, orderNo *string) error {
	query := `DECLARE @id varchar(8)
			EXEC @id = p21_get_counter 'WO', 1
			SELECT @id`

	err := m.bun.QueryRowContext(ctx, query).Scan(orderNo)
	if err != nil {
		return err
	}
	return nil
}
