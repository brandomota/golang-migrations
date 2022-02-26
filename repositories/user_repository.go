package repositories

import (
	"fmt"
	"github.com/brandomota/golang-migrations-example/models"
	"github.com/kisielk/sqlstruct"
)

func GetUserByID(id int, model *models.User) error {
	query := fmt.Sprintf(`SELECT * FROM "USERS" WHERE ID = %d`, id)
	var err error
	db, err := getDB()

	row, err := db.Query(query)
	for row.Next() {
		err = sqlstruct.Scan(model, row)
	}
	defer db.Close()

	return err
}
