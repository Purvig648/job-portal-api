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

func TestService_AddCompanyDetails(t *testing.T) {
	type args struct {
		ctx         context.Context
		companyData models.Company
	}
	tests := []struct {
		name             string
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "error in db of AddCompanyDetails",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error")
			},
		},
		{
			name: "success for AddCompanyDetails",
			args: args{
				ctx: context.Background(),
				companyData: models.Company{
					Name:     "Teksystem",
					Location: "bangalore",
				},
			},
			want: models.Company{
				Name:     "Teksystem",
				Location: "bangalore",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Teksystem",
					Location: "bangalore",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		mc := gomock.NewController(t)
		mockRepo := repository.NewMockUserRepo(mc)
		if tt.mockRepoResponse != nil {
			mockRepo.EXPECT().CreateCompany(tt.args.ctx, tt.args.companyData).Return(tt.mockRepoResponse()).AnyTimes()
		}
		s, _ := NewService(mockRepo, &auth.Auth{})
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AddCompanyDetails(tt.args.ctx, tt.args.companyData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewCompanyDetails(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "error in db of ViewCompanyDetails",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("test error")
			},
		},
		{
			name: "success for ViewCompanyDetails",
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: models.Company{
				Name:     "Teksystem",
				Location: "bangalore",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "Teksystem",
					Location: "bangalore",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewCompany(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})

			got, err := s.ViewCompanyDetails(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewCompanyDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllCompanies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Company
		wantErr          bool
		mockRepoResponse func() ([]models.Company, error)
	}{
		{
			name: "error in db of ViewAllCompanies",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Company, error) {
				return nil, errors.New("test error")
			},
		},
		{
			name: "success for ViewAllCompanies",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Company{
				{
					Name:     "Teksystem",
					Location: "bangalore",
				},
				{
					Name:     "IBM",
					Location: "bangalore",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Company, error) {
				return []models.Company{
					{
						Name:     "Teksystem",
						Location: "bangalore",
					},
					{
						Name:     "IBM",
						Location: "bangalore",
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
				mockRepo.EXPECT().ViewAllCompanies(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})

			got, err := s.ViewAllCompanies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}
