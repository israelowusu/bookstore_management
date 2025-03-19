# Bookstore Management API

A RESTful API built with Go for managing a bookstore's inventory. This project uses Gorilla Mux for routing, GORM with MySQL for database operations, and Swagger for API documentation.

## Features

- CRUD operations for books
- RESTful API endpoints
- MySQL database integration
- Swagger API documentation
- Structured project layout

## Prerequisites

- Go 1.24 or higher
- MySQL
- Git
- swag CLI tool (`go install github.com/swaggo/swag/cmd/swag@latest`)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/israelowusu/bookstore_management.git
cd bookstore_management
```

2. Install dependencies:
```bash
go mod download
```

3. Generate Swagger documentation:
```bash
swag init --parseDependency --parseInternal -g cmd/main/main.go
```

4. Configure the database:
   - Create a MySQL database
   - Update the database configuration in `pkg/config/app.go`

## Database Configuration

Update the MySQL connection string in `pkg/config/app.go`:

```go
mysql.Open("user:password@tcp(localhost:3306)/database?charset=utf8&parseTime=True&loc=Local")
```

Replace `user`, `password`, and `database` with your MySQL credentials.

## API Documentation

The API documentation is available through Swagger UI:

- Swagger UI: `http://localhost:8080/swagger/index.html`
- API JSON Spec: `http://localhost:8080/swagger/doc.json`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/book/` | Get all books |
| GET | `/book/{bookId}` | Get a specific book |
| POST | `/book/` | Create a new book |
| PUT | `/book/{bookId}` | Update a book |
| DELETE | `/book/{bookId}` | Delete a book |

## Running the Application

```bash
go run cmd/main/main.go
```

The server will start at `http://localhost:8080`

## Project Structure

```
├── cmd/
│   └── main/
│       └── main.go
├── pkg/
│   ├── config/
│   │   └── app.go
│   ├── controllers/
│   │   └── book-controller.go
│   ├── docs/
│   │   └── docs.go
│   ├── models/
│   │   └── book.go
│   ├── routes/
│   │   └── bookstore-routes.go
│   └── utils/
│       └── utils.go
└── README.md
```

## Development

To regenerate Swagger documentation after API changes:

```bash
swag init --parseDependency --parseInternal -g cmd/main/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.