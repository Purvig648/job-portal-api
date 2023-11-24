package models

type ForgetPass struct {
	Email       string `json:"email" gorm:"unique"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
}
