package prophet_21_1_4559

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvMast struct {
	bun.BaseModel                 `bun:"table:inv_mast"`
	InvMastUid                    int32           `bun:"inv_mast_uid,type:int,pk"`
	ItemId                        string          `bun:"item_id,type:varchar(40)"`
	ItemDesc                      string          `bun:"item_desc,type:varchar(40)"`
	DeleteFlag                    string          `bun:"delete_flag,type:char"`
	Weight                        sql.NullFloat64 `bun:"weight,type:decimal(19,12),nullzero"`
	NetWeight                     sql.NullFloat64 `bun:"net_weight,type:decimal(19,12),nullzero"`
	DateCreated                   time.Time       `bun:"date_created,type:datetime"`
	DateLastModified              time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy              string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Inactive                      string          `bun:"inactive,type:char"`
	Cube                          sql.NullFloat64 `bun:"cube,type:decimal(19,12),nullzero"`
	DateInactive                  sql.NullTime    `bun:"date_inactive,type:datetime,nullzero"`
	DateReactive                  sql.NullTime    `bun:"date_reactive,type:datetime,nullzero"`
	CatchWeightIndicator          string          `bun:"catch_weight_indicator,type:char"`
	PurchasingWeight              sql.NullFloat64 `bun:"purchasing_weight,type:decimal(19,12),nullzero"`
	ClassId1                      sql.NullString  `bun:"class_id1,type:varchar(255),nullzero"`
	ClassId2                      sql.NullString  `bun:"class_id2,type:varchar(255),nullzero"`
	ClassId3                      sql.NullString  `bun:"class_id3,type:varchar(255),nullzero"`
	ClassId4                      sql.NullString  `bun:"class_id4,type:varchar(255),nullzero"`
	ClassId5                      sql.NullString  `bun:"class_id5,type:varchar(255),nullzero"`
	UpcOrEan                      sql.NullString  `bun:"upc_or_ean,type:varchar(3),nullzero"`
	UpcOrEanId                    sql.NullString  `bun:"upc_or_ean_id,type:varchar(15),nullzero"`
	Serialized                    string          `bun:"serialized,type:char"`
	ProductType                   string          `bun:"product_type,type:char"`
	Hose                          string          `bun:"hose,type:char"`
	Fitting                       string          `bun:"fitting,type:char"`
	DLength                       sql.NullFloat64 `bun:"d_length,type:decimal(6,4),nullzero"`
	CatalogItem                   string          `bun:"catalog_item,type:char"`
	ShortCode                     sql.NullString  `bun:"short_code,type:varchar(30),nullzero"`
	TrackLots                     string          `bun:"track_lots,type:char"`
	Requisition                   string          `bun:"requisition,type:char,default:('N')"`
	Price1                        sql.NullFloat64 `bun:"price1,type:decimal(19,9),nullzero"`
	Price2                        sql.NullFloat64 `bun:"price2,type:decimal(19,9),nullzero"`
	Price3                        sql.NullFloat64 `bun:"price3,type:decimal(19,9),nullzero"`
	Price4                        sql.NullFloat64 `bun:"price4,type:decimal(19,9),nullzero"`
	Price5                        sql.NullFloat64 `bun:"price5,type:decimal(19,9),nullzero"`
	Price6                        sql.NullFloat64 `bun:"price6,type:decimal(19,9),nullzero"`
	Price7                        sql.NullFloat64 `bun:"price7,type:decimal(19,9),nullzero"`
	Price8                        sql.NullFloat64 `bun:"price8,type:decimal(19,9),nullzero"`
	Price9                        sql.NullFloat64 `bun:"price9,type:decimal(19,9),nullzero"`
	Price10                       sql.NullFloat64 `bun:"price10,type:decimal(19,9),nullzero"`
	DefaultProductGroup           sql.NullString  `bun:"default_product_group,type:varchar(8),nullzero"`
	DefaultSalesDiscountGroup     sql.NullString  `bun:"default_sales_discount_group,type:varchar(8),nullzero"`
	DefaultPurchaseDiscGroup      sql.NullString  `bun:"default_purchase_disc_group,type:varchar(8),nullzero"`
	SalesPricingUnit              sql.NullString  `bun:"sales_pricing_unit,type:varchar(8),nullzero"`
	SalesPricingUnitSize          sql.NullFloat64 `bun:"sales_pricing_unit_size,type:decimal(19,4),nullzero"`
	PurchasePricingUnit           sql.NullString  `bun:"purchase_pricing_unit,type:varchar(8),nullzero"`
	PurchasePricingUnitSize       sql.NullFloat64 `bun:"purchase_pricing_unit_size,type:decimal(19,4),nullzero"`
	ExtendedDesc                  sql.NullString  `bun:"extended_desc,type:varchar(255),nullzero"`
	CommissionClassId             sql.NullString  `bun:"commission_class_id,type:varchar(8),nullzero"`
	OtherChargeItem               sql.NullString  `bun:"other_charge_item,type:char,nullzero"`
	InvoiceType                   sql.NullString  `bun:"invoice_type,type:varchar(2),nullzero"`
	PickTicketType                sql.NullString  `bun:"pick_ticket_type,type:varchar(2),nullzero"`
	DefaultSellingUnit            sql.NullString  `bun:"default_selling_unit,type:varchar(8),nullzero"`
	UpdateViaPricingService       sql.NullString  `bun:"update_via_pricing_service,type:char,nullzero"`
	DefaultPurchasingUnit         sql.NullString  `bun:"default_purchasing_unit,type:varchar(8),nullzero"`
	Hi                            sql.NullFloat64 `bun:"hi,type:decimal(6,0),nullzero"`
	Ti                            sql.NullFloat64 `bun:"ti,type:decimal(6,0),nullzero"`
	ConvertQuantities             sql.NullString  `bun:"convert_quantities,type:char,nullzero"`
	HazMatFlag                    string          `bun:"haz_mat_flag,type:char,default:('N')"`
	ClassCode                     sql.NullFloat64 `bun:"class_code,type:decimal(6,2),nullzero"`
	ItemTermsDiscountPct          float64         `bun:"item_terms_discount_pct,type:decimal(19,2),default:(0)"`
	TpcxStatus                    sql.NullInt32   `bun:"tpcx_status,type:int,nullzero"`
	DefaultTransferUnit           sql.NullString  `bun:"default_transfer_unit,type:varchar(8),nullzero"`
	Keywords                      sql.NullString  `bun:"keywords,type:text(2147483647),nullzero"`
	FulltextTimestamp             *[]byte         `bun:"fulltext_timestamp,type:timestamp,nullzero"`
	VendorConsigned               string          `bun:"vendor_consigned,type:char,default:('N')"`
	CreatedBy                     sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	Disposition                   sql.NullString  `bun:"disposition,type:char,default:('N')"`
	SourceTypeCd                  sql.NullInt32   `bun:"source_type_cd,type:int,nullzero"`
	UseTagsFlag                   string          `bun:"use_tags_flag,type:char,default:('N')"`
	BaseUnit                      string          `bun:"base_unit,type:varchar(8)"`
	TagHoldClassUid               sql.NullInt32   `bun:"tag_hold_class_uid,type:int,nullzero"`
	ShippableUnitFlag             string          `bun:"shippable_unit_flag,type:char,default:('N')"`
	AutoAllocationFlag            sql.NullString  `bun:"auto_allocation_flag,type:char,nullzero"`
	ParkerProductCd               sql.NullString  `bun:"parker_product_cd,type:varchar(255),nullzero"`
	ParkerDivisionCd              sql.NullString  `bun:"parker_division_cd,type:varchar(255),nullzero"`
	IvaTaxableFlag                sql.NullString  `bun:"iva_taxable_flag,type:char,nullzero"`
	RestrictedFlag                sql.NullString  `bun:"restricted_flag,type:char,nullzero"`
	SoldOutsideUsFlag             sql.NullString  `bun:"sold_outside_us_flag,type:char,default:('N')"`
	CurrencyId                    sql.NullFloat64 `bun:"currency_id,type:decimal(19,0),nullzero"`
	ServiceTermsDiscountPct       float64         `bun:"service_terms_discount_pct,type:decimal(19,2),default:((0))"`
	ServiceCommissionClassId      sql.NullString  `bun:"service_commission_class_id,type:varchar(8),nullzero"`
	OverrideSpecificCosting       string          `bun:"override_specific_costing,type:char,default:('Y')"`
	ConfigurableFlag              sql.NullString  `bun:"configurable_flag,type:char,default:('N')"`
	LastPricingServiceDate        sql.NullTime    `bun:"last_pricing_service_date,type:datetime,nullzero"`
	CommodityCode                 sql.NullString  `bun:"commodity_code,type:varchar(255),nullzero"`
	DefaultPriceFamilyUid         sql.NullInt32   `bun:"default_price_family_uid,type:int,nullzero"`
	SpaItemFlag                   sql.NullString  `bun:"spa_item_flag,type:char,default:('N')"`
	AvailForSchDeliveryFlag       sql.NullString  `bun:"avail_for_sch_delivery_flag,type:char,nullzero"`
	ItemTypeCd                    sql.NullInt32   `bun:"item_type_cd,type:int,default:((300))"`
	CustParentInvMastUid          sql.NullInt32   `bun:"cust_parent_inv_mast_uid,type:int,nullzero"`
	UnspscCode                    sql.NullString  `bun:"unspsc_code,type:varchar(255),nullzero"`
	DciCode                       sql.NullString  `bun:"dci_code,type:varchar(255),nullzero"`
	EpaCertReqFlag                sql.NullString  `bun:"epa_cert_req_flag,type:char,default:('N')"`
	AiaEnabledFlag                sql.NullString  `bun:"aia_enabled_flag,type:char,nullzero"`
	AiaRemnantQty                 sql.NullFloat64 `bun:"aia_remnant_qty,type:decimal(19,9),nullzero"`
	ManufacturerProgramTypeUid    sql.NullInt32   `bun:"manufacturer_program_type_uid,type:int,nullzero"`
	ManufacturerProgramTypePct    sql.NullFloat64 `bun:"manufacturer_program_type_pct,type:decimal(19,2),nullzero"`
	Length                        sql.NullFloat64 `bun:"length,type:decimal(19,9),nullzero"`
	Width                         sql.NullFloat64 `bun:"width,type:decimal(19,9),nullzero"`
	Height                        sql.NullFloat64 `bun:"height,type:decimal(19,9),nullzero"`
	UseOcTaxRulesFlag             sql.NullString  `bun:"use_oc_tax_rules_flag,type:char,default:('N')"`
	TallyFlag                     sql.NullString  `bun:"tally_flag,type:char,nullzero"`
	OcPrintOnPickTicketFlag       sql.NullString  `bun:"oc_print_on_pick_ticket_flag,type:char,nullzero"`
	OcPrintOnInvoiceFlag          sql.NullString  `bun:"oc_print_on_invoice_flag,type:char,nullzero"`
	UseRevisionsFlag              sql.NullString  `bun:"use_revisions_flag,type:char,default:('N')"`
	SingleUseOrReusable           sql.NullString  `bun:"single_use_or_reusable,type:char,nullzero"`
	UnitconvOverrideOeFlag        sql.NullString  `bun:"unitconv_override_oe_flag,type:char,nullzero"`
	UnitconvOverridePurcFlag      sql.NullString  `bun:"unitconv_override_purc_flag,type:char,nullzero"`
	ItemSalesTaxClass             sql.NullString  `bun:"item_sales_tax_class,type:varchar(255),nullzero"`
	LifoPoolItemClass             sql.NullString  `bun:"lifo_pool_item_class,type:varchar(8),nullzero"`
	CountryOfOriginReqFlag        sql.NullString  `bun:"country_of_origin_req_flag,type:char,nullzero"`
	NmfcHdrUid                    sql.NullInt32   `bun:"nmfc_hdr_uid,type:int,nullzero"`
	CatchLotWeightFlag            sql.NullString  `bun:"catch_lot_weight_flag,type:char,default:('N')"`
	Ucc128PackTypeCd              sql.NullInt32   `bun:"ucc128_pack_type_cd,type:int,nullzero"`
	Ucc128StandardPackSize        sql.NullInt32   `bun:"ucc128_standard_pack_size,type:int,nullzero"`
	RedemptionItemFlag            sql.NullString  `bun:"redemption_item_flag,type:char,nullzero"`
	RedemptionValue               sql.NullFloat64 `bun:"redemption_value,type:decimal(19,9),nullzero"`
	BoFillCanadianPurchaseFlag    sql.NullString  `bun:"bo_fill_canadian_purchase_flag,type:char,default:('N')"`
	CountryOfOriginCode           sql.NullString  `bun:"country_of_origin_code,type:char(2),nullzero"`
	ApplyStateFuelSurchargeFlag   sql.NullString  `bun:"apply_state_fuel_surcharge_flag,type:char,default:('N')"`
	TypeOfSale                    sql.NullString  `bun:"type_of_sale,type:varchar(255),nullzero"`
	CommissionCostMultiplier      sql.NullFloat64 `bun:"commission_cost_multiplier,type:decimal(19,9),nullzero"`
	ItemLabelTypeCd               sql.NullInt32   `bun:"item_label_type_cd,type:int,nullzero"`
	DrumPickupFlag                sql.NullString  `bun:"drum_pickup_flag,type:char,default:('N')"`
	ItemNotes                     sql.NullString  `bun:"item_notes,type:varchar(255),nullzero"`
	RetailAccessories             sql.NullString  `bun:"retail_accessories,type:varchar(255),nullzero"`
	HazmatOrmdFlag                string          `bun:"hazmat_ormd_flag,type:char,default:('N')"`
	HazmatOrmdQty                 sql.NullInt32   `bun:"hazmat_ormd_qty,type:int,nullzero"`
	FunctionalClass               sql.NullString  `bun:"functional_class,type:varchar(10),nullzero"`
	SalesSetQuantity              sql.NullInt32   `bun:"sales_set_quantity,type:int,nullzero"`
	SuppressFromWebFlag           string          `bun:"suppress_from_web_flag,type:char,default:('N')"`
	VndrStock                     sql.NullInt32   `bun:"vndr_stock,type:int,nullzero"`
	PtlxStock                     sql.NullInt32   `bun:"ptlx_stock,type:int,nullzero"`
	PfStock                       sql.NullInt32   `bun:"pf_stock,type:int,nullzero"`
	RoundPrices                   sql.NullString  `bun:"round_prices,type:char,default:('N')"`
	PumpoffItem                   sql.NullString  `bun:"pumpoff_item,type:varchar(40),nullzero"`
	PumpoffUom                    sql.NullString  `bun:"pumpoff_uom,type:varchar(8),nullzero"`
	ConocoDivisionClassId         sql.NullString  `bun:"conoco_division_class_id,type:varchar(8),nullzero"`
	BuyGetClass                   sql.NullString  `bun:"buy_get_class,type:varchar(8),nullzero"`
	DiscontinuedDate              sql.NullTime    `bun:"discontinued_date,type:datetime,nullzero"`
	LabelGroupItemFlag            sql.NullString  `bun:"label_group_item_flag,type:char,nullzero"`
	CylinderItemFlag              sql.NullString  `bun:"cylinder_item_flag,type:char,nullzero"`
	CoreItemFlag                  sql.NullString  `bun:"core_item_flag,type:char,nullzero"`
	GenericItemDesc               sql.NullString  `bun:"generic_item_desc,type:varchar(40),nullzero"`
	AllowCustomDescFlag           sql.NullString  `bun:"allow_custom_desc_flag,type:char,default:('N')"`
	ManufacturerName              sql.NullString  `bun:"manufacturer_name,type:varchar(255),nullzero"`
	BrandName                     sql.NullString  `bun:"brand_name,type:varchar(255),nullzero"`
	PartNumber                    sql.NullString  `bun:"part_number,type:varchar(255),nullzero"`
	DoNotAutoAllocateDesc         sql.NullString  `bun:"do_not_auto_allocate_desc,type:varchar(255),nullzero"`
	ExcludeOrderLevelDiscountFlag sql.NullString  `bun:"exclude_order_level_discount_flag,type:char,nullzero"`
	SampleItemFlag                sql.NullString  `bun:"sample_item_flag,type:char,default:('N')"`
	SampleInvMastUid              sql.NullInt32   `bun:"sample_inv_mast_uid,type:int,nullzero"`
	CoreParentItemId              sql.NullString  `bun:"core_parent_item_id,type:varchar(40),nullzero"`
	AttributeGroupUid             sql.NullInt32   `bun:"attribute_group_uid,type:int,nullzero"`
	PartnerProgramLaborItemFlag   sql.NullString  `bun:"partner_program_labor_item_flag,type:char,nullzero"`
	BrandType                     sql.NullInt32   `bun:"brand_type,type:int,nullzero"`
	IqsItemFlag                   sql.NullString  `bun:"iqs_item_flag,type:char,nullzero"`
	ImportAsDirectShipFlag        sql.NullString  `bun:"import_as_direct_ship_flag,type:char,default:('N')"`
	DoNotFieldDestroyFlag         sql.NullString  `bun:"do_not_field_destroy_flag,type:char,nullzero"`
	InvMastUidDupUsage            sql.NullInt32   `bun:"inv_mast_uid_dup_usage,type:int,nullzero"`
	RolledItemFlag                sql.NullString  `bun:"rolled_item_flag,type:char,nullzero"`
	ExcludeFromEdi832Flag         sql.NullString  `bun:"exclude_from_edi832_flag,type:char,nullzero"`
	OrderHoldClassId              sql.NullString  `bun:"order_hold_class_id,type:varchar(255),nullzero"`
	DefaultDirectShipDispFlag     sql.NullString  `bun:"default_direct_ship_disp_flag,type:char,default:('N')"`
	BuyerId                       sql.NullString  `bun:"buyer_id,type:varchar(16),nullzero"`
	TariffSurchargePct            sql.NullFloat64 `bun:"tariff_surcharge_pct,type:decimal(19,9),nullzero"`
	TariffSurchargeInvMastUid     sql.NullInt32   `bun:"tariff_surcharge_inv_mast_uid,type:int,nullzero"`
	PackingWeightTrackingOption   sql.NullString  `bun:"packing_weight_tracking_option,type:char,nullzero"`
	ExcludeFromAutoShortBuyFlag   sql.NullString  `bun:"exclude_from_auto_short_buy_flag,type:char,nullzero"`
	ItemClassificationId          sql.NullString  `bun:"item_classification_id,type:varchar(255),nullzero"`
	HazmatFeeInvMastUid           sql.NullInt32   `bun:"hazmat_fee_inv_mast_uid,type:int,nullzero"`
	BolClass                      sql.NullString  `bun:"bol_class,type:varchar(8),nullzero"`
	RentalItemFlag                sql.NullString  `bun:"rental_item_flag,type:char,default:('N')"`
	RentalItemUid                 sql.NullInt32   `bun:"rental_item_uid,type:int,nullzero"`
	DfltDimensionScale            sql.NullString  `bun:"dflt_dimension_scale,type:varchar(2),nullzero"`
	EccEnabledFlag                string          `bun:"ecc_enabled_flag,type:char,default:('N')"`
	LocalItemFlag                 sql.NullString  `bun:"local_item_flag,type:char,default:('N')"`

	AlternateCodes []*AlternateCode `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvLocItems    []*InvLoc        `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
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
