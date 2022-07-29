package grpc

import (
	"context"

	"go-news-clean/internal/domain/usecase/news"

	pb "go-news-clean/internal/proto"
)

// type NewsServiceServer interface {
// 	// Получение списка новостей
// 	GetNews(context.Context, *NewsRequestParams) (*NewsList, error)
// 	// Получение детальной информации по новости
// 	GetOne(context.Context, *ObjectId) (*NewsObject, error)
// 	// Получение детальной информации по новости для отображения при переходе по письму
// 	GetOneBySlug(context.Context, *ObjectSlug) (*NewsObject, error)
// 	// Создание новости
// 	Create(context.Context, *RequestNewsObject) (*BaseResponse, error)
// 	// Обновление новости
// 	Update(context.Context, *RequestNewsObject) (*BaseResponse, error)
// 	// Удаление новости
// 	Delete(context.Context, *ObjectId) (*BaseResponse, error)
// 	//Получить ссылку на файл для скачивания
// 	GetFileLink(context.Context, *FileId) (*FileLink, error)
// 	mustEmbedUnimplementedNewsServiceServer()
// }

type newsServiceServer struct {
	pb.UnimplementedNewsServiceServer
	newsService news.NewsService
}

func NewNewsServiceServer(ns news.NewsService) *newsServiceServer {
	return &newsServiceServer{
		newsService: ns,
	}
}

func (ns *newsServiceServer) GetNews(context.Context, *pb.NewsRequestParams) (*pb.NewsList, error) {
	return nil, nil
}
