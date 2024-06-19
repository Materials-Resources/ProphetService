package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceHdr struct {
	bun.BaseModel         `bun:"table:job_price_hdr"`
	JobPriceHdrUid        int32     `bun:"job_price_hdr_uid,type:int,pk"`                   // System generated. Uniquely distinguishes each row.
	JobNo                 string    `bun:"job_no,type:varchar(8),unique"`                   // Unique & System Generated.
	JobDescription        string    `bun:"job_description,type:varchar(255),nullzero"`      // Description for the job.
	CompanyId             string    `bun:"company_id,type:varchar(8)"`                      // Unique code that identifies a company.
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0),nullzero"`         // Customer for this job?
	SalesLocId            float64   `bun:"sales_loc_id,type:decimal(19,0),nullzero"`        // Location identifier for this job
	ContractNo            string    `bun:"contract_no,type:varchar(255),nullzero"`          // Is unique within the location & customer
	ContactId             string    `bun:"contact_id,type:varchar(16),nullzero"`            // Who is the contact for this job?
	StartDate             time.Time `bun:"start_date,type:datetime"`                        // Starting date of the job
	EndDate               time.Time `bun:"end_date,type:datetime"`                          // Ending date of the job
	ShipToId              float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`          // Ship to id for the job
	Approved              string    `bun:"approved,type:char(1),nullzero"`                  // Approved Status of the job.(Y or N)
	Cancelled             string    `bun:"cancelled,type:char(1),nullzero"`                 // Is the job canceled? (Y Or N)
	Taker                 string    `bun:"taker,type:varchar(30)"`                          // Who took the job order?
	PoNo                  string    `bun:"po_no,type:varchar(50),nullzero"`                 // Po No for the job.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                        // Indicates current record status.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                // Indicates the date/time this record was last modified.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                      // Indicates the date/time this record was created.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`             // ID of the user who last maintained this record
	CorpAddressId         float64   `bun:"corp_address_id,type:decimal(19,0),nullzero"`     // Corporate Address for this Contract.
	SalesrepId            string    `bun:"salesrep_id,type:varchar(16),nullzero"`           // Sales Person identifier for this Contract.
	CurrencyLineUid       int32     `bun:"currency_line_uid,type:int,nullzero"`             // Corporate currency identifier for the table.
	CurrencyConversion    string    `bun:"currency_conversion,type:char(1),default:('Y')"`  // Tells a job/contract is available for currency conversion.
	ConsignmentFlag       string    `bun:"consignment_flag,type:char(1),default:('N')"`     // Indicates whether this record is for a consignment contract
	ContractTypeCd        int32     `bun:"contract_type_cd,type:int,default:((2005))"`      // Identifies the type of contract. Available values are Standard & Vendor.
	ExtendedDesc          string    `bun:"extended_desc,type:varchar(255),nullzero"`        // To store extended description.
	NoOfPeriods           int32     `bun:"no_of_periods,type:int,nullzero"`                 // Custom (F34176): defines the number of budget periods for this job.
	CadChaContractCd      int32     `bun:"cad_cha_contract_cd,type:int,nullzero"`           // Code to identify whether the contract uses CAD or CHA pricing
	CadChaSupplierId      float64   `bun:"cad_cha_supplier_id,type:decimal(19,0),nullzero"` // Determines the supplier associated with a CAD or CHA contract.
	CadChaQuoteNo         string    `bun:"cad_cha_quote_no,type:varchar(255),nullzero"`     // The number associated with the CAD or CHA agreement.
	CadChaLocationId      float64   `bun:"cad_cha_location_id,type:decimal(19,0),nullzero"` // The location to which virtual inventory will post for this contract's CAD/CHA transactions
	AnniversaryDate       time.Time `bun:"anniversary_date,type:datetime,nullzero"`         // Anniversary Date
	UseTotesFlag          string    `bun:"use_totes_flag,type:char(1),nullzero"`            // Indicates if quantity from HHBM is in totes or not
	HandHeldAddFlag       string    `bun:"hand_held_add_flag,type:char(1),nullzero"`        // Indicates quantities entered in HHBM should be added not replaced
	NationalAcctFlag      string    `bun:"national_acct_flag,type:char(1),default:('N')"`   // Custom (F63764): determines if this contract is for a national account.
	UseVmiFlag            string    `bun:"use_vmi_flag,type:char(1),nullzero"`              // flag to mark this contract to be used with vmi functionality
	AutoAddNewShiptosFlag string    `bun:"auto_add_new_shiptos_flag,type:char(1),nullzero"` // Custom column to indicate if new Ship tos for the Customers on the contract will be added automatically to the contract
}
