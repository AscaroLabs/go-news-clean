package news

import (
	"errors"
	access_entity "go-news-clean/internal/domain/entity/access"
	files_entity "go-news-clean/internal/domain/entity/files"
	"go-news-clean/internal/domain/entity/news"
	tags_entity "go-news-clean/internal/domain/entity/tags"
	"go-news-clean/internal/domain/usecase/access"
	"go-news-clean/internal/domain/usecase/files"
	"go-news-clean/internal/domain/usecase/tags"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
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

	log.Printf("[NewsService] Get News request")

	news_list, err := ns.newsRepository.GetNews(gdto, adto.Id)
	if err != nil {

		log.Printf("[NewsService] Error: %v", err)

		if errors.Is(err, pgx.ErrNoRows) {
			news_list = make([]*news.News, 0)
		} else {
			return nil, err
		}
	}

	log.Printf("[NewsService] Get entities news list without tags & files")

	for _, news_element := range news_list {
		tags_list, err := ns.tagService.GetByNews(news_element.Id)
		if err != nil {

			log.Printf("[NewsService] Error: %v", err)

			if errors.Is(err, pgx.ErrNoRows) {
				tags_list = make([]*tags_entity.Tag, 0)
			} else {
				return nil, err
			}
		}
		news_element.Tags = tags_list
	}

	log.Printf("[NewsService] Get tags")

	for _, news_element := range news_list {
		files_list, err := ns.fileService.GetByNews(news_element.Id)
		if err != nil {

			log.Printf("[NewsService] Error: %v", err)

			if errors.Is(err, pgx.ErrNoRows) {
				files_list = make([]*files_entity.File, 0)
			} else {
				return nil, err
			}
		}
		news_element.FilesInfo = files_list
	}

	log.Printf("[NewsService] Get files")

	return news_list, nil

}

func (ns *NewsService) GetOne(news_id uuid.UUID, adto access_entity.AccessDTO) (*news.News, error) {
	// проверяем может ли пользователь посмотреть эту новость
	err := ns.accessService.CanSeeNews(news_id, adto)
	if err != nil {
		return nil, err
	}
	news_object, err := ns.newsRepository.GetOne(news_id)
	if err != nil {
		return nil, err
	}
	tags_list, _ := ns.tagService.GetByNews(news_id)
	files_list, _ := ns.fileService.GetByNews(news_id)
	news_object.Tags = tags_list
	news_object.FilesInfo = files_list
	return news_object, nil
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
