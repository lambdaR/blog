# Users Service

The Users Service is responsible for managing user accounts in the Micro Blog system. It handles user creation, retrieval, updating, and deletion.

## Service Overview

The Users Service provides the following functionality:

- User registration
- User profile retrieval
- User profile updates
- User deletion
- User listing

## Implementation

### Main Service File

The main service file (`users/main.go`) initializes the service and registers the handler:

```go
package main

import (
    "go-micro.dev/v5"

    "github.com/micro/blog/users/handler"
    pb "github.com/micro/blog/users/proto"
)

func main() {
    service := micro.New("users")

    pb.RegisterUsersHandler(service.Server(), handler.New())

    service.Init()

    service.Run()
}
```

### Handler Implementation

The handler (`users/handler/users.go`) implements the service logic:

```go
package handler

import (
    "context"
    "encoding/json"

    "github.com/google/uuid"
    pb "github.com/micro/blog/users/proto"
    "go-micro.dev/v5/store"
)

var userStore = store.DefaultStore

type Handler struct{}

func New() *Handler {
    return &Handler{}
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
    id := uuid.New().String()

    user := &pb.User{
        Id:       id,
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password, // store hash
    }

    rsp.User = user

    // Save to store
    b, err := json.Marshal(user)
    if err == nil {
        _ = userStore.Write(&store.Record{Key: "user-" + user.Id, Value: b})
    }

    return nil
}

// Other methods: Read, Update, Delete, List
```

## Data Storage

The Users Service uses go-micro's built-in store interface for data persistence:

- Each user is stored as a JSON document
- User records are keyed by `user-{id}`
- The default store implementation is used (memory store in development)

## Protocol Definition

The service interface is defined in Protocol Buffers (`users/proto/users.proto`):

```protobuf
syntax = "proto3";

package users;

option go_package = "github.com/micro/blog/users/proto;users";

service Users {
  rpc Create(CreateRequest) returns (CreateResponse) {}
  rpc Read(ReadRequest) returns (ReadResponse) {}
  rpc Update(UpdateRequest) returns (UpdateResponse) {}
  rpc Delete(DeleteRequest) returns (DeleteResponse) {}
  rpc List(ListRequest) returns (ListResponse) {}
}

message User {
  string id = 1;
  string name = 2;
  string email = 3;
  string password = 4;
}

message CreateRequest {
  string name = 1;
  string email = 2;
  string password = 3;
}

message CreateResponse {
  User user = 1;
}

// Other request/response message definitions...
```

## Authentication

The Users Service stores password hashes but does not directly handle authentication:

- Password hashing is performed by the Web Service
- Session management is handled by the Web Service
- The Users Service only stores and retrieves user data

## Service Usage

Other services interact with the Users Service through the generated client:

```go
// Create a client
userClient := userProto.NewUsersService("users", service.Client())

// Call methods
resp, err := userClient.Read(context.Background(), &userProto.ReadRequest{
    Id: userId,
})
```