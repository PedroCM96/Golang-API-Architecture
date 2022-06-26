package models

import (
	"gorm.io/gorm"
	"time"
)

type Dummy struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `validate:"required,alpha,min=2"`
	Email     string `gorm:"unique;not null" validate:"required,email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewDummy(name string, email string) (*Dummy, *Dummy) {
	dummy := Dummy{
		Name:  name,
		Email: email,
	}

	return &dummy, nil
}
