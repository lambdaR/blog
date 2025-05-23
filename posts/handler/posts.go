package handler

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/google/uuid"
	pb "github.com/micro/blog/posts/proto"
	"go-micro.dev/v5/store"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

var postStore = store.DefaultStore

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	post := &pb.Post{
		Id:         uuid.New().String(),
		Title:      req.Title,
		Content:    req.Content,
		AuthorId:   req.AuthorId,
		AuthorName: req.AuthorName,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	// Extract first URL and fetch link preview
	url := extractFirstURL(req.Content)
	if url != "" {
		title, desc, image, err := fetchLinkPreview(url)
		if err == nil {
			post.LinkPreview = &pb.LinkPreview{
				Url:         url,
				Title:       title,
				Description: desc,
				Image:       image,
			}
		}
	}

	res.Post = post

	// Save to store
	b, err := json.Marshal(post)
	if err == nil {
		_ = postStore.Write(&store.Record{Key: "post-" + post.Id, Value: b})
	}

	return nil
}

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, res *pb.ReadResponse) error {
	rec, err := postStore.Read("post-" + req.Id)
	if err == nil && len(rec) > 0 {
		var post pb.Post
		if err := json.Unmarshal(rec[0].Value, &post); err == nil {
			res.Post = &post
			return nil
		}
	}
	res.Post = nil
	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest, res *pb.UpdateResponse) error {
	rec, err := postStore.Read("post-" + req.Id)
	if err != nil || len(rec) == 0 {
		res.Post = nil
		return nil
	}
	var post pb.Post
	if err := json.Unmarshal(rec[0].Value, &post); err != nil {
		res.Post = nil
		return nil
	}
	post.Title = req.Title
	post.Content = req.Content
	b, err := json.Marshal(&post)
	if err == nil {
		_ = postStore.Write(&store.Record{Key: "post-" + post.Id, Value: b})
	}
	res.Post = &post
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, res *pb.DeleteResponse) error {
	_ = postStore.Delete("post-" + req.Id)
	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	rec, err := postStore.Read("post-", store.ReadPrefix())
	if err == nil && len(rec) > 0 {
		var loadedPosts []*pb.Post
		for _, r := range rec {
			var p pb.Post
			if err := json.Unmarshal(r.Value, &p); err == nil {
				loadedPosts = append(loadedPosts, &p)
			}
		}
		// Sort posts by CreatedAt descending (reverse chronological)
		if len(loadedPosts) > 1 {
			sort.Slice(loadedPosts, func(i, j int) bool {
				return loadedPosts[i].CreatedAt > loadedPosts[j].CreatedAt
			})
		}
		res.Posts = loadedPosts
		res.Total = int32(len(loadedPosts))
		return nil
	}
	res.Posts = nil
	res.Total = 0
	return nil
}
