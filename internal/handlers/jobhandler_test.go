package handlers

import (
	"context"
	"errors"
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/models"
	"job-application-api/internal/service"
	"job-application-api/internal/service/mockmodels"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func Test_handler_AddJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "error in validating json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"cid": 1,
					"jobname": "Software Engineer",
					"minNoticePeriod": "30 days",
					"maxNoticePeriod": "60 days",
					"description": "This is a software engineer job description.",
					"minExperience": "2 years",
					"maxExperience": "5 years",
					"location": [1,2],
					"technologyStacks": [1,2],
					"qualifications": [1,2],
					"shifts": [1,2],
					"jobType": "Full-Time
					}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"provide valid details"}`,
		},
		{
			name: "Failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"cid": 1,
					"jobname": "Software Engineer",
					"minNoticePeriod": "30 days",
					"maxNoticePeriod": "60 days",
					"description": "This is a software engineer job description.",
					"minExperience": "2 years",
					"maxExperience": "5 years",
					"location": [1,2],
					"technologyStacks": [1,2],
					"qualifications": [1,2],
					"shifts": [1,2],
					"jobType": "Full-Time"
				}`))

				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any()).Return(models.ResponseJob{}, errors.New("test error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"test error"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
					"cid": 1,
					"jobname": "Software Engineer",
					"minNoticePeriod": "30 days",
					"maxNoticePeriod": "60 days",
					"description": "This is a software engineer job description.",
					"minExperience": "2 years",
					"maxExperience": "5 years",
					"location": [1,2],
					"technologyStacks": [1,2],
					"qualifications": [1,2],
					"shifts": [1,2],
					"jobType": "Full-Time"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)
				ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any()).Return(models.ResponseJob{
					Id: 1,
				}, nil)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":1}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := handler{
				service: ms,
			}
			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ProcessApplication(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "error in validating json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
					{
						"name": "Jeevan",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": "2",
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time"
						}
					},
					{
						"name": "purvi",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": 0,
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time
						}
					},
					{
						"name": "afthab",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": "3",
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time"
						}
					}
				]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"please provide all fields"}`,
		},
		{
			name: "Failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
						{
							"name": "Jeevan",
							"id": 1,
							"jobApplication": {
								"notice_period": "30",
								"experience": "2",
								"location": [
									1,
									2
								],
								"technology_stack": [
									1,
									2
								],
								"qualifications": [
									1,
									2
								],
								"shift": [
									1,
									2
								],
								"jobtype": "Part-time"
							}
						},
						{
							"name": "purvi",
							"id": 1,
							"jobApplication": {
								"notice_period": "30",
								"experience": "0",
								"location": [
									1,
									2
								],
								"technology_stack": [
									1,
									2
								],
								"qualifications": [
									1,
									2
								],
								"shift": [
									1,
									2
								],
								"jobtype": "Part-time"
							}
						},
						{
							"name": "afthab",
							"id": 1,
							"jobApplication": {
								"notice_period": "30",
								"experience": "3",
								"location": [
									1,
									2
								],
								"technology_stack": [
									1,
									2
								],
								"qualifications": [
									1,
									2
								],
								"shift": [
									1,
									2
								],
								"jobtype": "Part-time"
							}
						}
					]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				ms.EXPECT().FilterApplications(gomock.Any(), gomock.Any()).Return([]models.RespondJApplicant{}, errors.New("test error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"unable to filter records"}`,
		},
		{
			name: "Success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
					{
						"name": "Jeevan",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": "2",
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time"
						}
					},
					{
						"name": "purvi",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": "0",
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time"
						}
					},
					{
						"name": "afthab",
						"id": 1,
						"jobApplication": {
							"notice_period": "30",
							"experience": "3",
							"location": [
								1,
								2
							],
							"technology_stack": [
								1,
								2
							],
							"qualifications": [
								1,
								2
							],
							"shift": [
								1,
								2
							],
							"jobtype": "Part-time"
						}
					}
				]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdkey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := mockmodels.NewMockUserService(mc)

				ms.EXPECT().FilterApplications(gomock.Any(), gomock.Any()).Return([]models.RespondJApplicant{
					{
						Name: "afthab",
						Jid:  1,
						Jobs: models.UserApplicant{},
					},
				}, nil)
				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"name":"afthab","id":1,"jobApplication":{"notice_period":"","experience":"","location":null,"technology_stack":null,"qualifications":null,"shift":null,"jobtype":""}}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := handler{
				service: ms,
			}
			h.ProcessApplication(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
