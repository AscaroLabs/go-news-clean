package files

import (
	"go-news-clean/internal/domain/entity/files"

	"github.com/google/uuid"
)

type FileService struct {
	fileRepository files.FileRepository
}

func (fs *FileService) GetByNews(news_id uuid.UUID) ([]*files.File, error) {
	return fs.fileRepository.GetByNews(news_id)
}
