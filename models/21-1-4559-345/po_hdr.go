package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PoHdr struct {
	bun.BaseModel                 `bun:"table:po_hdr"`
	PoNo                          float64   `bun:"po_no,type:decimal(19,0),pk"`
	VendorId                      float64   `bun:"vendor_id,type:decimal(19,0)"`
	OrderDate                     time.Time `bun:"order_date,type:datetime,nullzero"`
	CompanyNo                     string    `bun:"company_no,type:varchar(8)"`
	Ship2Name                     string    `bun:"ship2_name,type:varchar(50),nullzero"`
	PackingSlipNumber             string    `bun:"packing_slip_number,type:varchar(50),nullzero"`
	Ship2Add1                     string    `bun:"ship2_add1,type:varchar(50),nullzero"`
	Ship2Add2                     string    `bun:"ship2_add2,type:varchar(50),nullzero"`
	Ship2City                     string    `bun:"ship2_city,type:varchar(50),nullzero"`
	Ship2State                    string    `bun:"ship2_state,type:varchar(50),nullzero"`
	Ship2Country                  string    `bun:"ship2_country,type:varchar(50),nullzero"`
	RequestedBy                   string    `bun:"requested_by,type:varchar(20),nullzero"`
	Fob                           string    `bun:"fob,type:varchar(20),nullzero"`
	DateDue                       time.Time `bun:"date_due,type:datetime,nullzero"`
	ReceiptDate                   time.Time `bun:"receipt_date,type:datetime,nullzero"`
	PoDesc                        string    `bun:"po_desc,type:varchar(255),nullzero"`
	Ship2Zip                      string    `bun:"ship2_zip,type:varchar(10),nullzero"`
	DeleteFlag                    string    `bun:"delete_flag,type:char"`
	Complete                      string    `bun:"complete,type:char,nullzero"`
	DateCreated                   time.Time `bun:"date_created,type:datetime"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30)"`
	LocationId                    float64   `bun:"location_id,type:decimal(19,0)"`
	Terms                         string    `bun:"terms,type:varchar(20),nullzero"`
	CarrierId                     string    `bun:"carrier_id,type:varchar(8),nullzero"`
	VouchFlag                     string    `bun:"vouch_flag,type:char,nullzero"`
	Period                        float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod                 float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	CancelFlag                    string    `bun:"cancel_flag,type:char,nullzero"`
	CurrencyId                    float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	Printed                       string    `bun:"printed,type:char,nullzero"`
	AckDate                       time.Time `bun:"ack_date,type:datetime,nullzero"`
	AckCode                       string    `bun:"ack_code,type:varchar(2),nullzero"`
	ClosedFlag                    string    `bun:"closed_flag,type:char,nullzero"`
	SupplierId                    float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	DivisionId                    float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	BranchId                      string    `bun:"branch_id,type:varchar(8)"`
	Approved                      string    `bun:"approved,type:char"`
	PurchaseGroupId               string    `bun:"purchase_group_id,type:varchar(8),nullzero"`
	PoClass1                      string    `bun:"po_class1,type:varchar(255),nullzero"`
	PoClass2                      string    `bun:"po_class2,type:varchar(255),nullzero"`
	PoClass3                      string    `bun:"po_class3,type:varchar(255),nullzero"`
	PoClass4                      string    `bun:"po_class4,type:varchar(255),nullzero"`
	PoClass5                      string    `bun:"po_class5,type:varchar(255),nullzero"`
	SalesOrderNumber              string    `bun:"sales_order_number,type:varchar(8),nullzero"`
	PoType                        string    `bun:"po_type,type:char"`
	ExchangeRate                  float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`
	ExternalPoNo                  string    `bun:"external_po_no,type:varchar(40),nullzero"`
	ExcludeFromLeadTime           string    `bun:"exclude_from_lead_time,type:char,default:('N')"`
	SourceType                    int32     `bun:"source_type,type:int,nullzero"`
	TransmissionMethod            int32     `bun:"transmission_method,type:int,nullzero"`
	PoHdrUid                      int32     `bun:"po_hdr_uid,type:int"`
	EdiEditCount                  int16     `bun:"edi_edit_count,type:smallint,default:(0)"`
	RevisedPo                     string    `bun:"revised_po,type:char,nullzero"`
	ContactId                     string    `bun:"contact_id,type:varchar(16),nullzero"`
	CreatedBy                     string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TotalIvaTaxAmt                float64   `bun:"total_iva_tax_amt,type:decimal(19,4),nullzero"`
	ExpediteFlag                  string    `bun:"expedite_flag,type:char,nullzero"`
	VendorPaysFreightFlag         string    `bun:"vendor_pays_freight_flag,type:char,nullzero"`
	RetrievedByWms                string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	SupplierReleaseNo             string    `bun:"supplier_release_no,type:varchar(255),nullzero"`
	HistoricalCompleteFlag        string    `bun:"historical_complete_flag,type:char,default:('N')"`
	ExpectedDate                  time.Time `bun:"expected_date,type:datetime,nullzero"`
	SentToCarrierDate             time.Time `bun:"sent_to_carrier_date,type:datetime,nullzero"`
	PrintCanadianB3FormsFlag      string    `bun:"print_canadian_b3_forms_flag,type:char,default:('N')"`
	CanadianB3FormsDate           time.Time `bun:"canadian_b3_forms_date,type:datetime,nullzero"`
	CadChaContractCd              int32     `bun:"cad_cha_contract_cd,type:int,nullzero"`
	CadChaQuoteNo                 string    `bun:"cad_cha_quote_no,type:varchar(255),nullzero"`
	TrackingNo                    string    `bun:"tracking_no,type:char(40),nullzero"`
	NetBillingFlag                string    `bun:"net_billing_flag,type:char,default:('N')"`
	Ship2Add3                     string    `bun:"ship2_add3,type:varchar(50),nullzero"`
	DoNotExportCarrierPoFlag      string    `bun:"do_not_export_carrier_po_flag,type:char,default:('N')"`
	PackingBasis                  string    `bun:"packing_basis,type:varchar(255),nullzero"`
	PackingBasisViolationCd       int32     `bun:"packing_basis_violation_cd,type:int,nullzero"`
	TransmitPrint                 string    `bun:"transmit_print,type:char,nullzero"`
	TransmitFax                   string    `bun:"transmit_fax,type:char,nullzero"`
	TransmitEmail                 string    `bun:"transmit_email,type:char,nullzero"`
	TransmitEdi                   string    `bun:"transmit_edi,type:char,nullzero"`
	DfltListPriceMultiplier       float64   `bun:"dflt_list_price_multiplier,type:decimal(19,9),nullzero"`
	ShipViaDesc                   string    `bun:"ship_via_desc,type:varchar(255),nullzero"`
	BlindShipFlag                 string    `bun:"blind_ship_flag,type:char,nullzero"`
	PoQtyChangeCount              int32     `bun:"po_qty_change_count,type:int,nullzero"`
	BulkDiscountFlag              string    `bun:"bulk_discount_flag,type:char,nullzero"`
	EstimatedValue                float64   `bun:"estimated_value,type:decimal(19,9),nullzero"`
	EstimatedValueEditFlag        string    `bun:"estimated_value_edit_flag,type:char,nullzero"`
	BidNo                         string    `bun:"bid_no,type:varchar(255),nullzero"`
	ExpectedShipDate              time.Time `bun:"expected_ship_date,type:datetime,nullzero"`
	AutoVouchExceptFlag           string    `bun:"auto_vouch_except_flag,type:char,default:('N')"`
	SubmittedToVendorDate         time.Time `bun:"submitted_to_vendor_date,type:datetime,nullzero"`
	ExclSendLinkedInfoToRfnavFlag string    `bun:"excl_send_linked_info_to_rfnav_flag,type:char,default:('N')"`
}
