# Service Communication

The Micro Blog project uses several communication patterns to enable services to interact with each other and with clients.

## RPC Communication

Services communicate with each other using Remote Procedure Calls (RPC) based on Protocol Buffers and gRPC, which is abstracted by the go-micro framework.

### Protocol Buffers

Protocol Buffers (protobuf) are used to define service interfaces and message types. For example, the Users service is defined in `users/proto/users.proto`:

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

// Request and response message definitions...
```

### Service Clients

Services communicate with each other by creating clients for the services they need to call:

```go
// In web/main.go
postClient := postProto.NewPostsService("posts", service.Client())
commentClient := commentProto.NewCommentsService("comments", service.Client())
userClient := userProto.NewUsersService("users", service.Client())
```

### Making RPC Calls

Services make RPC calls to other services using the generated client:

```go
// Example: Web service calling the Posts service
resp, err := postClient.List(context.Background(), &postProto.ListRequest{
    Page:  1,
    Limit: 10,
})
```

## REST API

The Web service exposes a REST API for external clients using the Gin framework:

```go
// Example: REST endpoint for listing posts
router.GET("/posts", func(c *gin.Context) {
    resp, err := postClient.List(context.Background(), &postProto.ListRequest{
        Page:  1,
        Limit: 10,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
})
```

## Service Discovery

Services discover each other using go-micro's built-in service registry:

1. **Registration**: Each service registers itself with a unique name
2. **Discovery**: Services look up other services by name
3. **Load Balancing**: Multiple instances of the same service can be load balanced automatically

## Data Serialization

Data is serialized in different formats depending on the communication channel:

- **Between Services**: Protocol Buffers (binary format)
- **REST API**: JSON
- **Web UI to API**: JSON

## Authentication and Sessions

The Web service manages user authentication and sessions:

1. **Session Storage**: Uses cookie-based sessions with the `gin-contrib/sessions` package
2. **Authentication Flow**: 
   - User logs in via the `/login` endpoint
   - Session is created with user ID and name
   - Subsequent requests include the session cookie
   - Middleware extracts user information from the session

## Error Handling

Error handling in service communication:

1. **Service Errors**: Returned as error objects in RPC responses
2. **REST API Errors**: Converted to appropriate HTTP status codes and JSON error messages
3. **Retry Logic**: Not implemented in the basic version but could be added for resilience