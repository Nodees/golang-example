package models

import (
	"core/utils"

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

type UserResponse struct {
	Name     string
	Username string
	Email    string
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    *user.Email,
	}
}

type Address struct {
	BaseModel
	Cep          string
	Street       string
	Neighborhood string
	City         string
	State        string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}
