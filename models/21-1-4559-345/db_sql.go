package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DbSql struct {
	bun.BaseModel   `bun:"table:db_sql"`
	LastSqlExecuted string    `bun:"last_sql_executed,type:varchar(50)"` // What is the name of the script that was executed?
	DateSqlExecuted time.Time `bun:"date_sql_executed,type:datetime"`    // When was this SQL script executed?
	SqlDescription  string    `bun:"sql_description,type:varchar(255)"`  // What did this script do?
	DbSqlUid        int32     `bun:"db_sql_uid,type:int,autoincrement,scanonly,pk"`
}
