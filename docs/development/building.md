# Building the Services

This guide explains how to build the Micro Blog services for deployment.

## Building with the Makefile

The project includes a Makefile with commands to build all services:

```bash
# Build all services
make build-all

# Build individual services
make build-users
make build-posts
make build-comments
make build-web
```

The built binaries will be placed in the `bin/` directory:

- `bin/users`: Users service binary
- `bin/posts`: Posts service binary
- `bin/comments`: Comments service binary
- `bin/web`: Web service binary

## Building Manually

You can also build each service manually:

```bash
# Build the users service
cd users
go build -o ../bin/users

# Build the posts service
cd posts
go build -o ../bin/posts

# Build the comments service
cd comments
go build -o ../bin/comments

# Build the web service
cd web
go build -o ../bin/web
```

## Cross-Compilation

To build for a different platform, you can use Go's cross-compilation capabilities:

```bash
# Example: Build for Linux on a different platform
GOOS=linux GOARCH=amd64 go build -o bin/users-linux-amd64 ./users

# Example: Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/users-windows-amd64.exe ./users

# Example: Build for macOS
GOOS=darwin GOARCH=amd64 go build -o bin/users-darwin-amd64 ./users
```

## Building Docker Images

You can build Docker images for each service:

1. Create a Dockerfile for each service:

```dockerfile
# Example Dockerfile for the users service
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o users ./users

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/users /app/
COPY --from=builder /app/users/static /app/static

EXPOSE 8080
CMD ["/app/users"]
```

2. Build the Docker images:

```bash
docker build -t micro-blog-users -f Dockerfile.users .
docker build -t micro-blog-posts -f Dockerfile.posts .
docker build -t micro-blog-comments -f Dockerfile.comments .
docker build -t micro-blog-web -f Dockerfile.web .
```

## Deployment Considerations

When building for deployment, consider the following:

1. **Configuration**: Ensure that services are configured to use appropriate data stores and service discovery mechanisms for production.

2. **Environment Variables**: Use environment variables for configuration in production environments.

3. **Static Files**: Make sure that the web service has access to the static files in the deployment environment.

4. **Security**: Use HTTPS in production and secure any sensitive endpoints.

5. **Monitoring**: Add monitoring and logging for production deployments.

## Continuous Integration

For continuous integration, you can use the Makefile commands in your CI pipeline:

```yaml
# Example GitHub Actions workflow
name: Build and Test

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.24
    - name: Generate Protocol Buffers
      run: make gen-proto
    - name: Build
      run: make build-all
    - name: Test
      run: go test ./...
```