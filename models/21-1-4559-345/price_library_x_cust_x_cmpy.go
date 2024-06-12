package model

type PriceLibraryXCustXCmpy struct {
	bun.BaseModel         `bun:"table:price_library_x_cust_x_cmpy"`
	PriceLibXCustXCmpyUid int32     `bun:"price_lib_x_cust_x_cmpy_uid,type:int,pk"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0)"`
	PriceLibraryUid       int32     `bun:"price_library_uid,type:int"`
	SequenceNumber        int16     `bun:"sequence_number,type:smallint"`
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	WebBasedPricing       string    `bun:"web_based_pricing,type:char,default:('N')"`
	DistributorNetFlag    string    `bun:"distributor_net_flag,type:char,default:('N')"`
}
