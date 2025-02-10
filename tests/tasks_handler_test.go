package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/natindo/GolangTestTasks/internal/handlers"
	"github.com/natindo/GolangTestTasks/internal/models"
	"github.com/natindo/GolangTestTasks/internal/repository"
)

type MockRepository struct {
	CreateTaskFunc func(ctx context.Context, task *models.Task) error
}

var _ repository.TaskRepository = (*MockRepository)(nil)

func (m *MockRepository) CreateTask(ctx context.Context, task *models.Task) error {
	return m.CreateTaskFunc(ctx, task)
}

func TestCreateTask(t *testing.T) {
	e := echo.New()

	mockRepo := &MockRepository{
		CreateTaskFunc: func(ctx context.Context, task *models.Task) error {
			task.ID = 1
			return nil
		},
	}

	payload := map[string]string{
		"title":       "Test task",
		"description": "Mock description",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	handler := handlers.CreateTaskHandler(mockRepo)
	err := handler(c)
	assert.NoError(t, err, "Query processing error")

	assert.Equal(t, http.StatusCreated, rec.Code, "The status code should be 201")

	var response models.Task
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "There shouldn't be an error when parsing the response")

	assert.Equal(t, int64(1), response.ID)
	assert.Equal(t, "Test task", response.Title)
	assert.Equal(t, "Mock description", response.Description)
}
