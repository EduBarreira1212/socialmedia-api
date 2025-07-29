````md
# 📦 Project Name

Social media API

---

## 📑 Table of Contents

- [About](#about)
- [Tech Stack](#tech-stack)
- [Features](#features)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Running the Project](#running-the-project)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [License](#license)

---

## 📖 About

This project is the backend of a social media, it's possible to make CRUD operations for users, followers and posts

---

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Database:** MYSQL
- **Authentication:** JWT
- **Others:** Docker

---

## ✨ Features

- [x] Feature 1 (e.g., User authentication with JWT)
- [x] Feature 2 (e.g., CRUD operations for users)
- [x] Feature 3 (e.g., CRUD operations for followers)
- [x] Feature 4 (e.g., CRUD operations for posts)
- [x] Feature 5 (e.g., To like a post)
- [x] Feature 6 (e.g., To unlike a post)

---

## ⚙️ Getting Started

### Prerequisites

- Go ≥ 1.21
- Docker & Docker Compose (if using)

### Clone the repository

```bash
git clone https://github.com/EduBarreira1212/socialmedia-api.git
cd socialmedia-api
```
````

---

## 📁 Project Structure

```
socialmedia-api/
├── cmd/                # Entry point (main.go)
├── config/             # Config files and loading logic
├── internal/           # Core business logic (handlers, services, models)
│   ├── handler/
│   ├── service/
│   └── repository/
├── docs/               # Swagger / API documentation
├── test/               # Integration / unit tests
├── .env                # Environment variables
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

---

## 🔐 Environment Variables

Copy the `.env.example` to `.env` and update the values:

```bash
cp .env.example .env
```

| Variable       | Description                   |
| -------------- | ----------------------------- |
| `PORT`         | Port to run the server        |
| `DATABASE_URL` | PostgreSQL connection string  |
| `JWT_SECRET`   | Secret key for JWT            |
| `REDIS_URL`    | Redis connection string (opt) |

---

## 🚀 Running the Project

### Run Locally

```bash
go run cmd/main.go
```

### Run with Docker

```bash
docker-compose up --build
```

---

## 🧪 Testing

```bash
go test ./...
```

Or with coverage:

```bash
go test -cover ./...
```

---

## 📚 API Documentation

You can access the Swagger UI at:

```
http://localhost:PORT/docs
```

To generate docs:

```bash
swag init -g cmd/main.go
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 🙌 Contributing

1. Fork it
2. Create your feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -m 'Add my feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Open a pull request

---

## ✉️ Contact

For suggestions, bugs, or questions, contact [edubarreira1212@gmail.com](mailto:edubarreira1212@gmail.com)
