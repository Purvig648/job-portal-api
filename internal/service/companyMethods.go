package service

import (
	"context"
	"job-application-api/internal/models"
)

func (s *Service) ViewJobDetailsById(ctx context.Context, cid uint64) (models.Job, error) {
	jobData, err := s.UserRepo.Viewjob(ctx, cid)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil

}

func (s *Service) ViewAllJobPostings(ctx context.Context) ([]models.Job, error) {
	jobData, err := s.UserRepo.ViewJobPostings(ctx)
	if err != nil {
		return nil, err
	}
	return jobData, nil

}

func (s *Service) ViewJobDetails(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.ViewJobs(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

func (s *Service) AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error) {
	companyData, err := s.UserRepo.CreateCompany(ctx, companyData)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error) {
	companyData, err := s.UserRepo.ViewCompany(ctx, cid)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {
	companyDetails, err := s.UserRepo.ViewAllCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return companyDetails, nil
}
func (s *Service) AddJobDetails(ctx context.Context, jobData models.Job) (models.Job, error) {
	jobData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil
}
