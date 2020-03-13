package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	username = "devuser"          // TODO: Load this from env
	password = "devpass"          // TODO: Load this from env
	hostname = "(localhost:3306)" // TODO: Load this from env
	database = "core"
)

var db *sqlx.DB

func New() (*sqlx.DB, error) {
	var err error
	if db == nil {
		db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", username, password, hostname, database))
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %w", err)
		}
	}

	err = db.Ping()
	if err != nil {
		db = nil
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
