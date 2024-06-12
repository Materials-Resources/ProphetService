package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseTransferGroup struct {
	bun.BaseModel             `bun:"table:purchase_transfer_group"`
	PurchaseTransferGroupId   string    `bun:"purchase_transfer_group_id,type:varchar(8),pk"`
	PurchaseTransferGroupDesc string    `bun:"purchase_transfer_group_desc,type:varchar(40)"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	RegionGroupType           string    `bun:"region_group_type,type:char,default:('N')"`
}
