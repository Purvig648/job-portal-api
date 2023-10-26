package service

import (
	"context"
	"job-application-api/internal/models"
	"job-application-api/internal/pkg"
)

func (s *Service) UserSignup(ctx context.Context, userData models.UserSignup) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		PasswordHash: hashedPass,
	}
	userDetails, err = s.UserRepo.CreateUser(userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil

}
