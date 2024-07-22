package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Address struct {
	bun.BaseModel               `bun:"table:address"`
	Id                          float64    `bun:"id,type:decimal(19,0),pk"`                                      // Unique identifier for the address
	Name                        string     `bun:"name,type:varchar(50)"`                                         // Name that identifies the address
	MailAddress1                *string    `bun:"mail_address1,type:varchar(50)"`                                // What is line 1 of the mailing address?
	MailAddress2                *string    `bun:"mail_address2,type:varchar(50)"`                                // What is line 2 of the mailing address?
	MailCity                    *string    `bun:"mail_city,type:varchar(50)"`                                    // What is the city for this mailing address?
	MailState                   *string    `bun:"mail_state,type:varchar(50)"`                                   // What is the postal state or province for this mailing address?
	MailPostalCode              *string    `bun:"mail_postal_code,type:varchar(10)"`                             // What is the postal code for this mailing address?
	MailCountry                 *string    `bun:"mail_country,type:varchar(50)"`                                 // What is the country for this mailing address?
	PhysAddress1                *string    `bun:"phys_address1,type:varchar(50)"`                                // What is the first line of the physical address for this address?
	PhysAddress2                *string    `bun:"phys_address2,type:varchar(50)"`                                // What is the second line of the physical address for this address?
	PhysCity                    *string    `bun:"phys_city,type:varchar(50)"`                                    // What is the city of the physical address for this address?
	PhysState                   *string    `bun:"phys_state,type:varchar(50)"`                                   // What is the state or province of the physical address for this address?
	PhysPostalCode              *string    `bun:"phys_postal_code,type:varchar(10)"`                             // What is the postal code of the physical address for this address?
	PhysCountry                 *string    `bun:"phys_country,type:varchar(50)"`                                 // What is the country of the physical address for this address?
	CentralPhoneNumber          *string    `bun:"central_phone_number,type:varchar(20)"`                         // What is the main telephone number for this address
	CentralFaxNumber            *string    `bun:"central_fax_number,type:varchar(20)"`                           // What is the main fax number for this address?
	FederalIdNumber             *string    `bun:"federal_id_number,type:varchar(255)"`                           // This column is unused.
	ResaleCertificate           *string    `bun:"resale_certificate,type:varchar(40)"`                           // This column is unused.
	UpsCode                     *string    `bun:"ups_code,type:varchar(40)"`                                     // This column is unused.
	CorpAddressId               *float64   `bun:"corp_address_id,type:decimal(19,0)"`                            // What is the corporate address id for this address?
	BillingAddressId            *float64   `bun:"billing_address_id,type:decimal(19,0)"`                         // What is the id of the billing address?
	CreditAddressId             *float64   `bun:"credit_address_id,type:decimal(19,0)"`                          // What is the credit address for this address?
	Customer                    *string    `bun:"customer,type:char(1)"`                                         // This column is unused.
	Vendor                      *string    `bun:"vendor,type:char(1)"`                                           // This column is unused.
	Employee                    *string    `bun:"employee,type:char(1)"`                                         // This column is unused.
	Prospect                    *string    `bun:"prospect,type:char(1)"`                                         // This column is unused.
	BillingAddress              *string    `bun:"billing_address,type:char(1)"`                                  // Where should the bills be sent to?
	ShippingAddress             *string    `bun:"shipping_address,type:char(1)"`                                 // This column is unused.
	PaymentAddress              *string    `bun:"payment_address,type:char(1)"`                                  // This column is unused.
	Incorporated                *string    `bun:"incorporated,type:char(1)"`                                     // Is this address incorporated?
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`     // ID of the user who last maintained this record
	DeleteFlag                  string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DefaultShipTime             *time.Time `bun:"default_ship_time,type:datetime"`                               // When should a shipment normally go out?
	PreferredLocationId         *float64   `bun:"preferred_location_id,type:decimal(19,0)"`                      // This column is unused.
	CarrierId                   *string    `bun:"carrier_id,type:varchar(10)"`                                   // What is the id of this carrier (if any)?
	DeliveryInstructions1       *string    `bun:"delivery_instructions1,type:varchar(20)"`                       // What are the special delivery instructions for this address -  if any?
	DeliveryInstructions2       *string    `bun:"delivery_instructions2,type:varchar(20)"`                       // What are the special delivery instructions for this address -  if any?
	DeliveryInstructions3       *string    `bun:"delivery_instructions3,type:varchar(20)"`                       // What are the special delivery instructions for this address -  if any?
	MorBegDelivery              *time.Time `bun:"mor_beg_delivery,type:datetime"`                                // When is the earliest that morning deliveries are accepted?
	MorEndDelivery              *time.Time `bun:"mor_end_delivery,type:datetime"`                                // When is the latest that morning deliveries are accepted?
	EveBegDelivery              *time.Time `bun:"eve_beg_delivery,type:datetime"`                                // When is the earliest that evening deliveries are accepted?
	EveEndDelivery              *time.Time `bun:"eve_end_delivery,type:datetime"`                                // When is the latest that evening deliveries are accepted?
	InventoryLocationFlag       *string    `bun:"inventory_location_flag,type:char(1)"`                          // This column is unused.
	ShowOutItems                *string    `bun:"show_out_items,type:char(1)"`                                   // This column is unused.
	StoreNo                     *string    `bun:"store_no,type:varchar(5)"`                                      // This column is unused.
	InvoiceType                 *string    `bun:"invoice_type,type:varchar(2)"`                                  // What type of invoice is this invoice?
	AddressIdString             *string    `bun:"address_id_string,type:varchar(22)"`                            // What is the unique identifer of this address -  in string form?
	CarrierFlag                 *string    `bun:"carrier_flag,type:char(1)"`                                     // Is this a carrier?
	DefaultCarrierId            *float64   `bun:"default_carrier_id,type:decimal(19,0)"`                         // What carrier is normally used?
	TradePercentDisc            *float64   `bun:"trade_percent_disc,type:decimal(19,0)"`                         // This column is unused.
	Class1Id                    *string    `bun:"class1_id,type:varchar(255)"`                                   // Address class 1
	Class2Id                    *string    `bun:"class2_id,type:varchar(255)"`                                   // Address class 2
	Class3Id                    *string    `bun:"class3_id,type:varchar(255)"`                                   // Address class 3
	Class4Id                    *string    `bun:"class4_id,type:varchar(255)"`                                   // Address class 4
	Class5Id                    *string    `bun:"class5_id,type:varchar(255)"`                                   // Address class 5
	CentralWattsNumber          *string    `bun:"central_watts_number,type:varchar(20)"`                         // What is the main WATTS number for this address?
	Alternative1099Name         *string    `bun:"alternative_1099_name,type:varchar(40)"`                        // An alternative name printed on IRS 1099 forms for this customer.
	NameControl                 *string    `bun:"name_control,type:varchar(4)"`                                  // This column is unused.
	DefaultShipToCompany        *string    `bun:"default_ship_to_company,type:varchar(8)"`                       // Which company should ship materials to this address?
	DefaultShipToBranch         *string    `bun:"default_ship_to_branch,type:varchar(8)"`                        // Which branch should ship materials to this address?
	ShipToPackingBasis          *string    `bun:"ship_to_packing_basis,type:varchar(16)"`                        // This column is unused.
	BillOfLadingType            *string    `bun:"bill_of_lading_type,type:varchar(2)"`                           // This column is unused.
	EmailAddress                *string    `bun:"email_address,type:varchar(255)"`                               // What is the email address for this contact?
	Url                         *string    `bun:"url,type:varchar(255)"`                                         // Web address for the particular address record - if available
	CarrierFedexFlag            string     `bun:"carrier_fedex_flag,type:char(1),default:('N')"`                 // Flag to indicate if a carrier is for Fedex use
	RouteviewRequireRoutingFlag string     `bun:"routeview_require_routing_flag,type:char(1),default:('N')"`     // Custom: determines if the Routeview (3rd party software) app will use this carrier when processing routing data.
	CarrierTypeCd               int32      `bun:"carrier_type_cd,type:int,default:((300))"`                      // Identifies the type of carrier associated with this address (will eventually supercede carrier_flag and carrier_fedex_flag)
	CarrierDoNotRouteFlag       *string    `bun:"carrier_do_not_route_flag,type:char(1),default:('N')"`          // This flag will be used to determine if a carrier can not be routed via geocom
	RoadnetDoNotRouteFlag       *string    `bun:"roadnet_do_not_route_flag,type:char(1)"`                        // Indicates that this carrier does not route to UPS Roadnet.
	CarrierStrategicFreightFlag *string    `bun:"carrier_strategic_freight_flag,type:char(1)"`                   // When checked, strategic freight charges will be computed for the given carrier.  If not checked, no strategic freight will be calculated.
	CarrierFixedFreightCharge   *float64   `bun:"carrier_fixed_freight_charge,type:decimal(19,2)"`               // Only available when carrier_strategic_freight_flag is not checked; allows distributor to charge a fixed rate for freight
	CarrierFreightEstPercentage *float64   `bun:"carrier_freight_est_percentage,type:decimal(19,9)"`             // Freight Estimated Percentage for the Carrier record
	ExpeditedFreightOptionCd    *int32     `bun:"expedited_freight_option_cd,type:int"`                          // Custom column indicates carrier's expedited freight option.
	OrderPriorityUid            *int32     `bun:"order_priority_uid,type:int"`                                   // holds information about order_priority
	DcDoNotRouteFlag            *string    `bun:"dc_do_not_route_flag,type:char(1)"`                             // Indicates if this carrier does uses Dispatch Compass routing.
	RoadnetPtPrintOption        *string    `bun:"roadnet_pt_print_option,type:char(1)"`                          // Indicates for a Carrier if a Pick ticket is to print prior to routing, so a user can indicate on a Ship To or Order if a Pick Ticket should print.
	SfdcAccountId               *string    `bun:"sfdc_account_id,type:varchar(255)"`                             // Salesforce.com account id - added by the import
	SfdcCreateDate              *time.Time `bun:"sfdc_create_date,type:datetime"`                                // Date the record was created in Salesforce.com - added by import.
	SfdcUpdateDate              *time.Time `bun:"sfdc_update_date,type:datetime"`                                // Date the record was updated in Salesforce.com - added by import.
	CarrierFixedFreightMarkup   *float64   `bun:"carrier_fixed_freight_markup,type:decimal(19,9)"`               // Fixed Freight Markup value
	PhysCounty                  *string    `bun:"phys_county,type:varchar(255)"`                                 // Physical county of the ship-to
	ScacCode                    *string    `bun:"scac_code,type:varchar(4)"`                                     // Standard Carrier Alpha Code
	CarrierPalletWeightLimit    *int32     `bun:"carrier_pallet_weight_limit,type:int"`                          // Custom: Max weight per pallet allowed.  Used to determin how many pallets use for freight charge.
	CarrierPalletFreightCharge  *float64   `bun:"carrier_pallet_freight_charge,type:decimal(19,4)"`              // Custom: Freight charge per pallet.  Used for outgoing freight when carrier_pallet_weigth_limit is set.
	FreightCodeUid              *int32     `bun:"freight_code_uid,type:int"`                                     // Freight code that applies to a carrier.
	CarrierTransitDays          *int32     `bun:"carrier_transit_days,type:int"`                                 // Transit Days for a Carrier
	FidelitoneCarrierId         *string    `bun:"fidelitone_carrier_id,type:varchar(255)"`                       // Custom: Code used by Fidelitone to identify this carrier
	MailAddress3                *string    `bun:"mail_address3,type:varchar(50)"`                                // Mailing address line 3
	PhysAddress3                *string    `bun:"phys_address3,type:varchar(50)"`                                // Physical address line 3
	ShippingRouteUid            *int32     `bun:"shipping_route_uid,type:int"`                                   // Unique identifier for the shipping route associated with this carrier.
	FreightSurchargeFlag        string     `bun:"freight_surcharge_flag,type:char(1),default:('N')"`             // Custom (F55882): determines if a carrier freight surcharge per package will be applied to shipments.
	FreightSurchargePkgWeight   float64    `bun:"freight_surcharge_pkg_weight,type:decimal(19,9),default:((0))"` // Custom (F55882): the value used to calculate the number of packages used to calculate a carrier freight surcharge.  Divided into the total shipment weight.
	FreightSurchargePerPkg      float64    `bun:"freight_surcharge_per_pkg,type:decimal(19,9),default:((0))"`    // Custom (F55882): the value used to calculate the total carrier freight surcharge to apply to a shipment.  Multiplied by the value of the total shipment weight divided by the freight_surcharge_pkg_weight.
	ExpressCarrierFlag          *string    `bun:"express_carrier_flag,type:char(1),default:('N')"`               // Indicate if carrier is express.
	Latitude                    *string    `bun:"latitude,type:varchar(20)"`
	Longitude                   *string    `bun:"longitude,type:varchar(20)"`
	CarrierProviderTypeUid      *int32     `bun:"carrier_provider_type_uid,type:int"` // If Agile Carrier, would have a Agile specific code.
	RateGroupId                 *int32     `bun:"rate_group_id,type:int"`
	LtlFreightCalcPercentage    *float64   `bun:"ltl_freight_calc_percentage,type:decimal(19,9),default:((0))"` // Percentage used by the system to calculate the correct LTL freight for the carrier.
	ParcelCarrierFlag           string     `bun:"parcel_carrier_flag,type:char(1),default:('N')"`               // Determines if a record is a parcel carrier (like UPS).
	DqDoNotRouteFlag            *string    `bun:"dq_do_not_route_flag,type:char(1)"`                            // Flag to not route DQ for this carrier
	DqPrintPtFlag               *string    `bun:"dq_print_pt_flag,type:char(1)"`                                // Flag to print DQ PTs regardless of other settings.
	TransportationTypeCode      *string    `bun:"transportation_type_code,type:varchar(255)"`                   // Custom: Indicates the transportation type code associated with carrier.
	DisplayInOeFlag             *string    `bun:"display_in_oe_flag,type:char(1)"`                              // Indicator for Trane whether to show field in OE window
	NeighborhoodCode            *string    `bun:"neighborhood_code,type:varchar(255)"`                          // MX: Neighborhood code associated with this address ID.
	CityCode                    *string    `bun:"city_code,type:varchar(255)"`                                  // MX: City code associated with this address ID.
	LocalityCode                *string    `bun:"locality_code,type:varchar(255)"`                              // MX: Locality code associated with this address ID.
}
