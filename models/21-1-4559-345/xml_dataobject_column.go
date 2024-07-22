package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDataobjectColumn struct {
	bun.BaseModel          `bun:"table:xml_dataobject_column"`
	XmlDataobjectColumnUid int32     `bun:"xml_dataobject_column_uid,type:int,pk"`  // Unique indentifier for rows within the table
	XmlDataobjectUid       int32     `bun:"xml_dataobject_uid,type:int"`            // Foreign key to unique identifier to parent records in xml_dataobject table
	ColumnName             string    `bun:"column_name,type:varchar(40)"`           // Column name of associated retrieved data
	ColumnLabel            string    `bun:"column_label,type:varchar(255)"`         // Display name of retrieved column
	ColumnId               int32     `bun:"column_id,type:int"`                     // Number that identifies sequence within select statement
	ColumnType             string    `bun:"column_type,type:varchar(15)"`           // Data type of retrieved column data
	ColumnDbName           string    `bun:"column_db_name,type:varchar(255)"`       // Name of column on database table rather than name used within retrieved result set
	EditRequired           string    `bun:"edit_required,type:varchar(3)"`          // Indicator of whether the data of the associated retrieved column is required
	DateCreated            time.Time `bun:"date_created,type:datetime"`             // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`       // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`    // ID of the user who last maintained this record
	CustomFlag             string    `bun:"custom_flag,type:char(1),default:('N')"` // Indicator that column is custom, not baseline
}
