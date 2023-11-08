package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

type UserSignup struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UserApplicant struct {
	// Name    string `json:"name" validate:"required"`
	// College string `json:"college_name" validate:"required"`
	// JobId   uint   `json:"job_id" validate:"required"`
	// Job     NewJob    `json:"job_app"`
	NoticePeriod    string `json:"notice_period"`
	Experience      string `json:"experience"`
	Location        []uint `json:"location"`
	TechnologyStack []uint `json:"technology_stack"`
	Qualifications  []uint `json:"qualifications"`
	Shift           []uint `json:"shift"`
	Jobtype         string `json:"jobtype"`
}
