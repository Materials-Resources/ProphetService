package gen

import "github.com/uptrace/bun"

type TempDefaultRedo struct {
	bun.BaseModel  `bun:"table:temp_default_redo"`
	DefaultId      int32  `bun:"default_id,type:int,autoincrement"`
	ColumnName     string `bun:"column_name,type:varchar(255)"`
	TableName      string `bun:"table_name,type:varchar(255)"`
	CurrentDefault string `bun:"current_default,type:varchar(512)"`
	NewDefault     string `bun:"new_default,type:varchar(512)"`
}
