package news

import "github.com/google/uuid"

type GetDTO struct {
	Id uuid.UUID
}

type CreateDTO struct {
}

type UpdateDTO struct {
	Id uuid.UUID
}

type DeleteDTO struct {
	Id uuid.UUID
}
