package grpc

import (
	"context"
	"log"

	"go-news-clean/internal/domain/entity/files"
	entity "go-news-clean/internal/domain/entity/news"
	"go-news-clean/internal/domain/usecase/news"
	"go-news-clean/pkg/auth"

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
	newsService *news.NewsService
}

func NewNewsServiceServer(ns *news.NewsService) *newsServiceServer {
	return &newsServiceServer{
		newsService: ns,
	}
}

func (ns *newsServiceServer) GetNews(ctx context.Context, r *pb.NewsRequestParams) (*pb.NewsList, error) {

	log.Printf("[REQ] Get News")

	gdto, err := r.ToGetDTO()
	if err != nil {

		log.Printf("[REQ] Error: %v", err)

		return nil, err
	}

	log.Printf("[REQ] GetDTO created")

	adto, err := auth.GetAccessDTOFromGRPCContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[REQ] AccessDTO created")

	news_list, err := ns.newsService.GetNews(gdto, adto)
	if err != nil {
		return nil, err
	}

	log.Printf("[REQ] Entities news list created")

	return ConvertEntityNewsListToPBNewsList(news_list), nil

}

func ConvertEntityNewsListToPBNewsList(nl []*entity.News) *pb.NewsList {
	news_list := make([]*pb.NewsObject, 0)
	for _, n := range nl {
		news_list = append(news_list, ConvertEntityNewsToPBNewsObject(n))
	}
	return &pb.NewsList{
		News:  news_list,
		Total: int32(len(news_list)),
	}
}

func ConvertEntityNewsToPBNewsObject(n *entity.News) *pb.NewsObject {
	return &pb.NewsObject{
		Id:         n.Id.String(),
		Title:      n.Title,
		Author:     n.Author,
		Active:     n.Active,
		ActiveFrom: n.ActiveFrom.Unix(),
		Text:       n.Text,
		TextJson:   n.TextJson,
		UserId:     n.UserId.String(),
		Tags:       ConvertEntityTagListToPBTagList(n.Tags),
		FilesInfo:  ConvertEntityFileListToPBFileList(n.FilesInfo),
	}
}

func ConvertEntityFileListToPBFileList(fl []*files.File) []*pb.File {
	files_list := make([]*pb.File, 0)
	for _, f := range fl {
		files_list = append(files_list, ConvertEntityFileToPBFile(f))
	}
	return files_list
}

func ConvertEntityFileToPBFile(f *files.File) *pb.File {
	return &pb.File{
		Id:         f.Id.String(),
		Name:       f.Name,
		Ext:        f.Ext,
		Base64:     f.Base64,
		DateCreate: f.DateCreate.Unix(),
		UserId:     f.UserId.String(),
	}
}
