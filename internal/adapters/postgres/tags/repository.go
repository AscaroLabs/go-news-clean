package tags

// type TagsRepository interface {
// 	GetAll() ([]Tag, error)
// 	GetByNews(news_id uuid.UUID) ([]*Tag, error)
// }

import (
	"context"
	"go-news-clean/internal/domain/entity/tags"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tagsRepository struct {
	connStr string
}

func NewTagsRepository(connStr string) *tagsRepository {
	return &tagsRepository{
		connStr: connStr,
	}
}

func (tr *tagsRepository) GetAll() ([]*tags.Tag, error) {

	pool, err := pgxpool.Connect(context.Background(), tr.connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		`select id, name from tags`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tag_list := make([]*tags.Tag, 0)

	for rows.Next() {
		var tag tags.Tag
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}
		tag_list = append(tag_list, &tag)
	}

	return tag_list, nil
}

func (tr *tagsRepository) GetByNews(news_id uuid.UUID) ([]*tags.Tag, error) {
	pool, err := pgxpool.Connect(context.Background(), tr.connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		`select tags.id, tags.name from tags join news_tags on tags.id=news_tags.tag_id where news_tags.news_id=$1`,
		news_id.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tag_list := make([]*tags.Tag, 0)

	for rows.Next() {
		var tag tags.Tag
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}
		tag_list = append(tag_list, &tag)
	}

	return tag_list, nil
}
