package repositories

import (
	"database/sql"
	"fmt"
	"socialmedia-api/src/models"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (repository User) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nickName, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdInserted), nil
}

func (repository User) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repository.db.Query(
		"SELECT id, name, nickName, email, createdAt FROM users WHERE name LIKE ? OR nickName LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository User) SearchById(ID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nickName, email, createdAt FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository User) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, nickName = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.NickName, user.Email, ID); err != nil {
		return err
	}

	return nil
}
