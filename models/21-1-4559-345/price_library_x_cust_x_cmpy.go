package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceLibraryXCustXCmpy struct {
	bun.BaseModel         `bun:"table:price_library_x_cust_x_cmpy"`
	PriceLibXCustXCmpyUid int32     `bun:"price_lib_x_cust_x_cmpy_uid,type:int,pk"`             // Internal record identifier
	CompanyId             string    `bun:"company_id,type:varchar(8),unique"`                   // Unique code that identifies a company.
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0),unique"`               // Customer ID
	PriceLibraryUid       int32     `bun:"price_library_uid,type:int"`                          // Internal record identifier for the Sales Pricing Library
	SequenceNumber        int16     `bun:"sequence_number,type:smallint,unique"`                // Sequence number of  this Sales Pricing Library for the Customer
	RowStatusFlag         int16     `bun:"row_status_flag,type:smallint"`                       // Indicates current record status.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                    // Indicates the date/time this record was last modified.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                          // Indicates the date/time this record was created.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`                 // ID of the user who last maintained this record
	WebBasedPricing       string    `bun:"web_based_pricing,type:char(1),default:('N'),unique"` // When Yes, signifies that the current row is used to price items placed via web based orders.
	DistributorNetFlag    string    `bun:"distributor_net_flag,type:char(1),default:('N')"`     // Distributor Net flag
}
