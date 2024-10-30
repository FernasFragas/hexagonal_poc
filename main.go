package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler, err := setupAdapters()
	if err != nil {
		e.Logger.Error(err.Error())
	}

	e.POST("/add", handler)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func setupAdapters() (echo.HandlerFunc, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	taskRepository := NewTaskRepository(db)

	createTaskFunc := NewTaskHTTPHandler(
		NewCreateTaskFunc(
			taskRepository.CreateTask,
			PrintTaskIDToConsole,
		),
	)

	return createTaskFunc, nil
}

// PrintTaskIDToConsole is a concrete implementation of NotifyAboutTaskChangeFunc
// that prints a message to the console when the task is saved
// it represents a secondary adapter on the right side of the hexagon since is triggered by the business logic
func PrintTaskIDToConsole(_ context.Context, id int) error {
	fmt.Printf("Task with ID %d was saved", id)

	return nil
}
