package services

import (
	"auth-server/internal/adapters/dto"
	"auth-server/internal/application/ports/incoming"
	"auth-server/internal/errs"
	"fmt"

	"auth-server/internal/application/ports/outgoing"
	"auth-server/internal/domain"
)

type AuthService struct {
	userRepository outgoing.UserRepositoryPort
	tokenGenerator outgoing.TokenGeneratorPort
}

func NewAuthService(
	userRepository outgoing.UserRepositoryPort,
	// tokenGenerator outgoing.TokenGeneratorPort,
) incoming.AuthPort {
	return &AuthService{
		userRepository: userRepository,
		//tokenGenerator: tokenGenerator,
	}
}

//func (a *AuthService) Login(request dto.LoginRequestDTO) (*domain.AuthToken, error) {
//	user, err := a.userRepository.GetByEmail(request.Email)
//	if err != nil {
//		return &domain.AuthToken{}, err
//	}
//	if !user.ValidatePassword(request.Password) {
//		return &domain.AuthToken{}, errors.New("invalid credentials")
//	}
//	output, err := a.tokenGenerator.GenerateToken(user)
//	if err != nil {
//		return nil, errors.New(err.Error())
//	}
//	return output, nil
//}

func (a *AuthService) Register(input dto.CreateUserInputDTO) (dto.UserOutputDTO, errs.AppError) {
	user, err := domain.NewUser(
		input.Name,
		input.Email,
		input.Password,
	)
	fmt.Printf("user: %v", user)
	if err != nil {
		return dto.UserOutputDTO{}, *errs.NewValidationError(err.Error())
	}

	userExisted, _ := a.userRepository.GetByName(user.GetName())
	if userExisted != nil {
		return dto.UserOutputDTO{}, *errs.NewConflictError(fmt.Sprintf("name %s already exists", user.GetName()))
	}

	userExisted, _ = a.userRepository.GetByEmail(user.GetEmail())
	if userExisted != nil {
		return dto.UserOutputDTO{}, *errs.NewConflictError(fmt.Sprintf("email %s already exists", user.GetEmail()))
	}

	if err = a.userRepository.Create(user); err != nil {
		return dto.UserOutputDTO{}, *errs.NewUnexpectedError(fmt.Sprintf("error: %s ", err.Error()))
	}

	outputDTO := dto.UserOutputDTO{
		ID:     user.GetID(),
		Name:   user.GetName(),
		Email:  user.GetEmail(),
		Status: user.GetStatus(),
	}
	return outputDTO, errs.AppError{}
}
