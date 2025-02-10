package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/natindo/GolangTestTasks/internal/handlers"
	"github.com/natindo/GolangTestTasks/internal/metrics"
	"github.com/natindo/GolangTestTasks/internal/repository"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPortStr := getEnv("DB_PORT", "5432")
	dbPort, _ := strconv.Atoi(dbPortStr)
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "example")
	dbName := getEnv("DB_NAME", "tasks_db")
	appPort := getEnv("APP_PORT", "8080")

	repo, err := repository.NewPostgresRepository(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer repo.Close()

	metrics.InitMetrics()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.POST("/tasks", handlers.CreateTaskHandler(repo))

	log.Printf("The service is listening on the port %s", appPort)
	if err := e.Start(":" + appPort); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultVal
}
