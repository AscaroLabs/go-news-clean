package grpc

import (
	"go-news-clean/internal/domain/usecase/news"
	"go-news-clean/internal/domain/usecase/tags"
	pb "go-news-clean/internal/proto"

	"google.golang.org/grpc"
)

func Register(s *grpc.Server, ns *news.NewsService, ts *tags.TagService) {
	pb.RegisterContentCheckServiceServer(s, &contentCheckServiceServer{})
	pb.RegisterNewsServiceServer(s, NewNewsServiceServer(ns))
	pb.RegisterTagServiceServer(s, NewTagServiceServer(ts))
}
