package files

import (
	"go-news-clean/internal/domain/entity/files"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// type FileRepository interface {
// 	GetByNews(news_id uuid.UUID) ([]*File, error)
// }

type fileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *fileRepository {
	return &fileRepository{
		db: db,
	}
}

func (fr *fileRepository) GetByNews(news_id uuid.UUID) ([]*files.File, error) {
	return nil, nil
}
