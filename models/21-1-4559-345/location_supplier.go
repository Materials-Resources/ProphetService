package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type LocationSupplier struct {
	bun.BaseModel           `bun:"table:location_supplier"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	LocationId              float64   `bun:"location_id,type:decimal(19,0),pk"`                         // What is the unique location identifier for this row
	SupplierId              float64   `bun:"supplier_id,type:decimal(19,0),pk"`                         // What supplier is participating in this relationship?
	ControlValue            string    `bun:"control_value,type:varchar(16),nullzero"`                   // Weight/Dollars/Volume
	TargetValue             float64   `bun:"target_value,type:decimal(19,9),nullzero"`                  // target amount in terms of control value to purchase on a PO
	ReviewCycle             float64   `bun:"review_cycle,type:decimal(3,0),nullzero"`                   // This column is unused.
	DefaultCarrier          float64   `bun:"default_carrier,type:decimal(19,0),nullzero"`               // Default carrier for shipments from this supplier to this location.
	DefaultShipInstructions string    `bun:"default_ship_instructions,type:varchar(255),nullzero"`      // Shipping instructions.
	AverageLeadTime         float64   `bun:"average_lead_time,type:decimal(3,0),nullzero"`              // What is the average lead time for this supplier?
	LeadTimeSource          string    `bun:"lead_time_source,type:varchar(8),nullzero"`                 // This column is unused.
	LeadTimeSafetyFactor    float64   `bun:"lead_time_safety_factor,type:decimal(19,9),nullzero"`       // This is added to the lead time to make sure a shipment arrives in time.
	DateOfLastReview        time.Time `bun:"date_of_last_review,type:datetime,nullzero"`                // Last time GPOR was ran for this location/supplier combo to determine requirements.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	FreightControlValue     string    `bun:"freight_control_value,type:varchar(16),nullzero"`           // This supplements the standard control value and allows the distributor to maintain separate thresholds for things like minimum order amonts and free freight.
	FreightTargetValue      float64   `bun:"freight_target_value,type:decimal(19,9),default:(0)"`       // This supplements the standard target  value and allows the distributor to maintain separate thresholds for things like minimum order amonts and free freight.
	CombineStockNsSpecialCd int32     `bun:"combine_stock_ns_special_cd,type:int,nullzero"`             // Override of supplier setting to default combine stocks and specials on same PO.
	DistributorAccountId    string    `bun:"distributor_account_id,type:varchar(255),nullzero"`         // Custom: Account for the distributor on the supplier system
	DistributorSoldToId     string    `bun:"distributor_sold_to_id,type:varchar(255),nullzero"`         // Custom: Sold To ID for the distributor on the supplier system
	DistributorShipToId     string    `bun:"distributor_ship_to_id,type:varchar(255),nullzero"`         // Custom: Ship To ID for the distributor on the supplier system
	PlantCode               string    `bun:"plant_code,type:varchar(255),nullzero"`                     // Custom: Code that is assigned to this location by the supplier.
	ShippingPoint           string    `bun:"shipping_point,type:varchar(255),nullzero"`                 // Custom: Shipping point code that is assigned to this location by the supplier.
	StorageLocation         string    `bun:"storage_location,type:varchar(255),nullzero"`               // Pegmost Storage Location
	CarrierCustomerId       string    `bun:"carrier_customer_id,type:varchar(255),nullzero"`            // Carrier Integration: Carrier assigned id for this company/location
	CarrierLocationId       string    `bun:"carrier_location_id,type:varchar(255),nullzero"`            // Carrier Integration: Carrier assigned id for this location
	C10SoldTo               string    `bun:"c10_sold_to,type:varchar(255),nullzero"`                    // Carrier Integration: Stored C10 Sold To in Location Tab in Supplier Maintenance
	C10ShipTo               string    `bun:"c10_ship_to,type:varchar(255),nullzero"`                    // Carrier Integration: Stored C10 Ship To in Location Tab in Supplier Maintenance
	SpaShipToId             string    `bun:"spa_ship_to_id,type:varchar(255),nullzero"`                 // Stored SPA Ship To ID in Location Tab in Supplier Maintenance
	ShipDays                int32     `bun:"ship_days,type:int,nullzero"`                               // Custom column to specify number of days until receipt from the supplier ship date at location level
	PegmostPaymentTypeId    float64   `bun:"pegmost_payment_type_id,type:decimal(19,0),nullzero"`       // Payment type for Pegmost Buyback
	PegmostVarianceAmount   float64   `bun:"pegmost_variance_amount,type:decimal(19,9),nullzero"`       // Invoice variance amount for Pegmost Buyback
	PegmostVarianceTypeFlag string    `bun:"pegmost_variance_type_flag,type:char(1),nullzero"`          // Variance type for pegmost variance amount (Amount or Percentage)
	PegmostBankNo           float64   `bun:"pegmost_bank_no,type:decimal(19,0),default:((0))"`          // Bank Number for Pegmost Buyback
	SupplierZoneCostId      string    `bun:"supplier_zone_cost_id,type:varchar(255),nullzero"`          // The zone for the location at this supplier
}
