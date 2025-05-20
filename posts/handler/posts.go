package handler

import (
	"context"

	pb "github.com/micro/blog/posts/proto"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

var posts = []*pb.Post{
	{
		Id:      "1",
		Title:   "Post 1",
		Content: "Content for post 1",
	},
	{
		Id:      "2",
		Title:   "Post 2",
		Content: "Content for post 2",
	},
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	post := &pb.Post{
		Content: req.Content,
	}

	posts = append(posts, post)

	res.Post = post

	return nil
}

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, res *pb.ReadResponse) error {
	for _, post := range posts {
		if post.Id == req.Id {
			res.Post = post
			return nil
		}
	}

	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest, res *pb.UpdateResponse) error {
	for _, post := range posts {
		if post.Id == req.Id {
			post.Title = req.Title
			post.Content = req.Content
			res.Post = post
			return nil
		}
	}

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	for i, post := range posts {
		if post.Id == req.Id {
			posts = append(posts[:i], posts[i+1:]...)
			return nil
		}
	}

	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	res.Posts = posts

	return nil
}
