package service

import (
	"context"
	"errors"
	"fmt"
	"job-application-api/internal/models"
	"job-application-api/internal/pkg"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) VerifyUser(ctx context.Context, data models.ForgetPass) error {
	userDetails, err := s.UserRepo.CheckEmail(ctx, data.Email)
	if err != nil {
		return errors.New("email not found")
	}
	if userDetails.DOB != data.DateOfBirth {
		return errors.New("incorrect data of birth,can't send otp")
	}
	otp, err := pkg.GenerateAndSendOtp(data.Email)
	if err != nil {
		return errors.New("could not send otp")
	}
	err = s.rdb.AddToCacheOtp(ctx, data.Email, otp)
	if err != nil {
		fmt.Println("]][]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]]", err)
		return errors.New("could not cache")
	}
	return nil
}

func (s *Service) UserLogin(ctx context.Context, userData models.UserLogin) (string, error) {
	var userDetails models.User
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	err = pkg.CheckHashedPassword(userData.Password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return "", errors.New("entered password is not wrong")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.a.GenerateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) UserSignup(ctx context.Context, userData models.UserSignup) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		DOB:          userData.DOB,
		PasswordHash: hashedPass,
	}
	userDetails, err = s.UserRepo.CreateUser(userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil

}
