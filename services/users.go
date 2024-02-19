package services

import (
	"context"

	"github.com/anukuljoshi/goweb/models"
)

type UserService struct {
	store *models.Queries
}

func (us *UserService) ListUsers() ([]models.User, error) {
	return us.store.ListUsers(context.Background())
}

func (us *UserService) CreateUser(u models.CreateUserParams) (models.User, error) {
	return us.store.CreateUser(context.Background(), u)
}
