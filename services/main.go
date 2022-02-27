package services

import (
	"github.com/brandomota/golang-migrations-example/repositories"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var logger = log.Default()

func RunMigrations() error {
	var err error
	db, err := repositories.GetDBInstance()
	if err != nil {
		return err
	}
	driver, _ := sqlite3.WithInstance(db, &sqlite3.Config{})

	migrations_instance, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3", driver)

	if err != nil {
		return err
	}

	err = migrations_instance.Up()

	if err != nil {
		if err.Error() != "no change" {
			return err
		} else {
			logger.Print("No new migrations to execute, skiping....")
			return nil
		}
	}

	return nil
}
