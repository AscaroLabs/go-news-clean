package auth

import (
	"context"
	"go-news-clean/internal/domain/entity/access"

	"github.com/google/uuid"
)

func GetAccessDTOFromGRPCContext(context.Context) (access.AccessDTO, error) {
	// Заглушка
	return access.AccessDTO{
		Id:   uuid.New(),
		Role: access.Admin,
	}, nil
}
