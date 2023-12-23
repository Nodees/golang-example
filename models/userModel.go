package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	AddressID int
	Address   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
}

type Address struct {
	gorm.Model
	Cep          string
	Street       string
	Neighborhood string
	City         string
	State        string
}
