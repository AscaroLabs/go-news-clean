package files

import "github.com/google/uuid"

type FileRepository interface {
	GetByNews(news_id uuid.UUID) ([]*File, error)
}
