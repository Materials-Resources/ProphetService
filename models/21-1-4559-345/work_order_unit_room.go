package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type WorkOrderUnitRoom struct {
	bun.BaseModel    `bun:"table:work_order_unit_room"`
	WorkOrderRoomUid int32     `bun:"work_order_room_uid,type:int,pk"`                              // Unique Id for the work order unit room table
	RoomDesc         string    `bun:"room_desc,type:varchar(255)"`                                  // The name of the room specified
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Status of this record
}
