package tags

import "github.com/google/uuid"

type TagsRepository interface {
	GetAll() ([]*Tag, error)
	GetByNews(news_id uuid.UUID) ([]*Tag, error)
}
