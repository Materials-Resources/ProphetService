package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type BankAccounts struct {
	bun.BaseModel             `bun:"table:bank_accounts"`
	CompanyNo                 string     `bun:"company_no,type:varchar(8),pk"`                             // Unique code that identifies a company.
	BankNo                    float64    `bun:"bank_no,type:decimal(19,0),pk"`                             // Default bank number
	TransitCode               string     `bun:"transit_code,type:varchar(15)"`                             // Enter the transit code
	AccountNo                 string     `bun:"account_no,type:varchar(18)"`                               // Enter an account number
	BankCode                  float64    `bun:"bank_code,type:decimal(18,0)"`                              // Enter the bank code
	FedResCode                string     `bun:"fed_res_code,type:varchar(10)"`                             // Enter the federal reserve code
	BankName                  *string    `bun:"bank_name,type:varchar(30)"`                                // Enter the bank name
	BankBranch                *string    `bun:"bank_branch,type:varchar(30)"`                              // Enter the bank branch
	BankAddress1              *string    `bun:"bank_address1,type:varchar(30)"`                            // Enter the bank address
	BankAddress2              *string    `bun:"bank_address2,type:varchar(30)"`                            // Enter the bank address
	LastCheckNo               float64    `bun:"last_check_no,type:decimal(18,0)"`                          // Enter the last check number
	GlAccountNo               string     `bun:"gl_account_no,type:varchar(32)"`                            // General Ledger Account number
	Balance                   float64    `bun:"balance,type:decimal(19,4),default:(0)"`                    // This is the account balance
	Active                    string     `bun:"active,type:char(1)"`                                       // When "active" is not checked -  the voucher will not
	DeleteFlag                string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated               time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CurrencyId                float64    `bun:"currency_id,type:decimal(19,0)"`                            // What is the unique currency identifier for this row
	CheckFormName             *string    `bun:"check_form_name,type:varchar(255)"`                         // Name of the form used to print checks if the defaults are not used
	BeginReconcileDate        *time.Time `bun:"begin_reconcile_date,type:datetime"`                        // The date that reconciliation functionality was turned on for the bank.
	ReconcileFlag             string     `bun:"reconcile_flag,type:char(1),default:('N')"`                 // Indicates whether the bank is enabled for reconciliation functionality.
	BeginReconciledThruPeriod *int16     `bun:"begin_reconciled_thru_period,type:smallint"`                // Indicates the period in which AP/AR was reconciled thru at the time when bank was initially marked to Reconcile (note: user specifies this).
	BeginReconciledThruYear   *int16     `bun:"begin_reconciled_thru_year,type:smallint"`                  // Indicates the year in which AP/AR was reconciled thru at the time when bank was initially marked to Reconcile (note: user specifies this).
	AddressTypeCd             int32      `bun:"address_type_cd,type:int"`                                  // Column will contain the code for the user preferred address type( physical or mailing) to be used when printing checks.
	UsePreprintedChecksFlag   *string    `bun:"use_preprinted_checks_flag,type:char(1)"`                   // Flag to indicate if the current bank account uses pre-printed checks.
	NumberOfSigs              *string    `bun:"number_of_sigs,type:char(1)"`                               // Number of signatures required on an A/P check
	Sig1Path                  *string    `bun:"sig1_path,type:varchar(255)"`                               // File path for activity for signature 1
	Sig2Path                  *string    `bun:"sig2_path,type:varchar(255)"`                               // File path for activity for signature 2
	AmountReqTwoSigs          *float64   `bun:"amount_req_two_sigs,type:decimal(19,9)"`                    // Two signatures required if check exceed this amount
	RoutingNumber             *string    `bun:"routing_number,type:varchar(255)"`                          // Bank routing number for this vendor
	CheckDigit                *string    `bun:"check_digit,type:char(1)"`                                  // Check digit that goes along with the routing number
	AchFileIdModifier         *string    `bun:"ach_file_id_modifier,type:char(1)"`                         // Alphanumeric modifier that is used to distinguish between multiple ACH files send to the same bank on the same day.
	AchFileIdModifierDate     *time.Time `bun:"ach_file_id_modifier_date,type:datetime"`                   // Last date that the file_id_modifier was updated. Lets us know if we have to use it or start over.
	CheckLanguageCd           *int32     `bun:"check_language_cd,type:int"`                                // The language in which the Check Amounts will be displayed
	ExportPath                *string    `bun:"export_path,type:varchar(255)"`                             // Default export path
	AchCompanyId              *string    `bun:"ach_company_id,type:varchar(40)"`                           // Company identifier for ACH payments.
	CeoCompany                *string    `bun:"ceo_company,type:varchar(40)"`                              // Identifier for credit card payments.
	CommercialCardNo          *string    `bun:"commercial_card_no,type:varchar(40)"`                       // Commercial card account number.
	EddBillerId               *string    `bun:"edd_biller_id,type:varchar(255)"`                           // Wells Fargo identifier for company.
	DocumentTemplateNo        *string    `bun:"document_template_no,type:varchar(40)"`                     // Determines check type for printing.
	BankCity                  *string    `bun:"bank_city,type:varchar(40)"`                                // City to be used for WF ACH data output.
	BankState                 *string    `bun:"bank_state,type:varchar(2)"`                                // State to be used for WF ACH data output.
	BankPostalCode            *string    `bun:"bank_postal_code,type:varchar(10)"`                         // Zip code to be used for WF ACH data output.
	RemoteId                  *string    `bun:"remote_id,type:varchar(8)"`                                 // Wells Fargo assigned RID.
	ApplicationId             *string    `bun:"application_id,type:varchar(8)"`                            // Wells Fargo assigned application BID.
	BankCountry               *string    `bun:"bank_country,type:char(2)"`                                 // Describes where the bank is located in order to tell if it should be used EFT or ACH
	BankCompanyIdNo           *string    `bun:"bank_company_id_no,type:varchar(255)"`                      // Company Identification Number for ACH for Bank Accounts
}
