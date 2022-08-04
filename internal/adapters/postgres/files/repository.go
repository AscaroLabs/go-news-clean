package files

import (
	"context"
	"go-news-clean/internal/domain/entity/files"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type fileRepository struct {
	connStr string
}

func NewFileRepository(connStr string) *fileRepository {
	return &fileRepository{
		connStr: connStr,
	}
}

func (fr *fileRepository) GetByNews(news_id uuid.UUID) ([]*files.File, error) {
	pool, err := pgxpool.Connect(context.Background(), fr.connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		`select id, name, ext, base64, dateCreate, userId from files where newsId=$1`,
		news_id.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files_list := make([]*files.File, 0)

	for rows.Next() {
		var file files.File
		err := rows.Scan(&file.Id, &file.Name, &file.Ext, &file.Base64, &file.DateCreate, &file.UserId)
		if err != nil {
			return nil, err
		}
		files_list = append(files_list, &file)
	}

	return files_list, nil
}
