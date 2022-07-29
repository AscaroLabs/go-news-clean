package tags

import (
	"go-news-clean/internal/domain/entity/tags"

	"github.com/google/uuid"
)

type TagService struct {
	tagRepository tags.TagsRepository
}

func NewTagService(tagRepository tags.TagsRepository) *TagService {
	return &TagService{
		tagRepository: tagRepository,
	}
}

func (ts *TagService) GetAll() ([]*tags.Tag, error) {
	tags_list, err := ts.tagRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return tags_list, nil
}

func (ts *TagService) GetByNews(news_id uuid.UUID) ([]*tags.Tag, error) {
	tags_list, err := ts.tagRepository.GetByNews(news_id)
	if err != nil {
		return nil, err
	}
	return tags_list, nil
}
