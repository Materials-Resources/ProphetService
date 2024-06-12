package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseTransferLocations struct {
	bun.BaseModel           `bun:"table:purchase_transfer_locations"`
	PurchaseTransferGroupId string    `bun:"purchase_transfer_group_id,type:varchar(8),pk"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`
	LocationId              float64   `bun:"location_id,type:decimal(19,0),pk"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	SequenceNumber          float64   `bun:"sequence_number,type:decimal(3,0),nullzero"`
}
