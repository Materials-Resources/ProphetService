package model

type XmlDataobjectColumn struct {
	bun.BaseModel          `bun:"table:xml_dataobject_column"`
	XmlDataobjectColumnUid int32     `bun:"xml_dataobject_column_uid,type:int,pk"`
	XmlDataobjectUid       int32     `bun:"xml_dataobject_uid,type:int"`
	ColumnName             string    `bun:"column_name,type:varchar(40)"`
	ColumnLabel            string    `bun:"column_label,type:varchar(255)"`
	ColumnId               int32     `bun:"column_id,type:int"`
	ColumnType             string    `bun:"column_type,type:varchar(15)"`
	ColumnDbName           string    `bun:"column_db_name,type:varchar(255)"`
	EditRequired           string    `bun:"edit_required,type:varchar(3)"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`
	CustomFlag             string    `bun:"custom_flag,type:char,default:('N')"`
}
