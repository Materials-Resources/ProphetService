package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"time"
)

type Address struct {
	bun.BaseModel               `bun:"table:address"`
	Id                          float64         `bun:"id,type:decimal(19,0),pk"`
	Name                        string          `bun:"name,type:varchar(50)"`
	MailAddress1                sql.NullString  `bun:"mail_address1,type:varchar(50),nullzero"`
	MailAddress2                sql.NullString  `bun:"mail_address2,type:varchar(50),nullzero"`
	MailCity                    sql.NullString  `bun:"mail_city,type:varchar(50),nullzero"`
	MailState                   sql.NullString  `bun:"mail_state,type:varchar(50),nullzero"`
	MailPostalCode              sql.NullString  `bun:"mail_postal_code,type:varchar(10),nullzero"`
	MailCountry                 sql.NullString  `bun:"mail_country,type:varchar(50),nullzero"`
	PhysAddress1                sql.NullString  `bun:"phys_address1,type:varchar(50),nullzero"`
	PhysAddress2                sql.NullString  `bun:"phys_address2,type:varchar(50),nullzero"`
	PhysCity                    sql.NullString  `bun:"phys_city,type:varchar(50),nullzero"`
	PhysState                   sql.NullString  `bun:"phys_state,type:varchar(50),nullzero"`
	PhysPostalCode              sql.NullString  `bun:"phys_postal_code,type:varchar(10),nullzero"`
	PhysCountry                 sql.NullString  `bun:"phys_country,type:varchar(50),nullzero"`
	CentralPhoneNumber          sql.NullString  `bun:"central_phone_number,type:varchar(20),nullzero"`
	CentralFaxNumber            sql.NullString  `bun:"central_fax_number,type:varchar(20),nullzero"`
	FederalIdNumber             sql.NullString  `bun:"federal_id_number,type:varchar(255),nullzero"`
	ResaleCertificate           sql.NullString  `bun:"resale_certificate,type:varchar(40),nullzero"`
	UpsCode                     sql.NullString  `bun:"ups_code,type:varchar(40),nullzero"`
	CorpAddressId               sql.NullFloat64 `bun:"corp_address_id,type:decimal(19,0),nullzero"`
	BillingAddressId            sql.NullFloat64 `bun:"billing_address_id,type:decimal(19,0),nullzero"`
	CreditAddressId             sql.NullFloat64 `bun:"credit_address_id,type:decimal(19,0),nullzero"`
	Customer                    sql.NullString  `bun:"customer,type:char,nullzero"`
	Vendor                      sql.NullString  `bun:"vendor,type:char,nullzero"`
	Employee                    sql.NullString  `bun:"employee,type:char,nullzero"`
	Prospect                    sql.NullString  `bun:"prospect,type:char,nullzero"`
	BillingAddress              sql.NullString  `bun:"billing_address,type:char,nullzero"`
	ShippingAddress             sql.NullString  `bun:"shipping_address,type:char,nullzero"`
	PaymentAddress              sql.NullString  `bun:"payment_address,type:char,nullzero"`
	Incorporated                sql.NullString  `bun:"incorporated,type:char,nullzero"`
	DateCreated                 time.Time       `bun:"date_created,type:datetime"`
	DateLastModified            time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DeleteFlag                  string          `bun:"delete_flag,type:char"`
	DefaultShipTime             sql.NullTime    `bun:"default_ship_time,type:datetime,nullzero"`
	PreferredLocationId         sql.NullFloat64 `bun:"preferred_location_id,type:decimal(19,0),nullzero"`
	CarrierId                   sql.NullString  `bun:"carrier_id,type:varchar(10),nullzero"`
	DeliveryInstructions1       sql.NullString  `bun:"delivery_instructions1,type:varchar(20),nullzero"`
	DeliveryInstructions2       sql.NullString  `bun:"delivery_instructions2,type:varchar(20),nullzero"`
	DeliveryInstructions3       sql.NullString  `bun:"delivery_instructions3,type:varchar(20),nullzero"`
	MorBegDelivery              sql.NullTime    `bun:"mor_beg_delivery,type:datetime,nullzero"`
	MorEndDelivery              sql.NullTime    `bun:"mor_end_delivery,type:datetime,nullzero"`
	EveBegDelivery              sql.NullTime    `bun:"eve_beg_delivery,type:datetime,nullzero"`
	EveEndDelivery              sql.NullTime    `bun:"eve_end_delivery,type:datetime,nullzero"`
	InventoryLocationFlag       sql.NullString  `bun:"inventory_location_flag,type:char,nullzero"`
	ShowOutItems                sql.NullString  `bun:"show_out_items,type:char,nullzero"`
	StoreNo                     sql.NullString  `bun:"store_no,type:varchar(5),nullzero"`
	InvoiceType                 sql.NullString  `bun:"invoice_type,type:varchar(2),nullzero"`
	AddressIdString             sql.NullString  `bun:"address_id_string,type:varchar(22),nullzero"`
	CarrierFlag                 sql.NullString  `bun:"carrier_flag,type:char,nullzero"`
	DefaultCarrierId            sql.NullFloat64 `bun:"default_carrier_id,type:decimal(19,0),nullzero"`
	TradePercentDisc            sql.NullFloat64 `bun:"trade_percent_disc,type:decimal(19,0),nullzero"`
	Class1Id                    sql.NullString  `bun:"class1_id,type:varchar(255),nullzero"`
	Class2Id                    sql.NullString  `bun:"class2_id,type:varchar(255),nullzero"`
	Class3Id                    sql.NullString  `bun:"class3_id,type:varchar(255),nullzero"`
	Class4Id                    sql.NullString  `bun:"class4_id,type:varchar(255),nullzero"`
	Class5Id                    sql.NullString  `bun:"class5_id,type:varchar(255),nullzero"`
	CentralWattsNumber          sql.NullString  `bun:"central_watts_number,type:varchar(20),nullzero"`
	Alternative1099Name         sql.NullString  `bun:"alternative_1099_name,type:varchar(40),nullzero"`
	NameControl                 sql.NullString  `bun:"name_control,type:varchar(4),nullzero"`
	DefaultShipToCompany        sql.NullString  `bun:"default_ship_to_company,type:varchar(8),nullzero"`
	DefaultShipToBranch         sql.NullString  `bun:"default_ship_to_branch,type:varchar(8),nullzero"`
	ShipToPackingBasis          sql.NullString  `bun:"ship_to_packing_basis,type:varchar(16),nullzero"`
	BillOfLadingType            sql.NullString  `bun:"bill_of_lading_type,type:varchar(2),nullzero"`
	EmailAddress                sql.NullString  `bun:"email_address,type:varchar(255),nullzero"`
	Url                         sql.NullString  `bun:"url,type:varchar(255),nullzero"`
	CarrierFedexFlag            string          `bun:"carrier_fedex_flag,type:char,default:('N')"`
	RouteviewRequireRoutingFlag string          `bun:"routeview_require_routing_flag,type:char,default:('N')"`
	CarrierTypeCd               int32           `bun:"carrier_type_cd,type:int,default:((300))"`
	CarrierDoNotRouteFlag       sql.NullString  `bun:"carrier_do_not_route_flag,type:char,default:('N')"`
	RoadnetDoNotRouteFlag       sql.NullString  `bun:"roadnet_do_not_route_flag,type:char,nullzero"`
	CarrierStrategicFreightFlag sql.NullString  `bun:"carrier_strategic_freight_flag,type:char,nullzero"`
	CarrierFixedFreightCharge   sql.NullFloat64 `bun:"carrier_fixed_freight_charge,type:decimal(19,2),nullzero"`
	CarrierFreightEstPercentage sql.NullFloat64 `bun:"carrier_freight_est_percentage,type:decimal(19,9),nullzero"`
	ExpeditedFreightOptionCd    sql.NullInt32   `bun:"expedited_freight_option_cd,type:int,nullzero"`
	OrderPriorityUid            sql.NullInt32   `bun:"order_priority_uid,type:int,nullzero"`
	DcDoNotRouteFlag            sql.NullString  `bun:"dc_do_not_route_flag,type:char,nullzero"`
	RoadnetPtPrintOption        sql.NullString  `bun:"roadnet_pt_print_option,type:char,nullzero"`
	SfdcAccountId               sql.NullString  `bun:"sfdc_account_id,type:varchar(255),nullzero"`
	SfdcCreateDate              sql.NullTime    `bun:"sfdc_create_date,type:datetime,nullzero"`
	SfdcUpdateDate              sql.NullTime    `bun:"sfdc_update_date,type:datetime,nullzero"`
	CarrierFixedFreightMarkup   sql.NullFloat64 `bun:"carrier_fixed_freight_markup,type:decimal(19,9),nullzero"`
	PhysCounty                  sql.NullString  `bun:"phys_county,type:varchar(255),nullzero"`
	ScacCode                    sql.NullString  `bun:"scac_code,type:varchar(4),nullzero"`
	CarrierPalletWeightLimit    sql.NullInt32   `bun:"carrier_pallet_weight_limit,type:int,nullzero"`
	CarrierPalletFreightCharge  sql.NullFloat64 `bun:"carrier_pallet_freight_charge,type:decimal(19,4),nullzero"`
	FreightCodeUid              sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	CarrierTransitDays          sql.NullInt32   `bun:"carrier_transit_days,type:int,nullzero"`
	FidelitoneCarrierId         sql.NullString  `bun:"fidelitone_carrier_id,type:varchar(255),nullzero"`
	MailAddress3                sql.NullString  `bun:"mail_address3,type:varchar(50),nullzero"`
	PhysAddress3                sql.NullString  `bun:"phys_address3,type:varchar(50),nullzero"`
	ShippingRouteUid            sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	FreightSurchargeFlag        string          `bun:"freight_surcharge_flag,type:char,default:('N')"`
	FreightSurchargePkgWeight   float64         `bun:"freight_surcharge_pkg_weight,type:decimal(19,9),default:((0))"`
	FreightSurchargePerPkg      float64         `bun:"freight_surcharge_per_pkg,type:decimal(19,9),default:((0))"`
	ExpressCarrierFlag          sql.NullString  `bun:"express_carrier_flag,type:char,default:('N')"`
	Latitude                    sql.NullString  `bun:"latitude,type:varchar(20),nullzero"`
	Longitude                   sql.NullString  `bun:"longitude,type:varchar(20),nullzero"`
	CarrierProviderTypeUid      sql.NullInt32   `bun:"carrier_provider_type_uid,type:int,nullzero"`
	RateGroupId                 sql.NullInt32   `bun:"rate_group_id,type:int,nullzero"`
	LtlFreightCalcPercentage    sql.NullFloat64 `bun:"ltl_freight_calc_percentage,type:decimal(19,9),default:((0))"`
	ParcelCarrierFlag           string          `bun:"parcel_carrier_flag,type:char,default:('N')"`
	DqDoNotRouteFlag            sql.NullString  `bun:"dq_do_not_route_flag,type:char,nullzero"`
	DqPrintPtFlag               sql.NullString  `bun:"dq_print_pt_flag,type:char,nullzero"`
	TransportationTypeCode      sql.NullString  `bun:"transportation_type_code,type:varchar(255),nullzero"`
	DisplayInOeFlag             sql.NullString  `bun:"display_in_oe_flag,type:char,nullzero"`
	NeighborhoodCode            sql.NullString  `bun:"neighborhood_code,type:varchar(255),nullzero"`
	CityCode                    sql.NullString  `bun:"city_code,type:varchar(255),nullzero"`
	LocalityCode                sql.NullString  `bun:"locality_code,type:varchar(255),nullzero"`
}

type AddressModel struct {
	bun bun.IDB
}

func (m *AddressModel) Get(ctx context.Context, id float64) (*Address, error) {
	address := new(Address)
	err := m.bun.NewSelect().Model(address).Where("id = ?", id).Scan(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return address, nil
}
