package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Return database instance
func GetDBInstance() (*sql.DB, error) {
	return getDB()
}

// Execute SQL Statement in database
func ExecCreateEntity(query string) (int, error) {
	var err error
	var id = 0
	db, err := getDB()
	if err != nil {
		return id, err
	}
	err = db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	return id, err
}

// Get Database object
func getDB() (*sql.DB, error) {

	db, err := sql.Open(os.Getenv("DRIVER_NAME"), os.Getenv("DATASOURCE_URL"))
	return db, err
}
