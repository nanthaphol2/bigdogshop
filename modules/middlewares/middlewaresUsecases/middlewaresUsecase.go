package middlewaresUsecases

import (
	"github.com/nanthaphol2/bigdogshop/modules/middlewares"
	"github.com/nanthaphol2/bigdogshop/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
	FindOneUserById(userId string) (*middlewares.UserRole, error)
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (u *middlewaresUsecase) FindRole() ([]*middlewares.Role, error) {
	roles, err := u.middlewaresRepository.FindRole()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (u *middlewaresUsecase) FindOneUserById(userId string) (*middlewares.UserRole, error) {
	roles, err := u.middlewaresRepository.FindOneUserById(userId)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
