package repository

import (
	"context"
	"errors"

	"github.com/braejan/platzi-go-rest-websockets/entities"
)

type UserRepository interface {
	CreateUser(ctx *context.Context, user *entities.User) (err error)
	GetUserByID(ctx *context.Context, userID string) (user *entities.User, err error)
	DeleteUserByID(ctx *context.Context, userID string) (err error)
	UpdateUser(ctx *context.Context, user *entities.User) (err error)
}

var userRepositoryImpl UserRepository

func SetRepository(repository UserRepository) {
	userRepositoryImpl = repository
}

func exists() (exists bool) {
	exists = userRepositoryImpl != nil
	return
}

func errorEmptyRepository() (err error) {
	if !exists() {
		err = errors.New("user repository has not be initialized")
	}
	return
}

func CreateUser(ctx *context.Context, user *entities.User) (err error) {
	err = errorEmptyRepository()
	if err != nil {
		return
	}
	return userRepositoryImpl.CreateUser(ctx, user)
}
func GetUserByID(ctx *context.Context, userID string) (user *entities.User, err error) {
	err = errorEmptyRepository()
	if err != nil {
		return
	}
	return userRepositoryImpl.GetUserByID(ctx, userID)
}
func DeleteUserByID(ctx *context.Context, userID string) (err error) {
	err = errorEmptyRepository()
	if err != nil {
		return
	}
	return userRepositoryImpl.DeleteUserByID(ctx, userID)
}
func UpdateUser(ctx *context.Context, user *entities.User) (err error) {
	err = errorEmptyRepository()
	if err != nil {
		return
	}
	return userRepositoryImpl.UpdateUser(ctx, user)
}
