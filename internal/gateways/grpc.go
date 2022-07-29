package gateways

import (
	handlers "go-news-clean/internal/adapters/handlers/v1/grpc"
	"go-news-clean/internal/domain/usecase/news"
	"go-news-clean/internal/domain/usecase/tags"

	"google.golang.org/grpc"
)

func NewGrpcServer(ns *news.NewsService, ts *tags.TagService) *grpc.Server {
	s := grpc.NewServer()
	handlers.Register(s, ns, ts)
	return s
}
