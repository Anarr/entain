package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func Migrate(gormDB *gorm.DB, dbName string, path string) error {
	db, err := gormDB.DB()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		SchemaName: dbName,
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+path, dbName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
