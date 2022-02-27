package main

import (
	"encoding/json"
	"fmt"
	"github.com/brandomota/golang-migrations-example/middlewares"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	models "github.com/brandomota/golang-migrations-example/models"
	services "github.com/brandomota/golang-migrations-example/services"
)

var logger = log.Default()

func handlerUsers(w http.ResponseWriter, req *http.Request) {
	var msg = ""
	switch req.Method {
	case "POST":
		var user models.User
		body, _ := ioutil.ReadAll(req.Body)
		errorParse := json.Unmarshal(body, &user)

		if errorParse != nil {
			msg = fmt.Sprintf("error on parse body: %s", errorParse)
			logger.Print(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		userCreated, errorCreateUser := services.AddUser(user)

		if errorCreateUser != nil {
			msg = fmt.Sprintf("error on user create: %s", errorCreateUser)
			logger.Print(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userCreated)
	case "DELETE":
		fmt.Println("2")
	case "GET":
		var stringId = req.URL.Path[:(strings.LastIndex(req.URL.Path, "/") + 1)]

		id, err := strconv.ParseInt(stringId, 0, 32)
		if err != nil {
			msg = fmt.Sprintf(`Invalid id %d`, id)
			http.Error(w, msg, http.StatusBadRequest)
		}
		//user, err := services.GetUser(id)
	default:
		msg = "Method not allowed"
		logger.Print(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Print("Error loading .env file")
	}
	logger.Print("Running migrations...")
	err = services.RunMigrations()
	if err != nil {
		logger.Fatalln(fmt.Sprintf("error on running migrations : %s", err))
	}
	server := http.NewServeMux()
	server.HandleFunc("/users/", handlerUsers)
	middleware := middlewares.NewLogger(server)
	logger.Print("server running...")
	http.ListenAndServe(":8000", middleware)

}
