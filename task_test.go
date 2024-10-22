package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewTaskHTTPHandler(t *testing.T) {
	// create a test for NewTaskHTTPHandler
	t.Run("Test NewTaskHTTPHandler", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Test Task","description":"Test Description"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, err := InitDB()
		assert.NoError(t, err)

		taskRepository := NewTaskRepository(db)

		if assert.NoError(t, NewTaskHTTPHandler(NewCreateTaskFunc(
			taskRepository.CreateTask(),
			func(ctx context.Context, id int) error {
				fmt.Printf("Task with ID %d was saved", id)

				return nil
			},
		),
		)(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

}
