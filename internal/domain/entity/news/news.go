package news

import (
	"go-news-clean/internal/domain/entity/files"
	"go-news-clean/internal/domain/entity/tags"
	"time"

	"github.com/google/uuid"
)

type News struct {
	// id повости
	Id uuid.UUID
	// Название
	Title string
	// Автор
	Author string
	// Активность (true - активна, false - черновик)
	Active bool
	// Дата начала активности
	ActiveFrom time.Time
	// Текстовое описание
	Text string
	// Текствое описание (для визуального редактора)
	TextJson string
	// Идентификатор пользователя
	UserId uuid.UUID
	// Список тегов
	Tags []*tags.Tag
	// Список прикрепленных файлов
	FilesInfo []*files.File
	// Важное
	IsImportant bool
}
