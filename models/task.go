package models

import (
	"fmt"
	"log"

	infra "github.com/tomazcx/go-todo-list/infra/db"
)

type Task struct {
	Id        uint
	Name      string
	Completed bool
}

func Index() ([]Task, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return nil, err
	}

	query := "SELECT * FROM task"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Error running the SQL query.")
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var item Task
		if err := rows.Scan(&item.Id, &item.Name, &item.Completed); err != nil {
			log.Fatal("Error running the SQL query.")
			return nil, err
		}

		tasks = append(tasks, item)
	}

	return tasks, nil

}

func CreateTask(name string) (Task, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return Task{}, err
	}
	query := "INSERT INTO task (name, completed) VALUES ($1, false) RETURNING id, name, completed"

	var task Task

	err = db.QueryRow(query, name).Scan(&task.Id, &task.Name, &task.Completed)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error inserting the item into the database: %v", err))
		return Task{}, err

	}

	return task, nil
}

func DeleteTask(id uint) error {

	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return err
	}

	query := "DELETE FROM task WHERE id=$1"
	db.Exec(query, id)

	return nil
}
