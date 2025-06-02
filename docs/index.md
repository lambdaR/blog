# Micro Blog

A simple blog system built using microservices architecture with go-micro v5.

![Micro Blog Architecture](https://github.com/user-attachments/assets/c7c74c36-32da-4120-af88-254d00ecf8c6)

## Project Overview

This project demonstrates how to build a complete blog application using microservices architecture with the go-micro v5 framework. The application consists of multiple independent services that work together to provide a full-featured blog platform.

### Key Features

- User management (registration, authentication)
- Post creation and management
- Comment system
- Tag management for posts
- REST API
- Simple web interface

## Go-Micro Framework

Go-Micro is a pluggable microservices framework for Go that provides the core requirements for distributed systems development. This project uses go-micro v5, which includes several key components:

- **Service**: The main building block that encapsulates a microservice
- **Registry**: Service discovery mechanism for registering and finding services
- **Client**: RPC client for making requests to services
- **Server**: RPC server for handling requests
- **Broker**: Asynchronous messaging between services
- **Transport**: Synchronous communication mechanism
- **Store**: Simple key-value storage interface

Go-Micro abstracts away the details of distributed systems, allowing developers to focus on business logic rather than boilerplate code for service discovery, load balancing, and communication.

## Microservices Architecture

The project is built using a microservices architecture, with each service responsible for a specific domain:

1. **Users Service**: Handles user management (create, read, update, delete)
2. **Posts Service**: Manages blog posts and tags (create, read, delete, list, tag management)
3. **Comments Service**: Manages comments on posts (create, read, delete, list)
4. **Web Service**: Provides a REST API that integrates all other services and serves a static web UI

## Technology Stack

- **Go**: Programming language (v1.24+)
- **go-micro v5**: Microservices framework
- **Protocol Buffers**: For service definitions and communication
- **Gin**: Web framework for the REST API
- **HTML/CSS/JavaScript**: For the static web UI

## Getting Started

To get started with the Micro Blog project:

1. [Set up your development environment](development/setup.md)
2. [Learn how to run the services](development/running.md)
3. [Explore the architecture](architecture/overview.md)
4. [Understand the API](api/rest.md)

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
└── web/                # REST API and static web UI
    ├── main.go         # REST API server
    └── static/         # Static web UI
```