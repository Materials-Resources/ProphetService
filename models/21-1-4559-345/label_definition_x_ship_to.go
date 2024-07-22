package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinitionXShipTo struct {
	bun.BaseModel             `bun:"table:label_definition_x_ship_to"`
	LabelDefinitionXShipToUid int32     `bun:"label_definition_x_ship_to_uid,type:int,pk"`                   // Unique ID for this record
	LabelDefinitionUid        int32     `bun:"label_definition_uid,type:int"`                                // UID of the label definition record
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`                                   // Company ID
	ShipToId                  float64   `bun:"ship_to_id,type:decimal(19,0)"`                                // Ship To ID
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`                                     // Status of this record
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
