package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OrderTypeValue struct {
	bun.BaseModel     `bun:"table:order_type_value"`
	OrderTypeValueUid int32     `bun:"order_type_value_uid,type:int,autoincrement,identity,pk"`      // Unique Identifier for table.
	OrderTypeId       int32     `bun:"order_type_id,type:int"`                                       // Identifier for the order_type - this will be 1 - 11.
	RequiredFlag      string    `bun:"required_flag,type:char(1),default:('N')"`                     // Determine if this value will be required in OE.
	Value             string    `bun:"value,type:varchar(40)"`                                       // This will hold the value to appear in the drop down
	DefaultColumnFlag string    `bun:"default_column_flag,type:char(1),default:('N')"`               // Determine if this record is to be the default value for the given order_type
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`                                     // Indicates row status of this record.
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
