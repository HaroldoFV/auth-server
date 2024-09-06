package incoming

import (
	"auth-server/internal/adapters/dto"
	"auth-server/internal/errs"
)

type AuthPort interface {
	//Login(request dto.LoginRequestDTO) (domain.AuthToken, error)
	Register(input dto.CreateUserInputDTO) (dto.UserOutputDTO, errs.AppError)
}
