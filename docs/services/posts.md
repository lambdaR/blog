# Posts Service

The Posts Service manages blog posts in the Micro Blog system. It handles post creation, retrieval, deletion, and tag management.

## Service Overview

The Posts Service provides the following functionality:

- Post creation
- Post retrieval
- Post deletion
- Post listing
- Tag management (adding, removing, listing tags)
- Filtering posts by tag

## Implementation

### Main Service File

The main service file (`posts/main.go`) initializes the service and registers the handler:

```go
package main

import (
    "go-micro.dev/v5"

    "github.com/micro/blog/posts/handler"
    pb "github.com/micro/blog/posts/proto"
)

func main() {
    service := micro.New("posts")

    pb.RegisterPostsHandler(service.Server(), handler.New())

    service.Init()

    service.Run()
}
```

### Handler Implementation

The handler (`posts/handler/posts.go`) implements the service logic for managing posts and tags:

```go
package handler

import (
    "context"
    "encoding/json"
    "time"

    "github.com/google/uuid"
    pb "github.com/micro/blog/posts/proto"
    "go-micro.dev/v5/store"
)

var postStore = store.DefaultStore

type Handler struct{}

func New() *Handler {
    return &Handler{}
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
    id := uuid.New().String()
    now := time.Now().Unix()

    post := &pb.Post{
        Id:         id,
        Title:      req.Title,
        Content:    req.Content,
        AuthorId:   req.AuthorId,
        AuthorName: req.AuthorName,
        CreatedAt:  now,
        Tags:       []string{},
    }

    rsp.Post = post

    // Save to store
    b, err := json.Marshal(post)
    if err == nil {
        _ = postStore.Write(&store.Record{Key: "post-" + post.Id, Value: b})
    }

    return nil
}

// Other methods: Read, Delete, List, TagPost, UntagPost, ListTags
```

## Tag Management

The Posts Service includes special methods for tag management:

```go
func (h *Handler) TagPost(ctx context.Context, req *pb.TagPostRequest, rsp *pb.TagPostResponse) error {
    // Read the post
    rec, err := postStore.Read("post-" + req.PostId)
    if err != nil || len(rec) == 0 {
        return err
    }

    // Unmarshal post
    var post pb.Post
    if err := json.Unmarshal(rec[0].Value, &post); err != nil {
        return err
    }

    // Check if tag already exists
    for _, tag := range post.Tags {
        if tag == req.Tag {
            rsp.Post = &post
            return nil // Tag already exists
        }
    }

    // Add the tag
    post.Tags = append(post.Tags, req.Tag)

    // Save the updated post
    b, err := json.Marshal(&post)
    if err != nil {
        return err
    }
    if err := postStore.Write(&store.Record{Key: "post-" + post.Id, Value: b}); err != nil {
        return err
    }

    rsp.Post = &post
    return nil
}
```

## Data Storage

The Posts Service uses go-micro's built-in store interface for data persistence:

- Each post is stored as a JSON document
- Post records are keyed by `post-{id}`
- Tags are stored as string arrays within each post document
- The default store implementation is used (memory store in development)

## Protocol Definition

The service interface is defined in Protocol Buffers (`posts/proto/posts.proto`):

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

// Request and response message definitions...
```

## Service Usage

Other services interact with the Posts Service through the generated client:

```go
// Create a client
postClient := postProto.NewPostsService("posts", service.Client())

// Call methods
resp, err := postClient.List(context.Background(), &postProto.ListRequest{
    Page:  1,
    Limit: 10,
})
```