# gRPC API

The Micro Blog system uses gRPC for internal communication between microservices. Each service exposes a gRPC API defined using Protocol Buffers.

## Protocol Definitions

### Users Service

The Users Service API is defined in `users/proto/users.proto`:

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

message ReadRequest {
  string id = 1;
}

message ReadResponse {
  User user = 1;
}

message UpdateRequest {
  string id = 1;
  string name = 2;
  string email = 3;
}

message UpdateResponse {
  User user = 1;
}

message DeleteRequest {
  string id = 1;
}

message DeleteResponse {}

message ListRequest {
  int32 page = 1;
  int32 limit = 2;
}

message ListResponse {
  repeated User users = 1;
  int32 total = 2;
}
```

### Posts Service

The Posts Service API is defined in `posts/proto/posts.proto`:

```protobuf
syntax = "proto3";

package posts;

option go_package = "github.com/micro/blog/posts/proto;posts";

service Posts {
  rpc Create(CreateRequest) returns (CreateResponse) {}
  rpc Read(ReadRequest) returns (ReadResponse) {}
  rpc Delete(DeleteRequest) returns (DeleteResponse) {}
  rpc List(ListRequest) returns (ListResponse) {}
  rpc TagPost(TagPostRequest) returns (TagPostResponse) {}
  rpc UntagPost(UntagPostRequest) returns (UntagPostResponse) {}
  rpc ListTags(ListTagsRequest) returns (ListTagsResponse) {}
}

message Post {
  string id = 1;
  string title = 2;
  string content = 3;
  string author_id = 4;
  string author_name = 5;
  int64 created_at = 6;
  repeated string tags = 7;
}

message CreateRequest {
  string title = 1;
  string content = 2;
  string author_id = 3;
  string author_name = 4;
}

message CreateResponse {
  Post post = 1;
}

message ReadRequest {
  string id = 1;
}

message ReadResponse {
  Post post = 1;
}

message DeleteRequest {
  string id = 1;
}

message DeleteResponse {}

message ListRequest {
  int32 page = 1;
  int32 limit = 2;
}

message ListResponse {
  repeated Post posts = 1;
  int32 total = 2;
}

message TagPostRequest {
  string post_id = 1;
  string tag = 2;
}

message TagPostResponse {
  Post post = 1;
}

message UntagPostRequest {
  string post_id = 1;
  string tag = 2;
}

message UntagPostResponse {
  Post post = 1;
}

message ListTagsRequest {
  string post_id = 1;
}

message ListTagsResponse {
  repeated string tags = 1;
}
```

### Comments Service

The Comments Service API is defined in `comments/proto/comments.proto`:

```protobuf
syntax = "proto3";

package comments;

option go_package = "github.com/micro/blog/comments/proto;comments";

service Comments {
  rpc Create(CreateRequest) returns (CreateResponse) {}
  rpc Read(ReadRequest) returns (ReadResponse) {}
  rpc Delete(DeleteRequest) returns (DeleteResponse) {}
  rpc List(ListRequest) returns (ListResponse) {}
}

message Comment {
  string id = 1;
  string content = 2;
  string author_id = 3;
  string author_name = 4;
  string post_id = 5;
  int64 created_at = 6;
}

message CreateRequest {
  string content = 1;
  string author_id = 2;
  string author_name = 3;
  string post_id = 4;
}

message CreateResponse {
  Comment comment = 1;
}

message ReadRequest {
  string id = 1;
}

message ReadResponse {
  Comment comment = 1;
}

message DeleteRequest {
  string id = 1;
}

message DeleteResponse {}

message ListRequest {
  string post_id = 1;
}

message ListResponse {
  repeated Comment comments = 1;
  int32 total = 2;
}
```

## Using the gRPC API

### Generating Client Code

To use the gRPC API, you need to generate client code from the Protocol Buffer definitions:

```bash
# Generate Go code
make gen-proto
```

This will generate:
- `*.pb.go`: Protocol Buffer message definitions
- `*.pb.micro.go`: go-micro service client and server code

### Creating a Client

To create a client for a service:

```go
import (
    "go-micro.dev/v5"
    userProto "github.com/micro/blog/users/proto"
)

// Create a new service
service := micro.NewService(micro.Name("client"))
service.Init()

// Create a client for the Users service
userClient := userProto.NewUsersService("users", service.Client())
```

### Making RPC Calls

To make RPC calls to a service:

```go
// Create a new user
createResp, err := userClient.Create(context.Background(), &userProto.CreateRequest{
    Name:  "User Name",
    Email: "user@example.com",
})
if err != nil {
    // Handle error
}
user := createResp.User

// Read a user
readResp, err := userClient.Read(context.Background(), &userProto.ReadRequest{
    Id: user.Id,
})
if err != nil {
    // Handle error
}
```

## Error Handling

gRPC errors are returned as `error` objects in Go. The error can be examined using the `status` package:

```go
import (
    "google.golang.org/grpc/status"
)

resp, err := userClient.Read(context.Background(), &userProto.ReadRequest{
    Id: "invalid-id",
})
if err != nil {
    if st, ok := status.FromError(err); ok {
        // Handle specific gRPC error
        switch st.Code() {
        case codes.NotFound:
            // Handle not found
        case codes.InvalidArgument:
            // Handle invalid argument
        default:
            // Handle other errors
        }
    } else {
        // Handle non-gRPC error
    }
}
```

## Service Discovery

go-micro handles service discovery automatically:

1. Services register themselves with the registry
2. Clients discover services by name
3. Load balancing is handled automatically for multiple instances

## Streaming

The current API does not use streaming, but go-micro supports gRPC streaming if needed in future versions.