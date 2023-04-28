package db

import (
	"database/sql"
	"log"
	"os"
	"time"
)

func Connect() *sql.DB {
	dsn := os.Getenv("DSN")
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, _ := openDB(dsn)
		if db != nil {
			log.Println("Connected to Postgres")
			return db
		}

		log.Println("Postgres not yet ready, sleep 2 seconds...")
		sleepTwoSeconds()
	}

	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sleepTwoSeconds() {
	time.Sleep(2 * time.Second)
}
