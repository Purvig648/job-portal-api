package repository

import (
	"context"
	"errors"
	"job-application-api/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) Viewjob(ctx context.Context, cid uint64) (models.Job, error) {
	var jobData models.Job
	result := r.DB.Where("id = ?", cid).First(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not find the job id")
	}
	return jobData, nil

}

func (r *Repo) ViewJobPostings(ctx context.Context) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find jobs")
	}
	return jobDetails, nil

}

func (r *Repo) ViewJobByCid(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.DB.Where("cid = ?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job for the cid")
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
