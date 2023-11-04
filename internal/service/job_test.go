package service

import (
	"context"
	"errors"
	"job-application-api/internal/auth"
	"job-application-api/internal/models"
	"job-application-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_ViewJobDetailsById(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error in db of ViewJobDetailsById",
			args: args{
				ctx: context.Background(),
				cid: 11,
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "success for ViewJobDetailsById",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want:    models.Job{},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().Viewjob(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobDetailsById(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobDetailsById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobDetailsById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewJobDetails(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db for ViewJobDetails",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("test error")
			},
		},
		{
			name: "success for ViewJobDetails",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: []models.Job{
				{
					Cid:     1,
					JobRole: "Developer",
					Salary:  "50000",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Cid:     1,
						JobRole: "Developer",
						Salary:  "50000",
					},
				}, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobByCid(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobDetails(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllJobPostings(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db for ViewAllJobPostings",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("test error")
			},
		},
		{
			name: "success for ViewAllJobPostings",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Job{
				{
					Cid:     1,
					JobRole: "Developer",
					Salary:  "20000",
				},
				{
					Cid:     1,
					JobRole: "Tester",
					Salary:  "23000",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Cid:     1,
						JobRole: "Developer",
						Salary:  "20000",
					},
					{
						Cid:     1,
						JobRole: "Tester",
						Salary:  "23000",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobPostings(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewAllJobPostings(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllJobPostings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllJobPostings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.Job
	}
	tests := []struct {
		name             string
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error in db for AddJobDetails",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "success for AddJobDetails",
			args: args{
				ctx: context.Background(),
				jobData: models.Job{
					Cid:     1,
					JobRole: "developer",
					Salary:  "200000",
				},
			},
			want: models.Job{
				Cid:     1,
				JobRole: "developer",
				Salary:  "200000",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Cid:     1,
					JobRole: "developer",
					Salary:  "200000",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(tt.args.ctx, tt.args.jobData).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
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
