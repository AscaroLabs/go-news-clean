package files

type FileDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Ext        string `json:"ext"`
	Bytes      []byte `json:"bytes"`
	DateCreate int64  `json:"date_create"`
	UserId     string `json:"user_id"`
	FileSize   int64  `json:"file_size"`
}

type File struct {
	// id файла
	Id string `json:"id,omitempty"`
	// Название
	Name string `json:"name,omitempty"`
	// Расширение
	Ext string `json:"ext,omitempty"`
	// base64 для файла
	Base64 string `json:"base64,omitempty"`
	// Дата создания
	DateCreate int64 `json:"dateCreate,omitempty"`
	// id пользователя, создавшего файл
	UserId string `json:"userId,omitempty"`
}
