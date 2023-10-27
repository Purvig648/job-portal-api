package repository

import (
	"context"
	"errors"
	"job-application-api/internal/models"

	"github.com/rs/zerolog/log"
)
func(r *Repo) Viewjob(ctx context.Context, cid uint64)(models.Job,error){
	var jobData models.Job
	result := r.DB.Where("id = ?", cid).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the company")
	}
	return jobData, nil

}

func (r *Repo) ViewJobPostings(ctx context.Context) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find companires")
	}
	return jobDetails, nil

}

func (r *Repo) ViewJobs(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Where("cid = ?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, result.Error
	}
	return jobDetails, nil

}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Job) (models.Job, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create the job")
	}
	return jobData, nil

}

func (r *Repo) CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	result := r.DB.Create(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the company")
	}
	return companyData, nil
}
func (r *Repo) ViewCompany(ctx context.Context, cid uint64) (models.Company, error) {
	var companyData models.Company
	result := r.DB.Where("id = ?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not find the company")
	}
	return companyData, nil
}
func (r *Repo) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {
	var userDetails []models.Company
	result := r.DB.Find(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find companires")
	}
	return userDetails, nil
}
