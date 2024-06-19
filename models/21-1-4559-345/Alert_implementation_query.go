package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertImplementationQuery struct {
	bun.BaseModel               `bun:"table:Alert_implementation_query"`
	AlertImplementationQueryUid int32     `bun:"alert_implementation_query_uid,type:int,pk"` // Unique Identifier for alert_implementation_query table.
	AlertImplementationUid      int32     `bun:"alert_implementation_uid,type:int"`          // Unique Identifier from alert_implementation table.
	ColumnId                    int32     `bun:"column_id,type:int"`                         // What is the unique identifier for this column?
	OperatorCd                  int32     `bun:"operator_cd,type:int"`                       // Code identifying what operator to use to build the query or where such as =, <>, >.
	ColumnValue                 string    `bun:"column_value,type:varchar(255)"`             // The value the user wants specified for a column when building the query or where clause.
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                 // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`           // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`        // ID of the user who last maintained this record
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`                   // Indicates current record status.
	ColumnValueDescription      string    `bun:"column_value_description,type:varchar(255)"` // Description of the column values.
}
