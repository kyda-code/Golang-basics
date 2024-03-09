# Go Tasks API Project

This project is a simple Go API that manages a task list. It provides a set of endpoints for CRUD operations on tasks.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Endpoints](#endpoints)
3. [Running the Application](#running-the-application)

## Project Structure

The application logic is mainly contained in the `Handler` struct and its methods.

Methods include:

- `TasksEndpoint`: This main function dispatches the various sub-routes to the appropriate HTTP handlers.
- `handleGetTasks`: This handler fetches all tasks from the DB and returns them in a JSON response.
- `handleGetTaskById`: This handler fetches a specific task using its ID and returns it in a JSON response.

## Endpoints

The API exposes the following endpoints:

- `GET /tasks`: Returns all tasks
- `GET /tasks/{id}`: Returns a specific task identified by `id`
- `POST /tasks`: Creates and saves a new task
- `PUT /tasks/{id}`: Updates the task identified by `id`
- `DELETE /tasks/{id}`: Deletes the task identified by `id`

## Running the Application

To run the application, make sure Go is installed on your system. Further, you need to have the Go project downloaded on your machine.

Here are the steps to run the application:

1. Open the terminal/command prompt.
2. Navigate to the directory containing the `main.go` file.
3. Execute `go run main.go` to start the server.
4. You can then make requests to the API at `localhost:8080/tasks`.

You can test the API using tools like Postman, cURL, etc.