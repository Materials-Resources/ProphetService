package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvMast struct {
	bun.BaseModel `bun:"table:inv_mast"`

	InvMastUid                    int32           `bun:"inv_mast_uid,type:int,pk"`
	ItemId                        string          `bun:"item_id,type:varchar(40)"`
	ItemDesc                      string          `bun:"item_desc,type:varchar(40)"`
	DeleteFlag                    string          `bun:"delete_flag,type:char"`
	Weight                        sql.NullFloat64 `bun:"weight,type:decimal(19,12)"`
	NetWeight                     sql.NullFloat64 `bun:"net_weight,type:decimal(19,12)"`
	DateCreated                   time.Time       `bun:"date_created,type:datetime"`
	DateLastModified              time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy              string          `bun:"last_maintained_by,type:varchar(30)"`
	Inactive                      string          `bun:"inactive,type:char"`
	Cube                          sql.NullFloat64 `bun:"cube,type:decimal(19,12)"`
	DateInactive                  sql.NullTime    `bun:"date_inactive,type:datetime"`
	DateReactive                  sql.NullTime    `bun:"date_reactive,type:datetime"`
	CatchWeightIndicator          string          `bun:"catch_weight_indicator,type:char"`
	PurchasingWeight              sql.NullFloat64 `bun:"purchasing_weight,type:decimal(19,12)"`
	ClassId1                      sql.NullString  `bun:"class_id1,type:varchar(255)"`
	ClassId2                      sql.NullString  `bun:"class_id2,type:varchar(255)"`
	ClassId3                      sql.NullString  `bun:"class_id3,type:varchar(255)"`
	ClassId4                      sql.NullString  `bun:"class_id4,type:varchar(255)"`
	ClassId5                      sql.NullString  `bun:"class_id5,type:varchar(255)"`
	UpcOrEan                      sql.NullString  `bun:"upc_or_ean,type:varchar(3)"`
	UpcOrEanId                    sql.NullString  `bun:"upc_or_ean_id,type:varchar(15)"`
	Serialized                    string          `bun:"serialized,type:char"`
	ProductType                   string          `bun:"product_type,type:char"`
	Hose                          string          `bun:"hose,type:char"`
	Fitting                       string          `bun:"fitting,type:char"`
	DLength                       sql.NullFloat64 `bun:"d_length,type:decimal(6,4)"`
	CatalogItem                   string          `bun:"catalog_item,type:char"`
	ShortCode                     sql.NullString  `bun:"short_code,type:varchar(30)"`
	TrackLots                     string          `bun:"track_lots,type:char"`
	Requisition                   string          `bun:"requisition,type:char"`
	Price1                        sql.NullFloat64 `bun:"price1,type:decimal(19,9)"`
	Price2                        sql.NullFloat64 `bun:"price2,type:decimal(19,9)"`
	Price3                        sql.NullFloat64 `bun:"price3,type:decimal(19,9)"`
	Price4                        sql.NullFloat64 `bun:"price4,type:decimal(19,9)"`
	Price5                        sql.NullFloat64 `bun:"price5,type:decimal(19,9)"`
	Price6                        sql.NullFloat64 `bun:"price6,type:decimal(19,9)"`
	Price7                        sql.NullFloat64 `bun:"price7,type:decimal(19,9)"`
	Price8                        sql.NullFloat64 `bun:"price8,type:decimal(19,9)"`
	Price9                        sql.NullFloat64 `bun:"price9,type:decimal(19,9)"`
	Price10                       sql.NullFloat64 `bun:"price10,type:decimal(19,9)"`
	DefaultProductGroup           sql.NullString  `bun:"default_product_group,type:varchar(8)"`
	DefaultSalesDiscountGroup     sql.NullString  `bun:"default_sales_discount_group,type:varchar(8)"`
	DefaultPurchaseDiscGroup      sql.NullString  `bun:"default_purchase_disc_group,type:varchar(8)"`
	SalesPricingUnit              sql.NullString  `bun:"sales_pricing_unit,type:varchar(8)"`
	SalesPricingUnitSize          sql.NullFloat64 `bun:"sales_pricing_unit_size,type:decimal(19,4)"`
	PurchasePricingUnit           sql.NullString  `bun:"purchase_pricing_unit,type:varchar(8)"`
	PurchasePricingUnitSize       sql.NullFloat64 `bun:"purchase_pricing_unit_size,type:decimal(19,4)"`
	ExtendedDesc                  sql.NullString  `bun:"extended_desc,type:varchar(255)"`
	CommissionClassId             sql.NullString  `bun:"commission_class_id,type:varchar(8)"`
	OtherChargeItem               sql.NullString  `bun:"other_charge_item,type:char"`
	InvoiceType                   sql.NullString  `bun:"invoice_type,type:varchar(2)"`
	PickTicketType                sql.NullString  `bun:"pick_ticket_type,type:varchar(2)"`
	DefaultSellingUnit            sql.NullString  `bun:"default_selling_unit,type:varchar(8)"`
	UpdateViaPricingService       sql.NullString  `bun:"update_via_pricing_service,type:char"`
	DefaultPurchasingUnit         sql.NullString  `bun:"default_purchasing_unit,type:varchar(8)"`
	Hi                            sql.NullFloat64 `bun:"hi,type:decimal(6,0)"`
	Ti                            sql.NullFloat64 `bun:"ti,type:decimal(6,0)"`
	ConvertQuantities             sql.NullString  `bun:"convert_quantities,type:char"`
	HazMatFlag                    string          `bun:"haz_mat_flag,type:char"`
	ClassCode                     sql.NullFloat64 `bun:"class_code,type:decimal(6,2)"`
	ItemTermsDiscountPct          float64         `bun:"item_terms_discount_pct,type:decimal(19,2)"`
	TpcxStatus                    sql.NullInt32   `bun:"tpcx_status,type:int"`
	DefaultTransferUnit           sql.NullString  `bun:"default_transfer_unit,type:varchar(8)"`
	Keywords                      sql.NullString  `bun:"keywords,type:text(2147483647)"`
	FulltextTimestamp             *[]byte         `bun:"fulltext_timestamp,type:timestamp"`
	VendorConsigned               string          `bun:"vendor_consigned,type:char"`
	CreatedBy                     sql.NullString  `bun:"created_by,type:varchar(255)"`
	Disposition                   sql.NullString  `bun:"disposition,type:char"`
	SourceTypeCd                  sql.NullInt32   `bun:"source_type_cd,type:int"`
	UseTagsFlag                   string          `bun:"use_tags_flag,type:char"`
	BaseUnit                      string          `bun:"base_unit,type:varchar(8)"`
	TagHoldClassUid               sql.NullInt32   `bun:"tag_hold_class_uid,type:int"`
	ShippableUnitFlag             string          `bun:"shippable_unit_flag,type:char"`
	AutoAllocationFlag            sql.NullString  `bun:"auto_allocation_flag,type:char"`
	ParkerProductCd               sql.NullString  `bun:"parker_product_cd,type:varchar(255)"`
	ParkerDivisionCd              sql.NullString  `bun:"parker_division_cd,type:varchar(255)"`
	IvaTaxableFlag                sql.NullString  `bun:"iva_taxable_flag,type:char"`
	RestrictedFlag                sql.NullString  `bun:"restricted_flag,type:char"`
	SoldOutsideUsFlag             sql.NullString  `bun:"sold_outside_us_flag,type:char"`
	CurrencyId                    sql.NullFloat64 `bun:"currency_id,type:decimal(19,0)"`
	ServiceTermsDiscountPct       float64         `bun:"service_terms_discount_pct,type:decimal(19,2)"`
	ServiceCommissionClassId      sql.NullString  `bun:"service_commission_class_id,type:varchar(8)"`
	OverrideSpecificCosting       string          `bun:"override_specific_costing,type:char"`
	ConfigurableFlag              sql.NullString  `bun:"configurable_flag,type:char"`
	LastPricingServiceDate        sql.NullTime    `bun:"last_pricing_service_date,type:datetime"`
	CommodityCode                 sql.NullString  `bun:"commodity_code,type:varchar(255)"`
	DefaultPriceFamilyUid         sql.NullInt32   `bun:"default_price_family_uid,type:int"`
	SpaItemFlag                   sql.NullString  `bun:"spa_item_flag,type:char"`
	AvailForSchDeliveryFlag       sql.NullString  `bun:"avail_for_sch_delivery_flag,type:char"`
	ItemTypeCd                    sql.NullInt32   `bun:"item_type_cd,type:int"`
	CustParentInvMastUid          sql.NullInt32   `bun:"cust_parent_inv_mast_uid,type:int"`
	UnspscCode                    sql.NullString  `bun:"unspsc_code,type:varchar(255)"`
	DciCode                       sql.NullString  `bun:"dci_code,type:varchar(255)"`
	EpaCertReqFlag                sql.NullString  `bun:"epa_cert_req_flag,type:char"`
	AiaEnabledFlag                sql.NullString  `bun:"aia_enabled_flag,type:char"`
	AiaRemnantQty                 sql.NullFloat64 `bun:"aia_remnant_qty,type:decimal(19,9)"`
	ManufacturerProgramTypeUid    sql.NullInt32   `bun:"manufacturer_program_type_uid,type:int"`
	ManufacturerProgramTypePct    sql.NullFloat64 `bun:"manufacturer_program_type_pct,type:decimal(19,2)"`
	Length                        sql.NullFloat64 `bun:"length,type:decimal(19,9)"`
	Width                         sql.NullFloat64 `bun:"width,type:decimal(19,9)"`
	Height                        sql.NullFloat64 `bun:"height,type:decimal(19,9)"`
	UseOcTaxRulesFlag             sql.NullString  `bun:"use_oc_tax_rules_flag,type:char"`
	TallyFlag                     sql.NullString  `bun:"tally_flag,type:char"`
	OcPrintOnPickTicketFlag       sql.NullString  `bun:"oc_print_on_pick_ticket_flag,type:char"`
	OcPrintOnInvoiceFlag          sql.NullString  `bun:"oc_print_on_invoice_flag,type:char"`
	UseRevisionsFlag              sql.NullString  `bun:"use_revisions_flag,type:char"`
	SingleUseOrReusable           sql.NullString  `bun:"single_use_or_reusable,type:char"`
	UnitconvOverrideOeFlag        sql.NullString  `bun:"unitconv_override_oe_flag,type:char"`
	UnitconvOverridePurcFlag      sql.NullString  `bun:"unitconv_override_purc_flag,type:char"`
	ItemSalesTaxClass             sql.NullString  `bun:"item_sales_tax_class,type:varchar(255)"`
	LifoPoolItemClass             sql.NullString  `bun:"lifo_pool_item_class,type:varchar(8)"`
	CountryOfOriginReqFlag        sql.NullString  `bun:"country_of_origin_req_flag,type:char"`
	NmfcHdrUid                    sql.NullInt32   `bun:"nmfc_hdr_uid,type:int"`
	CatchLotWeightFlag            sql.NullString  `bun:"catch_lot_weight_flag,type:char"`
	Ucc128PackTypeCd              sql.NullInt32   `bun:"ucc128_pack_type_cd,type:int"`
	Ucc128StandardPackSize        sql.NullInt32   `bun:"ucc128_standard_pack_size,type:int"`
	RedemptionItemFlag            sql.NullString  `bun:"redemption_item_flag,type:char"`
	RedemptionValue               sql.NullFloat64 `bun:"redemption_value,type:decimal(19,9)"`
	BoFillCanadianPurchaseFlag    sql.NullString  `bun:"bo_fill_canadian_purchase_flag,type:char"`
	CountryOfOriginCode           sql.NullString  `bun:"country_of_origin_code,type:char(2)"`
	ApplyStateFuelSurchargeFlag   sql.NullString  `bun:"apply_state_fuel_surcharge_flag,type:char"`
	TypeOfSale                    sql.NullString  `bun:"type_of_sale,type:varchar(255)"`
	CommissionCostMultiplier      sql.NullFloat64 `bun:"commission_cost_multiplier,type:decimal(19,9)"`
	ItemLabelTypeCd               sql.NullInt32   `bun:"item_label_type_cd,type:int"`
	DrumPickupFlag                sql.NullString  `bun:"drum_pickup_flag,type:char"`
	ItemNotes                     sql.NullString  `bun:"item_notes,type:varchar(255)"`
	RetailAccessories             sql.NullString  `bun:"retail_accessories,type:varchar(255)"`
	HazmatOrmdFlag                string          `bun:"hazmat_ormd_flag,type:char"`
	HazmatOrmdQty                 sql.NullInt32   `bun:"hazmat_ormd_qty,type:int"`
	FunctionalClass               sql.NullString  `bun:"functional_class,type:varchar(10)"`
	SalesSetQuantity              sql.NullInt32   `bun:"sales_set_quantity,type:int"`
	SuppressFromWebFlag           string          `bun:"suppress_from_web_flag,type:char"`
	VndrStock                     sql.NullInt32   `bun:"vndr_stock,type:int"`
	PtlxStock                     sql.NullInt32   `bun:"ptlx_stock,type:int"`
	PfStock                       sql.NullInt32   `bun:"pf_stock,type:int"`
	RoundPrices                   sql.NullString  `bun:"round_prices,type:char"`
	PumpoffItem                   sql.NullString  `bun:"pumpoff_item,type:varchar(40)"`
	PumpoffUom                    sql.NullString  `bun:"pumpoff_uom,type:varchar(8)"`
	ConocoDivisionClassId         sql.NullString  `bun:"conoco_division_class_id,type:varchar(8)"`
	BuyGetClass                   sql.NullString  `bun:"buy_get_class,type:varchar(8)"`
	DiscontinuedDate              sql.NullTime    `bun:"discontinued_date,type:datetime"`
	LabelGroupItemFlag            sql.NullString  `bun:"label_group_item_flag,type:char"`
	CylinderItemFlag              sql.NullString  `bun:"cylinder_item_flag,type:char"`
	CoreItemFlag                  sql.NullString  `bun:"core_item_flag,type:char"`
	GenericItemDesc               sql.NullString  `bun:"generic_item_desc,type:varchar(40)"`
	AllowCustomDescFlag           sql.NullString  `bun:"allow_custom_desc_flag,type:char"`
	ManufacturerName              sql.NullString  `bun:"manufacturer_name,type:varchar(255)"`
	BrandName                     sql.NullString  `bun:"brand_name,type:varchar(255)"`
	PartNumber                    sql.NullString  `bun:"part_number,type:varchar(255)"`
	DoNotAutoAllocateDesc         sql.NullString  `bun:"do_not_auto_allocate_desc,type:varchar(255)"`
	ExcludeOrderLevelDiscountFlag sql.NullString  `bun:"exclude_order_level_discount_flag,type:char"`
	SampleItemFlag                sql.NullString  `bun:"sample_item_flag,type:char"`
	SampleInvMastUid              sql.NullInt32   `bun:"sample_inv_mast_uid,type:int"`
	CoreParentItemId              sql.NullString  `bun:"core_parent_item_id,type:varchar(40)"`
	AttributeGroupUid             sql.NullInt32   `bun:"attribute_group_uid,type:int"`
	PartnerProgramLaborItemFlag   sql.NullString  `bun:"partner_program_labor_item_flag,type:char"`
	BrandType                     sql.NullInt32   `bun:"brand_type,type:int"`
	IqsItemFlag                   sql.NullString  `bun:"iqs_item_flag,type:char"`
	ImportAsDirectShipFlag        sql.NullString  `bun:"import_as_direct_ship_flag,type:char"`
	DoNotFieldDestroyFlag         sql.NullString  `bun:"do_not_field_destroy_flag,type:char"`
	InvMastUidDupUsage            sql.NullInt32   `bun:"inv_mast_uid_dup_usage,type:int"`
	RolledItemFlag                sql.NullString  `bun:"rolled_item_flag,type:char"`
	ExcludeFromEdi832Flag         sql.NullString  `bun:"exclude_from_edi832_flag,type:char"`
	OrderHoldClassId              sql.NullString  `bun:"order_hold_class_id,type:varchar(255)"`
	DefaultDirectShipDispFlag     sql.NullString  `bun:"default_direct_ship_disp_flag,type:char"`
	BuyerId                       sql.NullString  `bun:"buyer_id,type:varchar(16)"`
	TariffSurchargePct            sql.NullFloat64 `bun:"tariff_surcharge_pct,type:decimal(19,9)"`
	TariffSurchargeInvMastUid     sql.NullInt32   `bun:"tariff_surcharge_inv_mast_uid,type:int"`
	PackingWeightTrackingOption   sql.NullString  `bun:"packing_weight_tracking_option,type:char"`
	ExcludeFromAutoShortBuyFlag   sql.NullString  `bun:"exclude_from_auto_short_buy_flag,type:char"`
	ItemClassificationId          sql.NullString  `bun:"item_classification_id,type:varchar(255)"`
	HazmatFeeInvMastUid           sql.NullInt32   `bun:"hazmat_fee_inv_mast_uid,type:int"`
	BolClass                      sql.NullString  `bun:"bol_class,type:varchar(8)"`
	RentalItemFlag                sql.NullString  `bun:"rental_item_flag,type:char"`
	RentalItemUid                 sql.NullInt32   `bun:"rental_item_uid,type:int"`
	DfltDimensionScale            sql.NullString  `bun:"dflt_dimension_scale,type:varchar(2)"`
	EccEnabledFlag                string          `bun:"ecc_enabled_flag,type:char"`
	LocalItemFlag                 sql.NullString  `bun:"local_item_flag,type:char"`
	ProductTier                   sql.NullString  `bun:"product_tier,type:varchar(12)"`
	CfcLicenseReqFlag             sql.NullString  `bun:"cfc_license_req_flag,type:char"`
	EquipmentClass                sql.NullString  `bun:"equipment_class,type:varchar(255)"`
	SerialTrackingOptionCd        sql.NullInt32   `bun:"serial_tracking_option_cd,type:int"`
	CpqItemFlag                   sql.NullString  `bun:"cpq_item_flag,type:char"`
	CpqConfiguratorId             sql.NullInt32   `bun:"cpq_configurator_id,type:int"`
	SerializationRequiredFlag     string          `bun:"serialization_required_flag,type:char"`
	DefaultTransferIncrementQty   sql.NullFloat64 `bun:"default_transfer_increment_qty,type:decimal(19,9)"`
	JmDtsFulfillmentFlag          sql.NullString  `bun:"jm_dts_fulfillment_flag,type:char"`
	DefectiveRecallFlag           string          `bun:"defective_recall_flag,type:char"`

	InvLoc InvLoc `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

var _ bun.BeforeInsertHook = (*InvMast)(nil)

func (m *InvMast) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	m.DateCreated = time.Now()
	m.DateLastModified = time.Now()
	return nil
}

var _ bun.BeforeUpdateHook = (*InvMast)(nil)

func (m *InvMast) BeforeUpdate(ctx context.Context, query *bun.UpdateQuery) error {
	m.DateLastModified = time.Now()
	return nil
}
