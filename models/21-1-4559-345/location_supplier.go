package model

type LocationSupplier struct {
	bun.BaseModel           `bun:"table:location_supplier"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`
	LocationId              float64   `bun:"location_id,type:decimal(19,0),pk"`
	SupplierId              float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	ControlValue            string    `bun:"control_value,type:varchar(16),nullzero"`
	TargetValue             float64   `bun:"target_value,type:decimal(19,9),nullzero"`
	ReviewCycle             float64   `bun:"review_cycle,type:decimal(3,0),nullzero"`
	DefaultCarrier          float64   `bun:"default_carrier,type:decimal(19,0),nullzero"`
	DefaultShipInstructions string    `bun:"default_ship_instructions,type:varchar(255),nullzero"`
	AverageLeadTime         float64   `bun:"average_lead_time,type:decimal(3,0),nullzero"`
	LeadTimeSource          string    `bun:"lead_time_source,type:varchar(8),nullzero"`
	LeadTimeSafetyFactor    float64   `bun:"lead_time_safety_factor,type:decimal(19,9),nullzero"`
	DateOfLastReview        time.Time `bun:"date_of_last_review,type:datetime,nullzero"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	FreightControlValue     string    `bun:"freight_control_value,type:varchar(16),nullzero"`
	FreightTargetValue      float64   `bun:"freight_target_value,type:decimal(19,9),default:(0)"`
	CombineStockNsSpecialCd int32     `bun:"combine_stock_ns_special_cd,type:int,nullzero"`
	DistributorAccountId    string    `bun:"distributor_account_id,type:varchar(255),nullzero"`
	DistributorSoldToId     string    `bun:"distributor_sold_to_id,type:varchar(255),nullzero"`
	DistributorShipToId     string    `bun:"distributor_ship_to_id,type:varchar(255),nullzero"`
	PlantCode               string    `bun:"plant_code,type:varchar(255),nullzero"`
	ShippingPoint           string    `bun:"shipping_point,type:varchar(255),nullzero"`
	StorageLocation         string    `bun:"storage_location,type:varchar(255),nullzero"`
	CarrierCustomerId       string    `bun:"carrier_customer_id,type:varchar(255),nullzero"`
	CarrierLocationId       string    `bun:"carrier_location_id,type:varchar(255),nullzero"`
	C10SoldTo               string    `bun:"c10_sold_to,type:varchar(255),nullzero"`
	C10ShipTo               string    `bun:"c10_ship_to,type:varchar(255),nullzero"`
	SpaShipToId             string    `bun:"spa_ship_to_id,type:varchar(255),nullzero"`
	ShipDays                int32     `bun:"ship_days,type:int,nullzero"`
	PegmostPaymentTypeId    float64   `bun:"pegmost_payment_type_id,type:decimal(19,0),nullzero"`
	PegmostVarianceAmount   float64   `bun:"pegmost_variance_amount,type:decimal(19,9),nullzero"`
	PegmostVarianceTypeFlag string    `bun:"pegmost_variance_type_flag,type:char,nullzero"`
	PegmostBankNo           float64   `bun:"pegmost_bank_no,type:decimal(19,0),default:((0))"`
	SupplierZoneCostId      string    `bun:"supplier_zone_cost_id,type:varchar(255),nullzero"`
}
