package grpc

import (
	"go-news-clean/internal/domain/usecase/tags"

	pb "go-news-clean/internal/proto"
)

type tagServiceServer struct {
	pb.UnimplementedTagServiceServer
	tagService *tags.TagService
}

func NewTagServiceServer(ts *tags.TagService) *tagServiceServer {
	return &tagServiceServer{
		tagService: ts,
	}
}
