package app

import (
	accessRepository "go-news-clean/internal/adapters/postgres/access"
	filesRepository "go-news-clean/internal/adapters/postgres/files"
	newsRepository "go-news-clean/internal/adapters/postgres/news"
	tagsRepository "go-news-clean/internal/adapters/postgres/tags"
	"go-news-clean/internal/domain/usecase/access"
	"go-news-clean/internal/domain/usecase/files"
	"go-news-clean/internal/domain/usecase/news"
	"go-news-clean/internal/domain/usecase/tags"
	"go-news-clean/pkg/client"
	"go-news-clean/pkg/env"
	"log"
	"net"

	"go-news-clean/internal/gateways"
)

type App struct {
}

func (a *App) Run() error {
	ts := tags.NewTagService(tagsRepository.NewTagsRepository(client.ConnString))
	fs := files.NewFileService(filesRepository.NewFileRepository(client.ConnString))
	as := access.NewAccessService(accessRepository.NewAccessRepository(client.ConnString))
	ns := news.NewNewsService(newsRepository.NewNewsRepository(client.ConnString), ts, as, fs)
	grpc_server := gateways.NewGrpcServer(ns, ts)
	go func() {
		lis, err := net.Listen(
			"tcp",
			env.GrpcPort,
		)
		if err != nil {
			log.Fatalf("[gRPC] Failed to listen: %v", err)
		}
		log.Printf("[gRPC] Server listening at %v", lis.Addr())
		if err := grpc_server.Serve(lis); err != nil {
			log.Fatalf("[gRPC] Failed to serve: %v", err)
		}
	}()
	http_server := gateways.NewHTTPServer()
	go func() {
		err := http_server.Run()
		log.Printf("[HTTP] Server listening at %v", env.Port)
		if err != nil {
			log.Fatalf("[HTTP] Failed to run: %v", err)
		}
	}()
	return nil
}
