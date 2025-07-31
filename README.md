# Running Fund

A Go-based RESTful CRUD application for managing to-do list.

## 🛠️ Technologies Used

- Backend: Go

- Database: PostgreSQL

- Containerization: Docker

## ⚙️ Setup Instructions

## Prerequisites

Ensure you have the following installed:

[Go](https://go.dev/doc/install) Version 1.23 or higher

[Docker](https://docs.docker.com/engine/install/)

[Docker Compose](https://docs.docker.com/compose/install/)

## Local Development Instructions

Clone the repository:

```sh
git clone git@github.com:poomipat-k/Arise-Gin-GORM-CRUD.git
cd Arise-Gin-GORM-CRUD
```

Copy the example environment variables:

```sh
cp .env.example .env
```

Build and start the application using Docker Compose:

```sh
docker compose up
```

Access the application at `http://localhost:8080`

Running Tests
To run the tests:

```sh
cd ./backend
go test ./...
```