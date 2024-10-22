// Represents the driven adapter of the task repository.
package main

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		panic(err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS tasks (
     		id INTEGER PRIMARY KEY,
     		title TEXT,
     		description TEXT,
     		status TEXT
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) CreateTask() SaveTaskFunc {
	return func(ctx context.Context, task Task) error {
		_, err := t.db.Exec("INSERT INTO tasks (id, title, description, status) VALUES (?, ?, ?, ?)", task.ID, task.Title, task.Description, task.Status)
		if err != nil {
			return err
		}

		return nil
	}
}
