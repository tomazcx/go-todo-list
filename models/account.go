package models

import (
	"fmt"
	"log"

	infra "github.com/tomazcx/go-todo-list/infra/db"
)

type Account struct {
	Login    string
	Password string
}

func (a *Account) FindByLogin(login string) (Account, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the db instance.")
		return Account{}, err
	}

	var result Account
	query := "SELECT login, password FROM account WHERE login=$1"
	err = db.QueryRow(query, login).Scan(&result.Login, &result.Password)

	if err != nil {
		return Account{}, err
	}

	return result, nil

}

func (a *Account) Create(login, password string) (Account, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the db instance.")
		return Account{}, err
	}

	var result Account
	query := "INSERT INTO account (login, password) VALUES ($1, $2) RETURNING login, password"
	err = db.QueryRow(query, login, password).Scan(&result.Login, &result.Password)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating account: %v", err))
		return Account{}, err
	}

	return result, nil
}
