package model

type ShipToSalesrep struct {
	bun.BaseModel        `bun:"table:ship_to_salesrep"`
	CompanyId            string    `bun:"company_id,type:varchar(8),pk"`
	ShipToId             float64   `bun:"ship_to_id,type:decimal(19,0),pk"`
	SalesrepId           string    `bun:"salesrep_id,type:varchar(16),pk"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrimarySalesrep      string    `bun:"primary_salesrep,type:char,nullzero"`
	CommissionPercentage float64   `bun:"commission_percentage,type:decimal(19,4),default:((100))"`
	PrimaryServiceRep    string    `bun:"primary_service_rep,type:char,default:('N')"`
}
