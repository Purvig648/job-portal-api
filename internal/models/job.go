package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company Company `json:"-" gorm:"foreignKey:cid"`
	Cid     uint    `json:"cid"`
	JobRole string  `json:"job_role"`
	Salary  string  `json:"salary"`
}
