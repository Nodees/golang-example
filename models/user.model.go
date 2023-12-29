package models

import (
	"core/utils"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name      string   `gorm:"column:tx_name" validate:"required,min=8"`
	Username  string   `gorm:"column:tx_username" validate:"required,min=8" gorm:"unique"`
	Email     *string  `gorm:"column:tx_email"`
	Password  string   `gorm:"column:tx_password" binding:"required,min=8"`
	AddressID *int     `gorm:"column:id_address"`
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
