// HTTP driving adapter for task creation
// It represents a primary adapter on the left side of the hexagon
package main

import "github.com/labstack/echo/v4"

// NewTaskHTTPHandler creates a new HTTP handler for creating tasks
// Is a concrete implementation of the trigger to create a task
// Is a driving adapter because is responsible for initiating the business logic process
// in the hexagon it positions itself as a primary adapter on the left side of the hexagon
func NewTaskHTTPHandler(task CreateTaskFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskToCreate := TaskToCreate{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
		}

		if err := c.Bind(&taskToCreate); err != nil {
			return c.String(400, "Invalid request")
		}

		createdTask, err := task(c.Request().Context(), taskToCreate)
		if err != nil {
			return c.String(500, "Failed to create task")
		}

		return c.JSON(201, createdTask)
	}
}
