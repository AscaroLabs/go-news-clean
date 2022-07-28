package news

import (
	"go-news-clean/internal/domain/entity/news"

	"github.com/google/uuid"
)

type NewsService struct {
	// adapters.DataTransfer
	// cfg *config.Config
	newsRepository news.NewsRepository
}

func NewNewsService(newsRepository news.NewsRepository) *NewsService {
	return &NewsService{
		newsRepository: newsRepository,
	}
}

func (ns *NewsService) GetNews(gdto news.GetDTO) ([]*news.News, error) {
	// вопрос авторизации? (есть ли доступ к черновикам)
	return ns.newsRepository.GetNews(gdto)
}

func (ns *NewsService) GetOne(id uuid.UUID) (*news.News, error) {
	// вопрос авторизации? (есть ли доступ к черновикам)
	return ns.newsRepository.GetOne(id)
}
