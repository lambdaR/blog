# Microservices

The Micro Blog project is divided into several microservices, each responsible for a specific domain. This approach allows for better separation of concerns, independent development, and scalability.

## Service Breakdown

### Users Service

The Users Service is responsible for user management:

- User registration
- User authentication
- User profile management
- User data storage

**Key Components:**

- `users/main.go`: Service entry point
- `users/handler/users.go`: Request handlers
- `users/proto/users.proto`: Service definition

**Data Model:**
```go
message User {
  string id = 1;
  string name = 2;
  string email = 3;
  string password = 4; // Stored as hash
}
```

### Posts Service

The Posts Service manages blog posts and their tags:

- Post creation
- Post retrieval
- Post listing
- Tag management

**Key Components:**

- `posts/main.go`: Service entry point
- `posts/handler/posts.go`: Request handlers
- `posts/proto/posts.proto`: Service definition

**Data Model:**
```go
message Post {
  string id = 1;
  string title = 2;
  string content = 3;
  string author_id = 4;
  string author_name = 5;
  int64 created_at = 6;
  repeated string tags = 7;
}
```

### Comments Service

The Comments Service handles comments on blog posts:

- Comment creation
- Comment retrieval
- Comment listing by post

**Key Components:**

- `comments/main.go`: Service entry point
- `comments/handler/comments.go`: Request handlers
- `comments/proto/comments.proto`: Service definition

**Data Model:**
```go
message Comment {
  string id = 1;
  string content = 2;
  string author_id = 3;
  string author_name = 4;
  string post_id = 5;
  int64 created_at = 6;
}
```

### Web Service

The Web Service acts as an API gateway and serves the static web UI:

- REST API endpoints
- Service orchestration
- Static file serving
- Session management

**Key Components:**

- `web/main.go`: Service entry point and API implementation
- `web/static/`: Static web UI files

## Service Independence

Each microservice in the Micro Blog project:

1. **Has its own codebase**: Each service has its own directory with dedicated code
2. **Manages its own data**: Each service has its own data store
3. **Exposes a well-defined API**: Services communicate through Protocol Buffer definitions
4. **Can be deployed independently**: Services can be run and scaled separately

## Service Registration

All services register themselves with go-micro's service registry:

```go
service := micro.New("service-name")
```

This allows other services to discover and communicate with them without hardcoding addresses.