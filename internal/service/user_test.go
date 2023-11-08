package service

// import (
// 	"context"
// 	"errors"
// 	"job-application-api/internal/auth"
// 	"job-application-api/internal/models"
// 	"job-application-api/internal/repository"
// 	"reflect"
// 	"testing"

// 	"go.uber.org/mock/gomock"
// )

// func TestService_UserSignup(t *testing.T) {
// 	type args struct {
// 		ctx      context.Context
// 		userData models.UserSignup
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             models.User
// 		wantErr          bool
// 		mockRepoResponse func() (models.User, error)
// 	}{
// 		{
// 			name: "error in db",
// 			args: args{
// 				ctx:      context.Background(),
// 				userData: models.UserSignup{},
// 			},
// 			want:    models.User{},
// 			wantErr: true,
// 			mockRepoResponse: func() (models.User, error) {
// 				return models.User{}, errors.New("test error")
// 			},
// 		}, {
// 			name: "Success case",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserSignup{
// 					Name:     "purvi",
// 					Email:    "purvi458@gmail.com",
// 					Password: "12345",
// 				},
// 			},
// 			want: models.User{
// 				Name:         "purvi",
// 				Email:        "purvi458@gmail.com",
// 				PasswordHash: "12345"},
// 			wantErr: false,
// 			mockRepoResponse: func() (models.User, error) {
// 				return models.User{
// 					Name:         "purvi",
// 					Email:        "purvi458@gmail.com",
// 					PasswordHash: "12345"}, nil
// 			},
// 		}, {
// 			name: "Error generating hash",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserSignup{
// 					Name:  "purvi",
// 					Email: "purvi458@gmail.com",
// 				},
// 			},
// 			want:    models.User{},
// 			wantErr: true,
// 			mockRepoResponse: func() (models.User, error) {
// 				return models.User{}, errors.New("test error")
// 			},
// 		},
// 	}

// 	for _, tt := range tests {

// 		mc := gomock.NewController(t)
// 		mockRepo := repository.NewMockUserRepo(mc)
// 		if tt.mockRepoResponse != nil {
// 			mockRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
// 		}
// 		s, _ := NewService(mockRepo, &auth.Auth{})
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := s.UserSignup(tt.args.ctx, tt.args.userData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_UserLogin(t *testing.T) {
// 	type args struct {
// 		ctx      context.Context
// 		userData models.UserLogin
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             string
// 		wantErr          bool
// 		mockRepoResponse func() (models.User, error)
// 	}{
// 		{
// 			name: "error in validating mail",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserLogin{
// 					Email:    "purvig648@gmail.com",
// 					Password: "purvi@123",
// 				},
// 			},
// 			want:    "",
// 			wantErr: true,
// 			mockRepoResponse: func() (models.User, error) {
// 				return models.User{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "error in comparing hashed password",
// 			args: args{
// 				ctx: context.Background(),
// 				userData: models.UserLogin{
// 					Email:    "purvi@gmail.com",
// 					Password: "12345678",
// 				},
// 			},
// 			want:    "",
// 			wantErr: true,
// 			mockRepoResponse: func() (models.User, error) {
// 				return models.User{
// 					Email:        "purvi@gmail.com",
// 					PasswordHash: "12345678",
// 				}, nil
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockRepoResponse != nil {
// 				mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockRepoResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockRepo, &auth.Auth{})

// 			got, err := s.UserLogin(tt.args.ctx, tt.args.userData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Service.UserLogin() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
