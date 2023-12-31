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

	return nil
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		log.Fatalf("You need to connect to the database before trying to access its instance")
		return nil, errors.New("Connection not made")
	}

	return db, nil
}

func InitTables() error {
	query := "CREATE TABLE IF NOT EXISTS account (id SERIAL PRIMARY KEY, login VARCHAR(255), password VARCHAR(255))"
	_, err := db.Exec(query)

	if err != nil {
		log.Fatalf("Failed to create user table: %v", err)
		return err
	}

	query = "CREATE TABLE IF NOT EXISTS task (id SERIAL PRIMARY KEY, name VARCHAR(255), completed BOOLEAN, createdAt TIMESTAMP, user_id INTEGER, CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES account(id))"
	_, err = db.Exec(query)

	if err != nil {
		log.Fatalf("Failed to create task table: %v", err)
		return err
	}

	return nil

}
