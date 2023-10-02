package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/materials-resources/s_prophet/core/port/repository"
	"gorm.io/gorm"
)

type database struct {
	*sql.DB
}

func NewDB() (repository.Database, error) {
	db, err := sql.Open(
		"sqlserver",
		os.Getenv("DSN"),
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

func gormToRepositoryError(err error) error {
	if errors.Is(
		err,
		gorm.ErrRecordNotFound,
	) {
		return repository.NotFound
	}
	return err
}
