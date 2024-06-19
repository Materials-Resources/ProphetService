package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ShippingIntegrationMsgHandling struct {
	bun.BaseModel                     `bun:"table:shipping_integration_msg_handling"`
	ShippingIntegrationMsgHandlingUid int32     `bun:"shipping_integration_msg_handling_uid,type:int,autoincrement,pk"` // Unique identifier for the shipping_integration_msg_handling table.
	ShippingIntegrationTypeCd         int32     `bun:"shipping_integration_type_cd,type:int,unique"`                    // Code that identifies the shipping integration that returns the message code/id.
	MessageNumberOrId                 string    `bun:"message_number_or_id,type:varchar(255),unique"`                   // The message number or id returned from the shipping integration.
	IgnoreErrorFlag                   string    `bun:"ignore_error_flag,type:char(1),default:('N')"`                    // Whether or not the application will ignore (that is, bypass displaying to the user) the returned message.
	MessageDescription                string    `bun:"message_description,type:varchar(255),nullzero"`                  // An optional description of the message or message handling.
	RowStatusFlag                     int32     `bun:"row_status_flag,type:int,default:((704))"`                        // Status of this record
	DateCreated                       time.Time `bun:"date_created,type:datetime,default:(getdate())"`                  // Date and time the record was originally created
	CreatedBy                         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`            // User who created the record
	DateLastModified                  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`            // Date and time the record was modified
	LastMaintainedBy                  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`    // User who last changed the record
}
