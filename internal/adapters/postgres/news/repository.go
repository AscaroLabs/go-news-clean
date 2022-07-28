package news

import (
	"go-news-clean/internal/domain/entity/news"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type newsRepository struct {
	db *sqlx.DB
}

func NewNewsRepository(db *sqlx.DB) *newsRepository {
	return &newsRepository{
		db: db,
	}
}

// news.News возвращаются без файлов и тегов
func (nr *newsRepository) GetNews(gdto news.GetDTO, user_id uuid.UUID) ([]*news.News, error) {
	return nil, nil
}

// news.News возвращается без файлов и тегов
func (nr *newsRepository) GetOne(id uuid.UUID) (*news.News, error) {
	return nil, nil
}

// news.News возвращается без файлов и тегов
func (nr *newsRepository) GetOneBySlug(objectSlug string) (*news.News, error) {
	return nil, nil
}

func (nr *newsRepository) Create(cdto news.CreateDTO) error {
	return nil
}

func (nr *newsRepository) Update(udto news.UpdateDTO) error {
	return nil
}

func (nr *newsRepository) Delete(ddto news.DeleteDTO) error {
	return nil
}
