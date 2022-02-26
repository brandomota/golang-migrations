package main

import (
	"encoding/json"
	"fmt"
	"github.com/brandomota/golang-migrations-example/middlewares"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"

	models "github.com/brandomota/golang-migrations-example/models"
	services "github.com/brandomota/golang-migrations-example/services"
)

var logger = log.Default()

func handlerCreateUser(w http.ResponseWriter, req *http.Request) {
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
	default:
		msg = "Method not allowed"
		logger.Print(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Print("Error loading .env file")
	}
	server := http.NewServeMux()
	server.HandleFunc("/", handlerCreateUser)
	middleware := middlewares.NewLogger(server)
	logger.Print("server running...")
	http.ListenAndServe(":8000", middleware)

}
