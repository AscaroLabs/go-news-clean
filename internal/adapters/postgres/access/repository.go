package access

import (
	"go-news-clean/internal/domain/entity/access"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// type AccessRepository interface {
// 	CanEditNews(uuid.UUID, AccessDTO) (bool, error)
// 	CanDeleteNews(uuid.UUID, AccessDTO) (bool, error)
// 	CanSeeNews(uuid.UUID, AccessDTO) (bool, error)
// }

type accessRepository struct {
	db *sqlx.DB
}

func NewAccessRepository(db *sqlx.DB) *accessRepository {
	return &accessRepository{
		db: db,
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
