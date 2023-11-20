package service

import (
	"context"
	"errors"
	"job-application-api/internal/auth"
	"job-application-api/internal/caching"
	"job-application-api/internal/models"
	"job-application-api/internal/repository"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
	"go.uber.org/mock/gomock"
)

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.NewJob
	}
	tests := []struct {
		name             string
		args             args
		want             models.ResponseJob
		wantErr          bool
		mockRepoResponse func() (models.ResponseJob, error)
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Cid:             1,
					Jobname:         "Software Engineer",
					MinNoticePeriod: "30",
					MaxNoticePeriod: "60",
					Description:     "This is a software engineer job description.",
					MinExperience:   "2",
					MaxExperience:   "4",
					Location:        []uint{uint(1), uint(2)},
					TechnologyStack: []uint{uint(1), uint(2)},
					Qualifications:  []uint{uint(1), uint(2)},
					Shift:           []uint{uint(1), uint(2)},
					Jobtype:         "Full-time",
				},
			},
			want: models.ResponseJob{
				Id: 1,
			},
			wantErr: false,
			mockRepoResponse: func() (models.ResponseJob, error) {
				return models.ResponseJob{
					Id: 1,
				}, nil
			},
		},
		{
			name: "Failure",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					Cid:             1,
					Jobname:         "Software Engineer",
					MinNoticePeriod: "30",
					MaxNoticePeriod: "60",
					Description:     "This is a software engineer job description.",
					MinExperience:   "2",
					MaxExperience:   "4",
					Location:        []uint{uint(1), uint(2)},
					TechnologyStack: []uint{uint(1), uint(2)},
					Qualifications:  []uint{uint(1), uint(2)},
					Shift:           []uint{uint(1), uint(2)},
					Jobtype:         "Full-time",
				},
			},
			want:    models.ResponseJob{},
			wantErr: true,
			mockRepoResponse: func() (models.ResponseJob, error) {
				return models.ResponseJob{}, errors.New("error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRepo.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			s, err := NewService(mockRepo, &auth.Auth{}, &caching.MockCache{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestService_FilterApplications(t *testing.T) {
	type args struct {
		ctx            context.Context
		jobApplication []models.RespondJApplicant
	}
	tests := []struct {
		name              string
		args              args
		want              []models.RespondJApplicant
		wantErr           bool
		mockRepoResponse  func() (models.Job, error)
		mockCacheResponse func() (models.Job, error)
	}{
		{
			name: "Failure - Cache Miss",
			args: args{
				ctx: context.Background(),
				jobApplication: []models.RespondJApplicant{
					{
						Name: "Devops",
						Jid:  1,
						Jobs: models.UserApplicant{
							NoticePeriod:    "2",
							Experience:      "2",
							Location:        []uint{uint(1)},
							Qualifications:  []uint{uint(1)},
							Shift:           []uint{uint(1)},
							TechnologyStack: []uint{uint(1)},
							Jobtype:         "permanent",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockCacheResponse: func() (models.Job, error) {
				return models.Job{}, redis.Nil
			},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "Failure - Cache Error",
			args: args{
				ctx: context.Background(),
				jobApplication: []models.RespondJApplicant{
					{
						Name: "Devops",
						Jid:  1,
						Jobs: models.UserApplicant{
							NoticePeriod:    "2",
							Experience:      "2",
							Location:        []uint{uint(1)},
							Qualifications:  []uint{uint(1)},
							Shift:           []uint{uint(1)},
							TechnologyStack: []uint{uint(1)},
							Jobtype:         "permanent",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockCacheResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("cache error")
			},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCache := caching.NewMockCache(mc)
			mockCache.EXPECT().FetchCache(gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
			mockCache.EXPECT().AddCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test error")).AnyTimes()
			mockRepo := repository.NewMockUserRepo(mc)

			mockRepo.EXPECT().Viewjob(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			s, err := NewService(mockRepo, &auth.Auth{}, mockCache)
			if err != nil {
				t.Errorf("error initializing the repo")
				return
			}
			got, err := s.FilterApplications(tt.args.ctx, tt.args.jobApplication)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FilterApplications() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FilterApplications() = %v, want %v", got, tt.want)
			}
		})
	}
}
