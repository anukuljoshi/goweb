package services

import (
	"github.com/anukuljoshi/goweb/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	User *UserService
}

func NewService(dbPool *pgxpool.Pool) (*Service, error) {
	return &Service{
		User: &UserService{
			store: models.New(dbPool),
		},
	}, nil
}
