package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/natindo/GolangTestTasks/internal/metrics"
	"github.com/natindo/GolangTestTasks/internal/models"
	"github.com/natindo/GolangTestTasks/internal/repository"
)

func CreateTaskHandler(repo repository.TaskRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Incorrect input data"})
		}

		task := &models.Task{
			Title:       input.Title,
			Description: input.Description,
		}

		if err := repo.CreateTask(c.Request().Context(), task); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving tasks"})
		}

		metrics.TasksCreatedTotal.Inc()

		duration := time.Since(startTime).Seconds()
		metrics.TaskCreationDuration.Observe(duration)

		return c.JSON(http.StatusCreated, task)
	}
}
