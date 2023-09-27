package repository

import (
	"errors"
	"log"
	"os"

	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type database struct {
	*gorm.DB
}

func NewDB() (repository.Database, error) {
	db, err := gorm.Open(
		sqlserver.Open(os.Getenv("DSN")),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		},
	)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	return &database{
		db,
	}, nil
}

func (da database) Close() error {
	dbInstance, _ := da.DB.DB()

	return dbInstance.Close()
}

func (da database) GetDB() *gorm.DB {
	return da.DB
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
