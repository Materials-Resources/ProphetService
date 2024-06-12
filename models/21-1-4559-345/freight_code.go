package model

type FreightCode struct {
	bun.BaseModel                 `bun:"table:freight_code"`
	FreightCodeUid                int32     `bun:"freight_code_uid,type:int,pk"`
	CompanyId                     string    `bun:"company_id,type:varchar(8)"`
	FreightCd                     string    `bun:"freight_cd,type:varchar(30)"`
	FreightDesc                   string    `bun:"freight_desc,type:varchar(255)"`
	IncomingFreight               string    `bun:"incoming_freight,type:char,default:('N')"`
	OutgoingFreight               string    `bun:"outgoing_freight,type:char,default:('N')"`
	IncomingReduceCommission      string    `bun:"incoming_reduce_commission,type:char,default:('N')"`
	OutgoingIncreaseCommission    string    `bun:"outgoing_increase_commission,type:char,default:('N')"`
	ProrateMethodCodeNo           int32     `bun:"prorate_method_code_no,type:int,default:(0)"`
	TaxGroupId                    string    `bun:"tax_group_id,type:varchar(10),nullzero"`
	RevenueAccountNo              string    `bun:"revenue_account_no,type:varchar(32)"`
	RowStatus                     int32     `bun:"row_status,type:int,default:(0)"`
	DateCreated                   time.Time `bun:"date_created,type:datetime,nullzero"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime,nullzero"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	FreeFreightBasisCd            int32     `bun:"free_freight_basis_cd,type:int,nullzero"`
	FreeInFreightMin              float64   `bun:"free_in_freight_min,type:decimal(19,9),nullzero"`
	FreeOutFreightMin             float64   `bun:"free_out_freight_min,type:decimal(19,9),nullzero"`
	DirectShipFreeFreightFlag     string    `bun:"direct_ship_free_freight_flag,type:char,nullzero"`
	FreeInFreightMinWeb           float64   `bun:"free_in_freight_min_web,type:decimal(19,9),nullzero"`
	FreeOutFreightMinWeb          float64   `bun:"free_out_freight_min_web,type:decimal(19,9),nullzero"`
	HandlingChargeOptionCd        int32     `bun:"handling_charge_option_cd,type:int,nullzero"`
	ExternalTaxProductCodeIn      string    `bun:"external_tax_product_code_in,type:varchar(255),nullzero"`
	ExternalTaxProductCodeOut     string    `bun:"external_tax_product_code_out,type:varchar(255),nullzero"`
	IncomingIncreaseCommission    string    `bun:"incoming_increase_commission,type:char,default:('N')"`
	PaySpecialFlag                string    `bun:"pay_special_flag,type:char,nullzero"`
	SkipFirstShipmentFlag         string    `bun:"skip_first_shipment_flag,type:char,nullzero"`
	ExcludeFromSalesMasterInquiry string    `bun:"exclude_from_sales_master_inquiry,type:char,default:('N')"`
	DeductibleFlag                string    `bun:"deductible_flag,type:char,nullzero"`
	FreeColdFreight               string    `bun:"free_cold_freight,type:char,default:('N')"`
	FreeHazmatFreight             string    `bun:"free_hazmat_freight,type:char,default:('N')"`
	FreeExpressFreight            string    `bun:"free_express_freight,type:char,default:('N')"`
	FreeBulkFreight               string    `bun:"free_bulk_freight,type:char,default:('N')"`
	FedexPaymentMethod            int32     `bun:"fedex_payment_method,type:int,nullzero"`
	ExcludeDiscountedFreight      string    `bun:"exclude_discounted_freight,type:char,default:('N')"`
}
