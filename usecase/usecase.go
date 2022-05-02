package usecase

import (
	"context"

	useruc "github.com/braejan/platzi-go-rest-websockets/usecase/user"
)

type UseCases struct {
	User *useruc.UserUsecase
}

func NewUsecases(ctx *context.Context) (usecases *UseCases, err error) {
	userUsecase, err := useruc.NewUserUsecase(ctx)
	if err != nil {
		return
	}
	usecases = &UseCases{
		User: userUsecase,
	}
	return
}
