package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Carrier struct {
	bun.BaseModel           `bun:"table:carrier"`
	CarrierId               int32     `bun:"carrier_id,type:int,pk"`
	ExtTaxFreightTaxCodeIn  string    `bun:"ext_tax_freight_tax_code_in,type:varchar(255),nullzero"`
	ExtTaxFreightTaxCodeOut string    `bun:"ext_tax_freight_tax_code_out,type:varchar(255),nullzero"`
	DeliveryFeeInvMastUid   int32     `bun:"delivery_fee_inv_mast_uid,type:int,nullzero"`
	FuelSurchargeInvMastUid int32     `bun:"fuel_surcharge_inv_mast_uid,type:int,nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
