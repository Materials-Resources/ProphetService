package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type EpaymentResponseCodeInfo struct {
	bun.BaseModel             `bun:"table:epayment_response_code_info"`
	EpymntResponseCodeInfoUid int32     `bun:"epymnt_response_code_info_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the epayment_response_code_info table
	EpaymentIntegrationTypeCd int32     `bun:"epayment_integration_type_cd,type:int"`                            // Code that identifies the electronic payments integration that returns the response code
	ResponseCodeTypeCd        int32     `bun:"response_code_type_cd,type:int"`                                   // The type of response code (AVS, CVV, etc.) - from code_p21 table
	ResponseCode              string    `bun:"response_code,type:varchar(255)"`                                  // The identifier returned by the integration as the response code (typically one character)
	ResponseCodeDescription   string    `bun:"response_code_description,type:varchar(255)"`                      // A description of the response code (usually found in API specifications for the integration provider)
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:((704))"`                         // Status of this record
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
