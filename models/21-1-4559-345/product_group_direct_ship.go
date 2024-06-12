package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProductGroupDirectShip struct {
	bun.BaseModel             `bun:"table:product_group_direct_ship"`
	ProductGroupDirectShipUid int32     `bun:"product_group_direct_ship_uid,type:int,pk"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8)"`
	DirectShipStatus          string    `bun:"direct_ship_status,type:char,default:('S')"`
	DirectShipQty             float64   `bun:"direct_ship_qty,type:decimal(19,9),default:(0)"`
	DeleteFlag                string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OrderValue                float64   `bun:"order_value,type:decimal(19,9),default:((0))"`
}
