package usecase

import (
	"context"

	database "github.com/braejan/platzi-go-rest-websockets/database/postgres"
	"github.com/braejan/platzi-go-rest-websockets/entities"
	"github.com/braejan/platzi-go-rest-websockets/repository"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(ctx *context.Context) (userUsecase *UserUsecase, err error) {
	postgresRepository, err := database.NewPostgresRepository(ctx)
	if err != nil {
		return
	}
	userUsecase = &UserUsecase{
		repository: postgresRepository,
	}
	return
}

func (usecase *UserUsecase) CreateUser(ctx *context.Context, email string) (err error) {
	user := &entities.User{
		ID:    uuid.NewString(),
		Email: email,
	}
	return usecase.repository.CreateUser(ctx, user)
}
