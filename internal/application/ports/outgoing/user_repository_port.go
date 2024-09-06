package outgoing

import (
	"auth-server/internal/domain"
	"auth-server/internal/errs"
)

type UserRepositoryPort interface {
	Create(user *domain.User) error
	GetByName(username string) (*domain.User, *errs.AppError)
	GetByEmail(email string) (*domain.User, error)
}
