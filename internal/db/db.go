package db

import (
	"database/sql"
	"os"

	envsettings "github.com/kasariks/project_for_graduating/internal/env_settings"
	_ "modernc.org/sqlite"
)

const schema = "CREATE TABLE scheduler (" +
	"id INTEGER PRIMARY KEY AUTOINCREMENT," +
	"date CHAR(8) NOT NULL DEFAULT \"\"," +
	"title VARCHAR(255) NOT NULL," +
	"comment TEXT NOT NULL DEFAULT \"\"," +
	"repeat VARCHAR(255) NOT NULL DEFAULT \"\"" +
	");"

var db *sql.DB

func Init(dbFiles string) error {
	dbFile := envsettings.Env.DbPath
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	if install {
		if _, err := db.Exec(schema); err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return db.Close()
}
