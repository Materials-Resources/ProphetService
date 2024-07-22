package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Carrier struct {
	bun.BaseModel           `bun:"table:carrier"`
	CarrierId               int32     `bun:"carrier_id,type:int,pk"`                                       // Unique identifier for this record, matches corresponding address ID in address table.
	ExtTaxFreightTaxCodeIn  *string   `bun:"ext_tax_freight_tax_code_in,type:varchar(255)"`                // The freight tax code to be used for 3rd party tax systems on incoming freight
	ExtTaxFreightTaxCodeOut *string   `bun:"ext_tax_freight_tax_code_out,type:varchar(255)"`               // The freight tax code to be used for 3rd party tax systems on outgoing freight
	DeliveryFeeInvMastUid   *int32    `bun:"delivery_fee_inv_mast_uid,type:int"`                           // Other charge item to be used to add delivery fee to orders/shipments
	FuelSurchargeInvMastUid *int32    `bun:"fuel_surcharge_inv_mast_uid,type:int"`                         // Other charge item to be used to add fuel surcharges to orders/shipments
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
