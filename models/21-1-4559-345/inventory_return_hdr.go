package model

type InventoryReturnHdr struct {
	bun.BaseModel               `bun:"table:inventory_return_hdr"`
	InventoryReturnHdrUid       int32     `bun:"inventory_return_hdr_uid,type:int,pk"`
	ReturnNumber                int32     `bun:"return_number,type:int"`
	CompanyId                   string    `bun:"company_id,type:varchar(8)"`
	VendorId                    float64   `bun:"vendor_id,type:decimal(19,0)"`
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0)"`
	DivisionId                  float64   `bun:"division_id,type:decimal(19,0)"`
	ReturnDate                  time.Time `bun:"return_date,type:datetime,default:(getdate())"`
	ExpectedDate                time.Time `bun:"expected_date,type:datetime,default:(getdate())"`
	FreightAmount               float64   `bun:"freight_amount,type:decimal(19,4),default:(0)"`
	FreightAmountDisplay        float64   `bun:"freight_amount_display,type:decimal(19,4),default:(0)"`
	RestockingFeeAmount         float64   `bun:"restocking_fee_amount,type:decimal(19,4),default:(0)"`
	RestockingFeeAmountDisplay  float64   `bun:"restocking_fee_amount_display,type:decimal(19,4),default:(0)"`
	RmaNumber                   string    `bun:"rma_number,type:varchar(255),nullzero"`
	PoReferenceNumber           float64   `bun:"po_reference_number,type:decimal(19,0),nullzero"`
	Ship2Name                   string    `bun:"ship2_name,type:varchar(255)"`
	Ship2Address1               string    `bun:"ship2_address1,type:varchar(255),nullzero"`
	Ship2Address2               string    `bun:"ship2_address2,type:varchar(255),nullzero"`
	Ship2City                   string    `bun:"ship2_city,type:varchar(255),nullzero"`
	Ship2State                  string    `bun:"ship2_state,type:varchar(255),nullzero"`
	Ship2Zip                    string    `bun:"ship2_zip,type:varchar(255),nullzero"`
	Ship2Country                string    `bun:"ship2_country,type:varchar(255),nullzero"`
	CurrencyId                  float64   `bun:"currency_id,type:decimal(19,0),default:(1)"`
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LocationId                  float64   `bun:"location_id,type:decimal(19,0)"`
	BuyerId                     string    `bun:"buyer_id,type:varchar(16)"`
	Ship2Phone                  string    `bun:"ship2_phone,type:varchar(20),nullzero"`
	CarrierId                   float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`
	CarrierTrackingNo           string    `bun:"carrier_tracking_no,type:varchar(255),nullzero"`
	ChargeFreightToVendor       string    `bun:"charge_freight_to_vendor,type:char,default:('Y')"`
	ContactId                   string    `bun:"contact_id,type:varchar(16),nullzero"`
	RetrievedByWms              string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	CurrencyLineUid             int32     `bun:"currency_line_uid,type:int,nullzero"`
	InvReceiptsClearingAcct     string    `bun:"inv_receipts_clearing_acct,type:varchar(32),nullzero"`
	Ship2Address3               string    `bun:"ship2_address3,type:varchar(255),nullzero"`
	BlindShipFlag               string    `bun:"blind_ship_flag,type:char,nullzero"`
	Fob                         string    `bun:"fob,type:varchar(20),nullzero"`
	TrackingNoList              string    `bun:"tracking_no_list,type:varchar(1000),nullzero"`
	ExpediteNotes               string    `bun:"expedite_notes,type:varchar(8000),nullzero"`
	WarrantyInventoryReturnFlag string    `bun:"warranty_inventory_return_flag,type:char,nullzero"`
}
