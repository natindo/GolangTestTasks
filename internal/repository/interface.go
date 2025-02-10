package repository

import (
	"context"

	"github.com/natindo/GolangTestTasks/internal/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
}
