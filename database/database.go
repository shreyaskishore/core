package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/acm-uiuc/core/config"
)

var db *sqlx.DB

func New() (*sqlx.DB, error) {
	var err error
	if db == nil {
		username, err := config.GetConfigValue("DB_USER")
		if err != nil {
			return nil, fmt.Errorf("failed to get config value: %w", err)
		}

		password, err := config.GetConfigValue("DB_PASS")
		if err != nil {
			return nil, fmt.Errorf("failed to get config value: %w", err)
		}

		hostname, err := config.GetConfigValue("DB_HOST")
		if err != nil {
			return nil, fmt.Errorf("failed to get config value: %w", err)
		}

		database, err := config.GetConfigValue("DB_NAME")
		if err != nil {
			return nil, fmt.Errorf("failed to get config value: %w", err)
		}

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
