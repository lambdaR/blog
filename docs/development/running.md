# Running the Services

This guide explains how to run the Micro Blog services for development and testing.

## Running Options

You have two options to run the Micro Blog services:

1. Using the Micro CLI
2. Using the Makefile (plain Go)

## Option 1: Using Micro CLI

If you have the [Micro CLI](https://github.com/micro/micro) installed, you can run all services at once:

```bash
# From the project root
micro run --all
```

This will start all services defined in the project.

To run a specific service:

```bash
# Run just the users service
micro run users

# Run just the posts service
micro run posts

# Run just the comments service
micro run comments

# Run just the web service
micro run web
```

To stop services:

```bash
# Stop all services
micro kill --all

# Stop a specific service
micro kill users
```

## Option 2: Using the Makefile

The project includes a Makefile with commands to run the services using plain Go:

```bash
# Run all services in parallel
make run-all

# Run individual services
make run-users
make run-posts
make run-comments
make run-web
```

When using `make run-all`, all services will run in parallel, and the process will wait for all of them to complete.

## Running Services Manually

You can also run each service manually:

```bash
# Run the users service
cd users
go run main.go

# Run the posts service
cd posts
go run main.go

# Run the comments service
cd comments
go run main.go

# Run the web service
cd web
go run main.go
```

## Service Dependencies

The services have the following dependencies:

- **Web Service**: Depends on Users, Posts, and Comments services
- **Posts Service**: Independent
- **Comments Service**: Independent
- **Users Service**: Independent

For full functionality, all services should be running. However, you can run services independently for development and testing.

## Accessing the Application

Once all services are running:

1. The REST API will be available at http://localhost:42096
2. The web interface will be available at http://localhost:42096

## Verifying Services

To verify that services are running correctly:

1. Open http://localhost:42096 in your browser
2. You should see the blog interface
3. Try signing up for an account
4. Create a post
5. Add comments to posts

## Troubleshooting

If you encounter issues:

1. **Service not found errors**:
   - Make sure all services are running
   - Check that the service names match (users, posts, comments)

2. **Connection refused errors**:
   - Check that the services are running on the expected ports
   - Make sure no firewall is blocking the connections

3. **Web interface not loading**:
   - Check that the web service is running
   - Verify that static files are in the correct location (`web/static/`)

4. **Data persistence issues**:
   - Remember that the default store is in-memory
   - Data will be lost when services are restarted