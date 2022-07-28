package tags

// type TagsRepository interface {
// 	GetAll() ([]Tag, error)
// 	GetByNews(news_id uuid.UUID) ([]*Tag, error)
// }

import (
	"go-news-clean/internal/domain/entity/tags"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type tagsRepository struct {
	db *sqlx.DB
}

func NewTagsRepository(db *sqlx.DB) *tagsRepository {
	return &tagsRepository{
		db: db,
	}
}

func (tr *tagsRepository) GetAll() ([]tags.Tag, error) {
	return nil, nil
}

func (tr *tagsRepository) GetByNews(news_id uuid.UUID) ([]*tags.Tag, error) {
	return nil, nil
}
