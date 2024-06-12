package model

type ShippingRoute struct {
	bun.BaseModel            `bun:"table:shipping_route"`
	ShippingRouteUid         int32     `bun:"shipping_route_uid,type:int,pk"`
	RouteCode                string    `bun:"route_code,type:varchar(255)"`
	RouteDescription         string    `bun:"route_description,type:varchar(255),default:('Shipping Route')"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	RowStatusFlag            int16     `bun:"row_status_flag,type:smallint,default:(704)"`
	DeliveryZoneUid          int32     `bun:"delivery_zone_uid,type:int,nullzero"`
	DowRouteFlag             string    `bun:"dow_route_flag,type:char,nullzero"`
	TodRouteFlag             string    `bun:"tod_route_flag,type:char,nullzero"`
	ShippingProgramFlag      string    `bun:"shipping_program_flag,type:char,nullzero"`
	ShipProgPickingThreshold float64   `bun:"ship_prog_picking_threshold,type:decimal(19,9),nullzero"`
	ShipProgThresholdBasisCd int32     `bun:"ship_prog_threshold_basis_cd,type:int,nullzero"`
	ShipProgPickingDays      string    `bun:"ship_prog_picking_days,type:varchar(10),nullzero"`
	ShipProgFreightDiscPct   float64   `bun:"ship_prog_freight_disc_pct,type:decimal(19,9),nullzero"`
	DeductibleExemptFlag     string    `bun:"deductible_exempt_flag,type:char,nullzero"`
	DeliveryDaysOfWeek       string    `bun:"delivery_days_of_week,type:varchar(255),nullzero"`
}
