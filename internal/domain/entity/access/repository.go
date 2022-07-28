package access

import "github.com/google/uuid"

type AccessRepository interface {
	CanEditNews(uuid.UUID, AccessDTO) (bool, error)
	CanDeleteNews(uuid.UUID, AccessDTO) (bool, error)
	CanSeeNews(uuid.UUID, AccessDTO) (bool, error)
}
