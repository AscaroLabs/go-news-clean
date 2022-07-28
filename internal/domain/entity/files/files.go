package files

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	// id файла
	Id uuid.UUID
	// Название
	Name string
	// Расширение
	Ext string
	// base64 для файла
	Base64 string
	// Дата создания
	DateCreate time.Time
	// id пользователя, создавшего файл
	UserId uuid.UUID
}
