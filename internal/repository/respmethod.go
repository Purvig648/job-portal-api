package repository

import "job-application-api/internal/models"

func (r *Repo) CreateUser(userData models.User) (models.User, error) {
	result := r.DB.Create(&userData)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return userData, nil
}
