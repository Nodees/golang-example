package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name      string
	Username  string
	Email     *string
	Password  string `json:"password" binding:"required,min=8"`
	AddressID *int
	Address   *Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Address struct {
	BaseModel
	Cep          string
	Street       string
	Neighborhood string
	City         string
	State        string
}

func (u *User) BeforeSave(tx *gorm.DB) error {
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
