package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardProcessor struct {
	bun.BaseModel                `bun:"table:creditcard_processor"`
	CreditcardProcessorUid       int32     `bun:"creditcard_processor_uid,type:int,pk"`                        // UID of this table
	LocationName                 *string   `bun:"location_name,type:varchar(255)"`                             // The name of the third party processing center that the CommerceCenter uses to complete creidt card transactions
	ProcessorName                *string   `bun:"processor_name,type:varchar(255)"`                            // Credit card processor center name.
	ProcessorNumber              *string   `bun:"processor_number,type:varchar(255)"`                          // ID of the credit card processor center
	MerchantId                   *string   `bun:"merchant_id,type:varchar(255)"`                               // ID number assigned by the third party credit card processor to identify the merchant
	Merchantkey                  *string   `bun:"merchantkey,type:varchar(255)"`                               // A code assigned by the third party credit card processor to identify the merchant
	HelpDeskPhone1               *string   `bun:"help_desk_phone1,type:varchar(20)"`                           // Help desk phone number 1 of credit card processor center
	HelpDeskPhone2               *string   `bun:"help_desk_phone2,type:varchar(20)"`                           // Help desk phone number 2 of credit card processor center
	VoiceAuthPhone1              *string   `bun:"voice_auth_phone1,type:varchar(20)"`                          // Phone number 1 of credit card processor center for voice authorization
	VoiceAuthPhone2              *string   `bun:"voice_auth_phone2,type:varchar(20)"`                          // Phone number 2 of credit card processor center for voice authorization
	IndustryType                 *string   `bun:"industry_type,type:varchar(255)"`                             // Industry type of CommerceCenter user.  For example Direct marketing or Retail
	ProtobasePath                *string   `bun:"protobase_path,type:varchar(255)"`                            // The location of Protobase application
	PbAdminPath                  *string   `bun:"pb_admin_path,type:varchar(255)"`                             // The location of Protobase application
	SettlementBatchFilePath      *string   `bun:"settlement_batch_file_path,type:varchar(255)"`                // The location of batch settlement
	RequestFilesPath             *string   `bun:"request_files_path,type:varchar(255)"`                        // Credit card transaction request file location
	ArchiveFilesPath             *string   `bun:"archive_files_path,type:varchar(255)"`                        // The path to where the CommerceCenter archives credit card request files and reponse files
	TerminalId                   *string   `bun:"terminal_id,type:varchar(255)"`                               // The ID of the terminal initiating credit card transaction
	Freight                      string    `bun:"freight,type:varchar(20),default:('15')"`                     // This field defines the percentage of an order amount that is an average freight estimate for all credit card authorization orders
	Timeout                      int32     `bun:"timeout,type:int,default:(45)"`                               // The number of seconds to connect to Protobase before a transaction times out
	DateCreated                  time.Time `bun:"date_created,type:datetime"`                                  // Indicates the date/time this record was created.
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30)"`                         // ID of the user who last maintained this record
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`                            // Indicates the date/time this record was last modified.
	ProcessorTypeCd              int32     `bun:"processor_type_cd,type:int,default:((1089))"`                 // Indicates type of payment associated with this processor
	WebServerUrl                 *string   `bun:"web_server_url,type:varchar(255)"`                            // Realex Web Server URL
	SecretKey                    *string   `bun:"secret_key,type:varchar(255)"`                                // Realex secret key for computing hash values
	EpaymentIntegrationTypeCd    *int32    `bun:"epayment_integration_type_cd,type:int"`                       // Electronic Payment Integration type
	ReportingUrl                 *string   `bun:"reporting_url,type:varchar(255)"`                             // The payment processor URL for reporting requests/responses
	ServicesUrl                  *string   `bun:"services_url,type:varchar(255)"`                              // The payment processor URL for services requests/responses
	TransactionUrl               *string   `bun:"transaction_url,type:varchar(255)"`                           // The payment processor URL for transaction requests/responses
	DefaultAccountFlag           string    `bun:"default_account_flag,type:char(1),default:('N')"`             // Indicates that the record is the default merchant services account with the processor
	Level3TransactionSupportFlag string    `bun:"level_3_transaction_support_flag,type:char(1),default:('N')"` // Indicates whether the payment processor supports Level III transactions
	ForceLevel3Mastercard        string    `bun:"force_level_3_mastercard,type:char(1),default:('N')"`         // Force Level 3 to be included on Mastercard transactions
	CurrencyId                   *float64  `bun:"currency_id,type:decimal(19,0)"`                              // Credit Card Processor Currency
}
