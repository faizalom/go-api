# Go REST API: Workout Tracker

This project is a RESTful API for a workout tracking application, built in Go. It features a clean architecture, JWT-based authentication, and user management.

## Features

*   **Clean Architecture:** Follows a layered architecture (Handler, Service, Repository) with dependency injection.
*   **Modular Routing:** Routes are organized into modules, with each module handling its own dependencies.
*   **JWT Authentication:** Endpoints are secured using JWT, with token generation (`/login`) and middleware validation.
*   **User CRUD:** Full support for creating, retrieving, updating, deleting, and listing users.
*   **Configuration Management:** All settings are managed via a `config.yaml` file.
*   **Structured Logging:** Centralized logger with different levels (`Info`, `Error`).
*   **Database Migrations:** Schema changes are managed through SQL migration files.
*   **API Documentation:** OpenAPI 3.0 specification provided in `api/openapi.yaml`.

## Getting Started

### Prerequisites

*   Go (version 1.22+ recommended)
*   PostgreSQL
*   [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    ```

2.  **Install Migration Tool:**
    ```sh
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

3.  **Configure the Database:**
    *   Make sure your PostgreSQL server is running.
    *   Create a database (e.g., `CREATE DATABASE workout_db;`).
    *   Copy `configs/config.example.yaml` to `configs/config.local.yaml` and update the `dsn` with your database credentials.

4.  **Run Database Migrations:**
    ```sh
    migrate -database "postgres://admin:admin@localhost:5432/workout_db?sslmode=disable" -path migrations up
    ```
    *(Replace the DSN with your actual database connection string)*

5.  **Install Go Dependencies:**
    ```sh
    go mod tidy
    ```

6.  **Run the Server:**
    ```sh
    go run ./cmd/server
    ```
    The server will start and listen on `http://localhost:8080`.

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Public

*   `POST /login`: Get a new JWT.

### Protected (Requires `Authorization: Bearer <token>`)

*   `GET /profile`: Get the authenticated user's profile.
*   `GET /example`: An example protected route.
*   `GET /users`: List all users.
*   `POST /users`: Create a new user.
*   `GET /users/{id}`: Get a user by ID.
*   `PUT /users/{id}`: Update a user.
*   `DELETE /users/{id}`: Delete a user.
