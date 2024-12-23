package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func New() *sql.DB {
	dbFile := "scheduler.db"

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Fatalf("Файл базы данных %s не найден", dbFile)
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("init db", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("ping db", err)
	}

	return db
}
