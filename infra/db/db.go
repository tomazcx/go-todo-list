package infra

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Failed to ping to the database: %v", err)
		return err
	}

	query := "CREATE TABLE IF NOT EXISTS task (id SERIAL PRIMARY KEY, name VARCHAR(255), completed BOOLEAN)"
	_, err = db.Exec(query)

	if err != nil {
		log.Fatalf("Failed to create todo table: %v", err)
		return err
	}

	return nil
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		log.Fatalf("You need to connect to the database before trying to access its instance")
		return nil, errors.New("Connection not made")
	}

	return db, nil
}

func TodoTable() {
}
