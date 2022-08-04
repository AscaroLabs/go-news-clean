package access

import (
	"go-news-clean/internal/domain/entity/access"

	"github.com/google/uuid"
)

type accessRepository struct {
	connStr string
}

func NewAccessRepository(connStr string) *accessRepository {
	return &accessRepository{
		connStr: connStr,
	}
}

func (ar *accessRepository) CanEditNews(news_id uuid.UUID, adto access.AccessDTO) (bool, error) {
	return true, nil
}

func (ar *accessRepository) CanDeleteNews(news_id uuid.UUID, adto access.AccessDTO) (bool, error) {
	return true, nil
}

func (ar *accessRepository) CanSeeNews(news_id uuid.UUID, adto access.AccessDTO) (bool, error) {
	return true, nil
}
