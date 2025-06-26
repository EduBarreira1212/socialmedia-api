package database

import (
	"database/sql"
	"socialmedia-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Conect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DCS)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
