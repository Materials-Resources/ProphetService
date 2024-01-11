package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/materials-resources/s_prophet/config"
	"github.com/materials-resources/s_prophet/core/port/repository"
	_ "github.com/microsoft/go-mssqldb"
)

type database struct {
	*sql.DB
}

func NewDB(c *config.Config) (repository.Database, error) {

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DB,
	)
	db, err := sql.Open(
		"sqlserver",
		dsn,
	)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	return &database{
		db,
	}, nil
}

func (d database) Close() error {
	return d.DB.Close()
}

func (d database) GetDB() *sql.DB {
	return d.DB
}
