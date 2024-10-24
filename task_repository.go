// Represents the driven adapter of the task repository.
// it represents a secondary adapter on the right side of the hexagon
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

// CreateTask is a function that saves a task in the database
// it is a concrete implementation of saving a task of the application
// And is a driven adapter since is triggered and used by the business logic
func (t *TaskRepository) CreateTask() SaveTaskFunc {
	return func(ctx context.Context, task Task) error {
		_, err := t.db.Exec("INSERT INTO tasks (id, title, description, status) VALUES (?, ?, ?, ?)", task.ID, task.Title, task.Description, task.Status)
		if err != nil {
			return err
		}

		return nil
	}
}
