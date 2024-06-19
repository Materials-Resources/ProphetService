package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type MassUpdateDefinitionDetail struct {
	bun.BaseModel                 `bun:"table:mass_update_definition_detail"`
	MassUpdateDefinitionDetailUid int32     `bun:"mass_update_definition_detail_uid,type:int,autoincrement,scanonly,pk"` // Unique Identifier of mass_update_definition table
	MassUpdateDefinitionUid       int32     `bun:"mass_update_definition_uid,type:int,unique"`                           // Unique identifier of the table
	ColumnName                    string    `bun:"column_name,type:varchar(255),unique"`                                 // Column name which will be exposed to the mass update
	IncludeFlag                   string    `bun:"include_flag,type:char(1)"`                                            // Allow access to the mass update
	RowStatusFlag                 int32     `bun:"row_status_flag,type:int"`                                             // Row status
	DateCreated                   time.Time `bun:"date_created,type:datetime,default:(getdate())"`                       // Date and time the record was originally created
	CreatedBy                     string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                 // User who created the record
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                 // Date and time the record was modified
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`         // User who last changed the record
}
