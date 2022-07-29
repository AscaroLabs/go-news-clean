package news

import (
	access_entity "go-news-clean/internal/domain/entity/access"
	"go-news-clean/internal/domain/entity/news"
	"go-news-clean/internal/domain/usecase/access"
	"go-news-clean/internal/domain/usecase/files"
	"go-news-clean/internal/domain/usecase/tags"

	"github.com/google/uuid"
)

type NewsService struct {
	newsRepository news.NewsRepository
	tagService     *tags.TagService
	fileService    *files.FileService
	accessService  *access.AccessService
}

func NewNewsService(nr news.NewsRepository, ts *tags.TagService, as *access.AccessService, fs *files.FileService) *NewsService {
	return &NewsService{
		newsRepository: nr,
		tagService:     ts,
		accessService:  as,
		fileService:    fs,
	}
}

func (ns *NewsService) GetNews(gdto news.GetDTO, adto access_entity.AccessDTO) ([]*news.News, error) {
	news_list, err := ns.newsRepository.GetNews(gdto, adto.Id)
	if err != nil {
		return nil, err
	}

	for _, news_element := range news_list {
		tags_list, _ := ns.tagService.GetByNews(news_element.Id)
		news_element.Tags = tags_list
	}

	for _, news_element := range news_list {
		files_list, _ := ns.fileService.GetByNews(news_element.Id)
		news_element.FilesInfo = files_list
	}

	return news_list, nil

}

func (ns *NewsService) GetOne(news_id uuid.UUID, adto access_entity.AccessDTO) (*news.News, error) {
	// проверяем может ли пользователь посмотреть эту новость
	err := ns.accessService.CanSeeNews(news_id, adto)
	if err != nil {
		return nil, err
	}
	return ns.newsRepository.GetOne(news_id)
}

func (ns *NewsService) Create(cdto news.CreateDTO, adto access_entity.AccessDTO) error {
	// проверяем может ли пользователь создавать новости
	err := ns.accessService.CanCreateNews(adto)
	if err != nil {
		return err
	}
	err = ns.newsRepository.Create(cdto)
	if err != nil {
		return err
	}
	return nil
}

func (ns *NewsService) Update(udto news.UpdateDTO, adto access_entity.AccessDTO) error {
	// проверяем может ли пользователь изменить данную новость
	err := ns.accessService.CanEditNews(udto.Id, adto)
	if err != nil {
		return err
	}
	err = ns.newsRepository.Update(udto)
	if err != nil {
		return err
	}
	return nil
}

func (ns *NewsService) Delete(ddto news.DeleteDTO, adto access_entity.AccessDTO) error {
	// проверяем может ли пользователь удалить данную новость
	err := ns.accessService.CanDeleteNews(ddto.Id, adto)
	if err != nil {
		return err
	}
	err = ns.newsRepository.Delete(ddto)
	if err != nil {
		return err
	}
	return nil
}
