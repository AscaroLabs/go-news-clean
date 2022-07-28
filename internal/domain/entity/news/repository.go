package news

import (
	"github.com/google/uuid"
)

type NewsRepository interface {
	GetNews(GetDTO) ([]*News, error)
	GetOne(id uuid.UUID) (*News, error)
	GetOneBySlug(objectSlug string) (*News, error)
	Create(cdto CreateDTO) error
	Update(udto UpdateDTO) error
	Delete(ddto DeleteDTO) error
}
