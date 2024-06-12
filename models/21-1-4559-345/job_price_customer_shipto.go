package model

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceCustomerShipto struct {
	bun.BaseModel                `bun:"table:job_price_customer_shipto"`
	JobPriceCustShiptoUid        int32     `bun:"job_price_cust_shipto_uid,type:int,pk"`
	JobPriceHdrUid               int32     `bun:"job_price_hdr_uid,type:int"`
	CompanyId                    string    `bun:"company_id,type:varchar(8)"`
	CustomerId                   float64   `bun:"customer_id,type:decimal(19,0)"`
	ShipToId                     float64   `bun:"ship_to_id,type:decimal(19,0)"`
	PreferredLocationId          float64   `bun:"preferred_location_id,type:decimal(19,0),nullzero"`
	TermsId                      string    `bun:"terms_id,type:varchar(2),nullzero"`
	ConsolidatedInvoicing        string    `bun:"consolidated_invoicing,type:char,nullzero"`
	AcceptableWaitTime           float64   `bun:"acceptable_wait_time,type:decimal(19,0),nullzero"`
	DefaultDisposition           string    `bun:"default_disposition,type:char,nullzero"`
	DefaultCarrierId             float64   `bun:"default_carrier_id,type:decimal(19,0),nullzero"`
	FreightCdUid                 int32     `bun:"freight_cd_uid,type:int,nullzero"`
	ContractDollarLimit          float64   `bun:"contract_dollar_limit,type:decimal(19,9),nullzero"`
	ShipToContactId              string    `bun:"ship_to_contact_id,type:varchar(16),nullzero"`
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag                int32     `bun:"row_status_flag,type:int"`
	LocationId                   float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	ContractDollarLimitUsed      float64   `bun:"contract_dollar_limit_used,type:decimal(19,9),default:(0)"`
	CustPoNo                     string    `bun:"cust_po_no,type:varchar(255),nullzero"`
	ContractOrderLimit           int32     `bun:"contract_order_limit,type:int,nullzero"`
	ContractOrderLimitUsed       int32     `bun:"contract_order_limit_used,type:int,nullzero"`
	ActivationDate               time.Time `bun:"activation_date,type:datetime,nullzero"`
	ExpirationDate               time.Time `bun:"expiration_date,type:datetime,nullzero"`
	AutoAssignedFlag             string    `bun:"auto_assigned_flag,type:char,default:('N')"`
	ExcludeFromFreightFactorFlag string    `bun:"exclude_from_freight_factor_flag,type:char,default:('N')"`
	Multiplier                   float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
}
