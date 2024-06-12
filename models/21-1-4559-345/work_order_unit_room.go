package model

import (
	"github.com/uptrace/bun"
	"time"
)

type WorkOrderUnitRoom struct {
	bun.BaseModel    `bun:"table:work_order_unit_room"`
	WorkOrderRoomUid int32     `bun:"work_order_room_uid,type:int,pk"`
	RoomDesc         string    `bun:"room_desc,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
}
