package gen

import "github.com/uptrace/bun"

type P21DatabaseChanges struct {
	bun.BaseModel         `bun:"table:p21_database_changes"`
	P21DatabaseChangesUid int32  `bun:"p21_database_changes_uid,type:int,autoincrement,scanonly"`
	Change                string `bun:"change,type:varchar(50),nullzero"`       // This denotes is an object was added or removed.
	ObjectType            string `bun:"object_type,type:varchar(255),nullzero"` // This denotes the type of object that changed. Table, proc, view or column.
	ObjectName            string `bun:"object_name,type:varchar(255),nullzero"` // This is the object name that was changed.
	TableName             string `bun:"table_name,type:varchar(255),nullzero"`  // This is the table for the column that was changed. Will only be populated for column adds and removals.
	Version               string `bun:"version,type:varchar(255),nullzero"`     // This is the update version that was applied WHEN the object changed.
}
