package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ShippingIntegrationMsgHandling struct {
	bun.BaseModel                     `bun:"table:shipping_integration_msg_handling"`
	ShippingIntegrationMsgHandlingUid int32     `bun:"shipping_integration_msg_handling_uid,type:int,pk,identity"`
	ShippingIntegrationTypeCd         int32     `bun:"shipping_integration_type_cd,type:int"`
	MessageNumberOrId                 string    `bun:"message_number_or_id,type:varchar(255)"`
	IgnoreErrorFlag                   string    `bun:"ignore_error_flag,type:char,default:('N')"`
	MessageDescription                string    `bun:"message_description,type:varchar(255),nullzero"`
	RowStatusFlag                     int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated                       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
