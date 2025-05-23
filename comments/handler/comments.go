package handler

import (
	"context"
	"encoding/json"
	"sort"
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

	// Extract first URL and fetch link preview
	url := extractFirstURL(req.Content)
	if url != "" {
		title, desc, image, err := fetchLinkPreview(url)
		if err == nil {
			comment.LinkPreview = &pb.LinkPreview{
				Url:         url,
				Title:       title,
				Description: desc,
				Image:       image,
			}
		}
	}

	rsp.Comment = comment

	// Save to store
	b, err := json.Marshal(comment)
	if err == nil {
		_ = commentStore.Write(&store.Record{Key: "comment-" + comment.Id, Value: b})
	}

	return nil
}

func (h *Handler) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	rec, err := commentStore.Read("comment-" + req.Id)
	if err == nil && len(rec) > 0 {
		var comment pb.Comment
		if err := json.Unmarshal(rec[0].Value, &comment); err == nil {
			rsp.Comment = &comment
			return nil
		}
	}
	rsp.Comment = nil
	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	rec, err := commentStore.Read("comment-" + req.Id)
	if err != nil || len(rec) == 0 {
		rsp.Comment = nil
		return nil
	}
	var comment pb.Comment
	if err := json.Unmarshal(rec[0].Value, &comment); err != nil {
		rsp.Comment = nil
		return nil
	}
	comment.Content = req.Content
	comment.AuthorId = req.UserId // UpdateRequest now uses user_id
	comment.PostId = req.PostId
	comment.CreatedAt = comment.CreatedAt // preserve original timestamp
	b, err := json.Marshal(&comment)
	if err == nil {
		_ = commentStore.Write(&store.Record{Key: "comment-" + comment.Id, Value: b})
	}
	rsp.Comment = &comment
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	_ = commentStore.Delete("comment-" + req.Id)
	return nil
}

func (h *Handler) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	rec, err := commentStore.Read("comment-", store.ReadPrefix())
	if err == nil && len(rec) > 0 {
		var loadedComments []*pb.Comment
		for _, r := range rec {
			var c pb.Comment
			if err := json.Unmarshal(r.Value, &c); err == nil {
				if req.PostId == "" || c.PostId == req.PostId {
					loadedComments = append(loadedComments, &c)
				}
			}
		}
		// Sort comments by CreatedAt descending
		if len(loadedComments) > 1 {
			sort.Slice(loadedComments, func(i, j int) bool {
				return loadedComments[i].CreatedAt > loadedComments[j].CreatedAt
			})
		}
		rsp.Comments = loadedComments
		rsp.Total = int32(len(loadedComments))
		return nil
	}
	rsp.Comments = nil
	rsp.Total = 0
	return nil
}
