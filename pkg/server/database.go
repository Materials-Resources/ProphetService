package server

import (
	"database/sql"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func connectBun() *bun.DB {
	sqldb, err := sql.Open("sqlserver", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mssqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}

func connectSQL() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(os.Getenv("DSN")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	return db
}
