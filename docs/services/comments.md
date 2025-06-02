# Comments Service

The Comments Service manages comments on blog posts in the Micro Blog system. It handles comment creation, retrieval, deletion, and listing comments for specific posts.

## Service Overview

The Comments Service provides the following functionality:

- Comment creation
- Comment retrieval
- Comment deletion
- Comment listing (all comments or filtered by post)

## Implementation

### Main Service File

The main service file (`comments/main.go`) initializes the service and registers the handler:

```go
package main

import (
    "go-micro.dev/v5"

    "github.com/micro/blog/comments/handler"
    pb "github.com/micro/blog/comments/proto"
)

func main() {
    service := micro.New("comments")

    pb.RegisterCommentsHandler(service.Server(), handler.New())

    service.Init()

    service.Run()
}
```

### Handler Implementation

The handler (`comments/handler/comments.go`) implements the service logic:

```go
package handler

import (
    "context"
    "encoding/json"
    "time"

    "github.com/google/uuid"
    pb "github.com/micro/blog/comments/proto"
    "go-micro.dev/v5/store"
)

var commentStore = store.DefaultStore

type Handler struct{}

func New() *Handler {
    return &Handler{}
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
    id := uuid.New().String()
    now := time.Now().Unix()

    comment := &pb.Comment{
        Id:         id,
        Content:    req.Content,
        AuthorId:   req.AuthorId,
        AuthorName: req.AuthorName,
        PostId:     req.PostId,
        CreatedAt:  now,
    }

    rsp.Comment = comment

    // Save to store
    b, err := json.Marshal(comment)
    if err == nil {
        _ = commentStore.Write(&store.Record{Key: "comment-" + comment.Id, Value: b})
    }

    return nil
}

// Other methods: Read, Delete, List
```

### Filtering Comments by Post

The Comments Service includes functionality to filter comments by post ID:

```go
func (h *Handler) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
    var records []*store.Record
    var err error

    // If post_id is provided, list comments for that post
    if req.PostId != "" {
        // We need to scan all comments and filter by post_id
        records, err = commentStore.Read("comment-", store.ReadPrefix())
        if err != nil {
            return err
        }

        var comments []*pb.Comment
        for _, record := range records {
            var comment pb.Comment
            if err := json.Unmarshal(record.Value, &comment); err != nil {
                continue
            }
            
            // Filter by post_id
            if comment.PostId == req.PostId {
                comments = append(comments, &comment)
            }
        }

        rsp.Comments = comments
        rsp.Total = int32(len(comments))
        return nil
    }

    // Otherwise, list all comments
    records, err = commentStore.Read("comment-", store.ReadPrefix())
    if err != nil {
        return err
    }

    var comments []*pb.Comment
    for _, record := range records {
        var comment pb.Comment
        if err := json.Unmarshal(record.Value, &comment); err != nil {
            continue
        }
        comments = append(comments, &comment)
    }

    rsp.Comments = comments
    rsp.Total = int32(len(comments))
    return nil
}
```

## Data Storage

The Comments Service uses go-micro's built-in store interface for data persistence:

- Each comment is stored as a JSON document
- Comment records are keyed by `comment-{id}`
- The default store implementation is used (memory store in development)

## Protocol Definition

The service interface is defined in Protocol Buffers (`comments/proto/comments.proto`):

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

// Request and response message definitions...
```

## Service Usage

Other services interact with the Comments Service through the generated client:

```go
// Create a client
commentClient := commentProto.NewCommentsService("comments", service.Client())

// Call methods
resp, err := commentClient.List(context.Background(), &commentProto.ListRequest{
    PostId: postId,
})
```