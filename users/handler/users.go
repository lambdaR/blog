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

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	rec, err := userStore.Read("user-" + req.Id)
	if err == nil && len(rec) > 0 {
		var user pb.User
		if err := json.Unmarshal(rec[0].Value, &user); err == nil {
			rsp.User = &user
			return nil
		}
	}
	// Not found
	rsp.User = nil
	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	rec, err := userStore.Read("user-" + req.Id)
	if err != nil || len(rec) == 0 {
		rsp.User = nil
		return nil
	}
	var user pb.User
	if err := json.Unmarshal(rec[0].Value, &user); err != nil {
		rsp.User = nil
		return nil
	}
	user.Name = req.Name
	user.Email = req.Email
	b, err := json.Marshal(&user)
	if err == nil {
		_ = userStore.Write(&store.Record{Key: "user-" + user.Id, Value: b})
	}
	rsp.User = &user
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	_ = userStore.Delete("user-" + req.Id)
	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	rec, err := userStore.Read("user-", store.ReadPrefix())
	if err == nil && len(rec) > 0 {
		var loadedUsers []*pb.User
		for _, r := range rec {
			var u pb.User
			if err := json.Unmarshal(r.Value, &u); err == nil {
				loadedUsers = append(loadedUsers, &u)
			}
		}
		rsp.Users = loadedUsers
		rsp.Total = int32(len(loadedUsers))
		return nil
	}
	rsp.Users = nil
	rsp.Total = 0
	return nil
}
