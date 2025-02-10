# Task Management Service

## Описание
Сервис на Go для управления задачами. Он реализует REST API для создания задач, сохраняет их в базе данных PostgreSQL и экспортирует метрики в Prometheus, чтобы их можно было визуализировать в Grafana.

## Предварительные требования
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Быстрый старт

### 1. Клонирование репозитория

Клонируйте репозиторий на локальную машину:

```bash
git clone https://github.com/natindo/GolangTestTasks.git
cd task-management-service
```

### 2. Переменные окружения
Для настройки подключения к базе данных и параметров сервиса используются переменные окружения. В репозитории хранится файл `.env.example`, содержащий пример значений. После клонирования репозитория скопируйте его в `.env` и при необходимости измените значения.

```bash
cp .env.example .env
```

Пример `.env.example`:
```bash
# Настройки для PostgreSQL
DB_HOST=tasks-db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=example
DB_NAME=tasks_db

# Порт приложения
APP_PORT=8080
```

### 3. Запуск контейнеров
Соберите и запустите контейнеры с помощью Docker Compose:
```bash
docker-compose up --build
```

### 4. Проверка работы приложения
#### 4.1 Health Check
Откройте браузер и перейдите по адресу: http://localhost:8080/health Вы должны увидеть ответ
```bash
OK
```
#### 4.2 Отправка запроса
Отправьте запрос для создания новой задачи:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"title": "Test Task", "description": "This is a test task"}' \
  http://localhost:8080/tasks
```
Пример ожидаемого ответа:
```json
{
  "id": 1,
  "title": "Test Task",
  "description": "This is a test task",
  "created_at": "2025-01-01T10:00:00Z"
}
```
#### 4.3 Метрики
По адресу http://localhost:8080/metrics находятся метрики.
```
tasks_created_total
task_creation_duration_seconds
```
#### 4.4 Prometheus
Prometheus доступен по адресу: http://localhost:9090
#### 4.5 Grafana
Grafana доступна по адресу: http://localhost:3000

### 5. Запуск тестов
Для запуска Unit-тестов запустите команду в корне проекта:
```bash
go test ./...
```