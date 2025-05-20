package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v5"

	commentProto "github.com/micro/blog/comments/proto"
	postProto "github.com/micro/blog/posts/proto"
	userProto "github.com/micro/blog/users/proto"
)

func main() {

	service := micro.NewService(
		micro.Name("rest-api"),
	)
	service.Init()

	postClient := postProto.NewPostsService("posts", service.Client())
	commentClient := commentProto.NewCommentsService("comments", service.Client())
	userClient := userProto.NewUsersService("users", service.Client())

	log.Println("Starting REST API server on port 42096...")

	router := gin.Default()

	// === Posts endpoints ===
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

	router.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := postClient.Read(context.Background(), &postProto.ReadRequest{
			Id: id,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.POST("/posts", func(c *gin.Context) {
		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := postClient.Create(context.Background(), &postProto.CreateRequest{
			Title:   req.Title,
			Content: req.Content,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	})

	// === Comments endpoints ===
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

	router.POST("/comments", func(c *gin.Context) {
		var req struct {
			Content  string `json:"content"`
			AuthorId string `json:"author_id"`
			PostId   string `json:"post_id"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := commentClient.Create(context.Background(), &commentProto.CreateRequest{
			Content:  req.Content,
			AuthorId: req.AuthorId,
			PostId:   req.PostId,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	})

	// === Users endpoints ===
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

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := userClient.Read(context.Background(), &userProto.ReadRequest{
			Id: id,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	router.POST("/users", func(c *gin.Context) {
		var req struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := userClient.Create(context.Background(), &userProto.CreateRequest{
			Name:  req.Name,
			Email: req.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	})

	if err := router.Run(":42096"); err != nil {
		log.Fatal("Failed to run Gin server:", err)
	}
}
