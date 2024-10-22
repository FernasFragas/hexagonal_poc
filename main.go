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
		NewCreateTaskFunc(taskRepository.CreateTask(),
			func(ctx context.Context, id int) error {
				fmt.Printf("Task with ID %d was saved", id)

				return nil
			},
		),
	)

	return createTaskFunc, nil
}
