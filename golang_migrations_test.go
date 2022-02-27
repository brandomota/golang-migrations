package main

import (
	"database/sql"
	"github.com/brandomota/golang-migrations-example/models"
	"github.com/brandomota/golang-migrations-example/services"
	"github.com/go-testfixtures/testfixtures/v3"
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
	assert.Equal(t, userAdded.Name, user.Name, "nome  veio incorreto")
	assert.Equal(t, userAdded.Age, user.Age, "idade  veio incorreta")

}

func TestGetUserById(t *testing.T) {
	id := 100
	user, err := services.GetUserById(id)
	assert.Nil(t, err, "Erro deveria ser nulo")
	assert.Equal(t, user.Name, "usuario existente", "nome veio incorreto")
	assert.Equal(t, user.Age, 50, "idade veio incorreta")
}

func initDB(m *testing.M) error {
	var err error
	conn, _ = sql.Open("sqlite3", os.Getenv("DATASOURCE_URL"))

	driver, _ := sqlite3.WithInstance(conn, &sqlite3.Config{})

	migrations_instance, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3", driver)

	err = migrations_instance.Up()
	if err != nil {
		return err
	}

	err = addFixtureData()

	return err

}

func addFixtureData() error {
	var err error
	fixtures, err := testfixtures.New(
		testfixtures.Database(conn),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory("./fixtures"),
	)

	if err != nil {
		return err
	}
	err = fixtures.Load()
	return err
}

func downDB(m *testing.M) {
	conn.Close()
	err := os.Remove(os.Getenv("DATASOURCE_URL"))
	if err != nil {
		log.Default().Print(err)
	}
}

func TestMain(m *testing.M) {
	godotenv.Load("test.env")
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
