# Notification Service (Go)

This is a simple RESTful notification service written in Go. It provides API endpoints to send and retrieve notifications, using MongoDB for data persistence.

## Features

*   **REST API**: Simple and clean API for creating and listing notifications.
*   **MongoDB Integration**: Uses MongoDB as the backend database to store notification data.
*   **Configuration Management**: Easily configurable via environment variables.
*   **Dockerized**: Comes with a `Dockerfile` for building and running the service in a container.

## Project Structure

```
.
├── cmd/server/
│   └── main.go         # Application entry point
├── internal/
│   ├── api/            # HTTP handlers and router setup
│   ├── config/         # Configuration loading
│   ├── db/             # Database connection logic
│   ├── model/          # Data models (structs)
│   ├── repository/     # Data access layer (database operations)
│   └── service/        # Business logic
├── .dockerignore
├── .gitignore
├── Dockerfile
├── go.mod
└── go.sum
```

## API Endpoints

### Create a Notification

*   **URL**: `/notifications`
*   **Method**: `POST`
*   **Body**:

    ```json
    {
        "to": "user@example.com",
        "subject": "Hello!",
        "body": "This is the body of the notification."
    }
    ```

*   **Success Response**:
    *   **Code**: `201 Created`
    *   **Content**:
        ```json
        {
            "message": "Notification sent successfully"
        }
        ```

### Get All Notifications

*   **URL**: `/notifications`
*   **Method**: `GET`
*   **Success Response**:
    *   **Code**: `200 OK`
    *   **Content**: An array of notification objects.
        ```json
        [
            {
                "id": "6702a0b3e4b0c3a6b3e4b0c3",
                "to": "user@example.com",
                "subject": "Hello!",
                "body": "This is the body of the notification.",
                "createdAt": "2025-10-06T10:00:00Z"
            }
        ]
        ```

## Getting Started

### Prerequisites

*   Go (version 1.20 or later)
*   MongoDB
*   Docker (optional, for containerized deployment)

### Configuration

Create a `.env` file in the root of the project and add the following variables:

```env
# .env
MONGO_URI=mongodb://localhost:27017/notifications_db
PORT=8080
```

| Variable    | Description                                  | Default |
|-------------|----------------------------------------------|---------|
| `MONGO_URI` | The connection string for your MongoDB instance. | (none)  |
| `PORT`      | The port on which the service will run.      | `8080`  |


### Running Locally

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd notification-service-go
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Run the application:**
    ```sh
    go run cmd/server/main.go
    ```
    The service will start and be accessible at `http://localhost:8080`.

### Running with Docker

1.  **Build the Docker image:**
    ```sh
    docker build -t notification-service .
    ```

2.  **Run the Docker container:**
    
    *Note: Replace `<your-mongo-uri>` with your MongoDB connection string. If MongoDB is running on your host machine, you might need to use `host.docker.internal` (on Docker Desktop) or your machine's network IP instead of `localhost`.*

    ```sh
    docker run -p 8080:8080 --name notification-app -e MONGO_URI=<your-mongo-uri> -d notification-service
    ```