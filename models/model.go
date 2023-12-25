package models

type User struct {
	BaseModel
	Name      string
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
