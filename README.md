# user-service
A user microservice built with Golang (Gin framework), GORM ORM, and MySQL. Supports user creation and retrieval with Dockerized deployment.

# 🧑‍💻 User Microservice – Go + Gin + GORM + MySQL

This is a lightweight and production-ready **User Microservice** built with:

- ⚙️ **Golang** using the **Gin** web framework
- 🗃️ **MySQL** as the relational database
- 🧠 **GORM** for ORM-based database interaction
- 🐳 **Docker & Docker Compose** for containerization

---

## ✨ Features

- 🔐 User registration & authentication
- 🔍 Get user by ID/email
- 🧾 GORM auto-migration for database tables
- 📦 Clean architecture (handlers, services, repository)
- 🔄 Environment-based config (.env)
- 🐳 Containerized with Docker Compose

---

## 🚀 Tech Stack

| Component       | Tool                          |
|----------------|-------------------------------|
| Language        | Golang (1.20+)                |
| Web Framework   | Gin                           |
| ORM             | GORM                          |
| Database        | MySQL                         |
| Containerization| Docker, Docker Compose        |
| Env Config      | dotenv / os.Getenv            |

---

## 🧪 API Endpoints (example)

- `POST /users` – Create a new user
- `GET /users/:id` – Get user by ID
- `POST /login` – Authenticate user

---

## 🐳 Quick Start with Docker

```bash
# Start everything
docker-compose up --build

# Stop all
docker-compose down
