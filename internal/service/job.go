package service

import (
	"context"
	"job-application-api/internal/models"
)

func (s *Service) AddJobDetails(ctx context.Context, jobData models.NewJob) (models.ResponseJob, error) {

	createjobdetails := models.Job{
		Cid:             jobData.Cid,
		Jobname:         jobData.Jobname,
		Location:        jobData.Location,
		MinExperience:   jobData.MinExperience,
		MaxExperience:   jobData.MaxExperience,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,
		TechnologyStack: jobData.TechnologyStack,
		Qualifications:  jobData.Qualifications,
		Shift:           jobData.Shift,
		Jobtype:         jobData.Jobtype,
		Description:     jobData.Description,
	}
	job, err := s.UserRepo.CreateJob(ctx, createjobdetails)
	if err != nil {
		return models.ResponseJob{}, err
	}
	return job, nil
}
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
	jobData, err := s.UserRepo.ViewJobByCid(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
