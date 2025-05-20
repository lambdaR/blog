package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/micro/blog/users/proto"
)

type Handler struct {
	users map[string]*pb.User
}

func New() *Handler {

	h := &Handler{
		users: make(map[string]*pb.User),
	}

	// Sample users
	var sampleUsers = []*pb.User{
		{
			Id:   "user1",
			Name: "Morena",
		},
		{
			Id:   "user2",
			Name: "Simon",
		},
		{
			Id:   "user3",
			Name: "Carmelo",
		},
	}

	// Store sample users
	for _, user := range sampleUsers {
		h.users[user.Id] = user
	}

	return h
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	id := uuid.New().String()

	user := &pb.User{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	h.users[id] = user

	rsp.User = user
	return nil
}

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	user, exists := h.users[req.Id]
	if !exists {
		return nil
	}

	rsp.User = user
	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	user, exists := h.users[req.Id]
	if !exists {
		return nil
	}

	user.Name = req.Name
	user.Email = req.Email

	rsp.User = user
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	delete(h.users, req.Id)
	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	var users []*pb.User
	for _, user := range h.users {
		users = append(users, user)
	}

	rsp.Users = users
	rsp.Total = int32(len(users))

	return nil
}
