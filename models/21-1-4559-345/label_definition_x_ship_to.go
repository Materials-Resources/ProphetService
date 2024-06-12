package model

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinitionXShipTo struct {
	bun.BaseModel             `bun:"table:label_definition_x_ship_to"`
	LabelDefinitionXShipToUid int32     `bun:"label_definition_x_ship_to_uid,type:int,pk"`
	LabelDefinitionUid        int32     `bun:"label_definition_uid,type:int"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	ShipToId                  float64   `bun:"ship_to_id,type:decimal(19,0)"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
