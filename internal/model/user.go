package model

import uuid "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	IsValidated bool      `gorm:"column:is_validated"`
	Password    string    `gorm:"column:password"`
}
