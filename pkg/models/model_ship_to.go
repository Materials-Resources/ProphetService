package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"time"
)

type ShipTo struct {
	bun.BaseModel                 `bun:"table:ship_to"`
	CompanyId                     string          `bun:"company_id,type:varchar(8),pk"`
	ShipToId                      float64         `bun:"ship_to_id,type:decimal(19,0),pk"`
	CustomerId                    float64         `bun:"customer_id,type:decimal(19,0)"`
	DefaultBranch                 string          `bun:"default_branch,type:varchar(8)"`
	FederalExemptionNumber        sql.NullString  `bun:"federal_exemption_number,type:varchar(40),nullzero"`
	StateExemptionNumber          sql.NullString  `bun:"state_exemption_number,type:varchar(40),nullzero"`
	OtherExemptionNumber          sql.NullString  `bun:"other_exemption_number,type:varchar(40),nullzero"`
	PriceFileId                   sql.NullFloat64 `bun:"price_file_id,type:decimal(19,0),nullzero"`
	PreferredLocationId           sql.NullFloat64 `bun:"preferred_location_id,type:decimal(19,0),nullzero"`
	DefaultShipTime               sql.NullTime    `bun:"default_ship_time,type:datetime,nullzero"`
	DefaultCarrierId              sql.NullFloat64 `bun:"default_carrier_id,type:decimal(19,0),nullzero"`
	Fob                           sql.NullString  `bun:"fob,type:varchar(20),nullzero"`
	DeliveryInstructions          sql.NullString  `bun:"delivery_instructions,type:varchar(255),nullzero"`
	MorningBegDelivery            sql.NullTime    `bun:"morning_beg_delivery,type:datetime,nullzero"`
	MorningEndDelivery            sql.NullTime    `bun:"morning_end_delivery,type:datetime,nullzero"`
	EveningBegDelivery            sql.NullTime    `bun:"evening_beg_delivery,type:datetime,nullzero"`
	EveningEndDelivery            sql.NullTime    `bun:"evening_end_delivery,type:datetime,nullzero"`
	AcceptPartialOrders           string          `bun:"accept_partial_orders,type:char"`
	AcceptableWaitTime            sql.NullFloat64 `bun:"acceptable_wait_time,type:decimal(3,0),nullzero"`
	Class1Id                      sql.NullString  `bun:"class1_id,type:varchar(255),nullzero"`
	Class2Id                      sql.NullString  `bun:"class2_id,type:varchar(255),nullzero"`
	Class3Id                      sql.NullString  `bun:"class3_id,type:varchar(255),nullzero"`
	Class4Id                      sql.NullString  `bun:"class4_id,type:varchar(255),nullzero"`
	Class5Id                      sql.NullString  `bun:"class5_id,type:varchar(255),nullzero"`
	PackingBasis                  sql.NullString  `bun:"packing_basis,type:varchar(16),nullzero"`
	DateCreated                   time.Time       `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified              time.Time       `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy              string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DeleteFlag                    string          `bun:"delete_flag,type:char,default:('N')"`
	TermsId                       sql.NullString  `bun:"terms_id,type:varchar(2),nullzero"`
	InvoiceType                   sql.NullString  `bun:"invoice_type,type:varchar(2),nullzero"`
	BilledOnGrossNetQty           sql.NullString  `bun:"billed_on_gross_net_qty,type:char,nullzero"`
	StateExciseTaxExemptionNo     sql.NullString  `bun:"state_excise_tax_exemption_no,type:varchar(40),nullzero"`
	PickTicketType                sql.NullString  `bun:"pick_ticket_type,type:varchar(2),nullzero"`
	TaxGroupId                    sql.NullString  `bun:"tax_group_id,type:varchar(10),nullzero"`
	HandlingChargeReqFlag         string          `bun:"handling_charge_req_flag,type:char,default:('N')"`
	ThirdPartyBillingFlag         string          `bun:"third_party_billing_flag,type:char,default:('S')"`
	DaysEarly                     sql.NullInt32   `bun:"days_early,type:int,default:(0)"`
	DaysLate                      sql.NullInt32   `bun:"days_late,type:int,default:(0)"`
	TransitDays                   sql.NullInt32   `bun:"transit_days,type:int,default:(1)"`
	FreightCodeUid                sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	SignatureRequired             string          `bun:"signature_required,type:char,default:('N')"`
	ShippingRouteUid              sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	PdaOrderEntry                 string          `bun:"pda_order_entry,type:char,default:('N')"`
	PdaOelistCriteriaUid          sql.NullInt32   `bun:"pda_oelist_criteria_uid,type:int,nullzero"`
	ExcludeHoldFromPickTix        string          `bun:"exclude_hold_from_pick_tix,type:varchar,default:('N')"`
	ExcludeCanceldFromPackList    string          `bun:"exclude_canceld_from_pack_list,type:varchar,default:('N')"`
	ExcludeHoldFromPackList       string          `bun:"exclude_hold_from_pack_list,type:varchar,default:('N')"`
	ExcludeCanceldFromOrderAck    string          `bun:"exclude_canceld_from_order_ack,type:varchar,default:('N')"`
	ExcludeHoldFromOrderAck       string          `bun:"exclude_hold_from_order_ack,type:varchar,default:('N')"`
	IncludeNonAllocOnPickTix      int32           `bun:"include_non_alloc_on_pick_tix,type:int,default:(1277)"`
	ExcludeCanceldFromPickTix     string          `bun:"exclude_canceld_from_pick_tix,type:char,default:('N')"`
	CreatedBy                     sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	IncludeNonAllocOnPackList     int32           `bun:"include_non_alloc_on_pack_list,type:int"`
	PrintPackinglistInShipping    string          `bun:"print_packinglist_in_shipping,type:char,default:('N')"`
	PrintPricesOnPackinglist      string          `bun:"print_prices_on_packinglist,type:char,default:('Y')"`
	SourceTypeCd                  sql.NullInt32   `bun:"source_type_cd,type:int,nullzero"`
	DataIdentifierGroupUid        sql.NullInt32   `bun:"data_identifier_group_uid,type:int,nullzero"`
	FedexFreightMarkup            sql.NullFloat64 `bun:"fedex_freight_markup,type:decimal(19,9),nullzero"`
	CreditStatus                  sql.NullString  `bun:"credit_status,type:varchar(8),nullzero"`
	CreditLimit                   sql.NullFloat64 `bun:"credit_limit,type:decimal(19,2),nullzero"`
	CreditLimitUsed               sql.NullFloat64 `bun:"credit_limit_used,type:decimal(19,2),nullzero"`
	DefaultCustomerPoNo           sql.NullString  `bun:"default_customer_po_no,type:varchar(255),nullzero"`
	BudgetCodeApprovalFlag        sql.NullString  `bun:"budget_code_approval_flag,type:char,nullzero"`
	ServiceTermsId                sql.NullString  `bun:"service_terms_id,type:varchar(2),nullzero"`
	DefaultFreightMarkupPct       float64         `bun:"default_freight_markup_pct,type:decimal(19,4),default:((0))"`
	FreightChargeByMileHdrUid     sql.NullInt32   `bun:"freight_charge_by_mile_hdr_uid,type:int,nullzero"`
	FreightMileageAmt             sql.NullFloat64 `bun:"freight_mileage_amt,type:decimal(19,2),nullzero"`
	DegreeDaysFlag                sql.NullString  `bun:"degree_days_flag,type:char,nullzero"`
	CalendarDaysFlag              sql.NullString  `bun:"calendar_days_flag,type:char,nullzero"`
	HeatAndHotWaterCd             sql.NullInt32   `bun:"heat_and_hot_water_cd,type:int,nullzero"`
	DrumDepositExemptFlag         sql.NullString  `bun:"drum_deposit_exempt_flag,type:char,default:('N')"`
	TaxableFlag                   sql.NullString  `bun:"taxable_flag,type:char,nullzero"`
	BillHoldFlag                  sql.NullString  `bun:"bill_hold_flag,type:char,nullzero"`
	DisplayPoNoFlag               sql.NullString  `bun:"display_po_no_flag,type:char,nullzero"`
	PoNoRequiredFlag              sql.NullString  `bun:"po_no_required_flag,type:char,nullzero"`
	DisplayBadgeFlag              sql.NullString  `bun:"display_badge_flag,type:char,nullzero"`
	BadgeRequiredFlag             sql.NullString  `bun:"badge_required_flag,type:char,nullzero"`
	DelSchSourceLocation          sql.NullFloat64 `bun:"del_sch_source_location,type:decimal(19,0),nullzero"`
	ServiceSourceLocation         sql.NullFloat64 `bun:"service_source_location,type:decimal(19,0),nullzero"`
	DistributorAccountId          sql.NullString  `bun:"distributor_account_id,type:varchar(255),nullzero"`
	CardlockCalcPriceFlag         sql.NullString  `bun:"cardlock_calc_price_flag,type:char,default:('N')"`
	CourtesyAddressId             sql.NullFloat64 `bun:"courtesy_address_id,type:decimal(19,0),nullzero"`
	DfoaShipToId                  sql.NullString  `bun:"dfoa_ship_to_id,type:varchar(255),nullzero"`
	ErouterTranType               sql.NullString  `bun:"erouter_tran_type,type:char,nullzero"`
	CardlockDiscount              sql.NullFloat64 `bun:"cardlock_discount,type:decimal(19,9),nullzero"`
	OrderPriorityUid              sql.NullInt32   `bun:"order_priority_uid,type:int,nullzero"`
	DeliveryZoneUid               sql.NullInt32   `bun:"delivery_zone_uid,type:int,nullzero"`
	PlantCode                     sql.NullString  `bun:"plant_code,type:varchar(255),nullzero"`
	FreightChargeUid              sql.NullInt32   `bun:"freight_charge_uid,type:int,nullzero"`
	UpsRoadnetAcctTypeId          sql.NullString  `bun:"ups_roadnet_acct_type_id,type:varchar(255),nullzero"`
	UpsRoadnetDeliveryDays        sql.NullString  `bun:"ups_roadnet_delivery_days,type:varchar(255),nullzero"`
	UpsRoadnetZoneId              sql.NullString  `bun:"ups_roadnet_zone_id,type:varchar(255),nullzero"`
	UseCustUpsHandlngChrgFlag     string          `bun:"use_cust_ups_handlng_chrg_flag,type:varchar,default:('Y')"`
	UpsHandlingCharge             float64         `bun:"ups_handling_charge,type:decimal(19,2),default:((0.00))"`
	SalesTaxPayableAccountNo      sql.NullString  `bun:"sales_tax_payable_account_no,type:varchar(255),nullzero"`
	VertexTaxableFlag             sql.NullString  `bun:"vertex_taxable_flag,type:char,nullzero"`
	CustomerTaxClass              sql.NullString  `bun:"customer_tax_class,type:varchar(255),nullzero"`
	TaxExemptReasonUid            sql.NullInt32   `bun:"tax_exempt_reason_uid,type:int,nullzero"`
	ServiceCenterUid              sql.NullInt32   `bun:"service_center_uid,type:int,nullzero"`
	DunsNumber                    sql.NullString  `bun:"duns_number,type:varchar(20),nullzero"`
	SmallTruckFlag                sql.NullString  `bun:"small_truck_flag,type:char,nullzero"`
	DeliveryTimeOffset            sql.NullInt32   `bun:"delivery_time_offset,type:int,default:((0))"`
	TermsOfDeliveryCd             sql.NullInt32   `bun:"terms_of_delivery_cd,type:int,nullzero"`
	ModeOfTransportCd             sql.NullInt32   `bun:"mode_of_transport_cd,type:int,nullzero"`
	UpsRoadnetExcludeFlag         string          `bun:"ups_roadnet_exclude_flag,type:char,default:('N')"`
	DfoaSoldToId                  sql.NullString  `bun:"dfoa_sold_to_id,type:varchar(255),nullzero"`
	InvoiceBatchUid               sql.NullInt32   `bun:"invoice_batch_uid,type:int,nullzero"`
	DateAcctOpened                sql.NullTime    `bun:"date_acct_opened,type:datetime,nullzero"`
	CourtesyContractShipToId      sql.NullFloat64 `bun:"courtesy_contract_ship_to_id,type:decimal(19,0),nullzero"`
	UpsThirdPartyBillingNo        sql.NullString  `bun:"ups_third_party_billing_no,type:varchar(255),nullzero"`
	ReqLotDocWithInvoiceFlag      sql.NullString  `bun:"req_lot_doc_with_invoice_flag,type:char,nullzero"`
	ReqLotDocWithPackinglistFlag  sql.NullString  `bun:"req_lot_doc_with_packinglist_flag,type:char,nullzero"`
	LimitOnlineWarrantiesFlag     sql.NullString  `bun:"limit_online_warranties_flag,type:char,default:('N')"`
	RemoteMargin                  sql.NullFloat64 `bun:"remote_margin,type:decimal(19,9),nullzero"`
	CfnResaleFlag                 string          `bun:"cfn_resale_flag,type:char,default:('N')"`
	SeparateInvoiceFlag           string          `bun:"separate_invoice_flag,type:char,default:('N')"`
	DoNotAutoInvoiceFlag          string          `bun:"do_not_auto_invoice_flag,type:char,default:('N')"`
	SicCode                       sql.NullFloat64 `bun:"sic_code,type:decimal(6,0),nullzero"`
	ValvolineNumber               sql.NullString  `bun:"valvoline_number,type:varchar(255),nullzero"`
	ConocoShipToId                sql.NullString  `bun:"conoco_ship_to_id,type:varchar(255),nullzero"`
	ConocoSoldToId                sql.NullString  `bun:"conoco_sold_to_id,type:varchar(255),nullzero"`
	SalesMarketGroupUid           sql.NullInt32   `bun:"sales_market_group_uid,type:int,nullzero"`
	OrderByBinPickTicketFlag      sql.NullString  `bun:"order_by_bin_pick_ticket_flag,type:char,nullzero"`
	AltTaxRateEligibleFlag        sql.NullString  `bun:"alt_tax_rate_eligible_flag,type:char,nullzero"`
	ResidentialAddressFlag        sql.NullString  `bun:"residential_address_flag,type:char,nullzero"`
	ExcludePrintPackinglistInWwms sql.NullString  `bun:"exclude_print_packinglist_in_wwms,type:char,nullzero"`
	RecordUsageActualLocFlag      sql.NullString  `bun:"record_usage_actual_loc_flag,type:char,nullzero"`
	ShipToLinkingId               sql.NullInt32   `bun:"ship_to_linking_id,type:int,nullzero"`
	ServiceLevelAgreementUid      sql.NullInt32   `bun:"service_level_agreement_uid,type:int,nullzero"`
	ScanAndPackFlag               string          `bun:"scan_and_pack_flag,type:char,default:('N')"`
	LotAgeRestrictionMonths       sql.NullInt32   `bun:"lot_age_restriction_months,type:int,nullzero"`
	FreightForwardAddressId       sql.NullFloat64 `bun:"freight_forward_address_id,type:decimal(19,0),nullzero"`
	AsbDeliveryMethodUid          sql.NullInt32   `bun:"asb_delivery_method_uid,type:int,nullzero"`
	StopNo                        sql.NullInt32   `bun:"stop_no,type:int,nullzero"`
	AdvancedBillingFlag           sql.NullString  `bun:"advanced_billing_flag,type:char,nullzero"`
	TrackaboutBillByShipTo        sql.NullString  `bun:"trackabout_bill_by_ship_to,type:char,default:('N')"`
	FreeFreightDays               sql.NullInt32   `bun:"free_freight_days,type:int,nullzero"`
	ProtectedFlag                 sql.NullString  `bun:"protected_flag,type:char,default:('N')"`
}

type ShipToModel struct {
	bun bun.IDB
}

func (m *ShipToModel) Get(ctx context.Context, companyId string, shipToId float64) (*ShipTo, error) {
	var shipTo ShipTo
	err := m.bun.NewSelect().Model(&shipTo).Where("company_id = ? and ship_to_id = ?", companyId, shipToId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shipTo, nil
}
