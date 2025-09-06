# 📝 Task Manager API

A simple and robust RESTful API for managing tasks, built with Go.  
Easily create, read, update, and delete your tasks with a clean architecture and PostgreSQL as the database.

---

## 🚀 Features

- Create, read, update, and delete tasks
- Set due dates and completion status
- Input validation and structured API responses
- Environment-based configuration
- Dockerized for easy development
- Hot reload with [Air](https://github.com/air-verse/air)

---

## 📁 Project Structure

```
.
├── cmd/                # Application entrypoint
│   └── server/
│       └── main.go
├── internal/           # Application core
│   ├── config/         # Environment config loader
│   ├── db/             # Database connection
│   └── task/           # Task domain (handlers, models, repository)
├── pkg/                # Shared utilities
│   └── utils/
├── tmp/                # Build artifacts
├── .env                # Environment variables
├── .air.toml           # Air config for live reload
├── docker-compose.yml  # Docker Compose setup
├── go.mod / go.sum     # Go modules
└── README.md
```

---

## ⚡️ Quick Start

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/task-manager-api.git
cd task-manager-api
```

### 2. Configure Environment

Copy `.env.example` to `.env` and update the values as needed:

```sh
cp .env.example .env
```

### 3. Run with Docker

```sh
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

### 4. Run Locally (with hot reload)

Install [Air](https://github.com/cosmtrek/air):

```sh
go install github.com/air-verse/air@latest
```

Then start the server:

```sh
air
```

---

## 🛠️ API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/tasks`      | List all tasks    |
| GET    | `/tasks/{id}` | Get a task by ID  |
| POST   | `/tasks`      | Create a new task |
| PUT    | `/tasks/{id}` | Update a task     |
| DELETE | `/tasks/{id}` | Delete a task     |

### Example: Create Task

```http
POST /tasks
Content-Type: application/json

{
  "title": "Write documentation",
  "description": "Complete the project README",
  "dueDate": "06/09/2025"
}
```

---

## 🧩 Tech Stack

- **Go** — Backend language
- **PostgreSQL** — Database
- **Docker** — Containerization
- **Air** — Live reload for Go

---

## 📄 License

MIT License. See [LICENSE](LICENSE) for details.

---

## 🙌 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

> Made with ❤️ by Muhammad Putra Age
