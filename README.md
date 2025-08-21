# CMS API (Golang + Gin + PostgreSQL)

This is a simple **Content Management System (CMS) API** built with **Golang**, **Gin**, **PostgreSQL**, and **JWT authentication**.

## Features
- User registration & login (JWT auth)
- Authenticated endpoints
- Category CRUD
- Content CRUD
- Docker & docker-compose setup
- Auto database migrations via GORM

---

## 🚀 Getting Started

### Prerequisites
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### Run with Docker
```bash
docker-compose up --build
```

### Services
- API: http://localhost:8080
- PostgreSQL: localhost:5432
- user: postgres
- password: postgres
- database: cms


### API Endpoints
Auth
POST /api/register → Register user
POST /api/login → Login and get JWT token

Users
GET /api/me → Get current logged-in user (JWT required)

#### Categories (JWT required)
- POST /api/categories → Create category
- GET /api/categories → List categories
- GET /api/categories/:id → Get category by ID
- PUT /api/categories/:id → Update category
- DELETE /api/categories/:id → Delete category

#### Contents (JWT required)
- POST /api/contents → Create content
- GET /api/contents → List contents
- GET /api/contents/:id → Get content by ID
- PUT /api/contents/:id → Update content
- DELETE /api/contents/:id → Delete content

## Running Tests
```bash
go test ./tests -v
```

## Environment Variables
These are set in `docker-compose.yml`
```bash
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=cms
DB_PORT=5432

```

TODO Swagger