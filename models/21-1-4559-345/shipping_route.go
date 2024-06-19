package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ShippingRoute struct {
	bun.BaseModel            `bun:"table:shipping_route"`
	ShippingRouteUid         int32     `bun:"shipping_route_uid,type:int,pk"`                                 // Unique Identifier.
	RouteCode                string    `bun:"route_code,type:varchar(255)"`                                   // User defined code for the route.
	RouteDescription         string    `bun:"route_description,type:varchar(255),default:('Shipping Route')"` // User defined extended description for the route.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                     // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                               // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`      // ID of the user who last maintained this record
	RowStatusFlag            int16     `bun:"row_status_flag,type:smallint,default:(704)"`                    // Indicates current record status.
	DeliveryZoneUid          int32     `bun:"delivery_zone_uid,type:int,nullzero"`                            // Indicates if there is a delivery zone for this ship_to record - Foreign key to table Delivery_Zone
	DowRouteFlag             string    `bun:"dow_route_flag,type:char(1),nullzero"`                           // Indicates whether a shipping routes uses day of week logic.
	TodRouteFlag             string    `bun:"tod_route_flag,type:char(1),nullzero"`                           // Indicates whether a shipping route uses time of day logic.
	ShippingProgramFlag      string    `bun:"shipping_program_flag,type:char(1),nullzero"`                    // Indicates that this shipping route uses a shipping program that determines when associated orders should be picked
	ShipProgPickingThreshold float64   `bun:"ship_prog_picking_threshold,type:decimal(19,9),nullzero"`        // The threshold amount that determines whether orders with this route should be picked.
	ShipProgThresholdBasisCd int32     `bun:"ship_prog_threshold_basis_cd,type:int,nullzero"`                 // Code value indicating whether the threshold amount is in terms of weight or sales amount.
	ShipProgPickingDays      string    `bun:"ship_prog_picking_days,type:varchar(10),nullzero"`               // Concatenated list of day numbers for the days of the week that orders with this route will be picked.
	ShipProgFreightDiscPct   float64   `bun:"ship_prog_freight_disc_pct,type:decimal(19,9),nullzero"`         // Discount percentage applied to freight amounts when this shipping route is used.
	DeductibleExemptFlag     string    `bun:"deductible_exempt_flag,type:char(1),nullzero"`                   // Indicates whether this shipping route will not be using the deductible freight code. (Custom feature)
	DeliveryDaysOfWeek       string    `bun:"delivery_days_of_week,type:varchar(255),nullzero"`               // Custom: Concatenated and undelimited list of day numbers (1 = Sunday, 2 = Monday, etc) that indicates which days of week the specified shipping route will deliver.
}
