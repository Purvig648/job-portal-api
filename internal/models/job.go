package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}
type Job struct {
	gorm.Model
	Company Company `json:"-" gorm:"foreignKey:cid"`
	Cid     uint    `json:"cid"`
	JobRole string  `json:"job_role"`
	Salary  string  `json:"salary"`
}


