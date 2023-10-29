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

func (t *Task) Index() ([]Task, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return nil, err
	}

	query := "SELECT id, name, completed FROM task ORDER BY createdat"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error running the SQL query: %v", err))
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var item Task
		if err := rows.Scan(&item.Id, &item.Name, &item.Completed); err != nil {
			log.Fatal(fmt.Sprintf("Error scanning the row: %v", err))
			return nil, err
		}

		tasks = append(tasks, item)
	}

	return tasks, nil

}

func (t *Task) Create(name string) (Task, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return Task{}, err
	}
	query := "INSERT INTO task (name, completed, createdAt) VALUES ($1, false, NOW()) RETURNING id, name, completed"

	var task Task

	err = db.QueryRow(query, name).Scan(&task.Id, &task.Name, &task.Completed)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error inserting the item into the database: %v", err))
		return Task{}, err

	}

	return task, nil
}

func (t *Task) ToggleCompleted(taskId uint) (Task, error) {
	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return Task{}, err
	}

	var result Task
	query := "UPDATE task SET completed = NOT completed WHERE id=$1 RETURNING id, name, completed"

	err = db.QueryRow(query, taskId).Scan(&result.Id, &result.Name, &result.Completed)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error updating the task: %v", err))
		return Task{}, err
	}

	return result, nil

}

func (t *Task) Delete(id uint) error {

	db, err := infra.GetDB()

	if err != nil {
		log.Fatal("Error getting the DB instance.")
		return err
	}

	query := "DELETE FROM task WHERE id=$1"
	db.Exec(query, id)

	return nil
}
