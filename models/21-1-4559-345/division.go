package model

type Division struct {
	bun.BaseModel             `bun:"table:division"`
	SupplierId                float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	DivisionId                float64   `bun:"division_id,type:decimal(19,0),pk"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DivisionName              string    `bun:"division_name,type:varchar(255),nullzero"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ReturnDivisionFlag        string    `bun:"return_division_flag,type:char,default:('N')"`
	DefaultCarrierId          float64   `bun:"default_carrier_id,type:decimal(19,0),nullzero"`
	DefaultFob                string    `bun:"default_fob,type:varchar(20),nullzero"`
	ChargeFreightToVendorFlag string    `bun:"charge_freight_to_vendor_flag,type:char,default:('N')"`
	DefaultAuthorizationNo    string    `bun:"default_authorization_no,type:varchar(255),nullzero"`
}
