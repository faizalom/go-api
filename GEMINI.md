# Go REST API: Workout Tracker

This project is a RESTful API built in Go, developed interactively with Gemini. It serves as a robust foundation for a workout tracking application, featuring a clean architecture, JWT-based authentication, and user management.

## Key Features

*   **Clean Architecture**: Follows a layered architecture (Handler, Service, Repository) with Dependency Injection.
*   **JWT Authentication**: Secure endpoints using JWT, with token generation (`/login`) and middleware validation.
*   **User CRUD**: Endpoints for creating and retrieving users.
*   **Configuration Management**: All settings (server port, JWT secret, database connection) are managed via a `config.yaml` file.
*   **Structured Logging**: Centralized logger with different levels (`Info`, `Error`).
*   **Database Migrations**: Schema changes are managed through SQL migration files in the `/migrations` folder.

## Project Structure

```
/workout-api/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── auth_handler.go
│   │   ├── example_handler.go
│   │   ├── profile_handler.go
│   │   └── user_handler.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── logging.go
│   │   └── middleware.go
│   ├── model/
│   │   ├── claims.go
│   │   └── user.go
│   ├── repository/
│   │   ├── repo_a.go
│   │   ├── repo_b.go
│   │   └── user_repository.go
│   ├── router/
│   │   └── router.go
│   └── service/
│       ├── service_a.go
│       ├── service_b.go
│       └── user_service.go
├── migrations/
│   ├── 000001_create_users_table.down.sql
│   └── 000001_create_users_table.up.sql
├── pkg/
│   └── logger/
│       └── logger.go
├── .gitignore
├── go.mod
├── go.sum
└── GEMINI.md
```

## API Endpoints

### Public Routes

*   **`POST /login`**: Simulates a user login and returns a JWT.
*   **`POST /users`**: Creates a new user.
*   **`GET /users/{id}`**: Retrieves a user by their ID.

### Protected Routes (Requires `Authorization: Bearer <token>`)

*   **`GET /api/v1/profile`**: Returns the profile information for the authenticated user.
*   **`GET /example`**: An example protected route that demonstrates using multiple services.

## Setup and Running the Application

### 1. Prerequisites

*   Go (version 1.22+ recommended)
*   PostgreSQL
*   A database migration tool like `golang-migrate/migrate`.

### 2. Installation

1.  **Clone the repository** (if applicable).

2.  **Install Migration Tool**:
    ```sh
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

3.  **Configure the Database**:
    *   Ensure your PostgreSQL server is running.
    *   Create a database for the project (e.g., `CREATE DATABASE workout_db;`).
    *   Verify the `dsn` in `configs/config.yaml` matches your database credentials.

4.  **Run Database Migrations**:
    Execute the following command from the project root, replacing the DSN if it's different from the one in the config file.
    ```sh
    migrate -database "postgres://admin:admin@localhost:5432/workout_db?sslmode=disable" -path migrations up
    ```

5.  **Install Go Dependencies**:
    ```sh
    go mod tidy
    ```

6.  **Run the Server**:
    ```sh
    go run ./cmd/server
    ```
    The server will start and listen on `http://localhost:8080`.

---
