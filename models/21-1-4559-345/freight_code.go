package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FreightCode struct {
	bun.BaseModel                 `bun:"table:freight_code"`
	FreightCodeUid                int32     `bun:"freight_code_uid,type:int,pk"`                                 // Unique record identifier
	CompanyId                     string    `bun:"company_id,type:varchar(8)"`                                   // Unique code that identifies a company.
	FreightCd                     string    `bun:"freight_cd,type:varchar(30)"`                                  // Unique code that identifies a particular freight type
	FreightDesc                   string    `bun:"freight_desc,type:varchar(255)"`                               // Free form description describing the current freight code
	IncomingFreight               string    `bun:"incoming_freight,type:char(1),default:('N')"`                  // Indicates whether incoming freight is charged to the customer
	OutgoingFreight               string    `bun:"outgoing_freight,type:char(1),default:('N')"`                  // Indicates whether outgoing freight is charged to the customer
	IncomingReduceCommission      string    `bun:"incoming_reduce_commission,type:char(1),default:('N')"`        // Indicates whether commission cost is reduced by incoming freight
	OutgoingIncreaseCommission    string    `bun:"outgoing_increase_commission,type:char(1),default:('N')"`      // Indicates whether commission cost is increased by outgoing freight
	ProrateMethodCodeNo           int32     `bun:"prorate_method_code_no,type:int,default:(0)"`                  // The method by which freight is prorated accross invoice lines to affect commission cost
	TaxGroupId                    string    `bun:"tax_group_id,type:varchar(10),nullzero"`                       // Indicates the tax group identification.
	RevenueAccountNo              string    `bun:"revenue_account_no,type:varchar(32)"`                          // Account that shows the gross increase in your company’s income caused by freight
	RowStatus                     int32     `bun:"row_status,type:int,default:(0)"`                              // Indicates current record status.
	DateCreated                   time.Time `bun:"date_created,type:datetime,nullzero"`                          // Indicates the date/time this record was created.
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime,nullzero"`                    // Indicates the date/time this record was last modified.
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`    // ID of the user who last maintained this record
	FreeFreightBasisCd            int32     `bun:"free_freight_basis_cd,type:int,nullzero"`                      // What is free freight based on (order or shipment)
	FreeInFreightMin              float64   `bun:"free_in_freight_min,type:decimal(19,9),nullzero"`              // Incoming minimum for free freight
	FreeOutFreightMin             float64   `bun:"free_out_freight_min,type:decimal(19,9),nullzero"`             // Outgoing minimum for free freight
	DirectShipFreeFreightFlag     string    `bun:"direct_ship_free_freight_flag,type:char(1),nullzero"`          // Allow free freight on Direct Shipments flag
	FreeInFreightMinWeb           float64   `bun:"free_in_freight_min_web,type:decimal(19,9),nullzero"`          // Incoming minimum for free freight via web orders
	FreeOutFreightMinWeb          float64   `bun:"free_out_freight_min_web,type:decimal(19,9),nullzero"`         // Outgoing minimum for free freight via web orders
	HandlingChargeOptionCd        int32     `bun:"handling_charge_option_cd,type:int,nullzero"`                  // The value indicates how to process freight and handling charges for direct shipping
	ExternalTaxProductCodeIn      string    `bun:"external_tax_product_code_in,type:varchar(255),nullzero"`      // The product code to be used for 3rd party tax systems on incoming freight
	ExternalTaxProductCodeOut     string    `bun:"external_tax_product_code_out,type:varchar(255),nullzero"`     // The product code to be used for 3rd party tax systems on outgoing freight
	IncomingIncreaseCommission    string    `bun:"incoming_increase_commission,type:char(1),default:('N')"`      // Indicates whether commission cost is increased by incoming freight
	PaySpecialFlag                string    `bun:"pay_special_flag,type:char(1),nullzero"`                       // For the UPS ConnectShip integration, indicates that freight will be charged for 'M' class items only when the freight minimum is met
	SkipFirstShipmentFlag         string    `bun:"skip_first_shipment_flag,type:char(1),nullzero"`               // For the UPS ConnectShip integration, indicates that freight will not be paid for special items (as determiend by the item class) for the first shipment of each month
	ExcludeFromSalesMasterInquiry string    `bun:"exclude_from_sales_master_inquiry,type:char(1),default:('N')"` // Exclude this freight code from free freight calculations
	DeductibleFlag                string    `bun:"deductible_flag,type:char(1),nullzero"`                        // Indicates whether this freight code is deductible type (Custom Feature)
	FreeColdFreight               string    `bun:"free_cold_freight,type:char(1),default:('N')"`                 // Set a freight code to provide free cold box charge.
	FreeHazmatFreight             string    `bun:"free_hazmat_freight,type:char(1),default:('N')"`               // Set a freight code to provide free hazmat box charge.
	FreeExpressFreight            string    `bun:"free_express_freight,type:char(1),default:('N')"`              // Set a freight code to provide free express charge.
	FreeBulkFreight               string    `bun:"free_bulk_freight,type:char(1),default:('N')"`                 // Set a freight code to provide free bulk charge.
	FedexPaymentMethod            int32     `bun:"fedex_payment_method,type:int,nullzero"`                       // Fedex payment method
	ExcludeDiscountedFreight      string    `bun:"exclude_discounted_freight,type:char(1),default:('N')"`        // Exclude Discounted Freight Flag
}
