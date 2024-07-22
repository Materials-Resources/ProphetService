package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FcDataobjectColumn struct {
	bun.BaseModel         `bun:"table:fc_dataobject_column"`
	FcDataobjectColumnUid int32     `bun:"fc_dataobject_column_uid,type:int,pk"`                         // Unique Identifier for fc_dataobject_column record.
	FcDataobjectTableUid  int32     `bun:"fc_dataobject_table_uid,type:int"`                             // Unique Identifier for the table that this column has been created for.
	UserColumnName        string    `bun:"user_column_name,type:varchar(255)"`                           // A user-defined column name for this column/field.
	ColumnName            string    `bun:"column_name,type:varchar(255)"`                                // The column name for this record that will show up in the P21 application.
	UpdateableFlag        string    `bun:"updateable_flag,type:char(1),default:('N')"`                   // Flag to indicate if the column value will be updated to the database.
	UpdatewhereclauseFlag string    `bun:"updatewhereclause_flag,type:char(1),default:('N')"`            // Flag to indicate if the column value will be part of the where clause of the update statement.
	DataType              *string   `bun:"data_type,type:varchar(255)"`                                  // Data type for this column.
	TypeCd                int32     `bun:"type_cd,type:int"`                                             // Type of field, db column, db computed column, computed field.
	Expression            *string   `bun:"expression,type:varchar(8000)"`                                // The expression of a db computed column or a computed field.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                                     // Status flag indicating whether the record is active or not.
	Description           *string   `bun:"description,type:varchar(8000)"`                               // User friendly description for this column that can be used as a default label.
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
