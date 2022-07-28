package news

import (
	"github.com/google/uuid"
)

type NewsRepository interface {
	// GetNews получает новости по запросу и id пользователя
	GetNews(GetDTO, uuid.UUID) ([]*News, error)
	GetOne(id uuid.UUID) (*News, error)
	GetOneBySlug(objectSlug string) (*News, error)
	Create(cdto CreateDTO) error
	Update(udto UpdateDTO) error
	Delete(ddto DeleteDTO) error
}
