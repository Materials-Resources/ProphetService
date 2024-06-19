package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardTransresponse struct {
	bun.BaseModel              `bun:"table:creditcard_transresponse"`
	CreditcardTransresponseUid int32     `bun:"creditcard_transresponse_uid,type:int,pk"`                    // UID of the table
	ResponsetypeId             int32     `bun:"responsetype_id,type:int"`                                    // ID of a credit card transaction type.  It is from creditcard_transtype
	PaymentNumber              float64   `bun:"payment_number,type:decimal(19,0),nullzero"`                  // Reference number to payment information
	RespfileLocation           string    `bun:"respfile_location,type:varchar(255),nullzero"`                // The path to where response file was created
	ResponseStatus             string    `bun:"response_status,type:varchar(10),nullzero"`                   // Status of response from Protobase to CommerceCenter
	ResponseDate               time.Time `bun:"response_date,type:datetime,nullzero"`                        // Date and time of the response from Protobase
	PbResponseCode             string    `bun:"pb_response_code,type:varchar(50),nullzero"`                  // A response code of credit card transaction from Protobase
	PbResponseMessage          string    `bun:"pb_response_message,type:varchar(255),nullzero"`              // A response message of credit card transaction from Protobase
	HostRefNumber              string    `bun:"host_ref_number,type:varchar(50),nullzero"`                   // A reference number from third party processor center.  The number is used by the center to refer to the credit card transaction.
	HostResponseCode           string    `bun:"host_response_code,type:varchar(50),nullzero"`                // A response code of credit card transaction from third party credit card processor center
	HostResponseMessage        string    `bun:"host_response_message,type:varchar(255),nullzero"`            // A response message of credit card transaction from third party credit card processor center
	AvsReturnCode              string    `bun:"avs_return_code,type:varchar(50),nullzero"`                   // Address verification code returned from third party credit card processor center
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`              // Indicates the date/time this record was created.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"` // ID of the user who last maintained this record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`        // Indicates the date/time this record was last modified.
	CvvResponseCode            string    `bun:"cvv_response_code,type:varchar(50),nullzero"`                 // CVV Response Code returned by payment processing host
}
