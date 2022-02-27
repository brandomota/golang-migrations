package services

import (
	"fmt"

	"github.com/brandomota/golang-migrations-example/models"
	"github.com/brandomota/golang-migrations-example/repositories"
)

func AddUser(user models.User) (models.User, error) {
	query := fmt.Sprintf(`INSERT INTO "USERS" (NAME, AGE) VALUES ("%s", %d) RETURNING ID`, user.Name, user.Age)
	var userAdded models.User
	id, err := repositories.ExecCreateEntity(query)
	if err != nil {
		return userAdded, err
	}
	err2 := repositories.GetUserByID(id, &userAdded)
	if err2 != nil {
		return userAdded, err2
	}
	return userAdded, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := repositories.GetUserByID(id, &user)
	return user, err
}
