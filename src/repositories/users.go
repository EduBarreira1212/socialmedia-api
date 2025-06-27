package repositories

import (
	"database/sql"
	"socialmedia-api/src/models"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) Create(user models.User) (uint64, error) {
	return 0, nil
}
