// This is the inside of the hexagon, where the business logic is implemented.
// This code is decoupled from any Primary(Driving) and Secondary(Driven) Adapters
// specific implementations.
package main

import (
	"context"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
}

type TaskStatus string

const (
	DONE   TaskStatus = "DONE"
	UNDONE TaskStatus = "UNDONE"
)

type TaskToCreate struct {
	Title       string
	Description string
}

type CreateTaskFunc func(ctx context.Context, task TaskToCreate) (Task, error)

type SaveTaskFunc func(ctx context.Context, task Task) error

type NotifyAboutTaskSaveOrUpdatedFunc func(ctx context.Context, id int) error

func NewCreateTaskFunc(save SaveTaskFunc, notify NotifyAboutTaskSaveOrUpdatedFunc) CreateTaskFunc {
	return func(ctx context.Context, taskToSave TaskToCreate) (Task, error) {
		task := Task{
			ID:          time.Now().Nanosecond(),
			Title:       taskToSave.Title,
			Description: taskToSave.Description,
			Status:      UNDONE,
		}

		err := save(ctx, task)
		if err != nil {
			return Task{}, err
		}

		err = notify(ctx, task.ID)
		if err != nil {
			return Task{}, err
		}

		return task, nil
	}
}
