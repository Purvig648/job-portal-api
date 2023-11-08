package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company         Company           `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint              `json:"cid"`
	Jobname         string            `json:"jobname"`
	MinNoticePeriod string            `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod string            `json:"max_notice_period" validate:"required"`
	Description     string            `json:"description" validate:"required"`
	MinExperience   string            `json:"min_experience" validate:"required"`
	MaxExperience   string            `json:"max_experience" validate:"required"`
	Location        []Location        `json:"location" gorm:"many2many:job_location;"`
	TechnologyStack []TechnologyStack `json:"technology_stack" gorm:"many2many:job_techstack;"`
	Qualifications  []Qualification   `json:"qualifications" gorm:"many2many:job_qualification;"`
	Shift           []Shift           `json:"shift" gorm:"many2many:job_shift;" `
	Jobtype         string            `json:"jobtype" validate:"required"`
}
type ResponseJob struct {
	Id uint `json:"responseId"`
}

type Location struct {
	gorm.Model
	PlaceName string `json:"place_name"`
}

type TechnologyStack struct {
	gorm.Model
	StackName string `json:"stack_name"`
}

type Qualification struct {
	gorm.Model
	QualificationRequired string `json:"qualification_required"`
}

type Shift struct {
	gorm.Model
	ShiftType string `json:"shift_type"`
}

type NewJob struct {
	Company         Company           `json:"-" `
	Cid             uint              `json:"cid"`
	Jobname         string            `json:"jobname"`
	MinNoticePeriod string            `json:"min_notice_period"`
	MaxNoticePeriod string            `json:"max_notice_period"`
	Description     string            `json:"description"`
	MinExperience   string            `json:"min_experience"`
	MaxExperience   string            `json:"max_experience"`
	Location        []Location        `json:"location"`
	TechnologyStack []TechnologyStack `json:"technology_stack"`
	Qualifications  []Qualification   `json:"qualifications"`
	Shift           []Shift           `json:"shift"`
	Jobtype         string            `json:"jobtype"`
}
