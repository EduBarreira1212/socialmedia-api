package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"socialmedia-api/src/auth"
	"socialmedia-api/src/database"
	"socialmedia-api/src/models"
	"socialmedia-api/src/repositories"
	"socialmedia-api/src/responses"
	"socialmedia-api/src/security"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	parameters, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(parameters, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, err := repository.SearchById(ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userIDInToken != ID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible to update a different user"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.Update(ID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userIDInToken != ID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible to delete a different user"))
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.Delete(ID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

func ToFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible to follow yourself"))
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.ToFollow(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

func StopFollowingUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusBadRequest, errors.New("isn't possible to stop following yourself"))
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.StopFollowing(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	followers, err := repository.GetFollowers(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	following, err := repository.GetFollowing(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDInToken, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if userIDInToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible to update other user's password"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(body, &password); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	passwordInDB, err := repository.GetPassword(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordInDB, password.Current); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	passwordWithHash, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(passwordWithHash)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
