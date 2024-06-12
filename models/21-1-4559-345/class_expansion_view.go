package model

type ClassExpansionView struct {
	bun.BaseModel         `bun:"table:class_expansion_view"`
	ClassExpansionViewUid int32  `bun:"class_expansion_view_uid,type:int,identity"`
	SchemaName            string `bun:"schema_name,type:varchar(255)"`
	TableName             string `bun:"table_name,type:varchar(255)"`
	TableType             string `bun:"table_type,type:varchar(255)"`
	Done                  string `bun:"done,type:char"`
}
