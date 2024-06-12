package model

import (
	"github.com/uptrace/bun"
	"time"
)

type EpaymentResponseCodeInfo struct {
	bun.BaseModel             `bun:"table:epayment_response_code_info"`
	EpymntResponseCodeInfoUid int32     `bun:"epymnt_response_code_info_uid,type:int,pk,identity"`
	EpaymentIntegrationTypeCd int32     `bun:"epayment_integration_type_cd,type:int"`
	ResponseCodeTypeCd        int32     `bun:"response_code_type_cd,type:int"`
	ResponseCode              string    `bun:"response_code,type:varchar(255)"`
	ResponseCodeDescription   string    `bun:"response_code_description,type:varchar(255)"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
