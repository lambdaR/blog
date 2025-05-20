package handler

import (
	"context"
	"time"

	"github.com/google/uuid"
	pb "github.com/micro/blog/comments/proto"
)

type Handler struct {
	comments map[string]*pb.Comment
}

func New() *Handler {
	h := &Handler{
		comments: make(map[string]*pb.Comment),
	}

	// Add some example comments
	sampleComments := []*pb.Comment{
		{
			Id:        "1",
			Content:   "I really enjoyed reading this post",
			AuthorId:  "user1",
			PostId:    "1",
			CreatedAt: time.Now().Add(-24 * time.Hour).Unix(),
		},
		{
			Id:        "2",
			Content:   "I found this post very informative",
			AuthorId:  "user2",
			PostId:    "1",
			CreatedAt: time.Now().Add(-12 * time.Hour).Unix(),
		},
		{
			Id:        "3",
			Content:   "Great insights, thanks for sharing!",
			AuthorId:  "user1",
			PostId:    "2",
			CreatedAt: time.Now().Add(-2 * time.Hour).Unix(),
		},
	}

	// Store example comments
	for _, comment := range sampleComments {
		h.comments[comment.Id] = comment
	}

	return h
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	id := uuid.New().String()
	now := time.Now().Unix()

	comment := &pb.Comment{
		Id:        id,
		Content:   req.Content,
		AuthorId:  req.AuthorId,
		PostId:    req.PostId,
		CreatedAt: now,
	}

	h.comments[id] = comment

	rsp.Comment = comment
	return nil
}

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	comment, exists := h.comments[req.Id]
	if !exists {
		return nil // Return error in real implementation
	}

	rsp.Comment = comment
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	delete(h.comments, req.Id)
	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	var comments []*pb.Comment

	// Filter comments by post_id if provided
	for _, comment := range h.comments {
		if req.PostId == "" || comment.PostId == req.PostId {
			comments = append(comments, comment)
		}
	}

	rsp.Comments = comments
	rsp.Total = int32(len(comments))

	return nil
}
