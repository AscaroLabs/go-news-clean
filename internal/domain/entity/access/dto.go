package access

import "github.com/google/uuid"

type AccessDTO struct {
	Id   uuid.UUID
	Role Role
}
