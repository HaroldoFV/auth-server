package outgoing

import "auth-server/internal/domain"

type UserRepositoryPort interface {
	Create(user *domain.User) error
	GetByName(username string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
