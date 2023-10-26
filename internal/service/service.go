package service

import (
	"context"
	"errors"
	"job-application-api/internal/models"
	"job-application-api/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
}

type UserService interface {
	UserSignup(ctx context.Context, userData models.UserSignup) (models.User, error)
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be nil")
	}
	return &Service{
		UserRepo: userRepo,
	}, nil

}
