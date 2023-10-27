package service

import (
	"context"
	"errors"
	"job-application-api/internal/auth"
	"job-application-api/internal/models"
	"job-application-api/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	a        *auth.Auth
}

type UserService interface {
	UserSignup(ctx context.Context, userData models.UserSignup) (models.User, error)
	UserLogin(ctx context.Context, userData models.UserLogin) (string, error)
	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	AddJobDetails(ctx context.Context, jobData models.Job) (models.Job, error)
	ViewJobDetails(ctx context.Context, cid uint64) ([]models.Job, error)
	ViewAllJobPostings(ctx context.Context) ([]models.Job, error)
	ViewJobDetailsById(ctx context.Context, cid uint64) (models.Job, error)
}

func NewService(userRepo repository.UserRepo, a *auth.Auth) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be nil")
	}
	return &Service{
		UserRepo: userRepo,
		a:        a,
	}, nil

}
