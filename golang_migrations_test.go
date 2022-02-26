package main

import (
	"database/sql"
	"github.com/brandomota/golang-migrations-example/models"
	"github.com/brandomota/golang-migrations-example/services"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var conn *sql.DB

func TestAddUserInDatabase(t *testing.T) {
	var user models.User
	user.Name = "nome teste"
	user.Age = 1

	userAdded, err := services.AddUser(user)
	assert.Nil(t, err, "Erro deveria ser nulo")
	assert.Equal(t, userAdded.Id, 1, "id  veio incorreto")
	assert.Equal(t, userAdded.Name, user.Name, "id  veio incorreto")
	assert.Equal(t, userAdded.Age, user.Age, "id  veio incorreto")

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
	godotenv.Load()
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
