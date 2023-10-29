package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomazcx/go-todo-list/application"
	infra "github.com/tomazcx/go-todo-list/infra/db"
)

func main() {
	fmt.Println("Server is starting...")

	//Load the env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbhost := os.Getenv("DB_HOST")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbuser := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbpassword, dbuser, dbname)

	err := infra.ConnectToDB(connStr)

	if err == nil {
		fmt.Println("Connected to database!")
	}

	err = infra.InitTables()

	if err == nil {
		fmt.Println("Created tables!")
	}

	port := os.Getenv("PORT")

	application.Router()

	if port == "" {
		port = "3000"
	}

	fmt.Println("Server is now running at port " + port + " 🚀")
	http.ListenAndServe(":"+port, nil)
}
