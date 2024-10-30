package main

// This is the inside of the hexagon, where the business logic is implemented.
// This code is decoupled from any Primary(Driving) and Secondary(Driven) Adapters
// specific implementations.

import (
	"context"
	"time"
)

// Task represents a task in the system
type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
}

// TaskStatus represents the status of a task
type TaskStatus string

const (
	DONE   TaskStatus = "DONE"
	UNDONE TaskStatus = "UNDONE"
)

// TaskToCreate represents the data needed to create a task
type TaskToCreate struct {
	Title       string
	Description string
}

// CreateTaskFunc represents the business process of creating a task
type CreateTaskFunc func(ctx context.Context, task TaskToCreate) (Task, error)

// SaveTaskFunc represents the business process of saving a task
type SaveTaskFunc func(ctx context.Context, task Task) error

// NotifyAboutTaskChangeFunc represents the business process of generate a notification
type NotifyAboutTaskChangeFunc func(ctx context.Context, id int) error

// NewCreateTaskFunc creates a new CreateTaskFunc that saves a task and notifies about it
// every adapter will call this function to create a new CreateTaskFunc
// independent of the concrete technology used to save the task and notify about it
func NewCreateTaskFunc(save SaveTaskFunc, notify NotifyAboutTaskChangeFunc) CreateTaskFunc {
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
