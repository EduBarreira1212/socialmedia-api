package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"socialmedia-api/src/database"
	"socialmedia-api/src/models"
	"socialmedia-api/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	parameters, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(parameters, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted: %d", userId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users!"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
