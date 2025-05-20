# Micro Blog

A simple blog system using microservices architecture with go-micro v5.

## Services

The project consists of the following microservices:

1. **Users Service**: User management (create, read, update, delete)
2. **Posts Service**: Post management (create, read, delete, list)
3. **Comments Service**: Comment management (create, read, delete, list)
4. **Web Service**: REST API that uses all other services

## Getting Started

### Prerequisites

- Go 1.24 or higher
- go-micro v5

### Launching Services

Run each service in a separate terminal:

```bash
# Users Service
cd users
go run main.go

# Posts Service
cd posts
go run main.go

# Comments Service
cd comments
go run main.go

# Web API (REST)
cd web
go run main.go
```

The REST API will be available at http://localhost:42096

## API Endpoints

### Posts

- `GET /posts`: List all posts
- `GET /posts/:id`: Get a post by ID
- `POST /posts`: Create a new post
  ```json
  {
    "title": "Post title",
    "content": "Post content"
  }
  ```

### Comments

- `GET /comments?post_id=123`: List comments (optionally filtered by post_id)
- `POST /comments`: Add a comment
  ```json
  {
    "content": "Comment content",
    "author_id": "user_id",
    "post_id": "post_id"
  }
  ```

### Users

- `GET /users/:id`: Get a user by ID
- `POST /users`: Create a new user
  ```json
  {
    "name": "User Name",
    "email": "email@example.com"
  }
  ```

## Project Structure

```
blog/
├── comments/           # Comments service
│   ├── handler/        # Request handlers
│   ├── main.go         # Entry point
│   └── proto/          # Protobuf definitions
├── posts/              # Posts service
│   ├── handler/
│   ├── main.go
│   └── proto/
├── users/              # Users service
│   ├── handler/
│   ├── main.go
│   └── proto/
└── web/                # REST API
    └── main.go
```