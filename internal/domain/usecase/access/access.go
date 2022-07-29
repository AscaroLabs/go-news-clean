package access

import (
	"errors"
	"go-news-clean/internal/domain/entity/access"
	"log"

	"github.com/google/uuid"
)

var (
	ErrPermissionDenied   error = errors.New("permission denied")
	ErrAccessServiceError error = errors.New("problem with access providing")
)

type AccessService struct {
	accessRepository access.AccessRepository
}

func NewAccessService(accessRepository access.AccessRepository) *AccessService {
	return &AccessService{
		accessRepository: accessRepository,
	}
}

// CanCreateNews проверяет может ли пользователь создавать новости
func (as *AccessService) CanCreateNews(adto access.AccessDTO) error {
	switch adto.Role {
	case access.Admin:
		return nil
	case access.Employee:
		return nil
	case access.Dealer:
		return ErrPermissionDenied
	default:
		return ErrPermissionDenied
	}
}

// CanEditNews проверяет может ли пользователь редактировать новость
func (as *AccessService) CanEditNews(news_id uuid.UUID, adto access.AccessDTO) error {
	switch adto.Role {
	case access.Admin:
		return nil
	case access.Dealer:
		return ErrPermissionDenied
	case access.Employee:
		ok, err := as.accessRepository.CanEditNews(news_id, adto)
		if err != nil {
			log.Printf("[AccessService] Error: %v", err)
			return ErrAccessServiceError
		}
		if !ok {
			return ErrPermissionDenied
		}
		return nil
	default:
		return ErrPermissionDenied
	}
}

// CanDeleteNews проверят может ли пользователь удалить новость
func (as *AccessService) CanDeleteNews(news_id uuid.UUID, adto access.AccessDTO) error {
	switch adto.Role {
	case access.Admin:
		return nil
	case access.Dealer:
		return ErrPermissionDenied
	case access.Employee:
		ok, err := as.accessRepository.CanDeleteNews(news_id, adto)
		if err != nil {
			log.Printf("[AccessService] Error: %v", err)
			return ErrAccessServiceError
		}
		if !ok {
			return ErrPermissionDenied
		}
		return nil
	default:
		return ErrPermissionDenied
	}
}

func (as *AccessService) CanSeeNews(news_id uuid.UUID, adto access.AccessDTO) error {
	switch adto.Role {
	case access.Admin:
		return nil
	case access.Dealer:
		return ErrPermissionDenied
	case access.Employee:
		ok, err := as.accessRepository.CanSeeNews(news_id, adto)
		if err != nil {
			log.Printf("[AccessService] Error: %v", err)
			return ErrAccessServiceError
		}
		if !ok {
			return ErrPermissionDenied
		}
		return nil
	default:
		return ErrPermissionDenied
	}
}
