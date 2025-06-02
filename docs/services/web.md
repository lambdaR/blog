# Web Service

The Web Service acts as an API gateway for the Micro Blog system. It provides a REST API that integrates all other microservices and serves a static web UI.

## Service Overview

The Web Service provides the following functionality:

- REST API endpoints for all features
- Service orchestration (calling appropriate microservices)
- Authentication and session management
- Static file serving for the web UI

## Implementation

### Main Service File

The main service file (`web/main.go`) initializes the service, sets up the Gin router, and defines all API endpoints:

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "go-micro.dev/v5"
    "golang.org/x/crypto/bcrypt"

    commentProto "github.com/micro/blog/comments/proto"
    postProto "github.com/micro/blog/posts/proto"
    userProto "github.com/micro/blog/users/proto"
)

func main() {
    service := micro.NewService(
        micro.Name("rest-api"),
    )
    service.Init()

    // Create clients for all microservices
    postClient := postProto.NewPostsService("posts", service.Client())
    commentClient := commentProto.NewCommentsService("comments", service.Client())
    userClient := userProto.NewUsersService("users", service.Client())

    // Set up Gin router and session store
    router := gin.Default()
    sessionStore := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("session", sessionStore))

    // Define API endpoints...

    // Start the server
    if err := router.Run(":42096"); err != nil {
        log.Fatal("Failed to run Gin server:", err)
    }
}
```

### API Endpoints

The Web Service defines REST API endpoints for all features:

```go
// Posts endpoints
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

// Comments endpoints
router.GET("/comments", func(c *gin.Context) {
    postID := c.Query("post_id")
    resp, err := commentClient.List(context.Background(), &commentProto.ListRequest{
        PostId: postID,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
})

// Users endpoints
router.GET("/users", func(c *gin.Context) {
    resp, err := userClient.List(context.Background(), &userProto.ListRequest{
        Page:  1,
        Limit: 10,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
})

// Authentication endpoints
router.POST("/signup", func(c *gin.Context) {
    // Implementation...
})

router.POST("/login", func(c *gin.Context) {
    // Implementation...
})

router.POST("/logout", func(c *gin.Context) {
    // Implementation...
})
```

### Authentication and Session Management

The Web Service handles authentication and session management:

```go
// Password hashing
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Session middleware
router.Use(func(c *gin.Context) {
    sess := sessions.Default(c)
    userID := sess.Get("user_id")
    userName := sess.Get("user_name")
    if userID != nil && userName != nil {
        c.Set("user_id", userID)
        c.Set("user_name", userName)
    }
    c.Next()
})

// Login endpoint
router.POST("/login", func(c *gin.Context) {
    // Verify credentials
    // ...
    
    // Set session
    sess := sessions.Default(c)
    sess.Set("user_id", found.Id)
    sess.Set("user_name", found.Name)
    sess.Save()
    
    c.JSON(http.StatusOK, gin.H{"user": found})
})
```

### Static File Serving

The Web Service serves static files for the web UI:

```go
// Serve static files
router.Static("/static", "./static")

// Serve static HTML pages
router.GET("/", func(c *gin.Context) {
    c.File("./static/index.html")
})

router.GET("/login.html", func(c *gin.Context) {
    c.File("./static/login.html")
})

router.GET("/signup.html", func(c *gin.Context) {
    c.File("./static/signup.html")
})

router.GET("/@:username", func(c *gin.Context) {
    c.File("./static/profile.html")
})
```

## API Documentation

The Web Service exposes a REST API with the following endpoints:

### Posts

- `GET /posts`: List all posts
- `GET /posts/:id`: Get a post by ID
- `POST /posts`: Create a new post
- `GET /posts/by-tag/:tag`: Get posts with a specific tag

### Comments

- `GET /comments`: List all comments (optionally filter by `post_id` query param)
- `POST /comments`: Add a comment

### Users

- `GET /users`: List all users
- `GET /users/:id`: Get a user by ID
- `POST /users`: Create a new user
- `GET /users/me`: Get the current session user info

### Authentication

- `POST /signup`: Register a new user (and log in)
- `POST /login`: Log in as a user
- `POST /logout`: Log out the current user

### Tags

- `POST /posts/:id/tags`: Add a tag to a post
- `DELETE /posts/:id/tags/:tag`: Remove a tag from a post
- `GET /tags`: Get all available tags
- `GET /tags?post_id=:id`: Get tags for a specific post

## Web UI

The Web Service serves a static web UI from the `web/static/` directory:

- `index.html`: Main feed, create posts, view posts and comments
- `login.html`: User login page
- `signup.html`: User registration page
- `profile.html`: User profile, posts, and comments