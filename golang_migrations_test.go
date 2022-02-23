package golangmigrations

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
)

var conn *sql.DB

func TestMigrations(t *testing.T) {
	assert.Equal(t, true, true, "valor incorreto")
}

func initDB(m *testing.M) error {
	conn, _ = sql.Open("sqlite3", "./database.db")

	driver, _ := sqlite3.WithInstance(conn, &sqlite3.Config{})

	migrations_instance, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3", driver)

	err := migrations_instance.Up()

	return err

}

func downDB(m *testing.M) {
	conn.Close()
	err := os.Remove("database.db")
	if err != nil {
		log.Default().Print(err)
	}
}

func TestMain(m *testing.M) {
	err := initDB(m)
	if err != nil {
		downDB(m)
		log.Fatalln(err)

	} else {
		code := m.Run()
		downDB(m)
		os.Exit(code)
	}

}
