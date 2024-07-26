package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun/schema"
	"time"

	"github.com/uptrace/bun"
)

type OeLine struct {
	bun.BaseModel               `bun:"table:oe_line"`
	OrderNo                     string          `bun:"order_no,type:varchar(8),pk"`
	UnitPrice                   sql.NullFloat64 `bun:"unit_price,type:decimal(19,9),nullzero"`
	QtyOrdered                  sql.NullFloat64 `bun:"qty_ordered,type:decimal(19,9),nullzero"`
	QtyPerAssembly              float64         `bun:"qty_per_assembly,type:decimal(19,9)"`
	CompanyNo                   string          `bun:"company_no,type:varchar(8)"`
	DeleteFlag                  string          `bun:"delete_flag,type:char"`
	ManualPriceOveride          sql.NullString  `bun:"manual_price_overide,type:char,nullzero"`
	DateCreated                 time.Time       `bun:"date_created,type:datetime"`
	DateLastModified            time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string          `bun:"last_maintained_by,type:varchar(30)"`
	ExtendedPrice               sql.NullFloat64 `bun:"extended_price,type:decimal(19,4),nullzero"`
	SalesTax                    sql.NullFloat64 `bun:"sales_tax,type:decimal(19,4),nullzero"`
	LineNo                      float64         `bun:"line_no,type:decimal(19,0),pk"`
	Complete                    sql.NullString  `bun:"complete,type:char,nullzero"`
	UnitOfMeasure               sql.NullString  `bun:"unit_of_measure,type:varchar(8),nullzero"`
	BaseUtPrice                 sql.NullFloat64 `bun:"base_ut_price,type:decimal(19,9),nullzero"`
	CalcType                    sql.NullString  `bun:"calc_type,type:varchar(10),nullzero"`
	CalcValue                   sql.NullFloat64 `bun:"calc_value,type:decimal(19,4),nullzero"`
	Combinable                  sql.NullString  `bun:"combinable,type:char,nullzero"`
	Disposition                 sql.NullString  `bun:"disposition,type:char,nullzero"`
	QtyAllocated                sql.NullFloat64 `bun:"qty_allocated,type:decimal(19,9),nullzero"`
	QtyOnPickTickets            sql.NullFloat64 `bun:"qty_on_pick_tickets,type:decimal(19,9),nullzero"`
	QtyInvoiced                 sql.NullFloat64 `bun:"qty_invoiced,type:decimal(19,9),nullzero"`
	ExpediteDate                sql.NullTime    `bun:"expedite_date,type:datetime,nullzero"`
	RequiredDate                sql.NullTime    `bun:"required_date,type:datetime,nullzero"`
	SourceLocId                 sql.NullFloat64 `bun:"source_loc_id,type:decimal(19,0),nullzero"`
	ShipLocId                   sql.NullFloat64 `bun:"ship_loc_id,type:decimal(19,0),nullzero"`
	SupplierId                  sql.NullFloat64 `bun:"supplier_id,type:decimal(19,0),nullzero"`
	ProductGroupId              sql.NullString  `bun:"product_group_id,type:varchar(8),nullzero"`
	Assembly                    sql.NullString  `bun:"assembly,type:char,nullzero"`
	Scheduled                   sql.NullString  `bun:"scheduled,type:char,nullzero"`
	LotBill                     string          `bun:"lot_bill,type:char"`
	ExtendedDesc                sql.NullString  `bun:"extended_desc,type:varchar(255),nullzero"`
	UnitSize                    float64         `bun:"unit_size,type:decimal(19,4)"`
	UnitQuantity                float64         `bun:"unit_quantity,type:decimal(19,9)"`
	NextBreak                   sql.NullFloat64 `bun:"next_break,type:decimal(19,9),nullzero"`
	NextUtPrice                 sql.NullFloat64 `bun:"next_ut_price,type:decimal(19,9),nullzero"`
	CustomerPartNumber          string          `bun:"customer_part_number,type:varchar(40)"`
	TaxItem                     sql.NullString  `bun:"tax_item,type:char,nullzero"`
	OkToInterchange             sql.NullString  `bun:"ok_to_interchange,type:char,nullzero"`
	OtherCharge                 string          `bun:"other_charge,type:char"`
	SalesCost                   sql.NullFloat64 `bun:"sales_cost,type:decimal(19,9),nullzero"`
	CommissionCost              sql.NullFloat64 `bun:"commission_cost,type:decimal(19,9),nullzero"`
	RequestedDownpayment        sql.NullFloat64 `bun:"requested_downpayment,type:decimal(19,9),nullzero"`
	DownpaymentDate             sql.NullTime    `bun:"downpayment_date,type:datetime,nullzero"`
	WillCall                    sql.NullString  `bun:"will_call,type:char,nullzero"`
	OtherCost                   sql.NullFloat64 `bun:"other_cost,type:decimal(19,9),nullzero"`
	DownpaymentRemaining        sql.NullFloat64 `bun:"downpayment_remaining,type:decimal(19,9),nullzero"`
	CancelFlag                  sql.NullString  `bun:"cancel_flag,type:char,nullzero"`
	CommissionCostEdited        sql.NullString  `bun:"commission_cost_edited,type:char,nullzero"`
	OtherCostEdited             sql.NullString  `bun:"other_cost_edited,type:char,nullzero"`
	PoCost                      sql.NullFloat64 `bun:"po_cost,type:decimal(19,9),nullzero"`
	QtyCanceled                 sql.NullFloat64 `bun:"qty_canceled,type:decimal(19,9),nullzero"`
	PricingUnit                 sql.NullString  `bun:"pricing_unit,type:varchar(8),nullzero"`
	PricingUnitSize             sql.NullFloat64 `bun:"pricing_unit_size,type:decimal(19,4),nullzero"`
	QtyStaged                   sql.NullFloat64 `bun:"qty_staged,type:decimal(19,9),nullzero"`
	ItemSource                  int16           `bun:"item_source,type:smallint"`
	OeHdrUid                    int32           `bun:"oe_hdr_uid,type:int"`
	ParentOeLineUid             int32           `bun:"parent_oe_line_uid,type:int"`
	OeLineUid                   int32           `bun:"oe_line_uid,type:int"`
	PricePageUid                sql.NullInt32   `bun:"price_page_uid,type:int,nullzero"`
	PricingOption               int32           `bun:"pricing_option,type:int"`
	DetailType                  int32           `bun:"detail_type,type:int"`
	UserLineNo                  sql.NullString  `bun:"user_line_no,type:varchar(25),nullzero"`
	ItemTermsDiscountPct        float64         `bun:"item_terms_discount_pct,type:decimal(19,2),default:(0)"`
	DefaultInOe                 sql.NullString  `bun:"default_in_oe,type:char,default:('N')"`
	InvMastUid                  int32           `bun:"inv_mast_uid,type:int"`
	FreightIn                   float64         `bun:"freight_in,type:decimal(19,4),default:(0)"`
	PeerOeLineUid               int32           `bun:"peer_oe_line_uid,type:int,default:(0)"`
	PeerType                    int32           `bun:"peer_type,type:int,default:(0)"`
	SubstituteItem              sql.NullString  `bun:"substitute_item,type:char,default:('N')"`
	AllocateUsageToOriginalItem sql.NullString  `bun:"allocate_usage_to_original_item,type:char,default:('N')"`
	OrderCostEdited             string          `bun:"order_cost_edited,type:char,default:('N')"`
	CaptureUsage                string          `bun:"capture_usage,type:char"`
	PurchaseClassId             sql.NullString  `bun:"purchase_class_id,type:varchar(8),nullzero"`
	PickDate                    sql.NullTime    `bun:"pick_date,type:datetime,nullzero"`
	ShippingRouteUid            sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	InvXrefUid                  sql.NullInt32   `bun:"inv_xref_uid,type:int,nullzero"`
	HoseQtyNeeded               float64         `bun:"hose_qty_needed,type:decimal(19,9),default:(0)"`
	CreatedBy                   sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OriginalQtyOrdered          sql.NullFloat64 `bun:"original_qty_ordered,type:decimal(19,9),nullzero"`
	TagHoldClassUid             sql.NullInt32   `bun:"tag_hold_class_uid,type:int,nullzero"`
	CostCenter                  sql.NullString  `bun:"cost_center,type:varchar(255),nullzero"`
	DispositionEditedFlag       sql.NullString  `bun:"disposition_edited_flag,type:char,nullzero"`
	JobPriceHdrUid              sql.NullInt32   `bun:"job_price_hdr_uid,type:int,nullzero"`
	DeaRestrictionFailed        sql.NullString  `bun:"dea_restriction_failed,type:char,nullzero"`
	RestockFeePercentage        sql.NullFloat64 `bun:"restock_fee_percentage,type:decimal(19,9),nullzero"`
	CustomerConfiguredPrice     sql.NullFloat64 `bun:"customer_configured_price,type:decimal(19,9),nullzero"`
	CustomerPickedFlag          sql.NullString  `bun:"customer_picked_flag,type:char,nullzero"`
	JobPriceLineUid             sql.NullInt32   `bun:"job_price_line_uid,type:int,nullzero"`
	CustPoNo                    sql.NullString  `bun:"cust_po_no,type:varchar(255),nullzero"`
	JobPriceBinUid              sql.NullInt32   `bun:"job_price_bin_uid,type:int,nullzero"`
	UseContractCost             string          `bun:"use_contract_cost,type:char,default:('Y')"`
	CostPricePageUid            sql.NullInt32   `bun:"cost_price_page_uid,type:int,nullzero"`
	PackageTypeUid              sql.NullInt32   `bun:"package_type_uid,type:int,nullzero"`
	TagPackageQtyPer            sql.NullFloat64 `bun:"tag_package_qty_per,type:decimal(19,0),nullzero"`
	FreightCodeUid              sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	AcknowledgementDate         sql.NullTime    `bun:"acknowledgement_date,type:datetime,nullzero"`
	UsedSpecificCostFlag        string          `bun:"used_specific_cost_flag,type:char,default:('N')"`
	OperationUid                sql.NullInt32   `bun:"operation_uid,type:int,nullzero"`
	ServiceLaborUid             sql.NullInt32   `bun:"service_labor_uid,type:int,nullzero"`
	QuoteOutcomeCd              sql.NullInt32   `bun:"quote_outcome_cd,type:int,nullzero"`
	ItemCommissionClassId       sql.NullString  `bun:"item_commission_class_id,type:varchar(8),nullzero"`
	FromAltPrevRequestsFlag     sql.NullString  `bun:"from_alt_prev_requests_flag,type:char,nullzero"`
	UdCost                      sql.NullFloat64 `bun:"ud_cost,type:decimal(19,9),nullzero"`
	TankName                    sql.NullString  `bun:"tank_name,type:varchar(255),nullzero"`
	ItemBillToId                sql.NullFloat64 `bun:"item_bill_to_id,type:decimal(19,0),nullzero"`
	BuybackNo                   sql.NullString  `bun:"buyback_no,type:varchar(255),nullzero"`
	RoutingStatusFlag           sql.NullString  `bun:"routing_status_flag,type:char,nullzero"`
	SplitFlag                   sql.NullString  `bun:"split_flag,type:char,nullzero"`
	AdditionalDescription       sql.NullString  `bun:"additional_description,type:varchar(255),nullzero"`
	BillHoldFlag                sql.NullString  `bun:"bill_hold_flag,type:char,nullzero"`
	CorePrice                   sql.NullFloat64 `bun:"core_price,type:decimal(19,9),nullzero"`
	Clock                       sql.NullString  `bun:"clock,type:varchar(255),nullzero"`
	Cell                        sql.NullString  `bun:"cell,type:varchar(255),nullzero"`
	EnvironmentalFee            sql.NullFloat64 `bun:"environmental_fee,type:decimal(19,4),nullzero"`
	AdminFee                    sql.NullFloat64 `bun:"admin_fee,type:decimal(19,4),nullzero"`
	ExtDiscAmt                  sql.NullFloat64 `bun:"ext_disc_amt,type:decimal(19,9),nullzero"`
	SystemCalcUnitPrice         sql.NullFloat64 `bun:"system_calc_unit_price,type:decimal(19,9),nullzero"`
	Buyer                       sql.NullString  `bun:"buyer,type:varchar(255),nullzero"`
	Recipient                   sql.NullString  `bun:"recipient,type:varchar(255),nullzero"`
	StrategicUnitPrice          sql.NullFloat64 `bun:"strategic_unit_price,type:decimal(19,9),nullzero"`
	StrategicPricePageUid       sql.NullInt32   `bun:"strategic_price_page_uid,type:int,nullzero"`
	VerifiedFlag                sql.NullString  `bun:"verified_flag,type:char,nullzero"`
	VerifiedCode                sql.NullString  `bun:"verified_code,type:varchar(255),nullzero"`
	CoreStatusCd                sql.NullInt32   `bun:"core_status_cd,type:int,nullzero"`
	UsedStrategicPricingFlag    sql.NullString  `bun:"used_strategic_pricing_flag,type:char,nullzero"`
	LoanUid                     sql.NullInt32   `bun:"loan_uid,type:int,nullzero"`
	LoanItemUid                 sql.NullInt32   `bun:"loan_item_uid,type:int,nullzero"`
	OriginalUnitPrice           sql.NullFloat64 `bun:"original_unit_price,type:decimal(19,9),default:((0))"`
	OriginalFreightIn           sql.NullFloat64 `bun:"original_freight_in,type:decimal(19,9),nullzero"`
	StrategicListCostValue      sql.NullFloat64 `bun:"strategic_list_cost_value,type:decimal(19,9),nullzero"`
	StrategicListCostCd         sql.NullInt32   `bun:"strategic_list_cost_cd,type:int,nullzero"`
	SalesMarketGroupUid         sql.NullInt32   `bun:"sales_market_group_uid,type:int,nullzero"`
	SpecialItemIndicator        sql.NullInt32   `bun:"special_item_indicator,type:int,nullzero"`
	AwardedDate                 sql.NullTime    `bun:"awarded_date,type:datetime,nullzero"`
	SampleFlag                  sql.NullString  `bun:"sample_flag,type:char,nullzero"`
	RoutingAllocationChange     sql.NullFloat64 `bun:"routing_allocation_change,type:decimal(19,9),nullzero"`
	Price1                      sql.NullFloat64 `bun:"price1,type:decimal(19,9),nullzero"`
	UnitPickFee                 sql.NullFloat64 `bun:"unit_pick_fee,type:decimal(19,9),nullzero"`
	OeLineAltCode               sql.NullString  `bun:"oe_line_alt_code,type:varchar(40),nullzero"`
	RetailPrice                 sql.NullFloat64 `bun:"retail_price,type:decimal(19,9),nullzero"`
	QuotePriceOeLineUid         sql.NullInt32   `bun:"quote_price_oe_line_uid,type:int,nullzero"`
	CoreItemCost                sql.NullFloat64 `bun:"core_item_cost,type:decimal(19,9),nullzero"`
	ItemCommitmentDetailUid     sql.NullInt32   `bun:"item_commitment_detail_uid,type:int,nullzero"`
	SuppressCustomCost          sql.NullString  `bun:"suppress_custom_cost,type:char,nullzero"`
	BuyListApprovalInitial      sql.NullString  `bun:"buy_list_approval_initial,type:varchar(3),nullzero"`
	GeocomForcedSendFlag        sql.NullString  `bun:"geocom_forced_send_flag,type:char,default:('N')"`
	GeocomForcedQuantity        sql.NullFloat64 `bun:"geocom_forced_quantity,type:decimal(19,0),nullzero"`
	SuppressPickTicketIndicator sql.NullInt32   `bun:"suppress_pick_ticket_indicator,type:int,nullzero"`
	BeltingWidth                sql.NullFloat64 `bun:"belting_width,type:decimal(19,9),nullzero"`
	BeltingLength               sql.NullFloat64 `bun:"belting_length,type:decimal(19,9),nullzero"`
	BeltingNumCuts              sql.NullInt32   `bun:"belting_num_cuts,type:int,nullzero"`
	CarrierId                   sql.NullFloat64 `bun:"carrier_id,type:decimal(19,0),nullzero"`
	FreightInEditedFlag         sql.NullString  `bun:"freight_in_edited_flag,type:char,nullzero"`
	SalesDiscountGroupId        sql.NullString  `bun:"sales_discount_group_id,type:varchar(8),nullzero"`
	PriceFamilyUid              sql.NullInt32   `bun:"price_family_uid,type:int,nullzero"`
	CostCarrierContractLineUid  sql.NullInt32   `bun:"cost_carrier_contract_line_uid,type:int,nullzero"`
	PriceCarrierContractLineUid sql.NullInt32   `bun:"price_carrier_contract_line_uid,type:int,nullzero"`
	PumpOffFlag                 sql.NullString  `bun:"pump_off_flag,type:char,default:('N')"`
	CostCarrierContractZLineUid sql.NullInt32   `bun:"cost_carrier_contract_z_line_uid,type:int,nullzero"`
	TargetPrice                 sql.NullFloat64 `bun:"target_price,type:decimal(19,9),nullzero"`
	Edi830LastShipmentNo        sql.NullString  `bun:"edi830_last_shipment_no,type:varchar(255),nullzero"`
	Edi830LastReceiveQty        sql.NullFloat64 `bun:"edi830_last_receive_qty,type:decimal(19,9),nullzero"`
	Edi830LastReceiveDate       sql.NullTime    `bun:"edi830_last_receive_date,type:datetime,nullzero"`
	FaspacYtdQtyInvoiced        sql.NullFloat64 `bun:"faspac_ytd_qty_invoiced,type:decimal(19,9),nullzero"`
	CompetitorUid               sql.NullInt32   `bun:"competitor_uid,type:int,nullzero"`
	PriceAdjNote                sql.NullString  `bun:"price_adj_note,type:varchar(255),nullzero"`
	ImportedPrice               sql.NullFloat64 `bun:"imported_price,type:decimal(19,9),nullzero"`
	GlCode                      sql.NullString  `bun:"gl_code,type:varchar(255),nullzero"`
	RfqIndicatorFlag            sql.NullString  `bun:"rfq_indicator_flag,type:char,default:('Y')"`
	BuyGetRewardsFlag           sql.NullString  `bun:"buy_get_rewards_flag,type:char,default:('N')"`
	BuyGetXRewardsProgramUid    sql.NullInt32   `bun:"buy_get_x_rewards_program_uid,type:int,nullzero"`
	CarrierRebateCost           sql.NullFloat64 `bun:"carrier_rebate_cost,type:decimal(19,9),nullzero"`
	ExtendedDescLocation        sql.NullString  `bun:"extended_desc_location,type:varchar(255),nullzero"`
	ExtendedDescCustomer        sql.NullString  `bun:"extended_desc_customer,type:varchar(255),nullzero"`
	LineDiscount                sql.NullFloat64 `bun:"line_discount,type:decimal(19,9),nullzero"`
	LineDiscountDescription     sql.NullString  `bun:"line_discount_description,type:varchar(255),nullzero"`
	UnitDistributorNet          sql.NullFloat64 `bun:"unit_distributor_net,type:decimal(19,9),nullzero"`
	ExtendedDistributorNet      sql.NullFloat64 `bun:"extended_distributor_net,type:decimal(19,4),nullzero"`
	EcoFeeAmount                sql.NullFloat64 `bun:"eco_fee_amount,type:decimal(19,9),nullzero"`
	BypassProdOrderProcessing   sql.NullString  `bun:"bypass_prod_order_processing,type:char,nullzero"`
	GenericCustomDescription    sql.NullString  `bun:"generic_custom_description,type:varchar(255),nullzero"`
	LineSeqNo                   sql.NullInt32   `bun:"line_seq_no,type:int,nullzero"`
	OeBuyGetRewardsUid          sql.NullInt32   `bun:"oe_buy_get_rewards_uid,type:int,nullzero"`
	RestrictedByAddressFlag     sql.NullString  `bun:"restricted_by_address_flag,type:char,default:('N')"`
	SecondaryUnitPrice          sql.NullFloat64 `bun:"secondary_unit_price,type:decimal(19,9),nullzero"`
	SecondaryExtendedPrice      sql.NullFloat64 `bun:"secondary_extended_price,type:decimal(19,9),nullzero"`
	OriginalDisposition         sql.NullString  `bun:"original_disposition,type:char,nullzero"`
	SecondaryManualPriceOveride sql.NullString  `bun:"secondary_manual_price_overide,type:char,default:('N')"`
	ExtendedItemFlag            sql.NullString  `bun:"extended_item_flag,type:char,nullzero"`
	PrintDescOnFormsFlag        sql.NullString  `bun:"print_desc_on_forms_flag,type:char,nullzero"`
	PrintPartNo                 sql.NullString  `bun:"print_part_no,type:varchar(40),nullzero"`
	CourtesyCoreReturnFlag      sql.NullString  `bun:"courtesy_core_return_flag,type:char,nullzero"`
	CurrentbinUid               sql.NullInt32   `bun:"currentbin_uid,type:int,nullzero"`
	QaStatus                    sql.NullString  `bun:"qa_status,type:varchar(255),nullzero"`
	ServiceItemFlag             sql.NullString  `bun:"service_item_flag,type:char,nullzero"`
	PrediscountUnitPrice        sql.NullFloat64 `bun:"prediscount_unit_price,type:decimal(19,9),nullzero"`
	CustPercentageDisc          sql.NullFloat64 `bun:"cust_percentage_disc,type:decimal(19,9),nullzero"`
	OriginalQtyAllocated        sql.NullFloat64 `bun:"original_qty_allocated,type:decimal(19,9),default:((0))"`
	PriceLockFlag               sql.NullString  `bun:"price_lock_flag,type:char,nullzero"`
	ExtElectronicDiscAmt        sql.NullFloat64 `bun:"ext_electronic_disc_amt,type:decimal(19,9),nullzero"`
	ServiceItemId               sql.NullString  `bun:"service_item_id,type:varchar(255),nullzero"`
	PoContractNumber            sql.NullString  `bun:"po_contract_number,type:varchar(255),nullzero"`
	ExcludeFromEdi844867Flag    sql.NullString  `bun:"exclude_from_edi_844_867_flag,type:char,nullzero"`
	AddToOpenPtFlag             sql.NullString  `bun:"add_to_open_pt_flag,type:char,default:('N')"`
	SplitToOeLineUid            sql.NullInt32   `bun:"split_to_oe_line_uid,type:int,nullzero"`
	SplitFromOeLineUid          sql.NullInt32   `bun:"split_from_oe_line_uid,type:int,nullzero"`
	FullRolledItemCd            sql.NullInt32   `bun:"full_rolled_item_cd,type:int,nullzero"`
	RequiredTransferShipDate    sql.NullTime    `bun:"required_transfer_ship_date,type:datetime,nullzero"`
	ReasonId                    sql.NullString  `bun:"reason_id,type:varchar(5),nullzero"`
	DamagedItemSerialNumber     sql.NullString  `bun:"damaged_item_serial_number,type:varchar(40),nullzero"`
	ExportParkerPrinterName     sql.NullString  `bun:"export_parker_printer_name,type:varchar(255),nullzero"`
	ExportParkerLabelCopies     sql.NullInt32   `bun:"export_parker_label_copies,type:int,nullzero"`
	CreateTrackaboutLeaseFlag   sql.NullString  `bun:"create_trackabout_lease_flag,type:char,default:('N')"`
	JobPriceShipControlNoUid    sql.NullInt32   `bun:"job_price_ship_control_no_uid,type:int,nullzero"`
	PricingMultiplierFromLot    sql.NullFloat64 `bun:"pricing_multiplier_from_lot,type:decimal(19,9),nullzero"`
	RentalFlag                  sql.NullString  `bun:"rental_flag,type:char,default:('N')"`
	RentalReturnQty             sql.NullFloat64 `bun:"rental_return_qty,type:decimal(19,9),nullzero"`
	RentalPrice                 sql.NullFloat64 `bun:"rental_price,type:decimal(19,9),nullzero"`
	QuoteConversionListPrice    sql.NullFloat64 `bun:"quote_conversion_list_price,type:decimal(19,9),nullzero"`
	QuotedPrice                 sql.NullFloat64 `bun:"quoted_price,type:decimal(19,9),nullzero"`
	ConversionOrderCost         sql.NullFloat64 `bun:"conversion_order_cost,type:decimal(19,9),nullzero"`
	DirectShipFreightAmount     sql.NullFloat64 `bun:"direct_ship_freight_amount,type:decimal(19,4),nullzero"`
	ReturnCylinderQuantity      sql.NullInt32   `bun:"return_cylinder_quantity,type:int,nullzero"`
	LinkedChargeParentOeLineUid sql.NullInt32   `bun:"linked_charge_parent_oe_line_uid,type:int,nullzero"`
	LinkedChargeScaleQtyFlag    sql.NullString  `bun:"linked_charge_scale_qty_flag,type:char,nullzero"`
	SourceQuoteOeLineUid        sql.NullInt32   `bun:"source_quote_oe_line_uid,type:int,nullzero"`
	RestrictedClassUid          sql.NullInt32   `bun:"restricted_class_uid,type:int,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*OeLine)(nil)

func (m *OeLine) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OeLineModel struct {
	bun bun.IDB
}

type CreateOeLineParams struct {
	OrderNo              string
	UnitPrice            float64
	QtyOrdered           float64
	BaseUtPrice          float64
	CalcValue            float64
	SourceLocId          float64
	ShipLocId            float64
	SupplierId           float64
	ProductGroupId       string
	ExtendedDesc         string
	CustomerPartNumber   string
	CommissionCost       float64
	OeHdrUid             int32
	InvMastUid           int32
	SalesDiscountGroupId string
	UnitQuantity         float64
	UnitSize             float64
}

func (m *OeLineModel) Create(ctx context.Context, params CreateOeLineParams) (*OeLine, error) {
	lineSeqNo, err := m.GetLineSeqNo(ctx, params.OrderNo)
	if err != nil {
		switch {
		case errors.Is(err, ErrRecordNotFound):
			lineSeqNo = 1
		default:
			return nil, err
		}
	}

	lineSeqNo++

	oeLine := &OeLine{
		OrderNo:              params.OrderNo,
		UnitPrice:            sql.NullFloat64{Float64: params.UnitPrice, Valid: true},
		QtyPerAssembly:       0,
		LineNo:               float64(lineSeqNo),
		CompanyNo:            "MRS", // TODO set from config
		DeleteFlag:           "N",
		ManualPriceOveride:   sql.NullString{String: "N", Valid: true},
		ExtendedPrice:        sql.NullFloat64{Float64: 0, Valid: true},
		SalesTax:             sql.NullFloat64{Float64: 0, Valid: true},
		LineSeqNo:            sql.NullInt32{Int32: lineSeqNo, Valid: true},
		Complete:             sql.NullString{String: "N", Valid: true},
		UnitOfMeasure:        sql.NullString{String: "EA", Valid: true},
		BaseUtPrice:          sql.NullFloat64{Float64: params.BaseUtPrice, Valid: true},
		CalcValue:            sql.NullFloat64{Float64: params.CalcValue, Valid: true},
		Combinable:           sql.NullString{String: "N", Valid: true},
		Disposition:          sql.NullString{String: "N", Valid: true},
		ExpediteDate:         sql.NullTime{Time: time.Now(), Valid: true},
		RequiredDate:         sql.NullTime{Time: time.Now(), Valid: true},
		SourceLocId:          sql.NullFloat64{Float64: params.SourceLocId, Valid: true},
		ShipLocId:            sql.NullFloat64{Float64: params.ShipLocId, Valid: true},
		SupplierId:           sql.NullFloat64{Float64: params.SupplierId, Valid: true},
		ProductGroupId:       sql.NullString{String: params.ProductGroupId, Valid: true},
		Assembly:             sql.NullString{String: "N", Valid: true},
		Scheduled:            sql.NullString{String: "N", Valid: true},
		LotBill:              "N",
		ExtendedDesc:         sql.NullString{String: params.ExtendedDesc, Valid: true},
		UnitSize:             params.UnitSize,
		UnitQuantity:         params.UnitQuantity,
		CustomerPartNumber:   params.CustomerPartNumber,
		CommissionCost:       sql.NullFloat64{Float64: params.CommissionCost, Valid: true},
		OeHdrUid:             params.OeHdrUid,
		InvMastUid:           params.InvMastUid,
		SalesDiscountGroupId: sql.NullString{String: params.SalesDiscountGroupId, Valid: true},
	}

	err = m.generateOeLineUid(ctx, &oeLine.OeLineUid)
	if err != nil {
		return nil, err
	}

	err = m.Insert(ctx, oeLine)
	if err != nil {
		return nil, err
	}
	return oeLine, err
}
func (m *OeLineModel) Insert(ctx context.Context, oeLine *OeLine) error {
	_, err := m.bun.NewInsert().Model(oeLine).Exec(ctx)
	return err
}

func (m *OeLineModel) GetLineSeqNo(ctx context.Context, orderNo string) (int32, error) {
	query := `SELECT MAX(line_seq_no) FROM oe_line WHERE order_no = ?`
	var lineSeqNo sql.NullInt32
	err := m.bun.QueryRowContext(ctx, query, orderNo).Scan(&lineSeqNo)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ErrRecordNotFound
		}
		return 0, err
	}
	if lineSeqNo.Valid {
		return lineSeqNo.Int32, nil
	}
	return 0, nil
}

func (m *OeLineModel) CountByInvMastUid(ctx context.Context, invMastUid int32) (int32, error) {
	query := `SELECT COUNT(*) FROM oe_line WHERE inv_mast_uid = ?`
	var count int32
	err := m.bun.QueryRowContext(ctx, query, invMastUid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// generateOeLineUid returns a newly generated value for oe_line_uid
func (m *OeLineModel) generateOeLineUid(ctx context.Context, uid *int32) error {
	q := `
		DECLARE @oe_line_uid int
		EXEC @oe_line_uid = p21_get_counter 'oe_line', 1
		SELECT @oe_line_uid`
	err := m.bun.QueryRowContext(ctx, q).Scan(uid)
	if err != nil {
		return err
	}
	return nil
}
