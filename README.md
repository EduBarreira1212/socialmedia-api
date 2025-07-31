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
- Docker & Docker Compose

### Clone the repository

```bash
git clone https://github.com/EduBarreira1212/socialmedia-api.git
cd socialmedia-api
```

---

## 📁 Project Structure

```
socialmedia-api/
├── .mysql-data/            # MySQL volume data (used by Docker)
├── sql/                    # SQL scripts or migrations
├── src/                    # Source code
│   ├── auth/               # Authentication logic (e.g., login, signup)
│   ├── config/             # Configuration (e.g., env loading)
│   ├── controllers/        # Route handlers / controllers
│   ├── database/           # Database connection and setup
│   ├── middlewares/        # HTTP middlewares (e.g., logging, auth)
│   ├── models/             # Structs for DB models and DTOs
│   ├── repositories/       # Data access layer
│   ├── responses/          # Standard API response formatting
│   ├── router/             # Routing setup (e.g., mux, echo, gin)
│   └── security/           # Security utilities (e.g., hashing, JWT)
├── .env                    # Environment variables
├── .gitignore              # Git ignore rules
├── docker-compose.yml      # Docker services definition
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies checksums
├── main.go                 # Application entry point
└── README.md               # Project documentation
```

---

## 🔐 Environment Variables

Copy the `.env.example` to `.env` and update the values:

```bash
cp .env.example .env
```

| Variable              | Description                                                 |
| --------------------- | ----------------------------------------------------------- |
| `MYSQL_ROOT_PASSWORD` | Root password for the MySQL server                          |
| `MYSQL_DATABASE`      | Name of the default MySQL database                          |
| `MYSQL_USER`          | MySQL user name                                             |
| `MYSQL_PASSWORD`      | Password for the MySQL user                                 |
| `PORT`                | Port for running the Go server                              |
| `SECRET_KEY`          | Secret key used for signing JWT tokens or other secure data |

---

## 🚀 Running the Project

### Run Locally

```bash
go run main.go
```

### Run with Docker

```bash
docker-compose up --build
```

## ✅ To Do

Config docker to run the api and not only the database

Add Swagger documentation

Implement tests

Add rate limiting middleware

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
