package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomObjectsDetail struct {
	bun.BaseModel          `bun:"table:custom_objects_detail"`
	CustomObjectsDetailUid int32     `bun:"custom_objects_detail_uid,type:int,pk"`                        // Unique ID for this table
	CustomObjectsUid       int32     `bun:"custom_objects_uid,type:int,unique"`                           // FK reference to custom_objects table
	SequenceNo             int32     `bun:"sequence_no,type:int"`                                         // The sequence of this command in the dynachange modification
	ObjectName             string    `bun:"object_name,type:varchar(255),unique"`                         // The name of the object to be modified
	AttributeName          string    `bun:"attribute_name,type:varchar(255),unique"`                      // The name of the attribute to be modified
	AttributeValue         string    `bun:"attribute_value,type:varchar(2048)"`                           // The modified value of the attribute
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`                                     // The status of the row
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
