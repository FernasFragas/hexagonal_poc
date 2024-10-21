// HTTP driving adapter for task creation
package main

import "github.com/labstack/echo/v4"

func NewTaskHTTPHandler(task CreateTaskFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var taskToCreate TaskToCreate

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
