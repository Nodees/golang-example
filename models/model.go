package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name      string `json:"name" validate:"required,min=8"`
	Username  string `json:"username" validate:"required,min=8"`
	Email     *string
	Password  string `json:"password" binding:"required,min=8"`
	AddressID *int
	Address   *Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Address struct {
	BaseModel
	Cep          *string `json:"cep"`
	Street       *string `json:"street"`
	Neighborhood *string `json:"neighborhood"`
	City         *string `json:"city"`
	State        *string `json:"state"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
