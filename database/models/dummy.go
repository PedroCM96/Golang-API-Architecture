package models

import (
	"github.com/go-playground/validator/v10"
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

func NewDummy(name string, email string) (*Dummy, error) {
	d := Dummy{Name: name, Email: email}
	err := validator.New().Struct(d)

	if err != nil {
		return nil, err
	}

	return &d, nil
}
