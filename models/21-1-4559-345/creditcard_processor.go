package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardProcessor struct {
	bun.BaseModel                `bun:"table:creditcard_processor"`
	CreditcardProcessorUid       int32     `bun:"creditcard_processor_uid,type:int,pk"`
	LocationName                 string    `bun:"location_name,type:varchar(255),nullzero"`
	ProcessorName                string    `bun:"processor_name,type:varchar(255),nullzero"`
	ProcessorNumber              string    `bun:"processor_number,type:varchar(255),nullzero"`
	MerchantId                   string    `bun:"merchant_id,type:varchar(255),nullzero"`
	Merchantkey                  string    `bun:"merchantkey,type:varchar(255),nullzero"`
	HelpDeskPhone1               string    `bun:"help_desk_phone1,type:varchar(20),nullzero"`
	HelpDeskPhone2               string    `bun:"help_desk_phone2,type:varchar(20),nullzero"`
	VoiceAuthPhone1              string    `bun:"voice_auth_phone1,type:varchar(20),nullzero"`
	VoiceAuthPhone2              string    `bun:"voice_auth_phone2,type:varchar(20),nullzero"`
	IndustryType                 string    `bun:"industry_type,type:varchar(255),nullzero"`
	ProtobasePath                string    `bun:"protobase_path,type:varchar(255),nullzero"`
	PbAdminPath                  string    `bun:"pb_admin_path,type:varchar(255),nullzero"`
	SettlementBatchFilePath      string    `bun:"settlement_batch_file_path,type:varchar(255),nullzero"`
	RequestFilesPath             string    `bun:"request_files_path,type:varchar(255),nullzero"`
	ArchiveFilesPath             string    `bun:"archive_files_path,type:varchar(255),nullzero"`
	TerminalId                   string    `bun:"terminal_id,type:varchar(255),nullzero"`
	Freight                      string    `bun:"freight,type:varchar(20),default:('15')"`
	Timeout                      int32     `bun:"timeout,type:int,default:(45)"`
	DateCreated                  time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30)"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`
	ProcessorTypeCd              int32     `bun:"processor_type_cd,type:int,default:((1089))"`
	WebServerUrl                 string    `bun:"web_server_url,type:varchar(255),nullzero"`
	SecretKey                    string    `bun:"secret_key,type:varchar(255),nullzero"`
	EpaymentIntegrationTypeCd    int32     `bun:"epayment_integration_type_cd,type:int,nullzero"`
	ReportingUrl                 string    `bun:"reporting_url,type:varchar(255),nullzero"`
	ServicesUrl                  string    `bun:"services_url,type:varchar(255),nullzero"`
	TransactionUrl               string    `bun:"transaction_url,type:varchar(255),nullzero"`
	DefaultAccountFlag           string    `bun:"default_account_flag,type:char,default:('N')"`
	Level3TransactionSupportFlag string    `bun:"level_3_transaction_support_flag,type:char,default:('N')"`
	ForceLevel3Mastercard        string    `bun:"force_level_3_mastercard,type:char,default:('N')"`
	CurrencyId                   float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
}
