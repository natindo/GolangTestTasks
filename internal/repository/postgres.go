package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/natindo/GolangTestTasks/internal/models"
)

type PostgresRepository struct {
	conn *pgx.Conn
}

func NewPostgresRepository(host string, port int, user, password, dbName string) (*PostgresRepository, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, password, host, port, dbName)

	conn, err := ConnectPostgres(connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Postgres: %w", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);
	`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("ошибка при создании таблицы tasks: %w", err)
	}

	return &PostgresRepository{conn: conn}, nil
}

func (r *PostgresRepository) Close() error {
	return r.conn.Close(context.Background())
}

func (r *PostgresRepository) CreateTask(ctx context.Context, task *models.Task) error {
	task.CreatedAt = time.Now().UTC()

	query := `
		INSERT INTO tasks (title, description, created_at)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	if err := r.conn.QueryRow(ctx, query, task.Title, task.Description, task.CreatedAt).Scan(&task.ID); err != nil {
		return fmt.Errorf("ошибка при вставке задачи: %w", err)
	}

	return nil
}

func ConnectPostgres(connStr string) (*pgx.Conn, error) {
	cfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse config error: %w", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("pgx connect error: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("pgx ping error: %w", err)
	}

	return conn, nil
}
