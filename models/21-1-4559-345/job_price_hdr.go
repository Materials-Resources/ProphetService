package model

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceHdr struct {
	bun.BaseModel         `bun:"table:job_price_hdr"`
	JobPriceHdrUid        int32     `bun:"job_price_hdr_uid,type:int,pk"`
	JobNo                 string    `bun:"job_no,type:varchar(8)"`
	JobDescription        string    `bun:"job_description,type:varchar(255),nullzero"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	SalesLocId            float64   `bun:"sales_loc_id,type:decimal(19,0),nullzero"`
	ContractNo            string    `bun:"contract_no,type:varchar(255),nullzero"`
	ContactId             string    `bun:"contact_id,type:varchar(16),nullzero"`
	StartDate             time.Time `bun:"start_date,type:datetime"`
	EndDate               time.Time `bun:"end_date,type:datetime"`
	ShipToId              float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`
	Approved              string    `bun:"approved,type:char,nullzero"`
	Cancelled             string    `bun:"cancelled,type:char,nullzero"`
	Taker                 string    `bun:"taker,type:varchar(30)"`
	PoNo                  string    `bun:"po_no,type:varchar(50),nullzero"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	CorpAddressId         float64   `bun:"corp_address_id,type:decimal(19,0),nullzero"`
	SalesrepId            string    `bun:"salesrep_id,type:varchar(16),nullzero"`
	CurrencyLineUid       int32     `bun:"currency_line_uid,type:int,nullzero"`
	CurrencyConversion    string    `bun:"currency_conversion,type:char,default:('Y')"`
	ConsignmentFlag       string    `bun:"consignment_flag,type:char,default:('N')"`
	ContractTypeCd        int32     `bun:"contract_type_cd,type:int,default:((2005))"`
	ExtendedDesc          string    `bun:"extended_desc,type:varchar(255),nullzero"`
	NoOfPeriods           int32     `bun:"no_of_periods,type:int,nullzero"`
	CadChaContractCd      int32     `bun:"cad_cha_contract_cd,type:int,nullzero"`
	CadChaSupplierId      float64   `bun:"cad_cha_supplier_id,type:decimal(19,0),nullzero"`
	CadChaQuoteNo         string    `bun:"cad_cha_quote_no,type:varchar(255),nullzero"`
	CadChaLocationId      float64   `bun:"cad_cha_location_id,type:decimal(19,0),nullzero"`
	AnniversaryDate       time.Time `bun:"anniversary_date,type:datetime,nullzero"`
	UseTotesFlag          string    `bun:"use_totes_flag,type:char,nullzero"`
	HandHeldAddFlag       string    `bun:"hand_held_add_flag,type:char,nullzero"`
	NationalAcctFlag      string    `bun:"national_acct_flag,type:char,default:('N')"`
	UseVmiFlag            string    `bun:"use_vmi_flag,type:char,nullzero"`
	AutoAddNewShiptosFlag string    `bun:"auto_add_new_shiptos_flag,type:char,nullzero"`
}
